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
// A decision means the same on every result: approve agrees with
// the result signed, a verdict or the stop a receipt records, and reject
// judges it wrong. A sign-off releases or blocks nothing on its own.
//
// Commands (the same in every language of this repository):
//
//   claims-review setup                              owner key in VERDIFAX_OWNER_KEY
//   claims-review review DRAFT.md [--model ID]        model key in ANTHROPIC_API_KEY
//   claims-review review DRAFT.md --without-actor     the gate step names no agent
//   claims-review sign-off JOB_ID approve|reject [--note TEXT] [--approver NAME]
//   claims-review sign-off --receipt RUN_ID approve|reject [--note TEXT] [--approver NAME]
//   claims-review publish JOB_ID | publish --receipt RUN_ID
//   claims-review unpublish JOB_ID | unpublish --receipt RUN_ID
//   claims-review approver-add NAME | approver-revoke NAME
//   claims-review settings owner-replacement|verdifax-deletion allowed|blocked
//   claims-review request-key-rotate | request-key-clear      the agent key's request key
//   claims-review key-revoke                                  the owner revokes the agent key
//   claims-review delete-step JOB_ID STEP                     the owner deletes one step
//   claims-review webhook-subscribe URL | webhook-list | webhook-remove
//   claims-review feed
//   claims-review verify-commands JOB_ID
//
// Environment: VERDIFAX_BASE_URL (default https://api.verdifax.com).
// Files, in the working folder: demo-keys.json and approver-key.json (0600),
// approver.txt, account-history.json, records/<job_id>/*.json with payloads/,
// signals.json; approver-<name>-key.json (0600), approver-<name>.txt and
// approvers.json; account-keys.json; webhook.json (0600), webhooks.json and
// webhook-deliveries.json.

import { createHash } from "node:crypto";
import { existsSync, mkdirSync, readFileSync, realpathSync, unlinkSync, writeFileSync } from "node:fs";
import { dirname, join, resolve } from "node:path";
import { fileURLToPath } from "node:url";
import { APIError, VerdifaxClient, generateRequestKeyPair } from "verdifax";

const DEFAULT_BASE_URL = "https://api.verdifax.com";
const DEFAULT_MODEL = "claude-fable-5-1";
const FORM = "verdifax.claims-review.v1";
const RULE_SET = "v3";
const ACTOR_ID = "claims-review-agent";
const ROUTE_INTAKE = "claims-review/v1/intake";
const ROUTE_CLAIMS = "claims-review/v1/claims";
const ROUTE_RULE = "claims-review/v1/rule-check";
const ROUTE_VERDICT = "claims-review/v1/verdict";
const OBJECT_PREFIX = "claims-review/v1/";
const LEVELS = ["PUBLIC", "INTERNAL", "CONFIDENTIAL", "RESTRICTED", "TOP_SECRET", "TS_SCI"];
const MIN_VERIFIER = "v0.34.0";
// Anthropic's SDKs retry a dropped connection, 408, 409, 429 and any 5xx, but
// only twice by default, waiting half a second then one second. That is too
// short: on November 13, 2026 a 503 from the model API outlasted it and ended
// this program after twelve sealed steps, failing a readiness run.
// MODEL_RETRY_TRIES sets the number of tries for one model call; the SDK's
// wait grows to a cap of eight seconds, so an outage of about two minutes is
// waited out. A retry can bill a second model call when the first answer was
// lost on the way back; each step's measured cost shows it. All four languages
// hold the same number (demo-parity.json, R259).
const MODEL_RETRY_TRIES = 21;
// A stopped step's records live in records/stopped-<run_id>/, and the
// model's exact answers in answers/ beside each job's steps/.
const STOPPED_FOLDER_PREFIX = "stopped-";
const ANSWERS_FOLDER = "answers";
// The text the server's AI review checked for each step, as the server
// returned it once: a reader re-runs the review's check from these files.
const TRACES_FOLDER = "traces";
// The account policy setup activates; v1 to v3 stay for records sealed
// under them. v3 requires a reviewed AI answer on every step after the gate;
// v4 adds that every step names its agent (required_attested_fields), that
// each verdict needs a person's sign-off, and that every step after the gate
// passes its declared data checks.
const POLICY_FILE = "claims-review-v4.json";
// The rule set's pinned hash, which the rule_set_pinned data check compares.
const RULE_SET_SHA256 = "a2afb580a203ccbd084645cfbbce532dd7eaf1d2dbeaa49fd008f66a91f98ae7";
// The data checks every step declares in policy_context.data_quality, each
// computed: draft_unchanged (the draft re-read before the step has the hash
// sealed at intake) and rule_set_pinned (the rule set's hash is the pin).
const DATA_QUALITY_CHECKS = ["draft_unchanged", "rule_set_pinned"];
const DATA_QUALITY_SOURCE = "claims-review draft and rule set";
// The owner's two blocks on Verdifax, and their values.
const ACCOUNT_SETTINGS = ["owner-replacement", "verdifax-deletion"];
const SETTING_VALUES = ["allowed", "blocked"];
const APPROVER_NAME = /^[a-z0-9-]{1,40}$/;
// The AI's work on the steps after the gate: listing the draft's claims, and
// explaining the verdict code decided.
const CLAIMS_PROMPT = "claims-prompt-v1.txt";
const VERDICT_PROMPT = "verdict-prompt-v1.txt";
// A step whose record was deleted is written as its tombstone,
// steps/<step>.deleted.json: the run.deleted entry from the account feed,
// which verdifax-verify --job reads in the step's place (R250).
const DELETED_STEP_SUFFIX = ".deleted.json";
// A saved sign-off history or feed is checked as of when it was fetched, so
// every printed re-check states the limit on its head statement's age (R246).
const SAVED_HEAD_MAX_AGE = "168h";
const HISTORY_AGE_REASON = "(168h: a saved history is checked as of when it was fetched; raise the limit for an older copy, or run sign-off again to fetch it fresh)";
// Printed under the account-policy check of a verdict with no sign-off yet:
// policy v4 requires one, so the verifier exits 4 until it is given.
const SIGNOFF_PENDING_REASON = "(exit 4, ANCHORED, SIGN-OFF NOT SHOWN, until a sign-off is given: the policy requires one on the verdict; run sign-off, then verify-commands again)";
const FEED_AGE_REASON = "(168h: a saved feed is checked as of when it was fetched; raise the limit for an older copy, or run feed again to fetch it fresh)";
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
  const doc = JSON.parse(readFileSync(join(REPO, "rules", `${RULE_SET}.json`), "utf8"));
  return { rules: doc.rules, ruleSetHash: sha256Hex(canonical(doc)) };
}

