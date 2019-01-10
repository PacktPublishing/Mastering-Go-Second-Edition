package main

import (
	"fmt"
	"runtime"
)

func printStats(mem runtime.MemStats) {
	runtime.ReadMemStats(&mem)
	fmt.Println("mem.Alloc:", mem.Alloc)
	fmt.Println("mem.TotalAlloc:", mem.TotalAlloc)
	fmt.Println("mem.HeapAlloc:", mem.HeapAlloc)
	fmt.Println("mem.NumGC:", mem.NumGC)
	fmt.Println("-----")
}

func main() {
	var mem runtime.MemStats
	printStats(mem)

	s := make([]*int, 10000000)
	if s == nil {
		fmt.Println("Operation failed!")
	}
	printStats(mem)

	t := make([]*int, 10000000)
	if t == nil {
		fmt.Println("Operation failed!")
	}
	printStats(mem)
}
