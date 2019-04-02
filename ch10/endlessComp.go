package main

import (
	"fmt"
	"runtime"
)

func main() {
	var i byte
	go func() {
		for i = 0; i <= 255; i++ {
		}
	}()
	fmt.Println("Leaving goroutine!")
	runtime.Gosched()
	runtime.GC()

	fmt.Println("Good bye!")
}
