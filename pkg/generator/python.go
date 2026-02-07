package generator

import (
	"bytes"
	"dralgo/pkg/models"
	"fmt"
	"strings"
)

type PythonGenerator struct{}

func NewPythonGenerator() *PythonGenerator {
	return &PythonGenerator{}
}

func (g *PythonGenerator) Generate(flowchart *models.Flowchart) (string, error) {
	var buf bytes.Buffer

	buf.WriteString("def main():\n")

	g.generateNode(&buf, flowchart.Start, "    ")

	buf.WriteString("\nif __name__ == \"__main__\":\n")
	buf.WriteString("    main()\n")

	return buf.String(), nil
}

func (g *PythonGenerator) generateNode(buf *bytes.Buffer, node *models.Node, indent string) {
	if node == nil {
		return
	}

	switch node.Type {
	case models.NodeStart:
		buf.WriteString(fmt.Sprintf("%sprint(\"Starting flowchart...\")\n", indent))
	case models.NodeAction:
		buf.WriteString(fmt.Sprintf("%sprint(\"%s\")\n", indent, node.Content))
	case models.NodeAssignment:
		parts := strings.Split(node.Content, "<--")
		if len(parts) == 2 {
			variable := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			buf.WriteString(fmt.Sprintf("%s%s = %s\n", indent, variable, value))
		}
	case models.NodeDisplay:
		content := strings.TrimPrefix(node.Content, "Afficher")
		content = strings.TrimPrefix(content, " ")
		args := parseDisplayArgs(content)
		buf.WriteString(fmt.Sprintf("%sprint(%s)\n", indent, strings.Join(args, ", ")))
	case models.NodeRead:
		variable := strings.TrimSpace(strings.TrimPrefix(node.Content, "Lire "))
		buf.WriteString(fmt.Sprintf("%s%s = input()\n", indent, variable))
	case models.NodeReadInt:
		variable := strings.TrimSpace(strings.TrimPrefix(node.Content, "LireEntier "))
		buf.WriteString(fmt.Sprintf("\n%s# %s\n", indent, node.Content))
		buf.WriteString(fmt.Sprintf("%stry:\n", indent))
		buf.WriteString(fmt.Sprintf("%s    %s = int(input())\n", indent, variable))
		buf.WriteString(fmt.Sprintf("%sexcept ValueError:\n", indent))
		buf.WriteString(fmt.Sprintf("%s    print(\"La valeur entrée n'est pas un entier\")\n", indent))
		buf.WriteString(fmt.Sprintf("%s    import sys; sys.exit(1)\n", indent))
		buf.WriteString("\n")
	case models.NodeDecision:
		buf.WriteString(fmt.Sprintf("%sif %s:\n", indent, node.Content))
		if len(node.Next) > 0 {
			for _, edge := range node.Next {
				g.generateNode(buf, edge.Target, indent+"    ")
			}
		} else {
			buf.WriteString(fmt.Sprintf("%s    pass\n", indent))
		}
		if node.Else != nil {
			buf.WriteString(fmt.Sprintf("%selse:\n", indent))
			g.generateNode(buf, node.Else, indent+"    ")
		}
		return // Les enfants sont déjà gérés
	case models.NodeEnd:
		buf.WriteString(fmt.Sprintf("%sprint(\"Ending flowchart...\")\n", indent))
	}

	for _, edge := range node.Next {
		g.generateNode(buf, edge.Target, indent)
	}
}
