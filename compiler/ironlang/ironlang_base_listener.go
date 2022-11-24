// Code generated from /Users/thiagoluizrodrigues/go/src/iron-lang/IronLang.g4 by ANTLR 4.10.1. DO NOT EDIT.

package ironlang // IronLang
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseIronLangListener is a complete listener for a parse tree produced by IronLangParser.
type BaseIronLangListener struct{}

var _ IronLangListener = &BaseIronLangListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseIronLangListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseIronLangListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseIronLangListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseIronLangListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseIronLangListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseIronLangListener) ExitProgram(ctx *ProgramContext) {}

// EnterFuncMain is called when production funcMain is entered.
func (s *BaseIronLangListener) EnterFuncMain(ctx *FuncMainContext) {}

// ExitFuncMain is called when production funcMain is exited.
func (s *BaseIronLangListener) ExitFuncMain(ctx *FuncMainContext) {}

// EnterFunction is called when production function is entered.
func (s *BaseIronLangListener) EnterFunction(ctx *FunctionContext) {}

// ExitFunction is called when production function is exited.
func (s *BaseIronLangListener) ExitFunction(ctx *FunctionContext) {}

// EnterFuncType is called when production funcType is entered.
func (s *BaseIronLangListener) EnterFuncType(ctx *FuncTypeContext) {}

// ExitFuncType is called when production funcType is exited.
func (s *BaseIronLangListener) ExitFuncType(ctx *FuncTypeContext) {}

// EnterReturn is called when production return is entered.
func (s *BaseIronLangListener) EnterReturn(ctx *ReturnContext) {}

// ExitReturn is called when production return is exited.
func (s *BaseIronLangListener) ExitReturn(ctx *ReturnContext) {}

// EnterFuncScope is called when production funcScope is entered.
func (s *BaseIronLangListener) EnterFuncScope(ctx *FuncScopeContext) {}

// ExitFuncScope is called when production funcScope is exited.
func (s *BaseIronLangListener) ExitFuncScope(ctx *FuncScopeContext) {}

// EnterFuncCall is called when production funcCall is entered.
func (s *BaseIronLangListener) EnterFuncCall(ctx *FuncCallContext) {}

// ExitFuncCall is called when production funcCall is exited.
func (s *BaseIronLangListener) ExitFuncCall(ctx *FuncCallContext) {}

// EnterFuncCallArg is called when production funcCallArg is entered.
func (s *BaseIronLangListener) EnterFuncCallArg(ctx *FuncCallArgContext) {}

// ExitFuncCallArg is called when production funcCallArg is exited.
func (s *BaseIronLangListener) ExitFuncCallArg(ctx *FuncCallArgContext) {}

// EnterAnonimousFuncWithReturn is called when production anonimousFuncWithReturn is entered.
func (s *BaseIronLangListener) EnterAnonimousFuncWithReturn(ctx *AnonimousFuncWithReturnContext) {}

// ExitAnonimousFuncWithReturn is called when production anonimousFuncWithReturn is exited.
func (s *BaseIronLangListener) ExitAnonimousFuncWithReturn(ctx *AnonimousFuncWithReturnContext) {}

// EnterAnonimousFuncNoReturn is called when production anonimousFuncNoReturn is entered.
func (s *BaseIronLangListener) EnterAnonimousFuncNoReturn(ctx *AnonimousFuncNoReturnContext) {}

// ExitAnonimousFuncNoReturn is called when production anonimousFuncNoReturn is exited.
func (s *BaseIronLangListener) ExitAnonimousFuncNoReturn(ctx *AnonimousFuncNoReturnContext) {}

// EnterIfScope is called when production ifScope is entered.
func (s *BaseIronLangListener) EnterIfScope(ctx *IfScopeContext) {}

// ExitIfScope is called when production ifScope is exited.
func (s *BaseIronLangListener) ExitIfScope(ctx *IfScopeContext) {}

// EnterElseExpression is called when production elseExpression is entered.
func (s *BaseIronLangListener) EnterElseExpression(ctx *ElseExpressionContext) {}

// ExitElseExpression is called when production elseExpression is exited.
func (s *BaseIronLangListener) ExitElseExpression(ctx *ElseExpressionContext) {}

// EnterElseIfExpression is called when production elseIfExpression is entered.
func (s *BaseIronLangListener) EnterElseIfExpression(ctx *ElseIfExpressionContext) {}

// ExitElseIfExpression is called when production elseIfExpression is exited.
func (s *BaseIronLangListener) ExitElseIfExpression(ctx *ElseIfExpressionContext) {}

// EnterIfExpression is called when production ifExpression is entered.
func (s *BaseIronLangListener) EnterIfExpression(ctx *IfExpressionContext) {}

// ExitIfExpression is called when production ifExpression is exited.
func (s *BaseIronLangListener) ExitIfExpression(ctx *IfExpressionContext) {}

// EnterLoopScope is called when production loopScope is entered.
func (s *BaseIronLangListener) EnterLoopScope(ctx *LoopScopeContext) {}

// ExitLoopScope is called when production loopScope is exited.
func (s *BaseIronLangListener) ExitLoopScope(ctx *LoopScopeContext) {}

// EnterLoopDoWhile is called when production loopDoWhile is entered.
func (s *BaseIronLangListener) EnterLoopDoWhile(ctx *LoopDoWhileContext) {}

// ExitLoopDoWhile is called when production loopDoWhile is exited.
func (s *BaseIronLangListener) ExitLoopDoWhile(ctx *LoopDoWhileContext) {}

// EnterLoopWhile is called when production loopWhile is entered.
func (s *BaseIronLangListener) EnterLoopWhile(ctx *LoopWhileContext) {}

