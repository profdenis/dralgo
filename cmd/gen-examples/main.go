package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type example struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Source string `json:"source"`
}

var numberPrefix = regexp.MustCompile(`^\d+_`)

func main() {
	input := flag.String("in", "examples/algo", "Directory containing .algo files")
	output := flag.String("o", "web/examples.json", "Output JSON file")
	flag.Parse()

	paths, err := filepath.Glob(filepath.Join(*input, "*.algo"))
	if err != nil {
		fmt.Printf("Error listing examples: %v\n", err)
		os.Exit(1)
	}

	examples := make([]example, 0, len(paths))
	for _, path := range paths {
		data, err := os.ReadFile(path)
		if err != nil {
			fmt.Printf("Error reading %s: %v\n", path, err)
			os.Exit(1)
		}

		id := strings.TrimSuffix(filepath.Base(path), ".algo")
		examples = append(examples, example{
			ID:     id,
			Title:  titleFromID(id),
			Source: string(data),
		})
	}

	payload, err := json.MarshalIndent(examples, "", "  ")
	if err != nil {
		fmt.Printf("Error encoding examples: %v\n", err)
		os.Exit(1)
	}

	if err := os.MkdirAll(filepath.Dir(*output), 0755); err != nil {
		fmt.Printf("Error creating output directory: %v\n", err)
		os.Exit(1)
	}

	if err := os.WriteFile(*output, append(payload, '\n'), 0644); err != nil {
		fmt.Printf("Error writing output file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("%d exemples écrits dans %s\n", len(examples), *output)
}

func titleFromID(id string) string {
	name := strings.ReplaceAll(numberPrefix.ReplaceAllString(id, ""), "_", " ")
	if name == "" {
		return name
	}
	return strings.ToUpper(name[:1]) + name[1:]
}
