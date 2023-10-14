FROM golang:1.21 AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o ecommerce
FROM gcr.io/distroless/base
WORKDIR /app
COPY --from=builder /app/ecommerce .
CMD ["./ecommerce"]
