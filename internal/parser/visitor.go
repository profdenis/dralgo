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

	var thenBody []ast.Stmt
	if ctx.Block() != nil && ctx.Block().StmtList() != nil {
		thenBody = b.visitStmtList(ctx.Block().StmtList())
	}

	elsifs := make([]ast.ElseIfClause, 0, len(ctx.AllElseIfPart()))
	for _, e := range ctx.AllElseIfPart() {
		eprc := e.(antlr4.ParserRuleContext)
		var body []ast.Stmt
		if e.Block() != nil && e.Block().StmtList() != nil {
			body = b.visitStmtList(e.Block().StmtList())
		}
		elsifs = append(elsifs, ast.ElseIfClause{
			Cond: b.visitExpr(e.Expr()),
			Body: body,
			Span: b.spanOf(eprc),
		})
	}

	var elseBody []ast.Stmt
	if ctx.ElsePart() != nil && ctx.ElsePart().Block() != nil && ctx.ElsePart().Block().StmtList() != nil {
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
	var body []ast.Stmt
	if ctx.Block() != nil && ctx.Block().StmtList() != nil {
		body = b.visitStmtList(ctx.Block().StmtList())
	}
	return ast.WhileStmt{
		Cond: b.visitExpr(ctx.Expr()),
		Body: body,
		Span: b.spanOf(prc),
	}
}

// Tiny internal helper for unexpected token/operator situations.
func (b *astBuilder) internalErrf(format string, args ...any) {
	b.addErr(ParseError{Msg: "internal: " + fmt.Sprintf(format, args...)})
}
