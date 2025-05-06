#!/usr/bin/env bash

# Copyright 2017 The Kubernetes Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -o errexit
set -o nounset
set -o pipefail

SCRIPT_ROOT=$(dirname "${BASH_SOURCE[0]}")/..
CODEGEN_PKG=${CODEGEN_PKG:-$(cd "${SCRIPT_ROOT}"; ls -d -1 ./vendor/k8s.io/code-generator 2>/dev/null || echo ../code-generator)}

# generate the code with:
# --output-base    because this script should also be able to run inside the vendor dir of
#                  k8s.io/kubernetes. The output-base is needed for the generators to output into the vendor dir
#                  instead of the $GOPATH directly. For normal projects this can be dropped.
# 其中appcontroller/pkg/generated和appcontroller/pkg/apis的appcontroller为包名，appcontroller:v1alpha1为group和version，即appcontroller/pkg/apis下的目录。而--output-base表示生成代码的目录，--output-base应与项目的上级目录对齐，比如项目目录为/home/appcontroller，那么--output-base应为/home，如果--output-base为/tmp，那么生成的代码位置为/tmp/appcontroller/pkg/generated，所以基于此，项目目录名应与包名一致，否则生成代码位置将不能与项目对齐
bash "${CODEGEN_PKG}"/generate-groups.sh "deepcopy,client,informer,lister" \
  appcontroller/pkg/generated appcontroller/pkg/apis \
  appcontroller:v1alpha1 \
  --output-base "$(dirname "${BASH_SOURCE[0]}")/../../.." \
  --go-header-file "${SCRIPT_ROOT}"/hack/boilerplate.go.txt

# To use your own boilerplate text append:
#   --go-header-file "${SCRIPT_ROOT}"/hack/custom-boilerplate.go.txt
