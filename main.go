package main

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"iron-lang/compiler"
	"iron-lang/compiler/errors"
	"iron-lang/compiler/ironlang"
	"iron-lang/compiler/scopes"
	"log"
	"os"
	"strings"
)

func main() {
	//Lexical analysis
	is := antlr.NewInputStream("fn main() {let variavel int variavel = 20}")
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
	statics := compiler.NewSemanticAnalysis(scopes.NewScopesManager(), customSemanticErrorListener)
	statics.Visit(tree)
	errors.HasSemanticError(customSemanticErrorListener.Errors)

	//Code generator
	generator := compiler.NewLLVMCodeGenerator()
	generator.Visit(tree)
	println(generator.GetBuilder().String())
	NewFile(generator.GetBuilder())
}

func NewFile(strBuilder *strings.Builder) {
	f, err := os.Create("source.ll")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	_, err = f.WriteString(strBuilder.String())

	if err != nil {
		log.Fatal(err)
	}

}
