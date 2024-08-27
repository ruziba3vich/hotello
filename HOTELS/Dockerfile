FROM golang:1.23.0 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o hotels_service ./cmd/main

FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=builder /app/hotels_service /hotels_service

EXPOSE 7001

CMD ["/hotels_service"]
