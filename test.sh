#!/usr/bin/env nash

packages <= go list ./... | grep -v vendor
packages <= split($packages, "\n")
for package in $packages {
        go test -race $package
}
