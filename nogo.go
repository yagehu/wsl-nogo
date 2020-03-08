package wsl

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "wsl",
	Doc:  "unscientifically enforces vertical whitespace",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	p := NewProcessor()

	for _, f := range pass.Files {
		p.fileSet = pass.Fset
		p.file = f

		for _, d := range p.file.Decls {
			switch v := d.(type) {
			case *ast.FuncDecl:
				p.parseBlockBody(v.Name, v.Body)
			case *ast.GenDecl:
				// `go fmt` will handle proper spacing for GenDecl such as imports,
				// constants etc.
			default:
				p.addWarning(warnTypeNotImplement, d.Pos(), v)
			}
		}

		for _, result := range p.result {
			pass.Reportf(f.Package, result.String())
		}
	}

	return nil, nil
}
