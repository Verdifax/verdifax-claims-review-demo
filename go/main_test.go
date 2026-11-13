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

// The deleted-step cases every language runs, from one file (R250).
type deletedStepCases struct {
	JobID  string                     `json:"job_id"`
	Events []verdifax.SignalEvent     `json:"events"`
	Want   map[string]json.RawMessage `json:"want"`
}

func loadDeletedStepCases(t *testing.T) deletedStepCases {
	t.Helper()
	text, readErr := os.ReadFile(filepath.Join(repo, "testdata", "deleted-step-cases.json"))
	if readErr != nil {
		t.Fatal(readErr)
	}
	var cases deletedStepCases
	if parseErr := json.Unmarshal(text, &cases); parseErr != nil {
		t.Fatal(parseErr)
	}
	return cases
}

// parityConstants reads the constants demo-parity.json names.
func parityConstants(t *testing.T) map[string]any {
	t.Helper()
	text, readErr := os.ReadFile(filepath.Join(repo, "demo-parity.json"))
	if readErr != nil {
		t.Fatal(readErr)
	}
	var parity struct {
		Constants map[string]any `json:"constants"`
	}
	if parseErr := json.Unmarshal(text, &parity); parseErr != nil {
		t.Fatal(parseErr)
	}
	return parity.Constants
}

func sameJSON(t *testing.T, a, b []byte) bool {
	t.Helper()
	var x, y any
	if json.Unmarshal(a, &x) != nil || json.Unmarshal(b, &y) != nil {
		return false
	}
	return reflect.DeepEqual(x, y)
}

func TestADeletedStepIsReadFromItsJobsRunDeletedEventsOnly(t *testing.T) {
	cases := loadDeletedStepCases(t)
	got := deletedSteps(cases.Events, cases.JobID)
	if len(got) != len(cases.Want) {
		t.Fatalf("deleted steps %v, want %v", got, cases.Want)
	}
	for step, entry := range cases.Want {
		if !sameJSON(t, got[step], entry) {
			t.Fatalf("step %s: tombstone %s, want %s", step, got[step], entry)
		}
	}
	if other := deletedSteps(cases.Events, "claims-review-no-such-job"); len(other) != 0 {
		t.Fatalf("another job has no deleted steps here, got %v", other)
	}
	if parityConstants(t)["deleted_step_suffix"] != deletedStepSuffix {
		t.Fatalf("deletedStepSuffix %s differs from demo-parity.json", deletedStepSuffix)
	}
}

func TestADeletedStepsTombstoneIsWrittenBesideTheStepsAsTheServerSentIt(t *testing.T) {
	cases := loadDeletedStepCases(t)
	folder := t.TempDir()
	deleted := deletedSteps(cases.Events, cases.JobID)
	writeDeletedSteps(folder, deleted)
	files, _ := filepath.Glob(filepath.Join(folder, "steps", "*"))
	if len(files) != 2 || filepath.Base(files[0]) != "rule-02.deleted.json" || filepath.Base(files[1]) != "verdict.deleted.json" {
		t.Fatalf("steps/ holds %v, want rule-02.deleted.json and verdict.deleted.json", files)
	}
	for _, step := range []string{"rule-02", "verdict"} {
		written, readErr := os.ReadFile(filepath.Join(folder, "steps", step+deletedStepSuffix))
		if readErr != nil || string(written) != string(deleted[step])+"\n" {
			t.Fatalf("%s: the tombstone is not the server's exact bytes (%v):\n%s", step, readErr, written)
		}
		// The deleted step's sealed links are part of those bytes (round 61).
		if !bytes.Contains(written, []byte(`"job_links"`)) || !sameJSON(t, written, cases.Want[step]) {
			t.Fatalf("%s: the tombstone lacks the job_links the server sealed:\n%s", step, written)
		}
	}
}

func TestEveryPrintedRecheckOfASavedHistoryOrFeedStatesItsHeadAgeLimit(t *testing.T) {
	here, _ := os.Getwd()
	folder := t.TempDir()
	if chdirErr := os.Chdir(folder); chdirErr != nil {
		t.Fatal(chdirErr)
	}
	defer func() { _ = os.Chdir(here) }()
	must(os.MkdirAll(filepath.Join("records", "job-1"), 0o755))
	writeText(filepath.Join("records", "job-1", "signoff-history.json"), "[]")
	writeText("signals.json", "[]")
	reader, writer, pipeErr := os.Pipe()
	if pipeErr != nil {
		t.Fatal(pipeErr)
	}
	stdout := os.Stdout
	os.Stdout = writer
	printVerifyCommands("job-1")
	os.Stdout = stdout
	_ = writer.Close()
	var out bytes.Buffer
	_, _ = out.ReadFrom(reader)
	lines := strings.Split(strings.TrimRight(out.String(), "\n"), "\n")
	constants := parityConstants(t)
	if constants["min_verifier"] != minVerifier || lines[0] != fmt.Sprintf("verify with verdifax-verify %s or later:", minVerifier) {
		t.Fatalf("first line %q, min_verifier %v", lines[0], constants["min_verifier"])
	}
	if constants["saved_head_max_age"] != savedHeadMaxAge || constants["history_age_reason"] != historyAgeReason || constants["feed_age_reason"] != feedAgeReason {
		t.Fatal("the head age limit or its reasons differ from demo-parity.json")
	}
	want := map[string][2]string{
		"--signoff-history": {"  verdifax-verify --signoff-history " + filepath.Join("records", "job-1", "signoff-history.json") + " --signoff-head-max-age " + savedHeadMaxAge, "  " + historyAgeReason},
		"--signals":         {"  verdifax-verify --signals signals.json --feed-head-max-age " + savedHeadMaxAge, "  " + feedAgeReason},
	}
	for mode, pair := range want {
		found := false
		for i, line := range lines {
			if strings.Contains(line, mode) {
				found = line == pair[0] && i+1 < len(lines) && lines[i+1] == pair[1]
				break
			}
		}
		if !found {
			t.Fatalf("the %s re-check does not state its head age limit and reason:\n%s", mode, out.String())
		}
	}
}
