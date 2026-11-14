// claims-review: a small agent that checks a draft against 15 general writing
// rules, with every step sealed by Verdifax.
//
// A demo, not customer proof. A sealed record shows which rules were checked
// and what the model answered; it does not show the model's judgment was
// right. The 15 rules are general writing rules, not legal advice and not a
// compliance standard. The account policy judges what each step declared; it
// cannot show the declared facts are true. A sign-off shows a registered
// approver key signed that decision; it does not show who holds the key or
// that the review was careful.
// A decision means the same on every result: approve agrees with
// the result signed, a verdict or the stop a receipt records, and reject
// judges it wrong. A sign-off releases or blocks nothing on its own.
//
// Commands (the same in every language of this repository):
//
//	claims-review setup                              owner key in VERDIFAX_OWNER_KEY
//	claims-review review DRAFT.md [--model ID]        model key in ANTHROPIC_API_KEY
//	claims-review review DRAFT.md --without-actor     the gate step names no agent
//	claims-review sign-off JOB_ID approve|reject [--note TEXT] [--approver NAME]
//	claims-review sign-off --receipt RUN_ID approve|reject [--note TEXT] [--approver NAME]
//	claims-review publish JOB_ID | publish --receipt RUN_ID
//	claims-review unpublish JOB_ID | unpublish --receipt RUN_ID
//	claims-review approver-add NAME | approver-revoke NAME
//	claims-review settings owner-replacement|verdifax-deletion allowed|blocked
//	claims-review request-key-register KEY_FILE               the key Verdifax made for the chapter
//	claims-review request-key-rotate | request-key-clear [--holder]   the agent key's request key, or the holder's
//	claims-review key-revoke [--holder]                       the owner revokes the agent key, or the holder
//	claims-review delete-step JOB_ID STEP                     the owner deletes one step
//	claims-review webhook-subscribe URL | webhook-list | webhook-remove
//	claims-review feed
//	claims-review verify-commands JOB_ID
//
// Environment: VERDIFAX_BASE_URL (default https://api.verdifax.com).
// Files, in the working folder: demo-keys.json and approver-key.json (0600),
// approver.txt, account-history.json, records/<job_id>/*.json with payloads/,
// signals.json; approver-<name>-key.json (0600), approver-<name>.txt and
// approvers.json; account-keys.json; webhook.json (0600), webhooks.json and
// webhook-deliveries.json.
package main

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"

	verdifax "github.com/Verdifax/verdifax-sdk-go"
	"github.com/anthropics/anthropic-sdk-go/v2"
	"github.com/anthropics/anthropic-sdk-go/v2/option"
)

const (
	defaultBaseURL = "https://api.verdifax.com"
	defaultModel   = "claude-fable-5-1"
	form           = "verdifax.claims-review.v1"
	ruleSet        = "v3"
	actorID        = "claims-review-agent"
	routeIntake    = "claims-review/v1/intake"
	routeClaims    = "claims-review/v1/claims"
	routeRule      = "claims-review/v1/rule-check"
	routeVerdict   = "claims-review/v1/verdict"
	objectPrefix   = "claims-review/v1/"
	minVerifier    = "v0.35.0"
	// Anthropic's SDKs retry a dropped connection, 408, 409, 429 and any
	// 5xx, but only twice by default, waiting half a second then one
	// second. That is too short: on November 13, 2026 a 503 from the
	// model API outlasted it and ended this program after twelve sealed
	// steps, failing a readiness run. modelRetryTries sets the number of
	// tries for one model call; the SDK's wait grows to a cap of eight
	// seconds, so an outage of about two minutes is waited out. A retry
	// can bill a second model call when the first answer was lost on the
	// way back; each step's measured cost shows it. All four languages
	// hold the same number (demo-parity.json, R259).
	modelRetryTries = 21
	// A stopped step's records live in records/stopped-<run_id>/, and the
	// model's exact answers in answers/ beside each job's steps/.
	stoppedFolderPrefix = "stopped-"
	answersFolder       = "answers"
	// The text the server's AI review checked for each step, as the server
	// returned it once: a reader re-runs the review's check from these files.
	tracesFolder = "traces"
	// The account policy setup activates; v1 to v3 stay for records sealed
	// under them. v3 requires a reviewed AI answer on every step after the
	// gate; v4 adds that every step names its agent
	// (required_attested_fields), that each verdict needs a person's
	// sign-off, and that every step after the gate passes its declared data
	// checks.
	policyFile = "claims-review-v4.json"
	// The rule set's pinned hash, which the rule_set_pinned data check
	// compares.
	ruleSetSHA256 = "a2afb580a203ccbd084645cfbbce532dd7eaf1d2dbeaa49fd008f66a91f98ae7"
	// The source every step's data checks name.
	dataQualitySource = "claims-review draft and rule set"
	// The AI's work on the steps after the gate: listing the draft's claims,
	// and explaining the verdict code decided.
	claimsPrompt  = "claims-prompt-v1.txt"
	verdictPrompt = "verdict-prompt-v1.txt"
	// A step whose record was deleted is written as its tombstone,
	// steps/<step>.deleted.json: the run.deleted entry from the account
	// feed, which verdifax-verify --job reads in the step's place (R250).
	deletedStepSuffix = ".deleted.json"
	// A saved sign-off history or feed is checked as of when it was
	// fetched, so every printed re-check states the limit on its head
	// statement's age (R246).
	savedHeadMaxAge  = "168h"
	historyAgeReason = "(168h: a saved history is checked as of when it was fetched; raise the limit for an older copy, or run sign-off again to fetch it fresh)"
	// signoffPendingReason is printed under the account-policy check of a
	// verdict with no sign-off yet: policy v4 requires one, so the verifier
	// exits 4 until it is given.
	signoffPendingReason = "(exit 4, ANCHORED, SIGN-OFF NOT SHOWN, until a sign-off is given: the policy requires one on the verdict; run sign-off, then verify-commands again)"
	feedAgeReason        = "(168h: a saved feed is checked as of when it was fetched; raise the limit for an older copy, or run feed again to fetch it fresh)"
)

// The data checks every step declares in policy_context.data_quality, each
// computed: draft_unchanged (the draft re-read before the step has the hash
// sealed at intake) and rule_set_pinned (the rule set's hash is the pin).
var dataQualityChecks = []string{"draft_unchanged", "rule_set_pinned"}

// The owner's two blocks on Verdifax, and their values.
var (
	accountSettings = []string{"owner-replacement", "verdifax-deletion"}
	settingValues   = []string{"allowed", "blocked"}
	approverName    = regexp.MustCompile(`^[a-z0-9-]{1,40}$`)
)

var levels = map[string]bool{"PUBLIC": true, "INTERNAL": true, "CONFIDENTIAL": true, "RESTRICTED": true, "TOP_SECRET": true, "TS_SCI": true}

var (
	repo               = repoDir()
	programID          = sha256Hex([]byte("verdifax-claims-review-demo/program/v1"))
	registryRecordHash = sha256Hex([]byte("verdifax-claims-review-demo/registry/v1"))
	classificationLine = regexp.MustCompile(`^<!--\s*classification:\s*([A-Z_]+)\s*-->$`)
)

func sha256Hex(data []byte) string {
	sum := sha256.Sum256(data)
	return hex.EncodeToString(sum[:])
}

// canonical writes JSON with sorted keys and no spaces, as every language of
// this repository writes it. Go's encoder sorts map keys and escapes <, > and
// & by default, which the other languages do not, so HTML escaping is off.
func canonical(value any) string { // any: JSON of any shape, as the SDK takes it
	var out strings.Builder
	enc := json.NewEncoder(&out)
	enc.SetEscapeHTML(false)
	if encodeErr := enc.Encode(value); encodeErr != nil {
		fail("canonical JSON: " + encodeErr.Error())
	}
	return strings.TrimSuffix(out.String(), "\n")
}

func repoDir() string {
	exe, exeErr := os.Executable()
	candidates := []string{}
	if exeErr == nil {
		candidates = append(candidates, filepath.Join(filepath.Dir(exe), ".."))
	}
	if wd, wdErr := os.Getwd(); wdErr == nil {
		candidates = append(candidates, filepath.Join(wd, ".."), wd)
	}
	if fromEnv := os.Getenv("CLAIMS_REVIEW_REPO"); fromEnv != "" {
		candidates = append([]string{fromEnv}, candidates...)
	}
	for _, c := range candidates {
		if _, statErr := os.Stat(filepath.Join(c, "rules", ruleSet+".json")); statErr == nil {
			abs, _ := filepath.Abs(c)
			return abs
		}
	}
	fail("cannot find the repository's rules folder; set CLAIMS_REVIEW_REPO to the repository")
	return ""
}

func baseURL() string {
	if v := os.Getenv("VERDIFAX_BASE_URL"); v != "" {
		return strings.TrimRight(v, "/")
	}
	return defaultBaseURL
}

func isProduction() bool { return baseURL() == defaultBaseURL }

// failOut and exit are where fail writes and how it stops; the tests
// replace them to read the message the program would print.
var (
	failOut io.Writer = os.Stderr
	exit              = os.Exit
)

func fail(message string) {
	fmt.Fprintln(failOut, "claims-review:", message)
	exit(1)
}

func must(err error) {
	if err != nil {
		fail(err.Error())
	}
}

func readFile(path string) []byte {
	data, readErr := os.ReadFile(path)
	must(readErr)
	return data
}

func writePrivate(path string, text string) {
	must(os.WriteFile(path, []byte(text), 0o600))
}

func writeText(path string, text string) {
	must(os.MkdirAll(filepath.Dir(path), 0o755))
	must(os.WriteFile(path, []byte(text), 0o644))
}

