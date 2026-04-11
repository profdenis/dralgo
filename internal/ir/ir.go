package ir

import "dralgo/internal/ast"

type Flowchart struct {
	Start *Node
}

type NodeType int

const (
	NodeAction NodeType = iota
	NodeDecision
	NodeMerge
	NodeExit
)

type Node struct {
	ID   int
	Type NodeType

	// For NodeAction
	Stmts []ast.Stmt

	// For NodeDecision
	Cond ast.Expr

	// Connections
	Next *Edge // For Action, Merge
	Then *Edge // For Decision
	Else *Edge // For Decision
}

type Edge struct {
	Target *Node
	Label  string
}
