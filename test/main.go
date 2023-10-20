package main

import "fmt"

type name struct {
	firstName string
}

func (n *name) printFirstName() {
	fmt.Println(n.firstName)
}

func a() {
	a := 1
	fmt.Println(a)
	b()
	c()
}

func b() {
	c()
}

func c() {
	n := name{firstName: "Sasha"}
	n.printFirstName()
}
