package plugin

import (
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"testing"
)

func Test_groupVersionKindName_matches(t *testing.T) {
	type fields struct {
		Gvk  schema.GroupVersionKind
		Name string
	}
	tests := []struct {
		name   string
		fields fields
		gvkn   groupVersionKindName
		want   bool
	}{
		{
			name: "gvkns are identical",
			fields: fields{
				Name: "loadbalancer",
				Gvk: schema.GroupVersionKind{
					Kind:    "Service",
					Version: "v1",
					Group:   "Test",
				},
			},
			gvkn: groupVersionKindName{
				Name: "loadbalancer",
				Gvk: schema.GroupVersionKind{
					Kind:    "Service",
					Version: "v1",
					Group:   "Test",
				},
			},
			want: true,
		},
		{
			name: "gvkns do not match",
			fields: fields{
				Name: "loadbalancer",
				Gvk: schema.GroupVersionKind{
					Kind:    "Service",
					Version: "v1",
					Group:   "Test",
				},
			},
			gvkn: groupVersionKindName{
				Name: "loadbalancer",
				Gvk: schema.GroupVersionKind{
					Kind:    "Service",
					Version: "v2",
					Group:   "Test",
				},
			},
			want: false,
		},

		{
			name: "gvkn matches wildcard",
			fields: fields{
				Name: "loadbalancer",
				Gvk: schema.GroupVersionKind{
					Kind:    "Service",
					Version: "v1",
					Group:   "Test",
				},
			},
			gvkn: groupVersionKindName{
				Name: "*",
				Gvk: schema.GroupVersionKind{
					Kind:    "*",
					Version: "*",
					Group:   "*",
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := groupVersionKindName{
				Gvk:  tt.fields.Gvk,
				Name: tt.fields.Name,
			}
			if got := g.matches(tt.gvkn); got != tt.want {
				t.Errorf("matches() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewRestorePluginV2(t *testing.T) {
	t.Run("Create new restore plugin", func(t *testing.T) {
		log := logrus.Logger{}

		plugin := NewRestorePluginV2(&log)

		assert.NotNil(t, plugin)
	})
}

func TestAppliesTo(t *testing.T) {
	t.Run("Create new restore plugin", func(t *testing.T) {
		log := logrus.Logger{}

		sut := NewRestorePluginV2(&log)

		rs, err := sut.AppliesTo()

		assert.NoError(t, err)
		assert.NotNil(t, rs)
	})
}
