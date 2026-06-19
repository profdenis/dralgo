package js

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
	buf.WriteString("(async () => {\n")

	g.genNode(&buf, flow.Start, "    ", nil)

	buf.WriteString("})();\n")
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
			buf.WriteString(fmt.Sprintf("%swhile (%s) {\n", indent, g.exprStr(n.Cond)))
			g.visited[n.ID] = true
			g.genNode(buf, n.Then.Target, indent+"    ", n)
			g.visited[n.ID] = false
			buf.WriteString(indent + "}\n")

			if mergeNode != nil {
				g.genNode(buf, mergeNode, indent, stopAt)
			}
		} else {
			mergeNode := g.findMergeNode(n)
			buf.WriteString(fmt.Sprintf("%sif (%s) {\n", indent, g.exprStr(n.Cond)))
			if n.Then != nil {
				g.genNode(buf, n.Then.Target, indent+"    ", mergeNode)
			}
			buf.WriteString(indent + "}")

			if n.Else != nil {
				if n.Else.Target.Type == ir.NodeDecision && !g.isReachable(n.Else.Target.Then.Target, n.Else.Target) {
					buf.WriteString(" else if (" + g.exprStr(n.Else.Target.Cond) + ") {\n")
					g.visited[n.Else.Target.ID] = true
					if n.Else.Target.Then != nil {
						g.genNode(buf, n.Else.Target.Then.Target, indent+"    ", mergeNode)
					}
					buf.WriteString(indent + "}")
					if n.Else.Target.Else != nil {
						g.visited[n.Else.Target.ID] = false
						g.genElse(buf, n.Else.Target.Else.Target, indent, mergeNode)
					} else {
						buf.WriteString("\n")
					}
				} else if n.Else.Target == mergeNode {
					buf.WriteString("\n")
				} else {
					buf.WriteString(" else {\n")
					g.genNode(buf, n.Else.Target, indent+"    ", mergeNode)
					buf.WriteString(indent + "}\n")
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
		return
	}

	if n.Type == ir.NodeDecision && !g.isReachable(n.Then.Target, n) {
		buf.WriteString(" else if (" + g.exprStr(n.Cond) + ") {\n")
		g.visited[n.ID] = true
		if n.Then != nil {
			g.genNode(buf, n.Then.Target, indent+"    ", stopAt)
		}
		buf.WriteString(indent + "}")
		if n.Else != nil {
			g.visited[n.ID] = false
			g.genElse(buf, n.Else.Target, indent, stopAt)
		} else {
			buf.WriteString("\n")
		}
	} else if n == stopAt {
		buf.WriteString("\n")
	} else {
		buf.WriteString(" else {\n")
		g.genNode(buf, n, indent+"    ", stopAt)
		buf.WriteString(indent + "}\n")
	}
}

func (g *Generator) findMergeNode(n *ir.Node) *ir.Node {
	if n.Then == nil {
		return nil
	}
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
		kw := "let "
		if g.declaredVars[s.Name] {
			kw = ""
		}
		buf.WriteString(fmt.Sprintf("%s%s%s = %s;\n", indent, kw, s.Name, g.exprStr(s.Value)))
		g.declaredVars[s.Name] = true
	case ast.PrintStmt:
		args := make([]string, len(s.Args))
		for i, a := range s.Args {
			args[i] = g.exprStr(a)
		}
		buf.WriteString(fmt.Sprintf("%sconsole.log(%s);\n", indent, strings.Join(args, ", ")))
	case ast.ReadStmt:
		kw := "let "
		if g.declaredVars[s.Name] {
			kw = ""
		}
		buf.WriteString(fmt.Sprintf("%s%s%s = await new Promise(resolve => process.stdin.once('data', d => resolve(d.toString().trim())));\n", indent, kw, s.Name))
		g.declaredVars[s.Name] = true
	case ast.ReadIntStmt:
		kw := "let "
		if g.declaredVars[s.Name] {
			kw = ""
		}
		buf.WriteString(fmt.Sprintf("%s%s%s = parseInt(await new Promise(resolve => process.stdin.once('data', d => resolve(d.toString().trim()))));\n", indent, kw, s.Name))
		g.declaredVars[s.Name] = true
	case ast.IncDecStmt:
		buf.WriteString(fmt.Sprintf("%s%s%s;\n", indent, s.Name, string(s.Op)))
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
		if op == "**" {
			return fmt.Sprintf("Math.pow(%s, %s)", g.exprStr(x.LHS), g.exprStr(x.RHS))
		}
		return fmt.Sprintf("(%s %s %s)", g.exprStr(x.LHS), op, g.exprStr(x.RHS))
	case ast.UnaryExpr:
		return fmt.Sprintf("%s%s", string(x.Op), g.exprStr(x.X))
	case ast.ListExpr:
		elems := make([]string, len(x.Elems))
		for i, el := range x.Elems {
			elems[i] = g.exprStr(el)
		}
		return "[" + strings.Join(elems, ", ") + "]"
	default:
		return "undefined"
	}
}
