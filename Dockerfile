## Build Stage
FROM golang:1.17.0-alpine3.14 as builder

WORKDIR /go/src/github.com/backup-exporter
COPY . .

RUN go get ./..
RUN go build -o /backup-exporter ./

## Runtime Stage
FROM alpine:3.14

COPY --from=builder /backup-exporter/exporter /backup-exporter/exporter
RUN chmod +x /backup-exporter/exporter

CMD /backup-exporter/exporter