func writeJSON(path string, value any) {
	data, encodeErr := json.MarshalIndent(value, "", "  ")
	must(encodeErr)
	writeText(path, string(data)+"\n")
}

type rule struct {
	ID   string `json:"id"`
	Text string `json:"text"`
}

func loadRules() ([]rule, string) {
	var doc map[string]any
	must(json.Unmarshal(readFile(filepath.Join(repo, "rules", ruleSet+".json")), &doc))
	var typed struct {
		Rules []rule `json:"rules"`
	}
	must(json.Unmarshal(readFile(filepath.Join(repo, "rules", ruleSet+".json")), &typed))
	return typed.Rules, sha256Hex([]byte(canonical(doc)))
}

func loadPrompt() string { return loadPromptFile("prompt-v1.txt") }

func loadPromptFile(name string) string { return string(readFile(filepath.Join(repo, "rules", name))) }

func draftClassification(text string) string {
	first := strings.TrimSpace(strings.SplitN(text, "\n", 2)[0])
	level := "INTERNAL"
	if m := classificationLine.FindStringSubmatch(first); m != nil {
		level = m[1]
	}
	if !levels[level] {
		fail("the draft declares an unknown classification: " + level)
	}
	return level
}

type keyFile struct {
	BaseURL string `json:"base_url"`
	Owner   struct {
		Secret     string `json:"secret"`
		SigningKey string `json:"signing_key"`
	} `json:"owner"`
	Agent struct {
		ID         int64  `json:"id"`
		Secret     string `json:"secret"`
		SigningKey string `json:"signing_key"`
	} `json:"agent"`
	Feed struct {
		ID     int64  `json:"id"`
		Secret string `json:"secret"`
	} `json:"feed"`
	// Holder is the key Verdifax made at production tier for the
	// request-key chapter, once request-key-register has kept it.
	Holder *struct {
		ID         int64  `json:"id"`
		Secret     string `json:"secret"`
		SigningKey string `json:"signing_key"`
	} `json:"holder,omitempty"`
	PolicyDocumentHash string `json:"policy_document_hash"`
}

func loadKeys() keyFile {
	if _, statErr := os.Stat("demo-keys.json"); statErr != nil {
		fail("no demo-keys.json here; run `claims-review setup` first")
	}
	var keys keyFile
	must(json.Unmarshal(readFile("demo-keys.json"), &keys))
	return keys
}

// clientFor is a client for a key; one whose request key was cleared sends
// unsigned.
func clientFor(secret, signingHex string) *verdifax.Client {
	config := verdifax.Config{BaseURL: baseURL(), APIKey: secret}
	if signingHex != "" {
		key, decodeErr := hex.DecodeString(signingHex)
		must(decodeErr)
		config.SigningKey = verdifax.SigningKey(key)
	}
	client, newErr := verdifax.NewClient(config)
	must(newErr)
	return client
}

// usageReport is the model's own usage report for one call.
type usageReport struct {
	InputTokens  int64 `json:"input_tokens"`
	OutputTokens int64 `json:"output_tokens"`
}

// usageSubstance is the request_substance a step declares: the model's own
// usage report for the step's model call, so the server's budgets count real
// tokens. The gate step calls no model and declares none (nil).
func usageSubstance(usage *usageReport) map[string]any {
	if usage == nil {
		return nil
	}
	return map[string]any{"cost": map[string]any{"input_tokens": usage.InputTokens, "output_tokens": usage.OutputTokens,
		"total_tokens": usage.InputTokens + usage.OutputTokens}}
}

// dataQuality is the data checks a step declares, each computed (round 62).
func dataQuality(intakeDraftSHA256, draftSHA256Now, ruleSetHash, asOf string) *verdifax.DataQuality {
	result := func(ok bool) string {
		if ok {
			return "pass"
		}
		return "fail"
	}
	return &verdifax.DataQuality{Source: dataQualitySource, AsOf: asOf, Checks: []verdifax.DataCheck{
		{Name: dataQualityChecks[0], Result: result(draftSHA256Now == intakeDraftSHA256)},
		{Name: dataQualityChecks[1], Result: result(ruleSetHash == ruleSetSHA256)}}}
}

// settingError is why a settings command is refused before anything is sent,
// or "".
func settingError(setting, value string) string {
	if !containsString(accountSettings, setting) {
		return fmt.Sprintf("the setting must be %s, not %s", strings.Join(accountSettings, " or "), setting)
	}
	if !containsString(settingValues, value) {
		return fmt.Sprintf("the value must be %s, not %s", strings.Join(settingValues, " or "), value)
	}
	return ""
}

func containsString(list []string, want string) bool {
	for _, item := range list {
		if item == want {
			return true
		}
	}
	return false
}

func approverNameError(name string) string {
	if !approverName.MatchString(name) {
		return "an approver name is 1 to 40 lowercase letters, digits or dashes"
	}
	return ""
}

// approverFiles is a named approver's private key file (0600) and public key
// file.
func approverFiles(name string) (string, string) {
	return "approver-" + name + "-key.json", "approver-" + name + ".txt"
}

// A key that was ever on the trial tier must keep a request key, and the
// demo's accounts are trial accounts, so the agent key can replace its request
// key but never clear it. For the request-key chapter Verdifax makes one
// production-tier member key in the account, its creation sealed in the feed
// as Verdifax's: the holder. request-key-register keeps it in demo-keys.json,
// and --holder makes the request-key commands and key-revoke act on it.

// withSigningKey is the keys file, read as JSON, with one key's request key
// replaced, or removed when it was cleared (nil); role is "agent" or "holder".
func withSigningKey(keys map[string]any, role string, signingKeyHex *string) map[string]any {
	out := map[string]any{}
	for k, v := range keys {
		out[k] = v
	}
	entry := map[string]any{}
	if current, isMap := keys[role].(map[string]any); isMap {
		for k, v := range current {
			if k != "signing_key" {
				entry[k] = v
			}
		}
	}
	if signingKeyHex != nil {
		entry["signing_key"] = *signingKeyHex
	}
	out[role] = entry
	return out
}

// withAgentSigningKey is the keys file, read as JSON, with the agent key's
// request key replaced, or removed when it was cleared (nil).
func withAgentSigningKey(keys map[string]any, signingKeyHex *string) map[string]any {
	return withSigningKey(keys, "agent", signingKeyHex)
}

// keyFor is the key a request-key command or key-revoke acts on: the agent,
// or with --holder the key Verdifax made for the request-key chapter. It
// returns the role, the key's id, its secret and its request key.
func keyFor(keys keyFile, holder bool) (string, int64, string, string) {
	if !holder {
		return "agent", keys.Agent.ID, keys.Agent.Secret, keys.Agent.SigningKey
	}
	if keys.Holder == nil {
		fail("demo-keys.json holds no request-key holder; run request-key-register with the key Verdifax made first")
	}
	return "holder", keys.Holder.ID, keys.Holder.Secret, keys.Holder.SigningKey
}

// ── setup ──────────────────────────────────────────────────────────────────

func cmdSetup(ctx context.Context) {
	if _, statErr := os.Stat("demo-keys.json"); statErr == nil {
		fail("demo-keys.json exists here; this folder is already set up, and every program in this repository uses it")
	}
	ownerSecret := os.Getenv("VERDIFAX_OWNER_KEY")
	if ownerSecret == "" {
		fail("set VERDIFAX_OWNER_KEY to the account's owner key")
	}
	_, ownerSigner, genErr := verdifax.GenerateRequestKey()
	must(genErr)
	owner, ownerErr := verdifax.NewClient(verdifax.Config{BaseURL: baseURL(), APIKey: ownerSecret, SigningKey: ownerSigner})
	must(ownerErr)
	_, registerErr := owner.RegisterAccountRequestKey(ctx)
	must(registerErr)
	member, memberErr := owner.CreateAccountKey(ctx, &verdifax.CreateAccountKeyRequest{Name: "claims-review-agent"})
	must(memberErr)
	_, memberSigner, memberGenErr := verdifax.GenerateRequestKey()
	must(memberGenErr)
	agent, agentErr := verdifax.NewClient(verdifax.Config{BaseURL: baseURL(), APIKey: member.Secret, SigningKey: memberSigner})
	must(agentErr)
	_, agentRegisterErr := agent.RegisterRequestKey(ctx)
	must(agentRegisterErr)
	uploaded, uploadErr := owner.UploadAccountPolicy(ctx, string(readFile(filepath.Join(repo, "policy", policyFile))))
	must(uploadErr)
	_, activateErr := owner.ActivateAccountPolicy(ctx, uploaded.DocumentHash)
	must(activateErr)
	approverPub, approverKey, approverGenErr := verdifax.GenerateRequestKey()
	must(approverGenErr)
	_, approverErr := owner.RegisterApprover(ctx, &verdifax.RegisterApproverRequest{Name: "Claims review approver", ApproverKey: approverKey})
	must(approverErr)
	feed, feedErr := owner.CreateFeedKey(ctx, "claims-review-feed", nil)
	must(feedErr)
	history, historyErr := owner.GetAccountPolicyHistory(ctx)
	must(historyErr)
	writeText("account-history.json", string(history))
	writePrivate("approver-key.json", canonical(map[string]string{"private_key": hex.EncodeToString(approverKey)})+"\n")
	writeText("approver.txt", base64.StdEncoding.EncodeToString(approverPub))
	keys := map[string]any{
		"base_url":             baseURL(),
		"owner":                map[string]any{"secret": ownerSecret, "signing_key": hex.EncodeToString(ownerSigner)},
		"agent":                map[string]any{"id": member.ID, "secret": member.Secret, "signing_key": hex.EncodeToString(memberSigner)},
		"feed":                 map[string]any{"id": feed.ID, "secret": feed.Secret},
		"policy_document_hash": uploaded.DocumentHash,
	}
	data, encodeErr := json.MarshalIndent(keys, "", "  ")
	must(encodeErr)
	writePrivate("demo-keys.json", string(data)+"\n")
	fmt.Printf("setup done: agent key %d, feed key %d, policy %s\n", member.ID, feed.ID, uploaded.DocumentHash)
}

