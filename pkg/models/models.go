package models

type NodeType string

const (
	NodeStart      NodeType = "start"
	NodeEnd        NodeType = "end"
	NodeAction     NodeType = "action"
	NodeDecision   NodeType = "decision"
	NodeAssignment NodeType = "assignment"
	NodeDisplay    NodeType = "display"
	NodeRead       NodeType = "read"
	NodeReadInt    NodeType = "read_int"
)

type Node struct {
	ID      string
	Type    NodeType
	Content string
	Next    []*Edge
	Else    *Node // Nouveau champ pour le bras else
}

type Edge struct {
	Target *Node
	Label  string
}

type Flowchart struct {
	Nodes []*Node
	Start *Node
}
