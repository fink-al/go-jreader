HAS_GO := $(shell hash go > /dev/null 2>&1 && echo yes)
ifeq ($(HAS_GO), yes)
export GOPATH ?= $(shell go env GOPATH)
endif

GOTESTSUM_PACKAGE := gotest.tools/gotestsum@latest
GOFUMPT_PACKAGE := mvdan.cc/gofumpt@v0.5.0
GOLANGBADGE_PACKAGE := github.com/jpoles1/gopherbadger@latest
GO_VULNCHECK := golang.org/x/vuln/cmd/govulncheck@latest
GODOC_PACKAGE := golang.org/x/tools/cmd/godoc@latest
GOLANGCI_PACKAGE := github.com/golangci/golangci-lint/cmd/golangci-lint@v1.55.2

GOTESTFLAGS := -timeout 360s -cover -coverprofile .cover.out
ifeq ($(RACE_ENABLED),true)
GOTESTFLAGS += -race
endif
GO_TEST_CMD ?= go run $(GOTESTSUM_PACKAGE) --format-hide-empty-pkg --

GO_SOURCES := $(wildcard *.go) $(shell find . -path ./node_modules -prune -false -o -type f -name "*.go")

#########
# DEPS
#########

package-lock.json: package.json
	npm install --package-lock-only

node_modules: package-lock.json
	npm install --no-save
	@touch node_modules

.PHONY: deps
deps: node_modules
	go mod download

.PHONY: deps-tools
deps-tools: .golangci.yml
	go install $(GOLANGBADGE_PACKAGE)
	go install $(GOFUMPT_PACKAGE)

.PHONY: update-go
update-go:
	go get -u
	$(MAKE) --no-print-directory tidy


#########
# LINT
#########

.PHONY: lint-go-vuln
lint-go-vuln: $(GO_SOURCES)
	go run $(GO_VULNCHECK) ./...

.PHONY: lint-go
lint-go: $(GO_SOURCES)
	go run $(GOLANGCI_PACKAGE) --color=always run

.PHONY: lint
lint: lint-go lint-go-vuln

#########
# TEST
#########

.PHONY: test-backend
test-backend: $(GO_SOURCES)
	$(GO_TEST_CMD) $(GOTESTFLAGS) ./...

.PHONY: test-backend-verbose
test-backend-verbose: $(GO_SOURCES)
	go test -v -timeout 360s -cover ./... -coverprofile .cover.out

.PHONY: test-backend-badge
test-backend-badge:
	@go run $(GOLANGBADGE_PACKAGE) -md="README.md" -manualcov $(shell go tool cover -func .cover.out | grep total: | awk '{print $$3}' | sed -r 's/%//g') -png=false

.PHONY: test
test: test-backend test-backend-badge

.PHONY: test-verbose
test-verbose: test-backend-verbose

.PHONY: test-benchmark
test-benchmark: $(GO_SOURCES)
	go test -run ^$$ -bench ./...

#########
# MISC
#########

.PHONY: fmt
fmt:
	@go run $(GOFUMPT_PACKAGE) -l -w $(GO_SOURCES)

.PHONY: tidy
tidy:
	$(eval MIN_GO_VERSION := $(shell grep -Eo '^go\s+[0-9]+\.[0-9.]+' go.mod | cut -d' ' -f2))
	go mod tidy -compat=$(MIN_GO_VERSION)

#########
# VERSIONING
#########

.PHONY: patch
patch: node_modules
	npx versions -p patch
	@make readme-version

.PHONY: minor
minor: node_modules
	npx versions -p minor
	@make readme-version

.PHONY: major
major: node_modules
	npx versions -p major
	@make readme-version

.PHONY: readme-version
readme-version:
	@sed -i '' -E "s/img.shields.io\/badge\/Version-v[0-9]+\.[0-9]+\.[0-9]+/img.shields.io\/badge\/Version-$$(git describe --abbrev=0)/g" README.md
	@git add README.md

#########
# DOCUMENTATION
#########

.PHONY: godoc
godoc:
	@echo "Starting godoc server on http://localhost:6060"
	@go run $(GODOC_PACKAGE) -http=:6060
