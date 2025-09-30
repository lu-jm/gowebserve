FROM golang:1.24 AS builder

WORKDIR /app

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o my-go-app main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/my-go-app /app/.env /app/

EXPOSE 8080

CMD ["/app/my-go-app"]