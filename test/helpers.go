package main

import (
	"fmt"
)

func a() {
	a := 1
	fmt.Println(a)
	b()
	c()
}

func b() {
	c()
}

func c() {}
