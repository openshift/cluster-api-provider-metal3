BIN_DIR := bin

MANIFESTS_GEN := bin/manifests-gen
KUSTOMIZE := bin/kustomize

$(MANIFESTS_GEN): pkg = github.com/openshift/cluster-capi-operator/manifests-gen
$(KUSTOMIZE): pkg = sigs.k8s.io/kustomize/kustomize/v5

.PHONY: default
default: $(MANIFESTS_GEN) $(KUSTOMIZE)

$(BIN_DIR):
	mkdir -p $(BIN_DIR)

$(BIN_DIR)/%: _FORCE | $(BIN_DIR)
	go build -o $@ $(pkg)

.PHONY: _FORCE
