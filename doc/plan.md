## Document de migration “structure + ANTLR”

### Objectif
Réorganiser le module Go `dralgo` pour :
- isoler la **grammaire ANTLR** (`.g4`)
- isoler le **code généré** ANTLR (Go)
- introduire une couche **AST** (langage)
- introduire une couche **IR** (flowchart interne)
- garder des générateurs séparés (Python / Go / PlantUML)
- garder des binaires dans `cmd/...`

Ce document décrit **l’arborescence cible**, les **règles de dépendances** entre packages, et une **checklist** de migration.

---

## 1) Arborescence cible

Créer/obtenir cette structure :

```plain text
dralgo/
  cmd/
    dralgo/
      main.go
    puml2code/
      main.go

  grammar/
    dralgo/
      DrAlgo.g4

  internal/
    antlr/
      README.md                 # optionnel: notes generation
      # fichiers générés ANTLR Go (lexer/parser/visitor/listener)
      # ex: dralgo_lexer.go, dralgo_parser.go, dralgo_visitor.go, etc.

    parser/
      parser.go                 # API: ParseProgram(input string) (*ast.Program, error)
      errors.go                 # conversion erreurs ANTLR -> erreurs “propres”
      visitor.go                # Visitor ANTLR -> AST

    ast/
      nodes.go                  # types AST (Program, Stmt, Expr, etc.)
      span.go                   # positions (line/col/range) optionnel

    ir/
      ir.go                     # types IR (Flowchart, Blocks, Edges, etc.)
      lower.go                  # AST -> IR

    gen/
      go/
        generator.go            # IR -> Go
      python/
        generator.go            # IR -> Python
      plantuml/
        generator.go            # IR -> PlantUML

  examples/
  go.mod
  go.sum
```


### Notes
- `grammar/` contient **uniquement** les sources `.g4`.
- `internal/antlr/` contient **uniquement** le code Go généré par ANTLR (et éventuellement un README).
- `internal/parser/` est ton wrapper “propre” : c’est le seul endroit qui importe `internal/antlr`.
- Les générateurs n’importent **jamais** ANTLR directement ; ils travaillent sur `ast` ou `ir`.

---

## 2) Règles de dépendances (à respecter)

Dépendances autorisées :

- `cmd/...` → `internal/parser`, `internal/ir`, `internal/gen/...`
- `internal/parser` → `internal/antlr`, `internal/ast`
- `internal/ir` → `internal/ast`
- `internal/gen/...` → `internal/ir` (et éventuellement `internal/ast` si besoin)

Interdit (à éviter) :

- `internal/gen/...` → `internal/antlr`
- `internal/ast` → `internal/parser`
- `internal/antlr` → n’importe quel autre package (il ne doit dépendre de rien du projet)

But : éviter que le code généré “pollue” le design.

---

## 3) API à exposer (contrat minimal)

Créer une fonction de parsing centrale :

- Package : `internal/parser`
- Fonction : `ParseProgram(input string) (*ast.Program, error)`

Le CLI `cmd/dralgo/main.go` :
- lit le fichier pseudo-code
- appelle `parser.ParseProgram`
- convertit AST → IR (`ir.Lower(program)`)
- appelle le générateur choisi (`gen/python`, `gen/go`, `gen/plantuml`)

---

## 4) Étapes de migration (checklist)

### Étape A — Préparer la structure de dossiers
1. Créer les dossiers listés plus haut.
2. Ajouter un fichier vide `grammar/dralgo/DrAlgo.g4` (la grammaire complète sera collée ici).
3. Ajouter des fichiers “squelettes” Go (placeholders) dans `internal/parser`, `internal/ast`, `internal/ir`, `internal/gen/...`.

### Étape B — Génération ANTLR (code Go)
Objectif : générer le lexer/parser Go dans `internal/antlr`.

1. Installer/configurer ANTLR (hors Go) + runtime Go (dépendance Go).
2. Générer :
    - lexer
    - parser
    - visitor (recommandé)  
      dans `internal/antlr`.
3. Vérifier que `go test ./...` compile (même si ça ne fait rien encore).

**Décision repo** :
- soit commiter les fichiers générés dans git (souvent pratique),
- soit ne pas les commiter et fournir un script de génération + CI (plus strict).
  Recommandation pour un petit projet : **commiter le généré** au début.

### Étape C — AST minimal
Créer `internal/ast/nodes.go` avec :
- `type Program struct { Stmts []Stmt }`
- `Stmt` (interface) + structs : `AssignStmt`, `PrintStmt`, `IfStmt`, `WhileStmt`, `DoWhileStmt`, `ForStmt`, `ForEachStmt`, `IncDecStmt`, `ReadStmt`, `ReadIntStmt`
- `Expr` (interface) + structs : `BinaryExpr`, `UnaryExpr`, `IdentExpr`, `LiteralExpr{Int/Float/String}`, `ListExpr`

Optionnel : `Span` (line/col) attaché aux nœuds pour messages d’erreur.

### Étape D — Visitor ANTLR → AST
Dans `internal/parser/visitor.go` :
- implémenter un visitor qui transforme l’arbre ANTLR en AST.
- ne pas faire de génération de code ici.

