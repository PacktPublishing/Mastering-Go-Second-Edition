package main

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
)

func main() {
	myVersion := runtime.Version()
	major := strings.Split(myVersion, ".")[0][2]
	minor := strings.Split(myVersion, ".")[1]
	m1, _ := strconv.Atoi(string(major))
	m2, _ := strconv.Atoi(minor)

	if m1 == 1 && m2 < 8 {
		fmt.Println("Need Go version 1.8 or higher!")
		return
	}

	fmt.Println("You are using Go version 1.8 or higher!")
}
