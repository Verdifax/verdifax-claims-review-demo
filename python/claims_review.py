#!/usr/bin/env python3
"""claims-review: a small agent that checks a draft against 15 general writing
rules, with every step sealed by Verdifax.

A demo, not customer proof. A sealed record shows which rules were checked and
what the model answered; it does not show the model's judgment was right. The
15 rules are general writing rules, not legal advice and not a compliance
standard. The account policy judges what each step declared; it cannot show
the declared facts are true. A sign-off shows a registered approver key signed
that decision; it does not show who holds the key or that the review was
careful.

Commands (the same in every language of this repository):

  claims-review setup                              owner key in VERDIFAX_OWNER_KEY
  claims-review review DRAFT.md [--model ID]        model key in ANTHROPIC_API_KEY
  claims-review sign-off JOB_ID approve|reject [--note TEXT]
  claims-review sign-off --receipt RUN_ID reject
  claims-review feed
  claims-review verify-commands JOB_ID

Environment: VERDIFAX_BASE_URL (default https://api.verdifax.com).
Files, in the working folder: demo-keys.json and approver-key.json (0600),
approver.txt, account-history.json, records/<job_id>/*.json with payloads/,
signals.json.
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
RULE_SET = "v1"
ACTOR_ID = "claims-review-agent"
ROUTE_INTAKE = "claims-review/v1/intake"
ROUTE_RULE = "claims-review/v1/rule-check"
ROUTE_VERDICT = "claims-review/v1/verdict"
OBJECT_PREFIX = "claims-review/v1/"
LEVELS = ("PUBLIC", "INTERNAL", "CONFIDENTIAL", "RESTRICTED", "TOP_SECRET", "TS_SCI")
VERDICTS = ("pass", "fail", "unanswered")
MIN_VERIFIER = "v0.24.0"
HERE = Path(__file__).resolve().parent
REPO = HERE.parent


def sha256_hex(data: bytes) -> str:
    return hashlib.sha256(data).hexdigest()


def canonical(value: Any) -> str:
    return json.dumps(value, sort_keys=True, separators=(",", ":"), ensure_ascii=False)


# Fixed identifiers, derived from fixed text so every language seals the same.
PROGRAM_ID = sha256_hex(b"verdifax-claims-review-demo/program/v1")
REGISTRY_RECORD_HASH = sha256_hex(b"verdifax-claims-review-demo/registry/v1")


def load_rules() -> tuple[list[dict[str, str]], str]:
    doc = json.loads((REPO / "rules" / "v1.json").read_text(encoding="utf-8"))
    return doc["rules"], sha256_hex(canonical(doc).encode("utf-8"))


def load_prompt() -> str:
    return (REPO / "rules" / "prompt-v1.txt").read_text(encoding="utf-8")


def base_url() -> str:
    return os.environ.get("VERDIFAX_BASE_URL", DEFAULT_BASE_URL).rstrip("/")


def is_production() -> bool:
    return base_url() == DEFAULT_BASE_URL


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
    return verdifax.VerdifaxClient(base_url=base_url(), api_key=key["secret"],
                                   signing_key=bytes.fromhex(key["signing_key"]))


# ── setup ──────────────────────────────────────────────────────────────────

def cmd_setup(_: argparse.Namespace) -> None:
    if keys_path().exists():
        fail("demo-keys.json exists here; setup runs once per account and folder")
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
    policy_text = (REPO / "policy" / "claims-review-v1.json").read_text(encoding="utf-8")
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

class Answers:
    """Recorded answers for tests, refused against production."""

    def __init__(self, path: Optional[str]) -> None:
        self.recorded = None
        if path:
            if is_production():
                fail("--answers is for tests against a local server, never production")
            self.recorded = json.loads(Path(path).read_text(encoding="utf-8"))

    def ask(self, prompt: str, user_text: str, rule_id: str, model: str) -> dict[str, Any]:
        if self.recorded is not None:
            a = self.recorded[rule_id]
            return {"model": a.get("model", "recorded-answers"), "response_id": a.get("response_id", "recorded"),
                    "verdict": a["verdict"], "reason": a["reason"], "raw": canonical({"verdict": a["verdict"], "reason": a["reason"]}),
                    "input_tokens": 0, "output_tokens": 0}
        import anthropic
        from pydantic import BaseModel

        class RuleAnswer(BaseModel):
            verdict: str
            reason: str

        client = anthropic.Anthropic()
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


def measured_cost(model: str, input_tokens: int, output_tokens: int) -> Optional[float]:
    prices = json.loads((REPO / "rules" / "model-prices-v1.json").read_text(encoding="utf-8"))["prices"]
    if model not in prices:
        return None
    return input_tokens / 1e6 * prices[model]["input"] + output_tokens / 1e6 * prices[model]["output"]


# ── review ─────────────────────────────────────────────────────────────────

def cmd_review(args: argparse.Namespace) -> None:
    keys = load_keys()
    agent = client_for(keys["agent"])
    rules, rule_set_hash = load_rules()
    prompt = load_prompt()
    draft_path = Path(args.draft)
    draft = draft_path.read_text(encoding="utf-8")
    draft_hash = sha256_hex(draft.encode("utf-8"))
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
        intake = job.step("intake", after=[], payload=canonical(intake_payload), program_id=PROGRAM_ID, route_id=ROUTE_INTAKE,
                          registry_record_hash=REGISTRY_RECORD_HASH,
                          policy_context={"object_type": "text", "object_id": OBJECT_PREFIX + "intake", "classification": level})
    except APIError as exc:
        # The account policy refused the draft at intake (for example a
        # confidential draft), so it never reached a model. The refusal is
        # sealed as a receipt; nothing else runs for this draft.
        refuse(agent, job_id, "intake", {}, {"intake": intake_payload}, str(draft_path), draft_hash, exc.message)
        return
    payloads["intake"] = intake_payload
    run_ids = {"intake": intake.run_id}

    results: list[dict[str, Any]] = []
    tokens = {"input": 0, "output": 0}
    models: set[str] = set()
    for rule in rules:
        user_text = "Rule %s: %s\n\nDraft:\n%s" % (rule["id"], rule["text"], draft)
        prompt_hash = sha256_hex((prompt + "\n" + user_text).encode("utf-8"))
        answer = answers.ask(prompt, user_text, rule["id"], args.model)
        tokens["input"] += answer["input_tokens"]
        tokens["output"] += answer["output_tokens"]
        models.add(answer["model"])
        payload = {"form": FORM, "step": rule["id"], "rule_id": rule["id"], "rule_text": rule["text"], "draft_sha256": draft_hash,
                   "rule_set_hash": rule_set_hash, "model": answer["model"], "model_response_id": answer["response_id"],
                   "prompt_hash": prompt_hash, "verdict": answer["verdict"], "reason": answer["reason"]}
        try:
            receipt = job.step(rule["id"], after=["intake"], payload=canonical(payload), program_id=PROGRAM_ID, route_id=ROUTE_RULE,
                               registry_record_hash=REGISTRY_RECORD_HASH, ai_output_text=answer["raw"] or answer["reason"],
                               attested_context=verdifax.AttestedContext(attested=True, actor_id=ACTOR_ID, model_provider="anthropic",
                                                                         model_name=answer["model"], prompt_hash=prompt_hash,
                                                                         decision_kind="claims.rule_check", decision_result=answer["verdict"]),
                               policy_context={"object_type": "text", "object_id": OBJECT_PREFIX + rule["id"], "classification": level})
        except APIError as exc:
            # The account policy or the server's content screen refused this
            # step, so the job stops here: the refusal is sealed as a receipt,
            # the steps already sealed are kept beside it, and a person sees
            # the stop. The job never gets a verdict.
            payloads[rule["id"]] = payload
            refuse(agent, job_id, rule["id"], run_ids, payloads, str(draft_path), draft_hash, exc.message)
            return
        payloads[rule["id"]] = payload
        run_ids[rule["id"]] = receipt.run_id
        results.append({"rule_id": rule["id"], "verdict": answer["verdict"], "reason": answer["reason"]})
        print("%s %s: %s" % (rule["id"], answer["verdict"], answer["reason"]))

    failed = [r["rule_id"] for r in results if r["verdict"] != "pass"]
    verdict_payload = {"form": FORM, "step": "verdict", "draft_sha256": draft_hash, "rule_set_hash": rule_set_hash,
                       "verdict": "pass" if not failed else "fail", "failed_rules": failed, "results": results}
    final = job.step("verdict", after=[r["id"] for r in rules], final=True, payload=canonical(verdict_payload), program_id=PROGRAM_ID,
                     route_id=ROUTE_VERDICT, registry_record_hash=REGISTRY_RECORD_HASH,
                     attested_context=verdifax.AttestedContext(attested=True, actor_id=ACTOR_ID, decision_kind="claims.verdict",
                                                               decision_result=verdict_payload["verdict"]),
                     policy_context={"object_type": "text", "object_id": OBJECT_PREFIX + "verdict", "classification": level})
    payloads["verdict"] = verdict_payload
    run_ids["verdict"] = final.run_id

    save_job_records(agent, job_id, out, run_ids, payloads)
    cost = measured_cost(args.model, tokens["input"], tokens["output"]) if len(models) == 1 and args.model in models else None
    summary = {"job_id": job_id, "draft": str(draft_path), "draft_sha256": draft_hash, "rule_set_hash": rule_set_hash,
               "verdict": verdict_payload["verdict"], "failed_rules": failed, "models": sorted(models),
               "final_run_id": final.run_id, "run_ids": run_ids,
               "measured_tokens": tokens, "measured_cost_usd": cost}
    write_json(out / "job.json", summary)
    print("verdict %s (%d of %d rules failed); job %s; records in %s" % (summary["verdict"], len(failed), len(rules), job_id, out))
    if cost is not None:
        print("measured model use: %d input and %d output tokens, about $%.2f at published prices" % (tokens["input"], tokens["output"], cost))
    print_verify_commands(job_id)


def newest_refusal(agent: verdifax.VerdifaxClient) -> int:
    rows = [r for r in agent.export_runs().runs if r.status == "pepg_deny"]
    if not rows:
        fail("the step was refused but no refused run is listed for this key")
    return max(r.run_id for r in rows)


def refuse(agent: verdifax.VerdifaxClient, job_id: str, step: str, run_ids: dict[str, int], payloads: dict[str, Any],
           draft: str, draft_hash: str, reason: str) -> None:
    """Record a refused step: its sealed receipt, every record the job sealed
    before it, the exact payload bytes of every step including the refused
    one, and a summary naming the step. The same for intake and for any
    later step."""
    refused = newest_refusal(agent)
    folder = Path("records") / ("refused-%d" % refused)  # named by the run, so two refusals never share a folder
    write_json(folder / "receipt.json", agent.get_deny_receipt(refused))
    for sealed_step, run_id in run_ids.items():
        (folder / "steps").mkdir(parents=True, exist_ok=True)
        (folder / "steps" / ("%s.json" % sealed_step)).write_text(agent.get_run_artifacts(run_id), encoding="utf-8")
    (folder / "payloads").mkdir(parents=True, exist_ok=True)
    for name, payload in payloads.items():
        (folder / "payloads" / ("%s.json" % name)).write_text(canonical(payload), encoding="utf-8")
    write_json(folder / "job.json", {"refused_run_id": refused, "step": step, "job_id": job_id, "draft": draft,
                                     "draft_sha256": draft_hash, "reason": reason, "run_ids": run_ids})
    print("refused at %s: run %d, %s" % (step, refused, reason))
    print("records in %s" % folder)


def save_job_records(agent: verdifax.VerdifaxClient, job_id: str, out: Path, run_ids: dict[str, int], payloads: dict[str, Any]) -> None:
    rows = agent.export_runs(job_id=job_id).runs
    if len(rows) != len(run_ids):
        fail("the export lists %d records for %s, the program sealed %d" % (len(rows), job_id, len(run_ids)))
    # The sealed step records alone sit in steps/, the folder verdifax-verify
    # --job reads: it takes every .json file there as a record, so the
    # summary, the sign-offs and the payloads live beside that folder.
    (out / "steps").mkdir(parents=True, exist_ok=True)
    for row in rows:
        if row.status == "pepg_deny":
            write_json(out / "steps" / ("%s.json" % row.step_id), agent.get_deny_receipt(row.run_id))
        else:
            (out / "steps" / ("%s.json" % row.step_id)).write_text(agent.get_run_artifacts(row.run_id), encoding="utf-8")
    # The exact bytes each step sent, so a reader recomputes the sealed
    # payload hash from the file.
    (out / "payloads").mkdir(parents=True, exist_ok=True)
    for step, payload in payloads.items():
        (out / "payloads" / ("%s.json" % step)).write_text(canonical(payload), encoding="utf-8")


# ── sign-off ───────────────────────────────────────────────────────────────

def approver_private_key() -> bytes:
    path = Path("approver-key.json")
    if not path.exists():
        fail("no approver-key.json here; run `claims-review setup` first")
    return bytes.fromhex(json.loads(path.read_text(encoding="utf-8"))["private_key"])


def cmd_sign_off(args: argparse.Namespace) -> None:
    keys = load_keys()
    agent = client_for(keys["agent"])
    owner = client_for(keys["owner"])
    key = approver_private_key()
    if args.receipt:
        run_id = int(args.receipt)
        receipt = agent.get_deny_receipt(run_id)["deny_receipt"]
        target = {"kind": "receipt", "run_id": run_id, "digest": receipt["hash"]}
        folder = next((p for p in Path("records").glob("refused-*") if (p / "job.json").exists()
                       and json.loads((p / "job.json").read_text())["refused_run_id"] == run_id), Path("records") / ("receipt-%d" % run_id))
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
    print("verify: verdifax-verify --signoff %s --approver-key approver.txt %s" % (folder / "signoffs.json",
          folder / "receipt.json" if args.receipt else folder / "steps" / "verdict.json"))
    print("verify: verdifax-verify --signoff-history %s" % (folder / "signoff-history.json"))


# ── feed ───────────────────────────────────────────────────────────────────

# ── publish ────────────────────────────────────────────────────────────────
# The account's owner key makes a record public on its random link, or
# withdraws it. Each decision is sealed by the server as a run publication
# and placed in the account's feed, which `feed` then reads.

def publication_targets(job_id: Optional[str], receipt: Optional[str]) -> tuple[Path, dict[str, int]]:
    """The folder and the records a publish or unpublish acts on: every step
    of a job, or one refusal's receipt."""
    if receipt:
        run_id = int(receipt)
        folder = next((p for p in Path("records").glob("refused-*") if (p / "job.json").exists()
                       and json.loads((p / "job.json").read_text())["refused_run_id"] == run_id), Path("records") / ("receipt-%d" % run_id))
        return folder, {"receipt": run_id}
    folder = Path("records") / str(job_id)
    summary = json.loads((folder / "job.json").read_text(encoding="utf-8"))
    return folder, dict(summary["run_ids"])


