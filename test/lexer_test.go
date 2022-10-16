package test

import (
	"compiler-course/mycompiler/compiler/ironlang"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"testing"
)

func TestIdentifier(t *testing.T) {
	is := antlr.NewInputStream("2.tali_25ba 2 * 35.6 - (10 * 20) / 3")
	lexer := ironlang.NewIronLangLexer(is)

	for {
		t := lexer.NextToken()
		if t.GetTokenType() == antlr.TokenEOF {
			break
		}
		fmt.Printf("%s (%q)\n",
			lexer.SymbolicNames[t.GetTokenType()], t.GetText())
	}
}
