FROM golang:1.24-alpine AS builder
WORKDIR /go/src/BOT
COPY BOT/ .
RUN go build -o bot main.go

FROM alpine:latest
WORKDIR /app/PHP
COPY --from=builder /go/src/BOT/bot ./bot
EXPOSE 8080
ENTRYPOINT ["./bot"]
#ENTRYPOINT ["/bin/sh", "-c", "echo test && ls -l /app && cp /app/bot /app/PHP/bot && cd /app/PHP && ls -l /app/PHP && ls -l ."]

# not tested
#FROM php:7.3-cli
#WORKDIR /var/www/html
#EXPOSE 8000
#ENTRYPOINT ["php", "-S", "0.0.0.0:8000", "index.php"]
