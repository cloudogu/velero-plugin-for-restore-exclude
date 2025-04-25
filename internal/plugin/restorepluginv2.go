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
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	v1 "github.com/vmware-tanzu/velero/pkg/apis/velero/v1"
	"github.com/vmware-tanzu/velero/pkg/plugin/velero"
	riav2 "github.com/vmware-tanzu/velero/pkg/plugin/velero/restoreitemaction/v2"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	Namespace     = "ecosystem"
	ConfigMapName = "velero-plugin-for-restore-exclude-config"
)

// RestorePlugin is a restore item action plugin for Velero
type RestorePluginV2 struct {
	log       logrus.FieldLogger
	clientset configMapInterface
}

type ExcludeEntry struct {
	Group   string `yaml:"group"`
	Version string `yaml:"version"`
	Kind    string `yaml:"kind"`
	Name    string `yaml:"name"`
}

func (e ExcludeEntry) matches(item object) bool {
	return (item.GroupVersionKind().Group == e.Group || e.Group == "*") &&
		(item.GroupVersionKind().Version == e.Version || e.Version == "*") &&
		(item.GroupVersionKind().Kind == e.Kind || e.Kind == "*") &&
		(item.GetName() == e.Name || e.Name == "*")
}

type object interface {
	metaV1.Object
	schema.ObjectKind
}

// NewRestorePluginV2 instantiates a v2 RestorePlugin.
func NewRestorePluginV2(log logrus.FieldLogger, clientset configMapInterface) *RestorePluginV2 {
	return &RestorePluginV2{
		log:       log,
		clientset: clientset,
	}
}

// Name is required to implement the interface, but the Velero pod does not delegate this
// method -- it's used to tell velero what name it was registered under. The plugin implementation
// must define it, but it will never actually be called.
func (p *RestorePluginV2) Name() string {
	return "velero-plugin-for-restore-exclude"
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
	itemUnstructured, ok := input.Item.(*unstructured.Unstructured)
	if !ok {
		return nil, fmt.Errorf("failed to parse input")
	}

	gvkn := groupVersionKindName{
		Gvk:  itemUnstructured.GroupVersionKind(),
		Name: itemUnstructured.GetName(),
	}

	configMap, err := p.clientset.Get(context.Background(), ConfigMapName, metaV1.GetOptions{})
	if err != nil {
		return &velero.RestoreItemActionExecuteOutput{
			UpdatedItem: input.Item,
		}, fmt.Errorf("failed to get configmap: %w", err)
	}

	shouldBeExcludedString := configMap.Data["restore"]
	var exclude struct {
		Exclude []ExcludeEntry `yaml:"exclude"`
	}

	err = yaml.Unmarshal([]byte(shouldBeExcludedString), &exclude)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}
	for _, excludedElement := range exclude.Exclude {
		if excludedElement.matches(itemUnstructured) {
			return &velero.RestoreItemActionExecuteOutput{SkipRestore: true}, nil
		}
	}

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

