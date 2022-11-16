package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"iron-lang/compiler"
	"iron-lang/compiler/codegenerator"
	"iron-lang/compiler/errors"
	"iron-lang/compiler/ironlang"
	"iron-lang/compiler/scopes"
	"log"
	"os"
	"strings"
)

func ExampleWrite(data []byte) {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, data)
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}
	fmt.Printf("% x", buf.Bytes())
}

func main() {

}

func mainxx() {

	//Lexical analysis
	is := antlr.NewInputStream("fn main() {let arr = int[1,2,3,4,5,6] let bb = arr.map((val int) int -> {return val * 2}).filter((val int) int -> {}).map((val int) int -> {return val * 2})}")
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

	scopes := scopes.NewScopesManager()

	customSemanticErrorListener := &errors.CustomErrorListener{}
	statics := compiler.NewSemanticAnalysis(scopes, customSemanticErrorListener)
	statics.Visit(tree)
	errors.HasSemanticError(customSemanticErrorListener.Errors)

	//Code generator
	generator := codegenerator.NewClang(scopes.GetScopeLog())
	generator.Visit(tree)
	println(generator.GetBuilder().String())
	NewFile(generator.GetBuilder())
}

func NewFile(strBuilder *strings.Builder) {
	f, err := os.Create("source.tmp")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	_, err = f.WriteString(strBuilder.String())

	if err != nil {
		log.Fatal(err)
	}

}
