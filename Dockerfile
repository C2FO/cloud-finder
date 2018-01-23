FROM golang:1.9 as builder
WORKDIR /go/src/github.com/c2fo/cloud-finder
ADD . .
RUN CGO_ENABLED=0 go build -o cloud-finder ./cmd/cloud-finder/main.go

FROM alpine:latest
COPY --from=builder /go/src/github.com/c2fo/cloud-finder/cloud-finder /usr/local/bin/cloud-finder
ENTRYPOINT ["/usr/local/bin/cloud-finder"]
