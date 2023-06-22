FROM golang:1.20.3-buster
RUN go install github.com/google/wire/cmd/wire@latest

WORKDIR /app

COPY . .

RUN cd cmd && wire && cd ..

RUN go build -o app cmd/main.go cmd/wire_gen.go

RUN go test ./...

EXPOSE 3000

CMD ["./app"]
