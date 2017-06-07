PACKAGE = cibully
IMPORT_PATH := github.com/ahelal/ci-bully
DATE    ?= $(shell date +%FT%T%z)
VERSION ?= $(shell cat $(CURDIR)/.VERSION 2> /dev/null || echo v0)


GOLINT_COMMAND := $(shell command -v golint 2> /dev/null)
M = $(shell printf "\033[34;1m▶\033[0m")

include *.mk

.PHONY: all deps test lint

all: deps test lint dev

dev: set_dev build_linux build_darwin

release: set_release build_linux build_darwin

deps:
ifneq (,$(wildcard ./glide.yml))
	$(info $(M) Installing glide dependencies…)
	@go get -u github.com/Masterminds/glide
	@glide install
endif
	$(info $(M) GO get dependencies…)
	@#@go get

test:
	$(info $(M) GO get dependencies…)

set_release:
	$(eval VERSION_NAME := $(VERSION))
	$(eval BUILD_NAME := Release)

set_dev:
	$(eval VERSION_NAME := $(VERSION)-$(DATE))
	$(eval BUILD_NAME := DEV)
	$(eval PACKAGE := ${PACKAGE}_dev)

build_linux:
	$(info $(M) Building $(BUILD_NAME) linux version $(VERSION_NAME)…)
	@GOOS=linux GOARCH=amd64 go build \
		-tags release \
		-ldflags '-X main.version=$(VERSION_NAME)' \
		-o build/$(PACKAGE)_$(VERSION)_linux_amd64 *.go

build_darwin:
	$(info $(M) Building $(BUILD_NAME) darwin version $(VERSION_NAME)…)
	@GOOS=darwin GOARCH=amd64 go build \
		-tags release \
		-ldflags '-X main.version=$(VERSION_NAME)' \
		-o build/$(PACKAGE)_$(VERSION)_darwin_amd64 *.go

lint:
ifndef GOLINT_COMMAND
	$(info $(M) Installing golint…)
	@go get -u github.com/golang/lint/golint
endif
	$(info $(M) Linting…)
	@golint -set_exit_status $(go list ./... | grep -v /vendor/)
