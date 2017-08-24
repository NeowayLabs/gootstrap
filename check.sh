#!/bin/bash

set -o errexit
set -o nounset

projectdir=$(mktemp --dry-run -d $GOPATH/src/gootstraptest/XXXXXXX)
trap "rm -rf $projectdir" EXIT
mkdir -p $projectdir

echo
echo "testdir is: "$projectdir

echo "building"
go build .

echo "running gootstrap"
./gootstrap --output $projectdir --project gootstraptest --docker-registry gootstraptest

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
