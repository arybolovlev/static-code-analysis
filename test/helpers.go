package main

import (
	"fmt"

	"github.com/arybolovlev/static-code-analysis/test/cmd"
)

func a() {
	a := 1
	fmt.Println(a)
	b()
	c()
}

func b() {
	cmd.Run()
	c()
}

func c() {
	f := fullName{firstName: "Sasha"}
	// f := newFullName("Sasha")
	fmt.Println(f.getFirstName())
	const (
		version = "0.0.2"
	)
}

type fullName struct {
	firstName string
}

func newFullName(n string) fullName {
	return fullName{firstName: n}
}

func (fn *fullName) getFirstName() string {
	return fn.firstName
}

var (
	hello string
)

const (
	version = "0.0.1"
)
