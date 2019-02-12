package main

import (
	"fmt"
)

func main() {
	array := [12][4][7][10]float64{}
	x := len(array)
	y := len(array[0])
	z := len(array[0][0])
	w := len(array[0][0][0])
	fmt.Println("x:", x, "y:", y, "z:", z, "w:", w)
}
