#!/usr/bin/env python3
"""claims-review: a small agent that checks a draft against 15 general writing
rules, with every step sealed by Verdifax.

A demo, not customer proof. A sealed record shows which rules were checked and
what the model answered; it does not show the model's judgment was right. The
15 rules are general writing rules, not legal advice and not a compliance
standard. The account policy judges what each step declared; it cannot show
the declared facts are true. A sign-off shows a registered approver key signed
that decision; it does not show who holds the key or that the review was
careful. A decision means the same on every result: approve agrees with
the result signed, a verdict or the stop a receipt records, and reject
judges it wrong. A sign-off releases or blocks nothing on its own.

Commands (the same in every language of this repository):

  claims-review setup                              owner key in VERDIFAX_OWNER_KEY
  claims-review review DRAFT.md [--model ID]        model key in ANTHROPIC_API_KEY
  claims-review review DRAFT.md --without-actor     the gate step names no agent
  claims-review sign-off JOB_ID approve|reject [--note TEXT] [--approver NAME]
  claims-review sign-off --receipt RUN_ID approve|reject [--note TEXT] [--approver NAME]
  claims-review publish JOB_ID | publish --receipt RUN_ID
  claims-review unpublish JOB_ID | unpublish --receipt RUN_ID
  claims-review approver-add NAME | approver-revoke NAME
  claims-review settings owner-replacement|verdifax-deletion allowed|blocked
  claims-review request-key-register KEY_FILE               the key Verdifax made for the chapter
  claims-review request-key-rotate | request-key-clear [--holder]   the agent key's request key, or the holder's
  claims-review key-revoke [--holder]                       the owner revokes the agent key, or the holder
  claims-review delete-step JOB_ID STEP                     the owner deletes one step
  claims-review webhook-subscribe URL | webhook-list | webhook-remove
  claims-review feed
  claims-review verify-commands JOB_ID

Environment: VERDIFAX_BASE_URL (default https://api.verdifax.com);
CLAIMS_REVIEW_REGISTRY_RECORD_HASH, the registry record each step claims on a
server that enforces a program registry (default: the program's fixed value).
Files, in the working folder: demo-keys.json and approver-key.json (0600),
approver.txt, account-history.json, records/<job_id>/*.json with payloads/,
signals.json; approver-<name>-key.json (0600), approver-<name>.txt and
approvers.json; account-keys.json; webhook.json (0600),
webhooks.json and webhook-deliveries.json.
"""

from __future__ import annotations

import argparse
import base64
import datetime as dt
import hashlib
import json
import os
import re
import sys
from pathlib import Path
from typing import Any, Optional

import verdifax
from verdifax.exceptions import APIError

DEFAULT_BASE_URL = "https://api.verdifax.com"
DEFAULT_MODEL = "claude-fable-5-1"
FORM = "verdifax.claims-review.v1"
RULE_SET = "v3"
ACTOR_ID = "claims-review-agent"
ROUTE_INTAKE = "claims-review/v1/intake"
ROUTE_CLAIMS = "claims-review/v1/claims"
ROUTE_RULE = "claims-review/v1/rule-check"
ROUTE_VERDICT = "claims-review/v1/verdict"
OBJECT_PREFIX = "claims-review/v1/"
LEVELS = ("PUBLIC", "INTERNAL", "CONFIDENTIAL", "RESTRICTED", "TOP_SECRET", "TS_SCI")
VERDICTS = ("pass", "fail", "unanswered")
MIN_VERIFIER = "v0.36.0"
# Anthropic's SDKs retry a dropped connection, 408, 409, 429 and any 5xx, but
# only twice by default, waiting half a second then one second. That is too
# short: on November 13, 2026 a 503 from the model API outlasted it and ended
# this program after twelve sealed steps, failing a readiness run.
# MODEL_RETRY_TRIES sets the number of tries for one model call; the SDK's wait
# grows to a cap of eight seconds, so an outage of about two minutes is waited
# out. A retry can bill a second model call when the first answer was lost on
# the way back; each step's measured cost shows it. All four languages hold the
# same number (demo-parity.json, R259).
MODEL_RETRY_TRIES = 21
# A stopped step's records live in records/stopped-<run_id>/, and the
# model's exact answers in answers/ beside each job's steps/.
STOPPED_FOLDER_PREFIX = "stopped-"
ANSWERS_FOLDER = "answers"
# The text the server's AI review checked for each step, as the server
# returned it once: a reader re-runs the review's check from these files.
TRACES_FOLDER = "traces"
# The account policy setup activates; v1 to v3 stay for records sealed
# under them. v3 requires a reviewed AI answer on every step after the gate;
# v4 adds that every step names its agent (required_attested_fields), that
# each verdict needs a person's sign-off, and that every step after the gate
# passes its declared data checks.
POLICY_FILE = "claims-review-v4.json"
# The rule set's pinned hash, which the rule_set_pinned data check compares.
RULE_SET_SHA256 = "a2afb580a203ccbd084645cfbbce532dd7eaf1d2dbeaa49fd008f66a91f98ae7"
# The data checks every step declares in policy_context.data_quality, each
# computed: draft_unchanged (the draft re-read before the step has the hash
# sealed at intake) and rule_set_pinned (the rule set's hash is the pin).
DATA_QUALITY_CHECKS = ("draft_unchanged", "rule_set_pinned")
DATA_QUALITY_SOURCE = "claims-review draft and rule set"
# The owner's two blocks on Verdifax, and their values.
ACCOUNT_SETTINGS = ("owner-replacement", "verdifax-deletion")
SETTING_VALUES = ("allowed", "blocked")
APPROVER_NAME = re.compile(r"^[a-z0-9-]{1,40}$")
# The AI's work on the steps after the gate: listing the draft's claims,
# and explaining the verdict code decided.
CLAIMS_PROMPT = "claims-prompt-v1.txt"
VERDICT_PROMPT = "verdict-prompt-v1.txt"
# A step whose record was deleted is written as its tombstone,
# steps/<step>.deleted.json: the run.deleted entry from the account feed,
# which verdifax-verify --job reads in the step's place (R250).
DELETED_STEP_SUFFIX = ".deleted.json"
# A saved sign-off history or feed is checked as of when it was fetched, so
# every printed re-check states the limit on its head statement's age (R246).
SAVED_HEAD_MAX_AGE = "168h"
HISTORY_AGE_REASON = "(168h: a saved history is checked as of when it was fetched; raise the limit for an older copy, or run sign-off again to fetch it fresh)"
# Printed under the account-policy check of a verdict with no sign-off yet:
# policy v4 requires one, so the verifier exits 4 until it is given.
SIGNOFF_PENDING_REASON = "(exit 4, ANCHORED, SIGN-OFF NOT SHOWN, until a sign-off is given: the policy requires one on the verdict; run sign-off, then verify-commands again)"
FEED_AGE_REASON = "(168h: a saved feed is checked as of when it was fetched; raise the limit for an older copy, or run feed again to fetch it fresh)"
HERE = Path(__file__).resolve().parent
REPO = HERE.parent


def read_input(path):
    """A file's exact bytes as text. Text mode would turn Windows line endings
    into newlines and change the draft's sealed hash (R170)."""
    return Path(path).read_bytes().decode("utf-8")


def sha256_hex(data: bytes) -> str:
    return hashlib.sha256(data).hexdigest()


def canonical(value: Any) -> str:
    return json.dumps(value, sort_keys=True, separators=(",", ":"), ensure_ascii=False)