def publication(args: argparse.Namespace, publish: bool) -> None:
    keys = load_keys()
    owner = client_for(keys["owner"])
    folder, targets = publication_targets(args.job_id, args.receipt)
    published: dict[str, Any] = {}
    for name in sorted(targets):
        run_id = targets[name]
        answer = owner.publish_run(run_id) if publish else owner.unpublish_run(run_id)
        published[name] = {"run_id": run_id, "public": answer["public"], "public_bundle_url": answer.get("public_bundle_url"),
                           "verify_url": answer["verify_url"], "publication": answer["publication"]}
    write_json(folder / "published.json", published)
    print("%s %d record(s) in %s" % ("published" if publish else "withdrawn", len(published), folder))
    if publish:
        shown = published.get("verdict") or published.get("receipt") or published[sorted(published)[0]]
        print("public link: %s" % shown["verify_url"])
    print("each decision is an event in the account's feed; after `feed`, verify: verdifax-verify --signals signals.json")


def cmd_publish(args: argparse.Namespace) -> None:
    publication(args, True)


def cmd_unpublish(args: argparse.Namespace) -> None:
    publication(args, False)


def cmd_feed(_: argparse.Namespace) -> None:
    keys = load_keys()
    feed = verdifax.VerdifaxClient(base_url=base_url(), api_key=keys["feed"]["secret"])
    events: list[Any] = []
    after = 0
    while True:
        page = feed.get_signals(after=after, limit=500)
        events.extend(page["events"])
        if len(page["events"]) < 500:
            break
        after = page["next_after"]
    write_json(Path("signals.json"), events)
    kinds: dict[str, int] = {}
    for e in events:
        kinds[e["data"]["kind"]] = kinds.get(e["data"]["kind"], 0) + 1
    print("%d events in the feed: %s" % (len(events), ", ".join("%s %d" % kv for kv in sorted(kinds.items()))))
    print("verify: verdifax-verify --signals signals.json")


