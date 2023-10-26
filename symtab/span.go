package symtab

type Span struct {
	Start int `json:"start"`
	End   int `json:"end"`
}

func NewSpan(start int, end int) Span {
	return Span{
		Start: start,
		End:   end,
	}
}
