package template

const Makefile = `version ?= latest
img = {{.DockerImg}}:$(version)
imgdev = {{.DockerImg}}dev:$(version)
uid=$(shell id -u $$USER)
gid=$(shell id -g $$USER)
dockerbuilduser=--build-arg USER_ID=$(uid) --build-arg GROUP_ID=$(gid) --build-arg USER
wd=$(shell pwd)
modcachedir=$(wd)/.gomodcachedir
cachevol=$(modcachedir):/go/pkg/mod
appvol=$(wd):/app
run=docker run --rm -ti -v $(appvol) -v $(cachevol) $(imgdev)
cov=coverage.out
covhtml=coverage.html

all: check build

guard-%:
	@ if [ "${${*}}" = "" ]; then \
		echo "Variable '$*' not set"; \
		exit 1; \
	fi

# WHY: If cache dir does not exist it is mapped inside container as root
# If it exists it is mapped belonging to the non-root user inside the container
modcache:
	@mkdir -p $(modcachedir)

image: build
	docker build . -t $(img)

imagedev:
	@echo "building dev env"
	@docker build . -t $(imgdev) -f ./hack/Dockerfile $(dockerbuilduser) > .devimgbuild.logs
	@echo "dev env built"

release: guard-version publish
	git tag -a $(version) -m "Generated release "$(version)
	git push origin $(version)

publish: image
	docker push $(img)

build: modcache imagedev
	$(run) go build -v -ldflags "-X main.Version=$(version)" -o ./cmd/{{.Project}}/{{.Project}} ./cmd/{{.Project}}

check: modcache imagedev
	$(run) go test -timeout 60s -race -coverprofile=$(cov) ./...

coverage: modcache check
	$(run) go tool cover -html=$(cov) -o=$(covhtml)
	xdg-open coverage.html

analyze: modcache imagedev
	$(run) golangci-lint run ./...

modtidy: modcache imagedev
	$(run) go mod tidy

fmt: modcache imagedev
	$(run) gofmt -w -s -l .

githooks:
	@echo "copying git hooks"
	@mkdir -p .git/hooks
	@cp hack/githooks/pre-commit .git/hooks/pre-commit
	@chmod +x .git/hooks/pre-commit
	@echo "git hooks copied"

shell: modcache imagedev
	$(run) sh
`
