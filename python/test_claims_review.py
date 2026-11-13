"""Unit tests for the Python program: the same cases every language tests."""

import hashlib
import json
import pathlib
import types
import unittest
from unittest import mock

import claims_review as cr
from verdifax import APIError

HERE = pathlib.Path(__file__).resolve().parent
FIXED = {"b": 1, "a": "x", "c": [1, 2], "d": {"z": True, "y": None}, "e": "ü \"q\" \\ /"}
EXPECTED = '{"a":"x","b":1,"c":[1,2],"d":{"y":null,"z":true},"e":"ü \\"q\\" \\\\ /"}'
# The rule set in force and its hash are named once, in demo-parity.json.
PARITY = json.loads((HERE.parent / "demo-parity.json").read_text(encoding="utf-8"))
# The stop cases every language runs, from one file (R152).
STOPS = json.loads((HERE.parent / "testdata" / "stop-cases.json").read_text(encoding="utf-8"))


class KeptTexts(unittest.TestCase):
    """The review traces and model answers are kept byte for byte, so each
    file hashes to the value its record seals."""

    def test_each_text_is_written_exactly(self):
        import hashlib
        import tempfile
        with tempfile.TemporaryDirectory() as d:
            folder = pathlib.Path(d) / "traces"
            texts = {"rule-01": "1. The vault is locked.\n\n2. Caf\u00e9 \u201cquoted\u201d.\n", "rule-02": "no trailing newline"}
            cr.save_texts(folder, texts)
            for step, text in texts.items():
                kept = (folder / ("%s.txt" % step)).read_bytes()
                self.assertEqual(hashlib.sha256(kept).hexdigest(), hashlib.sha256(text.encode("utf-8")).hexdigest())
            cr.save_texts(pathlib.Path(d) / "none", {})
            self.assertFalse((pathlib.Path(d) / "none").exists(), "no folder is made when there is nothing to keep")


class CanonicalJSON(unittest.TestCase):
    def test_canonical_json_matches_every_language(self):
        self.assertEqual(cr.canonical(FIXED), EXPECTED)
        self.assertEqual(cr.sha256_hex(b"abc"), "ba7816bf8f01cfea414140de5dae2223b00361a396177a9cb410ff61f20015ad")

    def test_rule_set_hash(self):
        rules, rule_set_hash = cr.load_rules()
        self.assertEqual(len(rules), 15)
        self.assertEqual(cr.RULE_SET, PARITY["constants"]["rule_set"])
        self.assertEqual(rule_set_hash, PARITY["rule_set_sha256"])

    def test_classification_line(self):
        self.assertEqual(cr.draft_classification("# A draft\n\nText."), "INTERNAL")
        self.assertEqual(cr.draft_classification("<!-- classification: CONFIDENTIAL -->\n# Memo"), "CONFIDENTIAL")

    def test_parity_copy_matches_the_source(self):
        source = (HERE.parent / "demo-parity.json").read_text(encoding="utf-8")
        copy = (HERE / "testdata" / "demo-parity.json").read_text(encoding="utf-8")
        self.assertEqual(copy, source)
        parity = json.loads(source)
        self.assertEqual(parity["constants"]["default_model"], cr.DEFAULT_MODEL)
        self.assertEqual(parity["constants"]["form"], cr.FORM)
        self.assertEqual(int(parity["constants"]["model_retry_tries"]), cr.MODEL_RETRY_TRIES)

    def test_the_model_retry_budget_waits_out_an_outage(self):
        """R259: the SDK waits half a second before the first retry and
        doubles it to a cap of eight, so the shared budget must reach a
        minute of waiting. Round 61: the default of two retries waited about
        a second and a half, which a 503 from the model API outlasted."""
        waits = [min(8.0, 0.5 * 2 ** i) for i in range(cr.MODEL_RETRY_TRIES - 1)]
        self.assertGreaterEqual(sum(waits), 60)


class Failed(Exception):
    """Raised by the patched fail, so a test reads the message the program would print."""


