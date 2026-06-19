# Commandes de compilation pour DrAlgo

Ce document récapitule les commandes nécessaires pour compiler les différentes parties du projet DrAlgo.

## 1. Compilation WebAssembly (WASM)

Pour compiler la version WebAssembly utilisée dans l'interface web :

```bash
GOOS=js GOARCH=wasm go build -o main.wasm cmd/wasm/main.go
```

Un script utilitaire est également disponible à la racine :
```bash
./build_wasm.sh
```

## 2. Compilation de l'exécutable CLI (dralgo)

Pour compiler l'outil en ligne de commande principal :

```bash
go build -o dralgo ./cmd/dralgo/main.go
```

## 3. Régénération du Parser ANTLR (Go)

Si vous modifiez la grammaire (`grammar/dralgo/DrAlgo.g4`), vous devez régénérer les sources Go :

```bash
# Définir le chemin vers le JAR d'ANTLR (exemple)
export ANTLR_JAR="./antlr-4.13.1-complete.jar"
./scripts/generate_antlr_go_sources.sh
```

## 4. Autres commandes utiles

### Exécuter les tests
```bash
go test ./...
```

### Formater le code
```bash
go fmt ./...
```
