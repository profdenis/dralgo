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
	var buf bytes.Buffer
	buf.WriteString("@startuml\n")
	buf.WriteString("start\n")

	g.genNode(&buf, flow.Start)

	buf.WriteString("stop\n")
	buf.WriteString("@enduml\n")
	return buf.String(), nil
}

func (g *Generator) genNode(buf *bytes.Buffer, n *ir.Node) {
	if n == nil || g.visited[n.ID] {
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
			g.genNode(buf, n.Next.Target)
		}

	case ir.NodeMerge:
		if n.Next != nil {
			g.genNode(buf, n.Next.Target)
		}

	case ir.NodeDecision:
		// Check if it's a loop
		isLoop := g.isReachable(n.Then.Target, n)

		if isLoop {
			// Check if it's a 'tant que' (While) or 'fait ... tant que' (DoWhile)
			// For DoWhile, the entry is bodyStart, and the decision is after bodyEnd.
			// But our lowering might have slightly different structures.
			// In our lowering for DoWhile: current -> bodyStart -> bodyEnd -> decision -> (Then -> bodyStart, Else -> merge)
			// In our lowering for While: current -> decision -> (Then -> bodyStart -> bodyEnd -> decision, Else -> merge)

			// If we reached this decision node and it's a loop, let's see where we came from.
			// This is tricky without more context.
			// Let's use a simpler approach for PlantUML: just use if/else and explicit labels if needed?
			// No, PlantUML structured is better.

			// Let's look at the structure:
			// While loop: entry points to Decision.
			// DoWhile loop: entry points to Body, which then points to Decision.

			// If g.visited[n.ID] is false, we are entering the loop.
			buf.WriteString(fmt.Sprintf("while (%s) is (vrai)\n", g.exprStr(n.Cond)))
			g.visited[n.ID] = true
			g.genNode(buf, n.Then.Target)
			g.visited[n.ID] = false
			buf.WriteString("endwhile (faux)\n")

			if n.Else != nil {
				g.genNode(buf, n.Else.Target)
			}
		} else {
			buf.WriteString(fmt.Sprintf("if (%s) then (vrai)\n", g.exprStr(n.Cond)))
			if n.Then != nil {
				g.genNode(buf, n.Then.Target)
			}

			if n.Else != nil {
				// Elif optimization
				if n.Else.Target.Type == ir.NodeDecision && !g.isReachable(n.Else.Target.Then.Target, n.Else.Target) {
					buf.WriteString(fmt.Sprintf("(faux) elseif (%s) then (vrai)\n", g.exprStr(n.Else.Target.Cond)))
					g.visited[n.Else.Target.ID] = true
					if n.Else.Target.Then != nil {
						g.genNode(buf, n.Else.Target.Then.Target)
					}
					if n.Else.Target.Else != nil {
						g.visited[n.Else.Target.ID] = false
						g.genElse(buf, n.Else.Target.Else.Target)
					}
				} else if n.Else.Target.Type == ir.NodeMerge && n.Else.Target.Next == nil {
					buf.WriteString("endif\n")
				} else {
					buf.WriteString("else (faux)\n")
					g.genNode(buf, n.Else.Target)
					buf.WriteString("endif\n")
				}
			} else {
				buf.WriteString("endif\n")
			}
		}
	}
}

func (g *Generator) genElse(buf *bytes.Buffer, n *ir.Node) {
	if n == nil || g.visited[n.ID] {
		buf.WriteString("endif\n")
		return
	}

	if n.Type == ir.NodeDecision && !g.isReachable(n.Then.Target, n) {
		buf.WriteString(fmt.Sprintf("(faux) elseif (%s) then (vrai)\n", g.exprStr(n.Cond)))
		g.visited[n.ID] = true
		if n.Then != nil {
			g.genNode(buf, n.Then.Target)
		}
		if n.Else != nil {
			g.visited[n.ID] = false
			g.genElse(buf, n.Else.Target)
		} else {
			buf.WriteString("endif\n")
		}
	} else if n.Type == ir.NodeMerge && n.Next == nil {
		buf.WriteString("endif\n")
	} else {
		buf.WriteString("else (faux)\n")
		g.genNode(buf, n)
		buf.WriteString("endif\n")
	}
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
		return fmt.Sprintf("%s %s %s", g.exprStr(x.LHS), string(x.Op), g.exprStr(x.RHS))
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
