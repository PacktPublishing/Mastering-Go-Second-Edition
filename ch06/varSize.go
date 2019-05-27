package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"strconv"
)

var SIZE = 2
var GLOBAL = 0
var LOCAL = 0

type visitor struct {
	Package map[*ast.GenDecl]bool
}

func makeVisitor(f *ast.File) visitor {
	k1 := make(map[*ast.GenDecl]bool)
	for _, aa := range f.Decls {
		v, ok := aa.(*ast.GenDecl)
		if ok {
			k1[v] = true
		}
	}

	return visitor{k1}
}

func (v visitor) Visit(n ast.Node) ast.Visitor {
	if n == nil {
		return nil
	}

	switch d := n.(type) {
	case *ast.AssignStmt:
		if d.Tok != token.DEFINE {
			return v
		}

		for _, name := range d.Lhs {
			v.isItLocal(name)
		}
	case *ast.RangeStmt:
		v.isItLocal(d.Key)
		v.isItLocal(d.Value)
	case *ast.FuncDecl:
		if d.Recv != nil {
			v.CheckAll(d.Recv.List)
		}

		v.CheckAll(d.Type.Params.List)
		if d.Type.Results != nil {
			v.CheckAll(d.Type.Results.List)
		}
	case *ast.GenDecl:
		if d.Tok != token.VAR {
			return v
		}
		for _, spec := range d.Specs {
			value, ok := spec.(*ast.ValueSpec)
			if ok {
				for _, name := range value.Names {
					if name.Name == "_" {
						continue
					}
					if v.Package[d] {
						if len(name.Name) == SIZE {
							fmt.Printf("** %s\n", name.Name)
							GLOBAL++
						}
					} else {
						if len(name.Name) == SIZE {
							fmt.Printf("* %s\n", name.Name)
							LOCAL++
						}
					}
				}
			}
		}
	}
	return v
}

func (v visitor) isItLocal(n ast.Node) {
	identifier, ok := n.(*ast.Ident)
	if ok == false {
		return
	}

	if identifier.Name == "_" || identifier.Name == "" {
		return
	}

	if identifier.Obj != nil && identifier.Obj.Pos() == identifier.Pos() {
		if len(identifier.Name) == SIZE {
			fmt.Printf("* %s\n", identifier.Name)
			LOCAL++
		}
	}
}

func (v visitor) CheckAll(fs []*ast.Field) {
	for _, f := range fs {
		for _, name := range f.Names {
			v.isItLocal(name)
		}
	}
}

func main() {
	if len(os.Args) <= 2 {
		fmt.Println("Not enough arguments!")
		return
	}

	temp, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Using default SIZE:", SIZE)
	} else {
		SIZE = temp
	}

	var v visitor
	all := token.NewFileSet()
	for _, file := range os.Args[2:] {
		fmt.Println("Processing:", file)
		f, err := parser.ParseFile(all, file, nil, parser.AllErrors)
		if err != nil {
			fmt.Println(err)
			continue
		}

		v = makeVisitor(f)
		ast.Walk(v, f)
	}
	fmt.Printf("Local: %d, Global:%d with a length of %d.\n", LOCAL, GLOBAL, SIZE)
}
