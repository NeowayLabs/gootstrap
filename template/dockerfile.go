package template

const DockerfileDev = `FROM golang:{{.GoVersion}}-stretch

ENV GOLANG_CI_LINT_VERSION=v{{.CILintVersion}}

RUN cd /usr && \
    wget -O - -q https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s ${GOLANG_CI_LINT_VERSION}

ARG USER
ARG USER_ID
ARG GROUP_ID

RUN groupadd -f -g ${GROUP_ID} ${USER} && \
    useradd -m -g ${GROUP_ID} -u ${USER_ID} ${USER} || echo "user already exists"

USER ${USER_ID}:${GROUP_ID}

WORKDIR /app
`

const Dockerfile = `FROM alpine:{{.AlpineVersion}} as base

RUN apk --no-cache update && \
    apk --no-cache add ca-certificates tzdata && \
    rm -rf /var/cache/apk/*

RUN adduser -D -g '' appuser

COPY ./cmd/{{.Project}}/{{.Project}} /app/{{.Project}}

FROM scratch

COPY --from=base /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=base /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=base /etc/passwd /etc/passwd
COPY --from=base /etc/group /etc/group
COPY --from=base /app/{{.Project}} /app/{{.Project}}

# Use an unprivileged user.
USER appuser

ENTRYPOINT ["/app/{{.Project}}"]
`
