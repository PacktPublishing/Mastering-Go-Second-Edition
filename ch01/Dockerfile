FROM golang:alpine

RUN mkdir /files
COPY hw.go /files
WORKDIR /files

RUN go build -o /files/hw hw.go
ENTRYPOINT ["/files/hw"]
