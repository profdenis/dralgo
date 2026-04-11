## Next Step Spec for Junie: AST + Parser Wrapper (Go + ANTLR)

### Goal
Introduce a **stable Go AST** and a **parser wrapper** that:
- hides ANTLR generated types behind `internal/parser`
- returns `(*ast.Program, []error)` (multiple errors)
- builds expressions **robustly** (operators collected via `GetTokens()` and ordered by token index)

This step does **not** do IR lowering or code generation yet.

---

## 1) New Packages / Files

Create these packages and files:

```plain text
internal/
  ast/
    nodes.go
  parser/
    parser.go
    errors.go
    visitor.go
    visitor_expr.go        # optional split, but recommended for readability
    parser_test.go
```


### Dependency rules
- `internal/parser` may import:
    - `internal/antlr` (generated)
    - `internal/ast`
    - ANTLR runtime (`github.com/antlr4-go/antlr/v4`)
- `internal/ast` must not import anything from parser / antlr.

---

## 2) AST (internal/ast)

### Requirements
Define minimal AST types:
- `Program { Stmts []Stmt }`
- Statements: `AssignStmt`, `IncDecStmt`, `PrintStmt`, `ReadStmt`, `ReadIntStmt`, `IfStmt` (+ `ElseIfClause`), `WhileStmt`, `DoWhileStmt`, `ForStmt`, `ForEachStmt`
- Expressions: `IdentExpr`, `LiteralExpr` (`int/float/string`), `ListExpr`, `UnaryExpr`, `BinaryExpr`
- Optional but recommended: `Span` (line/col range) stored on every node

**Important design**: keep numeric literals as `Text string` initially (preserve underscores). Normalization/strconv can happen later.

Deliverable: `internal/ast/nodes.go`

---

## 3) Error collection (internal/parser/errors.go)

### Requirements
- Provide a parse error type with position:
    - `ParseError{Line, Column, Msg}`
- Provide an ANTLR error listener that collects errors into a slice
- We will return **all** collected errors (lexer + parser + visitor errors if any)

Deliverable: `internal/parser/errors.go`

---

## 4) Parser wrapper API (internal/parser/parser.go)

### Public API
Implement:

```textmate
func ParseProgram(input string) (*ast.Program, []error)
```


### Behavior
1. Create `antlr.NewInputStream(input)`
2. Create generated lexer: `gen.NewDrAlgoLexer(is)`
3. Attach error collector to lexer (`RemoveErrorListeners`, `AddErrorListener`)
4. Create token stream `antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)`
5. Create generated parser `gen.NewDrAlgoParser(stream)`
6. Attach error collector to parser
7. Parse: `tree := parser.Program()`
8. If lexer/parser errors exist:
    - still try building AST **only if** tree is non-nil (optional)
    - but always return errors (do not panic)
9. Convert parse tree → AST via `astBuilder`
10. Return `prog, errors` where errors contains:
    - lexer errors
    - parser errors
    - visitor/build errors (if any)

Deliverable: `internal/parser/parser.go`

---

## 5) AST builder (internal/parser/visitor.go)

### Structure
Create a type:

```textmate
type astBuilder struct {
  tokens antlr.TokenStream
  errs   []error
}
```


Implement methods like:
- `visitProgram(ctx gen.IProgramContext) *ast.Program`
- `visitStmtList(ctx gen.IStmtListContext) []ast.Stmt`
- `visitStmt(ctx gen.IStmtContext) ast.Stmt`
- plus one method per statement rule
- expression entrypoint: `visitExpr(ctx gen.IExprContext) ast.Expr`

### Span computation
Add helper:

```textmate
func (b *astBuilder) spanOf(ctx antlr.ParserRuleContext) ast.Span
```


Use `ctx.GetStart()` / `ctx.GetStop()` token line/col.

Deliverable: `internal/parser/visitor.go`

---

## 6) Expression building: robust operator handling using GetTokens (internal/parser/visitor_expr.go)

### Requirement
For rules like:

