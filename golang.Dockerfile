FROM golang:1.16.2

WORKDIR /application

COPY *.go /application
COPY go.* /application

RUN go build -o server

EXPOSE 8090

CMD ["/application/server"]
