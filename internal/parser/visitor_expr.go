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
	if ctx == nil {
		return ast.LiteralExpr{Kind: ast.LitInt, Text: "0"}
	}
	// expr : equality ;
	return b.visitEquality(ctx.Equality())
}

func (b *astBuilder) visitEquality(ctx gen.IEqualityContext) ast.Expr {
	if ctx == nil {
		return ast.LiteralExpr{Kind: ast.LitInt, Text: "0"}
	}
	comps := ctx.AllComparison()
	parts := make([]ast.Expr, 0, len(comps))
	for _, c := range comps {
		parts = append(parts, b.visitComparison(c))
	}

	ops := mergeOpsByIndex(
		ctx.AllEQ(),
		ctx.AllNEQ(),
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
	if ctx == nil {
		return ast.LiteralExpr{Kind: ast.LitInt, Text: "0"}
	}
	adds := ctx.AllAddition()
	parts := make([]ast.Expr, 0, len(adds))
	for _, a := range adds {
		parts = append(parts, b.visitAddition(a))
	}

	ops := mergeOpsByIndex(
		ctx.AllLT(),
		ctx.AllLTE(),
		ctx.AllGT(),
		ctx.AllGTE(),
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
	if ctx == nil {
		return ast.LiteralExpr{Kind: ast.LitInt, Text: "0"}
	}
	muls := ctx.AllMultiplication()
	parts := make([]ast.Expr, 0, len(muls))
	for _, m := range muls {
		parts = append(parts, b.visitMultiplication(m))
	}

	ops := mergeOpsByIndex(
		ctx.AllPLUS(),
		ctx.AllMINUS(),
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
	if ctx == nil {
		return ast.LiteralExpr{Kind: ast.LitInt, Text: "0"}
	}
	pows := ctx.AllPower()
	parts := make([]ast.Expr, 0, len(pows))
	for _, p := range pows {
		parts = append(parts, b.visitPower(p))
	}

	ops := mergeOpsByIndex(
		ctx.AllMUL(),
		ctx.AllDIV(),
		ctx.AllMOD(),
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
	if ctx == nil {
		return ast.LiteralExpr{Kind: ast.LitInt, Text: "0"}
	}
	// power : unary (POW power)?;
	left := b.visitUnary(ctx.Unary())
	if ctx.Power() == nil {
		return left
	}
	right := b.visitPower(ctx.Power())
	return ast.BinaryExpr{Op: ast.OpPow, LHS: left, RHS: right, Span: left.GetSpan()}
}

func (b *astBuilder) visitUnary(ctx gen.IUnaryContext) ast.Expr {
	if ctx == nil {
		return ast.LiteralExpr{Kind: ast.LitInt, Text: "0"}
	}
	if ctx.Primary() != nil {
		return b.visitPrimary(ctx.Primary())
	}

	// unary : (PLUS | MINUS) unary
	var opNode antlr4.TerminalNode
	if ctx.PLUS() != nil {
		opNode = ctx.PLUS()
	} else if ctx.MINUS() != nil {
		opNode = ctx.MINUS()
	}

	if opNode == nil {
		b.internalErrf("unary operator missing")
		return b.visitUnary(ctx.Unary())
	}
	op := opNode.GetText()

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
	if ctx == nil {
		return ast.LiteralExpr{Kind: ast.LitInt, Text: "0"}
	}
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
	if ctx == nil {
		return ast.ListExpr{}
	}
	prc := ctx.(antlr4.ParserRuleContext)
	span := b.spanOf(prc)

	var elems []ast.Expr
	for _, e := range ctx.AllExpr() {
		elems = append(elems, b.visitExpr(e))
	}
	return ast.ListExpr{Elems: elems, Span: span}
}

func (b *astBuilder) visitLiteral(ctx gen.ILiteralContext) ast.Expr {
	if ctx == nil {
		return ast.LiteralExpr{Kind: ast.LitInt, Text: "0"}
	}
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
