package ast

import "github.com/busylambda/raven/symtab"

type Expr interface{}

type Root struct {
	Name  string `json:"name"`
	Decls []Decl `json:"decls"`
}

func NewAstRoot(module string) *Root {
	return &Root{Name: module, Decls: []Decl{}}
}

type BlockStmt struct {
	Opening int    `json:"opening"`
	List    []Stmt `json:"list"`
	Closing int    `json:"closing"`
}

type FuncType struct {
	Params []FuncParam `json:"params"`
}

type FuncParam struct {
	Ident *Identifier `json:"ident"`
	Type  *Identifier `json:"type"`
}

type Identifier struct {
	Span symtab.Span `json:"span"`
	Name string      `json:"name"`
}

func NewIdentifier(name string, pos symtab.Span) *Identifier {
	return &Identifier{Name: name, Span: pos}
}

type ObjectKind int

const (
	FUNC ObjectKind = iota
	FUNC_PARAM
	FUNC_PARAM_TYPE
)

type BinOp int

const (
	PLUS BinOp = iota
	MINUS
	MULT
	DIVI
	MODU
)

type BinExpr struct {
	Op    *BinOp `json:"op"`
	Left  *Expr  `json:"left"`
	Right *Expr  `json:"right"`
}

type IdentExpr struct {
	Ident *Identifier `json:"ident"`
}

type StructExpr struct {
	Type   *Type
	Fields map[string]Expr `json:"fields"`
}
