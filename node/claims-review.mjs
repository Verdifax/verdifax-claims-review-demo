#!/usr/bin/env node
// claims-review: a small agent that checks a draft against 15 general writing
// rules, with every step sealed by Verdifax.
//
// A demo, not customer proof. A sealed record shows which rules were checked
// and what the model answered; it does not show the model's judgment was
// right. The 15 rules are general writing rules, not legal advice and not a
// compliance standard. The account policy judges what each step declared; it
// cannot show the declared facts are true. A sign-off shows a registered
// approver key signed that decision; it does not show who holds the key or
// that the review was careful.
//
// Commands (the same in every language of this repository):
//
//   claims-review setup                              owner key in VERDIFAX_OWNER_KEY
//   claims-review review DRAFT.md [--model ID]        model key in ANTHROPIC_API_KEY
//   claims-review sign-off JOB_ID approve|reject [--note TEXT]
//   claims-review sign-off --receipt RUN_ID reject
//   claims-review feed
//   claims-review verify-commands JOB_ID
//
// Environment: VERDIFAX_BASE_URL (default https://api.verdifax.com).
// Files, in the working folder: demo-keys.json and approver-key.json (0600),
// approver.txt, account-history.json, records/<job_id>/*.json with payloads/,
// signals.json.

import { createHash } from "node:crypto";
import { existsSync, mkdirSync, readFileSync, readdirSync, writeFileSync } from "node:fs";
import { dirname, join, resolve } from "node:path";
import { fileURLToPath } from "node:url";
import { APIError, VerdifaxClient, generateRequestKeyPair } from "verdifax";

const DEFAULT_BASE_URL = "https://api.verdifax.com";
const DEFAULT_MODEL = "claude-fable-5-1";
const FORM = "verdifax.claims-review.v1";
const RULE_SET = "v1";
const ACTOR_ID = "claims-review-agent";
const ROUTE_INTAKE = "claims-review/v1/intake";
const ROUTE_RULE = "claims-review/v1/rule-check";
const ROUTE_VERDICT = "claims-review/v1/verdict";
const OBJECT_PREFIX = "claims-review/v1/";
const LEVELS = ["PUBLIC", "INTERNAL", "CONFIDENTIAL", "RESTRICTED", "TOP_SECRET", "TS_SCI"];
const MIN_VERIFIER = "v0.24.0";
const REPO = resolve(dirname(fileURLToPath(import.meta.url)), "..");

const sha256Hex = (data) => createHash("sha256").update(data).digest("hex");

// Canonical JSON: sorted keys, no spaces, as every language of this
// repository writes it.
function canonical(value) {
  if (Array.isArray(value)) return "[" + value.map(canonical).join(",") + "]";
  if (value && typeof value === "object") {
    return "{" + Object.keys(value).sort().map((k) => JSON.stringify(k) + ":" + canonical(value[k])).join(",") + "}";
  }
  return JSON.stringify(value);
}

// Fixed identifiers, derived from fixed text so every language seals the same.
const PROGRAM_ID = sha256Hex("verdifax-claims-review-demo/program/v1");
const REGISTRY_RECORD_HASH = sha256Hex("verdifax-claims-review-demo/registry/v1");

const baseURL = () => (process.env.VERDIFAX_BASE_URL || DEFAULT_BASE_URL).replace(/\/+$/, "");
const isProduction = () => baseURL() === DEFAULT_BASE_URL;

function fail(message) {
  console.error("claims-review:", message);
  process.exit(1);
}

function loadRules() {
  const doc = JSON.parse(readFileSync(join(REPO, "rules", "v1.json"), "utf8"));
  return { rules: doc.rules, ruleSetHash: sha256Hex(canonical(doc)) };
}

const loadPrompt = () => readFileSync(join(REPO, "rules", "prompt-v1.txt"), "utf8");

function writePrivate(path, text) {
  writeFileSync(path, text, { mode: 0o600, flag: "w" });
}

function writeJSON(path, value) {
  mkdirSync(dirname(path), { recursive: true });
  writeFileSync(path, JSON.stringify(value, null, 2) + "\n");
}

function draftClassification(text) {
  const first = (text.split("\n")[0] || "").trim();
  const m = first.match(/^<!--\s*classification:\s*([A-Z_]+)\s*-->$/);
  const level = m ? m[1] : "INTERNAL";
  if (!LEVELS.includes(level)) fail("the draft declares an unknown classification: " + level);
  return level;
}

