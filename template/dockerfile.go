package template

const Dockerfile = `# ---------------------------------------------------------------------
#  The first stage container, for image base
# ---------------------------------------------------------------------
FROM golang:{{.GoVersion}}-stretch as base

ENV GOLANG_CI_LINT_VERSION=v{{.CILintVersion}}

RUN cd /usr && \
    wget -O - -q https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s ${GOLANG_CI_LINT_VERSION}

ARG USER_ID
ARG GROUP_ID

RUN groupadd -f -g ${GROUP_ID} appuser && \
    useradd -m -g ${GROUP_ID} -u ${USER_ID} appuser || echo "user already exists"

USER ${USER_ID}:${GROUP_ID}

WORKDIR /app

# ---------------------------------------------------------------------
#  The second stage container, for building the application
# ---------------------------------------------------------------------
FROM base AS builder

RUN apt-get update && \
    apt-get dist-upgrade -y && \
    apt-get install -y --no-install-recommends ca-certificates tzdata && \
	update-ca-certificates

RUN adduser --disabled-password --gecos '' appuser

WORKDIR $GOPATH/src/{{.Module}}

COPY . .

RUN go mod download

ARG VERSION

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-w -s -X main.Version=${VERSION}" -o /go/bin/{{.Project}} ./cmd/{{.Project}}

# ---------------------------------------------------------------------
#  The third stage container, for running the application
# --------------------------------------------------------------------
FROM scratch

COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group
COPY --from=builder /go/bin/{{.Project}} /bin/{{.Project}}

# Use an unprivileged user.
USER appuser

ENTRYPOINT ["/bin/{{.Project}}"]
`
