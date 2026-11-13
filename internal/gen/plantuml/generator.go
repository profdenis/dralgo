package plantuml

import (
	"bytes"
	"dralgo/internal/ast"
	"dralgo/internal/ir"
	"fmt"
	"strings"
)

type Generator struct {
	visited map[int]bool
}

func NewGenerator() *Generator {
	return &Generator{
		visited: make(map[int]bool),
	}
}

func (g *Generator) Generate(flow *ir.Flowchart) (string, error) {
	g.visited = make(map[int]bool)
	var buf bytes.Buffer
	buf.WriteString("@startuml\n")
	buf.WriteString("start\n")

	g.genNode(&buf, flow.Start, nil)

	buf.WriteString("stop\n")
	buf.WriteString("@enduml\n")
	return buf.String(), nil
}

func (g *Generator) genNode(buf *bytes.Buffer, n *ir.Node, stopAt *ir.Node) {
	if n == nil || g.visited[n.ID] || n == stopAt {
		return
	}

	switch n.Type {
	case ir.NodeExit:
		return

	case ir.NodeAction:
		for _, stmt := range n.Stmts {
			buf.WriteString(":")
			buf.WriteString(g.stmtStr(stmt))
			buf.WriteString(";\n")
		}
		if n.Next != nil {
			g.genNode(buf, n.Next.Target, stopAt)
		}

	case ir.NodeMerge:
		if n.Next != nil {
			g.genNode(buf, n.Next.Target, stopAt)
		}

	case ir.NodeDecision:
		// Check if it's a loop
		isLoop := g.isReachable(n.Then.Target, n)

		if isLoop {
			mergeNode := n.Else.Target
			buf.WriteString(fmt.Sprintf("while (%s) is (vrai)\n", g.exprStr(n.Cond)))
			g.visited[n.ID] = true
			g.genNode(buf, n.Then.Target, n)
			g.visited[n.ID] = false
			buf.WriteString("endwhile (faux)\n")

			if mergeNode != nil {
				g.genNode(buf, mergeNode, stopAt)
			}
		} else {
			mergeNode := g.findMergeNode(n)
			buf.WriteString(fmt.Sprintf("if (%s) then (vrai)\n", g.exprStr(n.Cond)))
			if n.Then != nil {
				g.genNode(buf, n.Then.Target, mergeNode)
			}

			if n.Else != nil {
				// Elif optimization
				if n.Else.Target.Type == ir.NodeDecision && !g.isReachable(n.Else.Target.Then.Target, n.Else.Target) {
					buf.WriteString(fmt.Sprintf("(faux) elseif (%s) then (vrai)\n", g.exprStr(n.Else.Target.Cond)))
					g.visited[n.Else.Target.ID] = true
					if n.Else.Target.Then != nil {
						g.genNode(buf, n.Else.Target.Then.Target, mergeNode)
					}
					if n.Else.Target.Else != nil {
						g.visited[n.Else.Target.ID] = false
						g.genElse(buf, n.Else.Target.Else.Target, mergeNode)
					}
				} else if n.Else.Target.Type == ir.NodeMerge && n.Else.Target.Next == nil && n.Else.Target == mergeNode {
					buf.WriteString("endif\n")
				} else {
					buf.WriteString("else (faux)\n")
					g.genNode(buf, n.Else.Target, mergeNode)
					buf.WriteString("endif\n")
				}
			} else {
				buf.WriteString("endif\n")
			}

			if mergeNode != nil {
				g.genNode(buf, mergeNode, stopAt)
			}
		}
	}
}

