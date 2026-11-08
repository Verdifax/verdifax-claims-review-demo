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
//
// Commands (the same in every language of this repository):
//
//	claims-review setup                              owner key in VERDIFAX_OWNER_KEY
//	claims-review review DRAFT.md [--model ID]        model key in ANTHROPIC_API_KEY
//	claims-review sign-off JOB_ID approve|reject [--note TEXT]
//	claims-review sign-off --receipt RUN_ID reject
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
	ruleSet        = "v1"
	actorID        = "claims-review-agent"
	routeIntake    = "claims-review/v1/intake"
	routeRule      = "claims-review/v1/rule-check"
	routeVerdict   = "claims-review/v1/verdict"
	objectPrefix   = "claims-review/v1/"
	minVerifier    = "v0.23.0"
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
		if _, statErr := os.Stat(filepath.Join(c, "rules", "v1.json")); statErr == nil {
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

func fail(message string) {
	fmt.Fprintln(os.Stderr, "claims-review:", message)
	os.Exit(1)
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
	must(json.Unmarshal(readFile(filepath.Join(repo, "rules", "v1.json")), &doc))
	var typed struct {
		Rules []rule `json:"rules"`
	}
	must(json.Unmarshal(readFile(filepath.Join(repo, "rules", "v1.json")), &typed))
	return typed.Rules, sha256Hex([]byte(canonical(doc)))
}

func loadPrompt() string { return string(readFile(filepath.Join(repo, "rules", "prompt-v1.txt"))) }

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
		fail("demo-keys.json exists here; setup runs once per account and folder")
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
	uploaded, uploadErr := owner.UploadAccountPolicy(ctx, string(readFile(filepath.Join(repo, "policy", "claims-review-v1.json"))))
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
		refuse(ctx, agent, jobID, "intake", map[string]int64{}, map[string]any{"intake": intakePayload}, draftPath, draftHash, apiErr.Message)
		return
	}
	must(intakeErr)
	payloads["intake"] = intakePayload
	runIDs := map[string]int64{"intake": intake.RunID}

	var results []map[string]any
	var failed []string
	var inputTokens, outputTokens int64
	models := map[string]bool{}
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
		receipt, stepErr := job.Step(ctx, r.ID, &verdifax.AttestRequest{PayloadText: canonical(payload), ProgramID: programID,
			RouteID: routeRule, RegistryRecordHash: registryRecordHash, AIOutputText: aiText,
			AttestedContext: map[string]any{"attested": true, "actor_id": actorID, "model_provider": "anthropic", "model_name": a.Model,
				"prompt_hash": promptHash, "decision_kind": "claims.rule_check", "decision_result": a.Verdict},
			PolicyContext: policyContext(r.ID, level)}, verdifax.WithAfter("intake"))
		if errors.As(stepErr, &apiErr) {
			// The account policy or the server's content screen refused this
			// step, so the job stops here: the refusal is sealed as a receipt,
			// the steps already sealed are kept beside it, and a person sees
			// the stop. The job never gets a verdict.
			payloads[r.ID] = payload
			refuse(ctx, agent, jobID, r.ID, runIDs, payloads, draftPath, draftHash, apiErr.Message)
			return
		}
		must(stepErr)
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
	verdictPayload := map[string]any{"form": form, "step": "verdict", "draft_sha256": draftHash, "rule_set_hash": ruleSetHash,
		"verdict": verdict, "failed_rules": failed, "results": results}
	final, finalErr := job.Step(ctx, "verdict", &verdifax.AttestRequest{PayloadText: canonical(verdictPayload), ProgramID: programID,
		RouteID: routeVerdict, RegistryRecordHash: registryRecordHash,
		AttestedContext: map[string]any{"attested": true, "actor_id": actorID, "decision_kind": "claims.verdict", "decision_result": verdict},
		PolicyContext:   policyContext("verdict", level)}, verdifax.WithAfter(ruleIDs...), verdifax.WithFinal())
	must(finalErr)
	payloads["verdict"] = verdictPayload
	runIDs["verdict"] = final.RunID

	saveJobRecords(ctx, agent, jobID, out, runIDs, payloads)
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

func newestRefusal(ctx context.Context, agent *verdifax.Client) int64 {
	export, exportErr := agent.ExportRuns(ctx, &verdifax.ExportRunsParams{})
	must(exportErr)
	var newest int64
	for _, row := range export.Runs {
		if row.Status == "pepg_deny" && row.RunID > newest {
			newest = row.RunID
		}
	}
	if newest == 0 {
		fail("the step was refused but no refused run is listed for this key")
	}
	return newest
}

