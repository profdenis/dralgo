package main

import (
	"fmt"
	"os"

	gen "dralgo/internal/antlr"
	antlr4 "github.com/antlr4-go/antlr/v4"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: debug_lexer <input_file>")
		os.Exit(1)
	}

	input, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}

	is := antlr4.NewInputStream(string(input))
	lexer := gen.NewDrAlgoLexer(is)

	for {
		t := lexer.NextToken()
		if t.GetTokenType() == antlr4.TokenEOF {
			break
		}
		fmt.Printf("Token: %q, Type: %d, Line: %d, Col: %d\n", t.GetText(), t.GetTokenType(), t.GetLine(), t.GetColumn())
	}
}
