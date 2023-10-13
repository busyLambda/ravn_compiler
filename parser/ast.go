package parser

type Decl interface{}
type Expr interface{}
type Stmt interface{}

type Root struct {
	Name  *Identifier
	Decls []Decl
}

type FuncDecl struct {
	Name *Identifier
	Type *FuncType
	Body *BlockStmt
}

type BlockStmt struct {
	Opening int
	List    []Stmt
	Closing int
}

type FuncType struct {
	Params *FieldList
}

type FieldList struct {
	Opening int
	Closing int
}

type Identifier struct {
	Span Span
	Name string
	Obj  *Object
}

func NewIdentifier(name string, pos Span, obj Object) *Identifier {
	return &Identifier{Name: name, Span: pos, Obj: &obj}
}

type Span struct {
	Start int
	End   int
}

type Object struct {
	Kind ObjectKind
	Name string
}

type ObjectKind int

const (
	FUNC ObjectKind = iota
)

type ExprStmt struct {
	Xpr Expr
}
