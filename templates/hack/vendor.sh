#!/usr/bin/env bash

set -o nounset
set -o errexit

go get -u ./...
godep save ./...
