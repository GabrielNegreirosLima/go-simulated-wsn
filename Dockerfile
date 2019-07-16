# Builder of binaries
FROM golang:1.8 AS builder

RUN mkdir -p /go-simulated-wsn
WORKDIR /go-simulated-wsn

COPY src/ /go-simulated-wsn/

RUN go build /go-simulated-wsn/sensor.go

# Light release
FROM alpine AS release

COPY --from=builder /go-simulated-wsn/sensor /

CMD ["/sensor"]
