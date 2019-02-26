package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
)

var SIZE = 2
var GLOBAL = 0
var LOCAL = 0

type visitor int

func (v visitor) Visit(n ast.Node) ast.Visitor {
	if n == nil {
		return nil
	}

	fmt.Printf("%s%T\n", strings.Repeat("\t", int(v)), n)
	return v + 1
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Not enough arguments!")
		return
	}

	all := token.NewFileSet()

	for _, file := range os.Args[1:] {
		fmt.Println("Processing:", file)
		var v visitor
		f, err := parser.ParseFile(all, file, nil, parser.AllErrors)
		if err != nil {
			fmt.Println(err)
			return
		}
		ast.Walk(v, f)
	}
}
