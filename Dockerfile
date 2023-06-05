FROM golang:1.20-buster

WORKDIR /app

COPY . .

RUN go build -o app cmd/main.go

EXPOSE 3000

CMD ["./app"]
