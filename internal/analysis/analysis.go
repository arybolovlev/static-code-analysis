package analysis

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/tools/go/packages"
)

func Analysis(path string) *Module {
	queue := make([]*ast.FuncDecl, 0)

	cfg := &packages.Config{
		Mode: packages.NeedModule,
		Dir:  path,
	}
	init, err := packages.Load(cfg)
	if err != nil {
		log.Fatal(err)
	}
	if packages.PrintErrors(init) > 0 {
		log.Fatal("Package contains errors")
	}
	modulePath := init[0].Module.Path

	module := &Module{Path: modulePath}
	fset := token.NewFileSet()

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

	var file *ast.File
	for pkgName, pkg := range pkgs {
		file = ast.MergePackageFiles(pkg, ast.FilterFuncDuplicates)

		// Imports
		imports := make(map[string]string)
		for _, imp := range file.Imports {
			path := strings.Trim(imp.Path.Value, `"`)
			name := filepath.Base(path)
			if imp.Name != nil {
				name = imp.Name.Name
			}
			imports[name] = path
		}
		//

		p := &Package{
			Name:      pkgName,
			Imports:   imports,
			Functions: make([]*Function, 0),
			Types:     make(map[string]struct{}),
		}
		module.Packages = append(module.Packages, p)
		ast.Inspect(file, func(n ast.Node) bool {
			switch x := n.(type) {
			// *ast.FuncDecl -- functions and methods
			case *ast.FuncDecl:
				p.Functions = append(p.Functions, &Function{
					Name:     x.Name.String(),
					CalledBy: make(map[string]struct{}),
					Calls:    make(map[string]struct{}),
				})
				queue = append(queue, x)
				return false
			case *ast.TypeSpec:
				p.Types[x.Name.String()] = struct{}{}
				return false

			}
			return true
		})

		for _, q := range queue {
			inspectFunc(q, p)
		}
	}

	return module
}

func inspectFunc(b *ast.FuncDecl, p *Package) {
	caller := p.GetFunction(b.Name.String())
	ast.Inspect(b.Body, func(n ast.Node) bool {
		switch t := n.(type) {
		case *ast.CallExpr:
			switch x := t.Fun.(type) {
			// *ast.Ident -- local function
			case *ast.Ident:
				callee := p.GetFunction(x.Name)
				if x.Obj != nil {
					switch o := x.Obj.Decl.(type) {
					case *ast.AssignStmt:
						ast.Inspect(o, func(n ast.Node) bool {
							switch x := n.(type) {
							case *ast.FuncDecl:
								callee = p.GetFunction(x.Name.String())
							}
							return true
						})
					}
				}
				if callee != nil {
					callee.CalledBy[caller.Name] = struct{}{}
					caller.Calls[callee.Name] = struct{}{}
				}
				return false
			// *ast.SelectorExpr -- imported functions / methods
			case *ast.SelectorExpr:
				callee := x.Sel.Name
				pkg := p.Imports[x.X.(*ast.Ident).Name]
				if x.X.(*ast.Ident).Obj != nil {
					switch o := x.X.(*ast.Ident).Obj.Decl.(type) {
					case *ast.AssignStmt:
						ast.Inspect(o, func(n ast.Node) bool {
							switch x := n.(type) {
							case *ast.CompositeLit:
								if _, ok := p.Types[x.Type.(*ast.Ident).String()]; ok {
									pkg = x.Type.(*ast.Ident).String()
								}
								return true
							}
							return true
						})
					}
				}
				caller.Calls[fmt.Sprintf("%s.%s", pkg, callee)] = struct{}{}
				return true
			}
		case *ast.Ident:
			callee := p.GetFunction(t.Name)
			if t.Obj != nil {
				switch o := t.Obj.Decl.(type) {
				case *ast.FuncDecl:
					callee = p.GetFunction(o.Name.String())
				}
				if callee != nil {
					callee.CalledBy[caller.Name] = struct{}{}
					caller.Calls[callee.Name] = struct{}{}
				}
			}
		}
		return true
	})
}
