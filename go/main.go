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
//	claims-review sign-off JOB_ID approve|reject [--note TEXT]
//	claims-review sign-off --receipt RUN_ID approve|reject [--note TEXT]
//	claims-review publish JOB_ID | publish --receipt RUN_ID
//	claims-review unpublish JOB_ID | unpublish --receipt RUN_ID
//	claims-review feed
//	claims-review verify-commands JOB_ID
//
// Environment: VERDIFAX_BASE_URL (default https://api.verdifax.com).
// Files, in the working folder: demo-keys.json and approver-key.json (0600),
// approver.txt, account-history.json, records/<job_id>/*.json with payloads/,
// signals.json.
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
	minVerifier    = "v0.28.0"
	// A stopped step's records live in records/stopped-<run_id>/, and the
	// model's exact answers in answers/ beside each job's steps/.
	stoppedFolderPrefix = "stopped-"
	answersFolder       = "answers"
	// The text the server's AI review checked for each step, as the server
	// returned it once: a reader re-runs the review's check from these files.
	tracesFolder = "traces"
	// The account policy setup activates; v1 and v2 stay for records sealed
	// under them. v3 requires a reviewed AI answer on every step after the
	// gate.
	policyFile = "claims-review-v3.json"
	// The AI's work on the steps after the gate: listing the draft's claims,
	// and explaining the verdict code decided.
	claimsPrompt  = "claims-prompt-v1.txt"
	verdictPrompt = "verdict-prompt-v1.txt"
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

func clientFor(secret, signingHex string) *verdifax.Client {
	key, decodeErr := hex.DecodeString(signingHex)
	must(decodeErr)
	client, newErr := verdifax.NewClient(verdifax.Config{BaseURL: baseURL(), APIKey: secret, SigningKey: verdifax.SigningKey(key)})
	must(newErr)
	return client
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
		out := structured{Model: recorded.Model, ResponseID: recorded.ResponseID, Raw: recordedRaw}
		if out.Model == "" {
			out.Model = "recorded-answers"
		}
		if out.ResponseID == "" {
			out.ResponseID = "recorded"
		}
		return out, true
	}
	client := anthropic.NewClient()
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
			Raw: canonical(map[string]string{"verdict": a.Verdict, "reason": a.Reason})}
	}
	client := anthropic.NewClient()
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

func policyContext(objectID, level string) *verdifax.PolicyContext {
	return &verdifax.PolicyContext{ObjectType: "text", ObjectID: objectPrefix + objectID, Classification: level}
}

