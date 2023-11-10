package codegen

import (
	"github.com/busylambda/raven/ast"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/types"
)

type LLVMModule struct {
	module *ir.Module
}

func InitLLVMModule() *LLVMModule {
	return &LLVMModule{module: ir.NewModule()}
}

func (lm *LLVMModule) Func(f ast.FuncDecl) *ir.Func {
	return lm.module.NewFunc(f.Name.Name, types.I32)
}
