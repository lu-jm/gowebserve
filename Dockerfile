# FROM golang:1.24 AS builder

# WORKDIR /app

# COPY . .

# RUN go build -o gomain main.go

FROM alpine:latest

WORKDIR /app

COPY gomain .
COPY .env .

EXPOSE 8080

CMD ["./gomain"]