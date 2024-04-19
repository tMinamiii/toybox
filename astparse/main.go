package main

import (
	"go/ast"
	"go/parser"
	"go/token"
)

func main() {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, "main.go", nil, parser.AllErrors)
	if err != nil {
		panic(err)
	}
	for _, d := range node.Decls {
		ast.Print(fset, d)
		// fmt.Println() // \n したい...
	}
}
