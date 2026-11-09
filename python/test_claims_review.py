"""Unit tests for the Python program: the same cases every language tests."""

import json
import pathlib
import unittest

import claims_review as cr

HERE = pathlib.Path(__file__).resolve().parent
FIXED = {"b": 1, "a": "x", "c": [1, 2], "d": {"z": True, "y": None}, "e": "ü \"q\" \\ /"}
EXPECTED = '{"a":"x","b":1,"c":[1,2],"d":{"y":null,"z":true},"e":"ü \\"q\\" \\\\ /"}'
# The rule set in force and its hash are named once, in demo-parity.json.
PARITY = json.loads((HERE.parent / "demo-parity.json").read_text(encoding="utf-8"))


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


if __name__ == "__main__":
    unittest.main()
