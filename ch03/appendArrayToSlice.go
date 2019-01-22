package main

import (
	"fmt"
)

func main() {
	s := []int{1, 2, 3}
	a := [3]int{4, 5, 6}

	ref := a[:]
	fmt.Println(ref)
	t := append(s, ref...)
	fmt.Println(t)
	s = append(s, ref...)
	fmt.Println(s)
}
