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

// EnterScope is called when production scope is entered.
func (s *BaseIronLangListener) EnterScope(ctx *ScopeContext) {}

// ExitScope is called when production scope is exited.
func (s *BaseIronLangListener) ExitScope(ctx *ScopeContext) {}

// EnterFuncCall is called when production funcCall is entered.
func (s *BaseIronLangListener) EnterFuncCall(ctx *FuncCallContext) {}

// ExitFuncCall is called when production funcCall is exited.
func (s *BaseIronLangListener) ExitFuncCall(ctx *FuncCallContext) {}

// EnterFuncCallArg is called when production funcCallArg is entered.
func (s *BaseIronLangListener) EnterFuncCallArg(ctx *FuncCallArgContext) {}

// ExitFuncCallArg is called when production funcCallArg is exited.
func (s *BaseIronLangListener) ExitFuncCallArg(ctx *FuncCallArgContext) {}

// EnterAnonimousFunc is called when production anonimousFunc is entered.
func (s *BaseIronLangListener) EnterAnonimousFunc(ctx *AnonimousFuncContext) {}

// ExitAnonimousFunc is called when production anonimousFunc is exited.
func (s *BaseIronLangListener) ExitAnonimousFunc(ctx *AnonimousFuncContext) {}

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

// EnterMathExpression is called when production mathExpression is entered.
func (s *BaseIronLangListener) EnterMathExpression(ctx *MathExpressionContext) {}

// ExitMathExpression is called when production mathExpression is exited.
func (s *BaseIronLangListener) ExitMathExpression(ctx *MathExpressionContext) {}

// EnterAtom is called when production atom is entered.
func (s *BaseIronLangListener) EnterAtom(ctx *AtomContext) {}

// ExitAtom is called when production atom is exited.
func (s *BaseIronLangListener) ExitAtom(ctx *AtomContext) {}