# Fixed identifiers, derived from fixed text so every language seals the same.
PROGRAM_ID = sha256_hex(b"verdifax-claims-review-demo/program/v1")
REGISTRY_RECORD_HASH = sha256_hex(b"verdifax-claims-review-demo/registry/v1")
REGISTRY_RECORD_ENV = "CLAIMS_REVIEW_REGISTRY_RECORD_HASH"


def registry_record_hash(given: Optional[str]) -> str:
    """The program registry record every step claims. A server that enforces
    a program registry, as Verdifax's hardware-attested endpoint does, admits
    a step only when it claims the record its operator registered for this
    program: CLAIMS_REVIEW_REGISTRY_RECORD_HASH names that record. Unset, the
    program claims its fixed value, which a server with no enforced registry
    seals as declared."""
    if not given:
        return REGISTRY_RECORD_HASH
    if not re.fullmatch(r"[0-9a-f]{64}", given):
        fail("CLAIMS_REVIEW_REGISTRY_RECORD_HASH must be 64 lowercase hex characters")
    return given


def load_rules() -> tuple[list[dict[str, str]], str]:
    doc = json.loads((REPO / "rules" / (RULE_SET + ".json")).read_text(encoding="utf-8"))
    return doc["rules"], sha256_hex(canonical(doc).encode("utf-8"))


def load_prompt(name: str = "prompt-v1.txt") -> str:
    return (REPO / "rules" / name).read_text(encoding="utf-8")


def base_url() -> str:
    return os.environ.get("VERDIFAX_BASE_URL", DEFAULT_BASE_URL).rstrip("/")


LOCAL_HOSTS = ("127.0.0.1", "localhost", "::1", "host.docker.internal")


def is_local_server() -> bool:
    """Whether the server is on this machine (or, from a container, its
    host). Recorded answers are allowed only there: every other server,
    Verdifax's hosted one or its hardware-attested one, is live, and a
    review against it asks the model."""
    from urllib.parse import urlsplit
    return (urlsplit(base_url()).hostname or "") in LOCAL_HOSTS


def write_private(path: Path, text: str) -> None:
    fd = os.open(path, os.O_WRONLY | os.O_CREAT | os.O_TRUNC, 0o600)
    with os.fdopen(fd, "w", encoding="utf-8") as f:
        f.write(text)


def write_json(path: Path, value: Any) -> None:
    path.parent.mkdir(parents=True, exist_ok=True)
    path.write_text(json.dumps(value, indent=2, ensure_ascii=False) + "\n", encoding="utf-8")


def fail(message: str) -> None:
    print("claims-review:", message, file=sys.stderr)
    sys.exit(1)


def draft_classification(text: str) -> str:
    first = text.splitlines()[0] if text else ""
    m = re.match(r"^<!--\s*classification:\s*([A-Z_]+)\s*-->$", first.strip())
    level = m.group(1) if m else "INTERNAL"
    if level not in LEVELS:
        fail("the draft declares an unknown classification: %s" % level)
    return level


def keys_path() -> Path:
    return Path("demo-keys.json")


def load_keys() -> dict[str, Any]:
    if not keys_path().exists():
        fail("no demo-keys.json here; run `claims-review setup` first")
    return json.loads(keys_path().read_text(encoding="utf-8"))


def client_for(key: dict[str, Any]) -> verdifax.VerdifaxClient:
    """A client for a key; one whose request key was cleared sends unsigned."""
    signing = bytes.fromhex(key["signing_key"]) if key.get("signing_key") else None
    return verdifax.VerdifaxClient(base_url=base_url(), api_key=key["secret"], signing_key=signing)


def usage_substance(usage: Optional[dict[str, int]]) -> Optional[dict[str, Any]]:
    """The request_substance a step declares: the model's own usage report for
    the step's model call, so the server's budgets count real tokens. The gate
    step calls no model and declares none (None)."""
    if usage is None:
        return None
    return {"cost": {"input_tokens": usage["input_tokens"], "output_tokens": usage["output_tokens"],
                     "total_tokens": usage["input_tokens"] + usage["output_tokens"]}}


def data_quality(intake_draft_sha256: str, draft_sha256_now: str, rule_set_hash: str, as_of: str) -> dict[str, Any]:
    """The data checks a step declares, each computed (round 62)."""
    def result(ok: bool) -> str:
        return "pass" if ok else "fail"
    return {"source": DATA_QUALITY_SOURCE, "as_of": as_of, "checks": [
        {"name": DATA_QUALITY_CHECKS[0], "result": result(draft_sha256_now == intake_draft_sha256)},
        {"name": DATA_QUALITY_CHECKS[1], "result": result(rule_set_hash == RULE_SET_SHA256)}]}


def setting_error(setting: str, value: str) -> Optional[str]:
    """Why a settings command is refused before anything is sent, or None."""
    if setting not in ACCOUNT_SETTINGS:
        return "the setting must be %s, not %s" % (" or ".join(ACCOUNT_SETTINGS), setting)
    if value not in SETTING_VALUES:
        return "the value must be %s, not %s" % (" or ".join(SETTING_VALUES), value)
    return None


def approver_name_error(name: str) -> Optional[str]:
    if not APPROVER_NAME.match(name):
        return "an approver name is 1 to 40 lowercase letters, digits or dashes"
    return None


def approver_files(name: str) -> tuple[Path, Path]:
    """A named approver's private key file (0600) and public key file."""
    return Path("approver-%s-key.json" % name), Path("approver-%s.txt" % name)



# A key that was ever on the trial tier must keep a request key, and the
# demo's accounts are trial accounts, so the agent key can replace its request
# key but never clear it. For the request-key chapter Verdifax makes one
# production-tier member key in the account, its creation sealed in the feed
# as Verdifax's: the holder. request-key-register keeps it in demo-keys.json,
# and --holder makes the request-key commands and key-revoke act on it.
def with_signing_key(keys: dict[str, Any], role: str, signing_key_hex: Optional[str]) -> dict[str, Any]:
    """The keys file with one key's request key replaced, or removed when it
    was cleared; role is "agent" or "holder"."""
    entry = {k: v for k, v in keys[role].items() if k != "signing_key"}
    if signing_key_hex is not None:
        entry["signing_key"] = signing_key_hex
    return dict(keys, **{role: entry})


def with_agent_signing_key(keys: dict[str, Any], signing_key_hex: Optional[str]) -> dict[str, Any]:
    """The keys file with the agent key's request key replaced, or removed
    when it was cleared."""
    return with_signing_key(keys, "agent", signing_key_hex)


def key_for(keys: dict[str, Any], holder: bool) -> tuple[str, dict[str, Any]]:
    """The key a request-key command or key-revoke acts on: the agent, or with
    --holder the key Verdifax made for the request-key chapter."""
    role = "holder" if holder else "agent"
    if role not in keys:
        fail("demo-keys.json holds no request-key holder; run request-key-register with the key Verdifax made first")
    return role, keys[role]


# ── setup ──────────────────────────────────────────────────────────────────

