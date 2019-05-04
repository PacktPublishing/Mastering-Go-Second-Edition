package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	idf "github.com/mactsouk/protobuf"
	"net"
)

type server struct{}

var port = "8080"

func main() {
	listen, err := net.Listen("tpc", port)
	if err != nil {
		fmt.Println(err)
		return
	}

}
