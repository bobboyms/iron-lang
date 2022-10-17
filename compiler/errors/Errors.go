package errors

import "github.com/antlr/antlr4/runtime/Go/antlr"

type DataError struct {
	Line   int
	Column int
	Msg    string
}

type CustomSyntaxError struct {
	line   int
	column int
	msg    string
}

func (c *CustomSyntaxError) GetData() DataError {
	return DataError{
		Line:   c.line,
		Column: c.column,
		Msg:    c.msg,
	}
}

func (c *CustomSyntaxError) Error() string {
	return c.msg
}

type CustomErrorListener struct {
	*antlr.DefaultErrorListener
	Errors []error
}

func (c *CustomErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	c.Errors = append(c.Errors, &CustomSyntaxError{
		line:   line,
		column: column,
		msg:    msg,
	})
}
