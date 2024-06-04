# https://github.com/aperturerobotics/template

# Projects can override PROJECT_DIR with the path to their project.
COMMON_DIR = $(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))
PROJECT_DIR := $(COMMON_DIR)
PROJECT_DIR_REL = $(shell realpath --relative-to $(COMMON_DIR) $(PROJECT_DIR))

TOOLS_DIR := .tools
TOOLS_BIN := $(TOOLS_DIR)/bin

PROJECT_TOOLS_DIR := $(PROJECT_DIR)/$(TOOLS_DIR)
PROJECT_TOOLS_DIR_REL = $(shell realpath --relative-to $(COMMON_DIR) $(PROJECT_TOOLS_DIR))

SHELL:=bash
MAKEFLAGS += --no-print-directory

export GO111MODULE=on
undefine GOARCH
undefine GOOS

.PHONY: all
all: protodeps

# Setup node_modules
$(PROJECT_DIR)/node_modules:
	cd $(PROJECT_DIR_REL); yarn install

# Setup the .tools directory to hold the build tools in the project repo
$(PROJECT_TOOLS_DIR):
	@cd $(PROJECT_DIR); go run -v github.com/aperturerobotics/common $(TOOLS_DIR)

.PHONY: tools
tools: $(PROJECT_TOOLS_DIR)

# Build tool rule
define build_tool
$(PROJECT_DIR)/$(1): $(PROJECT_TOOLS_DIR)
	cd $(PROJECT_TOOLS_DIR_REL); \
	go build -mod=readonly -v \
		-o ./bin/$(shell basename $(1)) \
		$(2)

.PHONY: $(1)
$(1): $(PROJECT_DIR)/$(1)
endef

# List of available Go tool binaries
PROTOWRAP=$(TOOLS_BIN)/protowrap
PROTOC_GEN_GO=$(TOOLS_BIN)/protoc-gen-go-lite
PROTOC_GEN_GO_STARPC=$(TOOLS_BIN)/protoc-gen-go-starpc
GOIMPORTS=$(TOOLS_BIN)/goimports
GOFUMPT=$(TOOLS_BIN)/gofumpt
GOLANGCI_LINT=$(TOOLS_BIN)/golangci-lint
GO_MOD_OUTDATED=$(TOOLS_BIN)/go-mod-outdated
GORELEASER=$(TOOLS_BIN)/goreleaser
WASMBROWSERTEST=$(TOOLS_BIN)/wasmbrowsertest

# Mappings for build tool to Go import path
$(eval $(call build_tool,$(PROTOC_GEN_GO),github.com/aperturerobotics/protobuf-go-lite/cmd/protoc-gen-go-lite))
$(eval $(call build_tool,$(PROTOC_GEN_GO_STARPC),github.com/aperturerobotics/starpc/cmd/protoc-gen-go-starpc))
$(eval $(call build_tool,$(GOIMPORTS),golang.org/x/tools/cmd/goimports))
$(eval $(call build_tool,$(GOFUMPT),mvdan.cc/gofumpt))
$(eval $(call build_tool,$(PROTOWRAP),github.com/aperturerobotics/goprotowrap/cmd/protowrap))
$(eval $(call build_tool,$(GOLANGCI_LINT),github.com/golangci/golangci-lint/cmd/golangci-lint))
$(eval $(call build_tool,$(GO_MOD_OUTDATED),github.com/psampaz/go-mod-outdated))
$(eval $(call build_tool,$(GORELEASER),github.com/goreleaser/goreleaser))
$(eval $(call build_tool,$(WASMBROWSERTEST),github.com/agnivade/wasmbrowsertest))

.PHONY: protodeps
protodeps: $(GOIMPORTS) $(PROTOWRAP) $(PROTOC_GEN_GO) $(PROTOC_GEN_GO_STARPC) $(PROJECT_DIR)/node_modules

