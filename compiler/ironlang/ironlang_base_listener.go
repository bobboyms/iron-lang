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

// EnterScope is called when production scope is entered.
func (s *BaseIronLangListener) EnterScope(ctx *ScopeContext) {}

// ExitScope is called when production scope is exited.
func (s *BaseIronLangListener) ExitScope(ctx *ScopeContext) {}

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
