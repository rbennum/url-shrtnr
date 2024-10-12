# Build stage
FROM golang:1.22 AS builder
WORKDIR /build
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o shortener-service .

# Final stage
FROM alpine:3.19
WORKDIR /app
COPY --from=builder /build/shortener-service .
COPY .env.local .env.local
COPY etc/migrations etc/migrations
RUN chmod +x /app/shortener-service && ls -l /app
EXPOSE 32123

CMD ["/app/shortener-service"]
