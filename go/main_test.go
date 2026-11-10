package main

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"strings"
	"testing"

	verdifax "github.com/Verdifax/verdifax-sdk-go"
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

// The stop cases every language runs, from one file (R152).
type stopCases struct {
	JobID string `json:"job_id"`
	Error struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	} `json:"error"`
	ExportRows []verdifax.ExportRow `json:"export_rows"`
	StoppedRun []struct {
		Step    string `json:"step"`
		RunID   int64  `json:"run_id"`
		Status  string `json:"status"`
		Failure string `json:"failure"`
	} `json:"stopped_run"`
	Receipts []struct {
		RunID  int64  `json:"run_id"`
		Status string `json:"status"`
		Field  string `json:"field"`
		Hash   string `json:"hash"`
	} `json:"receipts"`
	NotStopped struct {
		RunID   int64  `json:"run_id"`
		Status  string `json:"status"`
		Failure string `json:"failure"`
	} `json:"not_stopped"`
}

func loadStopCases(t *testing.T) stopCases {
	t.Helper()
	text, readErr := os.ReadFile(filepath.Join(repo, "testdata", "stop-cases.json"))
	if readErr != nil {
		t.Fatalf("read stop-cases.json: %v", readErr)
	}
	var cases stopCases
	if parseErr := json.Unmarshal(text, &cases); parseErr != nil {
		t.Fatalf("parse stop-cases.json: %v", parseErr)
	}
	return cases
}

// fakeStopServer answers as the server would from stop-cases.json. Each
// receipt route serves only runs stopped by its own kind, so a fetch of the
// wrong kind fails the test.
type fakeStopServer struct{ cases stopCases }

func (f fakeStopServer) status(runID int64) string {
	for _, r := range f.cases.Receipts {
		if r.RunID == runID {
			return r.Status
		}
	}
	return f.cases.NotStopped.Status
}

func (f fakeStopServer) route(kind string, runID int64) (map[string]any, error) {
	for _, r := range f.cases.Receipts {
		if r.RunID == runID && r.Status == kind {
			return map[string]any{r.Field: map[string]any{"hash": r.Hash}}, nil
		}
	}
	return nil, fmt.Errorf("the %s route was asked for run %d, which is %s", kind, runID, f.status(runID))
}

func (f fakeStopServer) GetRun(_ context.Context, runID int64) (*verdifax.Run, error) {
	return &verdifax.Run{ID: runID, Status: f.status(runID)}, nil
}

func (f fakeStopServer) ExportRuns(_ context.Context, params *verdifax.ExportRunsParams) (*verdifax.RunsExport, error) {
	if params.JobID != f.cases.JobID {
		return &verdifax.RunsExport{}, nil
	}
	return &verdifax.RunsExport{Runs: append([]verdifax.ExportRow(nil), f.cases.ExportRows...)}, nil
}

func (f fakeStopServer) GetDenyReceipt(_ context.Context, id int64) (map[string]any, error) {
	return f.route("pepg_deny", id)
}

func (f fakeStopServer) GetAivpT4HaltReceipt(_ context.Context, id int64) (map[string]any, error) {
	return f.route("aivp_t4_halt", id)
}

func (f fakeStopServer) GetCCVHaltReceipt(_ context.Context, id int64) (map[string]any, error) {
	return f.route("ccv_halt", id)
}

func (f fakeStopServer) GetMACCHaltReceipt(_ context.Context, id int64) (map[string]any, error) {
	return f.route("macc_halt", id)
}

func (f fakeStopServer) GetAttestationGapReceipt(_ context.Context, id int64) (map[string]any, error) {
	return f.route("attestation_gap", id)
}

type exited struct{}

// catchFail runs call and returns the message fail printed, or "" when fail
// was not called; fail's exit becomes a panic the test recovers.
func catchFail(t *testing.T, call func()) (message string) {
	t.Helper()
	var printed bytes.Buffer
	failOut, exit = &printed, func(int) { panic(exited{}) }
	t.Cleanup(func() { failOut, exit = os.Stderr, os.Exit })
	defer func() {
		if recovered := recover(); recovered != nil {
			if _, isExit := recovered.(exited); !isExit {
				panic(recovered)
			}
			message = strings.TrimSpace(strings.TrimPrefix(printed.String(), "claims-review: "))
		}
	}()
	call()
	return ""
}

