#!/usr/bin/env nash

fn cover(project) {
	cwd <= pwd

	chdir($GOPATH+"/src/"+$project)

	packages     <= go list ./... | grep -v vendor
	packages     <= split($packages, "\n")

	coveragefile = "coverage.txt"
	profilefile  = "profile.out"

	rm -f $coveragefile
	echo "mode: count" > $coveragefile

	for package in $packages {
		rm -f $profilefile
		go test -v -race "-coverprofile="+$profilefile -covermode=atomic $package

		_, status <= test -f $profilefile

		if $status == "0" {
			cat $profilefile | tail -n "+2" | tee --append $coveragefile >[1=]
		}
	}

	rm -f $profilefile
	go tool cover -func $coveragefile

	chdir($cwd)
}
