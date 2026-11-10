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
