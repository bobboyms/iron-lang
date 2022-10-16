package test

import (
	"compiler-course/mycompiler/compiler/ironlang"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"testing"
)

func TestParseNumber(t *testing.T) {
	is := antlr.NewInputStream("2 2 2")
	lexer := ironlang.NewIronLangLexer(is)

	for {
		t := lexer.NextToken()
		if t.GetTokenType() == antlr.TokenEOF {
			break
		}
		fmt.Printf("%s (%q)\n",
			lexer.SymbolicNames[t.GetTokenType()], t.GetText())
	}

	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := ironlang.NewIronLangParser(stream)
	parser.BuildParseTrees = true

}
