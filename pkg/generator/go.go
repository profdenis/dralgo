package generator

import (
	"bytes"
	"dralgo/pkg/models"
	"fmt"
	"strings"
)

type GoGenerator struct {
	declaredVars map[string]bool
}

func NewGoGenerator() *GoGenerator {
	return &GoGenerator{
		declaredVars: make(map[string]bool),
	}
}

func (g *GoGenerator) Generate(flowchart *models.Flowchart) (string, error) {
	var buf bytes.Buffer
	g.declaredVars = make(map[string]bool)

	buf.WriteString("package main\n\n")
	buf.WriteString("import \"fmt\"\n\n")
	buf.WriteString("func main() {\n")

	g.generateNode(&buf, flowchart.Start, "    ")

	buf.WriteString("}\n")

	return buf.String(), nil
}

func (g *GoGenerator) generateNode(buf *bytes.Buffer, node *models.Node, indent string) {
	if node == nil {
		return
	}

	switch node.Type {
	case models.NodeStart:
		buf.WriteString(fmt.Sprintf("%sfmt.Println(\"Starting flowchart...\")\n", indent))
	case models.NodeAction:
		buf.WriteString(fmt.Sprintf("%sfmt.Println(\"%s\")\n", indent, node.Content))
	case models.NodeAssignment:
		parts := strings.Split(node.Content, "<--")
		if len(parts) == 2 {
			variable := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			op := ":="
			if g.declaredVars[variable] {
				op = "="
			}
			buf.WriteString(fmt.Sprintf("%s%s %s %s\n", indent, variable, op, value))
			g.declaredVars[variable] = true
		}
	case models.NodeDisplay:
		content := strings.TrimPrefix(node.Content, "Afficher")
		content = strings.TrimPrefix(content, " ")
		args := parseDisplayArgs(content)
		buf.WriteString(fmt.Sprintf("%sfmt.Println(%s)\n", indent, strings.Join(args, ", ")))
	case models.NodeRead:
		variable := strings.TrimSpace(strings.TrimPrefix(node.Content, "Lire "))
		buf.WriteString(fmt.Sprintf("\n%s// %s\n", indent, node.Content))
		op := ":="
		if g.declaredVars[variable] {
			op = "="
		}
		tempVar := "tempString"
		if !g.declaredVars[tempVar] {
			buf.WriteString(fmt.Sprintf("%svar %s string\n", indent, tempVar))
			g.declaredVars[tempVar] = true
		}
		buf.WriteString(fmt.Sprintf("%s_, _ = fmt.Scanln(&%s)\n", indent, tempVar))
		buf.WriteString(fmt.Sprintf("%s%s %s %s\n", indent, variable, op, tempVar))
		g.declaredVars[variable] = true
		buf.WriteString("\n")
	case models.NodeReadInt:
		variable := strings.TrimSpace(strings.TrimPrefix(node.Content, "LireEntier "))
		buf.WriteString(fmt.Sprintf("\n%s// %s\n", indent, node.Content))
		op := ":="
		if g.declaredVars[variable] {
			op = "="
		}
		tempVar := "tempInt"
		if !g.declaredVars[tempVar] {
			buf.WriteString(fmt.Sprintf("%svar %s int\n", indent, tempVar))
			g.declaredVars[tempVar] = true
		}
		errOp := ":="
		if g.declaredVars["err"] {
			errOp = "="
		}
		buf.WriteString(fmt.Sprintf("%s_, %s %s fmt.Scanln(&%s)\n", indent, "err", errOp, tempVar))
		g.declaredVars["err"] = true
		buf.WriteString(fmt.Sprintf("%sif err != nil {\n", indent))
		buf.WriteString(fmt.Sprintf("%s    fmt.Println(\"La valeur entrée n'est pas un entier\")\n", indent))
		buf.WriteString(fmt.Sprintf("%s    return\n", indent))
		buf.WriteString(fmt.Sprintf("%s}\n", indent))
		buf.WriteString(fmt.Sprintf("%s%s %s %s\n", indent, variable, op, tempVar))
		g.declaredVars[variable] = true
		buf.WriteString("\n")
	case models.NodeDecision:
		buf.WriteString(fmt.Sprintf("%sif %s {\n", indent, node.Content))
		if len(node.Next) > 0 {
			for _, edge := range node.Next {
				g.generateNode(buf, edge.Target, indent+"    ")
			}
		}
		buf.WriteString(fmt.Sprintf("%s}", indent))
		if node.Else != nil {
			buf.WriteString(" else {\n")
			g.generateNode(buf, node.Else, indent+"    ")
			buf.WriteString(fmt.Sprintf("%s}\n", indent))
		} else {
			buf.WriteString("\n")
		}
		return
	case models.NodeEnd:
		buf.WriteString(fmt.Sprintf("%sfmt.Println(\"Ending flowchart...\")\n", indent))
	}

	for _, edge := range node.Next {
		g.generateNode(buf, edge.Target, indent)
	}
}
