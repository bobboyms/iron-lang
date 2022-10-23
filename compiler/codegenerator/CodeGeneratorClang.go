package codegenerator

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"iron-lang/compiler/ironlang"
	"iron-lang/compiler/scopes"
	"iron-lang/compiler/utils"
	"strings"
)

type ClangPlus struct {
	StrBuilder       *StringBuilder
	StrImportBuilder *StringBuilder
	StrFuncBuilder   *StringBuilder
	ScopeLog         *scopes.ScopeLog
}

func NewClang(scopeLog *scopes.ScopeLog) *ClangPlus {
	return &ClangPlus{
		StrBuilder:       NewStringBuilder(),
		StrImportBuilder: NewStringBuilder(),
		ScopeLog:         scopeLog,
	}
}

func (l *ClangPlus) Visit(tree antlr.ParseTree) {
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
	case *ironlang.FunctionContext:
		l.VisitFunction(val)
	case *ironlang.ReturnContext:
		l.VisitReturn(val)
	case *ironlang.FuncCallContext:
		l.VisitFuncCall(val)
	case *ironlang.FuncCallArgContext:
		l.VisitFuncCallArg(val)

	default:
		panic("Unknown context")
	}
}

func (l *ClangPlus) GetBuilder() *strings.Builder {
	builder := &strings.Builder{}
	builder.WriteString(l.StrImportBuilder.GetBuilder().String())
	if l.StrFuncBuilder != nil {
		builder.WriteString(l.StrFuncBuilder.GetBuilder().String())
	}
	builder.WriteString(l.StrBuilder.GetBuilder().String())
	return builder
}

func (l *ClangPlus) VisitProgram(ctx *ironlang.ProgramContext) {
	l.StrImportBuilder.WriteString("#include <iostream>\n")
	l.Visit(ctx.FuncMain())
	l.StrBuilder.WriteString("\n")

}

func (l *ClangPlus) VisitFuncMain(ctx *ironlang.FuncMainContext) {
	l.StrBuilder.WriteString("int main() ")
	l.StrBuilder.WriteString("{\n")
	l.Visit(ctx.Scope())
	l.StrBuilder.WriteString("std::cout << \"Hello world!\";\n")
	l.StrBuilder.WriteString("return 0;\n")
	l.StrBuilder.WriteString("}\n")
}

func (l *ClangPlus) VisitScope(ctx *ironlang.ScopeContext) {

	l.ScopeLog.EnterScope(utils.GetMD5Hash(ctx.GetText()))

	for _, variable := range ctx.AllVariable() {
		l.Visit(variable)
	}

	for _, assignment := range ctx.AllAssignment() {
		l.Visit(assignment)
	}

	builder := NewStringBuilder()
	for _, function := range ctx.AllFunction() {
		c := NewClang(l.ScopeLog)
		c.Visit(function)
		builder.WriteString(c.GetBuilder().String())
	}

	if len(builder.GetBuilder().String()) > 0 {
		l.StrFuncBuilder = NewStringBuilder()
		l.StrFuncBuilder.WriteString(builder.GetBuilder().String())
	}

	if ctx.Return() != nil {
		l.Visit(ctx.Return())
	}

	for _, funcCall := range ctx.AllFuncCall() {
		l.Visit(funcCall)
	}

	l.ScopeLog.ExitScope()

}

func (l *ClangPlus) VisitReturn(ctx *ironlang.ReturnContext) {
	l.StrBuilder.WriteString("return ")
	l.Visit(ctx.MathExpression())
	l.StrBuilder.WriteString(";\n")
}

func (l *ClangPlus) VisitFuncCall(ctx *ironlang.FuncCallContext) {
	l.StrBuilder.WriteString(ctx.IDENTIFIER().GetText())
	l.StrBuilder.WriteString("(")
	for i, funcCallArg := range ctx.AllFuncCallArg() {
		l.Visit(funcCallArg)
		if ctx.COMMA(i) != nil {
			l.StrBuilder.WriteString(",")
		}
	}
	l.StrBuilder.WriteString(")")

	v, _ := ctx.GetParent().(*ironlang.FuncCallArgContext)
	if v == nil {
		l.StrBuilder.WriteString(";\n")
	}

}

