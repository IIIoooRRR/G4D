package cmd

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "g4d",
	Doc:  "Analysing a bot's commands",
	Run:  run,
}

func run(pass *analysis.Pass) (any, error) {
	for _, file := range pass.Files {
		ast.Inspect(file, func(n ast.Node) bool {
			fn, ok := n.(*ast.FuncDecl)
			if !ok {
				return true
			}
			if fn.Name == nil {
				return true
			}
			if fn.Type == nil || fn.Body == nil {
				return true
			}
			if !isCommand(fn) {
				return true
			}
			if !hasGetFn(fn) {
				pass.Reportf(fn.Pos(), "no call \"parse.GetEvent[_]\" function in %s", fn.Name.Name)
				return false
			}
			return true
		})
	}

	return nil, nil
}

func isCommand(fn *ast.FuncDecl) bool {
	if fn.Type == nil {
		return false
	}
	if fn.Type.Params == nil || fn.Type.Params.List == nil {
		return false
	}
	if fn.Type.Results == nil || fn.Type.Results.List == nil {
		return false
	}

	paramCount := len(fn.Type.Params.List)
	resultCount := len(fn.Type.Results.List)

	if fn.Recv != nil {
		if resultCount != 1 {
			return false
		}
	} else {
		if paramCount != 2 {
			return false
		}
		if resultCount != 1 {
			return false
		}
	}

	if len(fn.Type.Params.List) < 2 {
		return false
	}

	param1Type := fn.Type.Params.List[0].Type
	param2Type := fn.Type.Params.List[1].Type

	isRawEvent := isPointerToType(param1Type, "parse", "RawEvent")
	isContext := isPointerToType(param2Type, "ctx", "Context")

	resultType := fn.Type.Results.List[0].Type
	isError := isErrorType(resultType)

	return isRawEvent && isContext && isError
}

func isPointerToType(expr ast.Expr, pkgName, typeName string) bool {
	star, ok := expr.(*ast.StarExpr)
	if !ok {
		return false
	}
	sel, ok := star.X.(*ast.SelectorExpr)
	if !ok {
		return false
	}

	ident, ok := sel.X.(*ast.Ident)
	if !ok {
		return false
	}

	return ident.Name == pkgName && sel.Sel.Name == typeName
}

func isErrorType(expr ast.Expr) bool {
	if ident, ok := expr.(*ast.Ident); ok {
		return ident.Name == "error"
	}
	return false
}

func hasGetFn(fn *ast.FuncDecl) bool {

	found := false
	ast.Inspect(fn.Body, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.CallExpr:
			if sel, ok := x.Fun.(*ast.SelectorExpr); ok {
				if sel.Sel.Name == "GetEvent" {
					if ident, ok := sel.X.(*ast.Ident); ok && ident.Name == "parse" {
						found = true
						return false
					}
				}
			}
			if idx, ok := x.Fun.(*ast.IndexExpr); ok {
				if sel, ok := idx.X.(*ast.SelectorExpr); ok {
					if sel.Sel.Name == "GetEvent" {
						if ident, ok := sel.X.(*ast.Ident); ok && ident.Name == "parse" {
							found = true
							return false
						}
					}
				}
			}
		}
		return true
	})
	return found
}
