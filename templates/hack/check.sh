#!/usr/bin/env bash

set -o errexit
set -o nounset

for d in $(go list ./... | grep -v vendor); do
    go test -v -race $d
done