class FakeAgent:
    """Answers as the server would from stop-cases.json. Each receipt route
    serves only runs stopped by its own kind, so a fetch of the wrong kind
    fails the test."""

    def __init__(self):
        self.status = {r["run_id"]: r["status"] for r in STOPS["receipts"]}
        self.status[STOPS["not_stopped"]["run_id"]] = STOPS["not_stopped"]["status"]
        self.hashes = {r["run_id"]: (r["field"], r["hash"]) for r in STOPS["receipts"]}
        for kind, method in (("pepg_deny", "get_deny_receipt"), ("aivp_t4_halt", "get_aivp_t4_halt_receipt"), ("ccv_halt", "get_ccv_halt_receipt"),
                             ("macc_halt", "get_macc_halt_receipt"), ("attestation_gap", "get_attestation_gap_receipt")):
            setattr(self, method, self._receipt_route(kind))

    def _receipt_route(self, kind):
        def fetch(run_id):
            if self.status.get(run_id) != kind:
                raise AssertionError("the %s route was asked for run %d, which is %s" % (kind, run_id, self.status.get(run_id)))
            field, digest = self.hashes[run_id]
            return {field: {"hash": digest}}
        return fetch

    def get_run(self, run_id):
        return types.SimpleNamespace(status=self.status[run_id])

    def export_runs(self, job_id):
        rows = STOPS["export_rows"] if job_id == STOPS["job_id"] else []
        return types.SimpleNamespace(runs=[types.SimpleNamespace(**r) for r in rows])


def failing(message):
    raise Failed(message)


class StopHandling(unittest.TestCase):
    def test_stop_kinds_match_the_parity_file(self):
        constants = PARITY["constants"]
        self.assertEqual(sorted(cr.STOP_KINDS), sorted(constants["stop_statuses"]))
        self.assertEqual(cr.STOP_RECEIPT_FIELDS, dict(zip(constants["stop_statuses"], constants["stop_receipt_fields"])))

    def test_stopped_run_is_found_by_job_and_step(self):
        error = APIError(STOPS["error"]["message"], STOPS["error"]["status"])
        for case in STOPS["stopped_run"]:
            with self.subTest(step=case["step"]), mock.patch.object(cr, "fail", side_effect=failing):
                if "failure" in case:
                    with self.assertRaises(Failed) as caught:
                        cr.stopped_run(FakeAgent(), STOPS["job_id"], case["step"], error)
                    self.assertEqual(str(caught.exception), case["failure"])
                else:
                    row = cr.stopped_run(FakeAgent(), STOPS["job_id"], case["step"], error)
                    self.assertEqual((row.run_id, row.status), (case["run_id"], case["status"]))

    def test_stop_receipt_hash_reads_the_receipt_of_its_kind(self):
        for case in STOPS["receipts"]:
            with self.subTest(status=case["status"]):
                self.assertEqual(cr.stop_receipt_hash(FakeAgent(), case["run_id"]), case["hash"])
                self.assertEqual(cr.stop_receipt(FakeAgent(), case["status"], case["run_id"]), {case["field"]: {"hash": case["hash"]}})

    def test_a_run_that_did_not_stop_has_no_receipt(self):
        with mock.patch.object(cr, "fail", side_effect=failing), self.assertRaises(Failed) as caught:
            cr.stop_receipt_hash(FakeAgent(), STOPS["not_stopped"]["run_id"])
        self.assertEqual(str(caught.exception), STOPS["not_stopped"]["failure"])


if __name__ == "__main__":
    unittest.main()


class InputBytes(unittest.TestCase):
    """R170: the program reads every input as its exact bytes, so the hash it
    seals is the one shasum prints, and the parity file lists every sample
    and policy."""

    def test_every_input_file_reads_as_its_exact_bytes(self):
        repo = HERE.parent
        listed = sorted(PARITY["input_sha256"])
        on_disk = sorted("%s/%s" % (d, p.name) for d in ("samples", "policy") for p in (repo / d).iterdir())
        self.assertEqual(listed, on_disk)
        for name in listed:
            read = hashlib.sha256(cr.read_input(repo / name).encode("utf-8")).hexdigest()
            self.assertEqual(read, PARITY["input_sha256"][name], name)


