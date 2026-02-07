# DrAlgo grammar & ANTLR generation

This project uses **ANTLR4** to parse the DrAlgo pseudo-code language.

## Locations

- Grammar source (edit this):
  - `grammar/dralgo/DrAlgo.g4`

- Generated Go parser/lexer (DO NOT edit manually; committed to git):
  - `internal/antlr/` (Go package name: `antlr`)

Only the wrapper package `internal/parser` should import `internal/antlr`.  
Generators should not depend on ANTLR types.

## Regenerating the parser (Linux)

### Prerequisites

- Java installed (to run the ANTLR tool)
- ANTLR tool jar available locally

Set the environment variable `ANTLR_JAR` to the ANTLR jar path:
```
sh
export ANTLR_JAR="/path/to/antlr-4.x-complete.jar"
```
Then run:
```
sh
./scripts/gen-antlr.sh
```
This will:
- delete old generated files in `internal/antlr/` (keeps `README.md` if present)
- generate lexer/parser/visitor for Go
- set the generated Go package to `antlr`
- run `gofmt` on the generated `.go` files

## Notes / Policy

- The contents of `internal/antlr/` are **generated and committed**.
- Do not hand-edit files in `internal/antlr/`.
- If you change the grammar, re-run generation and commit the updated generated sources.

## Troubleshooting

- If you see `ANTLR_JAR is not set`, export it as shown above.
- If compilation fails after regeneration, ensure:
  - the grammar file path is correct
  - the `-package antlr` option is used (so `package antlr` matches the folder import)
  - `go test ./...` runs cleanly