# Default protogen targets and arguments
PROTOGEN_TARGETS ?= ./*.proto
PROTOGEN_ARGS ?=
GO_LITE_OPT_FEATURES ?= marshal+unmarshal+size+equal+json+clone+text

.PHONY: genproto
genproto: protodeps
	@shopt -s globstar; \
	set -eo pipefail; \
	cd $(PROJECT_DIR); \
	export PROTOBUF_GO_TYPES_PKG=github.com/aperturerobotics/protobuf-go-lite/types; \
	export PROJECT=$$(go list -m); \
	export OUT=./vendor; \
	mkdir -p $${OUT}/$$(dirname $${PROJECT}); \
	rm -f ./vendor/$${PROJECT}; \
	ln -s $$(pwd) ./vendor/$${PROJECT}; \
	protogen() { \
		PROTO_FILES=$$(git ls-files "$$1"); \
		FMT_GO_FILES=(); \
		FMT_TS_FILES=(); \
		PROTOWRAP_OPTS=(); \
		if [ -f "go.mod" ]; then \
			PROTOWRAP_OPTS+=( \
				--plugin=$(PROTOC_GEN_GO) \
				--plugin=$(PROTOC_GEN_GO_STARPC) \
				--go-lite_out=$${OUT} \
				--go-lite_opt=features=$(GO_LITE_OPT_FEATURES) \
				--go-starpc_out=$${OUT} \
			); \
		fi; \
		if [ -f "package.json" ]; then \
			PROTOWRAP_OPTS+=( \
				--plugin=./node_modules/.bin/protoc-gen-es \
				--plugin=./node_modules/.bin/protoc-gen-es-starpc \
				--es-lite_out=$${OUT} \
				--es-lite_opt target=ts \
				--es-lite_opt ts_nocheck=false \
				--es-starpc_out=$${OUT} \
				--es-starpc_opt target=ts \
				--es-starpc_opt ts_nocheck=false \
			); \
		fi; \
		$(PROTOWRAP) \
			-I $${OUT} \
			"$${PROTOWRAP_OPTS[@]}" \
			--proto_path $${OUT} \
			--print_structure \
			--only_specified_files \
			$(PROTOGEN_ARGS) \
			$$(echo "$$PROTO_FILES" | xargs printf -- "./vendor/$${PROJECT}/%s "); \
		for proto_file in $${PROTO_FILES}; do \
			proto_dir=$$(dirname $$proto_file); \
			proto_name=$${proto_file%".proto"}; \
			GO_FILES=$$(git ls-files ":(glob)$${proto_dir}/${proto_name}*.pb.go"); \
			if [ -n "$$GO_FILES" ]; then FMT_GO_FILES+=($${GO_FILES[@]}); fi; \
			TS_FILES=$$(git ls-files ":(glob)$${proto_dir}/${proto_name}*.pb.ts"); \
			if [ -n "$$TS_FILES" ]; then FMT_TS_FILES+=($${TS_FILES[@]}); fi; \
			if [ -z "$$TS_FILES" ]; then continue; fi; \
			for ts_file in $${TS_FILES}; do \
				ts_file_dir=$$(dirname $$ts_file); \
				relative_path=$${ts_file_dir#"./"}; \
				depth=$$(echo $$relative_path | awk -F/ '{print NF+1}'); \
				prefix=$$(printf '../%0.s' $$(seq 1 $$depth)); \
				istmts=$$(grep -oE "from\s+\"$$prefix[^\"]+\"" $$ts_file) || continue; \
				if [ -z "$$istmts" ]; then continue; fi; \
				ipaths=$$(echo "$$istmts" | awk -F'"' '{print $$2}'); \
				for import_path in $$ipaths; do \
					rel_import_path=$$(realpath -s --relative-to=./vendor \
						"./vendor/$${PROJECT}/$${ts_file_dir}/$${import_path}"); \
					go_import_path=$$(echo $$rel_import_path | sed -e "s|^|@go/|"); \
					sed -i -e "s|$$import_path|$$go_import_path|g" $$ts_file; \
				done; \
			done; \
		done; \
		if [ -n "$${FMT_GO_FILES}" ]; then \
			$(GOIMPORTS) -w $${FMT_GO_FILES[@]}; \
		fi; \
		if [ -n "$${FMT_TS_FILES}" ]; then \
			prettier --config $(TOOLS_DIR)/.prettierrc.yaml -w $${FMT_TS_FILES[@]}; \
		fi; \
	}; \
	protogen "$(PROTOGEN_TARGETS)"; \
	rm -f ./vendor/$${PROJECT}

.PHONY: gen
gen: genproto

.PHONY: outdated
outdated: $(GO_MOD_OUTDATED)
	cd $(PROJECT_DIR); \
	go list -mod=mod -u -m -json all | $(GO_MOD_OUTDATED) -update -direct

.PHONY: list
list: $(GO_MOD_OUTDATED)
	cd $(PROJECT_DIR); \
	go list -mod=mod -u -m -json all | $(GO_MOD_OUTDATED)

.PHONY: lint
lint: $(GOLANGCI_LINT)
	cd $(PROJECT_DIR); \
	$(GOLANGCI_LINT) run

.PHONY: fix
fix: $(GOLANGCI_LINT)
	cd $(PROJECT_DIR); \
	$(GOLANGCI_LINT) run --fix

.PHONY: test
test:
	cd $(PROJECT_DIR); \
	go test -v ./...

.PHONY: test-browser
test-browser: $(WASMBROWSERTEST)
	cd $(PROJECT_DIR); \
	GOOS=js GOARCH=wasm go test -exec $(WASMBROWSERTEST) -v ./...

.PHONY: format
format: $(GOFUMPT) $(GOIMPORTS)
	cd $(PROJECT_DIR); \
	$(GOIMPORTS) -w ./; \
	$(GOFUMPT) -w ./

.PHONY: release
release: $(GORELEASER)
	cd $(PROJECT_DIR); \
	$(GORELEASER) release $(GORELEASER_OPTS)

.PHONY: release-bundl\e
release-bundle: $(GORELEASER)
	cd $(PROJECT_DIR); \
	$(GORELEASER) check; \
	$(GORELEASER) release --snapshot --clean --skip-publish $(GORELEASER_OPTS)

.PHONY: release-build
release-build: $(GORELEASER)
	cd $(PROJECT_DIR); \
	$(GORELEASER) check; \
	$(GORELEASER) build --single-target --snapshot --clean $(GORELEASER_OPTS)

.PHONY: release-check
release-check: $(GORELEASER)
	cd $(PROJECT_DIR); \
	$(GORELEASER) check
