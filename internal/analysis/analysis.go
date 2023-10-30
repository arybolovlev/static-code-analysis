package analysis

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"strings"

	"github.com/arybolovlev/static-code-analysis/internal/types"
)

func Analysis(path string) *types.Package {
	queue := make([]*ast.FuncDecl, 0)
	pkg := types.NewPackage()
	fset := token.NewFileSet()

	// move to a diff file filters.go and think about more pre-build filters
	filter := func(d os.FileInfo) bool {
		if strings.Contains(d.Name(), "_test.go") {
			return false
		}

		return true
	}

	pkgs, err := parser.ParseDir(fset, path, filter, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	for _, p := range pkgs {
		pkg.Name = p.Name
		// '_' represents a file name
		for _, file := range p.Files {
			// Imports in this file
			imp := make(map[string]string)
			for _, i := range file.Imports {
				var name string
				path := strings.Trim(i.Path.Value, `"`)
				if i.Name == nil {
					n := strings.Split(path, "/")
					name = n[len(n)-1]
				} else {
					name = i.Name.String()
				}
				imp[name] = path
			}

			// ast.Print(fset, file)
			ast.Inspect(file, func(n ast.Node) bool {
				switch x := n.(type) {
				// Either functions and methods declaration.
				// The choice depends on whether there is a receiver or not.
				// If there is a receiver, then it is a method. Otherwise, it is a function.
				case *ast.FuncDecl:
					nn := *types.NewNode(x.Name.String())
					nn.Imports = imp
					pkg.InsertNode(nn)
					queue = append(queue, x)
					return false
				case *ast.TypeSpec:
					pkg.InsertStruct(types.Struct{Name: x.Name.String()})
					return false
				}
				// Pick up next Node.
				return true
			})
		}
	}

	// fmt.Println("Queue length:", len(queue))
	for _, q := range queue {
		inspectFunc(q, pkg)
	}

	return pkg
}

func inspectFunc(b *ast.FuncDecl, s *types.Package) {
	vars := make(map[string]string)
	ast.Inspect(b.Body, func(n ast.Node) bool {
		caller := s.GetNode(b.Name.String())
		switch t := n.(type) {
		case *ast.CallExpr:
			switch x := t.Fun.(type) {
			// *ast.Ident -- local function
			case *ast.Ident:
				callee := s.GetNode(x.Name)
				// callee can be `nil` in the case of build in functions like `make`
				// need to add validation here and/or exclude build in functions from ending up in the list
				callee.CalledBy[caller.Name] = &types.Func{}
				//
				caller.Calls[callee.Name] = &types.Func{}
			// *ast.SelectorExpr -- imported functions
			case *ast.SelectorExpr:
				callee := x.Sel.Name
				pkg := x.X.(*ast.Ident).Name
				st := vars[pkg]
				caller.Calls[callee] = &types.Func{
					PkgName:    pkg,
					PkgPath:    caller.Imports[pkg],
					StructName: st,
				}
			}
		case *ast.AssignStmt:
			for i, a := range t.Lhs {
				switch x := t.Rhs[i].(type) {
				case *ast.CompositeLit:
					varName := a.(*ast.Ident).Name
					varValue := x.Type.(*ast.Ident).Name
					// fmt.Println(varName, varValue)
					vars[varName] = varValue
				}
			}
		}
		return true
	})
}
