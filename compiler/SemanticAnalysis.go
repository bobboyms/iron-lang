package compiler

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"iron-lang/compiler/ironlang"
)

type SemanticAnalysis struct {
	Scopes *Scopes
}

func NewSemanticAnalysis(scopes *Scopes) *SemanticAnalysis {
	return &SemanticAnalysis{
		Scopes: scopes,
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
	default:
		panic("Unknown context")
	}
}

func (s *SemanticAnalysis) VisitProgram(ctx *ironlang.ProgramContext) {
	s.Visit(ctx.FuncMain())
}

func (s *SemanticAnalysis) VisitFuncMain(ctx *ironlang.FuncMainContext) {
	s.Visit(ctx.Scope())
}

func (s *SemanticAnalysis) VisitScope(ctx *ironlang.ScopeContext) {
	s.Scopes.CreateNewScope()

	for _, assignment := range ctx.AllAssignment() {
		s.Visit(assignment)
	}

	//println(ctx.GetText())
}

func (s *SemanticAnalysis) VisitAssignment(ctx *ironlang.AssignmentContext) {
	println(ctx.IDENTIFIER().GetText())
	println(ctx.EQ().GetText())
	s.Visit(ctx.MathExpression())
}

func (s *SemanticAnalysis) VisitMathExpression(ctx *ironlang.MathExpressionContext) {
	println("Math")
	println(ctx.GetText())
}

func (s *SemanticAnalysis) VisitAtom(ctx *ironlang.AtomContext) {
	println(ctx.GetText())
}

//variable := assignment.GetChild(0).(*antlr.TerminalNodeImpl)
//println("Variable: " + variable.GetText())
//actual := s.Scopes.GetActualScope()
//
//if actual.GetSymbolTable(variable.GetText()) == nil {
//	panic("Variavel não declarada")
//}
//
//child := assignment.GetChild(2)
//atom := child.(*ironlang.MathExpressionContext)
//s.Visit(atom)
//s.VisitAtom(atom)
