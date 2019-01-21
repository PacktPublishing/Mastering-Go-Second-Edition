package main

import (
	"runtime"
)

type data struct {
	i, j int32
}

func main() {
	var N = 2000000
	var structure []data
	for i := 0; i < N; i++ {
		value := int32(i)
		structure = append(structure, data{value, value})
	}

	runtime.GC()
	_ = structure[0]
}
