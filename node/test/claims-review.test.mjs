// Unit tests for the Node program: the same cases every language tests.
import assert from "node:assert/strict";
import { readFileSync } from "node:fs";
import { dirname, join } from "node:path";
import { test } from "node:test";
import { fileURLToPath } from "node:url";
import { canonical, draftClassification, loadRules } from "../claims-review.mjs";

const HERE = dirname(fileURLToPath(import.meta.url));
const FIXED = { b: 1, a: "x", c: [1, 2], d: { z: true, y: null }, e: 'ü "q" \\ /' };
const EXPECTED = '{"a":"x","b":1,"c":[1,2],"d":{"y":null,"z":true},"e":"ü \\"q\\" \\\\ /"}';
const RULE_SET_HASH = "61110740e99a53d6fe48378a52a450f053ef2d887ef491be55801f75b4e4e1ad";

test("canonical JSON matches every language", () => {
  assert.equal(canonical(FIXED), EXPECTED);
});

test("the rule set hash", () => {
  const { rules, ruleSetHash } = loadRules();
  assert.equal(rules.length, 15);
  assert.equal(ruleSetHash, RULE_SET_HASH);
});

test("the classification line", () => {
  assert.equal(draftClassification("# A draft\n\nText."), "INTERNAL");
  assert.equal(draftClassification("<!-- classification: CONFIDENTIAL -->\n# Memo"), "CONFIDENTIAL");
});

test("the parity copy matches the source", () => {
  const source = readFileSync(join(HERE, "..", "..", "demo-parity.json"), "utf8");
  const copy = readFileSync(join(HERE, "..", "testdata", "demo-parity.json"), "utf8");
  assert.equal(copy, source);
});
