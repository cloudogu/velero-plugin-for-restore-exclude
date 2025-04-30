package plugin

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	v12 "github.com/vmware-tanzu/velero/pkg/apis/velero/v1"
	"github.com/vmware-tanzu/velero/pkg/plugin/velero"

	v1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

type sampleObject struct {
	name string
	gvk  schema.GroupVersionKind
}

func (s sampleObject) GetName() string {
	return s.name
}

func (s sampleObject) GroupVersionKind() schema.GroupVersionKind {
	return s.gvk
}

func Test_ExcludeEntry_matches(t *testing.T) {
	tests := []struct {
		name         string
		excludeEntry ExcludeEntry
		object       object
		want         bool
	}{
		{
			name: "gvkns are identical",
			excludeEntry: ExcludeEntry{
				Name:    "loadbalancer",
				Kind:    "Service",
				Version: "v1",
				Group:   "Test",
			},
			object: sampleObject{
				name: "loadbalancer",
				gvk: schema.GroupVersionKind{
					Kind:    "Service",
					Version: "v1",
					Group:   "Test",
				},
			},
			want: true,
		},
		{
			name: "gvkns do not match",
			excludeEntry: ExcludeEntry{
				Name:    "loadbalancer",
				Kind:    "Service",
				Version: "v1",
				Group:   "Test",
			},
			object: sampleObject{
				name: "loadbalancer",
				gvk: schema.GroupVersionKind{
					Kind:    "Service",
					Version: "v2",
					Group:   "Test",
				},
			},
			want: false,
		},

		{
			name: "object matches wildcard",
			excludeEntry: ExcludeEntry{
				Name:    "*",
				Kind:    "*",
				Version: "*",
				Group:   "*",
			},
			object: sampleObject{
				name: "loadbalancer",
				gvk: schema.GroupVersionKind{
					Kind:    "Service",
					Version: "v1",
					Group:   "Test",
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.excludeEntry.matches(tt.object); got != tt.want {
				t.Errorf("matches() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewRestorePluginV2(t *testing.T) {
	t.Run("Create new restore plugin", func(t *testing.T) {
		log := logrus.Logger{}
		clientMock := newMockConfigMapInterface(t)

		plugin := NewRestorePluginV2(&log, clientMock)

		assert.NotNil(t, plugin)
	})
}

func TestAppliesTo(t *testing.T) {
	t.Run("Create new restore plugin", func(t *testing.T) {
		log := logrus.Logger{}
		clientMock := newMockConfigMapInterface(t)

		sut := NewRestorePluginV2(&log, clientMock)

		rs, err := sut.AppliesTo()

		assert.NoError(t, err)
		assert.NotNil(t, rs)
	})
}

func TestRestorePluginV2_Execute(t *testing.T) {
	tests := []struct {
		name        string
		clientSetFn func(t *testing.T) configMapInterface
		input       *velero.RestoreItemActionExecuteInput
		want        *velero.RestoreItemActionExecuteOutput
		wantErr     assert.ErrorAssertionFunc
	}{
		{
			name: "should fail to parse input",
			clientSetFn: func(t *testing.T) configMapInterface {
				return newMockConfigMapInterface(t)
			},
			input: &velero.RestoreItemActionExecuteInput{
				Item: nil,
			},
			want: nil,
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.ErrorContains(t, err, "failed to parse input")
			},
		},
		{
			name: "should fail to get config map",
			clientSetFn: func(t *testing.T) configMapInterface {
				clientSet := newMockConfigMapInterface(t)
				clientSet.EXPECT().Get(context.Background(), "velero-plugin-for-restore-exclude-config", metaV1.GetOptions{}).Return(nil, assert.AnError)
				return clientSet
			},
			input: &velero.RestoreItemActionExecuteInput{
				Item: &unstructured.Unstructured{},
			},
			want: &velero.RestoreItemActionExecuteOutput{
				UpdatedItem: &unstructured.Unstructured{},
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.ErrorContains(t, err, "failed to get configmap: ")
			},
		},
		{
			name: "should fail to unmarshal config",
			clientSetFn: func(t *testing.T) configMapInterface {
				clientSet := newMockConfigMapInterface(t)
				data := make(map[string]string)
				data["restore"] = "|\n" +
					"exclude:" +
					"  - name: :::::"
				clientSet.EXPECT().Get(context.Background(), "velero-plugin-for-restore-exclude-config", metaV1.GetOptions{}).Return(&v1.ConfigMap{Data: data}, nil)
				return clientSet
			},
			input: &velero.RestoreItemActionExecuteInput{
				Item: &unstructured.Unstructured{},
			},
			want: nil,
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.ErrorContains(t, err, "failed to unmarshal config: ")
			},
		},
		{
			name: "exclude element",
			clientSetFn: func(t *testing.T) configMapInterface {
				clientSet := newMockConfigMapInterface(t)
				data := make(map[string]string)
				data["restore"] = "exclude:\n  - kind: 'Service'"
				clientSet.EXPECT().Get(mock.Anything, "velero-plugin-for-restore-exclude-config", mock.Anything).Return(&v1.ConfigMap{Data: data}, nil)
				return clientSet
			},
			input: &velero.RestoreItemActionExecuteInput{
				Item: &unstructured.Unstructured{
					map[string]interface{}{"kind": "Service"},
				},
			},
			want:    &velero.RestoreItemActionExecuteOutput{SkipRestore: true},
			wantErr: assert.NoError,
		},
		{
			name: "don't exclude element",
			clientSetFn: func(t *testing.T) configMapInterface {
				clientSet := newMockConfigMapInterface(t)
				data := make(map[string]string)
				data["restore"] = "exclude:\n  - name: 'ces-loadbalancer'\n    kind: 'Service'"
				clientSet.EXPECT().Get(context.Background(), "velero-plugin-for-restore-exclude-config", metaV1.GetOptions{}).Return(&v1.ConfigMap{Data: data}, nil)
				return clientSet
			},
			input: &velero.RestoreItemActionExecuteInput{
				Item: &unstructured.Unstructured{
					map[string]interface{}{"name": "ces-loadbalancer"},
				},
			},
			want: &velero.RestoreItemActionExecuteOutput{SkipRestore: false, UpdatedItem: &unstructured.Unstructured{
				map[string]interface{}{"name": "ces-loadbalancer"},
			}},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &RestorePluginV2{
				log:       &logrus.Logger{},
				clientset: tt.clientSetFn(t),
			}
			got, err := p.Execute(tt.input)
			if !tt.wantErr(t, err, fmt.Sprintf("Execute(%v)", tt.input)) {
				return
			}
			assert.Equalf(t, tt.want, got, "Execute(%v)", tt.input)
		})
	}
}

func TestRestorePluginV2_Progress(t *testing.T) {
	type args struct {
		operationID string
		restore     *v12.Restore
	}
	tests := []struct {
		name    string
		args    args
		want    velero.OperationProgress
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "should fail because operation ID empty",
			args: args{
				operationID: "",
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.ErrorContains(t, err, "Operation ID  is invalid.")
			},
		},
		{
			name: "should fail because operation ID contains multiple /",
			args: args{
				operationID: "2/2/2",
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.ErrorContains(t, err, "Operation ID 2/2/2 is invalid.")
			},
		},
		{
			name: "should fail because operation ID contains /",
			args: args{
				operationID: "2/2",
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.ErrorContains(t, err, "Operation ID 2/2 is invalid.")
			},
		},
		{
			name: "should fail because operation ID contains /",
			args: args{
				operationID: "2/300ms",
				restore: &v12.Restore{
					Status: v12.RestoreStatus{
						StartTimestamp: &metaV1.Time{
							Time: time.Now(),
						},
					},
				},
			},
			want: velero.OperationProgress{
				Completed:      false,
				OperationUnits: "seconds",
			},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &RestorePluginV2{
				log:       &logrus.Logger{},
				clientset: newMockConfigMapInterface(t),
			}
			got, err := p.Progress(tt.args.operationID, tt.args.restore)
			if !tt.wantErr(t, err, fmt.Sprintf("Progress(%v, %v)", tt.args.operationID, tt.args.restore)) {
				return
			}
			assert.NotNil(t, got)
		})
	}
}
