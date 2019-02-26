package main

import (
	"fmt"
	"go/scanner"
	"go/token"
	"io/ioutil"
	"os"
)

var KEYWORD = "var"
var COUNT = 0

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Not enough arguments!")
		return
	}

	for _, file := range os.Args[1:] {
		fmt.Println("Processing:", file)
		f, err := ioutil.ReadFile(file)
		if err != nil {
			fmt.Println(err)
			return
		}
		one := token.NewFileSet()
		files := one.AddFile(file, one.Base(), len(f))

		var myScanner scanner.Scanner
		myScanner.Init(files, f, nil, scanner.ScanComments)

		localCount := 0
		for {
			_, tok, lit := myScanner.Scan()
			if tok == token.EOF {
				break
			}

			if lit == KEYWORD {
				COUNT++
				localCount++
			}
		}
		fmt.Printf("Found _%s_ %d times\n", KEYWORD, localCount)
	}
	fmt.Printf("Found _%s_ %d times in total\n", KEYWORD, COUNT)
}
