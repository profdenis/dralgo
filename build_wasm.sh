#!/bin/bash
GOOS=js GOARCH=wasm go build -o main.wasm cmd/wasm/main.go
echo "WASM build complete: main.wasm"
