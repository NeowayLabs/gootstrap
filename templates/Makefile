version ?= latest
{{.Name}}img={{.DockerRegistry}}/{{.Name}}:$(version)
devimg={{.Name}}dev
rundev=docker run -ti --rm $(devimg)

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
	sudo chown -R $(USER):$(USER) ./vendor

build: imagedev
	$(rundev) go build ./cmd/{{.Name}}/main.go -ldflags "-X main.Version=$(version)" -v

analyze:
	$(rundev) ./hack/analyze.sh

check: imagedev
	$(rundev) ./hack/check.sh $(pkg) $(test)

coverage: imagedev
	$(rundev) ./hack/coverage.sh

coverage-show: coverage
	xdg-open fullcov.html

shelldev: imagedev
	$(rundev)