```
equality : comparison ((EQ | NEQ) comparison)* ;
comparison : addition ((LT | LTE | GT | GTE) addition)* ;
addition : multiplication ((PLUS | MINUS) multiplication)* ;
multiplication : power ((MUL | DIV | MOD) power)* ;
```


We must build a left-associative AST:
- Use the operand list, e.g. `ctx.AllComparison()`, `ctx.AllAddition()`, etc.
- Extract operators using `GetTokens(tokenType)` (NOT `GetChild()` string hacks)
- When multiple operator types exist in the same rule (`EQ` or `NEQ`), gather both lists and **merge them in source order** using token indices.

### Implementation pattern

#### A) Helper: merge terminal nodes by token order
Use token index to sort operators:

```textmate
// pseudo
opsEQ  := ctx.GetTokens(gen.DrAlgoParserEQ)
opsNEQ := ctx.GetTokens(gen.DrAlgoParserNEQ)

allOps := append(opsEQ, opsNEQ...)
sort.Slice(allOps, func(i,j int) bool {
  return allOps[i].GetSymbol().GetTokenIndex() < allOps[j].GetSymbol().GetTokenIndex()
})
```


Then you can fold:

- operands count = `len(parts)`
- ops count should be `len(parts)-1`
- op at position `i` applies between operand `i` and `i+1`

Validate counts; if mismatch, push an internal error into `b.errs`.

#### B) Single-operator-type rules
For `addition`, you can do:

```textmate
opsPlus  := ctx.GetTokens(gen.DrAlgoParserPLUS)
opsMinus := ctx.GetTokens(gen.DrAlgoParserMINUS)
merge by token index
```


Or if you prefer, keep `PLUS/MINUS` merged the same way.

#### C) Power rule
`power : unary (POW power)?;` is already explicit; read `ctx.POW()` / `ctx.Power()` and build `BinaryExpr{Op:"**"}` right-associative.

#### D) Unary rule
`unary : (PLUS|MINUS) unary | primary;`
Use `ctx.PLUS()` / `ctx.MINUS()` checks (or `GetTokens` is not needed here).

Deliverable: `internal/parser/visitor_expr.go`

---

## 7) Tests (internal/parser/parser_test.go)

### Add at least two tests

1) **Valid program** parses and returns `len(errs)==0` and `prog != nil`:
- Use a small sample with `if / sinon si / sinon / fin si` and one assignment.

2) **Invalid program** returns errors:
- e.g. missing `fin si`

Assertions:
- `errs != nil && len(errs) > 0`
- errors contain meaningful line/col (optional)

Run:
```shell script
go test ./internal/parser -run TestParseProgram
```


Deliverable: `internal/parser/parser_test.go`

---

## 8) Acceptance Criteria (Definition of Done)

- `go test ./...` passes
- `internal/parser.ParseProgram` exists and is used by tests
- Parser returns multiple errors (`[]error`) without panics
- Expression operator extraction is done using `GetTokens()` and merged by token index (robust ordering)
- No code outside `internal/parser` imports `internal/antlr`

---

## Notes / Gotchas

- ANTLR token type `EOF` is `-1`; do not index arrays with it (already fixed in debug tool)
- When merging token lists by index, make sure you’re using `TerminalNode.GetSymbol().GetTokenIndex()`
- Keep literals as text for now (`"1_000.0_5"`); parse into numeric values later in a semantic pass or during codegen


## Go Skeletons

Below are **ready-to-paste Go skeletons** for the AST-parsing step (multi-error API + 100% robust operator handling via `GetTokens()` ordered by token index). They’re written to drop into your project as-is and then you fill in the remaining statement/expression cases incrementally.

---

## 1) `internal/parser/errors.go` (multi-error collection)

```textmate
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
```


---

## 2) `internal/parser/parser.go` (public API: `(*ast.Program, []error)`)

```textmate
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
```


---

## 3) `internal/parser/visitor.go` (AST builder skeleton + spans + basic statements)

