FROM golang:1.23-alpine AS builder

WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o product-srv .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/product-srv .
COPY etc/ etc/

EXPOSE 8083
ENTRYPOINT ["./product-srv", "service", "-c", "etc/product.yaml"]
