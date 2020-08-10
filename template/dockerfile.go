package template

const Dockerfile = `# ---------------------------------------------------------------------
#  The first stage container, for image dev
# ---------------------------------------------------------------------
FROM golang:{{.GoVersion}}-stretch as base

RUN apt-get update && \
    apt-get dist-upgrade -y && \
    apt-get install -y --no-install-recommends ca-certificates tzdata && \
    update-ca-certificates

WORKDIR /app

COPY go.sum go.mod ./
RUN go mod download

COPY . .

# ---------------------------------------------------------------------
#  The second stage container, for building the application
# ---------------------------------------------------------------------
FROM base AS builder

ARG VERSION

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-w -s -X main.version=${VERSION}" -o /go/bin/{{.Project}} ./cmd/{{.Project}}

# ---------------------------------------------------------------------
#  The third stage container, for running the application
# --------------------------------------------------------------------
FROM alpine:{{.AlpineVersion}}

COPY --from=base /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=base /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=base /etc/passwd /etc/passwd
COPY --from=base /etc/group /etc/group
COPY --from=builder /go/bin/{{.Project}} /bin/{{.Project}}

# Use an unprivileged user.
USER nobody

ENTRYPOINT ["/bin/{{.Project}}"]
`
