#!/bin/bash

go get -u golang.org/x/lint/golint

GOLINT_PATH=$(go list -f {{.Target}} golang.org/x/lint/golint)
export PATH="$GOLINT_PATH:$PATH"

golint $1 > $2
