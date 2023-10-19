package parser

type Decl interface{}
type Expr interface{}
type Stmt interface{}

type Root struct {
	Name  string
	Decls []Decl
}

func NewAstRoot(module string) *Root {
	return &Root{Name: module, Decls: []Decl{}}
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
	Params []FuncParam
}

type FuncParam struct {
	Ident *Identifier
	Type  *Identifier
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
	FUNC_PARAM
	FUNC_PARAM_TYPE
)

type ExprStmt struct {
	Xpr Expr
}

type DeclStmt struct {
}

type DeclAssignStmt struct {
}
