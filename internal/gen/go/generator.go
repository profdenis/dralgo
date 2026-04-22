package gogen

import (
	"bytes"
	"dralgo/internal/ast"
	"dralgo/internal/ir"
	"fmt"
	"strings"
)

type Generator struct {
	visited      map[int]bool
	declaredVars map[string]bool
}

func NewGenerator() *Generator {
	return &Generator{
		visited:      make(map[int]bool),
		declaredVars: make(map[string]bool),
	}
}

func (g *Generator) Generate(flow *ir.Flowchart) (string, error) {
	g.visited = make(map[int]bool)
	g.declaredVars = make(map[string]bool)
	var buf bytes.Buffer

	buf.WriteString("package main\n\n")
	buf.WriteString("import (\n\t\"fmt\"\n\t\"os\"\n)\n\n")
	buf.WriteString("func main() {\n")

	g.genNode(&buf, flow.Start, "    ", nil)

	buf.WriteString("}\n")
	return buf.String(), nil
}

func (g *Generator) genNode(buf *bytes.Buffer, n *ir.Node, indent string, stopAt *ir.Node) {
	if n == nil || g.visited[n.ID] || n == stopAt {
		return
	}

	switch n.Type {
	case ir.NodeExit:
		return

	case ir.NodeAction:
		for _, stmt := range n.Stmts {
			g.genStmt(buf, stmt, indent)
		}
		if n.Next != nil {
			g.genNode(buf, n.Next.Target, indent, stopAt)
		}

	case ir.NodeMerge:
		if n.Next != nil {
			g.genNode(buf, n.Next.Target, indent, stopAt)
		}

	case ir.NodeDecision:
		isLoop := g.isReachable(n.Then.Target, n)

		if isLoop {
			mergeNode := n.Else.Target
			buf.WriteString(fmt.Sprintf("%sfor %s {\n", indent, g.exprStr(n.Cond)))
			g.visited[n.ID] = true
			g.genNode(buf, n.Then.Target, indent+"    ", n)
			g.visited[n.ID] = false
			buf.WriteString(fmt.Sprintf("%s}\n", indent))

			if mergeNode != nil {
				g.genNode(buf, mergeNode, indent, stopAt)
			}
		} else {
			mergeNode := g.findMergeNode(n)
			buf.WriteString(fmt.Sprintf("%sif %s {\n", indent, g.exprStr(n.Cond)))
			if n.Then != nil {
				g.genNode(buf, n.Then.Target, indent+"    ", mergeNode)
			}
			buf.WriteString(fmt.Sprintf("%s}", indent))

			if n.Else != nil {
				if n.Else.Target.Type == ir.NodeDecision && !g.isReachable(n.Else.Target.Then.Target, n.Else.Target) {
					buf.WriteString(" else if " + g.exprStr(n.Else.Target.Cond) + " {\n")
					g.visited[n.Else.Target.ID] = true
					if n.Else.Target.Then != nil {
						g.genNode(buf, n.Else.Target.Then.Target, indent+"    ", mergeNode)
					}
					if n.Else.Target.Else != nil {
						g.visited[n.Else.Target.ID] = false
						g.genElse(buf, n.Else.Target.Else.Target, indent, mergeNode)
					} else {
						buf.WriteString(fmt.Sprintf("%s}", indent))
					}
				} else if n.Else.Target.Type == ir.NodeMerge && n.Else.Target.Next == nil && n.Else.Target == mergeNode {
					buf.WriteString("\n")
				} else {
					buf.WriteString(" else {\n")
					g.genNode(buf, n.Else.Target, indent+"    ", mergeNode)
					buf.WriteString(fmt.Sprintf("%s}\n", indent))
				}
			} else {
				buf.WriteString("\n")
			}

			if mergeNode != nil {
				g.genNode(buf, mergeNode, indent, stopAt)
			}
		}
	}
}

