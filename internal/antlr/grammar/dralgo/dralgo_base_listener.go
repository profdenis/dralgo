// Code generated from grammar/dralgo/DrAlgo.g4 by ANTLR 4.13.1. DO NOT EDIT.

package antlr // DrAlgo
import "github.com/antlr4-go/antlr/v4"

// BaseDrAlgoListener is a complete listener for a parse tree produced by DrAlgoParser.
type BaseDrAlgoListener struct{}

var _ DrAlgoListener = &BaseDrAlgoListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseDrAlgoListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseDrAlgoListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseDrAlgoListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseDrAlgoListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseDrAlgoListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseDrAlgoListener) ExitProgram(ctx *ProgramContext) {}

// EnterStmtList is called when production stmtList is entered.
func (s *BaseDrAlgoListener) EnterStmtList(ctx *StmtListContext) {}

// ExitStmtList is called when production stmtList is exited.
func (s *BaseDrAlgoListener) ExitStmtList(ctx *StmtListContext) {}

// EnterStmt is called when production stmt is entered.
func (s *BaseDrAlgoListener) EnterStmt(ctx *StmtContext) {}

// ExitStmt is called when production stmt is exited.
func (s *BaseDrAlgoListener) ExitStmt(ctx *StmtContext) {}

// EnterAssignStmt is called when production assignStmt is entered.
func (s *BaseDrAlgoListener) EnterAssignStmt(ctx *AssignStmtContext) {}

// ExitAssignStmt is called when production assignStmt is exited.
func (s *BaseDrAlgoListener) ExitAssignStmt(ctx *AssignStmtContext) {}

// EnterIncDecStmt is called when production incDecStmt is entered.
func (s *BaseDrAlgoListener) EnterIncDecStmt(ctx *IncDecStmtContext) {}

// ExitIncDecStmt is called when production incDecStmt is exited.
func (s *BaseDrAlgoListener) ExitIncDecStmt(ctx *IncDecStmtContext) {}

// EnterPrintStmt is called when production printStmt is entered.
func (s *BaseDrAlgoListener) EnterPrintStmt(ctx *PrintStmtContext) {}

// ExitPrintStmt is called when production printStmt is exited.
func (s *BaseDrAlgoListener) ExitPrintStmt(ctx *PrintStmtContext) {}

// EnterPrintArgs is called when production printArgs is entered.
func (s *BaseDrAlgoListener) EnterPrintArgs(ctx *PrintArgsContext) {}

// ExitPrintArgs is called when production printArgs is exited.
func (s *BaseDrAlgoListener) ExitPrintArgs(ctx *PrintArgsContext) {}

// EnterReadStmt is called when production readStmt is entered.
func (s *BaseDrAlgoListener) EnterReadStmt(ctx *ReadStmtContext) {}

// ExitReadStmt is called when production readStmt is exited.
func (s *BaseDrAlgoListener) ExitReadStmt(ctx *ReadStmtContext) {}

// EnterReadIntStmt is called when production readIntStmt is entered.
func (s *BaseDrAlgoListener) EnterReadIntStmt(ctx *ReadIntStmtContext) {}

// ExitReadIntStmt is called when production readIntStmt is exited.
func (s *BaseDrAlgoListener) ExitReadIntStmt(ctx *ReadIntStmtContext) {}

// EnterIfStmt is called when production ifStmt is entered.
func (s *BaseDrAlgoListener) EnterIfStmt(ctx *IfStmtContext) {}

// ExitIfStmt is called when production ifStmt is exited.
func (s *BaseDrAlgoListener) ExitIfStmt(ctx *IfStmtContext) {}

// EnterElseIfPart is called when production elseIfPart is entered.
func (s *BaseDrAlgoListener) EnterElseIfPart(ctx *ElseIfPartContext) {}

// ExitElseIfPart is called when production elseIfPart is exited.
func (s *BaseDrAlgoListener) ExitElseIfPart(ctx *ElseIfPartContext) {}

// EnterElsePart is called when production elsePart is entered.
func (s *BaseDrAlgoListener) EnterElsePart(ctx *ElsePartContext) {}

// ExitElsePart is called when production elsePart is exited.
func (s *BaseDrAlgoListener) ExitElsePart(ctx *ElsePartContext) {}

// EnterWhileStmt is called when production whileStmt is entered.
func (s *BaseDrAlgoListener) EnterWhileStmt(ctx *WhileStmtContext) {}

// ExitWhileStmt is called when production whileStmt is exited.
func (s *BaseDrAlgoListener) ExitWhileStmt(ctx *WhileStmtContext) {}

// EnterDoWhileStmt is called when production doWhileStmt is entered.
func (s *BaseDrAlgoListener) EnterDoWhileStmt(ctx *DoWhileStmtContext) {}

