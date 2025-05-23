package noinlineerr

import (
	"go/ast"
	"strings"

	"golang.org/x/tools/go/analysis"
)

type analyzer struct{}

func NewAnalyzer() *analysis.Analyzer {
	a := &analyzer{}

	return &analysis.Analyzer{
		Name: "noinlineerr",
		Doc:  "Dissallows inline error handling using `if err := ...; err != nil",
		Run:  a.run,
	}
}

func (a *analyzer) run(pass *analysis.Pass) (any, error) {
	for _, file := range pass.Files {
		ast.Inspect(file, func(n ast.Node) bool {
			// look for if statements with an Init clause
			ifStmt, ok := n.(*ast.IfStmt)
			if !ok || ifStmt.Init == nil {
				return true
			}

			// check if the init cluase is an assignment
			assignStmt, ok := ifStmt.Init.(*ast.AssignStmt)
			if !ok {
				return true
			}

			// iterate over left-hand side variables of the assignment
			for _, lhs := range assignStmt.Lhs {
				ident, ok := lhs.(*ast.Ident)
				if !ok {
					continue
				}

				// confirm type is error
				obj := pass.TypesInfo.ObjectOf(ident)
				if obj == nil || !strings.Contains(obj.Type().String(), "error") {
					continue
				}

				// report usage of inline error assignment
				pass.Reportf(
					ident.Pos(),
					"avoid inline error handling using `if err := ...; err != nil; use plain assignment `err := ...",
				)
			}
			return true
		})
	}
	return nil, nil //nolint:nilnil // nothing to return
}
