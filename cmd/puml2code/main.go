package main

import (
	"dralgo/pkg/generator"
	"dralgo/pkg/plantuml"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	lang := flag.String("lang", "python", "Target language (python or go)")
	input := flag.String("input", "", "Input file (stdin if omitted)")
	output := flag.String("output", "", "Output file (stdout if omitted)")
	flag.Parse()

	var pumlContent []byte
	var err error

	if *input != "" {
		pumlContent, err = os.ReadFile(*input)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading input file: %v\n", err)
			os.Exit(1)
		}
	} else {
		pumlContent, err = io.ReadAll(os.Stdin)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading from stdin: %v\n", err)
			os.Exit(1)
		}
	}

	parser := plantuml.NewParser()
	flowchart, err := parser.Parse(string(pumlContent))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing PlantUML: %v\n", err)
		os.Exit(1)
	}

	var gen generator.Generator
	switch *lang {
	case "python":
		gen = generator.NewPythonGenerator()
	case "go":
		gen = generator.NewGoGenerator()
	default:
		fmt.Fprintf(os.Stderr, "Unsupported language: %s\n", *lang)
		os.Exit(1)
	}

	code, err := gen.Generate(flowchart)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error generating code: %v\n", err)
		os.Exit(1)
	}

	if *output != "" {
		err = os.WriteFile(*output, []byte(code), 0644)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error writing output file: %v\n", err)
			os.Exit(1)
		}
	} else {
		fmt.Println(code)
	}
}
