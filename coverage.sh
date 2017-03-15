#!/usr/bin/env nash

fn cover(project) {
	cwd <= pwd

	chdir($GOPATH+"/src/"+$project)

	packages     <= go list ./... | grep -v vendor
	packages     <= split($packages, "\n")

	coveragefile = "coverage.txt"

	rm -f $coveragefile
	echo "mode: count" > coverage.txt

	for package in $packages {
		canon <= canonPath($package)

		profilefile = "."+$canon+".profile"

		go test -v -race "-coverprofile="+$profilefile -covermode=atomic $package

		_, status <= test -f $profilefile

		if $status == "0" {
			cat $profilefile | tail -n "+2" | tee --append $coveragefile >[1=]

			rm $profilefile
		}
	}

	go tool cover -func coverage.txt

	chdir($cwd)
}

fn canonPath(package) {
	canon <= echo $package | sed "s#/#.#g"

	return $canon
}
