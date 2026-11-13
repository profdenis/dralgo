package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"dralgo/internal/gen"
	"dralgo/internal/ir"
	"dralgo/internal/parser"
)

func main() {
	lang := flag.String("lang", "python", fmt.Sprintf("Target language (%s)", strings.Join(gen.Languages(), ", ")))
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

	g, ok := gen.Lookup(*lang)
	if !ok {
		fmt.Printf("Unsupported language: %s\n", *lang)
		os.Exit(1)
	}

	generated, genErr := g.Generate(flow)
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
