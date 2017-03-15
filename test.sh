#!/usr/bin/env nash

packages <= go list ./... | grep -v vendor
for package in $packages {
        go test -race $package
}
