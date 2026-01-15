BIN_DIR := bin
TOOLS_DIR := tools

include provider-version.mk

$(RELEASE_DIR):
	mkdir -p $(RELEASE_DIR)/

KUSTOMIZE := tools/bin/kustomize
MANIFESTS_GEN := go run ./vendor/github.com/openshift/cluster-capi-operator/manifests-gen/

.PHONY: verify
verify: verify-ocp-manifests

.PHONY: verify-%
verify-%: ## Ensure no diff after running some other target
	$(MAKE) $*
	./verify-diff.sh

$(KUSTOMIZE):
	./tools/ensure-kustomize.sh

.PHONY: ocp-manifests
ocp-manifests: $(RELEASE_DIR) $(KUSTOMIZE) ## Builds openshift specific manifests
	# Build core-components.
	$(KUSTOMIZE) build ../config/default > infrastructure-components.yaml
	# Generate provider manifests.
	# TODO: load the provider-version dynamically at rebase time when this is invoked by the Rebase Bot during one of its lifecycle hooks.
	cd tools && $(MANIFESTS_GEN) --provider-name "metal3" --provider-type "InfrastructureProvider" --provider-version "${PROVIDER_VERSION}" --base-path "../../" --manifests-path "../manifests"
