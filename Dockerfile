FROM golang:1.16-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /app/bin main.go

# Path: Dockerfile

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/bin /app/bin

CMD ["/app/bin"]
