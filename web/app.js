import { encodePlantUML } from "./plantuml-encoder.js";
import { DiagramViewer } from "./diagram-viewer.js";
import { TextZoom } from "./text-zoom.js";

const go = new Go();

const exampleSelect = document.getElementById("example-select");
const algoInput = document.getElementById("algo-input");
const langSelect = document.getElementById("lang-select");
const generateButton = document.getElementById("generate-button");
const zoomInButton = document.getElementById("zoom-in-button");
const zoomOutButton = document.getElementById("zoom-out-button");
const zoomResetButton = document.getElementById("zoom-reset-button");
const diagramPanel = document.getElementById("diagram-panel");
const codeOutput = document.getElementById("code-output");
const resizer = document.getElementById("resizer");
const layout = document.getElementById("layout");
const editorPane = document.getElementById("editor-pane");
const mainResizer = document.getElementById("main-resizer");
const layoutToggleButton = document.getElementById("layout-toggle-button");
const themeToggleButton = document.getElementById("theme-toggle-button");

const examplesById = new Map();
const diagramViewer = new DiagramViewer(diagramPanel);
const textZooms = new Map();

document.querySelectorAll(".zoom-toolbar").forEach((toolbar) => {
    const target = document.getElementById(toolbar.dataset.zoomTarget);
    const textZoom = new TextZoom(target);
    textZooms.set(target.id, textZoom);

    toolbar.addEventListener("click", (e) => {
        const action = e.target.dataset.zoomAction;
        if (action === "in") textZoom.zoomIn();
        if (action === "out") textZoom.zoomOut();
        if (action === "reset") textZoom.reset();
    });
});

function loadWasm() {
    return WebAssembly.instantiateStreaming(fetch("main.wasm"), go.importObject).then((result) => {
        go.run(result.instance);
        console.log("WASM DrAlgo prêt");
    });
}

function loadExamples() {
    return fetch("examples.json")
        .then((response) => {
            if (!response.ok) {
                throw new Error(`HTTP ${response.status}`);
            }
            return response.json();
        })
        .then((examples) => {
            examples.forEach((example) => {
                examplesById.set(example.id, example);
                const option = document.createElement("option");
                option.value = example.id;
                option.textContent = example.title;
                exampleSelect.appendChild(option);
            });
        })
        .catch((error) => {
            console.error("Chargement des exemples impossible :", error);
            codeOutput.textContent = "Les exemples n'ont pas pu être chargés. Vous pouvez saisir votre algorithme directement.";
        });
}

function showDiagram(puml) {
    const url = "https://www.plantuml.com/plantuml/svg/" + encodePlantUML(puml);
    diagramViewer.setImage(url, "PlantUML Diagram");
}


function generate() {
    const input = algoInput.value;
    const lang = langSelect.value;

    codeOutput.textContent = "";

    if (typeof dralgoTranslate !== "function") {
        alert("Le module WebAssembly n'est pas encore prêt.");
        return;
    }

    if (input.trim() === "") {
        codeOutput.textContent = "Saisissez un algorithme avant de générer.";
        return;
    }

    const diagram = dralgoTranslate(input, "puml");
    if (!diagram.ok) {
        codeOutput.textContent = diagram.error;
        return;
    }
    showDiagram(diagram.code);

    if (lang === "puml") {
        codeOutput.textContent = diagram.code;
        return;
    }

    const result = dralgoTranslate(input, lang);
    codeOutput.textContent = result.ok ? result.code : result.error;
}

exampleSelect.addEventListener("change", () => {
    const example = examplesById.get(exampleSelect.value);
    algoInput.value = example ? example.source : "";
});

generateButton.addEventListener("click", generate);
zoomInButton.addEventListener("click", () => diagramViewer.zoomIn());
zoomOutButton.addEventListener("click", () => diagramViewer.zoomOut());
zoomResetButton.addEventListener("click", () => diagramViewer.reset());

// Logique des séparateurs (diagramme/code et éditeur/résultats)
const splits = [
    { handle: resizer, container: resizer.closest(".result-container"), pane: diagramPanel.closest(".diagram-column") },
    { handle: mainResizer, container: layout, pane: editorPane },
];
let activeSplit = null;

splits.forEach((split) => {
    split.handle.addEventListener("mousedown", () => {
        activeSplit = split;
        document.body.style.cursor = "col-resize";
    });
});

document.addEventListener("mousemove", (e) => {
    if (!activeSplit) return;

    const containerRect = activeSplit.container.getBoundingClientRect();
    const relativeX = e.clientX - containerRect.left;
    const percentage = (relativeX / containerRect.width) * 100;

    if (percentage > 10 && percentage < 90) {
        activeSplit.pane.style.flex = `0 0 ${percentage}%`;
    }
});

document.addEventListener("mouseup", () => {
    activeSplit = null;
    document.body.style.cursor = "default";
});

layoutToggleButton.addEventListener("click", () => {
    const sideBySide = !layout.classList.contains("side-by-side");
    layout.classList.toggle("side-by-side", sideBySide);
    editorPane.style.flex = "";
    layoutToggleButton.textContent = sideBySide ? "Mode vertical" : "Mode côte à côte";
    layoutToggleButton.setAttribute("aria-pressed", String(sideBySide));
});

const THEME_STORAGE_KEY = "dralgo-theme";

function applyTheme(theme) {
    const dark = theme === "dark";
    document.documentElement.dataset.theme = dark ? "dark" : "light";
    themeToggleButton.textContent = dark ? "Thème clair" : "Thème sombre";
    themeToggleButton.setAttribute("aria-pressed", String(dark));
}

function initialTheme() {
    try {
        const stored = localStorage.getItem(THEME_STORAGE_KEY);
        if (stored === "dark" || stored === "light") return stored;
    } catch (e) {
        console.warn("Lecture du thème impossible :", e);
    }
    return window.matchMedia("(prefers-color-scheme: dark)").matches ? "dark" : "light";
}

themeToggleButton.addEventListener("click", () => {
    const next = document.documentElement.dataset.theme === "dark" ? "light" : "dark";
    applyTheme(next);
    try {
        localStorage.setItem(THEME_STORAGE_KEY, next);
    } catch (e) {
        console.warn("Sauvegarde du thème impossible :", e);
    }
});

applyTheme(initialTheme());

loadExamples();
loadWasm();
