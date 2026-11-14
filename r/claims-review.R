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
# A decision means the same on every result: approve agrees with
# the result signed, a verdict or the stop a receipt records, and reject
# judges it wrong. A sign-off releases or blocks nothing on its own.
#
# Commands (the same in every language of this repository):
#
#   claims-review setup                              owner key in VERDIFAX_OWNER_KEY
#   claims-review review DRAFT.md [--model ID]        model key in ANTHROPIC_API_KEY
#   claims-review review DRAFT.md --without-actor     the gate step names no agent
#   claims-review sign-off JOB_ID approve|reject [--note TEXT] [--approver NAME]
#   claims-review sign-off --receipt RUN_ID approve|reject [--note TEXT] [--approver NAME]
#   claims-review publish JOB_ID | publish --receipt RUN_ID
#   claims-review unpublish JOB_ID | unpublish --receipt RUN_ID
#   claims-review approver-add NAME | approver-revoke NAME
#   claims-review settings owner-replacement|verdifax-deletion allowed|blocked
#   claims-review request-key-register KEY_FILE               the key Verdifax made for the chapter
#   claims-review request-key-rotate | request-key-clear [--holder]   the agent key's request key, or the holder's
#   claims-review key-revoke [--holder]                       the owner revokes the agent key, or the holder
#   claims-review delete-step JOB_ID STEP                     the owner deletes one step
#   claims-review webhook-subscribe URL | webhook-list | webhook-remove
#   claims-review feed
#   claims-review verify-commands JOB_ID
#
# Environment: VERDIFAX_BASE_URL (default https://api.verdifax.com).
# Files, in the working folder: demo-keys.json and approver-key.json (0600),
# approver.txt, account-history.json, records/<job_id>/*.json with payloads/,
# signals.json; approver-<name>-key.json (0600), approver-<name>.txt and
# approvers.json; account-keys.json; webhook.json (0600), webhooks.json and
# webhook-deliveries.json.

suppressPackageStartupMessages({
  library(verdifax)
  library(jsonlite)
})

DEFAULT_BASE_URL <- "https://api.verdifax.com"
DEFAULT_MODEL <- "claude-fable-5-1"
FORM <- "verdifax.claims-review.v1"
RULE_SET <- "v3"
ACTOR_ID <- "claims-review-agent"
ROUTE_INTAKE <- "claims-review/v1/intake"
ROUTE_CLAIMS <- "claims-review/v1/claims"
ROUTE_RULE <- "claims-review/v1/rule-check"
ROUTE_VERDICT <- "claims-review/v1/verdict"
OBJECT_PREFIX <- "claims-review/v1/"
LEVELS <- c("PUBLIC", "INTERNAL", "CONFIDENTIAL", "RESTRICTED", "TOP_SECRET", "TS_SCI")
MIN_VERIFIER <- "v0.35.0"
# A stopped step's records live in records/stopped-<run_id>/, and the
# model's exact answers in answers/ beside each job's steps/.
STOPPED_FOLDER_PREFIX <- "stopped-"
ANSWERS_FOLDER <- "answers"
# The text the server's AI review checked for each step, as the server
# returned it once: a reader re-runs the review's check from these files.
TRACES_FOLDER <- "traces"
# The account policy setup activates; v1 to v3 stay for records sealed
# under them. v3 requires a reviewed AI answer on every step after the gate;
# v4 adds that every step names its agent (required_attested_fields), that
# each verdict needs a person's sign-off, and that every step after the gate
# passes its declared data checks.
POLICY_FILE <- "claims-review-v4.json"
# The rule set's pinned hash, which the rule_set_pinned data check compares.
RULE_SET_SHA256 <- "a2afb580a203ccbd084645cfbbce532dd7eaf1d2dbeaa49fd008f66a91f98ae7"
# The data checks every step declares in policy_context.data_quality, each
# computed: draft_unchanged (the draft re-read before the step has the hash
# sealed at intake) and rule_set_pinned (the rule set's hash is the pin).
DATA_QUALITY_CHECKS <- c("draft_unchanged", "rule_set_pinned")
DATA_QUALITY_SOURCE <- "claims-review draft and rule set"
# The owner's two blocks on Verdifax, and their values.
ACCOUNT_SETTINGS <- c("owner-replacement", "verdifax-deletion")
SETTING_VALUES <- c("allowed", "blocked")
APPROVER_NAME <- "^[a-z0-9-]{1,40}$"
# The AI's work on the steps after the gate: listing the draft's claims, and
# explaining the verdict code decided.
CLAIMS_PROMPT <- "claims-prompt-v1.txt"
VERDICT_PROMPT <- "verdict-prompt-v1.txt"
# A step whose record was deleted is written as its tombstone,
# steps/<step>.deleted.json: the run.deleted entry from the account feed,
# which verdifax-verify --job reads in the step's place (R250).
DELETED_STEP_SUFFIX <- ".deleted.json"
# A saved sign-off history or feed is checked as of when it was fetched, so
# every printed re-check states the limit on its head statement's age (R246).
SAVED_HEAD_MAX_AGE <- "168h"
HISTORY_AGE_REASON <- "(168h: a saved history is checked as of when it was fetched; raise the limit for an older copy, or run sign-off again to fetch it fresh)"
# Printed under the account-policy check of a verdict with no sign-off yet:
# policy v4 requires one, so the verifier exits 4 until it is given.
SIGNOFF_PENDING_REASON <- "(exit 4, ANCHORED, SIGN-OFF NOT SHOWN, until a sign-off is given: the policy requires one on the verdict; run sign-off, then verify-commands again)"
FEED_AGE_REASON <- "(168h: a saved feed is checked as of when it was fetched; raise the limit for an older copy, or run feed again to fetch it fresh)"

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