func cmdReview(ctx context.Context, draftPath, model, answersPath string) {
	keys := loadKeys()
	agent := clientFor(keys.Agent.Secret, keys.Agent.SigningKey)
	rules, ruleSetHash := loadRules()
	prompt := loadPrompt()
	draft := string(readFile(draftPath))
	draftHash := sha256Hex([]byte(draft))
	level := draftClassification(draft)
	recorded := loadAnswers(answersPath)
	now := time.Now().UTC().Format("20060102T150405Z")
	jobID := fmt.Sprintf("claims-review-%s-%s", draftHash[:12], now)
	out := filepath.Join("records", jobID)
	payloads := map[string]any{}

	intakePayload := map[string]any{"form": form, "step": "intake", "draft_sha256": draftHash, "draft_bytes": len(draft),
		"rule_set": ruleSet, "rule_set_hash": ruleSetHash, "classification": level}
	job, jobErr := agent.Job(jobID)
	must(jobErr)
	intake, intakeErr := job.Step(ctx, "intake", &verdifax.AttestRequest{PayloadText: canonical(intakePayload), ProgramID: programID,
		RouteID: routeIntake, RegistryRecordHash: registryRecordHash, PolicyContext: policyContext("intake", level)}, verdifax.WithAfter())
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
		AttestedContext: map[string]any{"attested": true, "actor_id": actorID, "model_provider": "anthropic", "model_name": listed.Model,
			"prompt_hash": claimsPromptHash, "decision_kind": "claims.list", "decision_result": fmt.Sprint(len(claims))},
		PolicyContext: policyContext("claims", level)}, verdifax.WithAfter("intake"))
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
			AttestedContext: map[string]any{"attested": true, "actor_id": actorID, "model_provider": "anthropic", "model_name": a.Model,
				"prompt_hash": promptHash, "decision_kind": "claims.rule_check", "decision_result": a.Verdict},
			PolicyContext: policyContext(r.ID, level)}, verdifax.WithAfter("claims"))
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
		AttestedContext: map[string]any{"attested": true, "actor_id": actorID, "model_provider": "anthropic", "model_name": explained.Model,
			"prompt_hash": verdictPromptHash, "decision_kind": "claims.verdict", "decision_result": verdict},
		PolicyContext: policyContext("verdict", level)}, verdifax.WithAfter(ruleIDs...), verdifax.WithFinal())
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
	for sealedStep, runID := range runIDs {
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
	export, exportErr := agent.ExportRuns(ctx, &verdifax.ExportRunsParams{JobID: jobID})
	must(exportErr)
	if len(export.Runs) != len(runIDs) {
		fail(fmt.Sprintf("the export lists %d records for %s, the program sealed %d", len(export.Runs), jobID, len(runIDs)))
	}
	// The sealed step records alone sit in steps/, the folder verdifax-verify
	// --job reads: it takes every .json file there as a record, so the
	// summary, the sign-offs and the payloads live beside that folder.
	must(os.MkdirAll(filepath.Join(out, "payloads"), 0o755))
	for _, row := range export.Runs {
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

// ── sign-off ───────────────────────────────────────────────────────────────

func approverPrivateKey() verdifax.SigningKey {
	if _, statErr := os.Stat("approver-key.json"); statErr != nil {
		fail("no approver-key.json here; run `claims-review setup` first")
	}
	var stored struct {
		PrivateKey string `json:"private_key"`
	}
	must(json.Unmarshal(readFile("approver-key.json"), &stored))
	key, decodeErr := hex.DecodeString(stored.PrivateKey)
	must(decodeErr)
	return verdifax.SigningKey(key)
}

func cmdSignOff(ctx context.Context, jobID, decision, note, receiptRunID string) {
	keys := loadKeys()
	agent := clientFor(keys.Agent.Secret, keys.Agent.SigningKey)
	owner := clientFor(keys.Owner.Secret, keys.Owner.SigningKey)
	key := approverPrivateKey()
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
	fmt.Printf("verify: verdifax-verify --signoff %s --approver-key approver.txt %s\n", filepath.Join(folder, "signoffs.json"), filepath.Join(folder, recordFile))
	fmt.Printf("verify: verdifax-verify --signoff-history %s\n", filepath.Join(folder, "signoff-history.json"))
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

func cmdPublication(ctx context.Context, jobID, receipt string, publish bool) {
	keys := loadKeys()
	owner := clientFor(keys.Owner.Secret, keys.Owner.SigningKey)
	folder, targets := publicationTargets(jobID, receipt)
	names := make([]string, 0, len(targets))
	for name := range targets {
		names = append(names, name)
	}
	sort.Strings(names)
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
	fmt.Println("each decision is an event in the account's feed; after feed, verify: verdifax-verify --signals signals.json")
}

func cmdFeed(ctx context.Context) {
	keys := loadKeys()
	feed, newErr := verdifax.NewClient(verdifax.Config{BaseURL: baseURL(), APIKey: keys.Feed.Secret})
	must(newErr)
	var events []verdifax.SignalEvent
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
	if events == nil {
		events = []verdifax.SignalEvent{}
	}
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
	fmt.Println("verify: verdifax-verify --signals signals.json")
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
	fmt.Printf("  verdifax-verify --account-policy %s %s\n", filepath.Join(repo, "policy", policyFile), filepath.Join(folder, "steps", "verdict.json"))
	if _, statErr := os.Stat(filepath.Join(folder, "signoffs.json")); statErr == nil {
		fmt.Printf("  verdifax-verify --signoff %s --approver-key approver.txt %s\n", filepath.Join(folder, "signoffs.json"), filepath.Join(folder, "steps", "verdict.json"))
	}
	if _, statErr := os.Stat("signals.json"); statErr == nil {
		fmt.Println("  verdifax-verify --signals signals.json")
	}
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

func main() {
	ctx := context.Background()
	usage := "commands: setup, review DRAFT.md, sign-off JOB_ID approve|reject, sign-off --receipt RUN_ID approve|reject, publish JOB_ID, publish --receipt RUN_ID, unpublish JOB_ID, unpublish --receipt RUN_ID, feed, verify-commands JOB_ID"
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
		if model == "" {
			model = defaultModel
		}
		if len(rest) != 1 {
			fail("review takes one draft file")
		}
		cmdReview(ctx, rest[0], model, answers)
	case "sign-off":
		note, rest := takeOption(args, "--note")
		receipt, rest := takeOption(rest, "--receipt")
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
		cmdSignOff(ctx, jobID, decision, note, receipt)
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