def cmd_setup(_: argparse.Namespace) -> None:
    if keys_path().exists():
        fail("demo-keys.json exists here; this folder is already set up, and every program in this repository uses it")
    owner_secret = os.environ.get("VERDIFAX_OWNER_KEY", "")
    if not owner_secret:
        fail("set VERDIFAX_OWNER_KEY to the account's owner key")
    owner_signer = verdifax.generate_request_key()
    owner = verdifax.VerdifaxClient(base_url=base_url(), api_key=owner_secret, signing_key=owner_signer.private_key)
    owner.register_account_request_key()
    member = owner.create_account_key("claims-review-agent")
    member_signer = verdifax.generate_request_key()
    agent = verdifax.VerdifaxClient(base_url=base_url(), api_key=member["secret"], signing_key=member_signer.private_key)
    agent.register_request_key()
    policy_text = read_input(REPO / "policy" / POLICY_FILE)
    uploaded = owner.upload_account_policy(policy_text)
    owner.activate_account_policy(uploaded["document_hash"])
    approver = verdifax.generate_request_key()
    owner.register_approver("Claims review approver", approver.private_key)
    feed = owner.create_feed_key("claims-review-feed")
    write_json(Path("account-history.json"), owner.get_account_policy_history())
    write_private(Path("approver-key.json"), json.dumps({"private_key": approver.private_key.hex()}) + "\n")
    Path("approver.txt").write_text(base64.b64encode(approver.public_key).decode("ascii"), encoding="utf-8")
    write_private(keys_path(), json.dumps({
        "base_url": base_url(),
        "owner": {"secret": owner_secret, "signing_key": owner_signer.private_key.hex()},
        "agent": {"id": member["id"], "secret": member["secret"], "signing_key": member_signer.private_key.hex()},
        "feed": {"id": feed["id"], "secret": feed["secret"]},
        "policy_document_hash": uploaded["document_hash"],
    }, indent=2) + "\n")
    print("setup done: agent key %s, feed key %s, policy %s" % (member["id"], feed["id"], uploaded["document_hash"]))


# ── the model ──────────────────────────────────────────────────────────────

def recorded_usage(a: dict[str, Any]) -> dict[str, int]:
    """The model use a recorded answer declares, 0 when it names none. Only
    tests against a local server read recorded answers; the budget stops are
    rehearsed with answers that declare large use. Never part of the answer."""
    return {"input_tokens": int(a.get("input_tokens", 0)), "output_tokens": int(a.get("output_tokens", 0))}


class Answers:
    """Recorded answers for tests, refused against production."""

    def __init__(self, path: Optional[str]) -> None:
        self.recorded = None
        if path:
            if not is_local_server():
                fail("--answers is for tests against a local server, never a live one")
            self.recorded = json.loads(Path(path).read_text(encoding="utf-8"))

    def ask(self, prompt: str, user_text: str, rule_id: str, model: str) -> dict[str, Any]:
        if self.recorded is not None:
            a = self.recorded[rule_id]
            return {"model": a.get("model", "recorded-answers"), "response_id": a.get("response_id", "recorded"),
                    "verdict": a["verdict"], "reason": a["reason"], "raw": canonical({"verdict": a["verdict"], "reason": a["reason"]}),
                    **recorded_usage(a)}
        import anthropic
        from pydantic import BaseModel

        class RuleAnswer(BaseModel):
            verdict: str
            reason: str

        client = anthropic.Anthropic(max_retries=MODEL_RETRY_TRIES - 1)
        response = client.messages.parse(model=model, max_tokens=4000, system=prompt,
                                         messages=[{"role": "user", "content": user_text}], output_format=RuleAnswer)
        usage = {"input_tokens": response.usage.input_tokens, "output_tokens": response.usage.output_tokens}
        if response.stop_reason == "refusal" or response.parsed_output is None:
            return {"model": response.model, "response_id": response.id, "verdict": "unanswered",
                    "reason": "the model declined to answer", "raw": "", **usage}
        parsed = response.parsed_output
        verdict = parsed.verdict if parsed.verdict in ("pass", "fail") else "unanswered"
        raw = next((b.text for b in response.content if b.type == "text"), "")
        return {"model": response.model, "response_id": response.id, "verdict": verdict, "reason": parsed.reason, "raw": raw, **usage}


    def ask_claims(self, prompt: str, user_text: str, model: str) -> dict[str, Any]:
        """The draft's claims, as the model listed them. The exact answer is
        the step's AI output, reviewed and sealed."""
        if self.recorded is not None:
            a = self.recorded["claims"]
            return {"model": a.get("model", "recorded-answers"), "response_id": a.get("response_id", "recorded"),
                    "claims": a["claims"], "raw": canonical({"claims": a["claims"]}), **recorded_usage(a)}
        import anthropic
        from pydantic import BaseModel

        class ClaimList(BaseModel):
            claims: list[str]

        client = anthropic.Anthropic(max_retries=MODEL_RETRY_TRIES - 1)
        response = client.messages.parse(model=model, max_tokens=4000, system=prompt,
                                         messages=[{"role": "user", "content": user_text}], output_format=ClaimList)
        usage = {"input_tokens": response.usage.input_tokens, "output_tokens": response.usage.output_tokens}
        raw = next((b.text for b in response.content if b.type == "text"), "")
        if response.stop_reason == "refusal" or response.parsed_output is None:
            return {"model": response.model, "response_id": response.id, "claims": [],
                    "raw": raw or "the model declined to list the claims", **usage}
        return {"model": response.model, "response_id": response.id, "claims": response.parsed_output.claims, "raw": raw, **usage}

    def ask_explanation(self, prompt: str, user_text: str, model: str) -> dict[str, Any]:
        """The model's plain-words explanation of the verdict code decided.
        The exact answer is the verdict step's AI output, reviewed and sealed."""
        if self.recorded is not None:
            a = self.recorded["verdict"]
            return {"model": a.get("model", "recorded-answers"), "response_id": a.get("response_id", "recorded"),
                    "explanation": a["explanation"], "raw": canonical({"explanation": a["explanation"]}),
                    **recorded_usage(a)}
        import anthropic
        from pydantic import BaseModel

        class Explanation(BaseModel):
            explanation: str

        client = anthropic.Anthropic(max_retries=MODEL_RETRY_TRIES - 1)
        response = client.messages.parse(model=model, max_tokens=2000, system=prompt,
                                         messages=[{"role": "user", "content": user_text}], output_format=Explanation)
        usage = {"input_tokens": response.usage.input_tokens, "output_tokens": response.usage.output_tokens}
        raw = next((b.text for b in response.content if b.type == "text"), "")
        if response.stop_reason == "refusal" or response.parsed_output is None:
            return {"model": response.model, "response_id": response.id, "explanation": "",
                    "raw": raw or "the model declined to explain the verdict", **usage}
        return {"model": response.model, "response_id": response.id, "explanation": response.parsed_output.explanation,
                "raw": raw, **usage}


def measured_cost(model: str, input_tokens: int, output_tokens: int) -> Optional[float]:
    prices = json.loads((REPO / "rules" / "model-prices-v1.json").read_text(encoding="utf-8"))["prices"]
    if model not in prices:
        return None
    return input_tokens / 1e6 * prices[model]["input"] + output_tokens / 1e6 * prices[model]["output"]


# ── review ─────────────────────────────────────────────────────────────────

