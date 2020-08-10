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
make test

echo
echo "analyzing"
make static-analysis

echo
echo "basic functionalities seems intact"
