#!/bin/bash
set -e

mkdir -p web
GOOS=js GOARCH=wasm go build -o web/main.wasm cmd/wasm/main.go
cp "$(go env GOROOT)/lib/wasm/wasm_exec.js" web/wasm_exec.js
go run ./cmd/gen-examples -in examples/algo -o web/examples.json
echo "Build web terminé : web/main.wasm, web/wasm_exec.js, web/examples.json"
