package main

import (
	"flag"

	"github.com/arybolovlev/static-code-analysis/internal/analysis"
)

func main() {
	var path string
	flag.StringVar(&path, "path", ".", "-path=/path")

	flag.Parse()

	analysis.Analysis(path)
}