def cmd_review(args: argparse.Namespace) -> None:
    registry_record = registry_record_hash(os.environ.get(REGISTRY_RECORD_ENV))
    keys = load_keys()
    agent = client_for(keys["agent"])
    rules, rule_set_hash = load_rules()
    prompt = load_prompt()
    draft_path = Path(args.draft)
    draft = read_input(draft_path)
    draft_hash = sha256_hex(draft.encode("utf-8"))
    # The time the draft was read: the as_of of every step's data checks.
    read_at = dt.datetime.now(dt.timezone.utc).strftime("%Y-%m-%dT%H:%M:%SZ")

    def context(step: str) -> dict[str, Any]:
        """The step's policy context: its object, the draft's classification
        and the data checks, computed again now, the draft re-read."""
        now_hash = sha256_hex(read_input(draft_path).encode("utf-8"))
        return {"object_type": "text", "object_id": OBJECT_PREFIX + step, "classification": level,
                "data_quality": data_quality(draft_hash, now_hash, rule_set_hash, read_at)}
    level = draft_classification(draft)
    answers = Answers(args.answers)
    now = dt.datetime.now(dt.timezone.utc).strftime("%Y%m%dT%H%M%SZ")
    job_id = "claims-review-%s-%s" % (draft_hash[:12], now)
    out = Path("records") / job_id
    payloads: dict[str, dict[str, Any]] = {}

    intake_payload = {"form": FORM, "step": "intake", "draft_sha256": draft_hash, "draft_bytes": len(draft.encode("utf-8")),
                      "rule_set": RULE_SET, "rule_set_hash": rule_set_hash, "classification": level}
    try:
        job = agent.job(job_id)
        # The gate step names its agent, as policy v4 requires of every step;
        # --without-actor leaves the name out on purpose, so the policy
        # refuses the draft here, before any model is called.
        agent_named = None if args.without_actor else verdifax.AttestedContext(attested=True, actor_id=ACTOR_ID)
        intake = job.step("intake", after=[], payload=canonical(intake_payload), program_id=PROGRAM_ID, route_id=ROUTE_INTAKE,
                          registry_record_hash=registry_record, attested_context=agent_named,
                          policy_context=context("intake"))
    except APIError as exc:
        # The account policy refused the draft at intake (for example a
        # confidential draft), so it never reached a model. The refusal is
        # sealed as a receipt; nothing else runs for this draft.
        stop(agent, job_id, "intake", {}, {"intake": intake_payload}, {}, str(draft_path), draft_hash, exc, args.model,
             {"input": 0, "output": 0}, set(), {})
        return
    payloads["intake"] = intake_payload
    run_ids = {"intake": intake.run_id}

    results: list[dict[str, Any]] = []
    tokens = {"input": 0, "output": 0}
    models: set[str] = set()
    # The model's exact answer, the text Verdifax's AI review reads, kept as
    # sent so a reader recomputes the sealed ai_output_hash from the file.
    answer_texts: dict[str, str] = {}
    # The text the server's AI review checked, returned once with each step.
    trace_texts: dict[str, str] = {}

    # The gate passed, so the draft may reach a model. Every step from here
    # carries the model's answer, which Verdifax's AI review reads and the
    # record seals; the account policy refuses any step without one.
    claims_prompt = load_prompt(CLAIMS_PROMPT)
    claims_text = "Draft:\n%s" % draft
    claims_prompt_hash = sha256_hex((claims_prompt + "\n" + claims_text).encode("utf-8"))
    listed = answers.ask_claims(claims_prompt, claims_text, args.model)
    tokens["input"] += listed["input_tokens"]
    tokens["output"] += listed["output_tokens"]
    models.add(listed["model"])
    answer_texts["claims"] = listed["raw"]
    claims_payload = {"form": FORM, "step": "claims", "draft_sha256": draft_hash, "rule_set_hash": rule_set_hash,
                      "model": listed["model"], "model_response_id": listed["response_id"], "prompt_hash": claims_prompt_hash,
                      "claims": listed["claims"]}
    try:
        claims_step = job.step("claims", after=["intake"], payload=canonical(claims_payload), program_id=PROGRAM_ID,
                               route_id=ROUTE_CLAIMS, registry_record_hash=registry_record, ai_output_text=answer_texts["claims"],
                               request_substance=usage_substance(listed),
                               attested_context=verdifax.AttestedContext(attested=True, actor_id=ACTOR_ID, model_provider="anthropic",
                                                                         model_name=listed["model"], prompt_hash=claims_prompt_hash,
                                                                         decision_kind="claims.list",
                                                                         decision_result=str(len(listed["claims"]))),
                               policy_context=context("claims"))
    except APIError as exc:
        payloads["claims"] = claims_payload
        if exc.aivp_trace_text:
            trace_texts["claims"] = exc.aivp_trace_text
        stop(agent, job_id, "claims", run_ids, payloads, answer_texts, str(draft_path), draft_hash, exc, args.model, tokens, models,
             trace_texts)
        return
    if claims_step.aivp_trace_text:
        trace_texts["claims"] = claims_step.aivp_trace_text
    payloads["claims"] = claims_payload
    run_ids["claims"] = claims_step.run_id
    print("claims: %d listed" % len(listed["claims"]))

    for rule in rules:
        user_text = "Rule %s: %s\n\nDraft:\n%s" % (rule["id"], rule["text"], draft)
        prompt_hash = sha256_hex((prompt + "\n" + user_text).encode("utf-8"))
        answer = answers.ask(prompt, user_text, rule["id"], args.model)
        tokens["input"] += answer["input_tokens"]
        tokens["output"] += answer["output_tokens"]
        models.add(answer["model"])
        answer_texts[rule["id"]] = answer["raw"] or answer["reason"]
        payload = {"form": FORM, "step": rule["id"], "rule_id": rule["id"], "rule_text": rule["text"], "draft_sha256": draft_hash,
                   "rule_set_hash": rule_set_hash, "model": answer["model"], "model_response_id": answer["response_id"],
                   "prompt_hash": prompt_hash, "verdict": answer["verdict"], "reason": answer["reason"]}
        try:
            receipt = job.step(rule["id"], after=["claims"], payload=canonical(payload), program_id=PROGRAM_ID, route_id=ROUTE_RULE,
                               registry_record_hash=registry_record, ai_output_text=answer_texts[rule["id"]],
                               request_substance=usage_substance(answer),
                               attested_context=verdifax.AttestedContext(attested=True, actor_id=ACTOR_ID, model_provider="anthropic",
                                                                         model_name=answer["model"], prompt_hash=prompt_hash,
                                                                         decision_kind="claims.rule_check", decision_result=answer["verdict"]),
                               policy_context=context(rule["id"]))
        except APIError as exc:
            # The account policy, the server's content screen or its review
            # of the model's answer stopped this step, so the job stops here:
            # the stop is sealed as a receipt, the steps already sealed are
            # kept beside it, and a person sees the stop. The job never gets
            # a verdict.
            payloads[rule["id"]] = payload
            if exc.aivp_trace_text:
                trace_texts[rule["id"]] = exc.aivp_trace_text
            stop(agent, job_id, rule["id"], run_ids, payloads, answer_texts, str(draft_path), draft_hash, exc, args.model, tokens, models,
                 trace_texts)
            return
        if receipt.aivp_trace_text:
            trace_texts[rule["id"]] = receipt.aivp_trace_text
        payloads[rule["id"]] = payload
        run_ids[rule["id"]] = receipt.run_id
        results.append({"rule_id": rule["id"], "verdict": answer["verdict"], "reason": answer["reason"]})
        print("%s %s: %s" % (rule["id"], answer["verdict"], answer["reason"]))

    # The program decides the verdict from the rule answers, so anyone can
    # decide it again from the sealed records; the model only explains it.
    failed = [r["rule_id"] for r in results if r["verdict"] != "pass"]
    verdict = "pass" if not failed else "fail"
    verdict_prompt = load_prompt(VERDICT_PROMPT)
    verdict_text = "Verdict: %s\n\nRule results:\n%s" % (verdict, canonical(results))
    verdict_prompt_hash = sha256_hex((verdict_prompt + "\n" + verdict_text).encode("utf-8"))
    explained = answers.ask_explanation(verdict_prompt, verdict_text, args.model)
    tokens["input"] += explained["input_tokens"]
    tokens["output"] += explained["output_tokens"]
    models.add(explained["model"])
    answer_texts["verdict"] = explained["raw"]
    verdict_payload = {"form": FORM, "step": "verdict", "draft_sha256": draft_hash, "rule_set_hash": rule_set_hash,
                       "verdict": verdict, "failed_rules": failed, "results": results, "model": explained["model"],
                       "model_response_id": explained["response_id"], "prompt_hash": verdict_prompt_hash,
                       "explanation": explained["explanation"]}
    try:
        final = job.step("verdict", after=[r["id"] for r in rules], final=True, payload=canonical(verdict_payload), program_id=PROGRAM_ID,
                         route_id=ROUTE_VERDICT, registry_record_hash=registry_record, ai_output_text=answer_texts["verdict"],
                         request_substance=usage_substance(explained),
                         attested_context=verdifax.AttestedContext(attested=True, actor_id=ACTOR_ID, model_provider="anthropic",
                                                                   model_name=explained["model"], prompt_hash=verdict_prompt_hash,
                                                                   decision_kind="claims.verdict", decision_result=verdict),
                         policy_context=context("verdict"))
    except APIError as exc:
        payloads["verdict"] = verdict_payload
        if exc.aivp_trace_text:
            trace_texts["verdict"] = exc.aivp_trace_text
        stop(agent, job_id, "verdict", run_ids, payloads, answer_texts, str(draft_path), draft_hash, exc, args.model, tokens, models,
             trace_texts)
        return
    if final.aivp_trace_text:
        trace_texts["verdict"] = final.aivp_trace_text
    payloads["verdict"] = verdict_payload
    run_ids["verdict"] = final.run_id
    print("verdict explained: %s" % explained["explanation"])

    save_job_records(agent, job_id, out, run_ids, payloads)
    save_answers(out, answer_texts)
    save_texts(out / TRACES_FOLDER, trace_texts)
    cost = measured_cost(args.model, tokens["input"], tokens["output"]) if len(models) == 1 and args.model in models else None
    summary = {"job_id": job_id, "draft": str(draft_path), "draft_sha256": draft_hash, "rule_set_hash": rule_set_hash,
               "verdict": verdict, "failed_rules": failed, "models": sorted(models),
               "final_run_id": final.run_id, "run_ids": run_ids,
               "measured_tokens": tokens, "measured_cost_usd": cost}
    write_json(out / "job.json", summary)
    print("verdict %s (%d of %d rules failed); job %s; records in %s" % (summary["verdict"], len(failed), len(rules), job_id, out))
    if cost is not None:
        print("measured model use: %d input and %d output tokens, about $%.2f at published prices" % (tokens["input"], tokens["output"], cost))
    print_verify_commands(job_id)


