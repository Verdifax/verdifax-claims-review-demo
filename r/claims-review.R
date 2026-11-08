#!/usr/bin/env Rscript
# claims-review: a small agent that checks a draft against 15 general writing
# rules, with every step sealed by Verdifax.
#
# A demo, not customer proof. A sealed record shows which rules were checked
# and what the model answered; it does not show the model's judgment was
# right. The 15 rules are general writing rules, not legal advice and not a
# compliance standard. The account policy judges what each step declared; it
# cannot show the declared facts are true. A sign-off shows a registered
# approver key signed that decision; it does not show who holds the key or
# that the review was careful.
#
# Commands (the same in every language of this repository):
#
#   claims-review setup                              owner key in VERDIFAX_OWNER_KEY
#   claims-review review DRAFT.md [--model ID]        model key in ANTHROPIC_API_KEY
#   claims-review sign-off JOB_ID approve|reject [--note TEXT]
#   claims-review sign-off --receipt RUN_ID reject
#   claims-review feed
#   claims-review verify-commands JOB_ID
#
# Environment: VERDIFAX_BASE_URL (default https://api.verdifax.com).
# Files, in the working folder: demo-keys.json and approver-key.json (0600),
# approver.txt, account-history.json, records/<job_id>/*.json with payloads/,
# signals.json.

suppressPackageStartupMessages({
  library(verdifax)
  library(jsonlite)
})

DEFAULT_BASE_URL <- "https://api.verdifax.com"
DEFAULT_MODEL <- "claude-fable-5-1"
FORM <- "verdifax.claims-review.v1"
RULE_SET <- "v1"
ACTOR_ID <- "claims-review-agent"
ROUTE_INTAKE <- "claims-review/v1/intake"
ROUTE_RULE <- "claims-review/v1/rule-check"
ROUTE_VERDICT <- "claims-review/v1/verdict"
OBJECT_PREFIX <- "claims-review/v1/"
LEVELS <- c("PUBLIC", "INTERNAL", "CONFIDENTIAL", "RESTRICTED", "TOP_SECRET", "TS_SCI")
MIN_VERIFIER <- "v0.23.0"

script_dir <- function() {
  args <- commandArgs(trailingOnly = FALSE)
  file <- sub("^--file=", "", args[grepl("^--file=", args)])
  if (length(file) == 0) return(normalizePath("."))
  normalizePath(dirname(file[1]))
}
REPO <- normalizePath(file.path(script_dir(), ".."))

sha256_hex <- function(text) paste(as.character(unclass(openssl::sha256(charToRaw(enc2utf8(text))))), collapse = "")

# Canonical JSON: sorted keys, no spaces, as every language of this
# repository writes it. Scalars are written by jsonlite.
json_scalar <- function(v) as.character(jsonlite::toJSON(v, auto_unbox = TRUE, digits = NA))
canonical <- function(value) {
  if (is.null(value)) return("null")
  if (is.list(value) && !is.null(names(value))) {
    keys <- sort(names(value))
    parts <- vapply(keys, function(k) paste0(json_scalar(k), ":", canonical(value[[k]])), character(1))
    return(paste0("{", paste(parts, collapse = ","), "}"))
  }
  if (is.list(value)) {
    return(paste0("[", paste(vapply(value, canonical, character(1)), collapse = ","), "]"))
  }
  if (length(value) == 0) return("[]")
  if (length(value) > 1) return(paste0("[", paste(vapply(value, json_scalar, character(1)), collapse = ","), "]"))
  json_scalar(value)
}
as_list <- function(x) lapply(x, identity)  # a character vector that stays a JSON array

PROGRAM_ID <- sha256_hex("verdifax-claims-review-demo/program/v1")
REGISTRY_RECORD_HASH <- sha256_hex("verdifax-claims-review-demo/registry/v1")

base_url <- function() sub("/+$", "", Sys.getenv("VERDIFAX_BASE_URL", DEFAULT_BASE_URL))
is_production <- function() base_url() == DEFAULT_BASE_URL