function loadKeys() {
  if (!existsSync("demo-keys.json")) fail("no demo-keys.json here; run `claims-review setup` first");
  return JSON.parse(readFileSync("demo-keys.json", "utf8"));
}

const clientFor = (key) =>
  new VerdifaxClient({ baseURL: baseURL(), apiKey: key.secret, signingKey: Buffer.from(key.signing_key, "hex") });

// ── setup ──────────────────────────────────────────────────────────────────

async function cmdSetup() {
  if (existsSync("demo-keys.json")) fail("demo-keys.json exists here; setup runs once per account and folder");
  const ownerSecret = process.env.VERDIFAX_OWNER_KEY || "";
  if (!ownerSecret) fail("set VERDIFAX_OWNER_KEY to the account's owner key");
  const ownerSigner = generateRequestKeyPair();
  const owner = new VerdifaxClient({ baseURL: baseURL(), apiKey: ownerSecret, signingKey: ownerSigner.privateKey });
  await owner.registerAccountRequestKey();
  const member = await owner.createAccountKey("claims-review-agent");
  const memberSigner = generateRequestKeyPair();
  const agent = new VerdifaxClient({ baseURL: baseURL(), apiKey: member.secret, signingKey: memberSigner.privateKey });
  await agent.registerRequestKey();
  const policyText = readFileSync(join(REPO, "policy", "claims-review-v1.json"), "utf8");
  const uploaded = await owner.uploadAccountPolicy(policyText);
  await owner.activateAccountPolicy(uploaded.document_hash);
  const approver = generateRequestKeyPair();
  await owner.registerApprover("Claims review approver", approver.privateKey);
  const feed = await owner.createFeedKey("claims-review-feed");
  writeJSON("account-history.json", await owner.getAccountPolicyHistory());
  writePrivate("approver-key.json", JSON.stringify({ private_key: approver.privateKey.toString("hex") }) + "\n");
  writeFileSync("approver.txt", approver.publicKey.toString("base64"));
  writePrivate(
    "demo-keys.json",
    JSON.stringify(
      {
        base_url: baseURL(),
        owner: { secret: ownerSecret, signing_key: ownerSigner.privateKey.toString("hex") },
        agent: { id: member.id, secret: member.secret, signing_key: memberSigner.privateKey.toString("hex") },
        feed: { id: feed.id, secret: feed.secret },
        policy_document_hash: uploaded.document_hash,
      },
      null,
      2,
    ) + "\n",
  );
  console.log(`setup done: agent key ${member.id}, feed key ${feed.id}, policy ${uploaded.document_hash}`);
}

// ── the model ──────────────────────────────────────────────────────────────

// Recorded answers for tests, refused against production.
function loadAnswers(path) {
  if (!path) return null;
  if (isProduction()) fail("--answers is for tests against a local server, never production");
  return JSON.parse(readFileSync(path, "utf8"));
}

async function ask(recorded, prompt, userText, ruleId, model) {
  if (recorded) {
    const a = recorded[ruleId];
    return {
      model: a.model || "recorded-answers",
      responseId: a.response_id || "recorded",
      verdict: a.verdict,
      reason: a.reason,
      raw: canonical({ verdict: a.verdict, reason: a.reason }),
      inputTokens: 0,
      outputTokens: 0,
    };
  }
  const { default: Anthropic } = await import("@anthropic-ai/sdk");
  const { zodOutputFormat } = await import("@anthropic-ai/sdk/helpers/zod");
  const { z } = await import("zod");
  const RuleAnswer = z.object({ verdict: z.string(), reason: z.string() });
  const client = new Anthropic();
  const response = await client.messages.parse({
    model,
    max_tokens: 4000,
    system: prompt,
    messages: [{ role: "user", content: userText }],
    output_config: { format: zodOutputFormat(RuleAnswer) },
  });
  const usage = { inputTokens: response.usage.input_tokens, outputTokens: response.usage.output_tokens };
  if (response.stop_reason === "refusal" || !response.parsed_output) {
    return { model: response.model, responseId: response.id, verdict: "unanswered", reason: "the model declined to answer", raw: "", ...usage };
  }
  const parsed = response.parsed_output;
  const verdict = parsed.verdict === "pass" || parsed.verdict === "fail" ? parsed.verdict : "unanswered";
  const raw = (response.content.find((b) => b.type === "text") || {}).text || "";
  return { model: response.model, responseId: response.id, verdict, reason: parsed.reason, raw, ...usage };
}

