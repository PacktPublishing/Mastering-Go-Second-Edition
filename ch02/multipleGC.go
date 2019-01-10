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

// Create a structure

func main() {
	var mem runtime.MemStats
	printStats(mem)

	var s *int = nil
	for i := 0; i < 10000000; i++ {
		s = nil
	}
	printStats(mem)

	for i := 0; i < 10000000; i++ {
		s = nil
	}
	printStats(mem)

	fmt.Println(s)
}
