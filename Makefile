ARTIFACT_ID=velero-plugin-for-restore-exclude
IMAGE=cloudogu/${ARTIFACT_ID}:${VERSION}
GOTAG?=1.24
MAKEFILES_VERSION=9.9.1
MOCKERY_VERSION=v2.53.3
MOCKERY_IGNORED=vendor,build,docs,generated

VERSION=1.0.1

GOOS   ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)


include build/make/variables.mk
include build/make/self-update.mk
include build/make/dependencies-gomod.mk
include build/make/build.mk
include build/make/test-common.mk
include build/make/test-unit.mk
include build/make/static-analysis.mk
include build/make/clean.mk
include build/make/digital-signature.mk
include build/make/mocks.mk
include build/make/k8s.mk
include build/make/release.mk
