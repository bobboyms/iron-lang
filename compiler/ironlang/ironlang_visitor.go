// Code generated from /Users/thiagoluizrodrigues/go/src/iron-lang/IronLang.g4 by ANTLR 4.10.1. DO NOT EDIT.

package ironlang // IronLang
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by IronLangParser.
type IronLangVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by IronLangParser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by IronLangParser#funcMain.
	VisitFuncMain(ctx *FuncMainContext) interface{}

	// Visit a parse tree produced by IronLangParser#function.
	VisitFunction(ctx *FunctionContext) interface{}

	// Visit a parse tree produced by IronLangParser#funcType.
	VisitFuncType(ctx *FuncTypeContext) interface{}

	// Visit a parse tree produced by IronLangParser#return.
	VisitReturn(ctx *ReturnContext) interface{}

	// Visit a parse tree produced by IronLangParser#funcScope.
	VisitFuncScope(ctx *FuncScopeContext) interface{}

	// Visit a parse tree produced by IronLangParser#funcCall.
	VisitFuncCall(ctx *FuncCallContext) interface{}

	// Visit a parse tree produced by IronLangParser#funcCallArg.
	VisitFuncCallArg(ctx *FuncCallArgContext) interface{}

	// Visit a parse tree produced by IronLangParser#anonimousFuncWithReturn.
	VisitAnonimousFuncWithReturn(ctx *AnonimousFuncWithReturnContext) interface{}

	// Visit a parse tree produced by IronLangParser#anonimousFuncNoReturn.
	VisitAnonimousFuncNoReturn(ctx *AnonimousFuncNoReturnContext) interface{}

	// Visit a parse tree produced by IronLangParser#ifScope.
	VisitIfScope(ctx *IfScopeContext) interface{}

	// Visit a parse tree produced by IronLangParser#elseExpression.
	VisitElseExpression(ctx *ElseExpressionContext) interface{}

	// Visit a parse tree produced by IronLangParser#elseIfExpression.
	VisitElseIfExpression(ctx *ElseIfExpressionContext) interface{}

	// Visit a parse tree produced by IronLangParser#ifExpression.
	VisitIfExpression(ctx *IfExpressionContext) interface{}

	// Visit a parse tree produced by IronLangParser#loopScope.
	VisitLoopScope(ctx *LoopScopeContext) interface{}

	// Visit a parse tree produced by IronLangParser#loopDoWhile.
	VisitLoopDoWhile(ctx *LoopDoWhileContext) interface{}

	// Visit a parse tree produced by IronLangParser#loopWhile.
	VisitLoopWhile(ctx *LoopWhileContext) interface{}

	// Visit a parse tree produced by IronLangParser#loopForIn.
	VisitLoopForIn(ctx *LoopForInContext) interface{}

	// Visit a parse tree produced by IronLangParser#loopForI.
	VisitLoopForI(ctx *LoopForIContext) interface{}

	// Visit a parse tree produced by IronLangParser#bodyScope.
	VisitBodyScope(ctx *BodyScopeContext) interface{}

	// Visit a parse tree produced by IronLangParser#println.
	VisitPrintln(ctx *PrintlnContext) interface{}

	// Visit a parse tree produced by IronLangParser#variable.
	VisitVariable(ctx *VariableContext) interface{}

	// Visit a parse tree produced by IronLangParser#functionArgs.
	VisitFunctionArgs(ctx *FunctionArgsContext) interface{}

	// Visit a parse tree produced by IronLangParser#funcArg.
	VisitFuncArg(ctx *FuncArgContext) interface{}

	// Visit a parse tree produced by IronLangParser#dataTypes.
	VisitDataTypes(ctx *DataTypesContext) interface{}

	// Visit a parse tree produced by IronLangParser#assignment.
	VisitAssignment(ctx *AssignmentContext) interface{}

	// Visit a parse tree produced by IronLangParser#array.
	VisitArray(ctx *ArrayContext) interface{}

	// Visit a parse tree produced by IronLangParser#slice.
	VisitSlice(ctx *SliceContext) interface{}

	// Visit a parse tree produced by IronLangParser#forEach.
	VisitForEach(ctx *ForEachContext) interface{}

	// Visit a parse tree produced by IronLangParser#mapFilterReduce.
	VisitMapFilterReduce(ctx *MapFilterReduceContext) interface{}

	// Visit a parse tree produced by IronLangParser#map.
	VisitMap(ctx *MapContext) interface{}

	// Visit a parse tree produced by IronLangParser#filter.
	VisitFilter(ctx *FilterContext) interface{}

	// Visit a parse tree produced by IronLangParser#reduce.
	VisitReduce(ctx *ReduceContext) interface{}

	// Visit a parse tree produced by IronLangParser#relExpression.
	VisitRelExpression(ctx *RelExpressionContext) interface{}

	// Visit a parse tree produced by IronLangParser#mathExpression.
	VisitMathExpression(ctx *MathExpressionContext) interface{}

	// Visit a parse tree produced by IronLangParser#atom.
	VisitAtom(ctx *AtomContext) interface{}
}