# R250: a deleted step's tombstone, from the cases every language runs.
DELETED = json.loads((HERE.parent / "testdata" / "deleted-step-cases.json").read_text(encoding="utf-8"))


class DeletedSteps(unittest.TestCase):
    def test_a_deleted_step_is_read_from_its_jobs_run_deleted_events_only(self):
        self.assertEqual(cr.deleted_steps(DELETED["events"], DELETED["job_id"]), DELETED["want"])
        self.assertEqual(cr.deleted_steps(DELETED["events"], "claims-review-no-such-job"), {})
        self.assertEqual(cr.DELETED_STEP_SUFFIX, PARITY["constants"]["deleted_step_suffix"])

    def test_a_deleted_steps_tombstone_is_written_beside_the_steps_with_every_field(self):
        # The SDK hands over the feed as parsed JSON, so the program writes
        # every field again: each written tombstone must equal the served one
        # field for field, with the same JSON types, job_links included
        # (round 61), so it hashes to the same deletion_hash.
        import tempfile
        def typed(value):
            if isinstance(value, dict):
                return {k: typed(v) for k, v in value.items()}
            if isinstance(value, list):
                return [typed(v) for v in value]
            return (type(value).__name__, value)
        served = cr.deleted_steps(DELETED["events"], DELETED["job_id"])
        with tempfile.TemporaryDirectory() as d:
            cr.write_deleted_steps(pathlib.Path(d), served)
            written = sorted(p.name for p in (pathlib.Path(d) / "steps").iterdir())
            self.assertEqual(written, ["rule-02.deleted.json", "verdict.deleted.json"])
            for step, entry in DELETED["want"].items():
                kept = json.loads((pathlib.Path(d) / "steps" / (step + cr.DELETED_STEP_SUFFIX)).read_text(encoding="utf-8"))
                self.assertEqual(typed(kept), typed(entry))
                self.assertEqual(list(kept), list(entry))
                self.assertEqual(kept["job_links"], entry["job_links"])


class HeadAgeLimits(unittest.TestCase):
    """R246: every printed re-check of a saved history or feed states its
    head age limit, with the reason on the line after."""

    def test_printed_rechecks_state_their_head_age_limit(self):
        import contextlib
        import io
        import os
        import tempfile
        here = os.getcwd()
        printed = io.StringIO()
        with tempfile.TemporaryDirectory() as d:
            try:
                os.chdir(d)
                (pathlib.Path("records") / "job-1").mkdir(parents=True)
                (pathlib.Path("records") / "job-1" / "signoff-history.json").write_text("[]", encoding="utf-8")
                pathlib.Path("signals.json").write_text("[]", encoding="utf-8")
                with contextlib.redirect_stdout(printed):
                    cr.print_verify_commands("job-1")
            finally:
                os.chdir(here)
        lines = printed.getvalue().splitlines()
        constants = PARITY["constants"]
        self.assertEqual(cr.MIN_VERIFIER, constants["min_verifier"])
        self.assertEqual(lines[0], "verify with verdifax-verify %s or later:" % cr.MIN_VERIFIER)
        history = next(i for i, line in enumerate(lines) if "--signoff-history" in line)
        feed = next(i for i, line in enumerate(lines) if "--signals" in line)
        self.assertEqual(lines[history], "  verdifax-verify --signoff-history %s --signoff-head-max-age %s"
                         % (pathlib.Path("records") / "job-1" / "signoff-history.json", cr.SAVED_HEAD_MAX_AGE))
        self.assertEqual(lines[history + 1], "  " + cr.HISTORY_AGE_REASON)
        self.assertEqual(lines[feed], "  verdifax-verify --signals signals.json --feed-head-max-age %s" % cr.SAVED_HEAD_MAX_AGE)
        self.assertEqual(lines[feed + 1], "  " + cr.FEED_AGE_REASON)
        self.assertEqual((cr.SAVED_HEAD_MAX_AGE, cr.HISTORY_AGE_REASON, cr.FEED_AGE_REASON),
                         (constants["saved_head_max_age"], constants["history_age_reason"], constants["feed_age_reason"]))
