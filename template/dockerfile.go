package template

const DockerfileDev = `FROM golang:{{.GoVersion}}-stretch

ENV GOLANG_CI_LINT_VERSION=v{{.CILintVersion}}

RUN cd /usr && \
    wget -O - -q https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s ${GOLANG_CI_LINT_VERSION}

WORKDIR /app

COPY go.mod ./go.mod
COPY go.sum ./go.sum

RUN go mod download
`

const Dockerfile = `FROM alpine:{{.AlpineVersion}}

RUN apk --no-cache update && \
    apk --no-cache add ca-certificates && \
    rm -rf /var/cache/apk/*

COPY ./cmd/{{.Project}}/{{.Project}} /app/{{.Project}}

ENTRYPOINT ["/app/{{.Project}}"]
`
