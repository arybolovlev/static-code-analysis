package analysis

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"strings"
)

func Analysis(path string) *Nodes {
	queue := make([]*ast.FuncDecl, 0)
	nodesSet := NewNode()
	fset := token.NewFileSet()

	filter := func(d os.FileInfo) bool {
		if strings.Contains(d.Name(), "_test.go") {
			return false
		}

		return true
	}

	dir, err := parser.ParseDir(fset, path, filter, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	for _, d := range dir {
		// _ -- filename
		for _, file := range d.Files {
			ast.Inspect(file, func(n ast.Node) bool {
				switch x := n.(type) {
				// *ast.FuncDecl -- functions and methods
				case *ast.FuncDecl:
					nodesSet.InsertNode(Node{
						Name:     x.Name.String(),
						CalledBy: make(map[string]struct{}),
						Calls:    make(map[string]struct{}),
					})
					queue = append(queue, x)
					return false
				}
				return true
			})
		}
	}

	for _, q := range queue {
		inspectFunc(q, nodesSet)
	}

	return nodesSet
}

func inspectFunc(b *ast.FuncDecl, s *Nodes) {
	ast.Inspect(b.Body, func(n ast.Node) bool {
		caller := s.GetNode(b.Name.String())
		switch t := n.(type) {
		case *ast.CallExpr:
			switch x := t.Fun.(type) {
			// *ast.Ident -- local function
			case *ast.Ident:
				callee := s.GetNode(x.Name)
				callee.CalledBy[caller.Name] = struct{}{}

				caller.Calls[callee.Name] = struct{}{}
				return false
			// *ast.SelectorExpr -- imported functions
			case *ast.SelectorExpr:
				callee := x.Sel.Name
				pkg := x.X.(*ast.Ident).Name
				caller.Calls[fmt.Sprintf("%s.%s", pkg, callee)] = struct{}{}
				return false
			}
		}
		return true
	})
}
