package main

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"iron-lang/compiler/ironlang"
)

func main() {
	is := antlr.NewInputStream("fn main(){(2+2*3)-(3*9)}")
	lexer := ironlang.NewIronLangLexer(is)
	lexer.RemoveErrorListeners()
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
	parser.Program()
	//println(ip.GetText())
}
