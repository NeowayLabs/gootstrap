#!/bin/bash

set -o errexit
set -o nounset

projectdir=$(mktemp -d)
trap "rm -rf $projectdir" EXIT

echo
echo "testdir is: "$projectdir

echo "building"
go build ./cmd/gootstrap

echo "running gootstrap"
./gootstrap --output-dir $projectdir --module "whatever.com/group/project" --docker-image "group/project"

echo
echo "let the tests begin"
cd $projectdir

make build
make check
make analyze
make coverage
make image

echo
echo "basic functionalities seems intact"