Dans `internal/parser/parser.go` :
- instancier lexer/parser
- configurer error listener
- appeler la rule racine `program`
- passer l’arbre au visitor
- retourner `*ast.Program`

### Étape E — IR + Lowering (AST → IR)
Créer une IR simple :
- blocs séquentiels
- nœuds “decision” pour `if` et boucles
- edges conditionnels

Dans `internal/ir/lower.go` :
- transformer `ast.Program` en `ir.Flowchart`.

### Étape F — Générateurs (IR → sorties)
Implémenter trois générateurs :
- `internal/gen/plantuml`: IR → texte PlantUML
- `internal/gen/python`: IR → code Python
- `internal/gen/go`: IR → code Go

**Règle** : chaque générateur ne dépend que de l’IR (et d’outils stdlib).

### Étape G — CLI principal `cmd/dralgo`
Créer `cmd/dralgo/main.go` :
- flags : `-lang` (`python|go|plantuml`), `-input`, `-output`
- pipeline :
    1) read input
    2) `parser.ParseProgram`
    3) `ir.Lower`
    4) `generator.Generate`
    5) write output

Garder `cmd/puml2code` si nécessaire, mais l’objectif est que `cmd/dralgo` devienne l’entrée principale pour le pseudo-code.

---

## 5) Conventions de nommage

- Grammaire : `DrAlgo.g4`
- Package généré : `internal/antlr` (package name Go généré : `antlr` ou `dralgoantlr`, à fixer)
- `internal/parser` : package `parser`
- `internal/ast` : package `ast`
- `internal/ir` : package `ir`

---

## 6) Tests (minimum viable)
Ajouter dès que possible :
- `internal/parser/parser_test.go` :
    - parse un petit programme
    - vérifie le nombre de statements / quelques nœuds
- `internal/ir/lower_test.go` :
    - un `if` + une boucle
- `internal/gen/python/generator_test.go` :
    - snapshots simples

---

## 7) Définition “done” (critères d’acceptation)
- Le projet compile : `go test ./...`
- `cmd/dralgo` peut :
    - parser un exemple
    - produire une sortie (au moins PlantUML ou Python)
- Les packages respectent les dépendances (générateurs ne dépendent pas d’ANTLR)


## Addendum

### Décisions
1) **Le code ANTLR généré est commité** dans le dépôt.
2) Le code généré utilise le **package Go `antlr`** (donc dossier `internal/antlr` avec `package antlr`).

---

## 1) Emplacement des fichiers

### Source grammaire
- Fichier : `grammar/dralgo/DrAlgo.g4`

### Code Go généré (commité)
- Dossier : `internal/antlr/`
- Contenu : fichiers générés (lexer, parser, visitor/listener)
- Package : `antlr` (déclaré en tête de chaque fichier généré)

---

## 2) Génération — règles et commandes (à documenter / automatiser)

### Règles de génération
- Générer **Visitor** (recommandé) et éventuellement Listener (facultatif).
- La génération doit :
    - écrire les fichiers dans `internal/antlr`
    - forcer le package Go à `antlr`

### Script recommandé
Créer un script (ex : `scripts/gen-antlr.sh`) qui :
1) supprime l’ancien généré (optionnel mais propre)
2) lance ANTLR avec la cible Go
3) écrit dans `internal/antlr`
4) vérifie que les fichiers générés commencent bien par `package antlr`

> Le script est la “source de vérité” pour régénérer, même si on commit le généré.

---

## 3) README génération (bloc à inclure)

Ajouter un README (par ex `internal/antlr/README.md` ou `grammar/README.md`) avec :

- Où est la grammaire : `grammar/dralgo/DrAlgo.g4`
- Où est le généré : `internal/antlr`
- Comment régénérer (commande ou script)
- Mention : “les fichiers dans `internal/antlr` sont générés et commités”

---

## 4) Convention “ne pas importer ANTLR partout”

### Interdit
- `internal/gen/...` ne doit **jamais** importer `internal/antlr`.

### Seul import autorisé
- `internal/parser` est le seul package autorisé à importer `internal/antlr`.

Objectif : garder une séparation nette entre parsing et génération.

---

## 5) API finale attendue

### Parsing
- `internal/parser.ParseProgram(input string) (*ast.Program, error)`

### Pipeline CLI
- `cmd/dralgo/main.go` :
    1) read input
    2) parse → AST
    3) lower AST → IR
    4) generate output

---

## 6) Critères d’acceptation (Done)

- `internal/antlr` existe et contient du Go compilable avec `package antlr`
- `go test ./...` passe
- `cmd/dralgo` compile
- Le parsing ne dépend pas du code de génération (AST/IR/générateurs isolés)

---

## 7) Notes pour éviter les surprises

- Le nom `antlr` pour le package généré est OK, mais attention à ne pas le confondre mentalement avec le runtime ANTLR (souvent importé sous un nom proche).  
  Recommandation : dans le code wrapper (`internal/parser`), importer le runtime avec un alias clair si nécessaire (ex : `antlr4 "github.com/...antlr..."`) et importer le généré normalement (`dralgo/internal/antlr`).