const loadPrompt = (name = "prompt-v1.txt") => readFileSync(join(REPO, "rules", name), "utf8");

// A file's exact bytes as UTF-8 text, so a sealed draft or policy hash
// matches the file (R170).
function readInput(path) {
  return readFileSync(path).toString("utf8");
}

function writePrivate(path, text) {
  writeFileSync(path, text, { mode: 0o600, flag: "w" });
}

function writeJSON(path, value) {
  mkdirSync(dirname(path), { recursive: true });
  writeFileSync(path, exactText(value) + "\n");
}

// A value the SDK read from the server keeps the server's exact text in
// sealedJson (SDK 1.11.0), so a saved record is the record the server sent,
// every digit included; a list of such records is joined from their texts.
// Anything this program builds itself is written as JSON.
function exactText(value) {
  if (value && typeof value.sealedJson === "string") return value.sealedJson;
  if (Array.isArray(value) && value.length > 0 && value.every((v) => v && typeof v.sealedJson === "string")) {
    return "[" + value.map((v) => v.sealedJson).join(",") + "]";
  }
  return JSON.stringify(value, null, 2);
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

// A client for a key; one whose request key was cleared sends unsigned.
const clientFor = (key) =>
  new VerdifaxClient({
    baseURL: baseURL(), apiKey: key.secret, ...(key.signing_key ? { signingKey: Buffer.from(key.signing_key, "hex") } : {}),
  });

// The request_substance a step declares: the model's own usage report for the
// step's model call, so the server's budgets count real tokens. The gate step
// calls no model and declares none (null).
function usageSubstance(usage) {
  if (usage === null || usage === undefined) return null;
  return {
    cost: { input_tokens: usage.input_tokens, output_tokens: usage.output_tokens, total_tokens: usage.input_tokens + usage.output_tokens },
  };
}

// The data checks a step declares, each computed (round 62).
function dataQuality(intakeDraftSha256, draftSha256Now, ruleSetHash, asOf) {
  const result = (ok) => (ok ? "pass" : "fail");
  return {
    source: DATA_QUALITY_SOURCE, as_of: asOf, checks: [
      { name: DATA_QUALITY_CHECKS[0], result: result(draftSha256Now === intakeDraftSha256) },
      { name: DATA_QUALITY_CHECKS[1], result: result(ruleSetHash === RULE_SET_SHA256) },
    ],
  };
}

// Why a settings command is refused before anything is sent, or null.
function settingError(setting, value) {
  if (!ACCOUNT_SETTINGS.includes(setting)) return `the setting must be ${ACCOUNT_SETTINGS.join(" or ")}, not ${setting}`;
  if (!SETTING_VALUES.includes(value)) return `the value must be ${SETTING_VALUES.join(" or ")}, not ${value}`;
  return null;
}

const approverNameError = (name) => (APPROVER_NAME.test(name) ? null : "an approver name is 1 to 40 lowercase letters, digits or dashes");

// A named approver's private key file (0600) and public key file.
const approverFiles = (name) => [`approver-${name}-key.json`, `approver-${name}.txt`];

// The keys file with the agent key's request key replaced, or removed when it
// was cleared.
function withAgentSigningKey(keys, signingKeyHex) {
  const agent = { ...keys.agent };
  delete agent.signing_key;
  if (signingKeyHex !== null) agent.signing_key = signingKeyHex;
  return { ...keys, agent };
}

// ── setup ──────────────────────────────────────────────────────────────────

async function cmdSetup() {
  if (existsSync("demo-keys.json")) fail("demo-keys.json exists here; this folder is already set up, and every program in this repository uses it");
  const ownerSecret = process.env.VERDIFAX_OWNER_KEY || "";
  if (!ownerSecret) fail("set VERDIFAX_OWNER_KEY to the account's owner key");
  const ownerSigner = generateRequestKeyPair();
  const owner = new VerdifaxClient({ baseURL: baseURL(), apiKey: ownerSecret, signingKey: ownerSigner.privateKey });
  await owner.registerAccountRequestKey();
  const member = await owner.createAccountKey("claims-review-agent");
  const memberSigner = generateRequestKeyPair();
  const agent = new VerdifaxClient({ baseURL: baseURL(), apiKey: member.secret, signingKey: memberSigner.privateKey });
  await agent.registerRequestKey();
  const policyText = readInput(join(REPO, "policy", POLICY_FILE));
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
  const client = new Anthropic({ maxRetries: MODEL_RETRY_TRIES - 1 });
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

// askStructured asks the model for one structured answer. recordedKey names
// the recorded answer for tests; makeShape builds the zod object it parses;
// declined is the text kept as the AI output when the model declines, and
// empty the fields then.
async function askStructured(recorded, recordedKey, prompt, userText, model, maxTokens, makeShape, declined, empty) {
  if (recorded) {
    const a = recorded[recordedKey];
    const { model: m, response_id: r, ...fields } = a;
    return { model: m || "recorded-answers", responseId: r || "recorded", fields, raw: canonical(fields), inputTokens: 0, outputTokens: 0 };
  }
  const { default: Anthropic } = await import("@anthropic-ai/sdk");
  const { zodOutputFormat } = await import("@anthropic-ai/sdk/helpers/zod");
  const { z } = await import("zod");
  const client = new Anthropic({ maxRetries: MODEL_RETRY_TRIES - 1 });
  const response = await client.messages.parse({
    model,
    max_tokens: maxTokens,
    system: prompt,
    messages: [{ role: "user", content: userText }],
    output_config: { format: zodOutputFormat(makeShape(z)) },
  });
  const usage = { inputTokens: response.usage.input_tokens, outputTokens: response.usage.output_tokens };
  const raw = (response.content.find((b) => b.type === "text") || {}).text || "";
  if (response.stop_reason === "refusal" || !response.parsed_output) {
    return { model: response.model, responseId: response.id, fields: empty, raw: raw || declined, ...usage };
  }
  return { model: response.model, responseId: response.id, fields: response.parsed_output, raw, ...usage };
}

// The draft's claims, as the model listed them; the exact answer is the
// step's AI output, reviewed and sealed.
const askClaims = (recorded, prompt, userText, model) =>
  askStructured(recorded, "claims", prompt, userText, model, 4000, (z) => z.object({ claims: z.array(z.string()) }),
    "the model declined to list the claims", { claims: [] });

// The model's plain-words explanation of the verdict code decided.
const askExplanation = (recorded, prompt, userText, model) =>
  askStructured(recorded, "verdict", prompt, userText, model, 2000, (z) => z.object({ explanation: z.string() }),
    "the model declined to explain the verdict", { explanation: "" });

function measuredCost(model, inputTokens, outputTokens) {
  const prices = JSON.parse(readFileSync(join(REPO, "rules", "model-prices-v1.json"), "utf8")).prices;
  if (!prices[model]) return null;
  return (inputTokens / 1e6) * prices[model].input + (outputTokens / 1e6) * prices[model].output;
}

// ── review ─────────────────────────────────────────────────────────────────

function utcStamp() {
  return new Date().toISOString().replace(/[-:]/g, "").replace(/\.\d{3}Z$/, "Z");
}

async function cmdReview(draftPath, model, answersPath, withoutActor) {
  const keys = loadKeys();
  const agent = clientFor(keys.agent);
  const { rules, ruleSetHash } = loadRules();
  const prompt = loadPrompt();
  const draft = readInput(draftPath);
  const draftHash = sha256Hex(draft);
  const level = draftClassification(draft);
  // The time the draft was read: the as_of of every step's data checks.
  const readAt = new Date().toISOString().replace(/\.\d{3}Z$/, "Z");
  const recorded = loadAnswers(answersPath);
  const now = utcStamp();
  const jobId = `claims-review-${draftHash.slice(0, 12)}-${now}`;
  const out = join("records", jobId);
  const payloads = {};
  // The step's policy context: its object, the draft's classification and the
  // data checks, computed again now, the draft re-read.
  const policy = (objectId) => ({
    object_type: "text", object_id: OBJECT_PREFIX + objectId, classification: level,
    data_quality: dataQuality(draftHash, sha256Hex(readInput(draftPath)), ruleSetHash, readAt),
  });

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
      // The gate step names its agent, as policy v4 requires of every step;
      // --without-actor leaves the name out on purpose, so the policy
      // refuses the draft here, before any model is called.
      ...(withoutActor ? {} : { attested_context: { attested: true, actor_id: ACTOR_ID } }),
    });
  } catch (err) {
    if (!(err instanceof APIError)) throw err;
    // The account policy refused the draft at intake (for example a
    // confidential draft), so it never reached a model. The refusal is
    // sealed as a receipt; nothing else runs for this draft.
    await stop(agent, jobId, "intake", {}, { intake: intakePayload }, {}, draftPath, draftHash, err, model, { input: 0, output: 0 }, new Set(), {});
    return;
  }
  payloads.intake = intakePayload;
  const runIds = { intake: intake.run_id };

  const results = [];
  const tokens = { input: 0, output: 0 };
  const models = new Set();
  // The model's exact answer, the text Verdifax's AI review reads, kept as
  // sent so a reader recomputes the sealed ai_output_hash from the file.
  const answerTexts = {};
  // The text the server's AI review checked, returned once with each step.
  const traceTexts = {};

  // The gate passed, so the draft may reach a model. Every step from here
  // carries the model's answer, which Verdifax's AI review reads and the
  // record seals; the account policy refuses any step without one.
  const claimsPrompt = loadPrompt(CLAIMS_PROMPT);
  const claimsText = `Draft:\n${draft}`;
  const claimsPromptHash = sha256Hex(claimsPrompt + "\n" + claimsText);
  const listed = await askClaims(recorded, claimsPrompt, claimsText, model);
  tokens.input += listed.inputTokens;
  tokens.output += listed.outputTokens;
  models.add(listed.model);
  answerTexts.claims = listed.raw;
  const claimsPayload = {
    form: FORM, step: "claims", draft_sha256: draftHash, rule_set_hash: ruleSetHash, model: listed.model,
    model_response_id: listed.responseId, prompt_hash: claimsPromptHash, claims: listed.fields.claims,
  };
  let claimsStep;
  try {
    claimsStep = await job.step("claims", {
      after: ["intake"], payload_text: canonical(claimsPayload), program_id: PROGRAM_ID, route_id: ROUTE_CLAIMS,
      registry_record_hash: REGISTRY_RECORD_HASH, ai_output_text: answerTexts.claims,
      request_substance: usageSubstance({ input_tokens: listed.inputTokens, output_tokens: listed.outputTokens }),
      attested_context: {
        attested: true, actor_id: ACTOR_ID, model_provider: "anthropic", model_name: listed.model, prompt_hash: claimsPromptHash,
        decision_kind: "claims.list", decision_result: String(listed.fields.claims.length),
      },
      policy_context: policy("claims"),
    });
  } catch (err) {
    if (!(err instanceof APIError)) throw err;
    payloads.claims = claimsPayload;
    if (err.aivpTraceText) traceTexts.claims = err.aivpTraceText;
    await stop(agent, jobId, "claims", runIds, payloads, answerTexts, draftPath, draftHash, err, model, tokens, models, traceTexts);
    return;
  }
  if (claimsStep.aivp_trace_text) traceTexts.claims = claimsStep.aivp_trace_text;
  payloads.claims = claimsPayload;
  runIds.claims = claimsStep.run_id;
  console.log(`claims: ${listed.fields.claims.length} listed`);

  for (const rule of rules) {
    const userText = `Rule ${rule.id}: ${rule.text}\n\nDraft:\n${draft}`;
    const promptHash = sha256Hex(prompt + "\n" + userText);
    const answer = await ask(recorded, prompt, userText, rule.id, model);
    tokens.input += answer.inputTokens;
    tokens.output += answer.outputTokens;
    models.add(answer.model);
    answerTexts[rule.id] = answer.raw || answer.reason;
    const payload = {
      form: FORM, step: rule.id, rule_id: rule.id, rule_text: rule.text, draft_sha256: draftHash, rule_set_hash: ruleSetHash,
      model: answer.model, model_response_id: answer.responseId, prompt_hash: promptHash, verdict: answer.verdict, reason: answer.reason,
    };
    let receipt;
    try {
      receipt = await job.step(rule.id, {
        after: ["claims"], payload_text: canonical(payload), program_id: PROGRAM_ID, route_id: ROUTE_RULE,
        registry_record_hash: REGISTRY_RECORD_HASH, ai_output_text: answerTexts[rule.id],
        request_substance: usageSubstance({ input_tokens: answer.inputTokens, output_tokens: answer.outputTokens }),
        attested_context: {
          attested: true, actor_id: ACTOR_ID, model_provider: "anthropic", model_name: answer.model, prompt_hash: promptHash,
          decision_kind: "claims.rule_check", decision_result: answer.verdict,
        },
        policy_context: policy(rule.id),
      });
    } catch (err) {
      if (!(err instanceof APIError)) throw err;
      // The account policy, the server's content screen or its review of
      // the model's answer stopped this step, so the job stops here: the
      // stop is sealed as a receipt, the steps already sealed are kept
      // beside it, and a person sees the stop. The job never gets a verdict.
      payloads[rule.id] = payload;
      if (err.aivpTraceText) traceTexts[rule.id] = err.aivpTraceText;
      await stop(agent, jobId, rule.id, runIds, payloads, answerTexts, draftPath, draftHash, err, model, tokens, models, traceTexts);
      return;
    }
    if (receipt.aivp_trace_text) traceTexts[rule.id] = receipt.aivp_trace_text;
    payloads[rule.id] = payload;
    runIds[rule.id] = receipt.run_id;
    results.push({ rule_id: rule.id, verdict: answer.verdict, reason: answer.reason });
    console.log(`${rule.id} ${answer.verdict}: ${answer.reason}`);
  }

  // The program decides the verdict from the rule answers, so anyone can
  // decide it again from the sealed records; the model only explains it.
  const failed = results.filter((r) => r.verdict !== "pass").map((r) => r.rule_id);
  const verdict = failed.length === 0 ? "pass" : "fail";
  const verdictPrompt = loadPrompt(VERDICT_PROMPT);
  const verdictText = `Verdict: ${verdict}\n\nRule results:\n${canonical(results)}`;
  const verdictPromptHash = sha256Hex(verdictPrompt + "\n" + verdictText);
  const explained = await askExplanation(recorded, verdictPrompt, verdictText, model);
  tokens.input += explained.inputTokens;
  tokens.output += explained.outputTokens;
  models.add(explained.model);
  answerTexts.verdict = explained.raw;
  const verdictPayload = {
    form: FORM, step: "verdict", draft_sha256: draftHash, rule_set_hash: ruleSetHash, verdict, failed_rules: failed, results,
    model: explained.model, model_response_id: explained.responseId, prompt_hash: verdictPromptHash,
    explanation: explained.fields.explanation,
  };
  let final;
  try {
    final = await job.step("verdict", {
      after: rules.map((r) => r.id), final: true, payload_text: canonical(verdictPayload), program_id: PROGRAM_ID,
      route_id: ROUTE_VERDICT, registry_record_hash: REGISTRY_RECORD_HASH, ai_output_text: answerTexts.verdict,
      request_substance: usageSubstance({ input_tokens: explained.inputTokens, output_tokens: explained.outputTokens }),
      attested_context: {
        attested: true, actor_id: ACTOR_ID, model_provider: "anthropic", model_name: explained.model, prompt_hash: verdictPromptHash,
        decision_kind: "claims.verdict", decision_result: verdict,
      },
      policy_context: policy("verdict"),
    });
  } catch (err) {
    if (!(err instanceof APIError)) throw err;
    payloads.verdict = verdictPayload;
    if (err.aivpTraceText) traceTexts.verdict = err.aivpTraceText;
    await stop(agent, jobId, "verdict", runIds, payloads, answerTexts, draftPath, draftHash, err, model, tokens, models, traceTexts);
    return;
  }
  if (final.aivp_trace_text) traceTexts.verdict = final.aivp_trace_text;
  payloads.verdict = verdictPayload;
  runIds.verdict = final.run_id;
  console.log(`verdict explained: ${explained.fields.explanation}`);

  await saveJobRecords(agent, jobId, out, runIds, payloads);
  saveAnswers(out, answerTexts);
  saveTexts(join(out, TRACES_FOLDER), traceTexts);
  const cost = models.size === 1 && models.has(model) ? measuredCost(model, tokens.input, tokens.output) : null;
  const summary = {
    job_id: jobId, draft: draftPath, draft_sha256: draftHash, rule_set_hash: ruleSetHash, verdict,
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

// A step the server stopped is sealed as one run whose status names the kind
// of stop, with a receipt of that kind (R152). Any other error is not a stop:
// the program reports it as the server gave it and seals nothing for it.
const STOP_KINDS = {
  pepg_deny: "the account policy or the content screen",
  aivp_t4_halt: "the server's review of the model's answer",
  ccv_halt: "the per-run budget",
  macc_halt: "the cumulative budget",
  attestation_gap: "the policy's required attested fields",
};
const STOP_RECEIPT_FIELDS = {
  pepg_deny: "deny_receipt", aivp_t4_halt: "aivp_t4_halt_receipt", ccv_halt: "ccv_halt_receipt",
  macc_halt: "macc_halt_receipt", attestation_gap: "attestation_gap_receipt",
};

async function stopReceipt(agent, status, runId) {
  const fetch = {
    pepg_deny: (id) => agent.getDenyReceipt(id), aivp_t4_halt: (id) => agent.getAivpT4HaltReceipt(id),
    ccv_halt: (id) => agent.getCCVHaltReceipt(id), macc_halt: (id) => agent.getMACCHaltReceipt(id),
    attestation_gap: (id) => agent.getAttestationGapReceipt(id),
  }[status];
  return fetch(runId);
}

// The sealed hash of a stopped run's receipt, whatever kind of stop it was.
async function stopReceiptHash(agent, runId) {
  const { status } = await agent.getRun(runId);
  if (!(status in STOP_KINDS)) fail(`run ${runId} is ${status}, not a stopped run with a receipt`);
  return (await stopReceipt(agent, status, runId))[STOP_RECEIPT_FIELDS[status]].hash;
}

// The run the server sealed for this job's stopped step, found by its job
// and step, never by being the newest of the key.
async function stoppedRun(agent, jobId, step, err) {
  const rows = (await agent.exportRuns({ job_id: jobId })).runs.filter((r) => r.step_id === step && r.status in STOP_KINDS);
  if (rows.length === 0) {
    fail(`the server answered HTTP ${err.status} at ${step}: ${err.message}; it recorded no stop for this step, so nothing was sealed as one`);
  }
  return rows.reduce((a, b) => (b.run_id > a.run_id ? b : a));
}

// Record a stopped step: its sealed receipt, every record the job sealed
// before it, the exact payload bytes of every step including the stopped
// one, and a summary naming the step and the kind of stop. The same for
// intake and for any later step.
// The model's exact answers, one file per step, so each sealed
// ai_output_hash recomputes from its file.
// Writes each step's text to <folder>/<step>.txt, byte for byte.
function saveTexts(folder, texts) {
  if (Object.keys(texts).length === 0) return;
  mkdirSync(folder, { recursive: true });
  for (const [step, text] of Object.entries(texts)) writeFileSync(join(folder, `${step}.txt`), Buffer.from(text, "utf8"));
}

function saveAnswers(folder, answerTexts) {
  mkdirSync(join(folder, ANSWERS_FOLDER), { recursive: true });
  for (const [step, text] of Object.entries(answerTexts)) writeFileSync(join(folder, ANSWERS_FOLDER, `${step}.txt`), Buffer.from(text, "utf8"));
}

// The folder of the stopped step whose receipt is run runId.
function stoppedFolder(runId) {
  return join("records", `${STOPPED_FOLDER_PREFIX}${runId}`);
}

async function stop(agent, jobId, step, runIds, payloads, answerTexts, draftPath, draftHash, err, model, tokens, models, traceTexts) {
  const row = await stoppedRun(agent, jobId, step, err);
  const stopped = row.run_id;
  const folder = stoppedFolder(stopped); // named by the run, so two stops never share a folder
  writeJSON(join(folder, "receipt.json"), await stopReceipt(agent, row.status, stopped));
  mkdirSync(join(folder, "steps"), { recursive: true });
  // A sealed step whose record was deleted since is kept as its tombstone.
  const deleted = Object.keys(runIds).length > 0 ? deletedSteps(await readAccountFeed(loadKeys()), jobId) : {};
  writeDeletedSteps(folder, Object.fromEntries(Object.entries(deleted).filter(([s]) => s in runIds)));
  for (const [sealedStep, runId] of Object.entries(runIds)) {
    if (!(sealedStep in deleted)) writeFileSync(join(folder, "steps", `${sealedStep}.json`), await agent.getRunArtifacts(runId));
  }
  mkdirSync(join(folder, "payloads"), { recursive: true });
  for (const [name, payload] of Object.entries(payloads)) writeFileSync(join(folder, "payloads", `${name}.json`), canonical(payload));
  saveAnswers(folder, answerTexts);
  saveTexts(join(folder, TRACES_FOLDER), traceTexts);
  const cost = models.size === 1 && models.has(model) ? measuredCost(model, tokens.input, tokens.output) : null;
  writeJSON(join(folder, "job.json"), {
    stopped_run_id: stopped, step, stop_status: row.status, job_id: jobId, draft: draftPath, draft_sha256: draftHash,
    reason: err.message, run_ids: runIds, models: [...models].sort(), measured_tokens: tokens, measured_cost_usd: cost,
  });
  console.log(`stopped at ${step} by ${STOP_KINDS[row.status]}: run ${stopped}, ${err.message}`);
  console.log(`records in ${folder}`);
  if (cost !== null) {
    console.log(`measured model use before the stop: ${tokens.input} input and ${tokens.output} output tokens, about $${cost.toFixed(2)} at published prices`);
  }
}

async function saveJobRecords(agent, jobId, out, runIds, payloads) {
  // A deleted step's record is gone from the export; its tombstone in the
  // account feed stands in its place (R250).
  const deleted = Object.fromEntries(Object.entries(deletedSteps(await readAccountFeed(loadKeys()), jobId)).filter(([s]) => s in runIds));
  const rows = (await agent.exportRuns({ job_id: jobId })).runs.filter((r) => !(r.step_id in deleted));
  if (rows.length + Object.keys(deleted).length !== Object.keys(runIds).length) {
    fail(`the export lists ${rows.length} records and the feed ${Object.keys(deleted).length} deleted steps for ${jobId}, the program sealed ${Object.keys(runIds).length}`);
  }
  // The sealed step records alone sit in steps/, the folder verdifax-verify
  // --job reads: it takes every .json file there as a record, so the
  // summary, the sign-offs and the payloads live beside that folder.
  mkdirSync(join(out, "steps"), { recursive: true });
  mkdirSync(join(out, "payloads"), { recursive: true });
  writeDeletedSteps(out, deleted);
  for (const row of rows) {
    if (row.status in STOP_KINDS) writeJSON(join(out, "steps", `${row.step_id}.json`), await stopReceipt(agent, row.status, row.run_id));
    else writeFileSync(join(out, "steps", `${row.step_id}.json`), await agent.getRunArtifacts(row.run_id));
  }
  // The exact bytes each step sent, so a reader recomputes the sealed
  // payload hash from the file.
  for (const [step, payload] of Object.entries(payloads)) writeFileSync(join(out, "payloads", `${step}.json`), canonical(payload));
}

// The steps of job jobId whose records were deleted, each with its
// tombstone: the entry of every run.deleted event naming the job, by its
// step. A run.sealed event whose record was deleted carries the tombstone
// in data.deleted with a null entry; it is not the deletion and is passed
// over, as is a deletion of another job or of no job.
function deletedSteps(events, jobId) {
  const out = {};
  for (const e of events) {
    const entry = e?.data?.entry;
    if (e?.data?.kind !== "run.deleted" || !entry || entry.job_id !== jobId || !entry.step_id) continue;
    out[entry.step_id] = entry;
  }
  return out;
}

// Writes each deleted step's tombstone as <folder>/steps/<step>.deleted.json,
// the server's exact text, beside the step records verdifax-verify --job reads.
function writeDeletedSteps(folder, deleted) {
  for (const [step, entry] of Object.entries(deleted)) writeJSON(join(folder, "steps", `${step}${DELETED_STEP_SUFFIX}`), entry);
}

// The account's whole feed, read with the feed key, page by page.
async function readAccountFeed(keys) {
  const feed = new VerdifaxClient({ baseURL: baseURL(), apiKey: keys.feed.secret });
  const events = [];
  let after = 0;
  for (;;) {
    const page = await feed.getSignals({ after, limit: 500 });
    events.push(...page.events);
    if (page.events.length < 500) break;
    after = page.next_after;
  }
  return events;
}

// The re-check of a saved sign-off history or feed, with its head age limit,
// and the reason for the limit on the line after (R246).
const historyCheck = (path) => [`verdifax-verify --signoff-history ${path} --signoff-head-max-age ${SAVED_HEAD_MAX_AGE}`, `  ${HISTORY_AGE_REASON}`];
const feedCheck = () => [`verdifax-verify --signals signals.json --feed-head-max-age ${SAVED_HEAD_MAX_AGE}`, `  ${FEED_AGE_REASON}`];

// ── sign-off ───────────────────────────────────────────────────────────────

// The first approver's private key, or a named approver's.
function approverPrivateKey(name) {
  const path = name ? approverFiles(name)[0] : "approver-key.json";
  if (!existsSync(path)) fail(`no ${path} here; run \`claims-review ${name ? "approver-add " + name : "setup"}\` first`);
  return Buffer.from(JSON.parse(readFileSync(path, "utf8")).private_key, "hex");
}

async function cmdSignOff(jobId, decision, note, receiptRunId, approverName) {
  if (approverName && approverNameError(approverName)) fail(approverNameError(approverName));
  const keys = loadKeys();
  const agent = clientFor(keys.agent);
  const owner = clientFor(keys.owner);
  const key = approverPrivateKey(approverName);
  const approverFile = approverName ? approverFiles(approverName)[1] : "approver.txt";
  let target;
  let folder;
  let runId;
  if (receiptRunId) {
    runId = Number(receiptRunId);
    target = { kind: "receipt", run_id: runId, digest: await stopReceiptHash(agent, runId) };
    folder = stoppedFolder(runId);
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
  console.log(`verify: verdifax-verify --signoff ${join(folder, "signoffs.json")} --approver-key ${approverFile} ${recordFile}`);
  const [history, historyReason] = historyCheck(join(folder, "signoff-history.json"));
  console.log(`verify: ${history}`);
  console.log(historyReason);
}

// ── the account: approvers, settings, keys, deletion, webhooks ─────────────
// Each change is made with the owner key (the agent key for its own request
// key) and sealed by the server in the account's feed, which `feed` saves:
// the feed is the record of every change, so no command keeps its own.

const ownerClient = () => clientFor(loadKeys().owner);

// The id the server gave this approver when it was added, kept with its key.
function approverId(name) {
  const [privatePath] = approverFiles(name);
  if (!existsSync(privatePath)) fail(`no ${privatePath} here; run \`claims-review approver-add ${name}\` first`);
  return Number(JSON.parse(readFileSync(privatePath, "utf8")).id);
}

async function cmdApproverAdd(name) {
  if (approverNameError(name)) fail(approverNameError(name));
  const [privatePath, publicPath] = approverFiles(name);
  if (existsSync(privatePath)) fail(`${privatePath} exists here; this approver is already added`);
  const owner = ownerClient();
  const approver = generateRequestKeyPair();
  const answer = await owner.registerApprover(`Claims review approver ${name}`, approver.privateKey);
  writePrivate(privatePath, JSON.stringify({ private_key: approver.privateKey.toString("hex"), id: answer.id }) + "\n");
  writeFileSync(publicPath, approver.publicKey.toString("base64"));
  writeJSON("approvers.json", await owner.listApprovers());
  console.log(`approver ${name} added; its public key is in ${publicPath}`);
}

async function cmdApproverRevoke(name) {
  if (approverNameError(name)) fail(approverNameError(name));
  const owner = ownerClient();
  await owner.revokeApprover(approverId(name));
  writeJSON("approvers.json", await owner.listApprovers());
  console.log(`approver ${name} revoked; the sign-offs it gave stay as they were`);
}

async function cmdSettings(setting, value) {
  if (settingError(setting, value)) fail(settingError(setting, value));
  await ownerClient().setAccountSetting(setting, value);
  console.log(`${setting} is ${value}`);
}

async function cmdRequestKeyRotate() {
  const keys = loadKeys();
  const next = generateRequestKeyPair();
  await clientFor(keys.agent).rotateRequestKey(next.privateKey);
  writePrivate("demo-keys.json", JSON.stringify(withAgentSigningKey(keys, next.privateKey.toString("hex")), null, 2) + "\n");
  console.log("the agent key's request key was replaced, approved by the key it replaced");
}

async function cmdRequestKeyClear() {
  const keys = loadKeys();
  await clientFor(keys.agent).clearRequestKey();
  writePrivate("demo-keys.json", JSON.stringify(withAgentSigningKey(keys, null), null, 2) + "\n");
  console.log("the agent key's request key was cleared, approved by that key");
}

async function cmdKeyRevoke() {
  const keys = loadKeys();
  const owner = clientFor(keys.owner);
  await owner.revokeAccountKey(Number(keys.agent.id));
  writeJSON("account-keys.json", await owner.listAccountKeys());
  console.log(`agent key ${keys.agent.id} revoked; the records it sealed stay as they were`);
}

// A deleted step's record is replaced by its tombstone, which verdifax-verify
// --job reads in its place (R250).
function replaceWithTombstone(folder, step, tombstone) {
  writeDeletedSteps(folder, { [step]: tombstone });
  const record = join(folder, "steps", `${step}.json`);
  if (existsSync(record)) unlinkSync(record);
}

async function cmdDeleteStep(jobId, step) {
  const folder = join("records", jobId);
  const recordPath = join(folder, "steps", `${step}.json`);
  if (!existsSync(recordPath)) fail(`${recordPath} does not exist; delete-step needs the step's sealed record`);
  const record = JSON.parse(readFileSync(recordPath, "utf8"));
  await ownerClient().deleteRun(Number(record.run_id), record.manifest_hash);
  const tombstone = deletedSteps(await readAccountFeed(loadKeys()), jobId)[step];
  if (!tombstone) fail(`the feed holds no run.deleted event for ${step} of ${jobId}`);
  replaceWithTombstone(folder, step, tombstone);
  console.log(`step ${step} of ${jobId} deleted; its tombstone is ${join(folder, "steps", step + DELETED_STEP_SUFFIX)}`);
  console.log(`verify: verdifax-verify --job ${join(folder, "steps")}`);
}

async function cmdWebhookSubscribe(url) {
  const hook = await ownerClient().createWebhook(url);
  writePrivate("webhook.json", JSON.stringify(hook, null, 2) + "\n");
  console.log(`webhook ${hook.id} subscribed to every kind of event at ${url}; its secret is in webhook.json`);
}

// Removes the subscription webhook-subscribe made, as the owner, sealed in the
// feed as webhook.deleted. webhook.json is kept: its secret can be published
// once the subscription is gone, since it then signs nothing.
async function cmdWebhookRemove() {
  if (!existsSync("webhook.json")) fail("no webhook.json here; run `claims-review webhook-subscribe URL` first");
  const id = Number(JSON.parse(readFileSync("webhook.json", "utf8")).id);
  writeJSON("webhook-removed.json", await ownerClient().deleteWebhook(id));
  console.log(`webhook ${id} removed; its removal is sealed in the account's feed`);
}

async function cmdWebhookList() {
  // The lists alone, the same shape in every language: webhooks.json the
  // account's webhooks, webhook-deliveries.json each one's deliveries by id.
  const owner = ownerClient();
  const hooks = (await owner.listWebhooks()).webhooks || [];
  writeJSON("webhooks.json", hooks);
  const deliveries = {};
  for (const h of hooks) deliveries[String(h.id)] = (await owner.listWebhookDeliveries(Number(h.id), { limit: 500 })).deliveries || [];
  writeJSON("webhook-deliveries.json", deliveries);
  console.log(`${hooks.length} webhook(s); their deliveries are in webhook-deliveries.json`);
}

// ── feed ───────────────────────────────────────────────────────────────────

// ── publish ────────────────────────────────────────────────────────────────
// The account's owner key makes a record public on its random link, or
// withdraws it. Each decision is sealed by the server as a run publication
// and placed in the account's feed, which `feed` then reads.

// The folder and the records a publish or unpublish acts on: every step of
// a job, or a stopped step's receipt with every step the job sealed before
// it, so a reader sees what the job did before it stopped.
function publicationTargets(jobId, receipt) {
  if (receipt) {
    const runId = Number(receipt);
    const folder = stoppedFolder(runId);
    const summary = JSON.parse(readFileSync(join(folder, "job.json"), "utf8"));
    return { folder, targets: { ...summary.run_ids, receipt: runId } };
  }
  const folder = join("records", jobId);
  const summary = JSON.parse(readFileSync(join(folder, "job.json"), "utf8"));
  return { folder, targets: { ...summary.run_ids } };
}

// The records a publish or unpublish acts on, sorted: every target but a
// deleted step, which has no record left; its tombstone stands in its place in
// the feed (R250).
function publishable(folder, targets) {
  return Object.keys(targets).filter((n) => !existsSync(join(folder, "steps", `${n}${DELETED_STEP_SUFFIX}`))).sort();
}

async function publication(jobId, receipt, publish) {
  const keys = loadKeys();
  const owner = clientFor(keys.owner);
  const { folder, targets } = publicationTargets(jobId, receipt);
  // A repeated publish changes nothing and the server answers with no new
  // publication, so the sealed publication saved earlier is kept.
  const earlier = existsSync(join(folder, "published.json")) ? JSON.parse(readFileSync(join(folder, "published.json"), "utf8")) : {};
  const published = {};
  for (const name of publishable(folder, targets)) {
    const runId = targets[name];
    const answer = publish ? await owner.publishRun(runId) : await owner.unpublishRun(runId);
    // A completed run is served at its bundle address, a stopped run at its
    // receipt address (R156); each answer names only the one it has.
    let sealed = answer.publication;
    if (sealed === null && earlier[name]?.run_id === runId && earlier[name]?.public === answer.public) sealed = earlier[name].publication ?? null;
    published[name] = {
      run_id: runId, public: answer.public, public_bundle_url: answer.public_bundle_url ?? null,
      public_receipt_url: answer.public_receipt_url ?? null, verify_url: answer.verify_url, publication: sealed,
    };
  }
  writeJSON(join(folder, "published.json"), published);
  console.log(`${publish ? "published" : "withdrawn"} ${Object.keys(published).length} record(s) in ${folder}`);
  if (publish) {
    const shown = published.verdict || published.receipt || published[Object.keys(published).sort()[0]];
    console.log(`public link: ${shown.verify_url}`);
  }
  const [feed, feedReason] = feedCheck();
  console.log(`each decision is an event in the account's feed; after \`feed\`, verify: ${feed}`);
  console.log(feedReason);
}

async function cmdFeed() {
  const events = await readAccountFeed(loadKeys());
  writeJSON("signals.json", events);
  const kinds = {};
  for (const e of events) kinds[e.data.kind] = (kinds[e.data.kind] || 0) + 1;
  console.log(`${events.length} events in the feed: ${Object.keys(kinds).sort().map((k) => `${k} ${kinds[k]}`).join(", ")}`);
  const [feed, feedReason] = feedCheck();
  console.log(`verify: ${feed}`);
  console.log(feedReason);
}

// ── verify commands ────────────────────────────────────────────────────────

function printVerifyCommands(jobId) {
  const folder = join("records", jobId);
  console.log(`verify with verdifax-verify ${MIN_VERIFIER} or later:`);
  const traces = existsSync(join(folder, TRACES_FOLDER)) ? ` --aivp-traces ${join(folder, TRACES_FOLDER)}` : "";
  console.log(`  verdifax-verify --job ${join(folder, "steps")}${traces}`);
  if (existsSync(join(folder, "signoffs.json"))) {
    console.log(`  verdifax-verify --account-policy ${join(REPO, "policy", POLICY_FILE)} --signoff ${join(folder, "signoffs.json")} --approver-key approver.txt ${join(folder, "steps", "verdict.json")}`);
  } else {
    console.log(`  verdifax-verify --account-policy ${join(REPO, "policy", POLICY_FILE)} ${join(folder, "steps", "verdict.json")}`);
    console.log(`  ${SIGNOFF_PENDING_REASON}`);
  }
  if (existsSync(join(folder, "signoffs.json"))) {
    console.log(`  verdifax-verify --signoff ${join(folder, "signoffs.json")} --approver-key approver.txt ${join(folder, "steps", "verdict.json")}`);
  }
  if (existsSync(join(folder, "signoff-history.json"))) {
    const [history, historyReason] = historyCheck(join(folder, "signoff-history.json"));
    console.log(`  ${history}`);
    console.log(historyReason);
  }
  if (existsSync("signals.json")) {
    const [feed, feedReason] = feedCheck();
    console.log(`  ${feed}`);
    console.log(feedReason);
  }
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
      const withoutActor = args.includes("--without-actor");
      if (withoutActor) args.splice(args.indexOf("--without-actor"), 1);
      if (args.length !== 1) fail("review takes one draft file");
      return cmdReview(args[0], model, answers, withoutActor);
    }
    case "sign-off": {
      const note = option(args, "--note");
      const receipt = option(args, "--receipt");
      const approverName = option(args, "--approver");
      const decision = args.pop();
      const jobId = args.shift();
      if (!["approve", "reject"].includes(decision)) fail("sign-off needs approve or reject");
      if (!receipt && !jobId) fail("sign-off needs a JOB_ID, or --receipt RUN_ID");
      return cmdSignOff(jobId, decision, note, receipt, approverName);
    }
    case "publish":
    case "unpublish": {
      const receipt = option(args, "--receipt");
      const jobId = args.shift();
      if (!receipt && !jobId) fail(`${command} needs a JOB_ID, or --receipt RUN_ID`);
      return publication(jobId, receipt, command === "publish");
    }
    case "approver-add":
    case "approver-revoke":
      if (args.length !== 1) fail(`${command} takes a NAME`);
      return command === "approver-add" ? cmdApproverAdd(args[0]) : cmdApproverRevoke(args[0]);
    case "settings":
      if (args.length !== 2) fail("settings takes a SETTING and a VALUE");
      return cmdSettings(args[0], args[1]);
    case "request-key-rotate":
      return cmdRequestKeyRotate();
    case "request-key-clear":
      return cmdRequestKeyClear();
    case "key-revoke":
      return cmdKeyRevoke();
    case "delete-step":
      if (args.length !== 2) fail("delete-step takes a JOB_ID and a STEP");
      return cmdDeleteStep(args[0], args[1]);
    case "webhook-subscribe":
      if (args.length !== 1) fail("webhook-subscribe takes a URL");
      return cmdWebhookSubscribe(args[0]);
    case "webhook-list":
      return cmdWebhookList();
    case "webhook-remove":
      return cmdWebhookRemove();
    case "feed":
      return cmdFeed();
    case "verify-commands":
      if (args.length !== 1) fail("verify-commands takes a JOB_ID");
      return printVerifyCommands(args[0]);
    default:
      fail("commands: setup, review DRAFT.md, sign-off JOB_ID approve|reject, sign-off --receipt RUN_ID approve|reject, publish JOB_ID, publish --receipt RUN_ID, unpublish JOB_ID, unpublish --receipt RUN_ID, approver-add NAME, approver-revoke NAME, settings SETTING VALUE, request-key-rotate, request-key-clear, key-revoke, delete-step JOB_ID STEP, webhook-subscribe URL, webhook-list, webhook-remove, feed, verify-commands JOB_ID");
  }
}

