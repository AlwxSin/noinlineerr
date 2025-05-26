package noinlineerr

import (
	"go/ast"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

func NewAnalyzer() *analysis.Analyzer {
	return &analysis.Analyzer{
		Name:     "noinlineerr",
		Doc:      "Disallows inline error handling (`if err := ...; err != nil {`)",
		Run:      run,
		Requires: []*analysis.Analyzer{inspect.Analyzer},
	}
}

func run(pass *analysis.Pass) (any, error) {
	insp, ok := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	if !ok {
		return nil, nil //nolint:nilnil // nothing to return
	}

	nodeFilter := []ast.Node{
		(*ast.IfStmt)(nil),
	}

	insp.Preorder(nodeFilter, func(n ast.Node) {
		ifStmt, ok := n.(*ast.IfStmt)
		if !ok || ifStmt.Init == nil {
			return
		}

		// check if the init clause is an assignment
		assignStmt, ok := ifStmt.Init.(*ast.AssignStmt)
		if !ok {
			return
		}

		// iterate over left-hand side variables of the assignment
		for _, lhs := range assignStmt.Lhs {
			ident, ok := lhs.(*ast.Ident)
			if !ok {
				continue
			}

			// confirm type is error
			obj := pass.TypesInfo.ObjectOf(ident)
			if obj == nil || !strings.Contains(obj.Type().String(), "error") || ident.Name == "_" {
				continue
			}

			// report usage of inline error assignment
			pass.Reportf(
				ident.Pos(),
				"avoid inline error handling using `if err := ...; err != nil`; use plain assignment `err := ...`",
			)
		}
	})

	return nil, nil //nolint:nilnil // nothing to return
}
