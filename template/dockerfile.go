package template

import (
	"fmt"
)

type DockerfileDevCfg struct {
	GoVersion     string
	CILintVersion string
}

type DockerfileCfg struct {
	Project string
	AlpineVersion string
}

func DockerfileDev(cfg DockerfileDevCfg) (string, error) {
	name := fmt.Sprintf("dockerfiledev:%v", cfg)
	return apply(name, dockerfileDevTemplate, cfg)
}

func Dockerfile(cfg DockerfileCfg) (string, error) {
	name := fmt.Sprintf("dockerfile:%v", cfg)
	return apply(name, dockerfileTemplate, cfg)
}

const dockerfileDevTemplate = `FROM golang:{{.GoVersion}}-stretch

ENV GOLANG_CI_LINT_VERSION=v{{.CILintVersion}}

RUN cd /usr && \
    wget -O - -q https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s ${GOLANG_CI_LINT_VERSION}

WORKDIR /app

COPY go.mod ./go.mod
COPY go.sum ./go.sum

RUN go mod download
`

const dockerfileTemplate = `FROM alpine:{{.AlpineVersion}}

RUN apk --no-cache update && \
    apk --no-cache add ca-certificates && \
    rm -rf /var/cache/apk/*

COPY ./cmd/{{.Project}}/{{.Project}} /app/{{.Project}}

ENTRYPOINT ["/app/{{.Project}}"]
`
