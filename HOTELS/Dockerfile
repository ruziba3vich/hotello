FROM golang:1.23.0

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

WORKDIR /app/cmd

EXPOSE 7001

RUN go build -o main .

CMD ["/app/cmd/main"]
