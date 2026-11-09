package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

// The same cases every language tests.
const (
	expectedCanonical = `{"a":"x","b":1,"c":[1,2],"d":{"y":null,"z":true},"e":"ü \"q\" \\ /"}`
)

func TestCanonicalJSONMatchesEveryLanguage(t *testing.T) {
	fixed := map[string]any{"b": 1, "a": "x", "c": []int{1, 2}, "d": map[string]any{"z": true, "y": nil}, "e": "ü \"q\" \\ /"}
	if got := canonical(fixed); got != expectedCanonical {
		t.Fatalf("canonical = %s", got)
	}
}

// The rule set in force and its hash are named once, in demo-parity.json.
func TestRuleSetHash(t *testing.T) {
	text, readErr := os.ReadFile(filepath.Join(repo, "demo-parity.json"))
	if readErr != nil {
		t.Fatalf("read demo-parity.json: %v", readErr)
	}
	var parity struct {
		RuleSetSHA256 string `json:"rule_set_sha256"`
		Constants     struct {
			RuleSet string `json:"rule_set"`
		} `json:"constants"`
	}
	if parseErr := json.Unmarshal(text, &parity); parseErr != nil {
		t.Fatalf("parse demo-parity.json: %v", parseErr)
	}
	rules, hash := loadRules()
	if len(rules) != 15 || ruleSet != parity.Constants.RuleSet || hash != parity.RuleSetSHA256 {
		t.Fatalf("%d rules, rule set %s, hash %s; demo-parity.json names %s, %s", len(rules), ruleSet, hash, parity.Constants.RuleSet, parity.RuleSetSHA256)
	}
}

func TestClassificationLine(t *testing.T) {
	if draftClassification("# A draft\n\nText.") != "INTERNAL" || draftClassification("<!-- classification: CONFIDENTIAL -->\n# Memo") != "CONFIDENTIAL" {
		t.Fatal("the classification line is not read")
	}
}

func TestParityCopyMatchesTheSource(t *testing.T) {
	source, sourceErr := os.ReadFile(filepath.Join(repo, "demo-parity.json"))
	copyText, copyErr := os.ReadFile(filepath.Join("testdata", "demo-parity.json"))
	if sourceErr != nil || copyErr != nil || string(source) != string(copyText) {
		t.Fatalf("the parity copy differs from the source (%v %v)", sourceErr, copyErr)
	}
}
