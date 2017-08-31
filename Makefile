# A Self-Documenting Makefile: http://marmelab.com/blog/2016/02/29/auto-documented-makefile.html

PACKAGE = github.com/borisding1994/hathcoin
VERSION= `cat $(CURDIR)/VERSION`
COMMIT_HASH = `git rev-parse --short HEAD 2>/dev/null || echo "unreleased" `
LDFLAGS = -ldflags "-X ${PACKAGE}/cmd.CommitHash=${COMMIT_HASH} -X ${PACKAGE}/cmd.Version=${VERSION}"
DIST_DIR:=$(CURDIR)/dist
GO_PACKAGES = $(shell find . -name "*.go" | grep -vE ".git|.env|vendor" | sed 's:/[^/]*$$::' | sort | uniq)
GO_TEST_PACKAGES = $(shell find . -name "*_test.go" | grep -vE ".git|.env|vendor" | sed 's:/[^/]*$$::' | sort | uniq)

# allow user to override go executable by running as GOEXE=xxx make ... on unix-like systems
GOEXE ?= go

.PHONY: all clean check-required-toolset dep-install dep-update help build test lint
.DEFAULT_GOAL := help

all: dep-install lint test build

build: clean ## build HashCoin
	@mkdir -p ${DIST_DIR}
	@${GOEXE} build -race -x ${LDFLAGS} -o ${DIST_DIR}/hathcoin ${PACKAGE}
	mkdir -p ${DIST_DIR}/config
	mkdir -p ${DIST_DIR}/logs
	mkdir -p ${DIST_DIR}/db
	@cp $(CURDIR)/LICENSE ${DIST_DIR}/LICENSE
	@cp $(CURDIR)/config/hathcoin.toml ${DIST_DIR}/config/hathcoin.toml

test: ## run test
	@HAC_CONFIG=$(CURDIR)/config/hathcoin.toml ${GOEXE} test -v $(GO_TEST_PACKAGES) -outputdir $(CURDIR)

lint: ## run code lint
	@gometalinter.v1 --config .linter.conf --vendor ./...

check-required-toolset:
	@command -v dep > /dev/null || (echo "Install golang/dep..." && go get -u github.com/golang/dep/cmd/dep)
	@command -v gometalinter.v1 > /dev/null || (echo "Install gometalinter..." && go get -u gopkg.in/alecthomas/gometalinter.v1 && gometalinter.v1 --install)


dep-install: check-required-toolset ## install go dependencies
	dep ensure

dep-update: ## update go dependencies
	dep ensure -update

clean: ## clean the build artifacts
	@rm -rf $DIST_DIR/*

generate-protobuf: ## Compiling protocol buffers
	cd rpc;protoc --gofast_out=. --js_out=library=hathcoin,binary:. --python_out=. --ruby_out=. --java_out=. ./hathcoin.proto

help: ## help
	@echo "HathCoin Makefile Tasks list:"
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {sub("\\\\n",sprintf("\n%22c"," "), $$2);printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)
