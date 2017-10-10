#!/bin/bash

set -o errexit
set -o nounset

echo "performing static analysis on the code"

for pkg in $(go list ./...); do
    go vet $pkg
    staticcheck $pkg
    gosimple $pkg
    unused $pkg
done

echo "done"
