package ir

import (
	"dralgo/internal/ast"
)

type lowerer struct {
	nextID int
}

func Lower(prog *ast.Program) *Flowchart {
	l := &lowerer{nextID: 1}

	startNode := l.newNode(NodeAction)
	exitNode := l.newNode(NodeExit)

	last := l.lowerStmts(prog.Stmts, startNode, exitNode)
	last.Next = &Edge{Target: exitNode}

	return &Flowchart{Start: startNode}
}

func (l *lowerer) newNode(typ NodeType) *Node {
	n := &Node{
		ID:   l.nextID,
		Type: typ,
	}
	l.nextID++
	return n
}

func (l *lowerer) lowerStmts(stmts []ast.Stmt, current, exit *Node) *Node {
	for _, stmt := range stmts {
		switch s := stmt.(type) {
		case ast.IfStmt:
			current = l.lowerIf(s, current, exit)
		case ast.WhileStmt:
			current = l.lowerWhile(s, current, exit)
		case ast.DoWhileStmt:
			current = l.lowerDoWhile(s, current, exit)
		case ast.ForStmt:
			current = l.lowerFor(s, current, exit)
		case ast.ForEachStmt:
			current = l.lowerForEach(s, current, exit)
		default:
			// Regular statement (Assign, Print, Read, IncDec)
			// For simplicity, we group consecutive regular statements in the same Action node
			if current.Type == NodeAction {
				current.Stmts = append(current.Stmts, s)
			} else {
				next := l.newNode(NodeAction)
				next.Stmts = append(next.Stmts, s)
				l.connect(current, next)
				current = next
			}
		}
	}
	return current
}

func (l *lowerer) connect(from, to *Node) {
	if from.Type == NodeDecision {
		if from.Then == nil {
			from.Then = &Edge{Target: to, Label: "vrai"}
		} else {
			from.Else = &Edge{Target: to, Label: "faux"}
		}
	} else {
		from.Next = &Edge{Target: to}
	}
}

func (l *lowerer) lowerIf(s ast.IfStmt, current, exit *Node) *Node {
	// 1. Decision node
	decision := l.newNode(NodeDecision)
	decision.Cond = s.Cond
	l.connect(current, decision)

	// 2. Merge node at the end of IF
	merge := l.newNode(NodeMerge)

	// 3. Then branch
	thenStart := l.newNode(NodeAction)
	decision.Then = &Edge{Target: thenStart, Label: "vrai"}
	thenEnd := l.lowerStmts(s.Then, thenStart, exit)
	l.connect(thenEnd, merge)

	// 4. ElseIf / Else branches
	lastDecision := decision
	for _, elsif := range s.ElseIf {
		elsifDecision := l.newNode(NodeDecision)
		elsifDecision.Cond = elsif.Cond
		lastDecision.Else = &Edge{Target: elsifDecision, Label: "faux"}

		branchStart := l.newNode(NodeAction)
		elsifDecision.Then = &Edge{Target: branchStart, Label: "vrai"}
		branchEnd := l.lowerStmts(elsif.Body, branchStart, exit)
		l.connect(branchEnd, merge)

		lastDecision = elsifDecision
	}

	if len(s.Else) > 0 {
		elseStart := l.newNode(NodeAction)
		lastDecision.Else = &Edge{Target: elseStart, Label: "faux"}
		elseEnd := l.lowerStmts(s.Else, elseStart, exit)
		l.connect(elseEnd, merge)
	} else {
		lastDecision.Else = &Edge{Target: merge, Label: "faux"}
	}

	return merge
}

func (l *lowerer) lowerWhile(s ast.WhileStmt, current, exit *Node) *Node {
	decision := l.newNode(NodeDecision)
	decision.Cond = s.Cond
	l.connect(current, decision)

	bodyStart := l.newNode(NodeAction)
	decision.Then = &Edge{Target: bodyStart, Label: "vrai"}
	bodyEnd := l.lowerStmts(s.Body, bodyStart, exit)

	// Loop back
	l.connect(bodyEnd, decision)

	merge := l.newNode(NodeMerge)
	decision.Else = &Edge{Target: merge, Label: "faux"}

	return merge
}

func (l *lowerer) lowerDoWhile(s ast.DoWhileStmt, current, exit *Node) *Node {
	bodyStart := l.newNode(NodeAction)
	l.connect(current, bodyStart)
	bodyEnd := l.lowerStmts(s.Body, bodyStart, exit)

	decision := l.newNode(NodeDecision)
	decision.Cond = s.Cond
	l.connect(bodyEnd, decision)

	merge := l.newNode(NodeMerge)
	decision.Then = &Edge{Target: bodyStart, Label: "vrai"}
	decision.Else = &Edge{Target: merge, Label: "faux"}

	return merge
}

func (l *lowerer) lowerFor(s ast.ForStmt, current, exit *Node) *Node {
	if s.Init != nil {
		initNode := l.newNode(NodeAction)
		initNode.Stmts = []ast.Stmt{s.Init}
		l.connect(current, initNode)
		current = initNode
	}

	decision := l.newNode(NodeDecision)
	if s.Cond != nil {
		decision.Cond = s.Cond
	} else {
		// No condition means always true
		decision.Cond = ast.LiteralExpr{Kind: ast.LitInt, Text: "1"}
	}
	l.connect(current, decision)

	bodyStart := l.newNode(NodeAction)
	decision.Then = &Edge{Target: bodyStart, Label: "vrai"}
	bodyEnd := l.lowerStmts(s.Body, bodyStart, exit)

	if s.Post != nil {
		postNode := l.newNode(NodeAction)
		postNode.Stmts = []ast.Stmt{s.Post}
		l.connect(bodyEnd, postNode)
		bodyEnd = postNode
	}

	l.connect(bodyEnd, decision)

	merge := l.newNode(NodeMerge)
	decision.Else = &Edge{Target: merge, Label: "faux"}

	return merge
}

func (l *lowerer) lowerForEach(s ast.ForEachStmt, current, exit *Node) *Node {
	// Simplified lowering for ForEach: pretend it's a decision "has more items?"
	decision := l.newNode(NodeDecision)

	// Actually, let's just use the Iter expression as a placeholder
	decision.Cond = s.Iter

	l.connect(current, decision)

	bodyStart := l.newNode(NodeAction)
	decision.Then = &Edge{Target: bodyStart, Label: "vrai"}
	// We might need to inject the assignment "Var = next item" at bodyStart

	bodyEnd := l.lowerStmts(s.Body, bodyStart, exit)
	l.connect(bodyEnd, decision)

	merge := l.newNode(NodeMerge)
	decision.Else = &Edge{Target: merge, Label: "faux"}

	return merge
}
