# Copyright (c) 2018 Volterra, Inc. All rights reserved.

TOP := $(PWD)


# path from $(GOPATH)/src e.g. for ~/go/src/github.com/foo,
# GOSRC_PATH will be src/github.com/foo
GOSRC_PATH := $(shell echo $(TOP) | sed "s^.*src/\(.*\)^src/\1^")
SERVICE := $(shell basename ${TOP})

# Golang packages name, e.g. gopkg.volterra.us/vpm
PACKAGE = $(shell echo "${PWD}" | sed "s,$$GOPATH/src/,,g")

# Timeout for golang tests
TEST_TIMEOUT ?= 600s

# GO environemnt settings
GO_CMD=$(shell which richgo 2> /dev/null || which go)
CGO_ENABLED ?= 0
GOCACHE ?= ${PWD}/.cache/go
DEPCACHEDIR ?= ${PWD}/.cache/dep/
SOURCE_FILES=$(shell find ./ -name '*.go' | grep -v "./vendor/\|./pbgo/\|./.cache/")

# Copyright line is required in all files
COPYRIGHT="// Copyright (c) 20[1-9][0-9] Volterra, Inc. Licensed under APACHE LICENSE, VERSION 2.0"

# Local server with documentation is listening on this address
LISTEN_DOC = localhost:6060

# Identification for current build
BUILD_COMMENT ?= $(shell date -Iseconds; whoami)
BUILD_SHA ?= $(shell git rev-parse HEAD)
BUILD_VERSION ?= $(shell git rev-parse --abbrev-ref HEAD)
BUILD_IMAGE ?= ${DOCKER_REG}/ves.io/${SERVICE}:${ARTIFACT_TAG}

# Current user
UID=$(shell id -u)
GID=$(shell id -g)

# Detect operating system (only Darwin or Linux, others are unsupported)
UNAME_S := $(shell uname -s)
ifeq ($(UNAME_S),Linux)
	OS = linux
endif
ifeq ($(UNAME_S),Darwin)
	OS = darwin
endif

# all target runs full workflow
all: debug dep-ensure generate unittest build

# debug target is used to display configuration and importang variables
debug:
	@echo "OS: ${OS}"
	@echo "UID,GID: ${UID},${GID}"
	@echo "GOPATH: ${GOPATH}"
	@echo "TOP: ${TOP}"
	@echo "SERVICE: ${SERVICE}"
	@echo "PACKAGE: ${PACKAGE}"
	@echo "GOSRC_PATH: ${GOSRC_PATH}"
	@echo "GOCACHE: ${GOCACHE}"
	@echo "DEPCACHEDIR: ${DEPCACHEDIR}"

#
# Bare commands (prefixed with cmd)
#

.PHONY: cmd-go-mod-init
cmd-go-mod-init:
	GOPRIVATE=gopkg.volterra.us go mod init

.PHONY: cmd-go-mod-vendor
cmd-go-mod-vendor:
	GOPRIVATE=gopkg.volterra.us go mod vendor

.PHONY: cmd-go-mod-download
cmd-go-mod-download:
	GOPRIVATE=gopkg.volterra.us go mod download

.PHONY: cmd-go-mod-tidy
cmd-go-mod-tidy:
	GOPRIVATE=gopkg.volterra.us go mod tidy

.PHONY: cmd-go-mod-get
cmd-go-mod-get:
	GOPRIVATE=gopkg.volterra.us go get $(repo); go mod vendor

cmd-shell:
	sh

# unittest, lint, converage
cmd-test: cmd-lint cmd-unittest
cmd-lint:
	${GO_CMD} vet $$(go list ./...)

	golint $$(go list ./...)
	#gocyclo -over 23 -avg $$(go list -f "{{ .Dir }}" ./... | grep -v /pbgo)
	deadcode .
	gofmt -d -e -l -s $$(go list -f "{{ .Dir }}" ./...)

	@echo Checking copyright in files
	@err=0; \
	for f in ${SOURCE_FILES}; do \
		grep ${COPYRIGHT} $${f} > /dev/null; \
		if [ "$${?}" != "0" ]; then \
			echo "Missing copyright in $${f}"; \
			let "err++"; \
		fi \
	done; \
	if [ "$${err}" != "0" ]; then exit 1; fi

cmd-unittest:
	mkdir -p ./coverage_out/
	${GO_CMD} test -race -timeout ${TEST_TIMEOUT} -vet="" -coverprofile ./coverage_out/all -v $$(go list ./... | grep -v /pbgo)

	# split to generated code and rest
	cat ./coverage_out/all | grep -v "/pbgo/" > ./coverage_out/bl
	cat ./coverage_out/all | grep "/pbgo/\|^mode:" > ./coverage_out/generated

	# report for generated code
	${GO_CMD} tool cover -func ./coverage_out/generated | sed 's/total:\t\t/generated total:/g'

	# report for business logic
	${GO_CMD} tool cover -func ./coverage_out/bl

cmd-build-bin:
	mkdir -p "${GOBIN_OUT}"
	/bin/sh -c \
	'set -xe; for pkg in $$(go list ./cmd/...); do \
		pkgName=$$(go list -f "{{ .Name }}" $${pkg}); \
		if [ "$${pkgName}" = "main" ]; then \
			echo "Building $$(basename $${pkg}) ${BUILD_VERSION} ${BUILD_SHA} ${BUILD_COMMENT}"; \
			CGO_ENABLED=${CGO_ENABLED} GOOS=linux \
				go build -a -installsuffix cgo \
				-ldflags "-X \"${PACKAGE}/vendor/gopkg.volterra.us/stdlib/version.Version=${BUILD_VERSION}\" -X \"${PACKAGE}/vendor/gopkg.volterra.us/stdlib/version.SHA=${BUILD_SHA}\" -X \"${PACKAGE}/vendor/gopkg.volterra.us/stdlib/version.Comment=${BUILD_COMMENT}\"" \
					-o ${GOBIN_OUT}/$$(basename $${pkg}) \
					$${pkg}; \
		fi; \
	done'

cmd-clean:
	rm -rfv *_out/

cmd-doc:
	@echo "Documentation available at http://${LISTEN_DOC}/pkg/${PACKAGE}"
	godoc -http=${LISTEN_DOC}

install-tf-volterra:
	cd cmd/terraform-provider-volterra && go build .
	mkdir -p ~/.terraform.d/plugins/
	mv cmd/terraform-provider-volterra/terraform-provider-volterra ~/.terraform.d/plugins/