fail <- function(message) {
  cat("claims-review:", message, "\n", file = stderr())
  quit(status = 1)
}

read_text <- function(path) paste(readLines(path, warn = FALSE, encoding = "UTF-8"), collapse = "\n")
write_text <- function(path, text) {
  dir.create(dirname(path), showWarnings = FALSE, recursive = TRUE)
  writeLines(enc2utf8(text), path, useBytes = TRUE, sep = "")
}
write_private <- function(path, text) {
  write_text(path, text)
  Sys.chmod(path, "0600")
}
write_json <- function(path, value) {
  dir.create(dirname(path), showWarnings = FALSE, recursive = TRUE)
  writeLines(as.character(jsonlite::toJSON(value, auto_unbox = TRUE, pretty = TRUE, digits = NA, null = "null")), path)
}
hex <- function(raw) paste(sprintf("%02x", as.integer(raw)), collapse = "")
unhex <- function(text) as.raw(strtoi(substring(text, seq(1, nchar(text), 2), seq(2, nchar(text), 2)), 16L))

load_rules <- function() {
  doc <- jsonlite::fromJSON(file.path(REPO, "rules", "v1.json"), simplifyVector = FALSE)
  list(rules = doc$rules, rule_set_hash = sha256_hex(canonical(doc)))
}
load_prompt <- function() {
  # The exact bytes of the prompt file, so its hash matches every language.
  readChar(file.path(REPO, "rules", "prompt-v1.txt"), file.info(file.path(REPO, "rules", "prompt-v1.txt"))$size, useBytes = TRUE)
}

draft_classification <- function(text) {
  first <- trimws(strsplit(text, "\n", fixed = TRUE)[[1]][1])
  m <- regmatches(first, regexec("^<!--\\s*classification:\\s*([A-Z_]+)\\s*-->$", first))[[1]]
  level <- if (length(m) == 2) m[2] else "INTERNAL"
  if (!level %in% LEVELS) fail(paste("the draft declares an unknown classification:", level))
  level
}

load_keys <- function() {
  if (!file.exists("demo-keys.json")) fail("no demo-keys.json here; run `claims-review setup` first")
  jsonlite::fromJSON("demo-keys.json", simplifyVector = FALSE)
}
client_for <- function(key) verdifax_client(base_url = base_url(), api_key = key$secret, signing_key = unhex(key$signing_key))

# ── setup ──────────────────────────────────────────────────────────────────

cmd_setup <- function() {
  if (file.exists("demo-keys.json")) fail("demo-keys.json exists here; setup runs once per account and folder")
  owner_secret <- Sys.getenv("VERDIFAX_OWNER_KEY", "")
  if (owner_secret == "") fail("set VERDIFAX_OWNER_KEY to the account's owner key")
  owner_signer <- verdifax_generate_request_key()
  owner <- verdifax_client(base_url = base_url(), api_key = owner_secret, signing_key = owner_signer$private_key)
  verdifax_register_account_request_key(owner)
  member <- verdifax_create_account_key(owner, "claims-review-agent")
  member_signer <- verdifax_generate_request_key()
  agent <- verdifax_client(base_url = base_url(), api_key = member$secret, signing_key = member_signer$private_key)
  verdifax_register_request_key(agent)
  policy_text <- read_text(file.path(REPO, "policy", "claims-review-v1.json"))
  uploaded <- verdifax_upload_account_policy(owner, policy_text)
  verdifax_activate_account_policy(owner, uploaded$document_hash)
  approver <- verdifax_generate_request_key()
  verdifax_register_approver(owner, approver$private_key, "Claims review approver")
  feed <- verdifax_create_feed_key(owner, "claims-review-feed")
  write_text("account-history.json", verdifax_get_account_policy_history(owner))
  write_private("approver-key.json", paste0(canonical(list(private_key = hex(approver$private_key))), "\n"))
  write_text("approver.txt", jsonlite::base64_enc(approver$public_key))
  keys <- list(
    base_url = base_url(),
    owner = list(secret = owner_secret, signing_key = hex(owner_signer$private_key)),
    agent = list(id = member$id, secret = member$secret, signing_key = hex(member_signer$private_key)),
    feed = list(id = feed$id, secret = feed$secret),
    policy_document_hash = uploaded$document_hash
  )
  write_private("demo-keys.json", paste0(as.character(jsonlite::toJSON(keys, auto_unbox = TRUE, pretty = TRUE)), "\n"))
  cat(sprintf("setup done: agent key %s, feed key %s, policy %s\n", member$id, feed$id, uploaded$document_hash))
}

