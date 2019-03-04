package main

import (
	"fmt"
)

func returnPtr(x int) *int {
	y := x * x
	return &y
}

func main() {
	sq := returnPtr(10)
	fmt.Println("sq value:", *sq)
	fmt.Println("sq memory address:", sq)
}
