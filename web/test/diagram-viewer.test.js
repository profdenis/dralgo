import { test } from "node:test";
import assert from "node:assert/strict";

import { clampScale, isUrlTooLong, MIN_SCALE, MAX_SCALE, MAX_URL_LENGTH } from "../diagram-viewer.js";

test("clampScale borne l'échelle entre 0.25 et 4", () => {
    assert.equal(clampScale(MIN_SCALE / 10), MIN_SCALE);
    assert.equal(clampScale(MAX_SCALE * 10), MAX_SCALE);
    assert.equal(clampScale(1), 1);
});

test("isUrlTooLong détecte les URL au-delà du seuil", () => {
    assert.equal(isUrlTooLong("a".repeat(MAX_URL_LENGTH)), false);
    assert.equal(isUrlTooLong("a".repeat(MAX_URL_LENGTH + 1)), true);
});