# ── the model ──────────────────────────────────────────────────────────────

# Recorded answers for tests, refused against production.
load_answers <- function(path) {
  if (is.null(path)) return(NULL)
  if (is_production()) fail("--answers is for tests against a local server, never production")
  jsonlite::fromJSON(path, simplifyVector = FALSE)
}

ask <- function(recorded, prompt, user_text, rule_id, model) {
  if (!is.null(recorded)) {
    a <- recorded[[rule_id]]
    return(list(model = if (is.null(a$model)) "recorded-answers" else a$model,
                response_id = if (is.null(a$response_id)) "recorded" else a$response_id,
                verdict = a$verdict, reason = a$reason,
                raw = canonical(list(verdict = a$verdict, reason = a$reason)), input_tokens = 0L, output_tokens = 0L))
  }
  key <- Sys.getenv("ANTHROPIC_API_KEY", "")
  if (key == "") fail("set ANTHROPIC_API_KEY to call the model")
  body <- list(
    model = model, max_tokens = 4000L, system = prompt,
    messages = list(list(role = "user", content = user_text)),
    output_config = list(format = list(type = "json_schema", schema = list(
      type = "object",
      properties = list(verdict = list(type = "string"), reason = list(type = "string")),
      required = list("verdict", "reason"), additionalProperties = FALSE)))
  )
  req <- httr2::request(paste0(sub("/+$", "", Sys.getenv("ANTHROPIC_BASE_URL", "https://api.anthropic.com")), "/v1/messages"))
  req <- httr2::req_headers(req, `x-api-key` = key, `anthropic-version` = "2026-10-21", `content-type` = "application/json")
  req <- httr2::req_body_json(req, body, auto_unbox = TRUE)
  req <- httr2::req_timeout(req, 600)
  resp <- httr2::req_perform(req)
  parsed <- jsonlite::fromJSON(httr2::resp_body_string(resp), simplifyVector = FALSE)
  usage <- list(input_tokens = as.integer(parsed$usage$input_tokens), output_tokens = as.integer(parsed$usage$output_tokens))
  text_blocks <- Filter(function(b) identical(b$type, "text"), parsed$content)
  raw <- if (length(text_blocks) > 0) text_blocks[[1]]$text else ""
  answer <- tryCatch(jsonlite::fromJSON(raw, simplifyVector = FALSE), error = function(e) NULL)
  if (identical(parsed$stop_reason, "refusal") || is.null(answer) || is.null(answer$verdict) || is.null(answer$reason)) {
    return(c(list(model = parsed$model, response_id = parsed$id, verdict = "unanswered",
                  reason = "the model declined to answer", raw = ""), usage))
  }
  verdict <- if (answer$verdict %in% c("pass", "fail")) answer$verdict else "unanswered"
  c(list(model = parsed$model, response_id = parsed$id, verdict = verdict, reason = answer$reason, raw = raw), usage)
}

measured_cost <- function(model, input_tokens, output_tokens) {
  prices <- jsonlite::fromJSON(file.path(REPO, "rules", "model-prices-v1.json"), simplifyVector = FALSE)$prices
  if (is.null(prices[[model]])) return(NULL)
  input_tokens / 1e6 * prices[[model]]$input + output_tokens / 1e6 * prices[[model]]$output
}

