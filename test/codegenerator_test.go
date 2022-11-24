package test

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"iron-lang/compiler"
	"iron-lang/compiler/codegenerator"
	"iron-lang/compiler/errors"
	"iron-lang/compiler/ironlang"
	"iron-lang/compiler/scopes"
	"log"
	"os"
	"os/exec"
	"strings"
	"testing"
)

func NewFile(strBuilder *strings.Builder) {
	f, err := os.Create("source.tmp")

	if err != nil {
		log.Fatal(err)
	}

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(f)

	_, err = f.WriteString(strBuilder.String())

	if err != nil {
		log.Fatal(err)
	}

}

func Compiler(code string) {
	is := antlr.NewInputStream(code)
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

func RunMake() {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(path)

	e := exec.Command("make")
	e.Dir = path
	var out bytes.Buffer
	e.Stdout = &out
	err = e.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Output: %q\n", out.String())
}

func TestCodeGeneratorFunctions(t *testing.T) {

	f, err := os.Open("functions.ir")

	if err != nil {
		log.Fatal(err)
	}

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

		}
	}(f)

	scanner := bufio.NewScanner(f)

	builder := strings.Builder{}
	for scanner.Scan() {
		builder.WriteString(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	Compiler(builder.String())
	RunMake()

}

func TestCodeGeneratorMapFilterReduce(t *testing.T) {

	f, err := os.Open("map_filter_reduce.ir")

	if err != nil {
		log.Fatal(err)
	}

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

		}
	}(f)

	scanner := bufio.NewScanner(f)

	builder := strings.Builder{}
	for scanner.Scan() {
		builder.WriteString(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	Compiler(builder.String())
	RunMake()

}

func TestCodeGeneratorControlFlow(t *testing.T) {

	f, err := os.Open("control_flow.ir")

	if err != nil {
		log.Fatal(err)
	}

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

		}
	}(f)

	scanner := bufio.NewScanner(f)

	builder := strings.Builder{}
	for scanner.Scan() {
		builder.WriteString(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	Compiler(builder.String())
	RunMake()

}

func TestCodeGeneratorSlice(t *testing.T) {

	f, err := os.Open("slice.ir")

	if err != nil {
		log.Fatal(err)
	}

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

		}
	}(f)

	scanner := bufio.NewScanner(f)

	builder := strings.Builder{}
	for scanner.Scan() {
		builder.WriteString(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	Compiler(builder.String())
	RunMake()

}
