package ast

type Stmt interface{}

type ExprStmt struct {
	Xpr Expr `json:"expr"`
}

type DeclStmt struct {
	Left  *Expr       `json:"left"`
	Right *Expr       `json:"right"`
	Type  *Identifier `json:"type"`
}
