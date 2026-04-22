// Code generated from grammar/dralgo/DrAlgo.g4 by ANTLR 4.13.1. DO NOT EDIT.

package antlr // DrAlgo
import "github.com/antlr4-go/antlr/v4"

// DrAlgoListener is a complete listener for a parse tree produced by DrAlgoParser.
type DrAlgoListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterStmtList is called when entering the stmtList production.
	EnterStmtList(c *StmtListContext)

	// EnterStmt is called when entering the stmt production.
	EnterStmt(c *StmtContext)

	// EnterAssignStmt is called when entering the assignStmt production.
	EnterAssignStmt(c *AssignStmtContext)

	// EnterIncDecStmt is called when entering the incDecStmt production.
	EnterIncDecStmt(c *IncDecStmtContext)

	// EnterPrintStmt is called when entering the printStmt production.
	EnterPrintStmt(c *PrintStmtContext)

	// EnterPrintArgs is called when entering the printArgs production.
	EnterPrintArgs(c *PrintArgsContext)

	// EnterReadStmt is called when entering the readStmt production.
	EnterReadStmt(c *ReadStmtContext)

	// EnterReadIntStmt is called when entering the readIntStmt production.
	EnterReadIntStmt(c *ReadIntStmtContext)

	// EnterIfStmt is called when entering the ifStmt production.
	EnterIfStmt(c *IfStmtContext)

	// EnterElseIfPart is called when entering the elseIfPart production.
	EnterElseIfPart(c *ElseIfPartContext)

	// EnterElsePart is called when entering the elsePart production.
	EnterElsePart(c *ElsePartContext)

	// EnterWhileStmt is called when entering the whileStmt production.
	EnterWhileStmt(c *WhileStmtContext)

	// EnterDoWhileStmt is called when entering the doWhileStmt production.
	EnterDoWhileStmt(c *DoWhileStmtContext)

	// EnterForStmt is called when entering the forStmt production.
	EnterForStmt(c *ForStmtContext)

	// EnterForInit is called when entering the forInit production.
	EnterForInit(c *ForInitContext)

	// EnterForCond is called when entering the forCond production.
	EnterForCond(c *ForCondContext)

	// EnterForPost is called when entering the forPost production.
	EnterForPost(c *ForPostContext)

	// EnterForEachStmt is called when entering the forEachStmt production.
	EnterForEachStmt(c *ForEachStmtContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterSep is called when entering the sep production.
	EnterSep(c *SepContext)

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// EnterEquality is called when entering the equality production.
	EnterEquality(c *EqualityContext)

	// EnterComparison is called when entering the comparison production.
	EnterComparison(c *ComparisonContext)

	// EnterAddition is called when entering the addition production.
	EnterAddition(c *AdditionContext)

	// EnterMultiplication is called when entering the multiplication production.
	EnterMultiplication(c *MultiplicationContext)

	// EnterPower is called when entering the power production.
	EnterPower(c *PowerContext)

	// EnterUnary is called when entering the unary production.
	EnterUnary(c *UnaryContext)

	// EnterPrimary is called when entering the primary production.
	EnterPrimary(c *PrimaryContext)

	// EnterListLiteral is called when entering the listLiteral production.
	EnterListLiteral(c *ListLiteralContext)

	// EnterLiteral is called when entering the literal production.
	EnterLiteral(c *LiteralContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitStmtList is called when exiting the stmtList production.
	ExitStmtList(c *StmtListContext)

	// ExitStmt is called when exiting the stmt production.
	ExitStmt(c *StmtContext)

	// ExitAssignStmt is called when exiting the assignStmt production.
	ExitAssignStmt(c *AssignStmtContext)

	// ExitIncDecStmt is called when exiting the incDecStmt production.
	ExitIncDecStmt(c *IncDecStmtContext)

	// ExitPrintStmt is called when exiting the printStmt production.
	ExitPrintStmt(c *PrintStmtContext)

	// ExitPrintArgs is called when exiting the printArgs production.
	ExitPrintArgs(c *PrintArgsContext)

	// ExitReadStmt is called when exiting the readStmt production.
	ExitReadStmt(c *ReadStmtContext)

	// ExitReadIntStmt is called when exiting the readIntStmt production.
	ExitReadIntStmt(c *ReadIntStmtContext)

	// ExitIfStmt is called when exiting the ifStmt production.
	ExitIfStmt(c *IfStmtContext)

	// ExitElseIfPart is called when exiting the elseIfPart production.
	ExitElseIfPart(c *ElseIfPartContext)

	// ExitElsePart is called when exiting the elsePart production.
	ExitElsePart(c *ElsePartContext)

	// ExitWhileStmt is called when exiting the whileStmt production.
	ExitWhileStmt(c *WhileStmtContext)

	// ExitDoWhileStmt is called when exiting the doWhileStmt production.
	ExitDoWhileStmt(c *DoWhileStmtContext)

	// ExitForStmt is called when exiting the forStmt production.
	ExitForStmt(c *ForStmtContext)

	// ExitForInit is called when exiting the forInit production.
	ExitForInit(c *ForInitContext)

	// ExitForCond is called when exiting the forCond production.
	ExitForCond(c *ForCondContext)

	// ExitForPost is called when exiting the forPost production.
	ExitForPost(c *ForPostContext)

	// ExitForEachStmt is called when exiting the forEachStmt production.
	ExitForEachStmt(c *ForEachStmtContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitSep is called when exiting the sep production.
	ExitSep(c *SepContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)

	// ExitEquality is called when exiting the equality production.
	ExitEquality(c *EqualityContext)

	// ExitComparison is called when exiting the comparison production.
	ExitComparison(c *ComparisonContext)

	// ExitAddition is called when exiting the addition production.
	ExitAddition(c *AdditionContext)

	// ExitMultiplication is called when exiting the multiplication production.
	ExitMultiplication(c *MultiplicationContext)

	// ExitPower is called when exiting the power production.
	ExitPower(c *PowerContext)

	// ExitUnary is called when exiting the unary production.
	ExitUnary(c *UnaryContext)

	// ExitPrimary is called when exiting the primary production.
	ExitPrimary(c *PrimaryContext)

	// ExitListLiteral is called when exiting the listLiteral production.
	ExitListLiteral(c *ListLiteralContext)

	// ExitLiteral is called when exiting the literal production.
	ExitLiteral(c *LiteralContext)
}
