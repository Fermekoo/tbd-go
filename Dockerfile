FROM golang:1.23.4-alpine3.20 AS builder

WORKDIR /go/src/app

COPY . .

RUN GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -o /app ./...

FROM alpine:3
COPY --from=builder /app /app

CMD ["/app"]