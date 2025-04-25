/*
Copyright the Velero contributors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package plugin

import (
	"context"
	"fmt"
	"gopkg.in/yaml.v3"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	v1 "github.com/vmware-tanzu/velero/pkg/apis/velero/v1"
	"github.com/vmware-tanzu/velero/pkg/plugin/velero"
	riav2 "github.com/vmware-tanzu/velero/pkg/plugin/velero/restoreitemaction/v2"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	// If this annotation is found on the Velero Restore CR, then create an operation
	// that is considered done at backup start time + example RIA operation duration
	// If this annotation is not present, then operationID returned from Execute() will
	// be empty.
	// This annotation can also be set on the item, which overrides the restore CR value,
	// to allow for testing multiple action lengths
	AsyncRIADurationAnnotation = "velero.io/example-ria-operation-duration"
)

// RestorePlugin is a restore item action plugin for Velero
type RestorePluginV2 struct {
	log logrus.FieldLogger
}

type ExcludeEntry struct {
	Group   string `yaml:"group"`
	Version string `yaml:"version"`
	Kind    string `yaml:"kind"`
	Name    string `yaml:"name"`
}

type groupVersionKindName struct {
	Gvk  schema.GroupVersionKind
	Name string
}

// NewRestorePluginV2 instantiates a v2 RestorePlugin.
func NewRestorePluginV2(log logrus.FieldLogger) *RestorePluginV2 {
	return &RestorePluginV2{log: log}
}

// Name is required to implement the interface, but the Velero pod does not delegate this
// method -- it's used to tell velero what name it was registered under. The plugin implementation
// must define it, but it will never actually be called.
func (p *RestorePluginV2) Name() string {
	return "exampleRestorePlugin"
}

// AppliesTo returns information about which resources this action should be invoked for.
// The IncludedResources and ExcludedResources slices can include both resources
// and resources with group names. These work: "ingresses", "ingresses.extensions".
// A RestoreItemAction's Execute function will only be invoked on items that match the returned
// selector. A zero-valued ResourceSelector matches all resources.
func (p *RestorePluginV2) AppliesTo() (velero.ResourceSelector, error) {
	return velero.ResourceSelector{}, nil
}

// Execute allows the RestorePlugin to perform arbitrary logic with the item being restored,
// in this case, setting a custom annotation on the item being restored.
func (p *RestorePluginV2) Execute(input *velero.RestoreItemActionExecuteInput) (*velero.RestoreItemActionExecuteOutput, error) {
	p.log.Info("Hello from my RestorePlugin(v2)!")

	p.log.Info("-----Parse element to Unstructured-----")
	itemUnstructured, ok := input.Item.(*unstructured.Unstructured)
	if !ok {
		p.log.Errorf("failed to parse element")
	} else {
		p.log.Infof("parsed element: %q", itemUnstructured)
	}

	gvkn := groupVersionKindName{
		Gvk:  itemUnstructured.GroupVersionKind(),
		Name: itemUnstructured.GetName(),
	}
	p.log.Infof("gvkn: %s", gvkn)

	p.log.Info("-----Create kubernetes client-----")
	config, err := rest.InClusterConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to get cluster config: %w", err)
	}
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, fmt.Errorf("failed to create kubernetes client: %w", err)
	}

	p.log.Info("-----Read ConfigMap-----")
	configMap, err := clientSet.CoreV1().ConfigMaps("ecosystem").Get(context.TODO(), "velero-plugin-for-restore-exclude-config", metaV1.GetOptions{})
	if err != nil {
		return &velero.RestoreItemActionExecuteOutput{
			UpdatedItem: input.Item,
		}, fmt.Errorf("failed to get configmap: %w", err)
	}

	shouldBeExcludedString := configMap.Data["restore"]
	var exclude struct {
		Exclude []ExcludeEntry `yaml:"exclude"`
	}

	p.log.Info("-----Unmarshal configmap-----")
	err = yaml.Unmarshal([]byte(shouldBeExcludedString), &exclude)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}
	var shouldBeExcluded []groupVersionKindName
	for _, entry := range exclude.Exclude {
		gvk := schema.GroupVersionKind{
			Group:   entry.Group,
			Version: entry.Version,
			Kind:    entry.Kind,
		}
		shouldBeExcluded = append(shouldBeExcluded, groupVersionKindName{
			Gvk:  gvk,
			Name: entry.Name,
		})
	}

	p.log.Info("-----Check if element is excluded-----")
	for _, excludedElement := range shouldBeExcluded {
		if excludedElement.matches(gvkn) {
			p.log.Info("-----Exclude element-----")
			p.log.Info(gvkn.Name)
			return &velero.RestoreItemActionExecuteOutput{SkipRestore: true}, nil
		}
	}
	p.log.Info("-----Return result-----")
	return &velero.RestoreItemActionExecuteOutput{
		UpdatedItem: input.Item,
	}, nil
}

func (p *RestorePluginV2) Progress(operationID string, restore *v1.Restore) (velero.OperationProgress, error) {
	progress := velero.OperationProgress{}
	if operationID == "" {
		return progress, riav2.InvalidOperationIDError(operationID)
	}
	splitOp := strings.Split(operationID, "/")
	if len(splitOp) != 2 {
		return progress, riav2.InvalidOperationIDError(operationID)
	}
	duration, err := time.ParseDuration(splitOp[1])
	if err != nil {
		return progress, riav2.InvalidOperationIDError(operationID)
	}
	elapsed := time.Since(restore.Status.StartTimestamp.Time).Seconds()
	if elapsed >= duration.Seconds() {
		progress.Completed = true
		progress.NCompleted = int64(duration.Seconds())
	} else {
		progress.NCompleted = int64(elapsed)
	}
	progress.NTotal = int64(duration.Seconds())
	progress.OperationUnits = "seconds"
	progress.Updated = time.Now()

	return progress, nil
}

func (p *RestorePluginV2) Cancel(operationID string, restore *v1.Restore) error {
	return nil
}

func (p *RestorePluginV2) AreAdditionalItemsReady(additionalItems []velero.ResourceIdentifier, restore *v1.Restore) (bool, error) {
	return true, nil
}

func (g groupVersionKindName) matches(gvkn groupVersionKindName) bool {
	return (gvkn.Gvk.Group == g.Gvk.Group || g.Gvk.Group == "*" || g.Gvk.Group == "") &&
		(gvkn.Gvk.Version == g.Gvk.Version || g.Gvk.Version == "*" || g.Gvk.Version == "") &&
		(gvkn.Gvk.Kind == g.Gvk.Kind || g.Gvk.Kind == "*" || g.Gvk.Kind == "") &&
		(gvkn.Name == g.Name || g.Name == "*" || g.Name == "")
}

func groupVersionKind(resource unstructured.Unstructured) groupVersionKindName {
	return groupVersionKindName{
		Gvk:  resource.GroupVersionKind(),
		Name: resource.GetName(),
	}
}

func parseValue(input map[string]interface{}, key string) (string, map[string]interface{}, error) {
	val, exists := input[key]
	if exists {
		return "", nil, fmt.Errorf("key %q not found", key)
	}

	switch v := val.(type) {
	case string:
		return v, nil, nil
	case map[string]interface{}:
		return "", v, nil
	default:
		return "", nil, fmt.Errorf("unexpected type %T", val)
	}
}
