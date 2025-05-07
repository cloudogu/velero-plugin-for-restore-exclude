/*
This file is based on code originally licensed unter the Apache License, version 2.0.
It has been modified by Cloudogu GmbH and is distributed under the AGPL-3.0-only as part of the velero-plugin-for-restore-exclude Project.
Original code Copyright 2017, 2019 the Velero contributors.
Modification Copyright 2025 - present, Cloudogu GmbH
*/

package main

import (
	"fmt"
	"github.com/cloudogu/velero-plugin-for-restore-exclude/internal/plugin"
	"github.com/sirupsen/logrus"
	"github.com/vmware-tanzu/velero/pkg/plugin/framework"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func main() {
	framework.NewServer().
		RegisterRestoreItemActionV2("k8s.cloudogu.com/velero-plugin-for-restore-exclude", newRestorePluginV2).
		Serve()
}

func newRestorePluginV2(logger logrus.FieldLogger) (interface{}, error) {
	config, err := rest.InClusterConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to get cluster config: %w", err)
	}
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, fmt.Errorf("failed to create kubernetes client: %w", err)
	}
	return plugin.NewRestorePluginV2(logger, clientSet.CoreV1().ConfigMaps("ecosystem")), nil
}
