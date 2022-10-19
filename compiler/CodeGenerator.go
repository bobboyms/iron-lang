package compiler

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"iron-lang/compiler/ironlang"
	"strings"
)

type LLVMCodeGenerator struct {
	StrBuilder *strings.Builder
}

func NewLLVMCodeGenerator() *LLVMCodeGenerator {
	return &LLVMCodeGenerator{
		StrBuilder: &strings.Builder{},
	}
}

func (l *LLVMCodeGenerator) Visit(tree antlr.ParseTree) {
	switch val := tree.(type) {
	case *ironlang.ProgramContext:
		l.VisitProgram(val)
	case *ironlang.FuncMainContext:
		l.VisitFuncMain(val)
	case *ironlang.ScopeContext:
		l.VisitScope(val)
	case *ironlang.MathExpressionContext:
		l.VisitMathExpression(val)
	case *ironlang.AssignmentContext:
		l.VisitAssignment(val)
	case *ironlang.AtomContext:
		l.VisitAtom(val)
	case *ironlang.VariableContext:
		l.VisitVariable(val)
	case *ironlang.DataTypesContext:
		l.VisitDataTypes(val)
	default:
		panic("Unknown context")
	}
}

func (l *LLVMCodeGenerator) GetBuilder() *strings.Builder {
	return l.StrBuilder
}

func (l *LLVMCodeGenerator) VisitProgram(ctx *ironlang.ProgramContext) {
	l.Visit(ctx.FuncMain())
	l.StrBuilder.WriteString("")
}

func (l *LLVMCodeGenerator) VisitFuncMain(ctx *ironlang.FuncMainContext) {
	l.StrBuilder.WriteString("define i32 @main() #0 ")
	l.StrBuilder.WriteString("{\n")
	l.Visit(ctx.Scope())
	l.StrBuilder.WriteString("ret i32 0\n")
	l.StrBuilder.WriteString("}\n")
}

func (l *LLVMCodeGenerator) VisitScope(ctx *ironlang.ScopeContext) {
	for _, variable := range ctx.AllVariable() {
		l.Visit(variable)
	}

	for _, assignment := range ctx.AllAssignment() {
		l.Visit(assignment)
	}
}

func (l *LLVMCodeGenerator) VisitVariable(ctx *ironlang.VariableContext) {
	//%1 = alloca i32, align 4
	l.StrBuilder.WriteString("%" + ctx.IDENTIFIER().GetText() + " = alloca ")
	l.Visit(ctx.DataTypes())
}

func (l *LLVMCodeGenerator) VisitDataTypes(ctx *ironlang.DataTypesContext) {
	if ctx.TYPE_INT() != nil {
		l.StrBuilder.WriteString("i32, align 4\n")
	} else {
		l.StrBuilder.WriteString("float, align 4\n")
	}
}

func (l *LLVMCodeGenerator) VisitAssignment(ctx *ironlang.AssignmentContext) {
	println(ctx.GetText())
	ctx.IDENTIFIER().GetText()
	//store i32 135, i32* %2, align 4
	l.StrBuilder.WriteString("float, align 4\n")
}

func (l *LLVMCodeGenerator) VisitMathExpression(ctx *ironlang.MathExpressionContext) {
}

func (l *LLVMCodeGenerator) VisitAtom(ctx *ironlang.AtomContext) {
}