# A file's exact bytes as UTF-8 text: readLines would drop the final newline
# and change the draft's sealed hash (R170).
read_text <- function(path) {
  size <- file.info(path)$size
  if (is.na(size) || size == 0) return("")
  text <- readChar(path, size, useBytes = TRUE)
  Encoding(text) <- "UTF-8"
  text
}
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
  doc <- jsonlite::fromJSON(file.path(REPO, "rules", paste0(RULE_SET, ".json")), simplifyVector = FALSE)
  list(rules = doc$rules, rule_set_hash = sha256_hex(canonical(doc)))
}
load_prompt <- function(name = "prompt-v1.txt") {
  # The exact bytes of the prompt file, so its hash matches every language.
  readChar(file.path(REPO, "rules", name), file.info(file.path(REPO, "rules", name))$size, useBytes = TRUE)
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
# A client for a key; one whose request key was cleared sends unsigned.
client_for <- function(key) {
  if (is.null(key$signing_key) || !nzchar(key$signing_key)) return(verdifax_client(base_url = base_url(), api_key = key$secret))
  verdifax_client(base_url = base_url(), api_key = key$secret, signing_key = unhex(key$signing_key))
}

# The request_substance a step declares: the model's own usage report for the
# step's model call, so the server's budgets count real tokens. The gate step
# calls no model and declares none (NULL).
usage_substance <- function(usage) {
  if (is.null(usage)) return(NULL)
  list(cost = list(input_tokens = usage$input_tokens, output_tokens = usage$output_tokens,
                   total_tokens = usage$input_tokens + usage$output_tokens))
}

# The data checks a step declares, each computed (round 62).
data_quality <- function(intake_draft_sha256, draft_sha256_now, rule_set_hash, as_of) {
  result <- function(ok) if (ok) "pass" else "fail"
  list(source = DATA_QUALITY_SOURCE, as_of = as_of, checks = list(
    list(name = DATA_QUALITY_CHECKS[1], result = result(identical(draft_sha256_now, intake_draft_sha256))),
    list(name = DATA_QUALITY_CHECKS[2], result = result(identical(rule_set_hash, RULE_SET_SHA256)))))
}

# Why a settings command is refused before anything is sent, or NULL.
setting_error <- function(setting, value) {
  if (!setting %in% ACCOUNT_SETTINGS) return(sprintf("the setting must be %s, not %s", paste(ACCOUNT_SETTINGS, collapse = " or "), setting))
  if (!value %in% SETTING_VALUES) return(sprintf("the value must be %s, not %s", paste(SETTING_VALUES, collapse = " or "), value))
  NULL
}

approver_name_error <- function(name) {
  if (!grepl(APPROVER_NAME, name)) return("an approver name is 1 to 40 lowercase letters, digits or dashes")
  NULL
}

# A named approver's private key file (0600) and public key file.
approver_files <- function(name) c(paste0("approver-", name, "-key.json"), paste0("approver-", name, ".txt"))

# A key that was ever on the trial tier must keep a request key, and the
# demo's accounts are trial accounts, so the agent key can replace its request
# key but never clear it. For the request-key chapter Verdifax makes one
# production-tier member key in the account, its creation sealed in the feed
# as Verdifax's: the holder. request-key-register keeps it in demo-keys.json,
# and --holder makes the request-key commands and key-revoke act on it.

# The keys file with one key's request key replaced, or removed when it was
# cleared (NULL); role is "agent" or "holder".
with_signing_key <- function(keys, role, signing_key_hex) {
  entry <- keys[[role]]
  entry$signing_key <- NULL
  if (!is.null(signing_key_hex)) entry$signing_key <- signing_key_hex
  keys[[role]] <- entry
  keys
}

# The keys file with the agent key's request key replaced, or removed when it
# was cleared (NULL).
with_agent_signing_key <- function(keys, signing_key_hex) with_signing_key(keys, "agent", signing_key_hex)

# The key a request-key command or key-revoke acts on: the agent, or with
# --holder the key Verdifax made for the request-key chapter.
key_for <- function(keys, holder) {
  role <- if (isTRUE(holder)) "holder" else "agent"
  if (is.null(keys[[role]])) fail("demo-keys.json holds no request-key holder; run request-key-register with the key Verdifax made first")
  list(role = role, key = keys[[role]])
}

# ── setup ──────────────────────────────────────────────────────────────────

cmd_setup <- function() {
  if (file.exists("demo-keys.json")) fail("demo-keys.json exists here; this folder is already set up, and every program in this repository uses it")
  owner_secret <- Sys.getenv("VERDIFAX_OWNER_KEY", "")
  if (owner_secret == "") fail("set VERDIFAX_OWNER_KEY to the account's owner key")
  owner_signer <- verdifax_generate_request_key()
  owner <- verdifax_client(base_url = base_url(), api_key = owner_secret, signing_key = owner_signer$private_key)
  verdifax_register_account_request_key(owner)
  member <- verdifax_create_account_key(owner, "claims-review-agent")
  member_signer <- verdifax_generate_request_key()
  agent <- verdifax_client(base_url = base_url(), api_key = member$secret, signing_key = member_signer$private_key)
  verdifax_register_request_key(agent)
  policy_text <- read_text(file.path(REPO, "policy", POLICY_FILE))
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

# recorded_tokens is the model use a recorded answer declares, 0 when it
# names none. Only tests against a local server read recorded answers; the
# budget stops are rehearsed with answers that declare large use. Never part
# of the answer.
recorded_tokens <- function(a, field) if (is.null(a[[field]])) 0L else as.integer(a[[field]])

ask <- function(recorded, prompt, user_text, rule_id, model) {
  if (!is.null(recorded)) {
    a <- recorded[[rule_id]]
    return(list(model = if (is.null(a$model)) "recorded-answers" else a$model,
                response_id = if (is.null(a$response_id)) "recorded" else a$response_id,
                verdict = a$verdict, reason = a$reason,
                raw = canonical(list(verdict = a$verdict, reason = a$reason)),
                input_tokens = recorded_tokens(a, "input_tokens"), output_tokens = recorded_tokens(a, "output_tokens")))
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
  req <- model_request(body, key)
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

# model_request builds one call to the model's Messages API. Python, Node and
# Go call the model through Anthropic's official SDKs, which retry a
# temporary failure (a dropped connection, 408, 409, 429 or any 5xx, such as
# 529 overloaded) with growing waits and honor Retry-After. There is no
# official R SDK, so this request carries the same retry policy; without it
# one temporary error stopped the R demo where the other languages carry on.
# The SDKs retry only twice by default, waiting half a second then one
# second. That is too short: on November 13, 2026 a 503 from the model API
# outlasted it and ended the Go demo after twelve sealed steps, failing a
# readiness run. MODEL_RETRY_TRIES sets the number of tries for one model
# call in every language, and the wait grows to a cap of eight seconds, so an
# outage of about two minutes is waited out. A retry can bill a second model
# call when the first answer was lost on the way back; each step's measured
# cost shows it. All four languages hold the same number (demo-parity.json,
# R259).
MODEL_RETRY_TRIES <- 21L
MODEL_TRANSIENT_STATUS <- c(408L, 409L, 429L)
model_transient <- function(resp) {
  status <- httr2::resp_status(resp)
  status %in% MODEL_TRANSIENT_STATUS || status >= 500L
}
model_request <- function(body, key) {
  req <- httr2::request(paste0(sub("/+$", "", Sys.getenv("ANTHROPIC_BASE_URL", "https://api.anthropic.com")), "/v1/messages"))
  req <- httr2::req_headers(req, `x-api-key` = key, `anthropic-version` = "2026-10-21", `content-type` = "application/json")
  req <- httr2::req_body_json(req, body, auto_unbox = TRUE)
  req <- httr2::req_timeout(req, 600)
  httr2::req_retry(req, max_tries = MODEL_RETRY_TRIES, retry_on_failure = TRUE,
                   is_transient = model_transient,
                   backoff = function(attempt) min(8, 2^attempt))
}

# ask_structured asks the model for one answer in the JSON schema given. With
# a recorded answer it returns recorded_raw instead. When the model declines,
# raw is declined, so the step still carries the model's answer to review.
ask_structured <- function(recorded_entry, recorded_raw, prompt, user_text, model, max_tokens, schema, declined) {
  if (!is.null(recorded_entry)) {
    return(list(model = if (is.null(recorded_entry$model)) "recorded-answers" else recorded_entry$model,
                response_id = if (is.null(recorded_entry$response_id)) "recorded" else recorded_entry$response_id,
                raw = recorded_raw, ok = TRUE, input_tokens = recorded_tokens(recorded_entry, "input_tokens"),
                output_tokens = recorded_tokens(recorded_entry, "output_tokens")))
  }
  key <- Sys.getenv("ANTHROPIC_API_KEY", "")
  if (key == "") fail("set ANTHROPIC_API_KEY to call the model")
  body <- list(model = model, max_tokens = as.integer(max_tokens), system = prompt,
               messages = list(list(role = "user", content = user_text)),
               output_config = list(format = list(type = "json_schema", schema = schema)))
  req <- model_request(body, key)
  parsed <- jsonlite::fromJSON(httr2::resp_body_string(httr2::req_perform(req)), simplifyVector = FALSE)
  text_blocks <- Filter(function(b) identical(b$type, "text"), parsed$content)
  raw <- if (length(text_blocks) > 0) text_blocks[[1]]$text else ""
  ok <- !identical(parsed$stop_reason, "refusal") && nzchar(raw)
  list(model = parsed$model, response_id = parsed$id, raw = if (nzchar(raw)) raw else declined, ok = ok,
       input_tokens = as.integer(parsed$usage$input_tokens), output_tokens = as.integer(parsed$usage$output_tokens))
}

# The draft's claims, as the model listed them.
ask_claims <- function(recorded, prompt, user_text, model) {
  entry <- if (is.null(recorded)) NULL else recorded[["claims"]]
  recorded_raw <- if (is.null(entry)) "" else canonical(list(claims = as_list(unlist(entry$claims))))
  schema <- list(type = "object", properties = list(claims = list(type = "array", items = list(type = "string"))),
                 required = list("claims"), additionalProperties = FALSE)
  out <- ask_structured(entry, recorded_raw, prompt, user_text, model, 4000L, schema, "the model declined to list the claims")
  answer <- if (out$ok) tryCatch(jsonlite::fromJSON(out$raw, simplifyVector = FALSE), error = function(e) NULL) else NULL
  out$claims <- if (is.null(answer) || is.null(answer$claims)) character(0) else unlist(answer$claims)
  if (is.null(out$claims)) out$claims <- character(0)
  out
}

# The model's plain-words explanation of the verdict code decided.
ask_explanation <- function(recorded, prompt, user_text, model) {
  entry <- if (is.null(recorded)) NULL else recorded[["verdict"]]
  recorded_raw <- if (is.null(entry)) "" else canonical(list(explanation = entry$explanation))
  schema <- list(type = "object", properties = list(explanation = list(type = "string")),
                 required = list("explanation"), additionalProperties = FALSE)
  out <- ask_structured(entry, recorded_raw, prompt, user_text, model, 2000L, schema, "the model declined to explain the verdict")
  answer <- if (out$ok) tryCatch(jsonlite::fromJSON(out$raw, simplifyVector = FALSE), error = function(e) NULL) else NULL
  out$explanation <- if (is.null(answer) || is.null(answer$explanation)) "" else answer$explanation
  out
}

measured_cost <- function(model, input_tokens, output_tokens) {
  prices <- jsonlite::fromJSON(file.path(REPO, "rules", "model-prices-v1.json"), simplifyVector = FALSE)$prices
  if (is.null(prices[[model]])) return(NULL)
  input_tokens / 1e6 * prices[[model]]$input + output_tokens / 1e6 * prices[[model]]$output
}

# ── review ─────────────────────────────────────────────────────────────────

# A step's policy context: its object, the draft's classification and the
# step's data checks.
policy_context <- function(object_id, level, quality) {
  list(object_type = "text", object_id = paste0(OBJECT_PREFIX, object_id), classification = level, data_quality = quality)
}

cmd_review <- function(draft_path, model, answers_path, without_actor = FALSE) {
  keys <- load_keys()
  agent <- client_for(keys$agent)
  loaded <- load_rules()
  rules <- loaded$rules
  rule_set_hash <- loaded$rule_set_hash
  prompt <- load_prompt()
  draft <- read_text(draft_path)
  draft_hash <- sha256_hex(draft)
  level <- draft_classification(draft)
  # The time the draft was read: the as_of of every step's data checks.
  read_at <- format(Sys.time(), "%Y-%m-%dT%H:%M:%SZ", tz = "UTC")
  # The data checks, computed again now, the draft re-read.
  quality <- function() data_quality(draft_hash, sha256_hex(read_text(draft_path)), rule_set_hash, read_at)
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
                      # The gate step names its agent, as policy v4 requires of
                      # every step; --without-actor leaves the name out on
                      # purpose, so the policy refuses the draft here, before
                      # any model is called.
                      attested_context = if (without_actor) NULL else list(attested = TRUE, actor_id = ACTOR_ID),
                      policy_context = policy_context("intake", level, quality())),
    httr2_http = function(e) e)
  if (inherits(intake, "error")) {
    # The account policy refused the draft at intake (for example a
    # confidential draft), so it never reached a model. The refusal is
    # sealed as a receipt; nothing else runs for this draft.
    stop_job(agent, job_id, "intake", empty_object(), list(intake = intake_payload), list(), draft_path, draft_hash, intake,
             model, list(input = 0L, output = 0L), character(0), list())
    return(invisible(NULL))
  }
  payloads$intake <- intake_payload
  run_ids <- list(intake = intake$run_id)

  results <- list()
  tokens <- list(input = 0L, output = 0L)
  models <- character(0)
  # The model's exact answer, the text Verdifax's AI review reads, kept as
  # sent so a reader recomputes the sealed ai_output_hash from the file.
  answer_texts <- list()
  # The text the server's AI review checked, returned once with each step.
  trace_texts <- list()

  # The gate passed, so the draft may reach a model. Every step from here
  # carries the model's answer, which Verdifax's AI review reads and the
  # record seals; the account policy refuses any step without one.
  claims_prompt <- load_prompt(CLAIMS_PROMPT)
  claims_text <- paste0("Draft:\n", draft)
  claims_prompt_hash <- sha256_hex(paste0(claims_prompt, "\n", claims_text))
  listed <- ask_claims(recorded, claims_prompt, claims_text, model)
  tokens$input <- tokens$input + listed$input_tokens
  tokens$output <- tokens$output + listed$output_tokens
  models <- union(models, listed$model)
  answer_texts[["claims"]] <- listed$raw
  claims_payload <- list(form = FORM, step = "claims", draft_sha256 = draft_hash, rule_set_hash = rule_set_hash,
                         model = listed$model, model_response_id = listed$response_id, prompt_hash = claims_prompt_hash,
                         claims = as_list(listed$claims))
  claims_step <- tryCatch(
    verdifax_job_step(job, "claims", after = list("intake"), payload = canonical(claims_payload), program_id = PROGRAM_ID,
                      route_id = ROUTE_CLAIMS, registry_record_hash = REGISTRY_RECORD_HASH, ai_output_text = listed$raw,
                      request_substance = usage_substance(list(input_tokens = listed$input_tokens, output_tokens = listed$output_tokens)),
                      attested_context = list(attested = TRUE, actor_id = ACTOR_ID, model_provider = "anthropic",
                                              model_name = listed$model, prompt_hash = claims_prompt_hash,
                                              decision_kind = "claims.list", decision_result = as.character(length(listed$claims))),
                      policy_context = policy_context("claims", level, quality())),
    httr2_http = function(e) e)
  if (inherits(claims_step, "error")) {
    payloads$claims <- claims_payload
    if (!is.null(claims_step$aivp_trace_text)) trace_texts[["claims"]] <- claims_step$aivp_trace_text
    stop_job(agent, job_id, "claims", run_ids, payloads, answer_texts, draft_path, draft_hash, claims_step, model, tokens, models,
             trace_texts)
    return(invisible(NULL))
  }
  if (!is.null(claims_step$aivp_trace_text)) trace_texts[["claims"]] <- claims_step$aivp_trace_text
  payloads$claims <- claims_payload
  run_ids$claims <- claims_step$run_id
  cat(sprintf("claims: %d listed\n", length(listed$claims)))

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
    answer_texts[[rule$id]] <- ai_text
    receipt <- tryCatch(
      verdifax_job_step(job, rule$id, after = list("claims"), payload = canonical(payload), program_id = PROGRAM_ID,
                        route_id = ROUTE_RULE, registry_record_hash = REGISTRY_RECORD_HASH, ai_output_text = ai_text,
                        request_substance = usage_substance(list(input_tokens = answer$input_tokens, output_tokens = answer$output_tokens)),
                        attested_context = list(attested = TRUE, actor_id = ACTOR_ID, model_provider = "anthropic",
                                                model_name = answer$model, prompt_hash = prompt_hash,
                                                decision_kind = "claims.rule_check", decision_result = answer$verdict),
                        policy_context = policy_context(rule$id, level, quality())),
      httr2_http = function(e) e)
    if (inherits(receipt, "error")) {
      # The account policy, the server's content screen or its review of
      # the model's answer stopped this step, so the job stops here: the
      # stop is sealed as a receipt, the steps already sealed are kept
      # beside it, and a person sees the stop. The job never gets a verdict.
      payloads[[rule$id]] <- payload
      if (!is.null(receipt$aivp_trace_text)) trace_texts[[rule$id]] <- receipt$aivp_trace_text
      stop_job(agent, job_id, rule$id, run_ids, payloads, answer_texts, draft_path, draft_hash, receipt, model, tokens, models,
               trace_texts)
      return(invisible(NULL))
    }
    if (!is.null(receipt$aivp_trace_text)) trace_texts[[rule$id]] <- receipt$aivp_trace_text
    payloads[[rule$id]] <- payload
    run_ids[[rule$id]] <- receipt$run_id
    results[[length(results) + 1]] <- list(rule_id = rule$id, verdict = answer$verdict, reason = answer$reason)
    cat(sprintf("%s %s: %s\n", rule$id, answer$verdict, answer$reason))
  }

  failed <- unlist(lapply(results, function(r) if (r$verdict != "pass") r$rule_id else NULL))
  if (is.null(failed)) failed <- character(0)
  verdict <- if (length(failed) == 0) "pass" else "fail"
  # The program decides the verdict from the rule answers, so anyone can
  # decide it again from the sealed records; the model only explains it.
  verdict_prompt <- load_prompt(VERDICT_PROMPT)
  verdict_text <- sprintf("Verdict: %s\n\nRule results:\n%s", verdict, canonical(results))
  verdict_prompt_hash <- sha256_hex(paste0(verdict_prompt, "\n", verdict_text))
  explained <- ask_explanation(recorded, verdict_prompt, verdict_text, model)
  tokens$input <- tokens$input + explained$input_tokens
  tokens$output <- tokens$output + explained$output_tokens
  models <- union(models, explained$model)
  answer_texts[["verdict"]] <- explained$raw
  verdict_payload <- list(form = FORM, step = "verdict", draft_sha256 = draft_hash, rule_set_hash = rule_set_hash,
                          verdict = verdict, failed_rules = as_list(failed), results = results, model = explained$model,
                          model_response_id = explained$response_id, prompt_hash = verdict_prompt_hash,
                          explanation = explained$explanation)
  final <- tryCatch(
    verdifax_job_step(job, "verdict", after = as_list(vapply(rules, function(r) r$id, character(1))), final = TRUE,
                      payload = canonical(verdict_payload), program_id = PROGRAM_ID, route_id = ROUTE_VERDICT,
                      registry_record_hash = REGISTRY_RECORD_HASH, ai_output_text = explained$raw,
                      request_substance = usage_substance(list(input_tokens = explained$input_tokens, output_tokens = explained$output_tokens)),
                      attested_context = list(attested = TRUE, actor_id = ACTOR_ID, model_provider = "anthropic",
                                              model_name = explained$model, prompt_hash = verdict_prompt_hash,
                                              decision_kind = "claims.verdict", decision_result = verdict),
                      policy_context = policy_context("verdict", level, quality())),
    httr2_http = function(e) e)
  if (inherits(final, "error")) {
    payloads$verdict <- verdict_payload
    if (!is.null(final$aivp_trace_text)) trace_texts[["verdict"]] <- final$aivp_trace_text
    stop_job(agent, job_id, "verdict", run_ids, payloads, answer_texts, draft_path, draft_hash, final, model, tokens, models,
             trace_texts)
    return(invisible(NULL))
  }
  if (!is.null(final$aivp_trace_text)) trace_texts[["verdict"]] <- final$aivp_trace_text
  payloads$verdict <- verdict_payload
  run_ids$verdict <- final$run_id
  cat(sprintf("verdict explained: %s\n", explained$explanation))

  save_job_records(agent, job_id, out, run_ids, payloads)
  save_answers(out, answer_texts)
  save_texts(file.path(out, TRACES_FOLDER), trace_texts)
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

# A step the server stopped is sealed as one run whose status names the kind
# of stop, with a receipt of that kind (R152). Any other error is not a stop:
# the program reports it as the server gave it and seals nothing for it.
STOP_KINDS <- list(
  pepg_deny = "the account policy or the content screen",
  aivp_t4_halt = "the server's review of the model's answer",
  ccv_halt = "the per-run budget",
  macc_halt = "the cumulative budget",
  attestation_gap = "the policy's required attested fields"
)
STOP_RECEIPT_FIELDS <- list(pepg_deny = "deny_receipt", aivp_t4_halt = "aivp_t4_halt_receipt", ccv_halt = "ccv_halt_receipt",
                            macc_halt = "macc_halt_receipt", attestation_gap = "attestation_gap_receipt")

stop_receipt <- function(agent, status, run_id) {
  fetch <- list(pepg_deny = verdifax_get_deny_receipt, aivp_t4_halt = verdifax_get_aivp_t4_halt_receipt,
                ccv_halt = verdifax_get_ccv_halt_receipt, macc_halt = verdifax_get_macc_halt_receipt,
                attestation_gap = verdifax_get_attestation_gap_receipt)[[status]]
  fetch(agent, run_id)
}

# The sealed hash of a stopped run's receipt, whatever kind of stop it was.
stop_receipt_hash <- function(agent, run_id) {
  status <- verdifax_get_run(agent, run_id)$status
  if (is.null(STOP_KINDS[[status]])) fail(sprintf("run %d is %s, not a stopped run with a receipt", run_id, status))
  stop_receipt(agent, status, run_id)[[STOP_RECEIPT_FIELDS[[status]]]]$hash
}

# The HTTP status an error carries, or NA when it carries none.
error_status <- function(e) {
  if (!is.null(e$status)) return(e$status)
  if (!is.null(e$resp)) return(httr2::resp_status(e$resp))
  NA_integer_
}

# The run the server sealed for this job's stopped step, found by its job and
# step, never by being the newest of the key.
stopped_run <- function(agent, job_id, step, e) {
  rows <- Filter(function(r) identical(r$step_id, step) && !is.null(STOP_KINDS[[r$status]]),
                 verdifax_export_runs(agent, job_id = job_id)$runs)
  if (length(rows) == 0) {
    fail(sprintf("the server answered HTTP %s at %s: %s; it recorded no stop for this step, so nothing was sealed as one",
                 error_status(e), step, error_reason(e)))
  }
  rows[[which.max(vapply(rows, function(r) as.numeric(r$run_id), numeric(1)))]]
}

# The server's own message rides in the error's body; the headline is only
# the HTTP status.
error_reason <- function(e) if (length(e$body) > 0) paste(e$body, collapse = " ") else conditionMessage(e)

# An empty JSON object ({}), as the other programs write an empty run_ids.
empty_object <- function() structure(list(), names = character(0))

# The model's exact answers, one file per step, so each sealed
# ai_output_hash recomputes from its file.
# Writes each step's text to <folder>/<step>.txt, byte for byte.
save_texts <- function(folder, texts) {
  for (step in names(texts)) write_text(file.path(folder, paste0(step, ".txt")), texts[[step]])
}

save_answers <- function(folder, answer_texts) {
  for (step in names(answer_texts)) write_text(file.path(folder, ANSWERS_FOLDER, paste0(step, ".txt")), answer_texts[[step]])
}

# The folder of the stopped step whose receipt is run run_id.
stopped_folder <- function(run_id) file.path("records", paste0(STOPPED_FOLDER_PREFIX, run_id))

# Record a stopped step: its sealed receipt, every record the job sealed
# before it, the exact payload bytes and model answers of every step
# including the stopped one, the model use spent before the stop, and a
# summary naming the step and the kind of stop. The same for intake and for
# any later step.
stop_job <- function(agent, job_id, step, run_ids, payloads, answer_texts, draft_path, draft_hash, e, model, tokens, models,
                     trace_texts) {
  row <- stopped_run(agent, job_id, step, e)
  stopped <- row$run_id
  reason <- error_reason(e)
  folder <- stopped_folder(stopped)  # named by the run, so two stops never share a folder
  write_json(file.path(folder, "receipt.json"), stop_receipt(agent, row$status, stopped))
  # A sealed step whose record was deleted since is kept as its tombstone.
  deleted <- if (length(run_ids) > 0) sealed_only(deleted_steps(read_account_feed(load_keys()), job_id), run_ids) else list()
  write_deleted_steps(folder, deleted)
  for (sealed_step in setdiff(names(run_ids), names(deleted))) {
    write_text(file.path(folder, "steps", paste0(sealed_step, ".json")), verdifax_get_run_artifacts(agent, run_ids[[sealed_step]]))
  }
  for (name in names(payloads)) write_text(file.path(folder, "payloads", paste0(name, ".json")), canonical(payloads[[name]]))
  save_answers(folder, answer_texts)
  save_texts(file.path(folder, TRACES_FOLDER), trace_texts)
  cost <- if (length(models) == 1 && models == model) measured_cost(model, tokens$input, tokens$output) else NULL
  write_json(file.path(folder, "job.json"), list(stopped_run_id = stopped, step = step, stop_status = row$status, job_id = job_id,
                                                 draft = draft_path, draft_sha256 = draft_hash, reason = reason, run_ids = run_ids,
                                                 models = as_list(sort(models)), measured_tokens = tokens, measured_cost_usd = cost))
  cat(sprintf("stopped at %s by %s: run %d, %s\n", step, STOP_KINDS[[row$status]], stopped, reason))
  cat(sprintf("records in %s\n", folder))
  if (!is.null(cost)) {
    cat(sprintf("measured model use before the stop: %d input and %d output tokens, about $%.2f at published prices\n",
                as.integer(tokens$input), as.integer(tokens$output), cost))
  }
}

save_job_records <- function(agent, job_id, out, run_ids, payloads) {
  # A deleted step's record is gone from the export; its tombstone in the
  # account feed stands in its place (R250).
  deleted <- sealed_only(deleted_steps(read_account_feed(load_keys()), job_id), run_ids)
  rows <- Filter(function(r) !(r$step_id %in% names(deleted)), verdifax_export_runs(agent, job_id = job_id)$runs)
  if (length(rows) + length(deleted) != length(run_ids)) {
    fail(sprintf("the export lists %d records and the feed %d deleted steps for %s, the program sealed %d",
                 length(rows), length(deleted), job_id, length(run_ids)))
  }
  # The sealed step records alone sit in steps/, the folder verdifax-verify
  # --job reads: it takes every .json file there as a record, so the
  # summary, the sign-offs and the payloads live beside that folder.
  dir.create(file.path(out, "payloads"), showWarnings = FALSE, recursive = TRUE)
  write_deleted_steps(out, deleted)
  for (row in rows) {
    if (!is.null(STOP_KINDS[[row$status]])) {
      write_json(file.path(out, "steps", paste0(row$step_id, ".json")), stop_receipt(agent, row$status, row$run_id))
    } else {
      write_text(file.path(out, "steps", paste0(row$step_id, ".json")), verdifax_get_run_artifacts(agent, row$run_id))
    }
  }
  # The exact bytes each step sent, so a reader recomputes the sealed
  # payload hash from the file.
  for (step in names(payloads)) write_text(file.path(out, "payloads", paste0(step, ".json")), canonical(payloads[[step]]))
}

# The steps of job job_id whose records were deleted, each with its
# tombstone: the entry of every run.deleted event naming the job, by its
# step. A run.sealed event whose record was deleted carries the tombstone in
# data$deleted with a null entry; it is not the deletion and is passed over,
# as is a deletion of another job or of no job.
deleted_steps <- function(events, job_id) {
  out <- list()
  for (e in events) {
    entry <- e$data$entry
    if (!identical(e$data$kind, "run.deleted") || !is.list(entry) || !identical(entry$job_id, job_id) ||
        !is.character(entry$step_id) || !nzchar(entry$step_id)) next
    out[[entry$step_id]] <- entry
  }
  out
}

# The deleted steps this program sealed.
sealed_only <- function(deleted, run_ids) deleted[intersect(names(deleted), names(run_ids))]

# Writes each deleted step's tombstone as <folder>/steps/<step>.deleted.json,
# every field as the server sent it, beside the step records verdifax-verify
# --job reads.
write_deleted_steps <- function(folder, deleted) {
  for (step in names(deleted)) write_json(file.path(folder, "steps", paste0(step, DELETED_STEP_SUFFIX)), deleted[[step]])
}

# The account's whole feed, read with the feed key, page by page.
read_account_feed <- function(keys) {
  feed <- verdifax_client(base_url = base_url(), api_key = keys$feed$secret)
  events <- list()
  after <- 0
  repeat {
    page <- verdifax_get_signals(feed, after = after, limit = 500)
    events <- c(events, page$events)
    if (length(page$events) < 500) break
    after <- page$next_after
  }
  events
}

# The re-check of a saved sign-off history or feed, with its head age limit,
# and the reason for the limit on the line after (R246).
history_check <- function(path) c(sprintf("verdifax-verify --signoff-history %s --signoff-head-max-age %s", path, SAVED_HEAD_MAX_AGE),
                                  paste0("  ", HISTORY_AGE_REASON))
feed_check <- function() c(sprintf("verdifax-verify --signals signals.json --feed-head-max-age %s", SAVED_HEAD_MAX_AGE),
                           paste0("  ", FEED_AGE_REASON))

# ── sign-off ───────────────────────────────────────────────────────────────

# The first approver's private key, or a named approver's.
approver_private_key <- function(name = NULL) {
  path <- if (is.null(name)) "approver-key.json" else approver_files(name)[1]
  if (!file.exists(path)) fail(sprintf("no %s here; run `claims-review %s` first", path, if (is.null(name)) "setup" else paste("approver-add", name)))
  unhex(jsonlite::fromJSON(path, simplifyVector = FALSE)$private_key)
}

cmd_sign_off <- function(job_id, decision, note, receipt_run_id, approver = NULL) {
  if (!is.null(approver) && !is.null(approver_name_error(approver))) fail(approver_name_error(approver))
  keys <- load_keys()
  agent <- client_for(keys$agent)
  owner <- client_for(keys$owner)
  key <- approver_private_key(approver)
  approver_file <- if (is.null(approver)) "approver.txt" else approver_files(approver)[2]
  if (!is.null(receipt_run_id)) {
    run_id <- as.integer(receipt_run_id)
    target <- list(kind = "receipt", run_id = run_id, digest = stop_receipt_hash(agent, run_id))
    folder <- stopped_folder(run_id)
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
  # R272: a sign-off is checked with the account's whole sign-off history, so
  # a later decision cannot be left out; a rule that requires a sign-off, as
  # policy v4's verdict-step does, fails the check without it.
  cat(sprintf("verify: verdifax-verify --signoff %s --signoff-history %s --signoff-head-max-age %s --approver-key %s %s\n", file.path(folder, "signoffs.json"),
              file.path(folder, "signoff-history.json"), SAVED_HEAD_MAX_AGE, approver_file, file.path(folder, record_file)))
  history <- history_check(file.path(folder, "signoff-history.json"))
  cat(sprintf("verify: %s\n%s\n", history[1], history[2]))
}

# ── feed ───────────────────────────────────────────────────────────────────

# ── publish ────────────────────────────────────────────────────────────────
# The account's owner key makes a record public on its random link, or
# withdraws it. Each decision is sealed by the server as a run publication
# and placed in the account's feed, which `feed` then reads.

# The folder and the records a publish or unpublish acts on: every step of a
# job, or a stopped step's receipt with every step the job sealed before it,
# so a reader sees what the job did before it stopped.
publication_targets <- function(job_id, receipt) {
  if (!is.null(receipt)) {
    run_id <- as.integer(receipt)
    folder <- stopped_folder(run_id)
    summary <- jsonlite::fromJSON(file.path(folder, "job.json"), simplifyVector = FALSE)
    targets <- summary$run_ids
    targets$receipt <- run_id
    return(list(folder = folder, targets = targets))
  }
  folder <- file.path("records", job_id)
  summary <- jsonlite::fromJSON(file.path(folder, "job.json"), simplifyVector = FALSE)
  list(folder = folder, targets = summary$run_ids)
}

# The records a publish or unpublish acts on, sorted: every target but a
# deleted step, which has no record left; its tombstone stands in its place in
# the feed (R250).
publishable <- function(folder, targets) {
  kept <- names(targets)[!file.exists(file.path(folder, "steps", paste0(names(targets), DELETED_STEP_SUFFIX)))]
  sort(kept, method = "radix")
}

cmd_publication <- function(job_id, receipt, publish) {
  keys <- load_keys()
  owner <- client_for(keys$owner)
  where <- publication_targets(job_id, receipt)
  # A repeated publish changes nothing and the server answers with no new
  # publication, so the sealed publication saved earlier is kept.
  earlier_path <- file.path(where$folder, "published.json")
  earlier <- if (file.exists(earlier_path)) jsonlite::fromJSON(earlier_path, simplifyVector = FALSE) else list()
  published <- list()
  for (name in publishable(where$folder, where$targets)) {
    run_id <- as.integer(where$targets[[name]])
    answer <- if (publish) verdifax_publish_run(owner, run_id) else verdifax_unpublish_run(owner, run_id)
    # A completed run is served at its bundle address, a stopped run at its
    # receipt address (R156); each answer names only the one it has.
    sealed <- answer$publication
    before <- earlier[[name]]
    if (is.null(sealed) && !is.null(before) && identical(as.integer(before$run_id), run_id) && identical(before$public, answer$public)) {
      sealed <- before$publication
    }
    published[[name]] <- list(run_id = run_id, public = answer$public, public_bundle_url = answer$public_bundle_url,
                              public_receipt_url = answer$public_receipt_url,
                              verify_url = answer$verify_url, publication = sealed)
  }
  write_json(file.path(where$folder, "published.json"), published)
  cat(sprintf("%s %d record(s) in %s\n", if (publish) "published" else "withdrawn", length(published), where$folder))
  if (publish) {
    shown <- if (!is.null(published$verdict)) published$verdict else if (!is.null(published$receipt)) published$receipt else published[[sort(names(published))[1]]]
    cat(sprintf("public link: %s\n", shown$verify_url))
  }
  feed <- feed_check()
  cat(sprintf("each decision is an event in the account's feed; after `feed`, verify: %s\n%s\n", feed[1], feed[2]))
}

cmd_feed <- function() {
  events <- read_account_feed(load_keys())
  write_json("signals.json", events)
  kinds <- table(vapply(events, function(e) e$data$kind, character(1)))
  cat(sprintf("%d events in the feed: %s\n", length(events), paste(sprintf("%s %d", names(kinds), as.integer(kinds)), collapse = ", ")))
  feed <- feed_check()
  cat(sprintf("verify: %s\n%s\n", feed[1], feed[2]))
}

# ── verify commands ────────────────────────────────────────────────────────

print_verify_commands <- function(job_id) {
  folder <- file.path("records", job_id)
  cat(sprintf("verify with verdifax-verify %s or later:\n", MIN_VERIFIER))
  traces <- if (dir.exists(file.path(folder, TRACES_FOLDER))) paste0(" --aivp-traces ", file.path(folder, TRACES_FOLDER)) else ""
  cat(sprintf("  verdifax-verify --job %s%s\n", file.path(folder, "steps"), traces))
  if (file.exists(file.path(folder, "signoffs.json"))) {
    cat(sprintf("  verdifax-verify --account-policy %s --signoff %s --signoff-history %s --signoff-head-max-age %s --approver-key approver.txt %s\n",
                file.path(REPO, "policy", POLICY_FILE), file.path(folder, "signoffs.json"), file.path(folder, "signoff-history.json"), SAVED_HEAD_MAX_AGE,
                file.path(folder, "steps", "verdict.json")))
  } else {
    cat(sprintf("  verdifax-verify --account-policy %s %s\n", file.path(REPO, "policy", POLICY_FILE), file.path(folder, "steps", "verdict.json")))
    cat(paste0("  ", SIGNOFF_PENDING_REASON, "\n"))
  }
  if (file.exists(file.path(folder, "signoffs.json"))) {
    cat(sprintf("  verdifax-verify --signoff %s --signoff-history %s --signoff-head-max-age %s --approver-key approver.txt %s\n", file.path(folder, "signoffs.json"),
                file.path(folder, "signoff-history.json"), SAVED_HEAD_MAX_AGE, file.path(folder, "steps", "verdict.json")))
  }
  if (file.exists(file.path(folder, "signoff-history.json"))) {
    history <- history_check(file.path(folder, "signoff-history.json"))
    cat(sprintf("  %s\n%s\n", history[1], history[2]))
  }
  if (file.exists("signals.json")) {
    feed <- feed_check()
    cat(sprintf("  %s\n%s\n", feed[1], feed[2]))
  }
}

# ── the account: approvers, settings, keys, deletion, webhooks ─────────────
# Each change is made with the owner key (the agent key for its own request
# key) and sealed by the server in the account's feed, which `feed` saves:
# the feed is the record of every change, so no command keeps its own.

owner_client <- function() client_for(load_keys()$owner)

# The id the server gave this approver when it was added, kept with its key.
approver_id <- function(name) {
  private_path <- approver_files(name)[1]
  if (!file.exists(private_path)) fail(sprintf("no %s here; run `claims-review approver-add %s` first", private_path, name))
  as.integer(jsonlite::fromJSON(private_path, simplifyVector = FALSE)$id)
}

cmd_approver_add <- function(name) {
  if (!is.null(approver_name_error(name))) fail(approver_name_error(name))
  files <- approver_files(name)
  if (file.exists(files[1])) fail(sprintf("%s exists here; this approver is already added", files[1]))
  owner <- owner_client()
  approver <- verdifax_generate_request_key()
  added <- verdifax_register_approver(owner, approver$private_key, paste("Claims review approver", name))
  write_private(files[1], paste0(canonical(list(private_key = hex(approver$private_key), id = added$id)), "\n"))
  write_text(files[2], jsonlite::base64_enc(approver$public_key))
  write_json("approvers.json", verdifax_list_approvers(owner))
  cat(sprintf("approver %s added; its public key is in %s\n", name, files[2]))
}

cmd_approver_revoke <- function(name) {
  if (!is.null(approver_name_error(name))) fail(approver_name_error(name))
  owner <- owner_client()
  verdifax_revoke_approver(owner, approver_id(name))
  write_json("approvers.json", verdifax_list_approvers(owner))
  cat(sprintf("approver %s revoked; the sign-offs it gave stay as they were\n", name))
}

cmd_settings <- function(setting, value) {
  problem <- setting_error(setting, value)
  if (!is.null(problem)) fail(problem)
  verdifax_set_account_setting(owner_client(), setting, value)
  cat(sprintf("%s is %s\n", setting, value))
}

write_keys <- function(keys) {
  write_private("demo-keys.json", paste0(as.character(jsonlite::toJSON(keys, auto_unbox = TRUE, pretty = TRUE)), "\n"))
}

cmd_request_key_register <- function(key_file) {
  keys <- load_keys()
  if (!is.null(keys$holder)) fail("demo-keys.json already holds the request-key holder")
  made <- jsonlite::fromJSON(key_file, simplifyVector = FALSE)
  signer <- verdifax_generate_request_key()
  holder <- verdifax_client(base_url = base_url(), api_key = made$secret, signing_key = signer$private_key)
  verdifax_register_request_key(holder)
  keys$holder <- list(id = made$id, secret = made$secret, signing_key = hex(signer$private_key))
  write_keys(keys)
  cat(sprintf("a request key was registered on key %s, the key Verdifax made for the request-key chapter\n", made$id))
}

cmd_request_key_rotate <- function(holder) {
  keys <- load_keys()
  chosen <- key_for(keys, holder)
  next_key <- verdifax_generate_request_key()
  verdifax_rotate_request_key(client_for(chosen$key), next_key$private_key)
  write_keys(with_signing_key(keys, chosen$role, hex(next_key$private_key)))
  cat(sprintf("the %s key's request key was replaced, approved by the key it replaced\n", chosen$role))
}

cmd_request_key_clear <- function(holder) {
  keys <- load_keys()
  chosen <- key_for(keys, holder)
  verdifax_clear_request_key(client_for(chosen$key))
  write_keys(with_signing_key(keys, chosen$role, NULL))
  cat(sprintf("the %s key's request key was cleared, approved by that key\n", chosen$role))
}

cmd_key_revoke <- function(holder) {
  keys <- load_keys()
  chosen <- key_for(keys, holder)
  owner <- client_for(keys$owner)
  verdifax_revoke_account_key(owner, as.integer(chosen$key$id))
  write_json("account-keys.json", verdifax_list_account_keys(owner))
  cat(sprintf("%s key %s revoked; the records it sealed stay as they were\n", chosen$role, chosen$key$id))
}

# A deleted step's record is replaced by its tombstone, which verdifax-verify
# --job reads in its place (R250).
replace_with_tombstone <- function(folder, step, tombstone) {
  write_deleted_steps(folder, stats::setNames(list(tombstone), step))
  record <- file.path(folder, "steps", paste0(step, ".json"))
  if (file.exists(record)) file.remove(record)
  invisible(NULL)
}

cmd_delete_step <- function(job_id, step) {
  folder <- file.path("records", job_id)
  record_path <- file.path(folder, "steps", paste0(step, ".json"))
  if (!file.exists(record_path)) fail(sprintf("%s does not exist; delete-step needs the step's sealed record", record_path))
  record <- jsonlite::fromJSON(record_path, simplifyVector = FALSE)
  verdifax_delete_run(owner_client(), as.integer(record$run_id), record$manifest_hash)
  tombstone <- deleted_steps(read_account_feed(load_keys()), job_id)[[step]]
  if (is.null(tombstone)) fail(sprintf("the feed holds no run.deleted event for %s of %s", step, job_id))
  replace_with_tombstone(folder, step, tombstone)
  cat(sprintf("step %s of %s deleted; its tombstone is %s\n", step, job_id, file.path(folder, "steps", paste0(step, DELETED_STEP_SUFFIX))))
  cat(sprintf("verify: verdifax-verify --job %s\n", file.path(folder, "steps")))
}

cmd_webhook_subscribe <- function(url) {
  hook <- verdifax_create_webhook(owner_client(), url)
  write_private("webhook.json", paste0(as.character(jsonlite::toJSON(hook, auto_unbox = TRUE, pretty = TRUE, null = "null")), "\n"))
  cat(sprintf("webhook %s subscribed to every kind of event at %s; its secret is in webhook.json\n", hook$id, url))
}

# Removes the subscription webhook-subscribe made, as the owner, sealed in the
# feed as webhook.deleted. webhook.json is kept: its secret can be published
# once the subscription is gone, since it then signs nothing.
cmd_webhook_remove <- function() {
  if (!file.exists("webhook.json")) fail("no webhook.json here; run `claims-review webhook-subscribe URL` first")
  id <- as.integer(jsonlite::fromJSON("webhook.json", simplifyVector = FALSE)$id)
  write_json("webhook-removed.json", verdifax_delete_webhook(owner_client(), id))
  cat(sprintf("webhook %s removed; its removal is sealed in the account's feed\n", id))
}

# The lists alone, the same shape in every language: webhooks.json the
# account's webhooks, webhook-deliveries.json each one's deliveries by id.
cmd_webhook_list <- function() {
  owner <- owner_client()
  hooks <- verdifax_list_webhooks(owner)$webhooks
  if (is.null(hooks)) hooks <- list()
  write_json("webhooks.json", hooks)
  deliveries <- empty_object()
  for (h in hooks) {
    got <- verdifax_list_webhook_deliveries(owner, as.integer(h$id), limit = 500)$deliveries
    deliveries[[as.character(h$id)]] <- if (is.null(got)) list() else got
  }
  write_json("webhook-deliveries.json", deliveries)
  cat(sprintf("%d webhook(s); their deliveries are in webhook-deliveries.json\n", length(hooks)))
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
  if (length(args) == 0) fail("commands: setup, review DRAFT.md, sign-off JOB_ID approve|reject, sign-off --receipt RUN_ID approve|reject, publish JOB_ID, publish --receipt RUN_ID, unpublish JOB_ID, unpublish --receipt RUN_ID, approver-add NAME, approver-revoke NAME, settings SETTING VALUE, request-key-register KEY_FILE, request-key-rotate [--holder], request-key-clear [--holder], key-revoke [--holder], delete-step JOB_ID STEP, webhook-subscribe URL, webhook-list, webhook-remove, feed, verify-commands JOB_ID")
  command <- args[1]
  args <- args[-1]
  if (command == "setup") return(cmd_setup())
  if (command == "review") {
    model <- take_option(args, "--model"); args <- model$args
    answers <- take_option(args, "--answers"); args <- answers$args
    without_actor <- "--without-actor" %in% args
    args <- args[args != "--without-actor"]
    if (length(args) != 1) fail("review takes one draft file")
    return(cmd_review(args[1], if (is.null(model$value)) DEFAULT_MODEL else model$value, answers$value, without_actor))
  }
  if (command == "sign-off") {
    note <- take_option(args, "--note"); args <- note$args
    receipt <- take_option(args, "--receipt"); args <- receipt$args
    approver <- take_option(args, "--approver"); args <- approver$args
    if (length(args) == 0) fail("sign-off needs approve or reject")
    decision <- args[length(args)]
    job_id <- if (length(args) > 1) args[1] else NULL
    if (!decision %in% c("approve", "reject")) fail("sign-off needs approve or reject")
    if (is.null(receipt$value) && is.null(job_id)) fail("sign-off needs a JOB_ID, or --receipt RUN_ID")
    return(cmd_sign_off(job_id, decision, note$value, receipt$value, approver$value))
  }
  if (command %in% c("publish", "unpublish")) {
    receipt <- take_option(args, "--receipt"); args <- receipt$args
    job_id <- if (length(args) > 0) args[1] else NULL
    if (is.null(receipt$value) && is.null(job_id)) fail(paste(command, "needs a JOB_ID, or --receipt RUN_ID"))
    return(cmd_publication(job_id, receipt$value, command == "publish"))
  }
  if (command %in% c("approver-add", "approver-revoke")) {
    if (length(args) != 1) fail(paste(command, "takes a NAME"))
    return(if (command == "approver-add") cmd_approver_add(args[1]) else cmd_approver_revoke(args[1]))
  }
  if (command == "settings") {
    if (length(args) != 2) fail("settings takes a SETTING and a VALUE")
    return(cmd_settings(args[1], args[2]))
  }
  if (command == "request-key-register") {
    if (length(args) != 1) fail("request-key-register takes the KEY_FILE of the key Verdifax made")
    return(cmd_request_key_register(args[1]))
  }
  if (command %in% c("request-key-rotate", "request-key-clear", "key-revoke")) {
    holder <- "--holder" %in% args
    if (length(setdiff(args, "--holder")) != 0) fail(paste(command, "takes only --holder"))
    if (command == "request-key-rotate") return(cmd_request_key_rotate(holder))
    if (command == "request-key-clear") return(cmd_request_key_clear(holder))
    return(cmd_key_revoke(holder))
  }
  if (command == "delete-step") {
    if (length(args) != 2) fail("delete-step takes a JOB_ID and a STEP")
    return(cmd_delete_step(args[1], args[2]))
  }
  if (command == "webhook-subscribe") {
    if (length(args) != 1) fail("webhook-subscribe takes a URL")
    return(cmd_webhook_subscribe(args[1]))
  }
  if (command == "webhook-list") return(cmd_webhook_list())
  if (command == "webhook-remove") return(cmd_webhook_remove())
  if (command == "feed") return(cmd_feed())
  if (command == "verify-commands") {
    if (length(args) != 1) fail("verify-commands takes a JOB_ID")
    return(print_verify_commands(args[1]))
  }
  fail("commands: setup, review DRAFT.md, sign-off JOB_ID approve|reject, sign-off --receipt RUN_ID approve|reject, publish JOB_ID, publish --receipt RUN_ID, unpublish JOB_ID, unpublish --receipt RUN_ID, approver-add NAME, approver-revoke NAME, settings SETTING VALUE, request-key-register KEY_FILE, request-key-rotate [--holder], request-key-clear [--holder], key-revoke [--holder], delete-step JOB_ID STEP, webhook-subscribe URL, webhook-list, webhook-remove, feed, verify-commands JOB_ID")
}

if (!nzchar(Sys.getenv("CLAIMS_REVIEW_NO_MAIN"))) main()
