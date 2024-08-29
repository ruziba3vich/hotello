FROM golang:1.23.0 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

FROM debian:bullseye-slim

RUN apt-get update && apt-get install -y \
    ca-certificates \
    libglib2.0-0 \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /root/

COPY --from=builder /app/main .

EXPOSE 7000

CMD ["./main"]
