FROM golang:alpine

RUN mkdir /files
COPY kvTCP /files
WORKDIR /files

RUN go build -o /files/kvTCP kvTCP.go
ENTRYPOINT ["/files/kvTCP 80"]
