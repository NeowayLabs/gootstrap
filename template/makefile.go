package template

import (
	"fmt"
)

type MakefileCfg struct {
	Project   string
	DockerImg string
}

func Makefile(cfg MakefileCfg) (string, error) {
	name := fmt.Sprintf("makefile:%v", cfg)
	return apply(name, makefileTemplate, cfg)
}

const makefileTemplate = `version ?= latest
img = {{.DockerImg}}:$(version)
imgdev = {{.DockerImg}}dev:$(version)
run=docker run --rm -ti -v $(shell pwd):/app $(imgdev)
cov=coverage.out
covhtml=coverage.html

all: check build

guard-%:
	@ if [ "${${*}}" = "" ]; then \
		echo "Variable '$*' not set"; \
		exit 1; \
	fi

image: build
	docker build . -t $(img)

imagedev:
	docker build . -t $(imgdev) -f ./hack/Dockerfile

release: guard-version publish
	git tag -a $(version) -m "Generated release "$(version)
	git push origin $(version)

publish: image
	docker push $(img)

shell: image
	$(run) sh

build: imagedev
	$(run) go build -v -ldflags "-X main.Version=$(version)" -o ./cmd/{{.Project}}/{{.Project}} ./cmd/{{.Project}}

check: image
	$(run) go test -timeout 60s -race -coverprofile=$(cov) ./...

coverage: check
	$(run) go tool cover -html=$(cov) -o=$(covhtml)
	xdg-open coverage.html

analyze: image
	$(run) golangci-lint run ./...
`