function measuredCost(model, inputTokens, outputTokens) {
  const prices = JSON.parse(readFileSync(join(REPO, "rules", "model-prices-v1.json"), "utf8")).prices;
  if (!prices[model]) return null;
  return (inputTokens / 1e6) * prices[model].input + (outputTokens / 1e6) * prices[model].output;
}

// ── review ─────────────────────────────────────────────────────────────────

function utcStamp() {
  return new Date().toISOString().replace(/[-:]/g, "").replace(/\.\d{3}Z$/, "Z");
}

async function cmdReview(draftPath, model, answersPath) {
  const keys = loadKeys();
  const agent = clientFor(keys.agent);
  const { rules, ruleSetHash } = loadRules();
  const prompt = loadPrompt();
  const draft = readFileSync(draftPath, "utf8");
  const draftHash = sha256Hex(draft);
  const level = draftClassification(draft);
  const recorded = loadAnswers(answersPath);
  const now = utcStamp();
  const jobId = `claims-review-${draftHash.slice(0, 12)}-${now}`;
  const out = join("records", jobId);
  const payloads = {};
  const policy = (objectId) => ({ object_type: "text", object_id: OBJECT_PREFIX + objectId, classification: level });

  const intakePayload = {
    form: FORM, step: "intake", draft_sha256: draftHash, draft_bytes: Buffer.byteLength(draft, "utf8"),
    rule_set: RULE_SET, rule_set_hash: ruleSetHash, classification: level,
  };
  const job = agent.job(jobId);
  let intake;
  try {
    intake = await job.step("intake", {
      after: [], payload_text: canonical(intakePayload), program_id: PROGRAM_ID, route_id: ROUTE_INTAKE,
      registry_record_hash: REGISTRY_RECORD_HASH, policy_context: policy("intake"),
    });
  } catch (err) {
    if (!(err instanceof APIError)) throw err;
    // The account policy refused the draft at intake (for example a
    // confidential draft), so it never reached a model. The refusal is
    // sealed as a receipt; nothing else runs for this draft.
    await refuse(agent, jobId, "intake", {}, { intake: intakePayload }, draftPath, draftHash, err.message);
    return;
  }
  payloads.intake = intakePayload;
  const runIds = { intake: intake.run_id };

  const results = [];
  const tokens = { input: 0, output: 0 };
  const models = new Set();
  for (const rule of rules) {
    const userText = `Rule ${rule.id}: ${rule.text}\n\nDraft:\n${draft}`;
    const promptHash = sha256Hex(prompt + "\n" + userText);
    const answer = await ask(recorded, prompt, userText, rule.id, model);
    tokens.input += answer.inputTokens;
    tokens.output += answer.outputTokens;
    models.add(answer.model);
    const payload = {
      form: FORM, step: rule.id, rule_id: rule.id, rule_text: rule.text, draft_sha256: draftHash, rule_set_hash: ruleSetHash,
      model: answer.model, model_response_id: answer.responseId, prompt_hash: promptHash, verdict: answer.verdict, reason: answer.reason,
    };
    let receipt;
    try {
      receipt = await job.step(rule.id, {
        after: ["intake"], payload_text: canonical(payload), program_id: PROGRAM_ID, route_id: ROUTE_RULE,
        registry_record_hash: REGISTRY_RECORD_HASH, ai_output_text: answer.raw || answer.reason,
        attested_context: {
          attested: true, actor_id: ACTOR_ID, model_provider: "anthropic", model_name: answer.model, prompt_hash: promptHash,
          decision_kind: "claims.rule_check", decision_result: answer.verdict,
        },
        policy_context: policy(rule.id),
      });
    } catch (err) {
      if (!(err instanceof APIError)) throw err;
      // The account policy or the server's content screen refused this
      // step, so the job stops here: the refusal is sealed as a receipt,
      // the steps already sealed are kept beside it, and a person sees
      // the stop. The job never gets a verdict.
      payloads[rule.id] = payload;
      await refuse(agent, jobId, rule.id, runIds, payloads, draftPath, draftHash, err.message);
      return;
    }
    payloads[rule.id] = payload;
    runIds[rule.id] = receipt.run_id;
    results.push({ rule_id: rule.id, verdict: answer.verdict, reason: answer.reason });
    console.log(`${rule.id} ${answer.verdict}: ${answer.reason}`);
  }

  const failed = results.filter((r) => r.verdict !== "pass").map((r) => r.rule_id);
  const verdictPayload = {
    form: FORM, step: "verdict", draft_sha256: draftHash, rule_set_hash: ruleSetHash,
    verdict: failed.length === 0 ? "pass" : "fail", failed_rules: failed, results,
  };
  const final = await job.step("verdict", {
    after: rules.map((r) => r.id), final: true, payload_text: canonical(verdictPayload), program_id: PROGRAM_ID,
    route_id: ROUTE_VERDICT, registry_record_hash: REGISTRY_RECORD_HASH,
    attested_context: { attested: true, actor_id: ACTOR_ID, decision_kind: "claims.verdict", decision_result: verdictPayload.verdict },
    policy_context: policy("verdict"),
  });
  payloads.verdict = verdictPayload;
  runIds.verdict = final.run_id;

  await saveJobRecords(agent, jobId, out, runIds, payloads);
  const cost = models.size === 1 && models.has(model) ? measuredCost(model, tokens.input, tokens.output) : null;
  const summary = {
    job_id: jobId, draft: draftPath, draft_sha256: draftHash, rule_set_hash: ruleSetHash, verdict: verdictPayload.verdict,
    failed_rules: failed, models: [...models].sort(), final_run_id: final.run_id, run_ids: runIds,
    measured_tokens: tokens, measured_cost_usd: cost,
  };
  writeJSON(join(out, "job.json"), summary);
  console.log(`verdict ${summary.verdict} (${failed.length} of ${rules.length} rules failed); job ${jobId}; records in ${out}`);
  if (cost !== null) {
    console.log(`measured model use: ${tokens.input} input and ${tokens.output} output tokens, about $${cost.toFixed(2)} at published prices`);
  }
  printVerifyCommands(jobId);
}

