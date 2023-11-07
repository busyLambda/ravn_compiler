package parser

import "github.com/busylambda/raven/symtab"

type Decl interface{}
type Expr interface{}
type Stmt interface{}

type Root struct {
	Name  string `json:"name"`
	Decls []Decl `json:"decls"`
}

func NewAstRoot(module string) *Root {
	return &Root{Name: module, Decls: []Decl{}}
}

type FuncDecl struct {
	Name *Identifier `json:"name"`
	Type *FuncType   `json:"type"`
	Body *BlockStmt  `json:"body"`
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
	Obj  *Object     `json:"obj"`
}

func NewIdentifier(name string, pos symtab.Span, obj Object) *Identifier {
	return &Identifier{Name: name, Span: pos, Obj: &obj}
}

type Object struct {
	Kind ObjectKind `json:"kind"`
	Name string     `json:"name"`
}

type ObjectKind int

const (
	FUNC ObjectKind = iota
	FUNC_PARAM
	FUNC_PARAM_TYPE
)

type ExprStmt struct {
	Xpr Expr `json:"expr"`
}

type DeclStmt struct {
	Left  *Expr       `json:"left"`
	Right *Expr       `json:"right"`
	Type  *Identifier `json:"type"`
}

type IdentExpr struct {
	Ident *Identifier `json:"ident"`
}

type StructDecl struct {
	Name   *Identifier     `json:"name"`
	Fields map[string]Type `json:"fields"`
}

type StructExpr struct {
	Type   *Type
	Fields map[string]Expr `json:"fields"`
}