func (l *ClangPlus) VisitFuncCallArg(ctx *ironlang.FuncCallArgContext) {
	if ctx.MathExpression() != nil {
		l.Visit(ctx.MathExpression())
	}

	if ctx.FuncCall() != nil {
		l.Visit(ctx.FuncCall())
	}

}

func (l *ClangPlus) VisitFunction(ctx *ironlang.FunctionContext) {
	l.StrBuilder.WriteString("\n")
	if ctx.DataTypes() != nil {
		l.Visit(ctx.DataTypes())
	} else {
		l.StrBuilder.WriteString("void ")
	}

	l.StrBuilder.WriteString(ctx.IDENTIFIER().GetText())
	l.StrBuilder.WriteString("(")
	l.StrBuilder.WriteString(") ")
	l.StrBuilder.WriteString("{\n")
	l.Visit(ctx.Scope())
	l.StrBuilder.WriteString("}\n")
}

func (l *ClangPlus) VisitVariable(ctx *ironlang.VariableContext) {

	if ctx.DataTypes() != nil {
		l.Visit(ctx.DataTypes())
	} else {
		l.StrBuilder.WriteString("float ")
	}

	switch ctx.GetParent().(type) {
	case *ironlang.AssignmentContext:
		l.StrBuilder.WriteString(ctx.IDENTIFIER().GetText())
	default:
		l.StrBuilder.WriteString(ctx.IDENTIFIER().GetText() + ";\n")
	}

}

func (l *ClangPlus) VisitDataTypes(ctx *ironlang.DataTypesContext) {
	if ctx.TYPE_INT() != nil {
		l.StrBuilder.WriteString("int ")
	} else {
		l.StrBuilder.WriteString("float ")
	}
}

func (l *ClangPlus) VisitAssignment(ctx *ironlang.AssignmentContext) {

	if ctx.IDENTIFIER() != nil {
		l.StrBuilder.WriteString(ctx.IDENTIFIER().GetText() + " = ")
	}

	if ctx.Variable() != nil {
		l.Visit(ctx.Variable())
		l.StrBuilder.WriteString(" = ")
	}

	l.Visit(ctx.MathExpression())
	l.StrBuilder.WriteString(";\n")

}

func (l *ClangPlus) VisitMathExpression(ctx *ironlang.MathExpressionContext) {

	if ctx.L_PAREN() != nil && ctx.R_PAREN() != nil {
		l.StrBuilder.WriteString("(")
		l.Visit(ctx.AllMathExpression()[0])
		l.StrBuilder.WriteString(")")
	}

	if ctx.Atom() != nil {
		l.Visit(ctx.Atom())
	}

	if len(ctx.AllMathExpression()) == 2 {
		l.Visit(ctx.AllMathExpression()[0])
		if ctx.PLUS() != nil {
			l.StrBuilder.WriteString(" + ")
		}

		if ctx.MINUS() != nil {
			l.StrBuilder.WriteString(" - ")
		}

		if ctx.DIV() != nil {
			l.StrBuilder.WriteString(" / ")
		}

		if ctx.MULT() != nil {
			l.StrBuilder.WriteString(" * ")
		}

		l.Visit(ctx.AllMathExpression()[1])

	}

}

func (l *ClangPlus) VisitAtom(ctx *ironlang.AtomContext) {

	if ctx.INT_NUMBER() != nil {
		l.StrBuilder.WriteString(ctx.INT_NUMBER().GetText())
	} else if ctx.REAL_NUMBER() != nil {
		l.StrBuilder.WriteString(ctx.REAL_NUMBER().GetText())
	} else {
		l.StrBuilder.WriteString(ctx.IDENTIFIER().GetText())
	}

}