package main

import (
	"fmt"
	"os"
)

func function(i int) int {
	return i * i
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Need at least one argument.")
		return
	}

	i := 5
	fmt.Println("Debugging with Delve")
	fmt.Println(i)
	fmt.Println(function(i))

	for arg, _ := range os.Args[1:] {
		fmt.Println(arg)
	}
}