This is intentionally “minimum but functional”: it supports program + statement list + a set of statements, and delegates all expressions to `visitor_expr.go`.

```textmate
package parser

import (
	"fmt"

	antlr4 "github.com/antlr4-go/antlr/v4"

	gen "dralgo/internal/antlr"
	"dralgo/internal/ast"
)

type astBuilder struct {
	tokens antlr4.TokenStream
	errs   []error
}

func newASTBuilder(tokens antlr4.TokenStream) *astBuilder {
	return &astBuilder{tokens: tokens}
}

func (b *astBuilder) Errors() []error {
	if len(b.errs) == 0 {
		return nil
	}
	out := make([]error, len(b.errs))
	copy(out, b.errs)
	return out
}

func (b *astBuilder) addErr(err error) {
	if err != nil {
		b.errs = append(b.errs, err)
	}
}

func (b *astBuilder) spanOf(ctx antlr4.ParserRuleContext) ast.Span {
	start := ctx.GetStart()
	stop := ctx.GetStop()
	if start == nil || stop == nil {
		return ast.Span{}
	}
	// EndCol is best-effort. Good enough for diagnostics.
	return ast.Span{
		StartLine: start.GetLine(),
		StartCol:  start.GetColumn(),
		EndLine:   stop.GetLine(),
		EndCol:    stop.GetColumn() + len(stop.GetText()),
	}
}

func (b *astBuilder) visitProgram(ctx gen.IProgramContext) *ast.Program {
	prc, _ := ctx.(antlr4.ParserRuleContext)
	p := &ast.Program{Span: b.spanOf(prc)}
	if ctx.StmtList() != nil {
		p.Stmts = b.visitStmtList(ctx.StmtList())
	}
	return p
}

func (b *astBuilder) visitStmtList(ctx gen.IStmtListContext) []ast.Stmt {
	out := make([]ast.Stmt, 0, len(ctx.AllStmt()))
	for _, s := range ctx.AllStmt() {
		out = append(out, b.visitStmt(s))
	}
	return out
}

func (b *astBuilder) visitStmt(ctx gen.IStmtContext) ast.Stmt {
	switch {
	case ctx.AssignStmt() != nil:
		return b.visitAssign(ctx.AssignStmt())
	case ctx.IncDecStmt() != nil:
		return b.visitIncDec(ctx.IncDecStmt())
	case ctx.PrintStmt() != nil:
		return b.visitPrint(ctx.PrintStmt())
	case ctx.ReadStmt() != nil:
		return b.visitRead(ctx.ReadStmt())
	case ctx.ReadIntStmt() != nil:
		return b.visitReadInt(ctx.ReadIntStmt())
	case ctx.IfStmt() != nil:
		return b.visitIf(ctx.IfStmt())
	case ctx.WhileStmt() != nil:
		return b.visitWhile(ctx.WhileStmt())
	case ctx.DoWhileStmt() != nil:
		return b.visitDoWhile(ctx.DoWhileStmt())
	case ctx.ForStmt() != nil:
		return b.visitFor(ctx.ForStmt())
	case ctx.ForEachStmt() != nil:
		return b.visitForEach(ctx.ForEachStmt())
	default:
		// Should not happen if grammar and generated parser match.
		b.addErr(ParseError{Msg: "internal error: unknown statement variant"})
		return ast.PrintStmt{Args: nil}
	}
}

func (b *astBuilder) visitAssign(ctx gen.IAssignStmtContext) ast.Stmt {
	prc := ctx.(antlr4.ParserRuleContext)
	return ast.AssignStmt{
		Name:  ctx.IDENT().GetText(),
		Value: b.visitExpr(ctx.Expr()),
		Span:  b.spanOf(prc),
	}
}

func (b *astBuilder) visitIncDec(ctx gen.IIncDecStmtContext) ast.Stmt {
	prc := ctx.(antlr4.ParserRuleContext)
	op := ast.Inc
	if ctx.DECR() != nil {
		op = ast.Dec
	}
	return ast.IncDecStmt{
		Name: ctx.IDENT().GetText(),
		Op:   op,
		Span: b.spanOf(prc),
	}
}

func (b *astBuilder) visitPrint(ctx gen.IPrintStmtContext) ast.Stmt {
	prc := ctx.(antlr4.ParserRuleContext)
	var args []ast.Expr
	if ctx.PrintArgs() != nil {
		exprs := ctx.PrintArgs().AllExpr()
		args = make([]ast.Expr, 0, len(exprs))
		for _, e := range exprs {
			args = append(args, b.visitExpr(e))
		}
	}
	return ast.PrintStmt{Args: args, Span: b.spanOf(prc)}
}

func (b *astBuilder) visitRead(ctx gen.IReadStmtContext) ast.Stmt {
	prc := ctx.(antlr4.ParserRuleContext)
	return ast.ReadStmt{Name: ctx.IDENT().GetText(), Span: b.spanOf(prc)}
}

func (b *astBuilder) visitReadInt(ctx gen.IReadIntStmtContext) ast.Stmt {
	prc := ctx.(antlr4.ParserRuleContext)
	return ast.ReadIntStmt{Name: ctx.IDENT().GetText(), Span: b.spanOf(prc)}
}

func (b *astBuilder) visitIf(ctx gen.IIfStmtContext) ast.Stmt {
	prc := ctx.(antlr4.ParserRuleContext)

	thenBody := b.visitStmtList(ctx.Block().StmtList())

	elsifs := make([]ast.ElseIfClause, 0, len(ctx.AllElseIfPart()))
	for _, e := range ctx.AllElseIfPart() {
		eprc := e.(antlr4.ParserRuleContext)
		elsifs = append(elsifs, ast.ElseIfClause{
			Cond: b.visitExpr(e.Expr()),
			Body: b.visitStmtList(e.Block().StmtList()),
			Span: b.spanOf(eprc),
		})
	}

	var elseBody []ast.Stmt
	if ctx.ElsePart() != nil {
		elseBody = b.visitStmtList(ctx.ElsePart().Block().StmtList())
	}

	return ast.IfStmt{
		Cond:   b.visitExpr(ctx.Expr()),
		Then:   thenBody,
		ElseIf: elsifs,
		Else:   elseBody,
		Span:   b.spanOf(prc),
	}
}

func (b *astBuilder) visitWhile(ctx gen.IWhileStmtContext) ast.Stmt {
	prc := ctx.(antlr4.ParserRuleContext)
	return ast.WhileStmt{
		Cond: b.visitExpr(ctx.Expr()),
		Body: b.visitStmtList(ctx.Block().StmtList()),
		Span: b.spanOf(prc),
	}
}

func (b *astBuilder) visitDoWhile(ctx gen.IDoWhileStmtContext) ast.Stmt {
	prc := ctx.(antlr4.ParserRuleContext)
	return ast.DoWhileStmt{
		Body: b.visitStmtList(ctx.Block().StmtList()),
		Cond: b.visitExpr(ctx.Expr()),
		Span: b.spanOf(prc),
	}
}

func (b *astBuilder) visitFor(ctx gen.IForStmtContext) ast.Stmt {
	prc := ctx.(antlr4.ParserRuleContext)

	var init ast.Stmt
	if ctx.ForInit() != nil {
		init = b.visitAssign(ctx.ForInit().AssignStmt())
	}

	var cond ast.Expr
	if ctx.ForCond() != nil {
		cond = b.visitExpr(ctx.ForCond().Expr())
	}

	var post ast.Stmt
	if ctx.ForPost() != nil {
		switch {
		case ctx.ForPost().IncDecStmt() != nil:
			post = b.visitIncDec(ctx.ForPost().IncDecStmt())
		case ctx.ForPost().AssignStmt() != nil:
			post = b.visitAssign(ctx.ForPost().AssignStmt())
		default:
			b.addErr(ParseError{Msg: "internal error: unexpected forPost variant"})
		}
	}

	return ast.ForStmt{
		Init: init,
		Cond: cond,
		Post: post,
		Body: b.visitStmtList(ctx.Block().StmtList()),
		Span: b.spanOf(prc),
	}
}

func (b *astBuilder) visitForEach(ctx gen.IForEachStmtContext) ast.Stmt {
	prc := ctx.(antlr4.ParserRuleContext)
	return ast.ForEachStmt{
		Var:  ctx.IDENT().GetText(),
		Iter: b.visitExpr(ctx.Expr()),
		Body: b.visitStmtList(ctx.Block().StmtList()),
		Span: b.spanOf(prc),
	}
}

// Tiny internal helper for unexpected token/operator situations.
func (b *astBuilder) internalErrf(format string, args ...any) {
	b.addErr(ParseError{Msg: "internal: " + fmt.Sprintf(format, args...)})
}
```


