package symtab

import "fmt"

type Symbol struct {
	Span Span       `json:"span"`
	Type SymbolType `json:"type"`
	Ut   UsageType  `json:"usage_type"`
}

type SymOD struct {
	Span  Span
	Usage UsageType
}

func NewSymbol(span Span, ut UsageType) Symbol {
	return Symbol{
		Span: span,
		Ut:   ut,
	}
}

type UsageType int

func (ut *UsageType) String() string {
	switch *ut {
	case FUNC:
		return "FUNC"
	case TYPE:
		return "TYPE"
	case VAR:
		return "VAR"
	case PARAM:
		return "PARAM"
	case STRUCT:
		return "STRUCT"
	default:
		return "UNKNOWN"
	}
}

const (
	FUNC UsageType = iota
	TYPE
	VAR
	PARAM
	STRUCT
)

// This is a placeholder for now
type SymbolType int

type SymbolTable struct {
	Parent    *SymbolTable            `json:"parent"`
	Children  map[string]*SymbolTable `json:"children"`
	DeclTable map[string]Symbol       `json:"decls"`
	UsgeTable map[string]Symbol       `json:"usages"`
}

func NewSymTabRoot() *SymbolTable {
	return &SymbolTable{
		Parent:    nil,
		Children:  make(map[string]*SymbolTable),
		DeclTable: make(map[string]Symbol),
		UsgeTable: make(map[string]Symbol),
	}
}

func (st *SymbolTable) InsertDecl(s Symbol, name string) {
	st.DeclTable[fmt.Sprintf("%s:%s", name, s.Ut.String())] = s
}
func (st *SymbolTable) InsertUsage(s Symbol, name string) {
	st.UsgeTable[fmt.Sprintf("%s:%s", name, s.Ut.String())] = s
}
func (st *SymbolTable) AllocateChlid(name string) {
	st.Children[name] = &SymbolTable{
		Parent:    st,
		Children:  make(map[string]*SymbolTable),
		DeclTable: make(map[string]Symbol),
		UsgeTable: make(map[string]Symbol),
	}
}
