#!/usr/bin/env bash
set -euo pipefail

# Regenerate Go parser/lexer from the ANTLR grammar.
#
# Expected layout:
#   grammar/dralgo/DrAlgo.g4        (source grammar)
#   internal/antlr/                 (generated Go code, committed)
#
# Requirements:
#   - Java installed
#   - ANTLR tool available as a jar (set ANTLR_JAR env var)
#
# Example:
#   export ANTLR_JAR="$HOME/tools/antlr-4.x-complete.jar"
#   ./scripts/gen-antlr.sh

REPO_ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
GRAMMAR_FILE="${REPO_ROOT}/grammar/dralgo/DrAlgo.g4"
OUT_DIR="${REPO_ROOT}/internal/antlr"
PACKAGE_NAME="antlr"

if [[ ! -f "${GRAMMAR_FILE}" ]]; then
  echo "ERROR: Grammar file not found: ${GRAMMAR_FILE}" >&2
  exit 1
fi

: "${ANTLR_JAR:?ERROR: ANTLR_JAR is not set. Example: export ANTLR_JAR=/path/to/antlr-4.x-complete.jar}"

mkdir -p "${OUT_DIR}"

echo "==> Cleaning old generated files in: ${OUT_DIR}"
# Remove common generated file patterns; keep README.md (if present).
find "${OUT_DIR}" -maxdepth 1 -type f \
  \( -name "*.go" -o -name "*.tokens" -o -name "*.interp" \) \
  ! -name "README.md" \
  -print -delete

echo "==> Generating Go sources from: ${GRAMMAR_FILE}"
# -Dlanguage=Go : generate Go target
# -visitor      : generate visitor interfaces/base
# -no-listener  : reduce output surface (enable listener later if you want)
# -package      : set Go package name in generated files
java -jar "${ANTLR_JAR}" \
  -Dlanguage=Go \
  -visitor \
  -no-listener \
  -package "${PACKAGE_NAME}" \
  -o "${OUT_DIR}" \
  "${GRAMMAR_FILE}"

echo "==> Generated files:"
ls -1 "${OUT_DIR}"

echo "==> gofmt on generated .go files"
gofmt -w "${OUT_DIR}"/*.go

echo "==> Done. Remember: internal/antlr is generated and committed."
