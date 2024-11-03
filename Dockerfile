FROM  golang:1.22.8

WORKDIR /go/app

RUN go install github.com/matryer/moq@v0.4.0