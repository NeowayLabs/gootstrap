package template

const DockerfileDev = `FROM golang:{{.GoVersion}}-stretch

ENV GOLANG_CI_LINT_VERSION=v{{.CILintVersion}}

RUN cd /usr && \
    wget -O - -q https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s ${GOLANG_CI_LINT_VERSION}

WORKDIR /app

# WHY: Keep caches of packages inside the image so
# we dont need to download them all the time
COPY go.mod ./go.mod
COPY go.sum ./go.sum

RUN go mod download

# WHY: Since we map the host user to the container
# access to these dirs fails because of permission
# issues. This is usually a bad idea but since we just
# use this image locally to build the code it seems OK.
RUN mkdir -p /.cache && \
    chmod 777 -R /.cache && \
    chmod 777 -R /go
`

const Dockerfile = `FROM alpine:{{.AlpineVersion}}

RUN apk --no-cache update && \
    apk --no-cache add ca-certificates && \
    rm -rf /var/cache/apk/*

COPY ./cmd/{{.Project}}/{{.Project}} /app/{{.Project}}

ENTRYPOINT ["/app/{{.Project}}"]
`
