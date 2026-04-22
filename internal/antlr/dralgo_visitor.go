// Code generated from /home/denis/projets/profdenis/dralgo/grammar/dralgo/DrAlgo.g4 by ANTLR 4.13.1. DO NOT EDIT.

package antlr // DrAlgo
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by DrAlgoParser.
type DrAlgoVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by DrAlgoParser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by DrAlgoParser#stmtList.
	VisitStmtList(ctx *StmtListContext) interface{}

	// Visit a parse tree produced by DrAlgoParser#stmt.
	VisitStmt(ctx *StmtContext) interface{}

	// Visit a parse tree produced by DrAlgoParser#assignStmt.
	VisitAssignStmt(ctx *AssignStmtContext) interface{}

	// Visit a parse tree produced by DrAlgoParser#incDecStmt.
	VisitIncDecStmt(ctx *IncDecStmtContext) interface{}

	// Visit a parse tree produced by DrAlgoParser#printStmt.
	VisitPrintStmt(ctx *PrintStmtContext) interface{}

	// Visit a parse tree produced by DrAlgoParser#printArgs.
	VisitPrintArgs(ctx *PrintArgsContext) interface{}

	// Visit a parse tree produced by DrAlgoParser#readStmt.
	VisitReadStmt(ctx *ReadStmtContext) interface{}

	// Visit a parse tree produced by DrAlgoParser#readIntStmt.
	VisitReadIntStmt(ctx *ReadIntStmtContext) interface{}

	// Visit a parse tree produced by DrAlgoParser#ifStmt.
	VisitIfStmt(ctx *IfStmtContext) interface{}

	// Visit a parse tree produced by DrAlgoParser#elseIfPart.
	VisitElseIfPart(ctx *ElseIfPartContext) interface{}

	// Visit a parse tree produced by DrAlgoParser#elsePart.
	VisitElsePart(ctx *ElsePartContext) interface{}

	// Visit a parse tree produced by DrAlgoParser#whileStmt.
	VisitWhileStmt(ctx *WhileStmtContext) interface{}

	// Visit a parse tree produced by DrAlgoParser#block.
	VisitBlock(ctx *BlockContext) interface{}

	// Visit a parse tree produced by DrAlgoParser#sep.
	VisitSep(ctx *SepContext) interface{}

	// Visit a parse tree produced by DrAlgoParser#expr.
	VisitExpr(ctx *ExprContext) interface{}

	// Visit a parse tree produced by DrAlgoParser#equality.
	VisitEquality(ctx *EqualityContext) interface{}

	// Visit a parse tree produced by DrAlgoParser#comparison.
	VisitComparison(ctx *ComparisonContext) interface{}

	// Visit a parse tree produced by DrAlgoParser#addition.
	VisitAddition(ctx *AdditionContext) interface{}

	// Visit a parse tree produced by DrAlgoParser#multiplication.
	VisitMultiplication(ctx *MultiplicationContext) interface{}

	// Visit a parse tree produced by DrAlgoParser#power.
	VisitPower(ctx *PowerContext) interface{}

	// Visit a parse tree produced by DrAlgoParser#unary.
	VisitUnary(ctx *UnaryContext) interface{}

	// Visit a parse tree produced by DrAlgoParser#primary.
	VisitPrimary(ctx *PrimaryContext) interface{}

	// Visit a parse tree produced by DrAlgoParser#listLiteral.
	VisitListLiteral(ctx *ListLiteralContext) interface{}

	// Visit a parse tree produced by DrAlgoParser#literal.
	VisitLiteral(ctx *LiteralContext) interface{}
}