// ── the model ──────────────────────────────────────────────────────────────

type answer struct {
	Model, ResponseID, Verdict, Reason, Raw string
	InputTokens, OutputTokens               int64
}

type recordedAnswer struct {
	Model      string `json:"model"`
	ResponseID string `json:"response_id"`
	Verdict    string `json:"verdict"`
	Reason     string `json:"reason"`
	// Claims and Explanation are the recorded answers of the claims and
	// verdict steps.
	Claims      []string `json:"claims"`
	Explanation string   `json:"explanation"`
	// InputTokens and OutputTokens are the model use a recorded answer
	// declares, 0 when it names none: the budget stops are rehearsed against
	// a local server with answers that declare large use. Never part of the
	// answer.
	InputTokens  int64 `json:"input_tokens"`
	OutputTokens int64 `json:"output_tokens"`
}

// structured is one structured answer from the model: its exact text, which
// the step sends as its AI output, and where it came from.
type structured struct {
	Model, ResponseID, Raw    string
	InputTokens, OutputTokens int64
}

// askStructured asks the model for one answer in the JSON schema given. With
// recorded answers it returns recordedRaw instead. When the model declines,
// Raw is declined, so the step still carries the model's answer to review.
func askStructured(ctx context.Context, recorded *recordedAnswer, recordedRaw, prompt, userText, model string, maxTokens int64,
	schema map[string]any, declined string) (structured, bool) {
	if recorded != nil {
		out := structured{Model: recorded.Model, ResponseID: recorded.ResponseID, Raw: recordedRaw,
			InputTokens: recorded.InputTokens, OutputTokens: recorded.OutputTokens}
		if out.Model == "" {
			out.Model = "recorded-answers"
		}
		if out.ResponseID == "" {
			out.ResponseID = "recorded"
		}
		return out, true
	}
	client := anthropic.NewClient(option.WithMaxRetries(modelRetryTries - 1))
	response, callErr := client.Messages.New(ctx, anthropic.MessageNewParams{
		Model:        anthropic.Model(model),
		MaxTokens:    maxTokens,
		System:       []anthropic.TextBlockParam{{Text: prompt}},
		Messages:     []anthropic.MessageParam{anthropic.NewUserMessage(anthropic.NewTextBlock(userText))},
		OutputConfig: anthropic.OutputConfigParam{Format: anthropic.JSONOutputFormatParam{Schema: schema}},
	})
	must(callErr)
	out := structured{Model: string(response.Model), ResponseID: response.ID,
		InputTokens: response.Usage.InputTokens, OutputTokens: response.Usage.OutputTokens}
	for _, block := range response.Content {
		if block.Type == "text" {
			out.Raw = block.Text
			break
		}
	}
	if response.StopReason == anthropic.StopReasonRefusal || out.Raw == "" {
		if out.Raw == "" {
			out.Raw = declined
		}
		return out, false
	}
	return out, true
}

// askClaims is the draft's claims, as the model listed them.
func askClaims(ctx context.Context, recorded map[string]recordedAnswer, prompt, userText, model string) (structured, []string) {
	var rec *recordedAnswer
	raw := ""
	if recorded != nil {
		a := recorded["claims"]
		rec, raw = &a, canonical(map[string]any{"claims": a.Claims})
	}
	schema := map[string]any{"type": "object", "properties": map[string]any{"claims": map[string]any{"type": "array",
		"items": map[string]any{"type": "string"}}}, "required": []string{"claims"}, "additionalProperties": false}
	out, ok := askStructured(ctx, rec, raw, prompt, userText, model, 4000, schema, "the model declined to list the claims")
	var parsed struct {
		Claims []string `json:"claims"`
	}
	if !ok || json.Unmarshal([]byte(out.Raw), &parsed) != nil || parsed.Claims == nil {
		return out, []string{}
	}
	return out, parsed.Claims
}

// askExplanation is the model's plain-words explanation of the verdict code
// decided.
func askExplanation(ctx context.Context, recorded map[string]recordedAnswer, prompt, userText, model string) (structured, string) {
	var rec *recordedAnswer
	raw := ""
	if recorded != nil {
		a := recorded["verdict"]
		rec, raw = &a, canonical(map[string]any{"explanation": a.Explanation})
	}
	schema := map[string]any{"type": "object", "properties": map[string]any{"explanation": map[string]any{"type": "string"}},
		"required": []string{"explanation"}, "additionalProperties": false}
	out, ok := askStructured(ctx, rec, raw, prompt, userText, model, 2000, schema, "the model declined to explain the verdict")
	var parsed struct {
		Explanation string `json:"explanation"`
	}
	if !ok || json.Unmarshal([]byte(out.Raw), &parsed) != nil {
		return out, ""
	}
	return out, parsed.Explanation
}

// loadAnswers reads recorded answers for tests, refused against production.
func loadAnswers(path string) map[string]recordedAnswer {
	if path == "" {
		return nil
	}
	if isProduction() {
		fail("--answers is for tests against a local server, never production")
	}
	var recorded map[string]recordedAnswer
	must(json.Unmarshal(readFile(path), &recorded))
	return recorded
}

func ask(ctx context.Context, recorded map[string]recordedAnswer, prompt, userText, ruleID, model string) answer {
	if recorded != nil {
		a := recorded[ruleID]
		if a.Model == "" {
			a.Model = "recorded-answers"
		}
		if a.ResponseID == "" {
			a.ResponseID = "recorded"
		}
		return answer{Model: a.Model, ResponseID: a.ResponseID, Verdict: a.Verdict, Reason: a.Reason,
			Raw:         canonical(map[string]string{"verdict": a.Verdict, "reason": a.Reason}),
			InputTokens: a.InputTokens, OutputTokens: a.OutputTokens}
	}
	client := anthropic.NewClient(option.WithMaxRetries(modelRetryTries - 1))
	schema := map[string]any{
		"type":                 "object",
		"properties":           map[string]any{"verdict": map[string]any{"type": "string"}, "reason": map[string]any{"type": "string"}},
		"required":             []string{"verdict", "reason"},
		"additionalProperties": false,
	}
	response, callErr := client.Messages.New(ctx, anthropic.MessageNewParams{
		Model:        anthropic.Model(model),
		MaxTokens:    4000,
		System:       []anthropic.TextBlockParam{{Text: prompt}},
		Messages:     []anthropic.MessageParam{anthropic.NewUserMessage(anthropic.NewTextBlock(userText))},
		OutputConfig: anthropic.OutputConfigParam{Format: anthropic.JSONOutputFormatParam{Schema: schema}},
	})
	must(callErr)
	out := answer{Model: string(response.Model), ResponseID: response.ID,
		InputTokens: response.Usage.InputTokens, OutputTokens: response.Usage.OutputTokens}
	for _, block := range response.Content {
		if block.Type == "text" {
			out.Raw = block.Text
			break
		}
	}
	var parsed struct {
		Verdict string `json:"verdict"`
		Reason  string `json:"reason"`
	}
	if response.StopReason == anthropic.StopReasonRefusal || json.Unmarshal([]byte(out.Raw), &parsed) != nil || parsed.Reason == "" {
		out.Verdict, out.Reason, out.Raw = "unanswered", "the model declined to answer", ""
		return out
	}
	out.Verdict, out.Reason = parsed.Verdict, parsed.Reason
	if out.Verdict != "pass" && out.Verdict != "fail" {
		out.Verdict = "unanswered"
	}
	return out
}

func measuredCost(model string, inputTokens, outputTokens int64) (float64, bool) {
	var table struct {
		Prices map[string]struct {
			Input  float64 `json:"input"`
			Output float64 `json:"output"`
		} `json:"prices"`
	}
	must(json.Unmarshal(readFile(filepath.Join(repo, "rules", "model-prices-v1.json")), &table))
	price, known := table.Prices[model]
	if !known {
		return 0, false
	}
	return float64(inputTokens)/1e6*price.Input + float64(outputTokens)/1e6*price.Output, true
}

// ── review ─────────────────────────────────────────────────────────────────

// policyContext is a step's policy context: its object, the draft's
// classification and the step's data checks, sent as policy_context's
// object_type, object_id, classification and data_quality.
func policyContext(objectID, level string, quality *verdifax.DataQuality) *verdifax.PolicyContext {
	return &verdifax.PolicyContext{ObjectType: "text", ObjectID: objectPrefix + objectID, Classification: level, DataQuality: quality}
}

