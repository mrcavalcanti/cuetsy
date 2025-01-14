package ts

import (
	"fmt"
	"runtime"

	"github.com/mrcavalcanti/cuetsy/ts/ast"
)

type (
	File = ast.File
	Node = ast.Node
	Decl = ast.Decl
	Expr = ast.Expr
)

func Ident(name string) ast.Ident {
	return ast.Ident{Name: name}
}

func Names(names ...string) ast.Names {
	idents := make(ast.Idents, len(names))
	for i, n := range names {
		idents[i] = Ident(n)
	}

	return ast.Names{
		Idents: idents,
	}
}

func Union(elems ...Expr) Expr {
	switch len(elems) {
	case 0:
		return nil
	case 1:
		return elems[0]
	}

	var U Expr = elems[0]
	for _, e := range elems[1:] {
		U = ast.BinaryExpr{
			Op: "|",
			X:  U,
			Y:  e,
		}
	}

	return ast.ParenExpr{Expr: U}
}

func Export(decl ast.Decl) Decl {
	return ast.ExportStmt{Decl: decl}
}

func Raw(data string) ast.Raw {
	pc, file, no, ok := runtime.Caller(1)
	details := runtime.FuncForPC(pc)
	if ok && details != nil {
		fmt.Printf("fix: ts.Raw used by %s at %s#%d\n", details.Name(), file, no)
	}

	return ast.Raw{Data: data}
}

func Object(fields map[string]Expr) Expr {
	elems := make([]ast.KeyValueExpr, 0, len(fields))
	for k, v := range fields {
		elems = append(elems, ast.KeyValueExpr{
			Key:   Ident(k),
			Value: v,
		})
	}
	return ast.ObjectLit{Elems: elems}
}

func List(elems ...Expr) Expr {
	return ast.ListLit{Elems: elems}
}

func Null() Expr {
	return Ident("null")
}

func Str(s string) Expr {
	return ast.Str{Value: s}
}

// TODO: replace with generic num?
func Int(i int64) Expr {
	return ast.Num{N: i}
}
func Float(f float64) Expr {
	return ast.Num{N: f}
}

func Bool(b bool) Expr {
	if b {
		return Ident("true")
	}
	return Ident("false")
}
