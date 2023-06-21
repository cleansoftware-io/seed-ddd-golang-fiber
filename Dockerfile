FROM golang:1.20.3-buster

WORKDIR /app

COPY . .

RUN go build -o app cmd/main.go

RUN go test ./...

EXPOSE 3000

CMD ["./app"]