func cmdReview(ctx context.Context, draftPath, model, answersPath string, withoutActor bool) {
	keys := loadKeys()
	agent := clientFor(keys.Agent.Secret, keys.Agent.SigningKey)
	rules, ruleSetHash := loadRules()
	prompt := loadPrompt()
	draft := string(readFile(draftPath))
	draftHash := sha256Hex([]byte(draft))
	level := draftClassification(draft)
	// The time the draft was read: the as_of of every step's data checks.
	readAt := time.Now().UTC().Format("2006-01-02T15:04:05Z")
	// quality is the data checks, computed again now, the draft re-read.
	quality := func() *verdifax.DataQuality {
		return dataQuality(draftHash, sha256Hex(readFile(draftPath)), ruleSetHash, readAt)
	}
	recorded := loadAnswers(answersPath)
	now := time.Now().UTC().Format("20060102T150405Z")
	jobID := fmt.Sprintf("claims-review-%s-%s", draftHash[:12], now)
	out := filepath.Join("records", jobID)
	payloads := map[string]any{}

	intakePayload := map[string]any{"form": form, "step": "intake", "draft_sha256": draftHash, "draft_bytes": len(draft),
		"rule_set": ruleSet, "rule_set_hash": ruleSetHash, "classification": level}
	job, jobErr := agent.Job(jobID)
	must(jobErr)
	// The gate step names its agent, as policy v4 requires of every step;
	// --without-actor leaves the name out on purpose, so the policy refuses
	// the draft here, before any model is called.
	var agentNamed map[string]any
	if !withoutActor {
		agentNamed = map[string]any{"attested": true, "actor_id": actorID}
	}
	intake, intakeErr := job.Step(ctx, "intake", &verdifax.AttestRequest{PayloadText: canonical(intakePayload), ProgramID: programID,
		RouteID: routeIntake, RegistryRecordHash: registryRecordHash, AttestedContext: agentNamed,
		PolicyContext: policyContext("intake", level, quality())}, verdifax.WithAfter())
	var apiErr *verdifax.APIError
	if errors.As(intakeErr, &apiErr) {
		// The account policy refused the draft at intake (for example a
		// confidential draft), so it never reached a model. The refusal is
		// sealed as a receipt; nothing else runs for this draft.
		stop(ctx, agent, jobID, "intake", map[string]int64{}, map[string]any{"intake": intakePayload}, map[string]string{}, draftPath, draftHash,
			apiErr, model, 0, 0, map[string]bool{}, map[string]string{})
		return
	}
	must(intakeErr)
	payloads["intake"] = intakePayload
	runIDs := map[string]int64{"intake": intake.RunID}

	var results []map[string]any
	var failed []string
	var inputTokens, outputTokens int64
	models := map[string]bool{}
	// The model's exact answer, the text Verdifax's AI review reads, kept as
	// sent so a reader recomputes the sealed ai_output_hash from the file.
	answerTexts := map[string]string{}
	// The text the server's AI review checked, returned once with each step.
	traceTexts := map[string]string{}

	// The gate passed, so the draft may reach a model. Every step from here
	// carries the model's answer, which Verdifax's AI review reads and the
	// record seals; the account policy refuses any step without one.
	claimsPromptText := loadPromptFile(claimsPrompt)
	claimsText := "Draft:\n" + draft
	claimsPromptHash := sha256Hex([]byte(claimsPromptText + "\n" + claimsText))
	listed, claims := askClaims(ctx, recorded, claimsPromptText, claimsText, model)
	inputTokens += listed.InputTokens
	outputTokens += listed.OutputTokens
	models[listed.Model] = true
	answerTexts["claims"] = listed.Raw
	claimsPayload := map[string]any{"form": form, "step": "claims", "draft_sha256": draftHash, "rule_set_hash": ruleSetHash,
		"model": listed.Model, "model_response_id": listed.ResponseID, "prompt_hash": claimsPromptHash, "claims": claims}
	claimsStep, claimsErr := job.Step(ctx, "claims", &verdifax.AttestRequest{PayloadText: canonical(claimsPayload), ProgramID: programID,
		RouteID: routeClaims, RegistryRecordHash: registryRecordHash, AIOutputText: listed.Raw,
		RequestSubstance: usageSubstance(&usageReport{InputTokens: listed.InputTokens, OutputTokens: listed.OutputTokens}),
		AttestedContext: map[string]any{"attested": true, "actor_id": actorID, "model_provider": "anthropic", "model_name": listed.Model,
			"prompt_hash": claimsPromptHash, "decision_kind": "claims.list", "decision_result": fmt.Sprint(len(claims))},
		PolicyContext: policyContext("claims", level, quality())}, verdifax.WithAfter("intake"))
	if errors.As(claimsErr, &apiErr) {
		payloads["claims"] = claimsPayload
		if apiErr.AivpTraceText != "" {
			traceTexts["claims"] = apiErr.AivpTraceText
		}
		stop(ctx, agent, jobID, "claims", runIDs, payloads, answerTexts, draftPath, draftHash, apiErr, model, inputTokens, outputTokens, models,
			traceTexts)
		return
	}
	must(claimsErr)
	if claimsStep.AivpTraceText != "" {
		traceTexts["claims"] = claimsStep.AivpTraceText
	}
	payloads["claims"] = claimsPayload
	runIDs["claims"] = claimsStep.RunID
	fmt.Printf("claims: %d listed\n", len(claims))

	ruleIDs := make([]string, 0, len(rules))
	for _, r := range rules {
		ruleIDs = append(ruleIDs, r.ID)
		userText := fmt.Sprintf("Rule %s: %s\n\nDraft:\n%s", r.ID, r.Text, draft)
		promptHash := sha256Hex([]byte(prompt + "\n" + userText))
		a := ask(ctx, recorded, prompt, userText, r.ID, model)
		inputTokens += a.InputTokens
		outputTokens += a.OutputTokens
		models[a.Model] = true
		payload := map[string]any{"form": form, "step": r.ID, "rule_id": r.ID, "rule_text": r.Text, "draft_sha256": draftHash,
			"rule_set_hash": ruleSetHash, "model": a.Model, "model_response_id": a.ResponseID, "prompt_hash": promptHash,
			"verdict": a.Verdict, "reason": a.Reason}
		aiText := a.Raw
		if aiText == "" {
			aiText = a.Reason
		}
		answerTexts[r.ID] = aiText
		receipt, stepErr := job.Step(ctx, r.ID, &verdifax.AttestRequest{PayloadText: canonical(payload), ProgramID: programID,
			RouteID: routeRule, RegistryRecordHash: registryRecordHash, AIOutputText: aiText,
			RequestSubstance: usageSubstance(&usageReport{InputTokens: a.InputTokens, OutputTokens: a.OutputTokens}),
			AttestedContext: map[string]any{"attested": true, "actor_id": actorID, "model_provider": "anthropic", "model_name": a.Model,
				"prompt_hash": promptHash, "decision_kind": "claims.rule_check", "decision_result": a.Verdict},
			PolicyContext: policyContext(r.ID, level, quality())}, verdifax.WithAfter("claims"))
		if errors.As(stepErr, &apiErr) {
			// The account policy, the server's content screen or its review
			// of the model's answer stopped this step, so the job stops here:
			// the stop is sealed as a receipt, the steps already sealed are
			// kept beside it, and a person sees the stop. The job never gets
			// a verdict.
			payloads[r.ID] = payload
			if apiErr.AivpTraceText != "" {
				traceTexts[r.ID] = apiErr.AivpTraceText
			}
			stop(ctx, agent, jobID, r.ID, runIDs, payloads, answerTexts, draftPath, draftHash, apiErr, model, inputTokens, outputTokens, models,
				traceTexts)
			return
		}
		must(stepErr)
		if receipt.AivpTraceText != "" {
			traceTexts[r.ID] = receipt.AivpTraceText
		}
		payloads[r.ID] = payload
		runIDs[r.ID] = receipt.RunID
		results = append(results, map[string]any{"rule_id": r.ID, "verdict": a.Verdict, "reason": a.Reason})
		if a.Verdict != "pass" {
			failed = append(failed, r.ID)
		}
		fmt.Printf("%s %s: %s\n", r.ID, a.Verdict, a.Reason)
	}
	if failed == nil {
		failed = []string{}
	}
	verdict := "pass"
	if len(failed) > 0 {
		verdict = "fail"
	}
	// The program decides the verdict from the rule answers, so anyone can
	// decide it again from the sealed records; the model only explains it.
	verdictPromptText := loadPromptFile(verdictPrompt)
	if results == nil {
		results = []map[string]any{}
	}
	verdictText := fmt.Sprintf("Verdict: %s\n\nRule results:\n%s", verdict, canonical(results))
	verdictPromptHash := sha256Hex([]byte(verdictPromptText + "\n" + verdictText))
	explained, explanation := askExplanation(ctx, recorded, verdictPromptText, verdictText, model)
	inputTokens += explained.InputTokens
	outputTokens += explained.OutputTokens
	models[explained.Model] = true
	answerTexts["verdict"] = explained.Raw
	verdictPayload := map[string]any{"form": form, "step": "verdict", "draft_sha256": draftHash, "rule_set_hash": ruleSetHash,
		"verdict": verdict, "failed_rules": failed, "results": results, "model": explained.Model,
		"model_response_id": explained.ResponseID, "prompt_hash": verdictPromptHash, "explanation": explanation}
	final, finalErr := job.Step(ctx, "verdict", &verdifax.AttestRequest{PayloadText: canonical(verdictPayload), ProgramID: programID,
		RouteID: routeVerdict, RegistryRecordHash: registryRecordHash, AIOutputText: explained.Raw,
		RequestSubstance: usageSubstance(&usageReport{InputTokens: explained.InputTokens, OutputTokens: explained.OutputTokens}),
		AttestedContext: map[string]any{"attested": true, "actor_id": actorID, "model_provider": "anthropic", "model_name": explained.Model,
			"prompt_hash": verdictPromptHash, "decision_kind": "claims.verdict", "decision_result": verdict},
		PolicyContext: policyContext("verdict", level, quality())}, verdifax.WithAfter(ruleIDs...), verdifax.WithFinal())
	if errors.As(finalErr, &apiErr) {
		payloads["verdict"] = verdictPayload
		if apiErr.AivpTraceText != "" {
			traceTexts["verdict"] = apiErr.AivpTraceText
		}
		stop(ctx, agent, jobID, "verdict", runIDs, payloads, answerTexts, draftPath, draftHash, apiErr, model, inputTokens, outputTokens, models,
			traceTexts)
		return
	}
	must(finalErr)
	if final.AivpTraceText != "" {
		traceTexts["verdict"] = final.AivpTraceText
	}
	payloads["verdict"] = verdictPayload
	runIDs["verdict"] = final.RunID
	fmt.Printf("verdict explained: %s\n", explanation)

	saveJobRecords(ctx, agent, jobID, out, runIDs, payloads)
	saveAnswers(out, answerTexts)
	saveTexts(filepath.Join(out, tracesFolder), traceTexts)
	modelList := make([]string, 0, len(models))
	for m := range models {
		modelList = append(modelList, m)
	}
	sort.Strings(modelList)
	var cost any
	if len(modelList) == 1 && modelList[0] == model {
		if c, known := measuredCost(model, inputTokens, outputTokens); known {
			cost = c
		}
	}
	writeJSON(filepath.Join(out, "job.json"), map[string]any{"job_id": jobID, "draft": draftPath, "draft_sha256": draftHash,
		"rule_set_hash": ruleSetHash, "verdict": verdict, "failed_rules": failed, "models": modelList, "final_run_id": final.RunID,
		"run_ids": runIDs, "measured_tokens": map[string]int64{"input": inputTokens, "output": outputTokens}, "measured_cost_usd": cost})
	fmt.Printf("verdict %s (%d of %d rules failed); job %s; records in %s\n", verdict, len(failed), len(rules), jobID, out)
	if c, isCost := cost.(float64); isCost {
		fmt.Printf("measured model use: %d input and %d output tokens, about $%.2f at published prices\n", inputTokens, outputTokens, c)
	}
	printVerifyCommands(jobID)
}

