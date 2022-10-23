// Code generated from /Users/thiagoluizrodrigues/go/src/iron-lang/IronLang.g4 by ANTLR 4.10.1. DO NOT EDIT.

package ironlang // IronLang
import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseIronLangVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseIronLangVisitor) VisitProgram(ctx *ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIronLangVisitor) VisitFuncMain(ctx *FuncMainContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIronLangVisitor) VisitFunction(ctx *FunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIronLangVisitor) VisitReturn(ctx *ReturnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIronLangVisitor) VisitScope(ctx *ScopeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIronLangVisitor) VisitFuncCall(ctx *FuncCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIronLangVisitor) VisitFuncCallArg(ctx *FuncCallArgContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIronLangVisitor) VisitPrintln(ctx *PrintlnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIronLangVisitor) VisitVariable(ctx *VariableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIronLangVisitor) VisitFunctionArgs(ctx *FunctionArgsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIronLangVisitor) VisitFuncArg(ctx *FuncArgContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIronLangVisitor) VisitDataTypes(ctx *DataTypesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIronLangVisitor) VisitAssignment(ctx *AssignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIronLangVisitor) VisitMathExpression(ctx *MathExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIronLangVisitor) VisitAtom(ctx *AtomContext) interface{} {
	return v.VisitChildren(ctx)
}