async function newestRefusal(agent) {
  const rows = (await agent.exportRuns()).runs.filter((r) => r.status === "pepg_deny");
  if (rows.length === 0) fail("the step was refused but no refused run is listed for this key");
  return Math.max(...rows.map((r) => r.run_id));
}

// Record a refused step: its sealed receipt, every record the job sealed
// before it, the exact payload bytes of every step including the refused
// one, and a summary naming the step. The same for intake and for any
// later step.
async function refuse(agent, jobId, step, runIds, payloads, draftPath, draftHash, reason) {
  const refused = await newestRefusal(agent);
  const folder = join("records", `refused-${refused}`); // named by the run, so two refusals never share a folder
  writeJSON(join(folder, "receipt.json"), await agent.getDenyReceipt(refused));
  mkdirSync(join(folder, "steps"), { recursive: true });
  for (const [sealedStep, runId] of Object.entries(runIds)) writeFileSync(join(folder, "steps", `${sealedStep}.json`), await agent.getRunArtifacts(runId));
  mkdirSync(join(folder, "payloads"), { recursive: true });
  for (const [name, payload] of Object.entries(payloads)) writeFileSync(join(folder, "payloads", `${name}.json`), canonical(payload));
  writeJSON(join(folder, "job.json"), { refused_run_id: refused, step, job_id: jobId, draft: draftPath, draft_sha256: draftHash, reason, run_ids: runIds });
  console.log(`refused at ${step}: run ${refused}, ${reason}`);
  console.log(`records in ${folder}`);
}

async function saveJobRecords(agent, jobId, out, runIds, payloads) {
  const rows = (await agent.exportRuns({ job_id: jobId })).runs;
  if (rows.length !== Object.keys(runIds).length) {
    fail(`the export lists ${rows.length} records for ${jobId}, the program sealed ${Object.keys(runIds).length}`);
  }
  // The sealed step records alone sit in steps/, the folder verdifax-verify
  // --job reads: it takes every .json file there as a record, so the
  // summary, the sign-offs and the payloads live beside that folder.
  mkdirSync(join(out, "steps"), { recursive: true });
  mkdirSync(join(out, "payloads"), { recursive: true });
  for (const row of rows) {
    if (row.status === "pepg_deny") writeJSON(join(out, "steps", `${row.step_id}.json`), await agent.getDenyReceipt(row.run_id));
    else writeFileSync(join(out, "steps", `${row.step_id}.json`), await agent.getRunArtifacts(row.run_id));
  }
  // The exact bytes each step sent, so a reader recomputes the sealed
  // payload hash from the file.
  for (const [step, payload] of Object.entries(payloads)) writeFileSync(join(out, "payloads", `${step}.json`), canonical(payload));
}

