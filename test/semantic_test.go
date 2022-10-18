package test

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"iron-lang/compiler"
	"iron-lang/compiler/errors"
	"iron-lang/compiler/ironlang"
	"testing"
)

func CreateTest(sourceCode string) []*errors.CustomSyntaxError {
	is := antlr.NewInputStream(sourceCode)
	lexer := ironlang.NewIronLangLexer(is)
	customLexerErrorListener := &errors.CustomErrorListener{}
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(customLexerErrorListener)
	errors.HasLexerError(customLexerErrorListener.Errors)

	//Syntactic analysis
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := ironlang.NewIronLangParser(stream)
	customParserErrorListener := &errors.CustomErrorListener{}
	parser.RemoveErrorListeners()
	parser.AddErrorListener(customParserErrorListener)
	parser.BuildParseTrees = true
	tree := parser.Program()
	errors.HasParserError(customParserErrorListener.Errors)

	//Semantic analysis
	customSemanticErrorListener := &errors.CustomErrorListener{}
	statics := compiler.NewSemanticAnalysis(compiler.NewScopesManager(), customSemanticErrorListener)
	statics.Visit(tree)
	return customSemanticErrorListener.Errors
}

func TestSemanticVariableHasNotBeenDefined(t *testing.T) {
	var errs = CreateTest("fn main() {let x}")
	if len(errs) == 0 {
		t.Error("Variavel precisa de um tipo na sua declaração")
	}

	errs = CreateTest("fn main() {mut let x int}")
	if len(errs) != 0 {
		t.Error("Está correto semanticamente, não devia falhar")
	}

}

func TestSemanticVariableHasBeenDefined(t *testing.T) {
	var errs = CreateTest("fn main() {let x int let x float}")
	if len(errs) == 0 {
		t.Error("Variavel não pode ser declarada 2x dentro do mesmo escopo")
	}

	errs = CreateTest("fn main() {let x int mut let y float}")
	if len(errs) != 0 {
		t.Error("Tem que ser permitido declaração de veriaveis com nomes diferentes")
	}
}

func TestSemanticVariableNotDeclared(t *testing.T) {
	var errs = CreateTest("fn main() {x = 2 + 3}")
	if len(errs) == 0 {
		t.Error("A variavel precisa ser declarada para receber uma expressão")
	}

	errs = CreateTest("fn main() {let x int = 2+3*5}")
	if len(errs) != 0 {
		t.Error("Tem que ser permitido declaração de veriaveis com nomes diferentes")
	}

}

func TestSemanticVariableTypeInference(t *testing.T) {
	var errs = CreateTest("fn main() {let x = 2 + 3}")
	if len(errs) != 0 {
		t.Error("É preciso permitir a inferencia da tipos das variaveis")
	}

	errs = CreateTest("fn main() {let x = 2 + 3 mut let y = x}")
	if len(errs) != 0 {
		t.Error("Essa atribuição tem que ser permitida")
	}
}

func TestSemanticVariableManyScope(t *testing.T) {
	var errs = CreateTest("fn main() {{{{{y = 2}}}}}")
	if len(errs) == 0 {
		t.Error("A variavel tem que ser declarada antes do uso")
	}

	errs = CreateTest("fn main() {let n int {x = 2} }")
	if len(errs) == 0 {
		t.Error("A variavel tem que ser declarada antes do uso dento de scopos internos")
	}

	errs = CreateTest("fn main() {let p int {p = 2} }")
	if len(errs) != 0 {
		t.Error("Se a variavel não foi declarada no escopo local, tem que permitir a declaração no scopo superior")
	}
	errors.HasSemanticError(errs)

	errs = CreateTest("fn main() {let p int {let p = 2} }")
	if len(errs) != 0 {
		t.Error("Permitir a inferencia de tipos dentro de um scopo interno")
	}

	errs = CreateTest("fn main() {let j int {let p = 2} }")
	if len(errs) != 0 {
		t.Error("Permitir a inferencia de tipos dentro de um scopo interno")
	}

	errs = CreateTest("fn main() {let xt int {let xt = 2 {let nj = 2 {{{{let xt = 2}} nj = 2 }let nj float} xt = 15}} }")
	if len(errs) != 0 {
		t.Error("Permitir a inferencia de tipos dentro de muitos scopos internos")
	}

}
