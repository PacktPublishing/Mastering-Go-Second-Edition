FROM golang:alpine

RUN mkdir /files
COPY hw.go /files
WORKDIR /files

RUN go get github.com/lib/pq
RUN go build -o /files/hw hw.go

ENTRYPOINT ["/bin/bash"]