// ── sign-off ───────────────────────────────────────────────────────────────

function approverPrivateKey() {
  if (!existsSync("approver-key.json")) fail("no approver-key.json here; run `claims-review setup` first");
  return Buffer.from(JSON.parse(readFileSync("approver-key.json", "utf8")).private_key, "hex");
}

async function cmdSignOff(jobId, decision, note, receiptRunId) {
  const keys = loadKeys();
  const agent = clientFor(keys.agent);
  const owner = clientFor(keys.owner);
  const key = approverPrivateKey();
  let target;
  let folder;
  let runId;
  if (receiptRunId) {
    runId = Number(receiptRunId);
    const receipt = (await agent.getDenyReceipt(runId)).deny_receipt;
    target = { kind: "receipt", run_id: runId, digest: receipt.hash };
    folder = existsSync("records")
      ? readdirSync("records").filter((p) => p.startsWith("refused-")).map((p) => join("records", p))
          .find((p) => existsSync(join(p, "job.json")) && JSON.parse(readFileSync(join(p, "job.json"), "utf8")).refused_run_id === runId)
      : undefined;
    folder = folder || join("records", `receipt-${runId}`);
  } else {
    folder = join("records", jobId);
    const summary = JSON.parse(readFileSync(join(folder, "job.json"), "utf8"));
    runId = summary.final_run_id;
    const artifacts = JSON.parse(readFileSync(join(folder, "steps", "verdict.json"), "utf8"));
    target = { kind: "job", run_id: runId, digest: artifacts.bundle_anchor.anchored_digest, job_id: jobId };
  }
  const record = await agent.signOff(target, decision, key, note ? { note } : {});
  writeJSON(join(folder, "signoffs.json"), await agent.getSignOffs(runId));
  writeJSON(join(folder, "signoff-history.json"), await owner.getSignOffHistory());
  console.log(`sign-off ${decision} recorded as sequence ${record.sequence}, hash ${record.signoff_hash}`);
  const recordFile = receiptRunId ? join(folder, "receipt.json") : join(folder, "steps", "verdict.json");
  console.log(`verify: verdifax-verify --signoff ${join(folder, "signoffs.json")} --approver-key approver.txt ${recordFile}`);
  console.log(`verify: verdifax-verify --signoff-history ${join(folder, "signoff-history.json")}`);
}

// ── feed ───────────────────────────────────────────────────────────────────

// ── publish ────────────────────────────────────────────────────────────────
// The account's owner key makes a record public on its random link, or
// withdraws it. Each decision is sealed by the server as a run publication
// and placed in the account's feed, which `feed` then reads.

// The folder and the records a publish or unpublish acts on: every step of
// a job, or one refusal's receipt.
function publicationTargets(jobId, receipt) {
  if (receipt) {
    const runId = Number(receipt);
    const folder = (existsSync("records") ? readdirSync("records").filter((p) => p.startsWith("refused-")).map((p) => join("records", p)) : [])
      .find((p) => existsSync(join(p, "job.json")) && JSON.parse(readFileSync(join(p, "job.json"), "utf8")).refused_run_id === runId)
      || join("records", `receipt-${runId}`);
    return { folder, targets: { receipt: runId } };
  }
  const folder = join("records", jobId);
  const summary = JSON.parse(readFileSync(join(folder, "job.json"), "utf8"));
  return { folder, targets: { ...summary.run_ids } };
}

