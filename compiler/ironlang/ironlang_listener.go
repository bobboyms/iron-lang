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

	// EnterScope is called when entering the scope production.
	EnterScope(c *ScopeContext)

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

	// EnterMathExpression is called when entering the mathExpression production.
	EnterMathExpression(c *MathExpressionContext)

	// EnterAtom is called when entering the atom production.
	EnterAtom(c *AtomContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitFuncMain is called when exiting the funcMain production.
	ExitFuncMain(c *FuncMainContext)

	// ExitScope is called when exiting the scope production.
	ExitScope(c *ScopeContext)

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

	// ExitMathExpression is called when exiting the mathExpression production.
	ExitMathExpression(c *MathExpressionContext)

	// ExitAtom is called when exiting the atom production.
	ExitAtom(c *AtomContext)
}
