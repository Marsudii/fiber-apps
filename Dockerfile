# Build stage
FROM golang:1.21 AS builder
ENV GOTOOLCHAIN=auto

WORKDIR /app

COPY . .

# Build aplikasi agar bisa dijalankan di Linux Alpine
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app

# Runtime stage
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/app .

EXPOSE 8080

CMD ["./app"]