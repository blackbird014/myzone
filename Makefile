#!/usr/bin/make -f

BUILDDIR ?= $(CURDIR)/build
PACKAGES=$(shell go list ./... | grep -v '/simulation')
VERSION := $(shell echo $(shell git describe --tags --always) | sed 's/^v//')
COMMIT := $(shell git log -1 --format='%H')
BINARY = myzoned

# Check for Go installation
GO_BINARY := $(shell which go)
GOPATH := $(shell $(GO_BINARY) env GOPATH)
GOBIN := $(GOPATH)/bin

LDFLAGS = -X github.com/cosmos/cosmos-sdk/version.Name=MyZone \
	-X github.com/cosmos/cosmos-sdk/version.AppName=myzoned \
	-X github.com/cosmos/cosmos-sdk/version.Version=$(VERSION) \
	-X github.com/cosmos/cosmos-sdk/version.Commit=$(COMMIT)

BUILD_FLAGS := -ldflags '$(LDFLAGS)'

all: install lint test

build: go.sum
	@echo "Building myzoned binary..."
	@go build -mod=readonly $(BUILD_FLAGS) -o $(BUILDDIR)/$(BINARY) ./cmd/myzoned

install: go.sum
	@echo "Installing myzoned binary..."
	@go install -mod=readonly $(BUILD_FLAGS) ./cmd/myzoned

build-linux: go.sum
	@echo "Building myzoned binary for Linux..."
	@GOOS=linux GOARCH=amd64 go build -mod=readonly $(BUILD_FLAGS) -o $(BUILDDIR)/$(BINARY)-linux-amd64 ./cmd/myzoned

build-mac: go.sum
	@echo "Building myzoned binary for MacOS..."
	@GOOS=darwin GOARCH=amd64 go build -mod=readonly $(BUILD_FLAGS) -o $(BUILDDIR)/$(BINARY)-darwin-amd64 ./cmd/myzoned

test:
	@echo "Running tests..."
	@go test -mod=readonly -v $(PACKAGES)

test-race:
	@echo "Running tests with race detector..."
	@go test -mod=readonly -race -v $(PACKAGES)

lint:
	@echo "Running linter..."
	@go vet ./...
	@if command -v golangci-lint >/dev/null; then \
		golangci-lint run; \
	else \
		echo "Skipping golangci-lint check: not installed"; \
	fi

clean:
	@echo "Cleaning build directory..."
	@rm -rf $(BUILDDIR)/

go.sum: go.mod
	@echo "Ensuring dependencies are available..."
	@go mod verify
	@go mod tidy

proto-gen:
	@echo "Generating protobuf files..."
	@if ! command -v protoc > /dev/null; then \
		echo "Error: protoc is not installed, please install it first." >&2; \
		exit 1; \
	fi
	@if ! command -v protoc-gen-gocosmos > /dev/null; then \
		echo "Installing protoc-gen-gocosmos..."; \
		go install github.com/cosmos/gogoproto/protoc-gen-gocosmos; \
	fi
	@./scripts/protocgen.sh

.PHONY: all build install build-linux build-mac test test-race lint clean proto-gen 