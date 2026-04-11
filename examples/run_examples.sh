#!/usr/bin/env bash
set -euo pipefail

# Script pour exécuter dralgo sur tous les fichiers .algo d'un dossier.
# Usage: ./run_examples.sh <dossier_sources> <langage>

if [ "$#" -lt 2 ]; then
    echo "Usage: $0 <dossier_sources> <langage>"
    echo "Exemple: $0 algo python"
    exit 1
fi

SRC_DIR="$1"
LANG="$2"

# Détermination de l'extension de sortie
EXT=""
case "$LANG" in
    python)
        EXT="py"
        ;;
    puml|plantuml)
        EXT="puml"
        ;;
    *)
        echo "Erreur: Langage non supporté: $LANG"
        exit 1
        ;;
esac

# Chemin vers le binaire dralgo (on suppose qu'il est à la racine du projet ou dans le PATH)
# On cherche d'abord à la racine du projet par rapport à l'emplacement du script
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
DRALGO_BIN="${SCRIPT_DIR}/../dralgo"

if [ ! -f "$DRALGO_BIN" ]; then
    echo "Avertissement: $DRALGO_BIN non trouvé, tentative d'utiliser 'dralgo' du PATH..."
    DRALGO_BIN="dralgo"
fi

# Création du dossier de sortie
OUT_DIR="${SRC_DIR}/${LANG}"
mkdir -p "$OUT_DIR"

echo "==> Traitement des fichiers dans ${SRC_DIR} vers ${OUT_DIR} (lang: ${LANG})"

FILES_FOUND=0
for f in "${SRC_DIR}"/*.algo; do
    if [ -e "$f" ]; then
        FILES_FOUND=1
        FILENAME=$(basename "$f" .algo)
        OUTPUT_FILE="${OUT_DIR}/${FILENAME}.${EXT}"
        
        echo "  - ${FILENAME}.algo -> ${FILENAME}.${EXT}"
        "$DRALGO_BIN" -lang "$LANG" -o "$OUTPUT_FILE" "$f"
    fi
done

if [ "$FILES_FOUND" -eq 0 ]; then
    echo "Aucun fichier .algo trouvé dans ${SRC_DIR}"
fi

echo "==> Terminé."
