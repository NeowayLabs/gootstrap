version ?= latest
{{project}}img={{docker_registry}}/{{project}}:$(version)
devimg={{project}}dev
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
	docker push $({{project}}img)

image: build
	docker build -t $({{project}}img) .

imagedev:
	docker build -t $(devimg) -f ./hack/Dockerfile.dev .

vendor: imagedev
	$(rundev) ./hack/vendor.sh
	sudo chown -R $(USER):$(USER) ./vendor

build: imagedev
	$(rundev) ./hack/build.sh

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
