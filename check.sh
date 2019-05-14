#!/bin/bash

set -o errexit
set -o nounset

export GO111MODULE="on"

projectdir=$(mktemp -d)
trap "rm -rf $projectdir" EXIT

echo
echo "testdir is: "$projectdir

echo "building"
go build -o ./cmd/gootstrap/gootstrap ./cmd/gootstrap

echo "running gootstrap"
./cmd/gootstrap/gootstrap --output-dir $projectdir --module "whatever.com/group/project" --docker-image "group/project"

echo
echo "let the tests begin"
cd $projectdir

echo
echo "building"
make build

echo
echo "checking"
make check

echo
echo "analyzing"
make analyze

echo
echo "coverage"
make coverage

echo
echo "final image build"
make image

echo
echo "basic functionalities seems intact"
