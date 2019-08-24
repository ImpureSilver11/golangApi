FROM golang:latest

WORKDIR /go/src
ADD . /go

CMD ["go", "run", "main.go"]

RUN go get github.com/lib/pq