# A step the server stopped is sealed as one run whose status names the kind
# of stop, with a receipt of that kind (R152). Any other error is not a stop:
# the program reports it as the server gave it and seals nothing for it.
STOP_KINDS = {
    "pepg_deny": "the account policy or the content screen",
    "aivp_t4_halt": "the server's review of the model's answer",
    "ccv_halt": "the per-run budget",
    "macc_halt": "the cumulative budget",
    "attestation_gap": "the policy's required attested fields",
}
STOP_RECEIPT_FIELDS = {"pepg_deny": "deny_receipt", "aivp_t4_halt": "aivp_t4_halt_receipt", "ccv_halt": "ccv_halt_receipt",
                       "macc_halt": "macc_halt_receipt", "attestation_gap": "attestation_gap_receipt"}


def stop_receipt(agent: verdifax.VerdifaxClient, status: str, run_id: int) -> dict[str, Any]:
    fetch = {"pepg_deny": agent.get_deny_receipt, "aivp_t4_halt": agent.get_aivp_t4_halt_receipt,
             "ccv_halt": agent.get_ccv_halt_receipt, "macc_halt": agent.get_macc_halt_receipt,
             "attestation_gap": agent.get_attestation_gap_receipt}[status]
    return fetch(run_id)


def stop_receipt_hash(agent: verdifax.VerdifaxClient, run_id: int) -> str:
    """The sealed hash of a stopped run's receipt, whatever kind of stop it was."""
    status = agent.get_run(run_id).status
    if status not in STOP_KINDS:
        fail("run %d is %s, not a stopped run with a receipt" % (run_id, status))
    return stop_receipt(agent, status, run_id)[STOP_RECEIPT_FIELDS[status]]["hash"]


def stopped_run(agent: verdifax.VerdifaxClient, job_id: str, step: str, exc: APIError) -> Any:
    """The run the server sealed for this job's stopped step, found by its job
    and step, never by being the newest of the key."""
    rows = [r for r in agent.export_runs(job_id=job_id).runs if r.step_id == step and r.status in STOP_KINDS]
    if not rows:
        fail("the server answered HTTP %d at %s: %s; it recorded no stop for this step, so nothing was sealed as one"
             % (exc.status_code, step, exc.message))
    return max(rows, key=lambda r: r.run_id)


def save_texts(folder: Path, texts: dict[str, str]) -> None:
    """Writes each step's text to <folder>/<step>.txt, byte for byte."""
    if not texts:
        return
    folder.mkdir(parents=True, exist_ok=True)
    for step, text in texts.items():
        (folder / ("%s.txt" % step)).write_bytes(text.encode("utf-8"))


def save_answers(folder: Path, answer_texts: dict[str, str]) -> None:
    """The model's exact answers, one file per step, so each sealed
    ai_output_hash recomputes from its file."""
    (folder / ANSWERS_FOLDER).mkdir(parents=True, exist_ok=True)
    for step, text in answer_texts.items():
        (folder / ANSWERS_FOLDER / ("%s.txt" % step)).write_bytes(text.encode("utf-8"))


def stopped_folder(run_id: int) -> Path:
    """The folder of the stopped step whose receipt is run run_id."""
    return Path("records") / ("%s%d" % (STOPPED_FOLDER_PREFIX, run_id))


def stop(agent: verdifax.VerdifaxClient, job_id: str, step: str, run_ids: dict[str, int], payloads: dict[str, Any],
         answer_texts: dict[str, str], draft: str, draft_hash: str, exc: APIError, model: str, tokens: dict[str, int],
         models: set[str], trace_texts: dict[str, str]) -> None:
    """Record a stopped step: its sealed receipt, every record the job sealed
    before it, the exact payload bytes and model answers of every step
    including the stopped one, the model use spent before the stop, and a
    summary naming the step and the kind of stop. The same for intake and for
    any later step."""
    row = stopped_run(agent, job_id, step, exc)
    stopped = row.run_id
    folder = stopped_folder(stopped)  # named by the run, so two stops never share a folder
    write_json(folder / "receipt.json", stop_receipt(agent, row.status, stopped))
    # A sealed step whose record was deleted since is kept as its tombstone.
    deleted = deleted_steps(read_account_feed(load_keys()), job_id) if run_ids else {}
    write_deleted_steps(folder, {s: t for s, t in deleted.items() if s in run_ids})
    for sealed_step, run_id in run_ids.items():
        if sealed_step in deleted:
            continue
        (folder / "steps").mkdir(parents=True, exist_ok=True)
        (folder / "steps" / ("%s.json" % sealed_step)).write_text(agent.get_run_artifacts(run_id), encoding="utf-8")
    (folder / "payloads").mkdir(parents=True, exist_ok=True)
    for name, payload in payloads.items():
        (folder / "payloads" / ("%s.json" % name)).write_text(canonical(payload), encoding="utf-8")
    save_answers(folder, answer_texts)
    save_texts(folder / TRACES_FOLDER, trace_texts)
    cost = measured_cost(model, tokens["input"], tokens["output"]) if len(models) == 1 and model in models else None
    write_json(folder / "job.json", {"stopped_run_id": stopped, "step": step, "stop_status": row.status, "job_id": job_id,
                                     "draft": draft, "draft_sha256": draft_hash, "reason": exc.message, "run_ids": run_ids,
                                     "models": sorted(models), "measured_tokens": tokens, "measured_cost_usd": cost})
    print("stopped at %s by %s: run %d, %s" % (step, STOP_KINDS[row.status], stopped, exc.message))
    print("records in %s" % folder)
    if cost is not None:
        print("measured model use before the stop: %d input and %d output tokens, about $%.2f at published prices"
              % (tokens["input"], tokens["output"], cost))


