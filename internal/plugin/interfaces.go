package plugin

import v1 "k8s.io/client-go/kubernetes/typed/core/v1"

type coreV1Interface interface {
	v1.CoreV1Interface
}