// A step the server stopped is sealed as one run whose status names the kind
// of stop, with a receipt of that kind (R152). Any other error is not a stop:
// the program reports it as the server gave it and seals nothing for it.
var stopKinds = map[string]string{
	"pepg_deny":       "the account policy or the content screen",
	"aivp_t4_halt":    "the server's review of the model's answer",
	"ccv_halt":        "the per-run budget",
	"macc_halt":       "the cumulative budget",
	"attestation_gap": "the policy's required attested fields",
}

var stopReceiptFields = map[string]string{"pepg_deny": "deny_receipt", "aivp_t4_halt": "aivp_t4_halt_receipt",
	"ccv_halt": "ccv_halt_receipt", "macc_halt": "macc_halt_receipt", "attestation_gap": "attestation_gap_receipt"}

// stopServer is what the stop handling reads from the server; the
// client satisfies it, and the tests answer it from stop-cases.json.
type stopServer interface {
	GetRun(ctx context.Context, runID int64) (*verdifax.Run, error)
	ExportRuns(ctx context.Context, params *verdifax.ExportRunsParams) (*verdifax.RunsExport, error)
	GetDenyReceipt(ctx context.Context, runID int64) (map[string]any, error)
	GetAivpT4HaltReceipt(ctx context.Context, runID int64) (map[string]any, error)
	GetCCVHaltReceipt(ctx context.Context, runID int64) (map[string]any, error)
	GetMACCHaltReceipt(ctx context.Context, runID int64) (map[string]any, error)
	GetAttestationGapReceipt(ctx context.Context, runID int64) (map[string]any, error)
}

func stopReceipt(ctx context.Context, agent stopServer, status string, runID int64) map[string]any {
	fetch := map[string]func(context.Context, int64) (map[string]any, error){
		"pepg_deny": agent.GetDenyReceipt, "aivp_t4_halt": agent.GetAivpT4HaltReceipt, "ccv_halt": agent.GetCCVHaltReceipt,
		"macc_halt": agent.GetMACCHaltReceipt, "attestation_gap": agent.GetAttestationGapReceipt,
	}[status]
	receipt, receiptErr := fetch(ctx, runID)
	must(receiptErr)
	return receipt
}

// stopReceiptHash is the sealed hash of a stopped run's receipt, whatever
// kind of stop it was.
func stopReceiptHash(ctx context.Context, agent stopServer, runID int64) string {
	run, runErr := agent.GetRun(ctx, runID)
	must(runErr)
	if _, isStop := stopKinds[run.Status]; !isStop {
		fail(fmt.Sprintf("run %d is %s, not a stopped run with a receipt", runID, run.Status))
	}
	receipt, _ := stopReceipt(ctx, agent, run.Status, runID)[stopReceiptFields[run.Status]].(map[string]any)
	digest, _ := receipt["hash"].(string)
	return digest
}

// stoppedRun is the run the server sealed for this job's stopped step, found
// by its job and step, never by being the newest of the key.
func stoppedRun(ctx context.Context, agent stopServer, jobID, step string, apiErr *verdifax.APIError) verdifax.ExportRow {
	export, exportErr := agent.ExportRuns(ctx, &verdifax.ExportRunsParams{JobID: jobID})
	must(exportErr)
	var found verdifax.ExportRow
	for _, row := range export.Runs {
		if _, isStop := stopKinds[row.Status]; isStop && row.StepID == step && row.RunID > found.RunID {
			found = row
		}
	}
	if found.RunID == 0 {
		fail(fmt.Sprintf("the server answered HTTP %d at %s: %s; it recorded no stop for this step, so nothing was sealed as one",
			apiErr.Status, step, apiErr.Message))
	}
	return found
}

// refuse records a stopped step: its sealed receipt, every record the job
// sealed before it, the exact payload bytes of every step including the
// stopped one, and a summary naming the step and the kind of stop. The same
// for intake and for any later step.
func stop(ctx context.Context, agent *verdifax.Client, jobID, step string, runIDs map[string]int64, payloads map[string]any,
	answerTexts map[string]string, draftPath, draftHash string, apiErr *verdifax.APIError, model string, inputTokens, outputTokens int64,
	models map[string]bool, traceTexts map[string]string) {
	row := stoppedRun(ctx, agent, jobID, step, apiErr)
	stopped := row.RunID
	folder := stoppedFolder(stopped) // named by the run, so two stops never share a folder
	writeJSON(filepath.Join(folder, "receipt.json"), stopReceipt(ctx, agent, row.Status, stopped))
	// A sealed step whose record was deleted since is kept as its tombstone.
	deleted := map[string]json.RawMessage{}
	if len(runIDs) > 0 {
		deleted = sealedOnly(deletedSteps(readAccountFeed(ctx, loadKeys()), jobID), runIDs)
	}
	writeDeletedSteps(folder, deleted)
	for sealedStep, runID := range runIDs {
		if _, gone := deleted[sealedStep]; gone {
			continue
		}
		artifacts, artifactsErr := agent.GetRunArtifacts(ctx, runID)
		must(artifactsErr)
		writeText(filepath.Join(folder, "steps", sealedStep+".json"), string(artifacts))
	}
	for name, payload := range payloads {
		writeText(filepath.Join(folder, "payloads", name+".json"), canonical(payload))
	}
	saveAnswers(folder, answerTexts)
	saveTexts(filepath.Join(folder, tracesFolder), traceTexts)
	modelList := make([]string, 0, len(models))
	for m := range models {
		modelList = append(modelList, m)
	}
	sort.Strings(modelList)
	var cost any // any: a JSON null when the model's price is not known
	if len(modelList) == 1 && modelList[0] == model {
		if c, known := measuredCost(model, inputTokens, outputTokens); known {
			cost = c
		}
	}
	writeJSON(filepath.Join(folder, "job.json"), map[string]any{"stopped_run_id": stopped, "step": step, "stop_status": row.Status,
		"job_id": jobID, "draft": draftPath, "draft_sha256": draftHash, "reason": apiErr.Message, "run_ids": runIDs,
		"models": modelList, "measured_tokens": map[string]int64{"input": inputTokens, "output": outputTokens}, "measured_cost_usd": cost})
	fmt.Printf("stopped at %s by %s: run %d, %s\n", step, stopKinds[row.Status], stopped, apiErr.Message)
	fmt.Printf("records in %s\n", folder)
	if c, isCost := cost.(float64); isCost {
		fmt.Printf("measured model use before the stop: %d input and %d output tokens, about $%.2f at published prices\n", inputTokens, outputTokens, c)
	}
}

// saveAnswers writes the model's exact answers, one file per step, so each
// sealed ai_output_hash recomputes from its file.
// saveTexts writes each step's text to <folder>/<step>.txt, byte for byte.
func saveTexts(folder string, texts map[string]string) {
	for step, text := range texts {
		writeText(filepath.Join(folder, step+".txt"), text)
	}
}

func saveAnswers(folder string, answerTexts map[string]string) {
	for step, text := range answerTexts {
		writeText(filepath.Join(folder, answersFolder, step+".txt"), text)
	}
}

// stoppedFolder is the folder of the stopped step whose receipt is run runID.
func stoppedFolder(runID int64) string {
	return filepath.Join("records", fmt.Sprintf("%s%d", stoppedFolderPrefix, runID))
}

func saveJobRecords(ctx context.Context, agent *verdifax.Client, jobID, out string, runIDs map[string]int64, payloads map[string]any) {
	// A deleted step's record is gone from the export; its tombstone in the
	// account feed stands in its place (R250).
	deleted := sealedOnly(deletedSteps(readAccountFeed(ctx, loadKeys()), jobID), runIDs)
	export, exportErr := agent.ExportRuns(ctx, &verdifax.ExportRunsParams{JobID: jobID})
	must(exportErr)
	rows := make([]verdifax.ExportRow, 0, len(export.Runs))
	for _, row := range export.Runs {
		if _, gone := deleted[row.StepID]; !gone {
			rows = append(rows, row)
		}
	}
	if len(rows)+len(deleted) != len(runIDs) {
		fail(fmt.Sprintf("the export lists %d records and the feed %d deleted steps for %s, the program sealed %d", len(rows), len(deleted), jobID, len(runIDs)))
	}
	// The sealed step records alone sit in steps/, the folder verdifax-verify
	// --job reads: it takes every .json file there as a record, so the
	// summary, the sign-offs and the payloads live beside that folder.
	must(os.MkdirAll(filepath.Join(out, "payloads"), 0o755))
	writeDeletedSteps(out, deleted)
	for _, row := range rows {
		if _, isStop := stopKinds[row.Status]; isStop {
			writeJSON(filepath.Join(out, "steps", row.StepID+".json"), stopReceipt(ctx, agent, row.Status, row.RunID))
			continue
		}
		artifacts, artifactsErr := agent.GetRunArtifacts(ctx, row.RunID)
		must(artifactsErr)
		writeText(filepath.Join(out, "steps", row.StepID+".json"), string(artifacts))
	}
	// The exact bytes each step sent, so a reader recomputes the sealed
	// payload hash from the file.
	for step, payload := range payloads {
		writeText(filepath.Join(out, "payloads", step+".json"), canonical(payload))
	}
}

