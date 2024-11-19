FROM golang:1.23 AS builder
COPY . /go/src/github.com/skpr/crond
WORKDIR /go/src/github.com/skpr/crond
RUN CGO_ENABLED=0 go build -a -ldflags '-extldflags "-static"' -o skpr-crond github.com/skpr/crond/cmd/skpr-crond

FROM alpine:3.20
COPY --from=builder /go/src/github.com/skpr/crond/skpr-crond /usr/local/bin/skpr-crond
RUN chmod +x /usr/local/bin/skpr-crond
CMD ["/usr/local/bin/skpr-crond"]