// ExitLoopWhile is called when production loopWhile is exited.
func (s *BaseIronLangListener) ExitLoopWhile(ctx *LoopWhileContext) {}

// EnterLoopForIn is called when production loopForIn is entered.
func (s *BaseIronLangListener) EnterLoopForIn(ctx *LoopForInContext) {}

// ExitLoopForIn is called when production loopForIn is exited.
func (s *BaseIronLangListener) ExitLoopForIn(ctx *LoopForInContext) {}

// EnterLoopForI is called when production loopForI is entered.
func (s *BaseIronLangListener) EnterLoopForI(ctx *LoopForIContext) {}

// ExitLoopForI is called when production loopForI is exited.
func (s *BaseIronLangListener) ExitLoopForI(ctx *LoopForIContext) {}

// EnterBodyScope is called when production bodyScope is entered.
func (s *BaseIronLangListener) EnterBodyScope(ctx *BodyScopeContext) {}

// ExitBodyScope is called when production bodyScope is exited.
func (s *BaseIronLangListener) ExitBodyScope(ctx *BodyScopeContext) {}

// EnterPrintln is called when production println is entered.
func (s *BaseIronLangListener) EnterPrintln(ctx *PrintlnContext) {}

// ExitPrintln is called when production println is exited.
func (s *BaseIronLangListener) ExitPrintln(ctx *PrintlnContext) {}

// EnterVariable is called when production variable is entered.
func (s *BaseIronLangListener) EnterVariable(ctx *VariableContext) {}

// ExitVariable is called when production variable is exited.
func (s *BaseIronLangListener) ExitVariable(ctx *VariableContext) {}

// EnterFunctionArgs is called when production functionArgs is entered.
func (s *BaseIronLangListener) EnterFunctionArgs(ctx *FunctionArgsContext) {}

// ExitFunctionArgs is called when production functionArgs is exited.
func (s *BaseIronLangListener) ExitFunctionArgs(ctx *FunctionArgsContext) {}

// EnterFuncArg is called when production funcArg is entered.
func (s *BaseIronLangListener) EnterFuncArg(ctx *FuncArgContext) {}

// ExitFuncArg is called when production funcArg is exited.
func (s *BaseIronLangListener) ExitFuncArg(ctx *FuncArgContext) {}

// EnterDataTypes is called when production dataTypes is entered.
func (s *BaseIronLangListener) EnterDataTypes(ctx *DataTypesContext) {}

// ExitDataTypes is called when production dataTypes is exited.
func (s *BaseIronLangListener) ExitDataTypes(ctx *DataTypesContext) {}

// EnterAssignment is called when production assignment is entered.
func (s *BaseIronLangListener) EnterAssignment(ctx *AssignmentContext) {}

// ExitAssignment is called when production assignment is exited.
func (s *BaseIronLangListener) ExitAssignment(ctx *AssignmentContext) {}

// EnterArray is called when production array is entered.
func (s *BaseIronLangListener) EnterArray(ctx *ArrayContext) {}

// ExitArray is called when production array is exited.
func (s *BaseIronLangListener) ExitArray(ctx *ArrayContext) {}

// EnterSlice is called when production slice is entered.
func (s *BaseIronLangListener) EnterSlice(ctx *SliceContext) {}

// ExitSlice is called when production slice is exited.
func (s *BaseIronLangListener) ExitSlice(ctx *SliceContext) {}

// EnterForEach is called when production forEach is entered.
func (s *BaseIronLangListener) EnterForEach(ctx *ForEachContext) {}

// ExitForEach is called when production forEach is exited.
func (s *BaseIronLangListener) ExitForEach(ctx *ForEachContext) {}

// EnterMapFilterReduce is called when production mapFilterReduce is entered.
func (s *BaseIronLangListener) EnterMapFilterReduce(ctx *MapFilterReduceContext) {}

// ExitMapFilterReduce is called when production mapFilterReduce is exited.
func (s *BaseIronLangListener) ExitMapFilterReduce(ctx *MapFilterReduceContext) {}

// EnterMap is called when production map is entered.
func (s *BaseIronLangListener) EnterMap(ctx *MapContext) {}

// ExitMap is called when production map is exited.
func (s *BaseIronLangListener) ExitMap(ctx *MapContext) {}

// EnterFilter is called when production filter is entered.
func (s *BaseIronLangListener) EnterFilter(ctx *FilterContext) {}

// ExitFilter is called when production filter is exited.
func (s *BaseIronLangListener) ExitFilter(ctx *FilterContext) {}

// EnterReduce is called when production reduce is entered.
func (s *BaseIronLangListener) EnterReduce(ctx *ReduceContext) {}

// ExitReduce is called when production reduce is exited.
func (s *BaseIronLangListener) ExitReduce(ctx *ReduceContext) {}

// EnterRelExpression is called when production relExpression is entered.
func (s *BaseIronLangListener) EnterRelExpression(ctx *RelExpressionContext) {}

// ExitRelExpression is called when production relExpression is exited.
func (s *BaseIronLangListener) ExitRelExpression(ctx *RelExpressionContext) {}

// EnterMathExpression is called when production mathExpression is entered.
func (s *BaseIronLangListener) EnterMathExpression(ctx *MathExpressionContext) {}

// ExitMathExpression is called when production mathExpression is exited.
func (s *BaseIronLangListener) ExitMathExpression(ctx *MathExpressionContext) {}

// EnterAtom is called when production atom is entered.
func (s *BaseIronLangListener) EnterAtom(ctx *AtomContext) {}

// ExitAtom is called when production atom is exited.
func (s *BaseIronLangListener) ExitAtom(ctx *AtomContext) {}