// deletedSteps is the steps of job jobID whose records were deleted, each
// with its tombstone: the entry of every run.deleted event naming the job,
// by its step, as the exact bytes the server sent. A run.sealed event whose
// record was deleted carries the tombstone in data.deleted with a null
// entry; it is not the deletion and is passed over, as is a deletion of
// another job or of no job.
func deletedSteps(events []verdifax.SignalEvent, jobID string) map[string]json.RawMessage {
	out := map[string]json.RawMessage{}
	for _, e := range events {
		if e.Data.Kind != "run.deleted" {
			continue
		}
		var place struct {
			JobID  string `json:"job_id"`
			StepID string `json:"step_id"`
		}
		if json.Unmarshal(e.Data.Entry, &place) != nil || place.JobID != jobID || place.StepID == "" {
			continue
		}
		out[place.StepID] = e.Data.Entry
	}
	return out
}

// sealedOnly keeps the deleted steps this program sealed.
func sealedOnly(deleted map[string]json.RawMessage, runIDs map[string]int64) map[string]json.RawMessage {
	out := map[string]json.RawMessage{}
	for step, entry := range deleted {
		if _, sealed := runIDs[step]; sealed {
			out[step] = entry
		}
	}
	return out
}

// writeDeletedSteps writes each deleted step's tombstone as
// <folder>/steps/<step>.deleted.json, the server's exact bytes, beside the
// step records verdifax-verify --job reads.
func writeDeletedSteps(folder string, deleted map[string]json.RawMessage) {
	for step, entry := range deleted {
		writeText(filepath.Join(folder, "steps", step+deletedStepSuffix), string(entry)+"\n")
	}
}

// readAccountFeed is the account's whole feed, read with the feed key, page
// by page.
func readAccountFeed(ctx context.Context, keys keyFile) []verdifax.SignalEvent {
	feed, newErr := verdifax.NewClient(verdifax.Config{BaseURL: baseURL(), APIKey: keys.Feed.Secret})
	must(newErr)
	events := []verdifax.SignalEvent{}
	var after int64
	for {
		page, pageErr := feed.GetSignals(ctx, &verdifax.GetSignalsOptions{After: after, Limit: 500})
		must(pageErr)
		events = append(events, page.Events...)
		if len(page.Events) < 500 {
			break
		}
		after = page.NextAfter
	}
	return events
}

// historyCheck is the re-check of a saved sign-off history, with its head
// age limit, and the reason for the limit on the line after (R246).
func historyCheck(path string) (string, string) {
	return fmt.Sprintf("verdifax-verify --signoff-history %s --signoff-head-max-age %s", path, savedHeadMaxAge), "  " + historyAgeReason
}

// feedCheck is the re-check of the saved feed, with its head age limit, and
// the reason for the limit on the line after (R246).
func feedCheck() (string, string) {
	return "verdifax-verify --signals signals.json --feed-head-max-age " + savedHeadMaxAge, "  " + feedAgeReason
}

// ── sign-off ───────────────────────────────────────────────────────────────

// approverPrivateKey is the first approver's private key, or a named
// approver's.
func approverPrivateKey(name string) verdifax.SigningKey {
	path, how := "approver-key.json", "setup"
	if name != "" {
		path, _ = approverFiles(name)
		how = "approver-add " + name
	}
	if _, statErr := os.Stat(path); statErr != nil {
		fail(fmt.Sprintf("no %s here; run `claims-review %s` first", path, how))
	}
	var stored struct {
		PrivateKey string `json:"private_key"`
	}
	must(json.Unmarshal(readFile(path), &stored))
	key, decodeErr := hex.DecodeString(stored.PrivateKey)
	must(decodeErr)
	return verdifax.SigningKey(key)
}

func cmdSignOff(ctx context.Context, jobID, decision, note, receiptRunID, approver string) {
	if approver != "" && approverNameError(approver) != "" {
		fail(approverNameError(approver))
	}
	keys := loadKeys()
	agent := clientFor(keys.Agent.Secret, keys.Agent.SigningKey)
	owner := clientFor(keys.Owner.Secret, keys.Owner.SigningKey)
	key := approverPrivateKey(approver)
	approverFile := "approver.txt"
	if approver != "" {
		_, approverFile = approverFiles(approver)
	}
	var target verdifax.SignOffTarget
	var folder, recordFile string
	var runID int64
	if receiptRunID != "" {
		parsed, parseErr := strconv.ParseInt(receiptRunID, 10, 64)
		must(parseErr)
		runID = parsed
		target = verdifax.SignOffTarget{Kind: verdifax.SignOffTargetReceipt, RunID: runID, Digest: stopReceiptHash(ctx, agent, runID)}
		folder = stoppedFolder(runID)
		recordFile = "receipt.json"
	} else {
		folder = filepath.Join("records", jobID)
		var summary struct {
			FinalRunID int64 `json:"final_run_id"`
		}
		must(json.Unmarshal(readFile(filepath.Join(folder, "job.json")), &summary))
		runID = summary.FinalRunID
		var artifacts struct {
			BundleAnchor struct {
				AnchoredDigest string `json:"anchored_digest"`
			} `json:"bundle_anchor"`
		}
		must(json.Unmarshal(readFile(filepath.Join(folder, "steps", "verdict.json")), &artifacts))
		target = verdifax.SignOffTarget{Kind: verdifax.SignOffTargetJob, RunID: runID, Digest: artifacts.BundleAnchor.AnchoredDigest, JobID: jobID}
		recordFile = filepath.Join("steps", "verdict.json")
	}
	record, signErr := agent.SignOff(ctx, target, decision, key, &verdifax.SignOffOptions{Note: note})
	must(signErr)
	signOffs, listErr := agent.GetSignOffs(ctx, runID, nil)
	must(listErr)
	writeJSON(filepath.Join(folder, "signoffs.json"), signOffs)
	history, historyErr := owner.GetSignOffHistory(ctx)
	must(historyErr)
	writeText(filepath.Join(folder, "signoff-history.json"), string(history))
	fmt.Printf("sign-off %s recorded as sequence %d, hash %s\n", decision, record.Sequence, record.SignOffHash)
	// R272: a sign-off is checked with the account's whole sign-off history, so
	// a later decision cannot be left out; a rule that requires a sign-off, as
	// policy v4's verdict-step does, fails the check without it.
	fmt.Printf("verify: verdifax-verify --signoff %s --signoff-history %s --signoff-head-max-age %s --approver-key %s %s\n", filepath.Join(folder, "signoffs.json"),
		filepath.Join(folder, "signoff-history.json"), savedHeadMaxAge, approverFile, filepath.Join(folder, recordFile))
	historyLine, historyReason := historyCheck(filepath.Join(folder, "signoff-history.json"))
	fmt.Println("verify: " + historyLine)
	fmt.Println(historyReason)
}

// ── feed ───────────────────────────────────────────────────────────────────

// ── publish ────────────────────────────────────────────────────────────────
// The account's owner key makes a record public on its random link, or
// withdraws it. Each decision is sealed by the server as a run publication
// and placed in the account's feed, which feed then reads.

// publicationTargets is the folder and the records a publish or unpublish
// acts on: every step of a job, or a stopped step's receipt with every step
// the job sealed before it, so a reader sees what the job did before it
// stopped.
func publicationTargets(jobID, receipt string) (string, map[string]int64) {
	if receipt != "" {
		runID, parseErr := strconv.ParseInt(receipt, 10, 64)
		must(parseErr)
		folder := stoppedFolder(runID)
		var summary struct {
			RunIDs map[string]int64 `json:"run_ids"`
		}
		must(json.Unmarshal(readFile(filepath.Join(folder, "job.json")), &summary))
		targets := map[string]int64{"receipt": runID}
		for step, id := range summary.RunIDs {
			targets[step] = id
		}
		return folder, targets
	}
	folder := filepath.Join("records", jobID)
	var summary struct {
		RunIDs map[string]int64 `json:"run_ids"`
	}
	must(json.Unmarshal(readFile(filepath.Join(folder, "job.json")), &summary))
	return folder, summary.RunIDs
}

// publishable is the records a publish or unpublish acts on, sorted: every
// target but a deleted step, which has no record left; its tombstone stands in
// its place in the feed (R250).
func publishable(folder string, targets map[string]int64) []string {
	names := make([]string, 0, len(targets))
	for name := range targets {
		if _, statErr := os.Stat(filepath.Join(folder, "steps", name+deletedStepSuffix)); statErr == nil {
			continue
		}
		names = append(names, name)
	}
	sort.Strings(names)
	return names
}

