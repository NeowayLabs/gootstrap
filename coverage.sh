#!/usr/bin/env nash

packages <= go list ./... | grep -v vendor
echo "generating coverage for packages: " + $packages

coveragefile = "coverage.txt"
echo "" > $coveragefile

for package in $packages {
        profilefile = "." + $package + ".profile"
        go test -v -race -coverprofile=$profilefile -covermode=atomic $package
        -ls $profilefile
        if $status == "0" {
                cat $profilefile | tee --append $coveragefile
                rm $profilefile
        }
}
