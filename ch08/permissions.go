package main

import (
	"fmt"
	"os"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Printf("usage: permissions filename\n")
		return
	}

	filename := arguments[1]
	info, _ := os.Stat(filename)
	mode := info.Mode()
	fmt.Println(filename, "mode is", mode.String()[1:10])
}
