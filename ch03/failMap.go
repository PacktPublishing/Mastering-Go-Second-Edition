package main

import (
	"fmt"
)

func main() {

	var aMap map[string]int
	// aMap := map[string]int{}
	aMap = nil
	fmt.Println(aMap)
	aMap["test"] = 1

}
