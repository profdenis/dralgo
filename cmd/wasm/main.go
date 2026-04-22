package main

import (
	"fmt"
	"syscall/js"

	gogen "dralgo/internal/gen/go"
	"dralgo/internal/gen/java"
	jsgen "dralgo/internal/gen/js"
	"dralgo/internal/gen/plantuml"
	"dralgo/internal/gen/python"
	"dralgo/internal/ir"
	"dralgo/internal/parser"
)

func main() {
	fmt.Println("DrAlgo WASM Loaded")

	js.Global().Set("dralgoTranslate", js.FuncOf(translate))

	// Keep the Go program running
	select {}
}

func translate(this js.Value, args []js.Value) any {
	if len(args) < 2 {
		return "Erreur : argument(s) manquant(s) (input, lang)"
	}

	input := args[0].String()
	lang := args[1].String()

	prog, errs := parser.ParseProgram(input)
	if len(errs) > 0 {
		var errStr string
		for _, e := range errs {
			errStr += fmt.Sprintf("  %v\n", e)
		}
		return fmt.Sprintf("Erreurs de parsing :\n%s", errStr)
	}

	flow := ir.Lower(prog)

	var generated string
	var genErr error

	switch lang {
	case "python":
		g := python.NewGenerator()
		generated, genErr = g.Generate(flow)
	case "go":
		g := gogen.NewGenerator()
		generated, genErr = g.Generate(flow)
	case "js", "javascript":
		g := jsgen.NewGenerator()
		generated, genErr = g.Generate(flow)
	case "java":
		g := java.NewGenerator()
		generated, genErr = g.Generate(flow)
	case "puml", "plantuml":
		g := plantuml.NewGenerator()
		generated, genErr = g.Generate(flow)
	default:
		return fmt.Sprintf("Langage non supporté : %s", lang)
	}

	if genErr != nil {
		return fmt.Sprintf("Erreur lors de la génération : %v", genErr)
	}

	return generated
}