def save_job_records(agent: verdifax.VerdifaxClient, job_id: str, out: Path, run_ids: dict[str, int], payloads: dict[str, Any]) -> None:
    # A deleted step's record is gone from the export; its tombstone in the
    # account feed stands in its place (R250).
    deleted = {s: t for s, t in deleted_steps(read_account_feed(load_keys()), job_id).items() if s in run_ids}
    rows = [r for r in agent.export_runs(job_id=job_id).runs if r.step_id not in deleted]
    if len(rows) + len(deleted) != len(run_ids):
        fail("the export lists %d records and the feed %d deleted steps for %s, the program sealed %d"
             % (len(rows), len(deleted), job_id, len(run_ids)))
    # The sealed step records alone sit in steps/, the folder verdifax-verify
    # --job reads: it takes every .json file there as a record, so the
    # summary, the sign-offs and the payloads live beside that folder.
    (out / "steps").mkdir(parents=True, exist_ok=True)
    write_deleted_steps(out, deleted)
    for row in rows:
        if row.status in STOP_KINDS:
            write_json(out / "steps" / ("%s.json" % row.step_id), stop_receipt(agent, row.status, row.run_id))
        else:
            (out / "steps" / ("%s.json" % row.step_id)).write_text(agent.get_run_artifacts(row.run_id), encoding="utf-8")
    # The exact bytes each step sent, so a reader recomputes the sealed
    # payload hash from the file.
    (out / "payloads").mkdir(parents=True, exist_ok=True)
    for step, payload in payloads.items():
        (out / "payloads" / ("%s.json" % step)).write_text(canonical(payload), encoding="utf-8")


def deleted_steps(events: list[Any], job_id: str) -> dict[str, Any]:
    """The steps of job job_id whose records were deleted, each with its
    tombstone: the entry of every run.deleted event naming the job, by its
    step. A run.sealed event whose record was deleted carries the tombstone in
    data.deleted with a null entry; it is not the deletion and is passed over,
    as is a deletion of another job or of no job."""
    out: dict[str, Any] = {}
    for e in events:
        data = e.get("data") or {}
        entry = data.get("entry")
        if data.get("kind") != "run.deleted" or not entry or entry.get("job_id") != job_id or not entry.get("step_id"):
            continue
        out[entry["step_id"]] = entry
    return out


def write_deleted_steps(folder: Path, deleted: dict[str, Any]) -> None:
    """Writes each deleted step's tombstone as <folder>/steps/<step>.deleted.json,
    every field as the server sent it, beside the step records verdifax-verify
    --job reads."""
    for step, entry in deleted.items():
        write_json(folder / "steps" / ("%s%s" % (step, DELETED_STEP_SUFFIX)), entry)


def read_account_feed(keys: dict[str, Any]) -> list[Any]:
    """The account's whole feed, read with the feed key, page by page."""
    feed = verdifax.VerdifaxClient(base_url=base_url(), api_key=keys["feed"]["secret"])
    events: list[Any] = []
    after = 0
    while True:
        page = feed.get_signals(after=after, limit=500)
        events.extend(page["events"])
        if len(page["events"]) < 500:
            break
        after = page["next_after"]
    return events


def history_check(path: Path) -> tuple[str, str]:
    """The re-check of a saved sign-off history, with its head age limit, and
    the reason for the limit on the line after (R246)."""
    return ("verdifax-verify --signoff-history %s --signoff-head-max-age %s" % (path, SAVED_HEAD_MAX_AGE), "  " + HISTORY_AGE_REASON)


def feed_check() -> tuple[str, str]:
    """The re-check of the saved feed, with its head age limit, and the reason
    for the limit on the line after (R246)."""
    return ("verdifax-verify --signals signals.json --feed-head-max-age %s" % SAVED_HEAD_MAX_AGE, "  " + FEED_AGE_REASON)


# ── sign-off ───────────────────────────────────────────────────────────────

def approver_private_key(name: Optional[str] = None) -> bytes:
    """The first approver's private key, or a named approver's."""
    path = approver_files(name)[0] if name else Path("approver-key.json")
    if not path.exists():
        fail("no %s here; run `claims-review %s` first" % (path, "approver-add " + name if name else "setup"))
    return bytes.fromhex(json.loads(path.read_text(encoding="utf-8"))["private_key"])


def cmd_sign_off(args: argparse.Namespace) -> None:
    keys = load_keys()
    agent = client_for(keys["agent"])
    owner = client_for(keys["owner"])
    if args.approver and approver_name_error(args.approver):
        fail(approver_name_error(args.approver))
    key = approver_private_key(args.approver)
    approver_file = approver_files(args.approver)[1] if args.approver else Path("approver.txt")
    if args.receipt:
        run_id = int(args.receipt)
        target = {"kind": "receipt", "run_id": run_id, "digest": stop_receipt_hash(agent, run_id)}
        folder = stopped_folder(run_id)
    else:
        folder = Path("records") / args.job_id
        summary = json.loads((folder / "job.json").read_text(encoding="utf-8"))
        run_id = summary["final_run_id"]
        artifacts = json.loads((folder / "steps" / "verdict.json").read_text(encoding="utf-8"))
        target = {"kind": "job", "run_id": run_id, "digest": artifacts["bundle_anchor"]["anchored_digest"], "job_id": args.job_id}
    record = agent.sign_off(target, args.decision, key, note=args.note)
    write_json(folder / "signoffs.json", agent.get_signoffs(run_id))
    write_json(folder / "signoff-history.json", owner.get_signoff_history())
    print("sign-off %s recorded as sequence %s, hash %s" % (args.decision, record["sequence"], record["signoff_hash"]))
    # R272: a sign-off is checked with the account's whole sign-off history, so
    # a later decision cannot be left out; a rule that requires a sign-off,
    # as policy v4's verdict-step does, fails the check without it.
    print("verify: verdifax-verify --signoff %s --signoff-history %s --signoff-head-max-age %s --approver-key %s %s"
          % (folder / "signoffs.json", folder / "signoff-history.json", SAVED_HEAD_MAX_AGE, approver_file,
             folder / "receipt.json" if args.receipt else folder / "steps" / "verdict.json"))
    history, history_reason = history_check(folder / "signoff-history.json")
    print("verify: " + history)
    print(history_reason)


# ── the account: approvers, settings, keys, deletion, webhooks ─────────────
# Each change is made with the owner key (the agent key for its own request
# key) and sealed by the server in the account's feed, which `feed` saves:
# the feed is the record of every change, so no command keeps its own.

def owner_client() -> verdifax.VerdifaxClient:
    return client_for(load_keys()["owner"])


def approver_id(name: str) -> int:
    """The id the server gave this approver when it was added, kept with its key."""
    private = approver_files(name)[0]
    if not private.exists():
        fail("no %s here; run `claims-review approver-add %s` first" % (private, name))
    return int(json.loads(private.read_text(encoding="utf-8"))["id"])


