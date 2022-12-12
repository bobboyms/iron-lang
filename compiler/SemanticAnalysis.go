package compiler

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"iron-lang/compiler/errors"
	"iron-lang/compiler/ironlang"
	"iron-lang/compiler/scopes"
)

type SemanticAnalysis struct {
	ScopesManager *scopes.ScopeManager
	ErrorListener *errors.CustomErrorListener
}

func NewSemanticAnalysis(scopesManager *scopes.ScopeManager, errorListener *errors.CustomErrorListener) *SemanticAnalysis {
	return &SemanticAnalysis{
		ScopesManager: scopesManager,
		ErrorListener: errorListener,
	}
}

func (s *SemanticAnalysis) Visit(tree antlr.ParseTree) {
	//switch val := tree.(type) {
	//case *ironlang.ProgramContext:
	//	s.VisitProgram(val)
	//case *ironlang.FuncMainContext:
	//	s.VisitFuncMain(val)
	//case *ironlang.MathExpressionContext:
	//	s.VisitMathExpression(val)
	//case *ironlang.AssignmentContext:
	//	s.VisitAssignment(val)
	//case *ironlang.AtomContext:
	//	s.VisitAtom(val)
	//case *ironlang.VariableContext:
	//	s.VisitVariable(val)
	//case *ironlang.DataTypesContext:
	//	s.VisitDataTypes(val)
	//case *ironlang.FunctionContext:
	//	s.VisitFunction(val)
	//case *ironlang.ReturnContext:
	//	s.VisitReturn(val)
	//case *ironlang.FuncCallContext:
	//	s.VisitFuncCall(val)
	//case *ironlang.FuncCallArgContext:
	//	s.VisitFuncCallArg(val)
	//case *ironlang.BodyScopeContext:
	//	s.VisitBodyScope(val)
	//default:
	//	panic("Unknown context")
	//}
}

//func (s *SemanticAnalysis) VisitFunction(ctx *ironlang.FunctionContext) {
//
//}

func (s *SemanticAnalysis) VisitReturn(ctx *ironlang.ReturnDefinitionContext) {
}

func (s *SemanticAnalysis) VisitFuncCall(ctx *ironlang.FuncCallContext) {

}

func (s *SemanticAnalysis) VisitFuncCallArg(ctx *ironlang.FuncCallArgContext) {

}

func (s *SemanticAnalysis) insertNewError(identifier antlr.TerminalNode, codeError errors.CodeError) {
	s.ErrorListener.SyntaxError(
		nil,
		identifier,
		identifier.GetSymbol().GetLine(),
		identifier.GetSymbol().GetColumn(),
		errors.GetMessageError(codeError, identifier.GetText()), nil)
}

//func (s *SemanticAnalysis) VisitProgram(ctx *ironlang.ProgramContext) {
//	s.Visit(ctx.FuncMain())
//}
//
//func (s *SemanticAnalysis) VisitFuncMain(ctx *ironlang.FuncMainContext) {
//}

func (s *SemanticAnalysis) VisitBodyScope(ctx *ironlang.BodyScopeContext) {
	//if ctx.Variable() != nil {
	//	s.Visit(ctx.Variable())
	//}
	//
	//if ctx.Assignment() != nil {
	//	s.Visit(ctx.Assignment())
	//}

	//if ctx.Function() != nil {
	//	s.Visit(ctx.Function())
	//}
	//
	//if ctx.FuncCall() != nil {
	//	s.Visit(ctx.FuncCall())
	//}
}

//func (s *SemanticAnalysis) VisitScope(ctx *ironlang.FuncScopeContext) {
//
//	s.ScopesManager.CreateNewScope(utils.GetMD5Hash(ctx.GetText()))
//
//
//	//if ctx.Return() != nil {
//	//	s.Visit(ctx.Return())
//	//}
//
//	s.ScopesManager.DeleteActualScope()
//}

func (s *SemanticAnalysis) VisitVariable(ctx *ironlang.VariableContext) {

	//identifier := ctx.IDENTIFIER()
	//
	//if ctx.DataTypes() == nil {
	//	switch ctx.GetParent().(type) {
	//	case *ironlang.AssignmentContext:
	//	default:
	//		s.insertNewError(identifier, errors.VariableNotDefined)
	//	}
	//} else if s.ScopesManager.GetVariable(identifier.GetText()) == nil {
	//	dataType := ctx.DataTypes().(*ironlang.DataTypesContext)
	//
	//	if dataType.TYPE_INT() != nil {
	//		s.ScopesManager.RegisterVariable(
	//			identifier.GetText(),
	//			ironlang.IronLangParserTYPE_INT)
	//	} else {
	//		s.ScopesManager.RegisterVariable(
	//			identifier.GetText(),
	//			ironlang.IronLangParserTYPE_FLOAT)
	//	}
	//
	//} else {
	//	s.insertNewError(identifier, errors.VariableHasDeclared)
	//}
}

func (s *SemanticAnalysis) VisitDataTypes(ctx *ironlang.DataTypesContext) {
	ctx.TYPE_INT()
}

func (s *SemanticAnalysis) VisitAssignment(ctx *ironlang.AssignmentContext) {

}

func (s *SemanticAnalysis) VisitMathExpression(ctx *ironlang.MathExpressionContext) {
	//if ctx.Atom() != nil {
	//	s.Visit(ctx.Atom())
	//}
}

func (s *SemanticAnalysis) VisitAtom(ctx *ironlang.AtomContext) {
	println(ctx.GetText())
}
