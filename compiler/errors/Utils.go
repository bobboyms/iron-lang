package errors

import (
	"fmt"
	"log"
)

func HasParserError(errors []error) {
	if len(errors) > 0 {
		for _, err := range errors {
			syntaxError := err.(*CustomSyntaxError).GetData()
			fmt.Printf("Error line %v and colounm: %v message: %v\n", syntaxError.Line, syntaxError.Column, syntaxError.Msg)
		}
		log.Fatalln("Has syntactic error")
	}
}

func HasLexerError(errors []error) {
	if len(errors) > 0 {
		for _, err := range errors {
			syntaxError := err.(*CustomSyntaxError).GetData()
			fmt.Printf("Error line %v and colounm: %v message: %v\n", syntaxError.Line, syntaxError.Column, syntaxError.Msg)
		}
		log.Fatalln("Has lexer error")
	}
}
