FROM golang:alpine

WORKDIR /usr/src/app

COPY . .

RUN go build -o main ./cmd/main

ENTRYPOINT ["/usr/src/app/main"]

EXPOSE 3003