// refuse records a refused step: its sealed receipt, every record the job
// sealed before it, the exact payload bytes of every step including the
// refused one, and a summary naming the step. The same for intake and for
// any later step.
func refuse(ctx context.Context, agent *verdifax.Client, jobID, step string, runIDs map[string]int64, payloads map[string]any,
	draftPath, draftHash, reason string) {
	refused := newestRefusal(ctx, agent)
	folder := filepath.Join("records", fmt.Sprintf("refused-%d", refused)) // named by the run, so two refusals never share a folder
	receipt, receiptErr := agent.GetDenyReceipt(ctx, refused)
	must(receiptErr)
	writeJSON(filepath.Join(folder, "receipt.json"), receipt)
	for sealedStep, runID := range runIDs {
		artifacts, artifactsErr := agent.GetRunArtifacts(ctx, runID)
		must(artifactsErr)
		writeText(filepath.Join(folder, sealedStep+".json"), string(artifacts))
	}
	for name, payload := range payloads {
		writeText(filepath.Join(folder, "payloads", name+".json"), canonical(payload))
	}
	writeJSON(filepath.Join(folder, "job.json"), map[string]any{"refused_run_id": refused, "step": step, "job_id": jobID,
		"draft": draftPath, "draft_sha256": draftHash, "reason": reason, "run_ids": runIDs})
	fmt.Printf("refused at %s: run %d, %s\n", step, refused, reason)
	fmt.Printf("records in %s\n", folder)
}

func saveJobRecords(ctx context.Context, agent *verdifax.Client, jobID, out string, runIDs map[string]int64, payloads map[string]any) {
	export, exportErr := agent.ExportRuns(ctx, &verdifax.ExportRunsParams{JobID: jobID})
	must(exportErr)
	if len(export.Runs) != len(runIDs) {
		fail(fmt.Sprintf("the export lists %d records for %s, the program sealed %d", len(export.Runs), jobID, len(runIDs)))
	}
	must(os.MkdirAll(filepath.Join(out, "payloads"), 0o755))
	for _, row := range export.Runs {
		if row.Status == "pepg_deny" {
			receipt, receiptErr := agent.GetDenyReceipt(ctx, row.RunID)
			must(receiptErr)
			writeJSON(filepath.Join(out, row.StepID+".json"), receipt)
			continue
		}
		artifacts, artifactsErr := agent.GetRunArtifacts(ctx, row.RunID)
		must(artifactsErr)
		writeText(filepath.Join(out, row.StepID+".json"), string(artifacts))
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
		receipt, receiptErr := agent.GetDenyReceipt(ctx, runID)
		must(receiptErr)
		digest, _ := receipt["deny_receipt"].(map[string]any)["hash"].(string)
		target = verdifax.SignOffTarget{Kind: verdifax.SignOffTargetReceipt, RunID: runID, Digest: digest}
		folder = filepath.Join("records", fmt.Sprintf("receipt-%d", runID))
		matches, _ := filepath.Glob(filepath.Join("records", "refused-*", "job.json"))
		for _, m := range matches {
			var summary struct {
				RefusedRunID int64 `json:"refused_run_id"`
			}
			if json.Unmarshal(readFile(m), &summary) == nil && summary.RefusedRunID == runID {
				folder = filepath.Dir(m)
			}
		}
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
		must(json.Unmarshal(readFile(filepath.Join(folder, "verdict.json")), &artifacts))
		target = verdifax.SignOffTarget{Kind: verdifax.SignOffTargetJob, RunID: runID, Digest: artifacts.BundleAnchor.AnchoredDigest, JobID: jobID}
		recordFile = "verdict.json"
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
	fmt.Printf("  verdifax-verify --job %s\n", folder)
	fmt.Printf("  verdifax-verify --account-policy %s %s\n", filepath.Join(repo, "policy", "claims-review-v1.json"), filepath.Join(folder, "verdict.json"))
	if _, statErr := os.Stat(filepath.Join(folder, "signoffs.json")); statErr == nil {
		fmt.Printf("  verdifax-verify --signoff %s --approver-key approver.txt %s\n", filepath.Join(folder, "signoffs.json"), filepath.Join(folder, "verdict.json"))
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
	usage := "commands: setup, review DRAFT.md, sign-off JOB_ID approve|reject, sign-off --receipt RUN_ID reject, feed, verify-commands JOB_ID"
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
