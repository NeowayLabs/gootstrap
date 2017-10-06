FROM alpine:3.6

RUN apk --no-cache update && \
    apk --no-cache add ca-certificates && \
    rm -rf /var/cache/apk/*

COPY ./cmd/{{.Name}}/{{.Name}} /app/{{.Name}}

ENTRYPOINT ["/app/{{.Name}}"]
