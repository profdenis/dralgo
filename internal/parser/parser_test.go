package parser

import (
	"dralgo/internal/ast"
	"errors"
	"testing"
)

func TestParseProgram_Valid(t *testing.T) {
	input := `x = 2**10
si x < 0 alors
    Afficher "Négatif"
sinon si x > 1_000_000 alors
    Afficher "Trop grand"
sinon
    Afficher "Bien"
fin si
`
	prog, errs := ParseProgram(input)
	if len(errs) != 0 {
		t.Fatalf("expected no errors, got %d: %v", len(errs), errs[0])
	}
	if prog == nil {
		t.Fatal("prog is nil")
	}
	if len(prog.Stmts) == 0 {
		t.Fatal("expected non-empty program")
	}
}

func TestParseProgram_Expressions(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		validate func(*testing.T, *ast.Program)
	}{
		{
			name:  "operator precedence",
			input: `x = 1 + 2 * 3 ** 2 ** 2`,
			validate: func(t *testing.T, prog *ast.Program) {
				// x = 1 + (2 * (3 ** (2 ** 2)))
				assign := prog.Stmts[0].(ast.AssignStmt)
				add := assign.Value.(ast.BinaryExpr)
				if add.Op != ast.OpAdd {
					t.Errorf("expected +, got %v", add.Op)
				}
				mul := add.RHS.(ast.BinaryExpr)
				if mul.Op != ast.OpMul {
					t.Errorf("expected *, got %v", mul.Op)
				}
				pow1 := mul.RHS.(ast.BinaryExpr)
				if pow1.Op != ast.OpPow {
					t.Errorf("expected **, got %v", pow1.Op)
				}
				pow2 := pow1.RHS.(ast.BinaryExpr)
				if pow2.Op != ast.OpPow {
					t.Errorf("expected **, got %v", pow2.Op)
				}
			},
		},
		{
			name:  "associativity",
			input: `x = 10 - 5 - 2`,
			validate: func(t *testing.T, prog *ast.Program) {
				// (10 - 5) - 2
				assign := prog.Stmts[0].(ast.AssignStmt)
				sub1 := assign.Value.(ast.BinaryExpr)
				if sub1.Op != ast.OpSub {
					t.Errorf("expected -, got %v", sub1.Op)
				}
				sub2 := sub1.LHS.(ast.BinaryExpr)
				if sub2.Op != ast.OpSub {
					t.Errorf("expected -, got %v", sub2.Op)
				}
			},
		},
		{
			name:  "unary and parens",
			input: `x = -(1 + 2) * +3`,
			validate: func(t *testing.T, prog *ast.Program) {
				assign := prog.Stmts[0].(ast.AssignStmt)
				mul := assign.Value.(ast.BinaryExpr)
				neg := mul.LHS.(ast.UnaryExpr)
				if neg.Op != ast.OpNeg {
					t.Errorf("expected -, got %v", neg.Op)
				}
				pos := mul.RHS.(ast.UnaryExpr)
				if pos.Op != ast.OpPos {
					t.Errorf("expected +, got %v", pos.Op)
				}
			},
		},
		{
			name:  "list literals",
			input: `l = [1, 2 + 3, [4]]`,
			validate: func(t *testing.T, prog *ast.Program) {
				assign := prog.Stmts[0].(ast.AssignStmt)
				list := assign.Value.(ast.ListExpr)
				if len(list.Elems) != 3 {
					t.Fatalf("expected 3 elements, got %d", len(list.Elems))
				}
				if _, ok := list.Elems[2].(ast.ListExpr); !ok {
					t.Errorf("expected nested list, got %T", list.Elems[2])
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			prog, errs := ParseProgram(tt.input)
			if len(errs) != 0 {
				t.Fatalf("unexpected errors: %v", errs)
			}
			tt.validate(t, prog)
		})
	}
}

func TestParseProgram_Loops(t *testing.T) {
	input := `
tant que x > 0
    x = x - 1
fin tant que
`
	prog, errs := ParseProgram(input)
	if len(errs) != 0 {
		t.Fatalf("unexpected errors: %v", errs)
	}

	if len(prog.Stmts) != 1 {
		t.Fatalf("expected 1 statement, got %d", len(prog.Stmts))
	}

	if _, ok := prog.Stmts[0].(ast.WhileStmt); !ok {
		t.Errorf("stmt 0 is not WhileStmt, got %T", prog.Stmts[0])
	}
}

func TestParseProgram_Misc(t *testing.T) {
	input := `
// This is a comment
Afficher "Hello", 123
Lire nom
Lire entier age
`
	prog, errs := ParseProgram(input)
	if len(errs) != 0 {
		t.Fatalf("unexpected errors: %v", errs)
	}

	if len(prog.Stmts) != 3 {
		t.Fatalf("expected 3 statements, got %d", len(prog.Stmts))
	}
	printStmt := prog.Stmts[0].(ast.PrintStmt)
	if len(printStmt.Args) != 2 {
		t.Errorf("expected 2 print args, got %d", len(printStmt.Args))
	}
	if _, ok := prog.Stmts[1].(ast.ReadStmt); !ok {
		t.Errorf("stmt 1 is not ReadStmt, got %T", prog.Stmts[1])
	}
	if _, ok := prog.Stmts[2].(ast.ReadIntStmt); !ok {
		t.Errorf("stmt 2 is not ReadIntStmt, got %T", prog.Stmts[2])
	}
}

func TestParseProgram_Invalid_ReturnsErrors(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{"missing fin si", `si true alors Afficher 1`},
		{"bad expression", `x = 1 + * 2`},
		{"unclosed string", `Afficher "unterminated`},
		{"extra characters", `x = 1 }`},
		{"invalid assignment read", `x = Lire Entier`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, errs := ParseProgram(tt.input)
			if len(errs) == 0 {
				t.Fatalf("%s: expected errors for input %q, got none", tt.name, tt.input)
			}
		})
	}
}

func TestParseProgram_Position(t *testing.T) {
	input := `x = 1
y = * 2
`
	_, errs := ParseProgram(input)
	if len(errs) == 0 {
		t.Fatal("expected errors, got none")
	}

	var pe ParseError
	found := false
	for _, err := range errs {
		if errors.As(err, &pe) {
			if pe.Line == 2 {
				found = true
				break
			}
		}
	}
	if !found {
		t.Errorf("expected error on line 2, got errors: %v", errs)
	}
}