---

## 4) `internal/parser/visitor_expr.go` (100% robust operators via `GetTokens()`)

This file is the key piece you requested: it never depends on `GetChild()` hacks. It merges operator token lists in **source order** via `TokenIndex`.

```textmate
package parser

import (
	"sort"

	antlr4 "github.com/antlr4-go/antlr/v4"

	gen "dralgo/internal/antlr"
	"dralgo/internal/ast"
)

// ---- operator collection helpers ----

type opTok struct {
	Index int
	Text  string
}

func collectOpsByIndex(nodes []antlr4.TerminalNode) []opTok {
	out := make([]opTok, 0, len(nodes))
	for _, n := range nodes {
		if n == nil || n.GetSymbol() == nil {
			continue
		}
		out = append(out, opTok{
			Index: n.GetSymbol().GetTokenIndex(),
			Text:  n.GetText(),
		})
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Index < out[j].Index })
	return out
}

func mergeOpsByIndex(groups ...[]antlr4.TerminalNode) []opTok {
	var merged []antlr4.TerminalNode
	for _, g := range groups {
		merged = append(merged, g...)
	}
	return collectOpsByIndex(merged)
}

// foldLeft folds a sequence like: x op y op z into ((x op y) op z).
// parts must have length >= 1; ops must have length == len(parts)-1.
func (b *astBuilder) foldLeft(parts []ast.Expr, ops []opTok, mk func(op string, lhs, rhs ast.Expr) ast.Expr) ast.Expr {
	if len(parts) == 0 {
		b.internalErrf("foldLeft called with 0 parts")
		return ast.LiteralExpr{Kind: ast.LitInt, Text: "0"}
	}
	if len(ops) != len(parts)-1 {
		b.internalErrf("operator/operand mismatch: ops=%d parts=%d", len(ops), len(parts))
		// best-effort: return first part
		return parts[0]
	}
	acc := parts[0]
	for i := 0; i < len(ops); i++ {
		acc = mk(ops[i].Text, acc, parts[i+1])
	}
	return acc
}

// ---- expressions ----

func (b *astBuilder) visitExpr(ctx gen.IExprContext) ast.Expr {
	// expr : equality ;
	return b.visitEquality(ctx.Equality())
}

func (b *astBuilder) visitEquality(ctx gen.IEqualityContext) ast.Expr {
	comps := ctx.AllComparison()
	parts := make([]ast.Expr, 0, len(comps))
	for _, c := range comps {
		parts = append(parts, b.visitComparison(c))
	}

	ops := mergeOpsByIndex(
		ctx.GetTokens(gen.DrAlgoParserEQ),
		ctx.GetTokens(gen.DrAlgoParserNEQ),
	)

	return b.foldLeft(parts, ops, func(op string, lhs, rhs ast.Expr) ast.Expr {
		var bop ast.BinaryOp
		switch op {
		case "==":
			bop = ast.OpEQ
		case "!=":
			bop = ast.OpNEQ
		default:
			b.internalErrf("unexpected equality operator %q", op)
			bop = ast.OpEQ
		}
		return ast.BinaryExpr{Op: bop, LHS: lhs, RHS: rhs, Span: lhs.GetSpan()}
	})
}

func (b *astBuilder) visitComparison(ctx gen.IComparisonContext) ast.Expr {
	adds := ctx.AllAddition()
	parts := make([]ast.Expr, 0, len(adds))
	for _, a := range adds {
		parts = append(parts, b.visitAddition(a))
	}

	ops := mergeOpsByIndex(
		ctx.GetTokens(gen.DrAlgoParserLT),
		ctx.GetTokens(gen.DrAlgoParserLTE),
		ctx.GetTokens(gen.DrAlgoParserGT),
		ctx.GetTokens(gen.DrAlgoParserGTE),
	)

	return b.foldLeft(parts, ops, func(op string, lhs, rhs ast.Expr) ast.Expr {
		var bop ast.BinaryOp
		switch op {
		case "<":
			bop = ast.OpLT
		case "<=":
			bop = ast.OpLTE
		case ">":
			bop = ast.OpGT
		case ">=":
			bop = ast.OpGTE
		default:
			b.internalErrf("unexpected comparison operator %q", op)
			bop = ast.OpLT
		}
		return ast.BinaryExpr{Op: bop, LHS: lhs, RHS: rhs, Span: lhs.GetSpan()}
	})
}

func (b *astBuilder) visitAddition(ctx gen.IAdditionContext) ast.Expr {
	muls := ctx.AllMultiplication()
	parts := make([]ast.Expr, 0, len(muls))
	for _, m := range muls {
		parts = append(parts, b.visitMultiplication(m))
	}

	ops := mergeOpsByIndex(
		ctx.GetTokens(gen.DrAlgoParserPLUS),
		ctx.GetTokens(gen.DrAlgoParserMINUS),
	)

	return b.foldLeft(parts, ops, func(op string, lhs, rhs ast.Expr) ast.Expr {
		bop := ast.OpAdd
		if op == "-" {
			bop = ast.OpSub
		} else if op != "+" {
			b.internalErrf("unexpected addition operator %q", op)
		}
		return ast.BinaryExpr{Op: bop, LHS: lhs, RHS: rhs, Span: lhs.GetSpan()}
	})
}

func (b *astBuilder) visitMultiplication(ctx gen.IMultiplicationContext) ast.Expr {
	pows := ctx.AllPower()
	parts := make([]ast.Expr, 0, len(pows))
	for _, p := range pows {
		parts = append(parts, b.visitPower(p))
	}

	ops := mergeOpsByIndex(
		ctx.GetTokens(gen.DrAlgoParserMUL),
		ctx.GetTokens(gen.DrAlgoParserDIV),
		ctx.GetTokens(gen.DrAlgoParserMOD),
	)

	return b.foldLeft(parts, ops, func(op string, lhs, rhs ast.Expr) ast.Expr {
		var bop ast.BinaryOp
		switch op {
		case "*":
			bop = ast.OpMul
		case "/":
			bop = ast.OpDiv
		case "%":
			bop = ast.OpMod
		default:
			b.internalErrf("unexpected multiplication operator %q", op)
			bop = ast.OpMul
		}
		return ast.BinaryExpr{Op: bop, LHS: lhs, RHS: rhs, Span: lhs.GetSpan()}
	})
}

func (b *astBuilder) visitPower(ctx gen.IPowerContext) ast.Expr {
	// power : unary (POW power)?;
	left := b.visitUnary(ctx.Unary())
	if ctx.Power() == nil {
		return left
	}
	right := b.visitPower(ctx.Power())
	return ast.BinaryExpr{Op: ast.OpPow, LHS: left, RHS: right, Span: left.GetSpan()}
}

func (b *astBuilder) visitUnary(ctx gen.IUnaryContext) ast.Expr {
	if ctx.Primary() != nil {
		return b.visitPrimary(ctx.Primary())
	}

	// unary : (PLUS | MINUS) unary
	opNodes := mergeOpsByIndex(
		ctx.GetTokens(gen.DrAlgoParserPLUS),
		ctx.GetTokens(gen.DrAlgoParserMINUS),
	)
	if len(opNodes) == 0 {
		b.internalErrf("unary operator missing")
		return b.visitUnary(ctx.Unary())
	}
	op := opNodes[0].Text

	uop := ast.OpPos
	if op == "-" {
		uop = ast.OpNeg
	} else if op != "+" {
		b.internalErrf("unexpected unary operator %q", op)
	}

	x := b.visitUnary(ctx.Unary())
	return ast.UnaryExpr{Op: uop, X: x, Span: x.GetSpan()}
}

func (b *astBuilder) visitPrimary(ctx gen.IPrimaryContext) ast.Expr {
	prc := ctx.(antlr4.ParserRuleContext)
	span := b.spanOf(prc)

	switch {
	case ctx.IDENT() != nil:
		return ast.IdentExpr{Name: ctx.IDENT().GetText(), Span: span}

	case ctx.Literal() != nil:
		return b.visitLiteral(ctx.Literal())

	case ctx.ListLiteral() != nil:
		return b.visitListLiteral(ctx.ListLiteral())

	case ctx.Expr() != nil:
		// Parenthesized expression
		return b.visitExpr(ctx.Expr())

	default:
		b.internalErrf("unknown primary expression variant")
		return ast.LiteralExpr{Kind: ast.LitInt, Text: "0", Span: span}
	}
}

func (b *astBuilder) visitListLiteral(ctx gen.IListLiteralContext) ast.Expr {
	prc := ctx.(antlr4.ParserRuleContext)
	span := b.spanOf(prc)

	var elems []ast.Expr
	for _, e := range ctx.AllExpr() {
		elems = append(elems, b.visitExpr(e))
	}
	return ast.ListExpr{Elems: elems, Span: span}
}

func (b *astBuilder) visitLiteral(ctx gen.ILiteralContext) ast.Expr {
	prc := ctx.(antlr4.ParserRuleContext)
	span := b.spanOf(prc)

	switch {
	case ctx.INT() != nil:
		return ast.LiteralExpr{Kind: ast.LitInt, Text: ctx.INT().GetText(), Span: span}
	case ctx.FLOAT() != nil:
		return ast.LiteralExpr{Kind: ast.LitFloat, Text: ctx.FLOAT().GetText(), Span: span}
	case ctx.STRING() != nil:
		return ast.LiteralExpr{Kind: ast.LitString, Text: ctx.STRING().GetText(), Span: span}
	default:
		b.internalErrf("unknown literal variant")
		return ast.LiteralExpr{Kind: ast.LitInt, Text: "0", Span: span}
	}
}
```


