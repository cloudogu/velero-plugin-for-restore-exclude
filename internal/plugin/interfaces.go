package plugin

import v1 "k8s.io/client-go/kubernetes/typed/core/v1"

type configMapInterface interface {
	v1.ConfigMapInterface
}
