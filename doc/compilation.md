# Commandes de compilation pour DrAlgo

Ce document récapitule les commandes nécessaires pour compiler les différentes parties du projet DrAlgo.

## 1. Compilation WebAssembly (WASM)

Pour compiler la version WebAssembly utilisée dans l'interface web :

```bash
./build_wasm.sh
```

Le script produit dans le dossier `web/` :
- `main.wasm` : le traducteur compilé pour le navigateur ;
- `wasm_exec.js` : le runtime Go pour WebAssembly, copié depuis le SDK Go ;
- `examples.json` : la liste des exemples, générée depuis `examples/algo/*.algo` par `cmd/gen-examples`.

Pour régénérer uniquement la liste des exemples :
```bash
go run ./cmd/gen-examples -in examples/algo -o web/examples.json
```

L'interface web s'ouvre depuis le dossier `web/` (qui contient `index.html`) : il faut servir ce dossier, par exemple avec `python3 -m http.server --directory web`, pour que le chargement de `main.wasm` et `examples.json` fonctionne.

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
