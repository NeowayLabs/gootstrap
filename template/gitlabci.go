package template

const GitlabCI = `variables:
  DOCKER_DRIVER: overlay2
  DOCKER_HOST: tcp://docker:2376
  DOCKER_TLS_VERIFY: 1
  DOCKER_TLS_CERTDIR: /certs
  DOCKER_CERT_PATH: /certs/client

stages:
  - test

test:
  stage: test
  image:
    name: docker/compose:latest
    entrypoint: ["/bin/sh", "-c"]
  services:
    - docker:dind
  before_script:
    - apk add --no-cache make git musl-dev go
  script:
    - make test
`
