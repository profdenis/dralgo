package parser

import (
	antlr4 "github.com/antlr4-go/antlr/v4"

	gen "dralgo/internal/antlr"
	"dralgo/internal/ast"
)

// ParseProgram parses DrAlgo pseudo-code and returns an AST program plus *all* collected errors.
// It tries to build an AST even if parse errors exist (best-effort), but callers should treat
// any non-empty errs slice as "invalid program".
func ParseProgram(input string) (*ast.Program, []error) {
	var allErrs []error

	is := antlr4.NewInputStream(input)

	lexer := gen.NewDrAlgoLexer(is)
	lexErrs := newErrorCollector()
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(lexErrs)

	stream := antlr4.NewCommonTokenStream(lexer, antlr4.TokenDefaultChannel)

	p := gen.NewDrAlgoParser(stream)
	parseErrs := newErrorCollector()
	p.RemoveErrorListeners()
	p.AddErrorListener(parseErrs)

	tree := p.Program()

	allErrs = append(allErrs, lexErrs.Errors()...)
	allErrs = append(allErrs, parseErrs.Errors()...)

	// Best-effort AST build
	if tree == nil {
		allErrs = append(allErrs, ParseError{Msg: "internal error: parser returned nil parse tree"})
		return nil, allErrs
	}

	b := newASTBuilder(stream)
	prog := b.visitProgram(tree)
	allErrs = append(allErrs, b.Errors()...)

	return prog, allErrs
}
