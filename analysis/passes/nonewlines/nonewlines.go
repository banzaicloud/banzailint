package nonewlines

import (
	"go/ast"
	"go/token"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const Doc = `check that there are no unnecessary newlines in the code`

// nolint: gochecknoglobals,golint
var Analyzer = &analysis.Analyzer{
	Name:     "nonewlines",
	Doc:      Doc,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
	Run:      run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.BlockStmt)(nil),
		(*ast.CompositeLit)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		pos := n.Pos()
		var lbrace, rbrace token.Pos
		var size int
		var incomplete bool
		var nodeType string

		switch n := n.(type) {
		case *ast.BlockStmt:
			lbrace = n.Lbrace
			rbrace = n.Rbrace
			size = len(n.List)
			nodeType = "block statements"

		case *ast.CompositeLit:
			lbrace = n.Lbrace
			rbrace = n.Rbrace
			size = len(n.Elts)
			incomplete = n.Incomplete
			nodeType = "composite literals"
		}

		startPosition := pass.Fset.Position(lbrace)
		endPosition := pass.Fset.Position(rbrace)

		// Empty new line is allowed in an empty block statement
		// One line blocks are also valid
		if size == 0 || incomplete || startPosition.Line == endPosition.Line {
			return
		}

		nextLine := lbrace
		for pass.Fset.Position(nextLine).Line == startPosition.Line {
			nextLine++
		}

		previousLine := rbrace - 1
		for pass.Fset.Position(previousLine).Line == endPosition.Line {
			previousLine--
		}

		firstPos := pass.Fset.Position(nextLine)
		secondPos := pass.Fset.Position(nextLine + 1)

		lastButOnePos := pass.Fset.Position(previousLine - 1)
		lastPos := pass.Fset.Position(previousLine)

		if firstPos.Line != secondPos.Line || lastPos.Line != lastButOnePos.Line {
			pass.Reportf(pos, "%s should not start or end with empty lines", nodeType)
		}
	})

	return nil, nil
}
