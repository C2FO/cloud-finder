FROM golang:1.12 as builder
ENV GO111MODULE on
WORKDIR /cloud-finder
ADD . .
RUN CGO_ENABLED=0 go install ./cmd/cloud-finder/main.go

FROM alpine:latest
COPY --from=builder /go/bin/cloud-finder /usr/local/bin/cloud-finder
ENTRYPOINT ["/usr/local/bin/cloud-finder"]
