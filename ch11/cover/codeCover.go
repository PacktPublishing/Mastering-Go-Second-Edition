package main

import (
	"fmt"
)

func fibo1(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fibo1(n-1) + fibo1(n-2)
	}
}

func fibo2(n int) int {
	if n >= 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fibo1(n-1) + fibo1(n-2)
	}
}

func main() {
	fmt.Println(fibo1(40))
	fmt.Println(fibo2(40))
}
