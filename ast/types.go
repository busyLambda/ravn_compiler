package ast

import (
	"fmt"

	"github.com/llir/llvm/ir/types"
)

type Type interface {
	IsBuiltin() bool
	AsLLVMType() any
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

// I don't think there is a set type i can use here so it's gonna be any for now :c
func (n *NumericType) AsLLVMType() (G any) {
	// Convert the n to a types.Type
	switch n.Kind {
	case Float:
		return types.Float
	default:
		switch n.Bits {
		case B8:
			return types.I8
		case B16:
			return types.I16
		case B32:
			return types.I32
		case B64:
			return types.I64
		default:
			return types.I128
		}
	}
}

func (n *NumericType) String() string {
	return fmt.Sprintf("%s%v", n.Kind, n.Bits)
}

func (n *NumericType) IsBuiltin() bool {
	return true
}

type Str struct {
	Len uint64
}

func (s *Str) IsBuiltin() bool {
	return true
}

func (s *Str) AsLLVMType() (G any) {
	return types.NewArray(s.Len, types.I8)
}

type Bool struct{}

func (b *Bool) IsBuiltin() bool {
	return true
}
