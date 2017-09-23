version     ?= latest
{{.Name}}img = {{.DockerRegistry}}/{{.Name}}:$(version)
devimg       = {{.Name}}dev
GOPATH      ?= $(HOME)/go
packagename  = $(shell pwd | sed "s:"$(GOPATH)"/src/::")
workdir      = /go/src/$(packagename)
rundev       = docker run -ti --rm -v `pwd`:$(workdir) --workdir $(workdir) $(devimg)
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
	$(rundev) ./hack/vendor.sh
	chown -R $(USER):$(USER) ./vendor
	chown -R $(USER):$(USER) ./Godeps

build: imagedev
	$(rundev) go build -v -ldflags "-X main.Version=$(gitversion)" -o ./cmd/{{.Name}}/{{.Name}} ./cmd/{{.Name}}/main.go

analyze:
	$(rundev) ./hack/analyze.sh

check: imagedev
	$(rundev) ./hack/check.sh $(pkg) $(test)

coverage: imagedev
	$(rundev) ./hack/coverage.sh

coverage-show: coverage
	xdg-open fullcov.html

shell: imagedev
	$(rundev)
