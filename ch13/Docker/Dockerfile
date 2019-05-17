FROM golang:latest

RUN mkdir /files
COPY kvTCP.go /files
WORKDIR /files

RUN go build -o /files/kvTCP kvTCP.go
ENTRYPOINT ["/files/kvTCP","80"]
