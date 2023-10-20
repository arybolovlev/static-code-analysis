package analysis

import (
	"encoding/json"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
)

type ItemSet struct {
	Items []*Item `json:"items"`
}

type Item struct {
	Name   string   `json:"name"`
	Impact []string `json:"impact"`
}

func NewItemSet() *ItemSet {
	return &ItemSet{}
}

func (s *ItemSet) InsertItem(i Item) {
	s.Items = append(s.Items, &i)
}

func (s *ItemSet) Item(n string) *Item {
	for _, i := range s.Items {
		if i.Name == n {
			return i
		}
	}

	return nil
}

func (s *ItemSet) HasItem(n string) *Item {
	for _, i := range s.Items {
		if i.Name == n {
			return i
		}
	}

	return nil
}

func Analysis(path string) {
	itemSet := NewItemSet()
	fset := token.NewFileSet()
	dir, err := parser.ParseDir(fset, path, nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}
	for _, d := range dir {
		for _, file := range d.Files {
			ast.Inspect(file, func(n ast.Node) bool {
				if n == nil {
					return false
				}
				switch x := n.(type) {
				case *ast.FuncDecl:
					if v := itemSet.HasItem(x.Name.String()); v == nil {
						itemSet.InsertItem(Item{
							Name: x.Name.String(),
						})
					}
					inspectFunc(itemSet, x)
					return false
				}
				return true
			})
		}
	}

	o, err := json.Marshal(itemSet)
	if err != nil {
		log.Fatalf(err.Error())
		panic(err)
	}
	fmt.Println(string(o))
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
				v := s.HasItem(fname)
				if v == nil {
					s.InsertItem(Item{
						Name: fname,
					})
				}
				if i := s.Item(fname); i != nil {
					i.Impact = append(i.Impact, b.Name.String())
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
