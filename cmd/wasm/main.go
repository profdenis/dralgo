package main

import (
	"fmt"
	"syscall/js"

	"dralgo/internal/gen"
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
		return failure("argument(s) manquant(s) (input, lang)")
	}

	input := args[0].String()
	lang := args[1].String()

	prog, errs := parser.ParseProgram(input)
	if len(errs) > 0 {
		var errStr string
		for _, e := range errs {
			errStr += fmt.Sprintf("  %v\n", e)
		}
		return failure(fmt.Sprintf("Erreurs de parsing :\n%s", errStr))
	}

	flow := ir.Lower(prog)

	g, ok := gen.Lookup(lang)
	if !ok {
		return failure(fmt.Sprintf("Langage non supporté : %s", lang))
	}

	generated, genErr := g.Generate(flow)
	if genErr != nil {
		return failure(fmt.Sprintf("Erreur lors de la génération : %v", genErr))
	}

	return success(generated)
}

func success(code string) map[string]any {
	return map[string]any{"ok": true, "code": code, "error": ""}
}

func failure(message string) map[string]any {
	return map[string]any{"ok": false, "code": "", "error": message}
}
