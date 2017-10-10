version     ?= latest
{{.Name}}img = {{.DockerRegistry}}/{{.Name}}:$(version)
devimg       = {{.Name}}dev
GOPATH      ?= $(HOME)/go
packagename  = $(shell pwd | sed "s:"$(GOPATH)"/src/::")
workdir      = /go/src/$(packagename)
runargs      = --rm -v `pwd`:$(workdir) --workdir $(workdir) $(devimg)
runshell     = docker run $(runargs)
runcmd       = docker run -ti $(runargs)
gitversion   = $(version)

guard-%:
	@ if [ "${${*}}" = "" ]; then \
		echo "Variable '$*' not set"; \
		exit 1; \
	fi

release: guard-version publish
	git tag -a $(version) -m "Generated release "$(version)
	git push origin $(version)

publish: image
	docker push $({{.Name}}img)

image: build
	docker build -t $({{.Name}}img) .

imagedev:
	docker build -t $(devimg) -f ./hack/Dockerfile.dev .

vendor: imagedev
	$(runcmd) ./hack/vendor.sh
	sudo chown -R $(USER):$(USER) ./vendor
	sudo chown -R $(USER):$(USER) ./Godeps

build: imagedev
	$(runcmd) go build -v -ldflags "-X main.Version=$(gitversion)" -o ./cmd/{{.Name}}/{{.Name}} ./cmd/{{.Name}}/main.go

analyze:
	$(runshell) ./hack/analyze.sh

check: imagedev
	$(runshell) ./hack/check.sh $(pkg) $(test)

check-ci: imagedev
	$(runcmd) ./hack/check.sh $(pkg) $(test)

coverage: imagedev
	$(runshell) ./hack/coverage.sh

coverage-show: coverage
	xdg-open fullcov.html

shell: imagedev
	$(runshell)
