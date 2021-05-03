#!/bin/bash

# Install goreleaser
curl -sfL https://install.goreleaser.com/github.com/goreleaser/goreleaser.sh | sh

ARGS=$1

./bin/goreleaser ${ARGS}