func cmdPublication(ctx context.Context, jobID, receipt string, publish bool) {
	keys := loadKeys()
	owner := clientFor(keys.Owner.Secret, keys.Owner.SigningKey)
	folder, targets := publicationTargets(jobID, receipt)
	names := publishable(folder, targets)
	// A repeated publish changes nothing and the server answers with no new
	// publication, so the sealed publication saved earlier is kept.
	earlier := map[string]struct {
		RunID       int64           `json:"run_id"`
		Public      bool            `json:"public"`
		Publication json.RawMessage `json:"publication"`
	}{}
	if text, readErr := os.ReadFile(filepath.Join(folder, "published.json")); readErr == nil {
		must(json.Unmarshal(text, &earlier))
	}
	published := map[string]any{}
	var shown *verdifax.PublicationResult
	for _, name := range names {
		runID := targets[name]
		var answer *verdifax.PublicationResult
		var callErr error
		if publish {
			answer, callErr = owner.PublishRun(ctx, runID)
		} else {
			answer, callErr = owner.UnpublishRun(ctx, runID)
		}
		must(callErr)
		// A completed run is served at its bundle address, a stopped run at
		// its receipt address (R156); each answer names only the one it has.
		var bundleLink, receiptLink any // any: a JSON null when the answer names no such address
		if answer.PublicBundleURL != "" {
			bundleLink = answer.PublicBundleURL
		}
		if answer.PublicReceiptURL != "" {
			receiptLink = answer.PublicReceiptURL
		}
		var sealed any = answer.Publication // any: the sealed publication, or one saved earlier, as JSON
		if answer.Publication == nil {
			sealed = nil
			if e, saved := earlier[name]; saved && e.RunID == runID && e.Public == answer.Public && len(e.Publication) > 0 && string(e.Publication) != "null" {
				sealed = e.Publication
			}
		}
		published[name] = map[string]any{"run_id": runID, "public": answer.Public, "public_bundle_url": bundleLink,
			"public_receipt_url": receiptLink, "verify_url": answer.VerifyURL, "publication": sealed}
		if name == "verdict" || name == "receipt" || shown == nil {
			shown = answer
		}
	}
	writeJSON(filepath.Join(folder, "published.json"), published)
	verb := "withdrawn"
	if publish {
		verb = "published"
	}
	fmt.Printf("%s %d record(s) in %s\n", verb, len(published), folder)
	if publish && shown != nil {
		fmt.Printf("public link: %s\n", shown.VerifyURL)
	}
	feedLine, feedReason := feedCheck()
	fmt.Println("each decision is an event in the account's feed; after feed, verify: " + feedLine)
	fmt.Println(feedReason)
}

func cmdFeed(ctx context.Context) {
	events := readAccountFeed(ctx, loadKeys())
	writeJSON("signals.json", events)
	kinds := map[string]int{}
	for _, e := range events {
		kinds[e.Data.Kind]++
	}
	names := make([]string, 0, len(kinds))
	for k := range kinds {
		names = append(names, k)
	}
	sort.Strings(names)
	parts := make([]string, 0, len(names))
	for _, k := range names {
		parts = append(parts, fmt.Sprintf("%s %d", k, kinds[k]))
	}
	fmt.Printf("%d events in the feed: %s\n", len(events), strings.Join(parts, ", "))
	feedLine, feedReason := feedCheck()
	fmt.Println("verify: " + feedLine)
	fmt.Println(feedReason)
}

// ── verify commands ────────────────────────────────────────────────────────

func printVerifyCommands(jobID string) {
	folder := filepath.Join("records", jobID)
	fmt.Printf("verify with verdifax-verify %s or later:\n", minVerifier)
	traces := ""
	if info, statErr := os.Stat(filepath.Join(folder, tracesFolder)); statErr == nil && info.IsDir() {
		traces = " --aivp-traces " + filepath.Join(folder, tracesFolder)
	}
	fmt.Printf("  verdifax-verify --job %s%s\n", filepath.Join(folder, "steps"), traces)
	if _, statErr := os.Stat(filepath.Join(folder, "signoffs.json")); statErr == nil {
		fmt.Printf("  verdifax-verify --account-policy %s --signoff %s --signoff-history %s --signoff-head-max-age %s --approver-key approver.txt %s\n",
			filepath.Join(repo, "policy", policyFile), filepath.Join(folder, "signoffs.json"), filepath.Join(folder, "signoff-history.json"), savedHeadMaxAge,
			filepath.Join(folder, "steps", "verdict.json"))
	} else {
		fmt.Printf("  verdifax-verify --account-policy %s %s\n", filepath.Join(repo, "policy", policyFile), filepath.Join(folder, "steps", "verdict.json"))
		fmt.Println("  " + signoffPendingReason)
	}
	if _, statErr := os.Stat(filepath.Join(folder, "signoffs.json")); statErr == nil {
		fmt.Printf("  verdifax-verify --signoff %s --signoff-history %s --signoff-head-max-age %s --approver-key approver.txt %s\n", filepath.Join(folder, "signoffs.json"),
			filepath.Join(folder, "signoff-history.json"), savedHeadMaxAge, filepath.Join(folder, "steps", "verdict.json"))
	}
	if _, statErr := os.Stat(filepath.Join(folder, "signoff-history.json")); statErr == nil {
		history, historyReason := historyCheck(filepath.Join(folder, "signoff-history.json"))
		fmt.Println("  " + history)
		fmt.Println(historyReason)
	}
	if _, statErr := os.Stat("signals.json"); statErr == nil {
		feedLine, feedReason := feedCheck()
		fmt.Println("  " + feedLine)
		fmt.Println(feedReason)
	}
}

// ── the account: approvers, settings, keys, deletion, webhooks ─────────────
// Each change is made with the owner key (the agent key for its own request
// key) and sealed by the server in the account's feed, which `feed` saves:
// the feed is the record of every change, so no command keeps its own.

func ownerClient() *verdifax.Client {
	keys := loadKeys()
	return clientFor(keys.Owner.Secret, keys.Owner.SigningKey)
}

// approverID is the id the server gave this approver when it was added, kept
// with its key.
func approverID(name string) int64 {
	privatePath, _ := approverFiles(name)
	if _, statErr := os.Stat(privatePath); statErr != nil {
		fail(fmt.Sprintf("no %s here; run `claims-review approver-add %s` first", privatePath, name))
	}
	var stored struct {
		ID int64 `json:"id"`
	}
	must(json.Unmarshal(readFile(privatePath), &stored))
	return stored.ID
}

func cmdApproverAdd(ctx context.Context, name string) {
	if approverNameError(name) != "" {
		fail(approverNameError(name))
	}
	privatePath, publicPath := approverFiles(name)
	if _, statErr := os.Stat(privatePath); statErr == nil {
		fail(privatePath + " exists here; this approver is already added")
	}
	owner := ownerClient()
	public, private, genErr := verdifax.GenerateRequestKey()
	must(genErr)
	added, addErr := owner.RegisterApprover(ctx, &verdifax.RegisterApproverRequest{Name: "Claims review approver " + name, ApproverKey: private})
	must(addErr)
	writePrivate(privatePath, canonical(map[string]any{"private_key": hex.EncodeToString(private), "id": added.ID})+"\n")
	writeText(publicPath, base64.StdEncoding.EncodeToString(public))
	list, listErr := owner.ListApprovers(ctx)
	must(listErr)
	writeJSON("approvers.json", list)
	fmt.Printf("approver %s added; its public key is in %s\n", name, publicPath)
}

func cmdApproverRevoke(ctx context.Context, name string) {
	if approverNameError(name) != "" {
		fail(approverNameError(name))
	}
	owner := ownerClient()
	must(owner.RevokeApprover(ctx, approverID(name)))
	list, listErr := owner.ListApprovers(ctx)
	must(listErr)
	writeJSON("approvers.json", list)
	fmt.Printf("approver %s revoked; the sign-offs it gave stay as they were\n", name)
}

func cmdSettings(ctx context.Context, setting, value string) {
	if problem := settingError(setting, value); problem != "" {
		fail(problem)
	}
	_, setErr := ownerClient().SetAccountSetting(ctx, setting, value, nil)
	must(setErr)
	fmt.Printf("%s is %s\n", setting, value)
}

// readKeysJSON is demo-keys.json read as JSON, every field kept.
func readKeysJSON() map[string]any {
	loadKeys()
	var keys map[string]any
	must(json.Unmarshal(readFile("demo-keys.json"), &keys))
	return keys
}

func writeKeysJSON(keys map[string]any) {
	encoded, encodeErr := json.MarshalIndent(keys, "", "  ")
	must(encodeErr)
	writePrivate("demo-keys.json", string(encoded)+"\n")
}

func cmdRequestKeyRegister(ctx context.Context, keyFilePath string) {
	keys := loadKeys()
	if keys.Holder != nil {
		fail("demo-keys.json already holds the request-key holder")
	}
	var made struct {
		ID     int64  `json:"id"`
		Secret string `json:"secret"`
	}
	must(json.Unmarshal(readFile(keyFilePath), &made))
	_, signer, genErr := verdifax.GenerateRequestKey()
	must(genErr)
	holder, holderErr := verdifax.NewClient(verdifax.Config{BaseURL: baseURL(), APIKey: made.Secret, SigningKey: signer})
	must(holderErr)
	_, registerErr := holder.RegisterRequestKey(ctx)
	must(registerErr)
	all := readKeysJSON()
	all["holder"] = map[string]any{"id": made.ID, "secret": made.Secret, "signing_key": hex.EncodeToString(signer)}
	writeKeysJSON(all)
	fmt.Printf("a request key was registered on key %d, the key Verdifax made for the request-key chapter\n", made.ID)
}

