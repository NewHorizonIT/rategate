# Stage 1:
FROM golang:1.25.0-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/gateway/main.go

# Stage 2:
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .
RUN apk add --no-cache ca-certificates

RUN user -D user
USER user

EXPOSE 8080
CMD ["./main"]