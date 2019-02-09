package main

import (
	"fmt"
)

func main() {
	foo := [2][4][7]float64{}
	x := len(foo)       // x == 2
	y := len(foo[0])    // y == 4
	z := len(foo[0][0]) // z == 8
}
