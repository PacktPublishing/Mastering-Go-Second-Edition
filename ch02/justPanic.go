package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		panic("Not enough arguments!")
	}

	fmt.Println("Thanks for the argument(s)!")
}
