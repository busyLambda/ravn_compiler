package parser

import "fmt"

type Type interface {
	IsBuiltin() bool
}

// The set of bits allowed in the language i.e.: i<Bits> -> i32 etc...
type Bits int

const (
	B128 Bits = 128 // 128 bits
	B64       = 64  // 64 bits
	B32       = 32  // 32 bits
	B16       = 16  // 16 bits
	B8        = 8   // 8 bits
)

type NumericTypeKind string

const (
	Signed   NumericTypeKind = "i"
	Unsigned                 = "u"
	Float                    = "f"
)

type NumericType struct {
	Kind NumericTypeKind `json:"kind"`
	Bits Bits            `json:"bits"`
}

func (n *NumericType) String() string {
	return fmt.Sprintf("%s%v", n.Kind, n.Bits)
}

func (n *NumericType) IsBuiltin() bool {
	return true
}

type Str struct{}

func (s *Str) IsBuiltin() bool {
	return true
}

type Bool struct{}

func (b *Bool) IsBuiltin() bool {
	return true
}
