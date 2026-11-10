// Unit tests for the Node program: the same cases every language tests.
import assert from "node:assert/strict";
import { readdirSync, readFileSync } from "node:fs";
import { createHash } from "node:crypto";
import { dirname, join } from "node:path";
import { test } from "node:test";
import { fileURLToPath } from "node:url";
import {
  canonical, draftClassification, loadRules, readInput, RULE_SET, STOP_KINDS, STOP_RECEIPT_FIELDS, stopReceipt, stopReceiptHash, stoppedRun,
} from "../claims-review.mjs";

const HERE = dirname(fileURLToPath(import.meta.url));
const FIXED = { b: 1, a: "x", c: [1, 2], d: { z: true, y: null }, e: 'ü "q" \\ /' };
const EXPECTED = '{"a":"x","b":1,"c":[1,2],"d":{"y":null,"z":true},"e":"ü \\"q\\" \\\\ /"}';
// The rule set in force and its hash are named once, in demo-parity.json.
const PARITY = JSON.parse(readFileSync(join(HERE, "..", "..", "demo-parity.json"), "utf8"));
// The stop cases every language runs, from one file (R152).
const STOPS = JSON.parse(readFileSync(join(HERE, "..", "..", "testdata", "stop-cases.json"), "utf8"));

test("canonical JSON matches every language", () => {
  assert.equal(canonical(FIXED), EXPECTED);
});

test("the rule set hash", () => {
  const { rules, ruleSetHash } = loadRules();
  assert.equal(rules.length, 15);
  assert.equal(RULE_SET, PARITY.constants.rule_set);
  assert.equal(ruleSetHash, PARITY.rule_set_sha256);
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

// Answers as the server would from stop-cases.json. Each receipt route
// serves only runs stopped by its own kind, so a fetch of the wrong kind
// fails the test.
function fakeAgent() {
  const status = Object.fromEntries(STOPS.receipts.map((r) => [r.run_id, r.status]));
  status[STOPS.not_stopped.run_id] = STOPS.not_stopped.status;
  const route = (kind) => async (runId) => {
    if (status[runId] !== kind) throw new Error(`the ${kind} route was asked for run ${runId}, which is ${status[runId]}`);
    const r = STOPS.receipts.find((x) => x.run_id === runId);
    return { [r.field]: { hash: r.hash } };
  };
  return {
    getRun: async (runId) => ({ status: status[runId] }),
    exportRuns: async ({ job_id }) => ({ runs: job_id === STOPS.job_id ? STOPS.export_rows.map((r) => ({ ...r })) : [] }),
    getDenyReceipt: route("pepg_deny"), getAivpT4HaltReceipt: route("aivp_t4_halt"), getCCVHaltReceipt: route("ccv_halt"),
    getMACCHaltReceipt: route("macc_halt"), getAttestationGapReceipt: route("attestation_gap"),
  };
}

// fail() prints and exits; the test turns the exit into an error carrying
// the printed message.
function catchFail(t) {
  const printed = [];
  t.mock.method(console, "error", (...args) => printed.push(args.slice(1).join(" ")));
  t.mock.method(process, "exit", (code) => { throw new Error(`exit ${code}: ${printed.at(-1)}`); });
}

test("the stop kinds match the parity file", () => {
  const { stop_statuses: statuses, stop_receipt_fields: fields } = PARITY.constants;
  assert.deepEqual(Object.keys(STOP_KINDS).sort(), [...statuses].sort());
  assert.deepEqual(STOP_RECEIPT_FIELDS, Object.fromEntries(statuses.map((s, i) => [s, fields[i]])));
});

test("a stopped run is found by its job and step", async (t) => {
  catchFail(t);
  const err = { status: STOPS.error.status, message: STOPS.error.message };
  for (const c of STOPS.stopped_run) {
    if (c.failure) {
      await assert.rejects(stoppedRun(fakeAgent(), STOPS.job_id, c.step, err), { message: `exit 1: ${c.failure}` });
    } else {
      const row = await stoppedRun(fakeAgent(), STOPS.job_id, c.step, err);
      assert.deepEqual([row.run_id, row.status], [c.run_id, c.status]);
    }
  }
});

test("a stop's receipt hash is read from the receipt of its kind", async () => {
  for (const c of STOPS.receipts) {
    assert.equal(await stopReceiptHash(fakeAgent(), c.run_id), c.hash);
    assert.deepEqual(await stopReceipt(fakeAgent(), c.status, c.run_id), { [c.field]: { hash: c.hash } });
  }
});

test("a run that did not stop has no receipt", async (t) => {
  catchFail(t);
  await assert.rejects(stopReceiptHash(fakeAgent(), STOPS.not_stopped.run_id), { message: `exit 1: ${STOPS.not_stopped.failure}` });
});

// R170: the program reads every input as its exact bytes, so the hash it seals
// is the one shasum prints, and the parity file lists every sample and policy.
test("every input file reads as its exact bytes", () => {
  const repo = join(HERE, "..", "..");
  const listed = Object.keys(PARITY.input_sha256).sort();
  const onDisk = ["samples", "policy"].flatMap((dir) => readdirSync(join(repo, dir)).map((f) => `${dir}/${f}`)).sort();
  assert.deepEqual(listed, onDisk);
  for (const file of listed) {
    const read = createHash("sha256").update(Buffer.from(readInput(join(repo, file)), "utf8")).digest("hex");
    assert.equal(read, PARITY.input_sha256[file], file);
  }
});
