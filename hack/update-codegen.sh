#!/usr/bin/env bash

# Licensed Materials - Property of tenxcloud.com
# (C) Copyright 2018 Dreamxos. All Rights Reserved.
# 2018-09-18  @author gaozh

SCRIPT_ROOT=$(dirname ${BASH_SOURCE})/..
CODEGEN_PKG=${CODEGEN_PKG:-$(cd ${SCRIPT_ROOT}; ls -d -1 ./vendor/k8s.io/code-generator 2>/dev/null || echo ../code-generator)}

${CODEGEN_PKG}/generate-groups.sh \
  all \
  sample-operator/pkg/generated \
  sample-operator/pkg/apis \
  daas:v1 \
  --go-header-file ${SCRIPT_ROOT}/hack/boilerplate.go.txt

