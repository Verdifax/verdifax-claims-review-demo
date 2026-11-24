package main

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math"
	"net/http"
	"net/http/httptest"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"sort"
	"strconv"
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

// R259: the SDK waits half a second before the first retry and doubles it to a
// cap of eight, so the shared budget must reach a minute of waiting. Round 61:
// the default of two retries waited about a second and a half, which a 503
// from the model API outlasted.
func TestModelRetryBudgetWaitsOutAnOutage(t *testing.T) {
	text, readErr := os.ReadFile(filepath.Join(repo, "demo-parity.json"))
	if readErr != nil {
		t.Fatalf("read demo-parity.json: %v", readErr)
	}
	var parity struct {
		Constants struct {
			ModelRetryTries string `json:"model_retry_tries"`
		} `json:"constants"`
	}
	if parseErr := json.Unmarshal(text, &parity); parseErr != nil {
		t.Fatalf("parse demo-parity.json: %v", parseErr)
	}
	named, convErr := strconv.Atoi(parity.Constants.ModelRetryTries)
	if convErr != nil || named != modelRetryTries {
		t.Fatalf("demo-parity.json names %q tries, the program holds %d (%v)", parity.Constants.ModelRetryTries, modelRetryTries, convErr)
	}
	waited := 0.0
	for i := 0; i < modelRetryTries-1; i++ {
		waited += math.Min(8, 0.5*math.Pow(2, float64(i)))
	}
	if waited < 60 {
		t.Fatalf("a budget of %d tries waits %.1f seconds, under the 60 a model outage needs", modelRetryTries, waited)
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

// Round 62: the cases every language runs for the every-capability demo.
func readCases(t *testing.T, name string, into any) {
	t.Helper()
	text, readErr := os.ReadFile(filepath.Join(repo, "testdata", name))
	if readErr != nil {
		t.Fatalf("read %s: %v", name, readErr)
	}
	if decodeErr := json.Unmarshal(text, into); decodeErr != nil {
		t.Fatalf("decode %s: %v", name, decodeErr)
	}
}

// asJSON is a value as generic JSON, so it compares with a case's want.
func asJSON(t *testing.T, value any) any {
	t.Helper()
	encoded, encodeErr := json.Marshal(value)
	if encodeErr != nil {
		t.Fatal(encodeErr)
	}
	var out any
	if decodeErr := json.Unmarshal(encoded, &out); decodeErr != nil {
		t.Fatal(decodeErr)
	}
	return out
}

func TestUsageIsDeclaredForEveryModelStep(t *testing.T) {
	var cases struct {
		Cases []struct {
			Name          string       `json:"name"`
			Usage         *usageReport `json:"usage"`
			Want          any          `json:"want"`
			WantCanonical *string      `json:"want_canonical"`
		} `json:"cases"`
	}
	readCases(t, "usage-cases.json", &cases)
	for _, c := range cases.Cases {
		got := usageSubstance(c.Usage)
		if got == nil {
			if c.Want != nil || c.WantCanonical != nil {
				t.Errorf("%s: no usage declared, want %v", c.Name, c.Want)
			}
			continue
		}
		if !reflect.DeepEqual(asJSON(t, got), c.Want) || c.WantCanonical == nil || canonical(got) != *c.WantCanonical {
			t.Errorf("%s: %s, want %v", c.Name, canonical(got), c.Want)
		}
	}
	var recorded struct {
		Recorded []struct {
			Name   string          `json:"name"`
			Answer json.RawMessage `json:"answer"`
			Want   struct {
				InputTokens  int64 `json:"input_tokens"`
				OutputTokens int64 `json:"output_tokens"`
			} `json:"want"`
		} `json:"recorded"`
	}
	readCases(t, "usage-cases.json", &recorded)
	if len(recorded.Recorded) == 0 {
		t.Fatal("no recorded usage cases")
	}
	for _, c := range recorded.Recorded {
		var a recordedAnswer
		if err := json.Unmarshal(c.Answer, &a); err != nil {
			t.Fatal(err)
		}
		got := ask(context.Background(), map[string]recordedAnswer{"rule-01": a}, "", "", "rule-01", "")
		if got.InputTokens != c.Want.InputTokens || got.OutputTokens != c.Want.OutputTokens {
			t.Errorf("%s: recorded usage %d %d, want %d %d", c.Name, got.InputTokens, got.OutputTokens, c.Want.InputTokens, c.Want.OutputTokens)
		}
	}
}

func TestTheDataChecksAreComputed(t *testing.T) {
	var cases struct {
		RuleSetSHA256 string `json:"rule_set_sha256"`
		AsOf          string `json:"as_of"`
		Cases         []struct {
			Name          string `json:"name"`
			Intake        string `json:"intake_draft_sha256"`
			Now           string `json:"draft_sha256_now"`
			RuleSetHash   string `json:"rule_set_hash"`
			Want          any    `json:"want"`
			WantCanonical string `json:"want_canonical"`
		} `json:"cases"`
	}
	readCases(t, "data-quality-cases.json", &cases)
	if cases.RuleSetSHA256 != ruleSetSHA256 {
		t.Fatalf("the pinned rule set hash is %s, the cases say %s", ruleSetSHA256, cases.RuleSetSHA256)
	}
	for _, c := range cases.Cases {
		got := asJSON(t, dataQuality(c.Intake, c.Now, c.RuleSetHash, cases.AsOf))
		if !reflect.DeepEqual(got, c.Want) || canonical(got) != c.WantCanonical {
			t.Errorf("%s: %s, want %s", c.Name, canonical(got), c.WantCanonical)
		}
	}
}

func TestTheLoadedRuleSetIsThePinnedOneAndThePolicyIsV5(t *testing.T) {
	if _, hash := loadRules(); hash != ruleSetSHA256 {
		t.Fatalf("the rule set hashes to %s, the pin is %s", hash, ruleSetSHA256)
	}
	type match struct {
		Value string `json:"value"`
	}
	var policy struct {
		Required   []string `json:"required_attested_fields"`
		WatchSetup struct {
			Fields []string `json:"fields"`
		} `json:"watch_setup"`
		Rules []struct {
			Conditions []struct {
				Type              string   `json:"type"`
				PurposeMatch      match    `json:"purpose_match"`
				CategoriesAllowed []string `json:"categories_allowed"`
				Facts             []struct {
					Fact  string `json:"fact"`
					Match match  `json:"match"`
				} `json:"facts"`
			} `json:"conditions"`
		} `json:"rules"`
	}
	if decodeErr := json.Unmarshal(readFile(filepath.Join(repo, "policy", policyFile)), &policy); decodeErr != nil {
		t.Fatal(decodeErr)
	}
	if policyFile != parityConstants(t)["policy_file"] || !reflect.DeepEqual(policy.Required, []string{"actor_id"}) {
		t.Fatalf("policy %s requires %v", policyFile, policy.Required)
	}
	if !reflect.DeepEqual(policy.WatchSetup.Fields, []string{"client_name", "client_version"}) {
		t.Errorf("the policy watches %v", policy.WatchSetup.Fields)
	}
	uses, approved := 0, []string{}
	for _, r := range policy.Rules {
		for _, c := range r.Conditions {
			switch c.Type {
			case "data_use":
				uses++
				if c.PurposeMatch.Value != dataUsePurpose || !reflect.DeepEqual(c.CategoriesAllowed, []string{dataUseCategory}) {
					t.Errorf("a rule allows purpose %s and categories %v", c.PurposeMatch.Value, c.CategoriesAllowed)
				}
			case "attested_setup":
				for _, f := range c.Facts {
					approved = append(approved, f.Fact+"="+f.Match.Value)
				}
			}
		}
	}
	if uses != 4 || !reflect.DeepEqual(approved, []string{"client_version=" + clientVersion}) {
		t.Fatalf("%d data use conditions, approved setups %v", uses, approved)
	}
}

func TestDataUseAgentVersionsAndToolLists(t *testing.T) {
	constants := parityConstants(t)
	wantUse := map[string]any{"purpose": constants["data_use_purpose"], "basis": constants["data_use_basis"],
		"categories": []any{constants["data_use_category"]}}
	if got := asJSON(t, dataUse(dataUseCategory)); !reflect.DeepEqual(got, wantUse) {
		t.Errorf("data use %v, want %v", got, wantUse)
	}
	if clientName != constants["client_name"] || clientVersion != constants["client_version"] {
		t.Errorf("client %s %s", clientName, clientVersion)
	}
	if !reflect.DeepEqual(asJSON(t, grantableTools), constants["grantable_tools"]) {
		t.Errorf("grantable tools %v", grantableTools)
	}
	type values struct {
		Valid   []string `json:"valid"`
		Invalid []string `json:"invalid"`
		Error   string   `json:"error"`
	}
	var c struct {
		Categories values `json:"data_use_categories"`
		Versions   values `json:"agent_versions"`
		ToolLists  struct {
			Valid   []string `json:"valid"`
			Invalid []struct {
				Tools string `json:"tools"`
				Error string `json:"error"`
			} `json:"invalid"`
		} `json:"tool_lists"`
	}
	readCases(t, "command-cases.json", &c)
	if len(c.Categories.Invalid) == 0 || len(c.Versions.Invalid) == 0 || len(c.ToolLists.Invalid) == 0 {
		t.Fatal("command-cases.json holds no data use, version or tool list cases")
	}
	for _, check := range []struct {
		cases values
		why   func(string) string
	}{{c.Categories, dataUseError}, {c.Versions, clientVersionError}} {
		for _, v := range check.cases.Valid {
			if got := check.why(v); got != "" {
				t.Errorf("%q refused: %s", v, got)
			}
		}
		for _, v := range check.cases.Invalid {
			if got := check.why(v); got != check.cases.Error {
				t.Errorf("%q: %q, want %q", v, got, check.cases.Error)
			}
		}
	}
	for _, tools := range c.ToolLists.Valid {
		if got := toolsError(strings.Split(tools, ",")); got != "" {
			t.Errorf("tools %q refused: %s", tools, got)
		}
	}
	for _, bad := range c.ToolLists.Invalid {
		if got := toolsError(strings.Split(bad.Tools, ",")); got == "" || !strings.HasPrefix(got, bad.Error) {
			t.Errorf("tools %q: %q, want %q", bad.Tools, got, bad.Error)
		}
	}
}

// Every step names its agent, the agent's declared name and version, and its
// data use, in the fields demo-parity.json lists.
func TestEveryStepDeclaresItsAgentVersionAndDataUse(t *testing.T) {
	var parity struct {
		Attested map[string][]string `json:"attested_context"`
		Policy   struct {
			Fields []string `json:"fields"`
		} `json:"policy_context"`
	}
	if decodeErr := json.Unmarshal(readFile(filepath.Join(repo, "demo-parity.json")), &parity); decodeErr != nil {
		t.Fatal(decodeErr)
	}
	keysOf := func(value any) []string {
		out := []string{}
		for k := range asJSON(t, value).(map[string]any) {
			out = append(out, k)
		}
		sort.Strings(out)
		return out
	}
	sorted := func(list []string) []string {
		out := append([]string{}, list...)
		sort.Strings(out)
		return out
	}
	intake := attestedContext("6-preview", nil)
	if intake["client_name"] != clientName || intake["client_version"] != "6-preview" {
		t.Errorf("the gate step declares %v", intake)
	}
	if got := keysOf(intake); !reflect.DeepEqual(got, sorted(parity.Attested["intake"])) {
		t.Errorf("the gate step's attested fields are %v", got)
	}
	model := attestedContext(clientVersion, map[string]any{"model_provider": "anthropic", "model_name": "m", "prompt_hash": "h",
		"decision_kind": "claims.list", "decision_result": "1"})
	if got := keysOf(model); !reflect.DeepEqual(got, sorted(parity.Attested["claims"])) {
		t.Errorf("a model step's attested fields are %v", got)
	}
	step := policyContext("intake", "INTERNAL", dataQuality("a", "a", ruleSetSHA256, "2026-11-24T00:00:00Z"), dataUse(dataUseCategory))
	if got := keysOf(step); !reflect.DeepEqual(got, sorted(parity.Policy.Fields)) {
		t.Errorf("the policy context's fields are %v", got)
	}
}

// A bad value is refused before anything is read or sent, an empty one
// included: no key file exists in the folder these run in.
func TestBadDeclarationsAreRefusedBeforeAnythingIsSent(t *testing.T) {
	for _, c := range []struct {
		args []string
		want string
	}{
		{[]string{"review", "d.md", "--data-use-category", "Draft"}, dataUseError("Draft")},
		{[]string{"review", "d.md", "--data-use-category", ""}, dataUseError("")},
		{[]string{"review", "d.md", "--agent-version", ""}, clientVersionError("")},
		{[]string{"tool-key", "sign_off"}, toolsError([]string{"sign_off"})},
		{[]string{"review", "d.md", "--agent-version", "6-preview", "--data-use-category", "customer.records"}, "no demo-keys.json here"},
	} {
		run := exec.Command(os.Args[0], append([]string{"-test.run=TestHelperRunsTheProgram", "--"}, c.args...)...)
		run.Dir = t.TempDir()
		run.Env = append(os.Environ(), "CLAIMS_REVIEW_RUN_MAIN=1", "CLAIMS_REVIEW_REPO="+repo, "CLAIMS_REVIEW_REGISTRY_RECORD_HASH=")
		out, runErr := run.CombinedOutput()
		var exitErr *exec.ExitError
		if !errors.As(runErr, &exitErr) || exitErr.ExitCode() != 1 || !strings.Contains(string(out), "claims-review: "+c.want) {
			t.Errorf("%v: want exit 1 with %q, got %v: %s", c.args, c.want, runErr, out)
		}
	}
}

// The files hold what the server sent, every field and number kept, even one
// this SDK does not name.
func TestSummaryAndResponseTimesKeepEveryFieldTheServerSent(t *testing.T) {
	hash := strings.Repeat("a", 64)
	summary := `{"z_later_field":12345678901234567890,"format":"verdifax.period-summary.v1","account":"account-7","head_sequence":4,"head_hash":"` + hash +
		`","head_statement_hash":"` + hash + `","figures":[],"summary_hash":"` + hash + `"}`
	head := `{"z_later_field":"kept","format":"verdifax.feed-head.v1","account":"account-7","sequence":4,"head_hash":"` + hash +
		`","stated_at":"2026-11-24T00:00:00Z","head_statement_hash":"` + hash + `"}`
	report := `{"z_later_field":12345678901234567890,"form":"verdifax.response-times.v1","feed":[{"b":1,"a":2}]}`
	var asked []string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		asked = append(asked, r.Method+" "+r.URL.RequestURI())
		switch r.Method + " " + r.URL.Path {
		case "POST /account/summaries":
			_, _ = io.WriteString(w, `{"id":3,"reused":false,"summary":`+summary+`}`)
		case "GET /account/summaries/3":
			_, _ = io.WriteString(w, summary)
		case "GET /account/summaries/3/head":
			_, _ = io.WriteString(w, head)
		case "GET /account/summaries/3/report.pdf":
			_, _ = io.WriteString(w, "%PDF-1.7\n\xff\x00binary")
		case "GET /account/signoffs/response-times":
			_, _ = io.WriteString(w, report)
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer server.Close()
	t.Setenv("VERDIFAX_BASE_URL", server.URL)
	t.Chdir(t.TempDir())
	_, signer, genErr := verdifax.GenerateRequestKey()
	if genErr != nil {
		t.Fatal(genErr)
	}
	writePrivate("demo-keys.json", canonical(map[string]any{"owner": map[string]any{"secret": "vfx_case", "signing_key": hex.EncodeToString(signer)}}))
	const start, end = "2026-11-01T00:00:00Z", "2026-11-02T00:00:00Z"
	cmdSummary(context.Background(), start, end)
	cmdResponseTimes(context.Background(), start, end)
	for file, want := range map[string]string{"summary.json": summary, "summary-head.json": head, "response-times.json": report} {
		var got bytes.Buffer
		if compactErr := json.Compact(&got, readFile(file)); compactErr != nil || got.String() != want {
			t.Errorf("%s holds %s, the server sent %s (%v)", file, got.String(), want, compactErr)
		}
	}
	if pdf := readFile("summary.pdf"); string(pdf) != "%PDF-1.7\n\xff\x00binary" {
		t.Errorf("summary.pdf holds %q", pdf)
	}
	if last := asked[len(asked)-1]; !strings.Contains(last, "pack=1") {
		t.Errorf("the report was asked for without its feed: %s", last)
	}
}

type commandCases struct {
	Settings struct {
		Valid   [][2]string `json:"valid"`
		Invalid []struct {
			Args  [2]string `json:"args"`
			Error string    `json:"error"`
		} `json:"invalid"`
	} `json:"settings"`
	ApproverNames struct {
		Valid   []string `json:"valid"`
		Invalid []string `json:"invalid"`
		Error   string   `json:"error"`
		Files   struct {
			Name       string `json:"name"`
			PrivateKey string `json:"private_key"`
			PublicKey  string `json:"public_key"`
		} `json:"files"`
	} `json:"approver_names"`
	DeleteStep struct {
		JobID            string                 `json:"job_id"`
		Step             string                 `json:"step"`
		Record           json.RawMessage        `json:"record"`
		Feed             []verdifax.SignalEvent `json:"feed"`
		WantTombstone    any                    `json:"want_tombstone"`
		WantFiles        []string               `json:"want_files"`
		WantRemoved      []string               `json:"want_removed"`
		MissingStepError string                 `json:"missing_step_error"`
	} `json:"delete_step"`
	RequestKey struct {
		Keys        map[string]any `json:"keys"`
		RotatedTo   string         `json:"rotated_to"`
		AfterRotate map[string]any `json:"after_rotate"`
		AfterClear  map[string]any `json:"after_clear"`
		Holder      struct {
			Keys        map[string]any `json:"keys"`
			RotatedTo   string         `json:"rotated_to"`
			AfterRotate map[string]any `json:"after_rotate"`
			AfterClear  map[string]any `json:"after_clear"`
		} `json:"holder"`
	} `json:"request_key"`
}

func TestTheRegistryRecordAReviewClaims(t *testing.T) {
	var c struct {
		RegistryRecord struct {
			DefaultSource string   `json:"default_source"`
			Given         string   `json:"given"`
			Invalid       []string `json:"invalid"`
			Error         string   `json:"error"`
		} `json:"registry_record"`
	}
	readCases(t, "command-cases.json", &c)
	r := c.RegistryRecord
	if registryRecordHash != sha256Hex([]byte(r.DefaultSource)) {
		t.Fatalf("the fixed record is %s", registryRecordHash)
	}
	if got, msg := registryRecord(""); got != registryRecordHash || msg != "" {
		t.Errorf("unset: %q %q", got, msg)
	}
	if got, msg := registryRecord(r.Given); got != r.Given || msg != "" {
		t.Errorf("given: %q %q", got, msg)
	}
	for _, bad := range r.Invalid {
		if got, msg := registryRecord(bad); got != "" || msg != r.Error {
			t.Errorf("%q: %q %q, want the refusal %q", bad, got, msg, r.Error)
		}
	}
}

func TestTheRound62CommandsDecideAlike(t *testing.T) {
	var c commandCases
	readCases(t, "command-cases.json", &c)
	for _, v := range c.Settings.Valid {
		if got := settingError(v[0], v[1]); got != "" {
			t.Errorf("settings %v refused: %s", v, got)
		}
	}
	for _, v := range c.Settings.Invalid {
		if got := settingError(v.Args[0], v.Args[1]); got != v.Error {
			t.Errorf("settings %v: %q, want %q", v.Args, got, v.Error)
		}
	}
	for _, name := range c.ApproverNames.Valid {
		if approverNameError(name) != "" {
			t.Errorf("approver name %q refused", name)
		}
	}
	for _, name := range c.ApproverNames.Invalid {
		if approverNameError(name) != c.ApproverNames.Error {
			t.Errorf("approver name %q accepted", name)
		}
	}
	if private, public := approverFiles(c.ApproverNames.Files.Name); private != c.ApproverNames.Files.PrivateKey || public != c.ApproverNames.Files.PublicKey {
		t.Errorf("approver files %s %s", private, public)
	}
	rotated := c.RequestKey.RotatedTo
	if got := withAgentSigningKey(c.RequestKey.Keys, &rotated); !reflect.DeepEqual(asJSON(t, got), asJSON(t, c.RequestKey.AfterRotate)) {
		t.Errorf("after rotate: %v", got)
	}
	if got := withAgentSigningKey(c.RequestKey.Keys, nil); !reflect.DeepEqual(asJSON(t, got), asJSON(t, c.RequestKey.AfterClear)) {
		t.Errorf("after clear: %v", got)
	}
	h := c.RequestKey.Holder
	if got := withSigningKey(h.Keys, "holder", &h.RotatedTo); !reflect.DeepEqual(asJSON(t, got), asJSON(t, h.AfterRotate)) {
		t.Errorf("holder after rotate: %v", got)
	}
	if got := withSigningKey(h.Keys, "holder", nil); !reflect.DeepEqual(asJSON(t, got), asJSON(t, h.AfterClear)) {
		t.Errorf("holder after clear: %v", got)
	}
	var withHolder keyFile
	encoded, encodeErr := json.Marshal(h.Keys)
	if encodeErr != nil || json.Unmarshal(encoded, &withHolder) != nil {
		t.Fatal("the holder case is not a keys file")
	}
	if role, id, _, _ := keyFor(withHolder, true); role != "holder" || id != 9 {
		t.Errorf("--holder acts on %s %d, want the holder, key 9", role, id)
	}
	if role, id, _, _ := keyFor(withHolder, false); role != "agent" || id != 7 {
		t.Errorf("without --holder %s %d, want the agent, key 7", role, id)
	}
	agent := c.RequestKey.AfterClear["agent"].(map[string]any)
	if clientFor(agent["secret"].(string), "") == nil {
		t.Error("a key whose request key was cleared must still make a client, which sends unsigned")
	}
}

func TestADeletedStepIsReplacedByItsTombstone(t *testing.T) {
	var c commandCases
	readCases(t, "command-cases.json", &c)
	d := c.DeleteStep
	tombstone, found := deletedSteps(d.Feed, d.JobID)[d.Step]
	if !found || !reflect.DeepEqual(asJSON(t, tombstone), d.WantTombstone) {
		t.Fatalf("tombstone %s", tombstone)
	}
	folder := t.TempDir()
	if mkErr := os.MkdirAll(filepath.Join(folder, "steps"), 0o755); mkErr != nil {
		t.Fatal(mkErr)
	}
	if writeErr := os.WriteFile(filepath.Join(folder, "steps", d.Step+".json"), d.Record, 0o644); writeErr != nil {
		t.Fatal(writeErr)
	}
	replaceWithTombstone(folder, d.Step, tombstone)
	for _, f := range d.WantFiles {
		var got any
		if decodeErr := json.Unmarshal(readFile(filepath.Join(folder, f)), &got); decodeErr != nil || !reflect.DeepEqual(got, d.WantTombstone) {
			t.Errorf("%s: %v %v", f, got, decodeErr)
		}
	}
	for _, f := range d.WantRemoved {
		if _, statErr := os.Stat(filepath.Join(folder, f)); statErr == nil {
			t.Errorf("%s was not removed", f)
		}
	}
}

// TestHelperRunsTheProgram runs main with the arguments after "--" when the
// test binary is started by a test below; otherwise it does nothing.
func TestHelperRunsTheProgram(t *testing.T) {
	if os.Getenv("CLAIMS_REVIEW_RUN_MAIN") != "1" {
		return
	}
	for i, a := range os.Args {
		if a == "--" {
			os.Args = append([]string{"claims-review"}, os.Args[i+1:]...)
			main()
			os.Exit(0)
		}
	}
}

func TestDeleteStepNeedsTheSealedRecord(t *testing.T) {
	var c commandCases
	readCases(t, "command-cases.json", &c)
	run := exec.Command(os.Args[0], "-test.run=TestHelperRunsTheProgram", "--", "delete-step", c.DeleteStep.JobID, "rule-03")
	run.Dir = t.TempDir()
	run.Env = append(os.Environ(), "CLAIMS_REVIEW_RUN_MAIN=1", "CLAIMS_REVIEW_REPO="+repo)
	out, runErr := run.CombinedOutput()
	var exitErr *exec.ExitError
	if !errors.As(runErr, &exitErr) || exitErr.ExitCode() != 1 || !strings.Contains(string(out), c.DeleteStep.MissingStepError) {
		t.Fatalf("want exit 1 naming the missing record, got %v: %s", runErr, out)
	}
}

func TestPublicationSkipsADeletedStep(t *testing.T) {
	var c struct {
		Publication struct {
			Targets    map[string]int64 `json:"targets"`
			Tombstones []string         `json:"tombstones"`
			Want       []string         `json:"want"`
		} `json:"publication"`
	}
	readCases(t, "command-cases.json", &c)
	folder := t.TempDir()
	if mkErr := os.MkdirAll(filepath.Join(folder, "steps"), 0o755); mkErr != nil {
		t.Fatal(mkErr)
	}
	for _, step := range c.Publication.Tombstones {
		if writeErr := os.WriteFile(filepath.Join(folder, "steps", step+deletedStepSuffix), []byte("{}"), 0o644); writeErr != nil {
			t.Fatal(writeErr)
		}
	}
	if got := publishable(folder, c.Publication.Targets); !reflect.DeepEqual(got, c.Publication.Want) {
		t.Fatalf("publishable = %v, want %v", got, c.Publication.Want)
	}
}

func TestTheAccountPolicyCheckPassesTheSignOffsOnceGiven(t *testing.T) {
	var c struct {
		VerifyCommands struct {
			JobID           string   `json:"job_id"`
			WithoutSignoffs []string `json:"without_signoffs"`
			WithSignoffs    []string `json:"with_signoffs"`
		} `json:"verify_commands"`
	}
	readCases(t, "command-cases.json", &c)
	var parity struct {
		Constants map[string]any `json:"constants"`
	}
	if decodeErr := json.Unmarshal(readFile(filepath.Join(repo, "demo-parity.json")), &parity); decodeErr != nil {
		t.Fatal(decodeErr)
	}
	if parity.Constants["signoff_pending_reason"] != signoffPendingReason {
		t.Fatal("signoffPendingReason differs from demo-parity.json")
	}
	policy := filepath.Join(repo, "policy", policyFile)
	want := func(lines []string) string {
		out := []string{}
		for _, l := range lines {
			out = append(out, strings.ReplaceAll(strings.ReplaceAll(l, "POLICY", policy), "SIGNOFF_PENDING_REASON", signoffPendingReason))
		}
		return strings.Join(out, "\n")
	}
	printed := func() string {
		r, w, pipeErr := os.Pipe()
		if pipeErr != nil {
			t.Fatal(pipeErr)
		}
		stdout := os.Stdout
		os.Stdout = w
		printVerifyCommands(c.VerifyCommands.JobID)
		os.Stdout = stdout
		_ = w.Close()
		raw, readErr := io.ReadAll(r)
		if readErr != nil {
			t.Fatal(readErr)
		}
		kept := []string{}
		for _, l := range strings.Split(string(raw), "\n") {
			if strings.Contains(l, "--account-policy ") || strings.HasPrefix(l, "  verdifax-verify --signoff ") || l == "  "+signoffPendingReason {
				kept = append(kept, l)
			}
		}
		return strings.Join(kept, "\n")
	}
	here, wdErr := os.Getwd()
	if wdErr != nil {
		t.Fatal(wdErr)
	}
	dir := t.TempDir()
	if chErr := os.Chdir(dir); chErr != nil {
		t.Fatal(chErr)
	}
	defer func() { _ = os.Chdir(here) }()
	if mkErr := os.MkdirAll(filepath.Join("records", c.VerifyCommands.JobID, "steps"), 0o755); mkErr != nil {
		t.Fatal(mkErr)
	}
	if got := printed(); got != want(c.VerifyCommands.WithoutSignoffs) {
		t.Fatalf("without sign-offs:\n%s", got)
	}
	if writeErr := os.WriteFile(filepath.Join("records", c.VerifyCommands.JobID, "signoffs.json"), []byte("[]"), 0o644); writeErr != nil {
		t.Fatal(writeErr)
	}
	if got := printed(); got != want(c.VerifyCommands.WithSignoffs) {
		t.Fatalf("with sign-offs:\n%s", got)
	}
}

// Recorded answers are allowed only against a server on this machine or,
// from a container, its host; every other server is live.
func TestRecordedAnswersOnlyAgainstALocalServer(t *testing.T) {
	var c struct {
		LocalServer struct {
			Local []string `json:"local"`
			Live  []string `json:"live"`
			Error string   `json:"error"`
		} `json:"local_server"`
	}
	readCases(t, "command-cases.json", &c)
	for _, u := range c.LocalServer.Local {
		t.Setenv("VERDIFAX_BASE_URL", u)
		if !isLocalServer() {
			t.Errorf("%s is local", u)
		}
	}
	for _, u := range c.LocalServer.Live {
		t.Setenv("VERDIFAX_BASE_URL", u)
		if isLocalServer() {
			t.Errorf("%s is live", u)
		}
	}
	src, err := os.ReadFile("main.go")
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(src), "fail(\""+c.LocalServer.Error+"\")") {
		t.Errorf("a live server must end the program with %q", c.LocalServer.Error)
	}
}
