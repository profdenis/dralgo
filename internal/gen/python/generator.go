package python

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
	buf.WriteString("def main():\n")

	// Collect all nodes to handle loops and jumps via labels or while True
	// For this version, let's try to keep it simple and use recursive generation
	// for structured blocks, but we need to be careful with NodeMerge.

	g.genNode(&buf, flow.Start, "    ")

	buf.WriteString("\nif __name__ == \"__main__\":\n")
	buf.WriteString("    main()\n")
	return buf.String(), nil
}

func (g *Generator) genNode(buf *bytes.Buffer, n *ir.Node, indent string) {
	if n == nil || g.visited[n.ID] {
		return
	}

	// We mark as visited ONLY if it's a decision or merge that we want to avoid
	// re-processing in a way that causes infinite recursion.
	// But actually, structured lowering should mostly avoid cycles except for loop back-edges.

	switch n.Type {
	case ir.NodeExit:
		return

	case ir.NodeAction:
		for _, stmt := range n.Stmts {
			g.genStmt(buf, stmt, indent)
		}
		if n.Next != nil {
			g.genNode(buf, n.Next.Target, indent)
		}

	case ir.NodeMerge:
		// Just pass through
		if n.Next != nil {
			g.genNode(buf, n.Next.Target, indent)
		}

	case ir.NodeDecision:
		// Check if it's a loop (While/For/DoWhile) or an If
		// Heuristic: if it's a loop, one of the branches (or its descendants) leads back to this node.
		isLoop := g.isReachable(n.Then.Target, n)

		if isLoop {
			buf.WriteString(fmt.Sprintf("%swhile %s:\n", indent, g.exprStr(n.Cond)))
			g.visited[n.ID] = true // Mark to avoid recursion when coming back from body
			g.genNode(buf, n.Then.Target, indent+"    ")
			g.visited[n.ID] = false

			// After loop body, continue with Else branch (exit loop)
			if n.Else != nil {
				g.genNode(buf, n.Else.Target, indent)
			}
		} else {
			// It's an If
			buf.WriteString(fmt.Sprintf("%sif %s:\n", indent, g.exprStr(n.Cond)))
			if n.Then != nil {
				g.genNode(buf, n.Then.Target, indent+"    ")
			} else {
				buf.WriteString(indent + "    pass\n")
			}

			if n.Else != nil {
				// Special case for 'sinon si' (nested if) to avoid over-indentation
				if n.Else.Target.Type == ir.NodeDecision && !g.isReachable(n.Else.Target.Then.Target, n.Else.Target) {
					buf.WriteString(indent + "elif " + g.exprStr(n.Else.Target.Cond) + ":\n")
					g.visited[n.Else.Target.ID] = true
					if n.Else.Target.Then != nil {
						g.genNode(buf, n.Else.Target.Then.Target, indent+"    ")
					} else {
						buf.WriteString(indent + "    pass\n")
					}

					if n.Else.Target.Else != nil {
						// Recurse for further elifs or final else
						g.visited[n.Else.Target.ID] = false
						g.genElse(buf, n.Else.Target.Else.Target, indent)
					}
				} else if n.Else.Target.Type == ir.NodeMerge && n.Else.Target.Next == nil {
					// skip empty else
				} else {
					buf.WriteString(indent + "else:\n")
					g.genNode(buf, n.Else.Target, indent+"    ")
				}
			}
		}
	}
}

func (g *Generator) genElse(buf *bytes.Buffer, n *ir.Node, indent string) {
	if n == nil || g.visited[n.ID] {
		return
	}

	if n.Type == ir.NodeDecision && !g.isReachable(n.Then.Target, n) {
		buf.WriteString(indent + "elif " + g.exprStr(n.Cond) + ":\n")
		g.visited[n.ID] = true
		if n.Then != nil {
			g.genNode(buf, n.Then.Target, indent+"    ")
		} else {
			buf.WriteString(indent + "    pass\n")
		}
		if n.Else != nil {
			g.visited[n.ID] = false
			g.genElse(buf, n.Else.Target, indent)
		}
	} else if n.Type == ir.NodeMerge && n.Next == nil {
		// skip empty
	} else {
		buf.WriteString(indent + "else:\n")
		g.genNode(buf, n, indent+"    ")
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

func (g *Generator) genStmt(buf *bytes.Buffer, stmt ast.Stmt, indent string) {
	switch s := stmt.(type) {
	case ast.AssignStmt:
		buf.WriteString(fmt.Sprintf("%s%s = %s\n", indent, s.Name, g.exprStr(s.Value)))
	case ast.PrintStmt:
		args := make([]string, len(s.Args))
		for i, a := range s.Args {
			args[i] = g.exprStr(a)
		}
		buf.WriteString(fmt.Sprintf("%sprint(%s)\n", indent, strings.Join(args, ", ")))
	case ast.ReadStmt:
		buf.WriteString(fmt.Sprintf("%s%s = input()\n", indent, s.Name))
	case ast.ReadIntStmt:
		buf.WriteString(fmt.Sprintf("%stry:\n", indent))
		buf.WriteString(fmt.Sprintf("%s    %s = int(input())\n", indent, s.Name))
		buf.WriteString(fmt.Sprintf("%sexcept ValueError:\n", indent))
		buf.WriteString(fmt.Sprintf("%s    print(\"Erreur: un entier est attendu\")\n", indent))
		buf.WriteString(fmt.Sprintf("%s    exit(1)\n", indent))
	case ast.IncDecStmt:
		op := "+="
		if s.Op == ast.Dec {
			op = "-="
		}
		buf.WriteString(fmt.Sprintf("%s%s %s 1\n", indent, s.Name, op))
	}
}

func (g *Generator) exprStr(e ast.Expr) string {
	switch x := e.(type) {
	case ast.IdentExpr:
		return x.Name
	case ast.LiteralExpr:
		if x.Kind == ast.LitString {
			return x.Text
		}
		return strings.ReplaceAll(x.Text, "_", "")
	case ast.BinaryExpr:
		return fmt.Sprintf("(%s %s %s)", g.exprStr(x.LHS), string(x.Op), g.exprStr(x.RHS))
	case ast.UnaryExpr:
		return fmt.Sprintf("(%s%s)", string(x.Op), g.exprStr(x.X))
	case ast.ListExpr:
		elems := make([]string, len(x.Elems))
		for i, el := range x.Elems {
			elems[i] = g.exprStr(el)
		}
		return "[" + strings.Join(elems, ", ") + "]"
	default:
		return "None"
	}
}
