package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"strings"
)

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

	for _, file := range os.Args[1:] {
		fmt.Println("Processing:", file)
		one := token.NewFileSet()
		var v visitor
		f, err := parser.ParseFile(one, file, nil, parser.AllErrors)
		if err != nil {
			fmt.Println(err)
			return
		}
		ast.Walk(v, f)
	}
}
