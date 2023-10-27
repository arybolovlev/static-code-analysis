package analysis

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"strings"
)

func Analysis(path string) *NodesSet {
	itemSet := NewNodeSet()
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
					if v := itemSet.GetNode(x.Name.String()); v == nil {
						itemSet.InsertNode(Node{
							Name:   x.Name.String(),
							Impact: make(map[*Node]struct{}),
						})
					}
					// inspectFunc(itemSet, x)
					return false
				}
				return true
			})
		}
	}

	return itemSet
}

// func inspectFunc(s *NodesSet, b *ast.FuncDecl) {
// 	ast.Inspect(b.Body, func(n ast.Node) bool {
// 		if n == nil {
// 			return false
// 		}
// 		switch t := n.(type) {
// 		case *ast.CallExpr:
// 			switch x := t.Fun.(type) {
// 			// *ast.Ident -- local function
// 			case *ast.Ident:
// 				fnName := x.Name
// 				v := s.GetNode(fnName)
// 				if v == nil {
// 					s.InsertNode(Node{
// 						Name:   fnName,
// 						Impact: make(map[*Node]struct{}),
// 					})
// 				}
// 				if i := s.GetNode(fnName); i != nil {
// 					i.Impact[b.Name.String()] = struct{}{}
// 				}
// 				return false
// 			// *ast.SelectorExpr -- imported functions
// 			case *ast.SelectorExpr:
// 				return false
// 			}
// 		}
// 		return true
// 	})
// }
