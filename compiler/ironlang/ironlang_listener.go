// Code generated from /Users/thiagoluizrodrigues/go/src/iron-lang/IronLang.g4 by ANTLR 4.10.1. DO NOT EDIT.

package ironlang // IronLang
import "github.com/antlr/antlr4/runtime/Go/antlr"

// IronLangListener is a complete listener for a parse tree produced by IronLangParser.
type IronLangListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterFuncMain is called when entering the funcMain production.
	EnterFuncMain(c *FuncMainContext)

	// EnterFunction is called when entering the function production.
	EnterFunction(c *FunctionContext)

	// EnterFuncType is called when entering the funcType production.
	EnterFuncType(c *FuncTypeContext)

	// EnterReturn is called when entering the return production.
	EnterReturn(c *ReturnContext)

	// EnterFuncScope is called when entering the funcScope production.
	EnterFuncScope(c *FuncScopeContext)

	// EnterFuncCall is called when entering the funcCall production.
	EnterFuncCall(c *FuncCallContext)

	// EnterFuncCallArg is called when entering the funcCallArg production.
	EnterFuncCallArg(c *FuncCallArgContext)

	// EnterAnonimousFuncWithReturn is called when entering the anonimousFuncWithReturn production.
	EnterAnonimousFuncWithReturn(c *AnonimousFuncWithReturnContext)

	// EnterAnonimousFuncNoReturn is called when entering the anonimousFuncNoReturn production.
	EnterAnonimousFuncNoReturn(c *AnonimousFuncNoReturnContext)

	// EnterIfScope is called when entering the ifScope production.
	EnterIfScope(c *IfScopeContext)

	// EnterElseExpression is called when entering the elseExpression production.
	EnterElseExpression(c *ElseExpressionContext)

	// EnterElseIfExpression is called when entering the elseIfExpression production.
	EnterElseIfExpression(c *ElseIfExpressionContext)

	// EnterIfExpression is called when entering the ifExpression production.
	EnterIfExpression(c *IfExpressionContext)

	// EnterLoopScope is called when entering the loopScope production.
	EnterLoopScope(c *LoopScopeContext)

	// EnterLoopDoWhile is called when entering the loopDoWhile production.
	EnterLoopDoWhile(c *LoopDoWhileContext)

	// EnterLoopWhile is called when entering the loopWhile production.
	EnterLoopWhile(c *LoopWhileContext)

	// EnterLoopForIn is called when entering the loopForIn production.
	EnterLoopForIn(c *LoopForInContext)

	// EnterLoopForI is called when entering the loopForI production.
	EnterLoopForI(c *LoopForIContext)

	// EnterBodyScope is called when entering the bodyScope production.
	EnterBodyScope(c *BodyScopeContext)

	// EnterPrintln is called when entering the println production.
	EnterPrintln(c *PrintlnContext)

	// EnterVariable is called when entering the variable production.
	EnterVariable(c *VariableContext)

	// EnterFunctionArgs is called when entering the functionArgs production.
	EnterFunctionArgs(c *FunctionArgsContext)

	// EnterFuncArg is called when entering the funcArg production.
	EnterFuncArg(c *FuncArgContext)

	// EnterDataTypes is called when entering the dataTypes production.
	EnterDataTypes(c *DataTypesContext)

	// EnterAssignment is called when entering the assignment production.
	EnterAssignment(c *AssignmentContext)

	// EnterArray is called when entering the array production.
	EnterArray(c *ArrayContext)

	// EnterSlice is called when entering the slice production.
	EnterSlice(c *SliceContext)

	// EnterForEach is called when entering the forEach production.
	EnterForEach(c *ForEachContext)

	// EnterMapFilterReduce is called when entering the mapFilterReduce production.
	EnterMapFilterReduce(c *MapFilterReduceContext)

	// EnterMap is called when entering the map production.
	EnterMap(c *MapContext)

	// EnterFilter is called when entering the filter production.
	EnterFilter(c *FilterContext)

	// EnterReduce is called when entering the reduce production.
	EnterReduce(c *ReduceContext)

	// EnterRelExpression is called when entering the relExpression production.
	EnterRelExpression(c *RelExpressionContext)

	// EnterMathExpression is called when entering the mathExpression production.
	EnterMathExpression(c *MathExpressionContext)

	// EnterAtom is called when entering the atom production.
	EnterAtom(c *AtomContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitFuncMain is called when exiting the funcMain production.
	ExitFuncMain(c *FuncMainContext)

	// ExitFunction is called when exiting the function production.
	ExitFunction(c *FunctionContext)

	// ExitFuncType is called when exiting the funcType production.
	ExitFuncType(c *FuncTypeContext)

	// ExitReturn is called when exiting the return production.
	ExitReturn(c *ReturnContext)

	// ExitFuncScope is called when exiting the funcScope production.
	ExitFuncScope(c *FuncScopeContext)

	// ExitFuncCall is called when exiting the funcCall production.
	ExitFuncCall(c *FuncCallContext)

	// ExitFuncCallArg is called when exiting the funcCallArg production.
	ExitFuncCallArg(c *FuncCallArgContext)

	// ExitAnonimousFuncWithReturn is called when exiting the anonimousFuncWithReturn production.
	ExitAnonimousFuncWithReturn(c *AnonimousFuncWithReturnContext)

	// ExitAnonimousFuncNoReturn is called when exiting the anonimousFuncNoReturn production.
	ExitAnonimousFuncNoReturn(c *AnonimousFuncNoReturnContext)

	// ExitIfScope is called when exiting the ifScope production.
	ExitIfScope(c *IfScopeContext)

	// ExitElseExpression is called when exiting the elseExpression production.
	ExitElseExpression(c *ElseExpressionContext)

	// ExitElseIfExpression is called when exiting the elseIfExpression production.
	ExitElseIfExpression(c *ElseIfExpressionContext)

	// ExitIfExpression is called when exiting the ifExpression production.
	ExitIfExpression(c *IfExpressionContext)

	// ExitLoopScope is called when exiting the loopScope production.
	ExitLoopScope(c *LoopScopeContext)

	// ExitLoopDoWhile is called when exiting the loopDoWhile production.
	ExitLoopDoWhile(c *LoopDoWhileContext)

	// ExitLoopWhile is called when exiting the loopWhile production.
	ExitLoopWhile(c *LoopWhileContext)

	// ExitLoopForIn is called when exiting the loopForIn production.
	ExitLoopForIn(c *LoopForInContext)

	// ExitLoopForI is called when exiting the loopForI production.
	ExitLoopForI(c *LoopForIContext)

	// ExitBodyScope is called when exiting the bodyScope production.
	ExitBodyScope(c *BodyScopeContext)

	// ExitPrintln is called when exiting the println production.
	ExitPrintln(c *PrintlnContext)

	// ExitVariable is called when exiting the variable production.
	ExitVariable(c *VariableContext)

	// ExitFunctionArgs is called when exiting the functionArgs production.
	ExitFunctionArgs(c *FunctionArgsContext)

	// ExitFuncArg is called when exiting the funcArg production.
	ExitFuncArg(c *FuncArgContext)

	// ExitDataTypes is called when exiting the dataTypes production.
	ExitDataTypes(c *DataTypesContext)

	// ExitAssignment is called when exiting the assignment production.
	ExitAssignment(c *AssignmentContext)

	// ExitArray is called when exiting the array production.
	ExitArray(c *ArrayContext)

	// ExitSlice is called when exiting the slice production.
	ExitSlice(c *SliceContext)

	// ExitForEach is called when exiting the forEach production.
	ExitForEach(c *ForEachContext)

	// ExitMapFilterReduce is called when exiting the mapFilterReduce production.
	ExitMapFilterReduce(c *MapFilterReduceContext)

	// ExitMap is called when exiting the map production.
	ExitMap(c *MapContext)

	// ExitFilter is called when exiting the filter production.
	ExitFilter(c *FilterContext)

	// ExitReduce is called when exiting the reduce production.
	ExitReduce(c *ReduceContext)

	// ExitRelExpression is called when exiting the relExpression production.
	ExitRelExpression(c *RelExpressionContext)

	// ExitMathExpression is called when exiting the mathExpression production.
	ExitMathExpression(c *MathExpressionContext)

	// ExitAtom is called when exiting the atom production.
	ExitAtom(c *AtomContext)
}