func (g *Generator) genElse(buf *bytes.Buffer, n *ir.Node, stopAt *ir.Node) {
	if n == nil || g.visited[n.ID] || n == stopAt {
		return
	}

	if n.Type == ir.NodeDecision && !g.isReachable(n.Then.Target, n) {
		buf.WriteString(fmt.Sprintf("(faux) elseif (%s) then (vrai)\n", g.exprStr(n.Cond)))
		g.visited[n.ID] = true
		if n.Then != nil {
			g.genNode(buf, n.Then.Target, stopAt)
		}
		if n.Else != nil {
			g.visited[n.ID] = false
			g.genElse(buf, n.Else.Target, stopAt)
		} else {
			buf.WriteString("endif\n")
		}
	} else if n.Type == ir.NodeMerge && n.Next == nil && n == stopAt {
		buf.WriteString("endif\n")
	} else {
		buf.WriteString("else (faux)\n")
		g.genNode(buf, n, stopAt)
		buf.WriteString("endif\n")
	}
}

func (g *Generator) findMergeNode(n *ir.Node) *ir.Node {
	// Simple heuristic: follow Then branch until we hit a Merge node
	// This works because of how our IR is constructed in lower.go
	curr := n.Then.Target
	visited := make(map[int]bool)
	for curr != nil && !visited[curr.ID] {
		if curr.Type == ir.NodeMerge {
			return curr
		}
		visited[curr.ID] = true
		if curr.Next != nil {
			curr = curr.Next.Target
		} else if curr.Type == ir.NodeDecision {
			// For nested decisions, we need to be careful.
			// The structured IR ensures there is one merge node.
			// Let's just follow one branch.
			curr = curr.Then.Target
		} else {
			break
		}
	}
	return nil
}

func (g *Generator) isReachable(start, target *ir.Node) bool {
	visited := make(map[int]bool)
	var stack []*ir.Node
	stack = append(stack, start)

	for len(stack) > 0 {
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if curr == nil {
			continue
		}
		if curr == target {
			return true
		}
		if visited[curr.ID] {
			continue
		}
		visited[curr.ID] = true

		if curr.Next != nil {
			stack = append(stack, curr.Next.Target)
		}
		if curr.Then != nil {
			stack = append(stack, curr.Then.Target)
		}
		if curr.Else != nil {
			stack = append(stack, curr.Else.Target)
		}
	}
	return false
}

func (g *Generator) stmtStr(stmt ast.Stmt) string {
	switch s := stmt.(type) {
	case ast.AssignStmt:
		return fmt.Sprintf("%s = %s", s.Name, g.exprStr(s.Value))
	case ast.PrintStmt:
		args := make([]string, len(s.Args))
		for i, a := range s.Args {
			args[i] = g.exprStr(a)
		}
		return fmt.Sprintf("Afficher %s", strings.Join(args, ", "))
	case ast.ReadStmt:
		return fmt.Sprintf("Lire %s", s.Name)
	case ast.ReadIntStmt:
		return fmt.Sprintf("Lire entier %s", s.Name)
	case ast.IncDecStmt:
		return fmt.Sprintf("%s%s", s.Name, string(s.Op))
	default:
		return ""
	}
}

func (g *Generator) exprStr(e ast.Expr) string {
	switch x := e.(type) {
	case ast.IdentExpr:
		return x.Name
	case ast.LiteralExpr:
		return x.Text
	case ast.BinaryExpr:
		lhs := g.exprStr(x.LHS)
		if ast.NeedsParens(x.Op, x.LHS, false) {
			lhs = "(" + lhs + ")"
		}
		rhs := g.exprStr(x.RHS)
		if ast.NeedsParens(x.Op, x.RHS, true) {
			rhs = "(" + rhs + ")"
		}
		return fmt.Sprintf("%s %s %s", lhs, string(x.Op), rhs)
	case ast.UnaryExpr:
		return fmt.Sprintf("%s%s", string(x.Op), g.exprStr(x.X))
	case ast.ListExpr:
		elems := make([]string, len(x.Elems))
		for i, el := range x.Elems {
			elems[i] = g.exprStr(el)
		}
		return "[" + strings.Join(elems, ", ") + "]"
	default:
		return "?"
	}
}
