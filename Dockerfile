FROM golang:1.8

RUN mkdir -p /go-simulated-wsn
WORKDIR /go-simulated-wsn

ADD . /go-simulated-wsn

RUN go run /go-simulated-wsn/src/sensor.go
