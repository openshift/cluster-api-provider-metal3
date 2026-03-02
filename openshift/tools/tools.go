//go:build tools
// +build tools
package tools

import (
	_ "github.com/openshift/cluster-capi-operator/manifests-gen"
	_ "sigs.k8s.io/kustomize/kustomize/v5"
)
