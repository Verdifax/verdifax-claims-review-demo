// Unit tests for the Node program: the same cases every language tests.
import assert from "node:assert/strict";
import { mkdirSync, mkdtempSync, readdirSync, readFileSync, rmSync, writeFileSync } from "node:fs";
import { tmpdir } from "node:os";
import { createHash } from "node:crypto";
import { dirname, join } from "node:path";
import { test } from "node:test";
import { fileURLToPath } from "node:url";
import {
  canonical, DELETED_STEP_SUFFIX, deletedSteps, draftClassification, FEED_AGE_REASON, HISTORY_AGE_REASON, loadRules, MIN_VERIFIER, printVerifyCommands,
  MODEL_RETRY_TRIES, readInput, RULE_SET, SAVED_HEAD_MAX_AGE, STOP_KINDS, STOP_RECEIPT_FIELDS, stopReceipt, stopReceiptHash, stoppedRun, writeDeletedSteps,
  approverFiles, approverNameError, clientFor, DATA_QUALITY_CHECKS, DATA_QUALITY_SOURCE, dataQuality, POLICY_FILE, REPO, replaceWithTombstone, RULE_SET_SHA256,
  settingError, usageSubstance, withAgentSigningKey, publishable, SIGNOFF_PENDING_REASON, recordedUsage, withSigningKey, keyFor,
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

// R259: the SDK waits half a second before the first retry and doubles it to a
// cap of eight, so the shared budget must reach a minute of waiting. Round 61:
// the default of two retries waited about a second and a half, which a 503
// from the model API outlasted.
test("the model retry budget waits out an outage", () => {
  assert.equal(Number(PARITY.constants.model_retry_tries), MODEL_RETRY_TRIES);
  let waited = 0;
  for (let i = 0; i < MODEL_RETRY_TRIES - 1; i += 1) waited += Math.min(8, 0.5 * 2 ** i);
  assert.ok(waited >= 60, `the budget waits ${waited} seconds`);
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

// R250: a deleted step's tombstone, from the cases every language runs.
const DELETED = JSON.parse(readFileSync(join(HERE, "..", "..", "testdata", "deleted-step-cases.json"), "utf8"));

test("a deleted step is read from its job's run.deleted events only", () => {
  assert.deepEqual(deletedSteps(DELETED.events, DELETED.job_id), DELETED.want);
  assert.deepEqual(deletedSteps(DELETED.events, "claims-review-no-such-job"), {});
  assert.equal(DELETED_STEP_SUFFIX, PARITY.constants.deleted_step_suffix);
});

test("a deleted step's tombstone is written beside the steps, as the server sent it, job_links included", () => {
  const folder = mkdtempSync(join(tmpdir(), "claims-review-"));
  try {
    const texts = {};
    const deleted = {};
    for (const [step, served] of Object.entries(DELETED.want)) {
      const entry = structuredClone(served);
      // The SDK keeps the server's exact text; that text is what is written.
      texts[step] = JSON.stringify(entry);
      Object.defineProperty(entry, "sealedJson", { value: texts[step], enumerable: false });
      deleted[step] = entry;
    }
    writeDeletedSteps(folder, deleted);
    assert.deepEqual(readdirSync(join(folder, "steps")).sort(), ["rule-02.deleted.json", "verdict.deleted.json"]);
    for (const step of Object.keys(texts)) {
      const written = readFileSync(join(folder, "steps", `${step}${DELETED_STEP_SUFFIX}`), "utf8");
      assert.equal(written, texts[step] + "\n");
      assert.deepEqual(JSON.parse(written).job_links, DELETED.want[step].job_links);
    }
  } finally {
    rmSync(folder, { recursive: true, force: true });
  }
});

test("every printed re-check of a saved history or feed states its head age limit", (t) => {
  const folder = mkdtempSync(join(tmpdir(), "claims-review-"));
  const here = process.cwd();
  const printed = [];
  t.mock.method(console, "log", (line) => printed.push(line));
  try {
    process.chdir(folder);
    mkdirSync(join("records", "job-1"), { recursive: true });
    writeFileSync(join("records", "job-1", "signoff-history.json"), "[]");
    writeFileSync("signals.json", "[]");
    printVerifyCommands("job-1");
  } finally {
    process.chdir(here);
    rmSync(folder, { recursive: true, force: true });
  }
  assert.equal(MIN_VERIFIER, PARITY.constants.min_verifier);
  assert.equal(printed[0], `verify with verdifax-verify ${MIN_VERIFIER} or later:`);
  const history = printed.findIndex((l) => l.includes("--signoff-history"));
  const feed = printed.findIndex((l) => l.includes("--signals"));
  assert.equal(printed[history], `  verdifax-verify --signoff-history ${join("records", "job-1", "signoff-history.json")} --signoff-head-max-age ${SAVED_HEAD_MAX_AGE}`);
  assert.equal(printed[history + 1], `  ${HISTORY_AGE_REASON}`);
  assert.equal(printed[feed], `  verdifax-verify --signals signals.json --feed-head-max-age ${SAVED_HEAD_MAX_AGE}`);
  assert.equal(printed[feed + 1], `  ${FEED_AGE_REASON}`);
  assert.equal(SAVED_HEAD_MAX_AGE, PARITY.constants.saved_head_max_age);
  assert.equal(HISTORY_AGE_REASON, PARITY.constants.history_age_reason);
  assert.equal(FEED_AGE_REASON, PARITY.constants.feed_age_reason);
});

// Round 62: the cases every language runs for the every-capability demo.
const USAGE = JSON.parse(readFileSync(join(HERE, "..", "..", "testdata", "usage-cases.json"), "utf8"));
const QUALITY = JSON.parse(readFileSync(join(HERE, "..", "..", "testdata", "data-quality-cases.json"), "utf8"));
const COMMANDS = JSON.parse(readFileSync(join(HERE, "..", "..", "testdata", "command-cases.json"), "utf8"));

test("usage is declared for every model step", () => {
  for (const c of USAGE.cases) {
    const got = usageSubstance(c.usage);
    assert.deepEqual(got, c.want, c.name);
    assert.equal(got === null ? null : canonical(got), c.want_canonical, c.name);
  }
  for (const c of USAGE.recorded) {
    assert.deepEqual(recordedUsage(c.answer), { inputTokens: c.want.input_tokens, outputTokens: c.want.output_tokens }, c.name);
  }
});

test("the data checks are computed", () => {
  assert.equal(RULE_SET_SHA256, QUALITY.rule_set_sha256);
  assert.equal(RULE_SET_SHA256, PARITY.constants.rule_set_sha256);
  assert.deepEqual(DATA_QUALITY_CHECKS, PARITY.constants.data_quality_checks);
  assert.equal(DATA_QUALITY_SOURCE, PARITY.constants.data_quality_source);
  for (const c of QUALITY.cases) {
    const got = dataQuality(c.intake_draft_sha256, c.draft_sha256_now, c.rule_set_hash, QUALITY.as_of);
    assert.deepEqual(got, c.want, c.name);
    assert.equal(canonical(got), c.want_canonical, c.name);
  }
});

test("the loaded rule set is the pinned one", () => {
  assert.equal(loadRules().ruleSetHash, RULE_SET_SHA256);
});

test("the policy is v4", () => {
  assert.equal(POLICY_FILE, PARITY.constants.policy_file);
  const policy = JSON.parse(readFileSync(join(REPO, "policy", POLICY_FILE), "utf8"));
  assert.deepEqual(policy.required_attested_fields, ["actor_id"]);
});

test("settings are checked before sending", () => {
  for (const [setting, value] of COMMANDS.settings.valid) assert.equal(settingError(setting, value), null);
  for (const c of COMMANDS.settings.invalid) assert.equal(settingError(...c.args), c.error);
});

test("approver names and files", () => {
  for (const name of COMMANDS.approver_names.valid) assert.equal(approverNameError(name), null, name);
  for (const name of COMMANDS.approver_names.invalid) assert.equal(approverNameError(name), COMMANDS.approver_names.error, name);
  const f = COMMANDS.approver_names.files;
  assert.deepEqual(approverFiles(f.name), [f.private_key, f.public_key]);
});

test("the agent's request key in the keys file", () => {
  const c = COMMANDS.request_key;
  assert.deepEqual(withAgentSigningKey(c.keys, c.rotated_to), c.after_rotate);
  assert.deepEqual(withAgentSigningKey(c.keys, null), c.after_clear);
  assert.ok(clientFor(c.after_clear.agent), "a key whose request key was cleared still makes a client, which sends unsigned");
  const h = c.holder;
  assert.deepEqual(withSigningKey(h.keys, "holder", h.rotated_to), h.after_rotate);
  assert.deepEqual(withSigningKey(h.keys, "holder", null), h.after_clear);
  assert.deepEqual(keyFor(h.keys, true), ["holder", h.keys.holder]);
  assert.deepEqual(keyFor(h.keys, false), ["agent", h.keys.agent]);
});

test("a deleted step is replaced by its tombstone", () => {
  const c = COMMANDS.delete_step;
  assert.deepEqual(deletedSteps(c.feed, c.job_id)[c.step], c.want_tombstone);
  const folder = mkdtempSync(join(tmpdir(), "claims-review-"));
  try {
    mkdirSync(join(folder, "steps"));
    writeFileSync(join(folder, "steps", `${c.step}.json`), JSON.stringify(c.record));
    replaceWithTombstone(folder, c.step, c.want_tombstone);
    for (const f of c.want_files) assert.deepEqual(JSON.parse(readFileSync(join(folder, f), "utf8")), c.want_tombstone);
    for (const f of c.want_removed) assert.ok(!readdirSync(join(folder, "steps")).includes(f.split("/").pop()), f);
  } finally {
    rmSync(folder, { recursive: true, force: true });
  }
});

test("delete-step needs the step's sealed record", async () => {
  const { spawnSync } = await import("node:child_process");
  const c = COMMANDS.delete_step;
  const folder = mkdtempSync(join(tmpdir(), "claims-review-"));
  try {
    const run = spawnSync(process.execPath, [join(HERE, "..", "claims-review.mjs"), "delete-step", c.job_id, "rule-03"], { cwd: folder, encoding: "utf8" });
    assert.equal(run.status, 1);
    assert.ok(run.stderr.includes(c.missing_step_error), run.stderr);
  } finally {
    rmSync(folder, { recursive: true, force: true });
  }
});

test("publication skips a deleted step", () => {
  const c = COMMANDS.publication;
  const folder = mkdtempSync(join(tmpdir(), "claims-review-"));
  try {
    mkdirSync(join(folder, "steps"));
    for (const step of c.tombstones) writeFileSync(join(folder, "steps", `${step}${DELETED_STEP_SUFFIX}`), "{}");
    assert.deepEqual(publishable(folder, c.targets), c.want);
  } finally {
    rmSync(folder, { recursive: true, force: true });
  }
});

test("the account-policy check passes the sign-offs once given", (t) => {
  const c = COMMANDS.verify_commands;
  assert.equal(SIGNOFF_PENDING_REASON, PARITY.constants.signoff_pending_reason);
  const policy = join(REPO, "policy", POLICY_FILE);
  const want = (lines) => lines.map((l) => l.replace("POLICY", policy).replace("SIGNOFF_PENDING_REASON", SIGNOFF_PENDING_REASON));
  const folder = mkdtempSync(join(tmpdir(), "claims-review-"));
  const here = process.cwd();
  const printed = () => {
    const lines = [];
    t.mock.method(console, "log", (line) => lines.push(line));
    printVerifyCommands(c.job_id);
    t.mock.restoreAll();
    return lines.filter((l) => l.includes("--account-policy ") || l.startsWith("  verdifax-verify --signoff ") || l === `  ${SIGNOFF_PENDING_REASON}`);
  };
  try {
    process.chdir(folder);
    mkdirSync(join("records", c.job_id, "steps"), { recursive: true });
    assert.deepEqual(printed(), want(c.without_signoffs));
    writeFileSync(join("records", c.job_id, "signoffs.json"), "[]");
    assert.deepEqual(printed(), want(c.with_signoffs));
  } finally {
    process.chdir(here);
    rmSync(folder, { recursive: true, force: true });
  }
});
