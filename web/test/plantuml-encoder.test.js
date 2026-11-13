import { test } from "node:test";
import assert from "node:assert/strict";
import zlib from "node:zlib";

import { encodePlantUML } from "../plantuml-encoder.js";

globalThis.pako = {
    deflateRaw: (bytes, options) => zlib.deflateRawSync(bytes, options),
};

// Copie de l'encodage d'origine (ancien index.html), pour la non-régression.
function legacyEncode64(data) {
    let r = "";
    for (let i = 0; i < data.length; i += 3) {
        if (i + 2 < data.length) {
            r += legacyAppend3bytes(data.charCodeAt(i), data.charCodeAt(i + 1), data.charCodeAt(i + 2));
        } else if (i + 1 < data.length) {
            r += legacyAppend3bytes(data.charCodeAt(i), data.charCodeAt(i + 1), 0);
        } else {
            r += legacyAppend3bytes(data.charCodeAt(i), 0, 0);
        }
    }
    return r;
}

function legacyAppend3bytes(b1, b2, b3) {
    const c1 = b1 >> 2;
    const c2 = ((b1 & 0x3) << 4) | (b2 >> 4);
    const c3 = ((b2 & 0xF) << 2) | (b3 >> 6);
    const c4 = b3 & 0x3F;
    let r = "";
    r += legacyEncode6bit(c1 & 0x3F);
    r += legacyEncode6bit(c2 & 0x3F);
    r += legacyEncode6bit(c3 & 0x3F);
    r += legacyEncode6bit(c4 & 0x3F);
    return r;
}

function legacyEncode6bit(b) {
    if (b < 10) return String.fromCharCode(48 + b);
    b -= 10;
    if (b < 26) return String.fromCharCode(65 + b);
    b -= 26;
    if (b < 26) return String.fromCharCode(97 + b);
    b -= 26;
    if (b === 0) return '-';
    if (b === 1) return '_';
    return '?';
}

function legacyCompress(s) {
    const bytes = new TextEncoder().encode(s);
    const compressed = globalThis.pako.deflateRaw(bytes, { level: 9 });
    let binary = "";
    for (let i = 0; i < compressed.byteLength; i++) {
        binary += String.fromCharCode(compressed[i]);
    }
    return legacyEncode64(binary);
}

const samples = [
    "",
    "a",
    "ab",
    "abc",
    "@startuml\nstart\n:x = 1;\nstop\n@enduml",
    "@startuml\nstart\nif (score >= 90) then (oui)\n:Excellent;\nendif\nstop\n@enduml",
];

for (const sample of samples) {
    test(`encodePlantUML garde l'encodage d'origine pour ${JSON.stringify(sample.slice(0, 20))}`, () => {
        assert.equal(encodePlantUML(sample), legacyCompress(sample));
    });
}
