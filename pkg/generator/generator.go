package generator

import (
	"dralgo/pkg/models"
	"strings"
)

type Generator interface {
	Generate(flowchart *models.Flowchart) (string, error)
}

func parseDisplayArgs(content string) []string {
	var args []string
	var current strings.Builder
	inQuotes := false

	for _, r := range content {
		switch r {
		case '"':
			inQuotes = !inQuotes
			current.WriteRune(r)
		case ' ':
			if inQuotes {
				current.WriteRune(r)
			} else {
				if current.Len() > 0 {
					args = append(args, current.String())
					current.Reset()
				}
			}
		default:
			current.WriteRune(r)
		}
	}
	if current.Len() > 0 {
		args = append(args, current.String())
	}
	return args
}
