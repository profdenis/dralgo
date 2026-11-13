package gen

import (
	"dralgo/internal/ir"

	gogen "dralgo/internal/gen/go"
	"dralgo/internal/gen/java"
	jsgen "dralgo/internal/gen/js"
	"dralgo/internal/gen/plantuml"
	"dralgo/internal/gen/python"
)

type Generator interface {
	Generate(flow *ir.Flowchart) (string, error)
}

type entry struct {
	name    string
	aliases []string
	newGen  func() Generator
}

var entries = []entry{
	{name: "puml", aliases: []string{"plantuml"}, newGen: func() Generator { return plantuml.NewGenerator() }},
	{name: "python", newGen: func() Generator { return python.NewGenerator() }},
	{name: "go", newGen: func() Generator { return gogen.NewGenerator() }},
	{name: "js", aliases: []string{"javascript"}, newGen: func() Generator { return jsgen.NewGenerator() }},
	{name: "java", newGen: func() Generator { return java.NewGenerator() }},
}

func Lookup(lang string) (Generator, bool) {
	for _, e := range entries {
		if e.name == lang {
			return e.newGen(), true
		}
		for _, alias := range e.aliases {
			if alias == lang {
				return e.newGen(), true
			}
		}
	}
	return nil, false
}

func Languages() []string {
	names := make([]string, 0, len(entries))
	for _, e := range entries {
		names = append(names, e.name)
	}
	return names
}