# ── review ─────────────────────────────────────────────────────────────────

policy_context <- function(object_id, level) list(object_type = "text", object_id = paste0(OBJECT_PREFIX, object_id), classification = level)

cmd_review <- function(draft_path, model, answers_path) {
  keys <- load_keys()
  agent <- client_for(keys$agent)
  loaded <- load_rules()
  rules <- loaded$rules
  rule_set_hash <- loaded$rule_set_hash
  prompt <- load_prompt()
  draft <- read_text(draft_path)
  draft_hash <- sha256_hex(draft)
  level <- draft_classification(draft)
  recorded <- load_answers(answers_path)
  now <- format(Sys.time(), "%Y%m%dT%H%M%SZ", tz = "UTC")
  job_id <- sprintf("claims-review-%s-%s", substr(draft_hash, 1, 12), now)
  out <- file.path("records", job_id)
  payloads <- list()

  intake_payload <- list(form = FORM, step = "intake", draft_sha256 = draft_hash, draft_bytes = nchar(draft, type = "bytes"),
                         rule_set = RULE_SET, rule_set_hash = rule_set_hash, classification = level)
  job <- verdifax_job(agent, job_id)
  intake <- tryCatch(
    verdifax_job_step(job, "intake", after = list(), payload = canonical(intake_payload), program_id = PROGRAM_ID,
                      route_id = ROUTE_INTAKE, registry_record_hash = REGISTRY_RECORD_HASH,
                      policy_context = policy_context("intake", level)),
    httr2_http = function(e) e)
  if (inherits(intake, "error")) {
    # The account policy refused the draft at intake (for example a
    # confidential draft), so it never reached a model. The refusal is
    # sealed as a receipt; nothing else runs for this draft.
    refuse(agent, job_id, "intake", empty_object(), list(intake = intake_payload), draft_path, draft_hash, error_reason(intake))
    return(invisible(NULL))
  }
  payloads$intake <- intake_payload
  run_ids <- list(intake = intake$run_id)

  results <- list()
  tokens <- list(input = 0L, output = 0L)
  models <- character(0)
  for (rule in rules) {
    user_text <- sprintf("Rule %s: %s\n\nDraft:\n%s", rule$id, rule$text, draft)
    prompt_hash <- sha256_hex(paste0(prompt, "\n", user_text))
    answer <- ask(recorded, prompt, user_text, rule$id, model)
    tokens$input <- tokens$input + answer$input_tokens
    tokens$output <- tokens$output + answer$output_tokens
    models <- union(models, answer$model)
    payload <- list(form = FORM, step = rule$id, rule_id = rule$id, rule_text = rule$text, draft_sha256 = draft_hash,
                    rule_set_hash = rule_set_hash, model = answer$model, model_response_id = answer$response_id,
                    prompt_hash = prompt_hash, verdict = answer$verdict, reason = answer$reason)
    ai_text <- if (nzchar(answer$raw)) answer$raw else answer$reason
    receipt <- tryCatch(
      verdifax_job_step(job, rule$id, after = list("intake"), payload = canonical(payload), program_id = PROGRAM_ID,
                        route_id = ROUTE_RULE, registry_record_hash = REGISTRY_RECORD_HASH, ai_output_text = ai_text,
                        attested_context = list(attested = TRUE, actor_id = ACTOR_ID, model_provider = "anthropic",
                                                model_name = answer$model, prompt_hash = prompt_hash,
                                                decision_kind = "claims.rule_check", decision_result = answer$verdict),
                        policy_context = policy_context(rule$id, level)),
      httr2_http = function(e) e)
    if (inherits(receipt, "error")) {
      # The account policy or the server's content screen refused this
      # step, so the job stops here: the refusal is sealed as a receipt,
      # the steps already sealed are kept beside it, and a person sees
      # the stop. The job never gets a verdict.
      payloads[[rule$id]] <- payload
      refuse(agent, job_id, rule$id, run_ids, payloads, draft_path, draft_hash, error_reason(receipt))
      return(invisible(NULL))
    }
    payloads[[rule$id]] <- payload
    run_ids[[rule$id]] <- receipt$run_id
    results[[length(results) + 1]] <- list(rule_id = rule$id, verdict = answer$verdict, reason = answer$reason)
    cat(sprintf("%s %s: %s\n", rule$id, answer$verdict, answer$reason))
  }

  failed <- unlist(lapply(results, function(r) if (r$verdict != "pass") r$rule_id else NULL))
  if (is.null(failed)) failed <- character(0)
  verdict <- if (length(failed) == 0) "pass" else "fail"
  verdict_payload <- list(form = FORM, step = "verdict", draft_sha256 = draft_hash, rule_set_hash = rule_set_hash,
                          verdict = verdict, failed_rules = as_list(failed), results = results)
  final <- verdifax_job_step(job, "verdict", after = as_list(vapply(rules, function(r) r$id, character(1))), final = TRUE,
                             payload = canonical(verdict_payload), program_id = PROGRAM_ID, route_id = ROUTE_VERDICT,
                             registry_record_hash = REGISTRY_RECORD_HASH,
                             attested_context = list(attested = TRUE, actor_id = ACTOR_ID, decision_kind = "claims.verdict",
                                                     decision_result = verdict),
                             policy_context = policy_context("verdict", level))
  payloads$verdict <- verdict_payload
  run_ids$verdict <- final$run_id

  save_job_records(agent, job_id, out, run_ids, payloads)
  cost <- if (length(models) == 1 && models == model) measured_cost(model, tokens$input, tokens$output) else NULL
  summary <- list(job_id = job_id, draft = draft_path, draft_sha256 = draft_hash, rule_set_hash = rule_set_hash,
                  verdict = verdict, failed_rules = as_list(failed), models = as_list(sort(models)),
                  final_run_id = final$run_id, run_ids = run_ids, measured_tokens = tokens, measured_cost_usd = cost)
  write_json(file.path(out, "job.json"), summary)
  cat(sprintf("verdict %s (%d of %d rules failed); job %s; records in %s\n", verdict, length(failed), length(rules), job_id, out))
  if (!is.null(cost)) {
    cat(sprintf("measured model use: %d input and %d output tokens, about $%.2f at published prices\n", tokens$input, tokens$output, cost))
  }
  print_verify_commands(job_id)
}

