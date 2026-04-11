package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	antlr4 "github.com/antlr4-go/antlr/v4"

	gen "dralgo/internal/antlr"
)

func tokenName(lexer *gen.DrAlgoLexer, tok antlr4.Token) string {
	tt := tok.GetTokenType()
	if tt == antlr4.TokenEOF {
		return "EOF"
	}

	// Symbolic names are indexed by token type (>= 0).
	// Be defensive in case tt is out of range.
	if tt >= 0 && tt < len(lexer.SymbolicNames) {
		if name := lexer.SymbolicNames[tt]; name != "" {
			return name
		}
	}
	if tt >= 0 && tt < len(lexer.LiteralNames) {
		if lit := lexer.LiteralNames[tt]; lit != "" {
			return lit
		}
	}
	return fmt.Sprintf("TOKEN_%d", tt)
}

func main() {
	inputPath := flag.String("input", "", "Input file (stdin if omitted)")
	mode := flag.String("mode", "tokens", "tokens|parse")
	flag.Parse()

	var data []byte
	var err error
	if *inputPath == "" {
		data, err = io.ReadAll(os.Stdin)
	} else {
		data, err = os.ReadFile(*inputPath)
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, "read:", err)
		os.Exit(1)
	}

	is := antlr4.NewInputStream(string(data))
	lexer := gen.NewDrAlgoLexer(is)
	stream := antlr4.NewCommonTokenStream(lexer, antlr4.TokenDefaultChannel)

	if *mode == "tokens" {
		stream.Fill()
		for _, tok := range stream.GetAllTokens() {
			fmt.Printf("%-20s %q\n", tokenName(lexer, tok), tok.GetText())
		}
		return
	}

	parser := gen.NewDrAlgoParser(stream)
	tree := parser.Program()
	fmt.Println(tree.ToStringTree(parser.RuleNames, parser))
}