// ExitDoWhileStmt is called when production doWhileStmt is exited.
func (s *BaseDrAlgoListener) ExitDoWhileStmt(ctx *DoWhileStmtContext) {}

// EnterForStmt is called when production forStmt is entered.
func (s *BaseDrAlgoListener) EnterForStmt(ctx *ForStmtContext) {}

// ExitForStmt is called when production forStmt is exited.
func (s *BaseDrAlgoListener) ExitForStmt(ctx *ForStmtContext) {}

// EnterForInit is called when production forInit is entered.
func (s *BaseDrAlgoListener) EnterForInit(ctx *ForInitContext) {}

// ExitForInit is called when production forInit is exited.
func (s *BaseDrAlgoListener) ExitForInit(ctx *ForInitContext) {}

// EnterForCond is called when production forCond is entered.
func (s *BaseDrAlgoListener) EnterForCond(ctx *ForCondContext) {}

// ExitForCond is called when production forCond is exited.
func (s *BaseDrAlgoListener) ExitForCond(ctx *ForCondContext) {}

// EnterForPost is called when production forPost is entered.
func (s *BaseDrAlgoListener) EnterForPost(ctx *ForPostContext) {}

// ExitForPost is called when production forPost is exited.
func (s *BaseDrAlgoListener) ExitForPost(ctx *ForPostContext) {}

// EnterForEachStmt is called when production forEachStmt is entered.
func (s *BaseDrAlgoListener) EnterForEachStmt(ctx *ForEachStmtContext) {}

// ExitForEachStmt is called when production forEachStmt is exited.
func (s *BaseDrAlgoListener) ExitForEachStmt(ctx *ForEachStmtContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseDrAlgoListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseDrAlgoListener) ExitBlock(ctx *BlockContext) {}

// EnterSep is called when production sep is entered.
func (s *BaseDrAlgoListener) EnterSep(ctx *SepContext) {}

// ExitSep is called when production sep is exited.
func (s *BaseDrAlgoListener) ExitSep(ctx *SepContext) {}

// EnterExpr is called when production expr is entered.
func (s *BaseDrAlgoListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BaseDrAlgoListener) ExitExpr(ctx *ExprContext) {}

// EnterEquality is called when production equality is entered.
func (s *BaseDrAlgoListener) EnterEquality(ctx *EqualityContext) {}

// ExitEquality is called when production equality is exited.
func (s *BaseDrAlgoListener) ExitEquality(ctx *EqualityContext) {}

// EnterComparison is called when production comparison is entered.
func (s *BaseDrAlgoListener) EnterComparison(ctx *ComparisonContext) {}

// ExitComparison is called when production comparison is exited.
func (s *BaseDrAlgoListener) ExitComparison(ctx *ComparisonContext) {}

// EnterAddition is called when production addition is entered.
func (s *BaseDrAlgoListener) EnterAddition(ctx *AdditionContext) {}

// ExitAddition is called when production addition is exited.
func (s *BaseDrAlgoListener) ExitAddition(ctx *AdditionContext) {}

// EnterMultiplication is called when production multiplication is entered.
func (s *BaseDrAlgoListener) EnterMultiplication(ctx *MultiplicationContext) {}

// ExitMultiplication is called when production multiplication is exited.
func (s *BaseDrAlgoListener) ExitMultiplication(ctx *MultiplicationContext) {}

// EnterPower is called when production power is entered.
func (s *BaseDrAlgoListener) EnterPower(ctx *PowerContext) {}

// ExitPower is called when production power is exited.
func (s *BaseDrAlgoListener) ExitPower(ctx *PowerContext) {}

// EnterUnary is called when production unary is entered.
func (s *BaseDrAlgoListener) EnterUnary(ctx *UnaryContext) {}

// ExitUnary is called when production unary is exited.
func (s *BaseDrAlgoListener) ExitUnary(ctx *UnaryContext) {}

// EnterPrimary is called when production primary is entered.
func (s *BaseDrAlgoListener) EnterPrimary(ctx *PrimaryContext) {}

// ExitPrimary is called when production primary is exited.
func (s *BaseDrAlgoListener) ExitPrimary(ctx *PrimaryContext) {}

// EnterListLiteral is called when production listLiteral is entered.
func (s *BaseDrAlgoListener) EnterListLiteral(ctx *ListLiteralContext) {}

// ExitListLiteral is called when production listLiteral is exited.
func (s *BaseDrAlgoListener) ExitListLiteral(ctx *ListLiteralContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BaseDrAlgoListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BaseDrAlgoListener) ExitLiteral(ctx *LiteralContext) {}
