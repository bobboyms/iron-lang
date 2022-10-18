package compiler

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"iron-lang/compiler/errors"
	"iron-lang/compiler/ironlang"
)

type SemanticAnalysis struct {
	ScopesManager *ScopesManager
	ErrorListener *errors.CustomErrorListener
}

func NewSemanticAnalysis(scopesManager *ScopesManager, errorListener *errors.CustomErrorListener) *SemanticAnalysis {
	return &SemanticAnalysis{
		ScopesManager: scopesManager,
		ErrorListener: errorListener,
	}
}

func (s *SemanticAnalysis) Visit(tree antlr.ParseTree) {
	switch val := tree.(type) {
	case *ironlang.ProgramContext:
		s.VisitProgram(val)
	case *ironlang.FuncMainContext:
		s.VisitFuncMain(val)
	case *ironlang.ScopeContext:
		s.VisitScope(val)
	case *ironlang.MathExpressionContext:
		s.VisitMathExpression(val)
	case *ironlang.AssignmentContext:
		s.VisitAssignment(val)
	case *ironlang.AtomContext:
		s.VisitAtom(val)
	case *ironlang.VariableContext:
		s.VisitVariable(val)
	default:
		panic("Unknown context")
	}
}

func (s *SemanticAnalysis) insertNewError(identifier antlr.TerminalNode, codeError errors.CodeError) {
	s.ErrorListener.SyntaxError(
		nil,
		identifier,
		identifier.GetSymbol().GetLine(),
		identifier.GetSymbol().GetColumn(),
		errors.GetMessageError(codeError, identifier.GetText()), nil)
}

func (s *SemanticAnalysis) VisitProgram(ctx *ironlang.ProgramContext) {
	s.Visit(ctx.FuncMain())
}

func (s *SemanticAnalysis) VisitFuncMain(ctx *ironlang.FuncMainContext) {
	s.Visit(ctx.Scope())
}

func (s *SemanticAnalysis) VisitScope(ctx *ironlang.ScopeContext) {

	s.ScopesManager.CreateNewScope()

	for _, variable := range ctx.AllVariable() {
		s.Visit(variable)
	}

	for _, assignment := range ctx.AllAssignment() {
		s.Visit(assignment)
	}

	for _, scope := range ctx.AllScope() {
		s.Visit(scope)
	}

	s.ScopesManager.DeleteActualScope()
}

func (s *SemanticAnalysis) VisitVariable(ctx *ironlang.VariableContext) {

	identifier := ctx.IDENTIFIER()

	if ctx.DataTypes() == nil {
		switch ctx.GetParent().(type) {
		case *ironlang.AssignmentContext:
		default:
			s.insertNewError(identifier, errors.VariableNotDefined)
		}
	} else if s.ScopesManager.GetActualScope().GetSymbolTable(identifier.GetText()) == nil {
		s.ScopesManager.GetActualScope().Insert(
			identifier.GetText(),
			ironlang.IronLangParserTYPE_INT)
	} else {
		s.insertNewError(identifier, errors.VariableHasDeclared)
	}
}

func (s *SemanticAnalysis) VisitDataTypes(ctx *ironlang.DataTypesContext) {
	ctx.TYPE_INT()
}

func (s *SemanticAnalysis) VisitAssignment(ctx *ironlang.AssignmentContext) {

	if ctx.IDENTIFIER() != nil {
		identifier := ctx.IDENTIFIER()
		if s.ScopesManager.NoHasSymbolTableHigherScopes(identifier.GetText()) {
			s.insertNewError(identifier, errors.VariableNotDeclared)
		}
	} else {
		s.Visit(ctx.Variable())
	}

	s.Visit(ctx.MathExpression())

}

func (s *SemanticAnalysis) VisitMathExpression(ctx *ironlang.MathExpressionContext) {
	println("Math")
	if ctx.Atom() != nil {
		s.Visit(ctx.Atom())
	}
}

func (s *SemanticAnalysis) VisitAtom(ctx *ironlang.AtomContext) {
	println(ctx.GetText())
}
