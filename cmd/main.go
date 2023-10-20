package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"

	"github.com/arybolovlev/static-code-analysis/internal/analysis"
)

func main() {
	var path string
	flag.StringVar(&path, "path", ".", "-path=/path")
	var jsonOutput bool
	flag.BoolVar(&jsonOutput, "json", true, "-json")
	flag.Parse()

	o := analysis.Analysis(path)

	if jsonOutput {
		jo, err := json.Marshal(o)
		if err != nil {
			log.Fatalf(err.Error())
		}
		fmt.Println(string(jo))
	}
}