# ── verify commands ────────────────────────────────────────────────────────

def print_verify_commands(job_id: str) -> None:
    folder = Path("records") / job_id
    print("verify with verdifax-verify %s or later:" % MIN_VERIFIER)
    print("  verdifax-verify --job %s" % (folder / "steps"))
    print("  verdifax-verify --account-policy %s %s" % (REPO / "policy" / "claims-review-v1.json", folder / "steps" / "verdict.json"))
    if (folder / "signoffs.json").exists():
        print("  verdifax-verify --signoff %s --approver-key approver.txt %s" % (folder / "signoffs.json", folder / "steps" / "verdict.json"))
    if Path("signals.json").exists():
        print("  verdifax-verify --signals signals.json")


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
    review.set_defaults(fn=cmd_review)
    sign = sub.add_parser("sign-off")
    sign.add_argument("job_id", nargs="?")
    sign.add_argument("decision", choices=("approve", "reject"))
    sign.add_argument("--note")
    sign.add_argument("--receipt", help="sign off a refusal's receipt by run id")
    sign.set_defaults(fn=cmd_sign_off)
    for name, fn in (("publish", cmd_publish), ("unpublish", cmd_unpublish)):
        pub = sub.add_parser(name)
        pub.add_argument("job_id", nargs="?")
        pub.add_argument("--receipt", help="act on a refusal's receipt by run id")
        pub.set_defaults(fn=fn)
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
