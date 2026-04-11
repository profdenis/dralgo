package parser

import (
	"fmt"

	antlr4 "github.com/antlr4-go/antlr/v4"
)

type ParseError struct {
	Line   int
	Column int
	Msg    string
}

func (e ParseError) Error() string {
	if e.Line > 0 {
		return fmt.Sprintf("parse error at %d:%d: %s", e.Line, e.Column, e.Msg)
	}
	return "parse error: " + e.Msg
}

type errorCollector struct {
	*antlr4.DefaultErrorListener
	errs []error
}

func newErrorCollector() *errorCollector {
	return &errorCollector{DefaultErrorListener: &antlr4.DefaultErrorListener{}}
}

func (l *errorCollector) Errors() []error {
	if len(l.errs) == 0 {
		return nil
	}
	out := make([]error, len(l.errs))
	copy(out, l.errs)
	return out
}

func (l *errorCollector) SyntaxError(_ antlr4.Recognizer, _ any, line, column int, msg string, _ antlr4.RecognitionException) {
	l.errs = append(l.errs, ParseError{Line: line, Column: column, Msg: msg})
}
