# Licensed Materials - Property of tenxcloud.com
# (C) Copyright 2018 Dreamxos. All Rights Reserved.
# 2018-10-08  @author gaozh

VERSION ?= v3.1
PREFIX ?= 192.168.1.52/tenx_containers

BINARY_BUILD_IMAGE = golang:1.9.2-alpine3.7
DOCKER_BUILD_IMAGE = $(PREFIX)/sample-operator:$(VERSION)

all: build-image

build-binary: clean
	docker run --rm -v $(shell pwd):/go/src/$(shell basename $(shell pwd)) \
	-w /go/src/$(shell basename $(shell pwd))/cmd/controller/ ${BINARY_BUILD_IMAGE} \
	/bin/sh -c "go build -v -o ../../deploy/docker/controller"

build-image: build-binary
	docker build -t ${DOCKER_BUILD_IMAGE} deploy/docker/

clean:
	rm -rf deploy/docker/controller

.PHONY: all build-binary build-image clean