func TestStopKindsMatchTheParityFile(t *testing.T) {
	text, readErr := os.ReadFile(filepath.Join(repo, "demo-parity.json"))
	if readErr != nil {
		t.Fatalf("read demo-parity.json: %v", readErr)
	}
	var parity struct {
		Constants struct {
			StopStatuses      []string `json:"stop_statuses"`
			StopReceiptFields []string `json:"stop_receipt_fields"`
		} `json:"constants"`
	}
	if parseErr := json.Unmarshal(text, &parity); parseErr != nil {
		t.Fatalf("parse demo-parity.json: %v", parseErr)
	}
	want := map[string]string{}
	for i, status := range parity.Constants.StopStatuses {
		want[status] = parity.Constants.StopReceiptFields[i]
	}
	kinds := make([]string, 0, len(stopKinds))
	for kind := range stopKinds {
		kinds = append(kinds, kind)
	}
	statuses := append([]string(nil), parity.Constants.StopStatuses...)
	sort.Strings(kinds)
	sort.Strings(statuses)
	if !reflect.DeepEqual(kinds, statuses) || !reflect.DeepEqual(stopReceiptFields, want) {
		t.Fatalf("stop kinds %v and receipt fields %v; demo-parity.json names %v and %v", kinds, stopReceiptFields, statuses, want)
	}
}

func TestStoppedRunIsFoundByJobAndStep(t *testing.T) {
	cases := loadStopCases(t)
	apiErr := &verdifax.APIError{Status: cases.Error.Status, Message: cases.Error.Message}
	for _, c := range cases.StoppedRun {
		var row verdifax.ExportRow
		message := catchFail(t, func() { row = stoppedRun(context.Background(), fakeStopServer{cases}, cases.JobID, c.Step, apiErr) })
		switch {
		case c.Failure != "" && message != c.Failure:
			t.Errorf("%s: fail printed %q, want %q", c.Step, message, c.Failure)
		case c.Failure == "" && (message != "" || row.RunID != c.RunID || row.Status != c.Status):
			t.Errorf("%s: got run %d %s (fail %q), want run %d %s", c.Step, row.RunID, row.Status, message, c.RunID, c.Status)
		}
	}
}

func TestStopReceiptHashReadsTheReceiptOfItsKind(t *testing.T) {
	cases := loadStopCases(t)
	for _, c := range cases.Receipts {
		var digest string
		var receipt map[string]any
		message := catchFail(t, func() {
			digest = stopReceiptHash(context.Background(), fakeStopServer{cases}, c.RunID)
			receipt = stopReceipt(context.Background(), fakeStopServer{cases}, c.Status, c.RunID)
		})
		want := map[string]any{c.Field: map[string]any{"hash": c.Hash}}
		if message != "" || digest != c.Hash || !reflect.DeepEqual(receipt, want) {
			t.Errorf("%s: hash %q, receipt %v (fail %q); want %q", c.Status, digest, receipt, message, c.Hash)
		}
	}
}

func TestARunThatDidNotStopHasNoReceipt(t *testing.T) {
	cases := loadStopCases(t)
	message := catchFail(t, func() { stopReceiptHash(context.Background(), fakeStopServer{cases}, cases.NotStopped.RunID) })
	if message != cases.NotStopped.Failure {
		t.Fatalf("fail printed %q, want %q", message, cases.NotStopped.Failure)
	}
}

// R170: the program reads every input as its exact bytes, so the hash it
// seals is the one shasum prints, and the parity file lists every sample and
// policy.
func TestEveryInputFileReadsAsItsExactBytes(t *testing.T) {
	text, readErr := os.ReadFile(filepath.Join(repo, "demo-parity.json"))
	if readErr != nil {
		t.Fatalf("read demo-parity.json: %v", readErr)
	}
	var parity struct {
		InputSHA256 map[string]string `json:"input_sha256"`
	}
	if parseErr := json.Unmarshal(text, &parity); parseErr != nil {
		t.Fatalf("parse demo-parity.json: %v", parseErr)
	}
	var listed, onDisk []string
	for name := range parity.InputSHA256 {
		listed = append(listed, name)
	}
	for _, dir := range []string{"samples", "policy"} {
		entries, dirErr := os.ReadDir(filepath.Join(repo, dir))
		if dirErr != nil {
			t.Fatalf("read %s: %v", dir, dirErr)
		}
		for _, entry := range entries {
			onDisk = append(onDisk, dir+"/"+entry.Name())
		}
	}
	sort.Strings(listed)
	sort.Strings(onDisk)
	if !reflect.DeepEqual(listed, onDisk) {
		t.Fatalf("demo-parity.json lists %v; samples and policy hold %v", listed, onDisk)
	}
	for _, name := range listed {
		sum := sha256.Sum256(readFile(filepath.Join(repo, name)))
		if got := hex.EncodeToString(sum[:]); got != parity.InputSHA256[name] {
			t.Fatalf("%s reads as %s, demo-parity.json names %s", name, got, parity.InputSHA256[name])
		}
	}
}
