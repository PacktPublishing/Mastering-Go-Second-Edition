package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
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

	switch t := n.(type) {
	case *ast.AssignStmt:
		if t.Tok != token.DEFINE {
			return v + 1
		}
		for _, name := range t.Lhs {
			fmt.Println("name:", name)
		}
	case *ast.RangeStmt:
		fmt.Println("t.Tok:", t.Key)
		if len(string(t.Tok)) == SIZE {
			LOCAL++
		}
	case *ast.FuncDecl:

	}

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
			fmt.Println("Parsing error:", err)
			return
		}
		ast.Walk(v, f)
	}

	fmt.Printf("Local: %d, Global:%d with size %d\n", LOCAL, GLOBAL, SIZE)
}
