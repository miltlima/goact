package dockerfiles

const Golang = `FROM golang:latest

WORKDIR /app

COPY . /app

RUN groupadd -r golangapp && useradd -r -g golangapp golangapp
USER golangapp

RUN go build -o app

CMD ["./app"]`
