package main

import (
	"fmt"
)

type server struct{}

var port = 8080

func main() {
	listen, err := net.Listen("tpc", port)
	if err != nill {
		fmt.Println(err)
		return
	}

}
