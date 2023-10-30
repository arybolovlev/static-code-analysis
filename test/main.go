package main

import (
	"fmt"

	cd "github.com/arybolovlev/static-code-analysis/test/cmd"
)

func main() {
	cd.Run()
	a()
	fmt.Println("Hello")
}
