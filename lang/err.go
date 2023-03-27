package lang

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

type OgaSyntaxError struct {
	line, column int
	msg          string
}

func (e *OgaSyntaxError) Error() string {
	return fmt.Sprintf("line %d:%d %s", e.line, e.column, e.msg)
}

type OgaErrorListener struct {
	*antlr.DefaultErrorListener
	Errors []error
}

func (l *OgaErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	l.Errors = append(l.Errors,
		&OgaSyntaxError{
			line:   line,
			column: column,
			msg:    msg,
		})
}

func (l *OgaErrorListener) HasErrors() bool {
	return len(l.Errors) > 0
}
