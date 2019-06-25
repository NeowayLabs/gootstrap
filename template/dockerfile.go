package template

const DockerfileDev = `FROM golang:{{.GoVersion}}-stretch

ENV GOLANG_CI_LINT_VERSION=v{{.CILintVersion}}

RUN cd /usr && \
    wget -O - -q https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s ${GOLANG_CI_LINT_VERSION}

ARG USER
ARG USER_ID
ARG GROUP_ID

RUN groupadd -g ${GROUP_ID} ${USER} && \
    useradd -m -g ${GROUP_ID} -u ${USER_ID} ${USER}

USER ${USER_ID}:${GROUP_ID}

WORKDIR /app
`

const Dockerfile = `FROM alpine:{{.AlpineVersion}}

RUN apk --no-cache update && \
    apk --no-cache add ca-certificates tzdata && \
    rm -rf /var/cache/apk/*

RUN adduser -D -g '' appuser

COPY ./cmd/{{.Project}}/{{.Project}} /app/{{.Project}}

# Use an unprivileged user.
USER appuser

ENTRYPOINT ["/app/{{.Project}}"]
`
