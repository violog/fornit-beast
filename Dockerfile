FROM golang:1.24-alpine AS builder
WORKDIR /go/src/BOT
COPY BOT/ .
RUN go build -o bot.bin main.go

FROM alpine:latest
WORKDIR /app/PHP
COPY --from=builder /go/src/BOT/bot.bin /bot.bin
EXPOSE 8080
ENTRYPOINT ["/bin/sh", "-c", "cp /bot.bin /app/PHP/bot.bin && ./bot.bin"]