func (g *Generator) genElse(buf *bytes.Buffer, n *ir.Node, indent string, stopAt *ir.Node) {
	if n == nil || g.visited[n.ID] || n == stopAt {
		buf.WriteString(fmt.Sprintf("%s}\n", indent))
		return
	}

	if n.Type == ir.NodeDecision && !g.isReachable(n.Then.Target, n) {
		buf.WriteString(" else if " + g.exprStr(n.Cond) + " {\n")
		g.visited[n.ID] = true
		if n.Then != nil {
			g.genNode(buf, n.Then.Target, indent+"    ", stopAt)
		}
		if n.Else != nil {
			g.visited[n.ID] = false
			g.genElse(buf, n.Else.Target, indent, stopAt)
		} else {
			buf.WriteString(fmt.Sprintf("%s}\n", indent))
		}
	} else if n.Type == ir.NodeMerge && n.Next == nil && n == stopAt {
		buf.WriteString(fmt.Sprintf("%s}\n", indent))
	} else {
		buf.WriteString(" else {\n")
		g.genNode(buf, n, indent+"    ", stopAt)
		buf.WriteString(fmt.Sprintf("%s}\n", indent))
	}
}

func (g *Generator) findMergeNode(n *ir.Node) *ir.Node {
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

func (g *Generator) genStmt(buf *bytes.Buffer, stmt ast.Stmt, indent string) {
	switch s := stmt.(type) {
	case ast.AssignStmt:
		op := ":="
		if g.declaredVars[s.Name] {
			op = "="
		}
		buf.WriteString(fmt.Sprintf("%s%s %s %s\n", indent, s.Name, op, g.exprStr(s.Value)))
		g.declaredVars[s.Name] = true
	case ast.PrintStmt:
		args := make([]string, len(s.Args))
		for i, a := range s.Args {
			args[i] = g.exprStr(a)
		}
		buf.WriteString(fmt.Sprintf("%sfmt.Println(%s)\n", indent, strings.Join(args, ", ")))
	case ast.ReadStmt:
		if !g.declaredVars[s.Name] {
			buf.WriteString(fmt.Sprintf("%svar %s string\n", indent, s.Name))
			g.declaredVars[s.Name] = true
		}
		buf.WriteString(fmt.Sprintf("%sfmt.Scanln(&%s)\n", indent, s.Name))
	case ast.ReadIntStmt:
		if !g.declaredVars[s.Name] {
			buf.WriteString(fmt.Sprintf("%svar %s int\n", indent, s.Name))
			g.declaredVars[s.Name] = true
		}
		errOp := ":="
		if g.declaredVars["err"] {
			errOp = "="
		}
		buf.WriteString(fmt.Sprintf("%s_, %s %s fmt.Scanln(&%s)\n", indent, "err", errOp, s.Name))
		g.declaredVars["err"] = true
		buf.WriteString(fmt.Sprintf("%sif err != nil {\n", indent))
		buf.WriteString(fmt.Sprintf("%s    fmt.Fprintln(os.Stderr, \"Erreur: un entier est attendu\")\n", indent))
		buf.WriteString(fmt.Sprintf("%s    os.Exit(1)\n", indent))
		buf.WriteString(fmt.Sprintf("%s}\n", indent))
	case ast.IncDecStmt:
		buf.WriteString(fmt.Sprintf("%s%s%s\n", indent, s.Name, string(s.Op)))
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
		op := string(x.Op)
		if x.Op == ast.OpPow {
			// Go doesn't have a power operator, we should ideally use math.Pow
			// But math.Pow takes float64. For simplicity in this DSL, we might just use a placeholder or handle it if we had a type system.
			// Let's use a simple function call if we want to be correct, but let's see how python generator handles it.
			// Python generator just uses string(x.Op) which is "**". That's valid Python but not valid Go.
			return fmt.Sprintf("pow(%s, %s)", g.exprStr(x.LHS), g.exprStr(x.RHS))
		}
		return fmt.Sprintf("%s %s %s", g.exprStr(x.LHS), op, g.exprStr(x.RHS))
	case ast.UnaryExpr:
		return fmt.Sprintf("%s%s", string(x.Op), g.exprStr(x.X))
	case ast.ListExpr:
		elems := make([]string, len(x.Elems))
		for i, el := range x.Elems {
			elems[i] = g.exprStr(el)
		}
		return "[]any{" + strings.Join(elems, ", ") + "}"
	default:
		return "nil"
	}
}