**Note:** I used `Span: lhs.GetSpan()` / `x.GetSpan()` to keep the skeleton small. If you prefer accurate spans for binary/unary nodes (covering both sides), add a helper to merge spans.

---

## 5) `internal/parser/parser_test.go` (smoke tests for multi-error API)

```textmate
package parser

import "testing"

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

func TestParseProgram_Invalid_ReturnsErrors(t *testing.T) {
	input := `si 1 < 2 alors
    Afficher "ok"
`
	_, errs := ParseProgram(input)
	if len(errs) == 0 {
		t.Fatal("expected errors, got none")
	}
}
```


Run:

```shell script
go test ./internal/parser -run TestParseProgram
```


###  minimal internal/ast/nodes.go

```go
package ast

// Span identifies a source location (best-effort).
type Span struct {
	StartLine int
	StartCol  int
	EndLine   int
	EndCol    int
}

// ----- Program -----

type Program struct {
	Stmts []Stmt
	Span  Span
}

// ----- Statements -----

type Stmt interface {
	isStmt()
	GetSpan() Span
}

type AssignStmt struct {
	Name  string
	Value Expr
	Span  Span
}

func (AssignStmt) isStmt()         {}
func (s AssignStmt) GetSpan() Span { return s.Span }

type IncDecOp string

const (
	Inc IncDecOp = "++"
	Dec IncDecOp = "--"
)

type IncDecStmt struct {
	Name string
	Op   IncDecOp
	Span Span
}

func (IncDecStmt) isStmt()         {}
func (s IncDecStmt) GetSpan() Span { return s.Span }

type PrintStmt struct {
	Args []Expr
	Span Span
}

func (PrintStmt) isStmt()         {}
func (s PrintStmt) GetSpan() Span { return s.Span }

type ReadStmt struct {
	Name string
	Span Span
}

func (ReadStmt) isStmt()         {}
func (s ReadStmt) GetSpan() Span { return s.Span }

type ReadIntStmt struct {
	Name string
	Span Span
}

func (ReadIntStmt) isStmt()         {}
func (s ReadIntStmt) GetSpan() Span { return s.Span }

type ElseIfClause struct {
	Cond Expr
	Body []Stmt
	Span Span
}

type IfStmt struct {
	Cond   Expr
	Then   []Stmt
	ElseIf []ElseIfClause
	Else   []Stmt // nil if absent
	Span   Span
}

func (IfStmt) isStmt()         {}
func (s IfStmt) GetSpan() Span { return s.Span }

type WhileStmt struct {
	Cond Expr
	Body []Stmt
	Span Span
}

func (WhileStmt) isStmt()         {}
func (s WhileStmt) GetSpan() Span { return s.Span }

type DoWhileStmt struct {
	Body []Stmt
	Cond Expr
	Span Span
}

func (DoWhileStmt) isStmt()         {}
func (s DoWhileStmt) GetSpan() Span { return s.Span }

type ForStmt struct {
	Init Stmt // nil or AssignStmt
	Cond Expr // nil allowed by grammar
	Post Stmt // nil or AssignStmt/IncDecStmt
	Body []Stmt
	Span Span
}

func (ForStmt) isStmt()         {}
func (s ForStmt) GetSpan() Span { return s.Span }

type ForEachStmt struct {
	Var  string
	Iter Expr
	Body []Stmt
	Span Span
}

func (ForEachStmt) isStmt()         {}
func (s ForEachStmt) GetSpan() Span { return s.Span }

// ----- Expressions -----

type Expr interface {
	isExpr()
	GetSpan() Span
}

type IdentExpr struct {
	Name string
	Span Span
}

func (IdentExpr) isExpr()         {}
func (e IdentExpr) GetSpan() Span { return e.Span }

type LiteralKind string

const (
	LitInt    LiteralKind = "int"
	LitFloat  LiteralKind = "float"
	LitString LiteralKind = "string"
)

type LiteralExpr struct {
	Kind LiteralKind
	Text string // Keep original token text (e.g. "1_000.0_5")
	Span Span
}

func (LiteralExpr) isExpr()         {}
func (e LiteralExpr) GetSpan() Span { return e.Span }

type ListExpr struct {
	Elems []Expr
	Span  Span
}

func (ListExpr) isExpr()         {}
func (e ListExpr) GetSpan() Span { return e.Span }

type UnaryOp string

const (
	OpPos UnaryOp = "+"
	OpNeg UnaryOp = "-"
)

type UnaryExpr struct {
	Op   UnaryOp
	X    Expr
	Span Span
}

func (UnaryExpr) isExpr()         {}
func (e UnaryExpr) GetSpan() Span { return e.Span }

type BinaryOp string

const (
	OpPow BinaryOp = "**"

	OpMul BinaryOp = "*"
	OpDiv BinaryOp = "/"
	OpMod BinaryOp = "%"

	OpAdd BinaryOp = "+"
	OpSub BinaryOp = "-"

	OpLT  BinaryOp = "<"
	OpLTE BinaryOp = "<="
	OpGT  BinaryOp = ">"
	OpGTE BinaryOp = ">="

	OpEQ  BinaryOp = "=="
	OpNEQ BinaryOp = "!="
)

type BinaryExpr struct {
	Op   BinaryOp
	LHS  Expr
	RHS  Expr
	Span Span
}

func (BinaryExpr) isExpr()         {}
func (e BinaryExpr) GetSpan() Span { return e.Span }
```

