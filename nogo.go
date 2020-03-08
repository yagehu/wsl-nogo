package wsl

import (
	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "wsl",
	Doc:  "unscientifically enforces vertical whitespace",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	pass.Reportf(pass.Files[0].Package, "test")
	return nil, nil
}
