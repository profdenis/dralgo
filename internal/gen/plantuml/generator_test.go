package plantuml

import (
	"strings"
	"testing"

	"dralgo/internal/ir"
	"dralgo/internal/parser"
)

func TestGenerate_KeepsParenthesesForPrecedence(t *testing.T) {
	prog, errs := parser.ParseProgram("a = 10\nb = 3\nd = (a + b) * 2")
	if len(errs) > 0 {
		t.Fatalf("parse errors: %v", errs)
	}

	out, err := NewGenerator().Generate(ir.Lower(prog))
	if err != nil {
		t.Fatalf("Generate error: %v", err)
	}

	if !strings.Contains(out, "d = (a + b) * 2") {
		t.Errorf("expected parentheses to be kept, got:\n%s", out)
	}
}
