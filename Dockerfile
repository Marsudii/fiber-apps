# Gunakan image Go sebagai build stage
FROM golang:1.21 AS builder

# Set work directory dalam container
WORKDIR /app

# Copy semua file ke dalam container
COPY . .

# Download module Go dan build aplikasi
RUN go mod tidy
RUN go build -o app

# Stage kedua menggunakan image ringan
FROM alpine:latest

WORKDIR /root/

# Copy aplikasi dari builder stage
COPY --from=builder /app/app .

# Set port aplikasi
EXPOSE 8080

# Jalankan aplikasi
CMD ["./app"]