#FROM golang:1.24-alpine AS builder
#WORKDIR /go/src/BOT
#COPY BOT/ .
#RUN go build -o BOT main.go

#FROM alpine:latest
#WORKDIR /app/PHP
#COPY PHP/ .
#COPY --from=builder /go/src/BOT/bot .
# COPY --from=builder /go/src/BOT ../.
#EXPOSE 8080
#ENTRYPOINT ["./BOT"]

# not tested
FROM php:7.3-cli
WORKDIR /var/www/html
EXPOSE 8000
ENTRYPOINT ["php", "-S", "0.0.0.0:8000", "index.php"]
