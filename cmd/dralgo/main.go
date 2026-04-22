package main

import (
	"flag"
	"fmt"
	"os"

	gogen "dralgo/internal/gen/go"
	"dralgo/internal/gen/java"
	jsgen "dralgo/internal/gen/js"
	"dralgo/internal/gen/plantuml"
	"dralgo/internal/gen/python"
	"dralgo/internal/ir"
	"dralgo/internal/parser"
)

func main() {
	lang := flag.String("lang", "python", "Target language (python, puml)")
	output := flag.String("o", "", "Output file")
	flag.Parse()

	if flag.NArg() < 1 {
		fmt.Println("Usage: dralgo [options] <input_file>")
		os.Exit(1)
	}

	inputPath := flag.Arg(0)
	input, err := os.ReadFile(inputPath)
	if err != nil {
		fmt.Printf("Error reading input file: %v\n", err)
		os.Exit(1)
	}

	prog, errs := parser.ParseProgram(string(input))
	if len(errs) > 0 {
		fmt.Printf("Errors found during parsing:\n")
		for _, e := range errs {
			fmt.Printf("  %v\n", e)
		}
		os.Exit(1)
	}

	flow := ir.Lower(prog)

	var generated string
	var genErr error

	switch *lang {
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
		fmt.Printf("Unsupported language: %s\n", *lang)
		os.Exit(1)
	}

	if genErr != nil {
		fmt.Printf("Error during code generation: %v\n", genErr)
		os.Exit(1)
	}

	if *output != "" {
		err = os.WriteFile(*output, []byte(generated), 0644)
		if err != nil {
			fmt.Printf("Error writing output file: %v\n", err)
			os.Exit(1)
		}
	} else {
		fmt.Print(generated)
	}
}
