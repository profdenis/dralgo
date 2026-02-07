package plantuml

import (
	"bufio"
	"dralgo/pkg/models"
	"fmt"
	"strings"
)

type Parser struct{}

func NewParser() *Parser {
	return &Parser{}
}

func (p *Parser) Parse(input string) (*models.Flowchart, error) {
	scanner := bufio.NewScanner(strings.NewReader(input))
	flowchart := &models.Flowchart{
		Nodes: []*models.Node{},
	}

	var currentNode *models.Node
	var decisionStack []*models.Node
	var inElse bool

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "@start") || strings.HasPrefix(line, "@end") {
			continue
		}

		if line == "start" {
			node := &models.Node{ID: "start", Type: models.NodeStart, Content: "start"}
			flowchart.Nodes = append(flowchart.Nodes, node)
			flowchart.Start = node
			currentNode = node
			continue
		}

		if line == "stop" || line == "end" {
			node := &models.Node{ID: "end", Type: models.NodeEnd, Content: "end"}
			flowchart.Nodes = append(flowchart.Nodes, node)
			if currentNode != nil {
				if inElse && len(decisionStack) > 0 {
					decisionStack[len(decisionStack)-1].Else = node
				} else {
					currentNode.Next = append(currentNode.Next, &models.Edge{Target: node})
				}
			}
			currentNode = nil
			continue
		}

		if strings.HasPrefix(line, ":") && strings.HasSuffix(line, ";") {
			content := strings.Trim(line, ":;")
			nodeType := models.NodeAction
			if strings.Contains(content, "<--") {
				nodeType = models.NodeAssignment
			} else if content == "Afficher" || strings.HasPrefix(content, "Afficher ") {
				nodeType = models.NodeDisplay
			} else if strings.HasPrefix(content, "LireEntier ") {
				nodeType = models.NodeReadInt
			} else if strings.HasPrefix(content, "Lire ") {
				nodeType = models.NodeRead
			}
			node := &models.Node{
				ID:      fmt.Sprintf("node_%d", len(flowchart.Nodes)),
				Type:    nodeType,
				Content: content,
			}
			flowchart.Nodes = append(flowchart.Nodes, node)
			if currentNode != nil {
				if inElse && len(decisionStack) > 0 && decisionStack[len(decisionStack)-1].Else == nil {
					decisionStack[len(decisionStack)-1].Else = node
				} else {
					currentNode.Next = append(currentNode.Next, &models.Edge{Target: node})
				}
			}
			currentNode = node
			continue
		}

		if strings.HasPrefix(line, "if (") && strings.HasSuffix(line, ") then") {
			condition := line[strings.Index(line, "(")+1 : strings.LastIndex(line, ")")]
			condition = strings.TrimSuffix(condition, "?")
			node := &models.Node{
				ID:      fmt.Sprintf("node_%d", len(flowchart.Nodes)),
				Type:    models.NodeDecision,
				Content: condition,
			}
			flowchart.Nodes = append(flowchart.Nodes, node)
			if currentNode != nil {
				currentNode.Next = append(currentNode.Next, &models.Edge{Target: node})
			}
			decisionStack = append(decisionStack, node)
			currentNode = node
			inElse = false
			continue
		}

		if line == "else" {
			inElse = true
			if len(decisionStack) > 0 {
				currentNode = decisionStack[len(decisionStack)-1]
			}
			continue
		}

		if line == "endif" {
			inElse = false
			if len(decisionStack) > 0 {
				decisionStack = decisionStack[:len(decisionStack)-1]
			}
			continue
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return flowchart, nil
}
