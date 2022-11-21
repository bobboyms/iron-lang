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

func (v *BaseIronLangVisitor) VisitFuncType(ctx *FuncTypeContext) interface{} {
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

func (v *BaseIronLangVisitor) VisitAnonimousFunc(ctx *AnonimousFuncContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIronLangVisitor) VisitIfScope(ctx *IfScopeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIronLangVisitor) VisitElseExpression(ctx *ElseExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIronLangVisitor) VisitElseIfExpression(ctx *ElseIfExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIronLangVisitor) VisitIfExpression(ctx *IfExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIronLangVisitor) VisitLoopWhile(ctx *LoopWhileContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIronLangVisitor) VisitLoopForIn(ctx *LoopForInContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIronLangVisitor) VisitBodyScope(ctx *BodyScopeContext) interface{} {
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

func (v *BaseIronLangVisitor) VisitArray(ctx *ArrayContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIronLangVisitor) VisitForEach(ctx *ForEachContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIronLangVisitor) VisitMapFilterReduce(ctx *MapFilterReduceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIronLangVisitor) VisitMap(ctx *MapContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIronLangVisitor) VisitFilter(ctx *FilterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIronLangVisitor) VisitReduce(ctx *ReduceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIronLangVisitor) VisitRelExpression(ctx *RelExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIronLangVisitor) VisitMathExpression(ctx *MathExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIronLangVisitor) VisitAtom(ctx *AtomContext) interface{} {
	return v.VisitChildren(ctx)
}
