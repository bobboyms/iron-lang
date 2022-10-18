package errors

import (
	"fmt"
	"log"
)

func HasLexerError(errors []*CustomSyntaxError) {
	if len(errors) > 0 {
		log.Println("HAS LEXER ERRORS")
		for _, err := range errors {
			msg := fmt.Sprintf("Error line: %v and colounm: %v message: %v", err.Line, err.Column, err.Msg)
			log.Println(msg)
		}
		log.Fatalln("Compiler exit")
	}
}

func HasParserError(errors []*CustomSyntaxError) {

	if len(errors) > 0 {
		log.Println("HAS SYNTACTIC ERRORS")
		for _, err := range errors {
			msg := fmt.Sprintf("Error line: %v and colounm: %v message: %v", err.Line, err.Column, err.Msg)
			log.Println(msg)
		}
		log.Fatalln("Compiler exit")
	}
}

func HasSemanticError(errors []*CustomSyntaxError) {
	if len(errors) > 0 {
		log.Println("HAS SEMANTIC ERRORS")
		for _, err := range errors {
			msg := fmt.Sprintf("Error line: %v and colounm: %v message: %v", err.Line, err.Column, err.Msg)
			log.Println(msg)
		}
		log.Fatalln("Compiler exit")
	}
}
