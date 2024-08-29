FROM golang:1.23.0 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o booking-service .

FROM debian:bullseye-slim

COPY --from=builder /app/booking-service /bin/booking-service

EXPOSE 7002

CMD ["/bin/booking-service"]
