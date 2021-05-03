#!/bin/bash

set -e

COMMA_DELIMITED_PATHS=$1
OUTPUT_PATH=$2
CONFIG_PATH=$3

#### Installation ####
# binary will be $(go env GOPATH)/bin/golangci-lint
# curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.39.0

# or install it into ./bin/
# curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.39.0

# In alpine linux (as it does not come with curl by default)
# wget -O- -nv https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.39.0

#### Install Dependencies ####
go mod download

#### Run ####

run_lint()
{
  $(go env GOPATH)/bin/golangci-lint run ${COMMA_DELIMITED_PATHS//,/} $( if [ -n "$CONFIG_PATH" ]; then echo "--config $CONFIG_PATH"; fi ) --timeout "5m" --out-format checkstyle || true
}
echo "$(run_lint)" > ${OUTPUT_PATH}
