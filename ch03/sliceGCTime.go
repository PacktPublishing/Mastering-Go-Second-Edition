package main

import (
	"fmt"
	"runtime"
	"time"
)

type data struct {
	i, j int
}

func main() {
	var N = 40000000
	var structure []data

	for i := 0; i < N; i++ {
		value := int(i)
		structure = append(structure, data{value, value})
	}

	start := time.Now()
	runtime.GC()
	duration := time.Since(start)

	fmt.Println("It took GC()", duration, "to finish")
	_ = structure[0]
}
