BIN_DIR := bin
TOOLS_DIR := tools

include provider-version.mk

MANIFESTS_GEN := $(TOOLS_DIR)/bin/manifests-gen
KUSTOMIZE := $(TOOLS_DIR)/bin/kustomize

$(TOOLS_DIR)/bin/%:
	$(MAKE) -C $(TOOLS_DIR) bin/$*

.PHONY: verify
verify: verify-ocp-manifests

.PHONY: verify-%
verify-%:
	$(MAKE) $*
	./verify-diff.sh

.PHONY: update-manifests-gen
update-manifests-gen:
	cd tools && go get github.com/openshift/cluster-capi-operator/manifests-gen@main github.com/openshift/cluster-capi-operator@main && go mod tidy && go mod vendor

.PHONY: remediation-crds
remediation-crds: $(KUSTOMIZE) ## Generates remediation CRDs for CVO deployment into openshift/manifests/
	$(KUSTOMIZE) build --load-restrictor LoadRestrictionsNone remediation-crds/metal3remediation -o manifests/0000_30_machine-api-operator_07_metal3remediation.crd.yaml
	$(KUSTOMIZE) build --load-restrictor LoadRestrictionsNone remediation-crds/metal3remediationtemplate -o manifests/0000_30_machine-api-operator_07_metal3remediationtemplate.crd.yaml

.PHONY: ocp-manifests
ocp-manifests: remediation-crds $(MANIFESTS_GEN) ## Builds openshift specific manifests
	$(MANIFESTS_GEN) \
		--manifests-path "capi-operator-manifests" \
		--kustomize-dir "../openshift" \
		--name cluster-api-provider-metal3 \
		--install-order 20 \
		--attribute type=infrastructure \
		--attribute version="${PROVIDER_VERSION}" \
		--self-image-ref registry.ci.openshift.org/openshift:baremetal-cluster-api-controllers \
		--platform BareMetal \
		--protect-cluster-resource metal3cluster
