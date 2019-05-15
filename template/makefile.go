package template

const Makefile = `version ?= latest
img = {{.DockerImg}}:$(version)
imgdev = {{.DockerImg}}dev:$(version)
uid=$(shell id -u $$USER)
gid=$(shell id -g $$USER)
dockerbuilduser=--build-arg USER_ID=$(uid) --build-arg GROUP_ID=$(gid) --build-arg USER
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
	docker build . -t $(imgdev) -f ./hack/Dockerfile $(dockerbuilduser)

release: guard-version publish
	git tag -a $(version) -m "Generated release "$(version)
	git push origin $(version)

publish: image
	docker push $(img)

shell: imagedev
	$(run) sh

build: imagedev
	$(run) go build -v -ldflags "-X main.Version=$(version)" -o ./cmd/{{.Project}}/{{.Project}} ./cmd/{{.Project}}

check: imagedev
	$(run) go test -timeout 60s -race -coverprofile=$(cov) ./...

coverage: check
	$(run) go tool cover -html=$(cov) -o=$(covhtml)
	xdg-open coverage.html

analyze: imagedev
	$(run) golangci-lint run ./...
`