def cmd_approver_add(args: argparse.Namespace) -> None:
    if approver_name_error(args.name):
        fail(approver_name_error(args.name))
    private, public = approver_files(args.name)
    if private.exists():
        fail("%s exists here; this approver is already added" % private)
    owner = owner_client()
    approver = verdifax.generate_request_key()
    answer = owner.register_approver("Claims review approver %s" % args.name, approver.private_key)
    write_private(private, json.dumps({"private_key": approver.private_key.hex(), "id": answer["id"]}) + "\n")
    public.write_text(base64.b64encode(approver.public_key).decode("ascii"), encoding="utf-8")
    write_json(Path("approvers.json"), owner.list_approvers())
    print("approver %s added; its public key is in %s" % (args.name, public))


def cmd_approver_revoke(args: argparse.Namespace) -> None:
    if approver_name_error(args.name):
        fail(approver_name_error(args.name))
    owner = owner_client()
    owner.revoke_approver(approver_id(args.name))
    write_json(Path("approvers.json"), owner.list_approvers())
    print("approver %s revoked; the sign-offs it gave stay as they were" % args.name)


def cmd_settings(args: argparse.Namespace) -> None:
    if setting_error(args.setting, args.value):
        fail(setting_error(args.setting, args.value))
    owner_client().set_account_setting(args.setting, args.value)
    print("%s is %s" % (args.setting, args.value))


def cmd_request_key_register(args: argparse.Namespace) -> None:
    keys = load_keys()
    if "holder" in keys:
        fail("demo-keys.json already holds the request-key holder")
    made = json.loads(Path(args.key_file).read_text(encoding="utf-8"))
    signer = verdifax.generate_request_key()
    holder = verdifax.VerdifaxClient(base_url=base_url(), api_key=made["secret"], signing_key=signer.private_key)
    holder.register_request_key()
    entry = {"id": made["id"], "secret": made["secret"], "signing_key": signer.private_key.hex()}
    write_private(keys_path(), json.dumps(dict(keys, holder=entry), indent=2) + "\n")
    print("a request key was registered on key %s, the key Verdifax made for the request-key chapter" % made["id"])


def cmd_request_key_rotate(args: argparse.Namespace) -> None:
    keys = load_keys()
    role, key = key_for(keys, args.holder)
    new = verdifax.generate_request_key()
    client_for(key).rotate_request_key(new.private_key)
    write_private(keys_path(), json.dumps(with_signing_key(keys, role, new.private_key.hex()), indent=2) + "\n")
    print("the %s key's request key was replaced, approved by the key it replaced" % role)


def cmd_request_key_clear(args: argparse.Namespace) -> None:
    keys = load_keys()
    role, key = key_for(keys, args.holder)
    client_for(key).clear_request_key()
    write_private(keys_path(), json.dumps(with_signing_key(keys, role, None), indent=2) + "\n")
    print("the %s key's request key was cleared, approved by that key" % role)


def cmd_key_revoke(args: argparse.Namespace) -> None:
    keys = load_keys()
    role, key = key_for(keys, args.holder)
    owner = client_for(keys["owner"])
    owner.revoke_account_key(int(key["id"]))
    write_json(Path("account-keys.json"), owner.list_account_keys())
    print("%s key %s revoked; the records it sealed stay as they were" % (role, key["id"]))


def replace_with_tombstone(folder: Path, step: str, tombstone: dict[str, Any]) -> None:
    """A deleted step's record is replaced by its tombstone, which
    verdifax-verify --job reads in its place (R250)."""
    write_deleted_steps(folder, {step: tombstone})
    record = folder / "steps" / ("%s.json" % step)
    if record.exists():
        record.unlink()


def cmd_delete_step(args: argparse.Namespace) -> None:
    folder = Path("records") / args.job_id
    record_path = folder / "steps" / ("%s.json" % args.step)
    if not record_path.exists():
        fail("%s does not exist; delete-step needs the step's sealed record" % record_path)
    record = json.loads(record_path.read_text(encoding="utf-8"))
    owner_client().delete_run(int(record["run_id"]), record["manifest_hash"])
    tombstone = deleted_steps(read_account_feed(load_keys()), args.job_id).get(args.step)
    if tombstone is None:
        fail("the feed holds no run.deleted event for %s of %s" % (args.step, args.job_id))
    replace_with_tombstone(folder, args.step, tombstone)
    print("step %s of %s deleted; its tombstone is %s" % (args.step, args.job_id, folder / "steps" / (args.step + DELETED_STEP_SUFFIX)))
    print("verify: verdifax-verify --job %s" % (folder / "steps"))


def cmd_webhook_subscribe(args: argparse.Namespace) -> None:
    hook = owner_client().create_webhook(args.url)
    write_private(Path("webhook.json"), json.dumps(hook, indent=2) + "\n")
    print("webhook %s subscribed to every kind of event at %s; its secret is in webhook.json" % (hook["id"], args.url))


def cmd_webhook_list(_: argparse.Namespace) -> None:
    # The lists alone, the same shape in every language: webhooks.json the
    # account's webhooks, webhook-deliveries.json each one's deliveries by id.
    owner = owner_client()
    hooks = owner.list_webhooks().get("webhooks") or []
    write_json(Path("webhooks.json"), hooks)
    deliveries = {str(h["id"]): owner.list_webhook_deliveries(int(h["id"]), limit=500).get("deliveries") or [] for h in hooks}
    write_json(Path("webhook-deliveries.json"), deliveries)
    print("%d webhook(s); their deliveries are in webhook-deliveries.json" % len(hooks))


# ── feed ───────────────────────────────────────────────────────────────────

def cmd_webhook_remove(_: argparse.Namespace) -> None:
    """Removes the subscription webhook-subscribe made, as the owner, sealed in
    the feed as webhook.deleted. webhook.json is kept: its secret can be
    published once the subscription is gone, since it then signs nothing."""
    path = Path("webhook.json")
    if not path.exists():
        fail("no webhook.json here; run `claims-review webhook-subscribe URL` first")
    hook_id = int(json.loads(path.read_text(encoding="utf-8"))["id"])
    write_json(Path("webhook-removed.json"), owner_client().delete_webhook(hook_id))
    print("webhook %d removed; its removal is sealed in the account's feed" % hook_id)


# ── publish ────────────────────────────────────────────────────────────────
# The account's owner key makes a record public on its random link, or
# withdraws it. Each decision is sealed by the server as a run publication
# and placed in the account's feed, which `feed` then reads.

def publication_targets(job_id: Optional[str], receipt: Optional[str]) -> tuple[Path, dict[str, int]]:
    """The folder and the records a publish or unpublish acts on: every step
    of a job, or a stopped step's receipt with every step the job sealed
    before it, so a reader sees what the job did before it stopped."""
    if receipt:
        run_id = int(receipt)
        folder = stopped_folder(run_id)
        summary = json.loads((folder / "job.json").read_text(encoding="utf-8"))
        return folder, dict(summary["run_ids"], receipt=run_id)
    folder = Path("records") / str(job_id)
    summary = json.loads((folder / "job.json").read_text(encoding="utf-8"))
    return folder, dict(summary["run_ids"])


def publishable(folder: Path, targets: dict[str, int]) -> list[str]:
    """The records a publish or unpublish acts on, sorted: every target but a
    deleted step, which has no record left; its tombstone stands in its place
    in the feed (R250)."""
    return sorted(n for n in targets if not (folder / "steps" / ("%s%s" % (n, DELETED_STEP_SUFFIX))).exists())