func cmdRequestKeyRotate(ctx context.Context, holder bool) {
	role, _, secret, signing := keyFor(loadKeys(), holder)
	_, next, genErr := verdifax.GenerateRequestKey()
	must(genErr)
	_, rotateErr := clientFor(secret, signing).RotateRequestKey(ctx, next)
	must(rotateErr)
	nextHex := hex.EncodeToString(next)
	writeKeysJSON(withSigningKey(readKeysJSON(), role, &nextHex))
	fmt.Printf("the %s key's request key was replaced, approved by the key it replaced\n", role)
}

func cmdRequestKeyClear(ctx context.Context, holder bool) {
	role, _, secret, signing := keyFor(loadKeys(), holder)
	must(clientFor(secret, signing).ClearRequestKey(ctx))
	writeKeysJSON(withSigningKey(readKeysJSON(), role, nil))
	fmt.Printf("the %s key's request key was cleared, approved by that key\n", role)
}

func cmdKeyRevoke(ctx context.Context, holder bool) {
	keys := loadKeys()
	role, id, _, _ := keyFor(keys, holder)
	owner := clientFor(keys.Owner.Secret, keys.Owner.SigningKey)
	must(owner.RevokeAccountKey(ctx, id))
	list, listErr := owner.ListAccountKeys(ctx)
	must(listErr)
	writeJSON("account-keys.json", list)
	fmt.Printf("%s key %d revoked; the records it sealed stay as they were\n", role, id)
}

// replaceWithTombstone replaces a deleted step's record by its tombstone,
// which verdifax-verify --job reads in its place (R250).
func replaceWithTombstone(folder, step string, tombstone json.RawMessage) {
	writeDeletedSteps(folder, map[string]json.RawMessage{step: tombstone})
	record := filepath.Join(folder, "steps", step+".json")
	if _, statErr := os.Stat(record); statErr == nil {
		must(os.Remove(record))
	}
}

func cmdDeleteStep(ctx context.Context, jobID, step string) {
	folder := filepath.Join("records", jobID)
	recordPath := filepath.Join(folder, "steps", step+".json")
	if _, statErr := os.Stat(recordPath); statErr != nil {
		fail(recordPath + " does not exist; delete-step needs the step's sealed record")
	}
	var record struct {
		RunID        int64  `json:"run_id"`
		ManifestHash string `json:"manifest_hash"`
	}
	must(json.Unmarshal(readFile(recordPath), &record))
	_, deleteErr := ownerClient().DeleteRun(ctx, record.RunID, record.ManifestHash, nil)
	must(deleteErr)
	tombstone, found := deletedSteps(readAccountFeed(ctx, loadKeys()), jobID)[step]
	if !found {
		fail(fmt.Sprintf("the feed holds no run.deleted event for %s of %s", step, jobID))
	}
	replaceWithTombstone(folder, step, tombstone)
	fmt.Printf("step %s of %s deleted; its tombstone is %s\n", step, jobID, filepath.Join(folder, "steps", step+deletedStepSuffix))
	fmt.Printf("verify: verdifax-verify --job %s\n", filepath.Join(folder, "steps"))
}

func cmdWebhookSubscribe(ctx context.Context, url string) {
	hook, createErr := ownerClient().CreateWebhook(ctx, url, nil, nil)
	must(createErr)
	encoded, encodeErr := json.MarshalIndent(hook, "", "  ")
	must(encodeErr)
	writePrivate("webhook.json", string(encoded)+"\n")
	fmt.Printf("webhook %d subscribed to every kind of event at %s; its secret is in webhook.json\n", hook.ID, url)
}

// cmdWebhookRemove removes the subscription webhook-subscribe made, as the
// owner, sealed in the feed as webhook.deleted. webhook.json is kept: its
// secret can be published once the subscription is gone, since it then signs
// nothing.
func cmdWebhookRemove(ctx context.Context) {
	if _, statErr := os.Stat("webhook.json"); statErr != nil {
		fail("no webhook.json here; run `claims-review webhook-subscribe URL` first")
	}
	var hook struct {
		ID int64 `json:"id"`
	}
	must(json.Unmarshal(readFile("webhook.json"), &hook))
	removed, removeErr := ownerClient().DeleteWebhook(ctx, hook.ID, nil)
	must(removeErr)
	writeJSON("webhook-removed.json", removed)
	fmt.Printf("webhook %d removed; its removal is sealed in the account's feed\n", hook.ID)
}

// cmdWebhookList saves the lists alone, the same shape in every language:
// webhooks.json the account's webhooks, webhook-deliveries.json each one's
// deliveries by id.
func cmdWebhookList(ctx context.Context) {
	owner := ownerClient()
	hooks, listErr := owner.ListWebhooks(ctx)
	must(listErr)
	if hooks == nil {
		hooks = []verdifax.Webhook{}
	}
	writeJSON("webhooks.json", hooks)
	deliveries := map[string]any{}
	for _, h := range hooks {
		got, deliveriesErr := owner.ListWebhookDeliveries(ctx, h.ID, 500)
		must(deliveriesErr)
		if got == nil {
			got = []verdifax.WebhookDelivery{}
		}
		deliveries[strconv.FormatInt(h.ID, 10)] = got
	}
	writeJSON("webhook-deliveries.json", deliveries)
	fmt.Printf("%d webhook(s); their deliveries are in webhook-deliveries.json\n", len(hooks))
}

// ── main ───────────────────────────────────────────────────────────────────

func takeOption(args []string, name string) (string, []string) {
	for i, a := range args {
		if a == name {
			if i+1 >= len(args) {
				fail(name + " needs a value")
			}
			return args[i+1], append(append([]string{}, args[:i]...), args[i+2:]...)
		}
	}
	return "", args
}

// takeFlag reports whether a flag with no value was given, and the args
// without it.
func takeFlag(args []string, name string) (bool, []string) {
	for i, a := range args {
		if a == name {
			return true, append(append([]string{}, args[:i]...), args[i+1:]...)
		}
	}
	return false, args
}

func main() {
	ctx := context.Background()
	usage := "commands: setup, review DRAFT.md, sign-off JOB_ID approve|reject, sign-off --receipt RUN_ID approve|reject, publish JOB_ID, publish --receipt RUN_ID, unpublish JOB_ID, unpublish --receipt RUN_ID, approver-add NAME, approver-revoke NAME, settings SETTING VALUE, request-key-register KEY_FILE, request-key-rotate [--holder], request-key-clear [--holder], key-revoke [--holder], delete-step JOB_ID STEP, webhook-subscribe URL, webhook-list, webhook-remove, feed, verify-commands JOB_ID"
	if len(os.Args) < 2 {
		fail(usage)
	}
	command, args := os.Args[1], os.Args[2:]
	switch command {
	case "setup":
		cmdSetup(ctx)
	case "review":
		model, rest := takeOption(args, "--model")
		answers, rest := takeOption(rest, "--answers")
		withoutActor, rest := takeFlag(rest, "--without-actor")
		if model == "" {
			model = defaultModel
		}
		if len(rest) != 1 {
			fail("review takes one draft file")
		}
		cmdReview(ctx, rest[0], model, answers, withoutActor)
	case "sign-off":
		note, rest := takeOption(args, "--note")
		receipt, rest := takeOption(rest, "--receipt")
		approver, rest := takeOption(rest, "--approver")
		if len(rest) == 0 {
			fail("sign-off needs approve or reject")
		}
		decision := rest[len(rest)-1]
		jobID := ""
		if len(rest) > 1 {
			jobID = rest[0]
		}
		if decision != "approve" && decision != "reject" {
			fail("sign-off needs approve or reject")
		}
		if receipt == "" && jobID == "" {
			fail("sign-off needs a JOB_ID, or --receipt RUN_ID")
		}
		cmdSignOff(ctx, jobID, decision, note, receipt, approver)
	case "publish", "unpublish":
		receipt, rest := takeOption(args, "--receipt")
		jobID := ""
		if len(rest) > 0 {
			jobID = rest[0]
		}
		if receipt == "" && jobID == "" {
			fail(command + " needs a JOB_ID, or --receipt RUN_ID")
		}
		cmdPublication(ctx, jobID, receipt, command == "publish")
	case "approver-add", "approver-revoke":
		if len(args) != 1 {
			fail(command + " takes a NAME")
		}
		if command == "approver-add" {
			cmdApproverAdd(ctx, args[0])
		} else {
			cmdApproverRevoke(ctx, args[0])
		}
	case "settings":
		if len(args) != 2 {
			fail("settings takes a SETTING and a VALUE")
		}
		cmdSettings(ctx, args[0], args[1])
	case "request-key-register":
		if len(args) != 1 {
			fail("request-key-register takes the KEY_FILE of the key Verdifax made")
		}
		cmdRequestKeyRegister(ctx, args[0])
	case "request-key-rotate", "request-key-clear", "key-revoke":
		holder, rest := takeFlag(args, "--holder")
		if len(rest) != 0 {
			fail(command + " takes only --holder")
		}
		switch command {
		case "request-key-rotate":
			cmdRequestKeyRotate(ctx, holder)
		case "request-key-clear":
			cmdRequestKeyClear(ctx, holder)
		default:
			cmdKeyRevoke(ctx, holder)
		}
	case "delete-step":
		if len(args) != 2 {
			fail("delete-step takes a JOB_ID and a STEP")
		}
		cmdDeleteStep(ctx, args[0], args[1])
	case "webhook-subscribe":
		if len(args) != 1 {
			fail("webhook-subscribe takes a URL")
		}
		cmdWebhookSubscribe(ctx, args[0])
	case "webhook-list":
		cmdWebhookList(ctx)
	case "webhook-remove":
		cmdWebhookRemove(ctx)
	case "feed":
		cmdFeed(ctx)
	case "verify-commands":
		if len(args) != 1 {
			fail("verify-commands takes a JOB_ID")
		}
		printVerifyCommands(args[0])
	default:
		fail(usage)
	}
}
