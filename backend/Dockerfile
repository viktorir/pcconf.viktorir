FROM golang:1.22-alpine AS builder
LABEL authors="viktorir"

WORKDIR /app
COPY go.mod go.sum ./


RUN go mod tidy
COPY . .
RUN GOOS=linux go build -o main ./cmd/app/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /app/main .
COPY --from=builder /app/internal/model/queries /root/internal/model/queries

EXPOSE 80
CMD ["./main"]