package analysis

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"strings"
)

func Analysis(path string) *ItemSet {
	itemSet := NewItemSet()
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
				if n == nil {
					return false
				}
				switch x := n.(type) {
				// *ast.FuncDecl -- functions and methods
				case *ast.FuncDecl:
					if v := itemSet.Item(x.Name.String()); v == nil {
						itemSet.InsertItem(Item{
							Name:   x.Name.String(),
							Impact: make(map[string]struct{}),
						})
					}
					inspectFunc(itemSet, x)
					return false
				}
				return true
			})
		}
	}

	return itemSet
}

func inspectFunc(s *ItemSet, b *ast.FuncDecl) {
	ast.Inspect(b.Body, func(n ast.Node) bool {
		if n == nil {
			return false
		}
		switch t := n.(type) {
		case *ast.CallExpr:
			switch x := t.Fun.(type) {
			// *ast.Ident -- local function
			case *ast.Ident:
				fname := x.Name
				v := s.Item(fname)
				if v == nil {
					s.InsertItem(Item{
						Name:   fname,
						Impact: make(map[string]struct{}),
					})
				}
				if i := s.Item(fname); i != nil {
					i.Impact[b.Name.String()] = struct{}{}
				}
				return false
			// *ast.SelectorExpr -- imported functions
			case *ast.SelectorExpr:
				return false
			}
		}
		return true
	})
}
