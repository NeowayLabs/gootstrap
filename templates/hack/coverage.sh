#!/usr/bin/env bash

set -o errexit
set -o nounset

coveragemode="mode: atomic"
echo $coveragemode > coverage.txt

for d in $(go list ./... | grep -v vendor); do
    go test -v -race -coverprofile=profile.out -covermode=atomic $d
    if [ -f profile.out ]; then
        grep -h -v "^${coveragemode}" profile.out >> coverage.txt
        rm profile.out
    fi
done

go tool cover -html=coverage.txt -o=coverage.html