export {
  canonical, DELETED_STEP_SUFFIX, deletedSteps, draftClassification, FEED_AGE_REASON, HISTORY_AGE_REASON, loadRules, MIN_VERIFIER, printVerifyCommands,
  MODEL_RETRY_TRIES, readInput, RULE_SET, SAVED_HEAD_MAX_AGE, sha256Hex, STOP_KINDS, STOP_RECEIPT_FIELDS, stopReceipt, stopReceiptHash, stoppedRun, writeDeletedSteps,
  approverFiles, approverNameError, clientFor, DATA_QUALITY_CHECKS, DATA_QUALITY_SOURCE, dataQuality, POLICY_FILE, REPO, replaceWithTombstone, RULE_SET_SHA256,
  settingError, usageSubstance, withAgentSigningKey, publishable, SIGNOFF_PENDING_REASON,
};

// Node loads the program by its real path, so the path it was started with is
// compared as a real path too: through a linked folder (macOS's /tmp is one)
// the two spellings differ, and the program would otherwise do nothing (R165).
if (process.argv[1] && realpathSync(resolve(process.argv[1])) === realpathSync(fileURLToPath(import.meta.url))) {
  main().catch((err) => {
    console.error("claims-review:", err instanceof Error ? err.message : err);
    process.exit(1);
  });
}