newest_refusal <- function(agent) {
  rows <- verdifax_export_runs(agent)$runs
  ids <- unlist(lapply(rows, function(r) if (identical(r$status, "pepg_deny")) r$run_id else NULL))
  if (length(ids) == 0) fail("the step was refused but no refused run is listed for this key")
  max(ids)
}

# The server's own message rides in the error's body; the headline is only
# the HTTP status.
error_reason <- function(e) if (length(e$body) > 0) paste(e$body, collapse = " ") else conditionMessage(e)

# An empty JSON object ({}), as the other programs write an empty run_ids.
empty_object <- function() structure(list(), names = character(0))

# Record a refused step: its sealed receipt, every record the job sealed
# before it, the exact payload bytes of every step including the refused
# one, and a summary naming the step. The same for intake and for any
# later step.
refuse <- function(agent, job_id, step, run_ids, payloads, draft_path, draft_hash, reason) {
  refused <- newest_refusal(agent)
  folder <- file.path("records", paste0("refused-", refused))  # named by the run, so two refusals never share a folder
  write_json(file.path(folder, "receipt.json"), verdifax_get_deny_receipt(agent, refused))
  for (sealed_step in names(run_ids)) {
    write_text(file.path(folder, "steps", paste0(sealed_step, ".json")), verdifax_get_run_artifacts(agent, run_ids[[sealed_step]]))
  }
  for (name in names(payloads)) write_text(file.path(folder, "payloads", paste0(name, ".json")), canonical(payloads[[name]]))
  write_json(file.path(folder, "job.json"), list(refused_run_id = refused, step = step, job_id = job_id, draft = draft_path,
                                                 draft_sha256 = draft_hash, reason = reason, run_ids = run_ids))
  cat(sprintf("refused at %s: run %d, %s\n", step, refused, reason))
  cat(sprintf("records in %s\n", folder))
}

