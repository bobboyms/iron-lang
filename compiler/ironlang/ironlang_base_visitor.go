// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package ironlang // IronLang
import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

type BaseIronLangVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseIronLangVisitor) VisitProgram(ctx *ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIronLangVisitor) VisitBodyProgram(ctx *BodyProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIronLangVisitor) VisitFunctionReturn(ctx *FunctionReturnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIronLangVisitor) VisitFunctionNoReturn(ctx *FunctionNoReturnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIronLangVisitor) VisitFuncType(ctx *FuncTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIronLangVisitor) VisitImplConstructor(ctx *ImplConstructorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIronLangVisitor) VisitReturnDefinition(ctx *ReturnDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIronLangVisitor) VisitFuncCall(ctx *FuncCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIronLangVisitor) VisitFuncCallArg(ctx *FuncCallArgContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIronLangVisitor) VisitAnonimousFuncWithReturn(ctx *AnonimousFuncWithReturnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIronLangVisitor) VisitAnonimousFuncNoReturn(ctx *AnonimousFuncNoReturnContext) interface{} {
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

func (v *BaseIronLangVisitor) VisitLoopScope(ctx *LoopScopeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIronLangVisitor) VisitLoopDoWhile(ctx *LoopDoWhileContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIronLangVisitor) VisitLoopWhile(ctx *LoopWhileContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIronLangVisitor) VisitLoopForIn(ctx *LoopForInContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIronLangVisitor) VisitLoopForI(ctx *LoopForIContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIronLangVisitor) VisitBodyScope(ctx *BodyScopeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIronLangVisitor) VisitTrait(ctx *TraitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIronLangVisitor) VisitImpl(ctx *ImplContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIronLangVisitor) VisitInitImpl(ctx *InitImplContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIronLangVisitor) VisitStructDefinition(ctx *StructDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIronLangVisitor) VisitDefinitionVariables(ctx *DefinitionVariablesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIronLangVisitor) VisitInitStruct(ctx *InitStructContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIronLangVisitor) VisitInitStructBody(ctx *InitStructBodyContext) interface{} {
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

func (v *BaseIronLangVisitor) VisitDataTypes(ctx *DataTypesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIronLangVisitor) VisitAssignment(ctx *AssignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIronLangVisitor) VisitLeftAssignment(ctx *LeftAssignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIronLangVisitor) VisitRigthAssignment(ctx *RigthAssignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIronLangVisitor) VisitArray(ctx *ArrayContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIronLangVisitor) VisitSlice(ctx *SliceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIronLangVisitor) VisitForEach(ctx *ForEachContext) interface{} {
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

func (v *BaseIronLangVisitor) VisitExpression(ctx *ExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIronLangVisitor) VisitMathExpression(ctx *MathExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIronLangVisitor) VisitAtom(ctx *AtomContext) interface{} {
	return v.VisitChildren(ctx)
}
