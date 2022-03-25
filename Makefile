PROJECT_NAME := "github.com/suanve/su.proxy"
PKG := "$(PROJECT_NAME)"
# GO_FILES := $(shell find . -name '*.go' | grep -v /vendor/ | grep -v _test.go)

.PHONY: all dep lint vet test test-coverage build clean

all: build

# dep: ## Get the dependencies
# 	@go mod download

vet: ## Run go vet
	@go vet ${PKG_LIST}

build: 
	go build -o build/main .

clean: ## Remove previous build
	@rm -f ./build