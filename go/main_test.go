package main

import (
	"os"
	"path/filepath"
	"testing"
)

// The same cases every language tests.
const (
	expectedCanonical = `{"a":"x","b":1,"c":[1,2],"d":{"y":null,"z":true},"e":"ü \"q\" \\ /"}`
	expectedRuleSet   = "61110740e99a53d6fe48378a52a450f053ef2d887ef491be55801f75b4e4e1ad"
)

func TestCanonicalJSONMatchesEveryLanguage(t *testing.T) {
	fixed := map[string]any{"b": 1, "a": "x", "c": []int{1, 2}, "d": map[string]any{"z": true, "y": nil}, "e": "ü \"q\" \\ /"}
	if got := canonical(fixed); got != expectedCanonical {
		t.Fatalf("canonical = %s", got)
	}
}

func TestRuleSetHash(t *testing.T) {
	rules, hash := loadRules()
	if len(rules) != 15 || hash != expectedRuleSet {
		t.Fatalf("%d rules, hash %s", len(rules), hash)
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
