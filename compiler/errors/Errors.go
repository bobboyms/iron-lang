package errors

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

type CustomSyntaxError struct {
	Line   int
	Column int
	Msg    string
}

type CustomErrorListener struct {
	*antlr.DefaultErrorListener
	Errors []*CustomSyntaxError
}

func (c *CustomErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	c.Errors = append(c.Errors, &CustomSyntaxError{
		Line:   line,
		Column: column,
		Msg:    msg,
	})
}
