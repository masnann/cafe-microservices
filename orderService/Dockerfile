# Menggunakan gambar dasar golang untuk build stage
FROM golang:1.16-alpine AS build

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy seluruh kode sumber
COPY . .

# Build aplikasi
RUN go build -o order-service .

# Stage kedua, menggunakan gambar dasar alpine untuk menjalankan aplikasi
FROM alpine:latest

# Set working directory
WORKDIR /root/

# Copy binary dari build stage
COPY --from=build /app/order-service .

# Copy .env file
COPY .env .env

# Expose port yang digunakan
EXPOSE 8080

# Jalankan aplikasi
CMD ["./order-service"]
