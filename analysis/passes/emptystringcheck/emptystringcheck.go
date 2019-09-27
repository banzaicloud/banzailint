package emptystringcheck

import (
	"go/ast"
	"go/token"
	"go/types"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const Doc = `check that empty strings are not checked for by comparing their length to 0`

// nolint: gochecknoglobals,golint
var Analyzer = &analysis.Analyzer{
	Name:     "emptystringcheck",
	Doc:      Doc,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
	Run:      run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.BinaryExpr)(nil),
	}
	inspect.Preorder(nodeFilter, func(n ast.Node) {
		binExpr := n.(*ast.BinaryExpr)
		if isComparisionOp(binExpr.Op) {
			if isLenOfString(pass.TypesInfo, binExpr.X) && isZero(binExpr.Y) {
				pass.Reportf(binExpr.Pos(), "should not check for empty string by length")
			} else if isLenOfString(pass.TypesInfo, binExpr.Y) && isZero(binExpr.X) {
				pass.Reportf(binExpr.Pos(), "should not check for empty string by length")
			}
		}
	})
	return nil, nil
}

func isLenOfString(typesInfo *types.Info, expr ast.Expr) bool {
	if callExpr, ok := expr.(*ast.CallExpr); ok {
		if ident, ok := callExpr.Fun.(*ast.Ident); ok && ident.Name == "len" && len(callExpr.Args) == 1 {
			if typ := typesInfo.TypeOf(callExpr.Args[0]); typ != nil {
				if basic, ok := typ.(*types.Basic); ok {
					return basic.Kind() == types.String || basic.Kind() == types.UntypedString
				}
			}
		}
	}
	return false
}

func isZero(expr ast.Expr) bool {
	basicLit, ok := expr.(*ast.BasicLit)
	return ok && basicLit.Kind == token.INT && basicLit.Value == "0"
}

func isComparisionOp(tok token.Token) bool {
	switch tok {
	case token.EQL, token.LSS, token.GTR, token.NEQ, token.LEQ, token.GEQ:
		return true
	}
	return false
}