def publication(args: argparse.Namespace, publish: bool) -> None:
    keys = load_keys()
    owner = client_for(keys["owner"])
    folder, targets = publication_targets(args.job_id, args.receipt)
    # A repeated publish changes nothing and the server answers with no new
    # publication, so the sealed publication saved earlier is kept.
    earlier = json.loads((folder / "published.json").read_text(encoding="utf-8")) if (folder / "published.json").exists() else {}
    published: dict[str, Any] = {}
    for name in publishable(folder, targets):
        run_id = targets[name]
        answer = owner.publish_run(run_id) if publish else owner.unpublish_run(run_id)
        # A completed run is served at its bundle address, a stopped run at
        # its receipt address (R156); each answer names only the one it has.
        sealed = answer["publication"]
        if sealed is None and earlier.get(name, {}).get("run_id") == run_id and earlier[name].get("public") == answer["public"]:
            sealed = earlier[name].get("publication")
        published[name] = {"run_id": run_id, "public": answer["public"], "public_bundle_url": answer.get("public_bundle_url"),
                           "public_receipt_url": answer.get("public_receipt_url"),
                           "verify_url": answer["verify_url"], "publication": sealed}
    write_json(folder / "published.json", published)
    print("%s %d record(s) in %s" % ("published" if publish else "withdrawn", len(published), folder))
    if publish:
        shown = published.get("verdict") or published.get("receipt") or published[sorted(published)[0]]
        print("public link: %s" % shown["verify_url"])
    feed, feed_reason = feed_check()
    print("each decision is an event in the account's feed; after `feed`, verify: " + feed)
    print(feed_reason)


def cmd_publish(args: argparse.Namespace) -> None:
    publication(args, True)


def cmd_unpublish(args: argparse.Namespace) -> None:
    publication(args, False)


def cmd_feed(_: argparse.Namespace) -> None:
    events = read_account_feed(load_keys())
    write_json(Path("signals.json"), events)
    kinds: dict[str, int] = {}
    for e in events:
        kinds[e["data"]["kind"]] = kinds.get(e["data"]["kind"], 0) + 1
    print("%d events in the feed: %s" % (len(events), ", ".join("%s %d" % kv for kv in sorted(kinds.items()))))
    feed, feed_reason = feed_check()
    print("verify: " + feed)
    print(feed_reason)


# ── verify commands ────────────────────────────────────────────────────────

def print_verify_commands(job_id: str) -> None:
    folder = Path("records") / job_id
    print("verify with verdifax-verify %s or later:" % MIN_VERIFIER)
    traces = " --aivp-traces %s" % (folder / TRACES_FOLDER) if (folder / TRACES_FOLDER).is_dir() else ""
    print("  verdifax-verify --job %s%s" % (folder / "steps", traces))
    if (folder / "signoffs.json").exists():
        print("  verdifax-verify --account-policy %s --signoff %s --signoff-history %s --signoff-head-max-age %s --approver-key approver.txt %s"
              % (REPO / "policy" / POLICY_FILE, folder / "signoffs.json", folder / "signoff-history.json", SAVED_HEAD_MAX_AGE,
                 folder / "steps" / "verdict.json"))
    else:
        print("  verdifax-verify --account-policy %s %s" % (REPO / "policy" / POLICY_FILE, folder / "steps" / "verdict.json"))
        print("  " + SIGNOFF_PENDING_REASON)
    if (folder / "signoffs.json").exists():
        print("  verdifax-verify --signoff %s --signoff-history %s --signoff-head-max-age %s --approver-key approver.txt %s"
              % (folder / "signoffs.json", folder / "signoff-history.json", SAVED_HEAD_MAX_AGE, folder / "steps" / "verdict.json"))
    if (folder / "signoff-history.json").exists():
        history, history_reason = history_check(folder / "signoff-history.json")
        print("  " + history)
        print(history_reason)
    if Path("signals.json").exists():
        feed, feed_reason = feed_check()
        print("  " + feed)
        print(feed_reason)


def cmd_verify_commands(args: argparse.Namespace) -> None:
    print_verify_commands(args.job_id)


def main(argv: Optional[list[str]] = None) -> None:
    parser = argparse.ArgumentParser(prog="claims-review", description=__doc__, formatter_class=argparse.RawDescriptionHelpFormatter)
    sub = parser.add_subparsers(dest="command", required=True)
    sub.add_parser("setup").set_defaults(fn=cmd_setup)
    review = sub.add_parser("review")
    review.add_argument("draft")
    review.add_argument("--model", default=DEFAULT_MODEL)
    review.add_argument("--answers", help="recorded answers for tests against a local server")
    review.add_argument("--without-actor", action="store_true", help="send the gate step with no agent name, so policy v4 refuses it")
    review.set_defaults(fn=cmd_review)
    sign = sub.add_parser("sign-off")
    sign.add_argument("job_id", nargs="?")
    sign.add_argument("decision", choices=("approve", "reject"))
    sign.add_argument("--note")
    sign.add_argument("--receipt", help="sign off a stopped step's receipt by run id")
    sign.add_argument("--approver", help="sign with an approver added by approver-add, not the first")
    sign.set_defaults(fn=cmd_sign_off)
    for name, fn in (("publish", cmd_publish), ("unpublish", cmd_unpublish)):
        pub = sub.add_parser(name)
        pub.add_argument("job_id", nargs="?")
        pub.add_argument("--receipt", help="act on a stopped step's receipt, and the steps sealed before it, by run id")
        pub.set_defaults(fn=fn)
    for name, fn in (("approver-add", cmd_approver_add), ("approver-revoke", cmd_approver_revoke)):
        cmd = sub.add_parser(name)
        cmd.add_argument("name")
        cmd.set_defaults(fn=fn)
    settings = sub.add_parser("settings")
    settings.add_argument("setting")
    settings.add_argument("value")
    settings.set_defaults(fn=cmd_settings)
    register = sub.add_parser("request-key-register")
    register.add_argument("key_file", help="the key Verdifax made for the request-key chapter, as JSON with id and secret")
    register.set_defaults(fn=cmd_request_key_register)
    for name, fn in (("request-key-rotate", cmd_request_key_rotate), ("request-key-clear", cmd_request_key_clear),
                     ("key-revoke", cmd_key_revoke)):
        holder_cmd = sub.add_parser(name)
        holder_cmd.add_argument("--holder", action="store_true", help="act on the key Verdifax made for the request-key chapter")
        holder_cmd.set_defaults(fn=fn)
    delete = sub.add_parser("delete-step")
    delete.add_argument("job_id")
    delete.add_argument("step")
    delete.set_defaults(fn=cmd_delete_step)
    hook = sub.add_parser("webhook-subscribe")
    hook.add_argument("url")
    hook.set_defaults(fn=cmd_webhook_subscribe)
    sub.add_parser("webhook-list").set_defaults(fn=cmd_webhook_list)
    sub.add_parser("webhook-remove").set_defaults(fn=cmd_webhook_remove)
    sub.add_parser("feed").set_defaults(fn=cmd_feed)
    verify = sub.add_parser("verify-commands")
    verify.add_argument("job_id")
    verify.set_defaults(fn=cmd_verify_commands)
    args = parser.parse_args(argv)
    if args.command in ("sign-off", "publish", "unpublish") and not args.receipt and not args.job_id:
        parser.error("%s needs a JOB_ID, or --receipt RUN_ID" % args.command)
    args.fn(args)


if __name__ == "__main__":
    main()
