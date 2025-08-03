FROM golang:1.24 AS builder

WORKDIR /app

COPY . .

RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -o /stress-test-cli ./cmd/server/

FROM alpine:3.18

RUN apk add --no-cache ca-certificates

COPY --from=builder /stress-test-cli /usr/local/bin/

ENTRYPOINT ["/usr/local/bin/stress-test-cli"]