async function publication(jobId, receipt, publish) {
  const keys = loadKeys();
  const owner = clientFor(keys.owner);
  const { folder, targets } = publicationTargets(jobId, receipt);
  const published = {};
  for (const name of Object.keys(targets).sort()) {
    const runId = targets[name];
    const answer = publish ? await owner.publishRun(runId) : await owner.unpublishRun(runId);
    published[name] = { run_id: runId, public: answer.public, public_bundle_url: answer.public_bundle_url ?? null, verify_url: answer.verify_url, publication: answer.publication };
  }
  writeJSON(join(folder, "published.json"), published);
  console.log(`${publish ? "published" : "withdrawn"} ${Object.keys(published).length} record(s) in ${folder}`);
  if (publish) {
    const shown = published.verdict || published.receipt || published[Object.keys(published).sort()[0]];
    console.log(`public link: ${shown.verify_url}`);
  }
  console.log("each decision is an event in the account's feed; after `feed`, verify: verdifax-verify --signals signals.json");
}

async function cmdFeed() {
  const keys = loadKeys();
  const feed = new VerdifaxClient({ baseURL: baseURL(), apiKey: keys.feed.secret });
  const events = [];
  let after = 0;
  for (;;) {
    const page = await feed.getSignals({ after, limit: 500 });
    events.push(...page.events);
    if (page.events.length < 500) break;
    after = page.next_after;
  }
  writeJSON("signals.json", events);
  const kinds = {};
  for (const e of events) kinds[e.data.kind] = (kinds[e.data.kind] || 0) + 1;
  console.log(`${events.length} events in the feed: ${Object.keys(kinds).sort().map((k) => `${k} ${kinds[k]}`).join(", ")}`);
  console.log("verify: verdifax-verify --signals signals.json");
}

// ── verify commands ────────────────────────────────────────────────────────

function printVerifyCommands(jobId) {
  const folder = join("records", jobId);
  console.log(`verify with verdifax-verify ${MIN_VERIFIER} or later:`);
  console.log(`  verdifax-verify --job ${join(folder, "steps")}`);
  console.log(`  verdifax-verify --account-policy ${join(REPO, "policy", "claims-review-v1.json")} ${join(folder, "steps", "verdict.json")}`);
  if (existsSync(join(folder, "signoffs.json"))) {
    console.log(`  verdifax-verify --signoff ${join(folder, "signoffs.json")} --approver-key approver.txt ${join(folder, "steps", "verdict.json")}`);
  }
  if (existsSync("signals.json")) console.log("  verdifax-verify --signals signals.json");
}

// ── main ───────────────────────────────────────────────────────────────────

function option(args, name) {
  const i = args.indexOf(name);
  if (i < 0) return undefined;
  const value = args[i + 1];
  if (value === undefined) fail(`${name} needs a value`);
  args.splice(i, 2);
  return value;
}

async function main() {
  const args = process.argv.slice(2);
  const command = args.shift();
  switch (command) {
    case "setup":
      return cmdSetup();
    case "review": {
      const model = option(args, "--model") || DEFAULT_MODEL;
      const answers = option(args, "--answers");
      if (args.length !== 1) fail("review takes one draft file");
      return cmdReview(args[0], model, answers);
    }
    case "sign-off": {
      const note = option(args, "--note");
      const receipt = option(args, "--receipt");
      const decision = args.pop();
      const jobId = args.shift();
      if (!["approve", "reject"].includes(decision)) fail("sign-off needs approve or reject");
      if (!receipt && !jobId) fail("sign-off needs a JOB_ID, or --receipt RUN_ID");
      return cmdSignOff(jobId, decision, note, receipt);
    }
    case "publish":
    case "unpublish": {
      const receipt = option(args, "--receipt");
      const jobId = args.shift();
      if (!receipt && !jobId) fail(`${command} needs a JOB_ID, or --receipt RUN_ID`);
      return publication(jobId, receipt, command === "publish");
    }
    case "feed":
      return cmdFeed();
    case "verify-commands":
      if (args.length !== 1) fail("verify-commands takes a JOB_ID");
      return printVerifyCommands(args[0]);
    default:
      fail("commands: setup, review DRAFT.md, sign-off JOB_ID approve|reject, sign-off --receipt RUN_ID reject, publish JOB_ID, publish --receipt RUN_ID, unpublish JOB_ID, feed, verify-commands JOB_ID");
  }
}

export { canonical, draftClassification, loadRules, sha256Hex };

if (process.argv[1] && resolve(process.argv[1]) === fileURLToPath(import.meta.url)) {
  main().catch((err) => {
    console.error("claims-review:", err instanceof Error ? err.message : err);
    process.exit(1);
  });
}
