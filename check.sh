#!/bin/bash

set -o errexit
set -o nounset

projectdir=$(mktemp -d)
trap "rm -rf $projectdir" EXIT

echo
echo "testdir is: "$projectdir

echo "building"
go build -o ./cmd/gootstrap/gootstrap ./cmd/gootstrap

echo "running gootstrap"
./cmd/gootstrap/gootstrap -output-dir $projectdir -module "whatever.com/group/project" -image "group/project"

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
echo "formatting"
make fmt

echo
echo "mod tidy"
make modtidy

echo
echo "githooks"
make githooks

echo
echo "final image build"
make image

echo
echo "basic functionalities seems intact"
