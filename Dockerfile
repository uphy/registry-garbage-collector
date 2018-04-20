# Build latest registry
FROM golang:1.9 as builder
RUN go get github.com/docker/distribution/cmd/registry
WORKDIR /go/src/github.com/docker/distribution/cmd/registry
RUN CGO_ENABLED=0 go build -o /registry .

# Create image
FROM alpine:3.4
RUN set -ex \
    && apk add --no-cache ca-certificates apache2-utils
COPY --from=builder /registry /bin/registry
RUN chmod +x /bin/registry
RUN mkdir -p /var/lib/registry
COPY config.yml /etc/docker/registry/config.yml

ENTRYPOINT [ "/bin/registry", "garbage-collect" ]
CMD [ "-m", "/etc/docker/registry/config.yml" ]