save_job_records <- function(agent, job_id, out, run_ids, payloads) {
  rows <- verdifax_export_runs(agent, job_id = job_id)$runs
  if (length(rows) != length(run_ids)) {
    fail(sprintf("the export lists %d records for %s, the program sealed %d", length(rows), job_id, length(run_ids)))
  }
  # The sealed step records alone sit in steps/, the folder verdifax-verify
  # --job reads: it takes every .json file there as a record, so the
  # summary, the sign-offs and the payloads live beside that folder.
  dir.create(file.path(out, "payloads"), showWarnings = FALSE, recursive = TRUE)
  for (row in rows) {
    if (identical(row$status, "pepg_deny")) {
      write_json(file.path(out, "steps", paste0(row$step_id, ".json")), verdifax_get_deny_receipt(agent, row$run_id))
    } else {
      write_text(file.path(out, "steps", paste0(row$step_id, ".json")), verdifax_get_run_artifacts(agent, row$run_id))
    }
  }
  # The exact bytes each step sent, so a reader recomputes the sealed
  # payload hash from the file.
  for (step in names(payloads)) write_text(file.path(out, "payloads", paste0(step, ".json")), canonical(payloads[[step]]))
}

# ── sign-off ───────────────────────────────────────────────────────────────

approver_private_key <- function() {
  if (!file.exists("approver-key.json")) fail("no approver-key.json here; run `claims-review setup` first")
  unhex(jsonlite::fromJSON("approver-key.json", simplifyVector = FALSE)$private_key)
}

cmd_sign_off <- function(job_id, decision, note, receipt_run_id) {
  keys <- load_keys()
  agent <- client_for(keys$agent)
  owner <- client_for(keys$owner)
  key <- approver_private_key()
  if (!is.null(receipt_run_id)) {
    run_id <- as.integer(receipt_run_id)
    receipt <- verdifax_get_deny_receipt(agent, run_id)$deny_receipt
    target <- list(kind = "receipt", run_id = run_id, digest = receipt$hash)
    folders <- Filter(function(p) file.exists(file.path(p, "job.json")) &&
                        identical(as.integer(jsonlite::fromJSON(file.path(p, "job.json"))$refused_run_id), run_id),
                      list.files("records", pattern = "^refused-", full.names = TRUE))
    folder <- if (length(folders) > 0) folders[1] else file.path("records", paste0("receipt-", run_id))
    record_file <- "receipt.json"
  } else {
    folder <- file.path("records", job_id)
    summary <- jsonlite::fromJSON(file.path(folder, "job.json"), simplifyVector = FALSE)
    run_id <- as.integer(summary$final_run_id)
    artifacts <- jsonlite::fromJSON(file.path(folder, "steps", "verdict.json"), simplifyVector = FALSE)
    target <- list(kind = "job", run_id = run_id, digest = artifacts$bundle_anchor$anchored_digest, job_id = job_id)
    record_file <- file.path("steps", "verdict.json")
  }
  record <- verdifax_sign_off(agent, target, decision, key, note = note)
  write_text(file.path(folder, "signoffs.json"), verdifax_get_signoffs(agent, run_id))
  write_text(file.path(folder, "signoff-history.json"), verdifax_get_signoff_history(owner))
  cat(sprintf("sign-off %s recorded as sequence %s, hash %s\n", decision, record$sequence, record$signoff_hash))
  cat(sprintf("verify: verdifax-verify --signoff %s --approver-key approver.txt %s\n", file.path(folder, "signoffs.json"), file.path(folder, record_file)))
  cat(sprintf("verify: verdifax-verify --signoff-history %s\n", file.path(folder, "signoff-history.json")))
}

