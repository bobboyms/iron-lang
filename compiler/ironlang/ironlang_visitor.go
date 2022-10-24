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

	// Visit a parse tree produced by IronLangParser#return.
	VisitReturn(ctx *ReturnContext) interface{}

	// Visit a parse tree produced by IronLangParser#scope.
	VisitScope(ctx *ScopeContext) interface{}

	// Visit a parse tree produced by IronLangParser#funcCall.
	VisitFuncCall(ctx *FuncCallContext) interface{}

	// Visit a parse tree produced by IronLangParser#funcCallArg.
	VisitFuncCallArg(ctx *FuncCallArgContext) interface{}

	// Visit a parse tree produced by IronLangParser#anonimousFunc.
	VisitAnonimousFunc(ctx *AnonimousFuncContext) interface{}

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

	// Visit a parse tree produced by IronLangParser#mathExpression.
	VisitMathExpression(ctx *MathExpressionContext) interface{}

	// Visit a parse tree produced by IronLangParser#atom.
	VisitAtom(ctx *AtomContext) interface{}
}
