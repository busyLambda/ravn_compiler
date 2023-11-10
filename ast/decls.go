package ast

type Decl interface{}

type FuncDecl struct {
	Name *Identifier `json:"name"`
	Type *FuncType   `json:"type"`
	Body *BlockStmt  `json:"body"`
}

type StructDecl struct {
	Name   *Identifier     `json:"name"`
	Fields map[string]Type `json:"fields"`
}