# ── feed ───────────────────────────────────────────────────────────────────

cmd_feed <- function() {
  keys <- load_keys()
  feed <- verdifax_client(base_url = base_url(), api_key = keys$feed$secret)
  events <- list()
  after <- 0
  repeat {
    page <- verdifax_get_signals(feed, after = after, limit = 500)
    events <- c(events, page$events)
    if (length(page$events) < 500) break
    after <- page$next_after
  }
  write_json("signals.json", events)
  kinds <- table(vapply(events, function(e) e$data$kind, character(1)))
  cat(sprintf("%d events in the feed: %s\n", length(events), paste(sprintf("%s %d", names(kinds), as.integer(kinds)), collapse = ", ")))
  cat("verify: verdifax-verify --signals signals.json\n")
}

# ── verify commands ────────────────────────────────────────────────────────

print_verify_commands <- function(job_id) {
  folder <- file.path("records", job_id)
  cat(sprintf("verify with verdifax-verify %s or later:\n", MIN_VERIFIER))
  cat(sprintf("  verdifax-verify --job %s\n", file.path(folder, "steps")))
  cat(sprintf("  verdifax-verify --account-policy %s %s\n", file.path(REPO, "policy", "claims-review-v1.json"), file.path(folder, "steps", "verdict.json")))
  if (file.exists(file.path(folder, "signoffs.json"))) {
    cat(sprintf("  verdifax-verify --signoff %s --approver-key approver.txt %s\n", file.path(folder, "signoffs.json"), file.path(folder, "steps", "verdict.json")))
  }
  if (file.exists("signals.json")) cat("  verdifax-verify --signals signals.json\n")
}

# ── main ───────────────────────────────────────────────────────────────────

take_option <- function(args, name) {
  i <- match(name, args)
  if (is.na(i)) return(list(value = NULL, args = args))
  if (i == length(args)) fail(paste(name, "needs a value"))
  list(value = args[i + 1], args = args[-c(i, i + 1)])
}

main <- function() {
  args <- commandArgs(trailingOnly = TRUE)
  if (length(args) == 0) fail("commands: setup, review DRAFT.md, sign-off JOB_ID approve|reject, sign-off --receipt RUN_ID reject, feed, verify-commands JOB_ID")
  command <- args[1]
  args <- args[-1]
  if (command == "setup") return(cmd_setup())
  if (command == "review") {
    model <- take_option(args, "--model"); args <- model$args
    answers <- take_option(args, "--answers"); args <- answers$args
    if (length(args) != 1) fail("review takes one draft file")
    return(cmd_review(args[1], if (is.null(model$value)) DEFAULT_MODEL else model$value, answers$value))
  }
  if (command == "sign-off") {
    note <- take_option(args, "--note"); args <- note$args
    receipt <- take_option(args, "--receipt"); args <- receipt$args
    if (length(args) == 0) fail("sign-off needs approve or reject")
    decision <- args[length(args)]
    job_id <- if (length(args) > 1) args[1] else NULL
    if (!decision %in% c("approve", "reject")) fail("sign-off needs approve or reject")
    if (is.null(receipt$value) && is.null(job_id)) fail("sign-off needs a JOB_ID, or --receipt RUN_ID")
    return(cmd_sign_off(job_id, decision, note$value, receipt$value))
  }
  if (command == "feed") return(cmd_feed())
  if (command == "verify-commands") {
    if (length(args) != 1) fail("verify-commands takes a JOB_ID")
    return(print_verify_commands(args[1]))
  }
  fail("commands: setup, review DRAFT.md, sign-off JOB_ID approve|reject, sign-off --receipt RUN_ID reject, feed, verify-commands JOB_ID")
}

if (!nzchar(Sys.getenv("CLAIMS_REVIEW_NO_MAIN"))) main()
