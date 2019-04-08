package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

func main() {
	arguments := os.Args

	if len(arguments) == 1 {
		fmt.Println(os.Args[0], "n1, n2, [n]")
		return
	}

	var wg sync.WaitGroup
	for _, arg := range arguments[1:] {
		n, err := strconv.Atoi(arg)
		if err != nil || n < 0 {
			fmt.Print(". ")
			continue
		}

		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			time.Sleep(time.Duration(n) * time.Second)
			fmt.Print(n, " ")
		}(n)
	}

	wg.Wait()
	fmt.Println()
}
