package main

import (
	"flag"
	"fmt"
	"sync"
	"sync/atomic"
)

type atomicCounter struct {
	val int64
}

func (c *atomicCounter) Value() int64 {
	return atomic.LoadInt64(&c.val)
}

func main() {
	minusX := flag.Int("x", 100, "Goroutines")
	minusY := flag.Int("y", 200, "Value")
	flag.Parse()
	X := *minusX
	Y := *minusY

	var waitGroup sync.WaitGroup
	counter := atomicCounter{}
	for i := 0; i < X; i++ {
		waitGroup.Add(1)
		go func(no int) {
			defer waitGroup.Done()
			for i := 0; i < Y; i++ {
				atomic.AddInt64(&counter.val, 1)
			}
		}(i)
	}

	waitGroup.Wait()
	fmt.Println(counter.Value())
}
