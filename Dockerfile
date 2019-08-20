FROM golang:latest

WORKDIR /go/src
ADD . /go

CMD ["go", "run", "main.go"]
