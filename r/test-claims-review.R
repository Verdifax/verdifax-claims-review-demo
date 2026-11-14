# Unit tests for the R program: the same cases every language tests.
Sys.setenv(CLAIMS_REVIEW_NO_MAIN = "1")
here <- normalizePath(dirname(sub("^--file=", "", grep("^--file=", commandArgs(trailingOnly = FALSE), value = TRUE)[1])))
source(file.path(here, "claims-review.R"))
check <- function(ok, what) if (!isTRUE(ok)) stop("FAIL: ", what, call. = FALSE) else cat("ok   ", what, "\n")

fixed <- list(b = 1L, a = "x", c = list(1L, 2L), d = list(z = TRUE, y = NULL), e = "ü \"q\" \\ /")
expected <- '{"a":"x","b":1,"c":[1,2],"d":{"y":null,"z":true},"e":"ü \\"q\\" \\\\ /"}'
check(identical(canonical(fixed), expected), "canonical JSON matches every language")

# The rule set in force and its hash are named once, in demo-parity.json.
parity <- jsonlite::fromJSON(file.path(REPO, "demo-parity.json"), simplifyVector = FALSE)
loaded <- load_rules()
check(length(loaded$rules) == 15 && RULE_SET == parity$constants$rule_set &&
        loaded$rule_set_hash == parity$rule_set_sha256, "the rule set hash")

check(draft_classification("# A draft\n\nText.") == "INTERNAL" &&
        draft_classification("<!-- classification: CONFIDENTIAL -->\n# Memo") == "CONFIDENTIAL", "the classification line")

source_text <- readChar(file.path(REPO, "demo-parity.json"), file.info(file.path(REPO, "demo-parity.json"))$size, useBytes = TRUE)
copy_text <- readChar(file.path(here, "testdata", "demo-parity.json"), file.info(file.path(here, "testdata", "demo-parity.json"))$size, useBytes = TRUE)
check(identical(source_text, copy_text), "the parity copy matches the source")

# The model call retries a temporary failure as the official SDKs the other
# three languages use do (round 58: one temporary error stopped the R demo in
# a readiness run while Python, Node and Go would have retried), with the
# shared budget every language holds and waits short enough to reach it
# (round 61: the SDKs' default of two retries waited about a second and a
# half, which a 503 from the model API outlasted; R259).
req <- model_request(list(model = "m", max_tokens = 1L), "k")
tries <- as.integer(jsonlite::fromJSON(source_text)$constants$model_retry_tries)
waits <- vapply(seq_len(tries - 1L), req$policies$retry_backoff, 0)
check(identical(req$policies$retry_max_tries, MODEL_RETRY_TRIES) &&
        identical(MODEL_RETRY_TRIES, tries) && isTRUE(req$policies$retry_on_failure) &&
        sum(waits) >= 60,
      "the model call retries the shared number of times, after a dropped connection too")
transient <- vapply(c(408L, 409L, 429L, 500L, 503L, 529L), function(code) model_transient(httr2::response(status_code = code)), TRUE)
permanent <- vapply(c(400L, 401L, 403L, 404L, 422L), function(code) model_transient(httr2::response(status_code = code)), TRUE)
check(all(transient) && !any(permanent), "only a temporary model error is retried")

# The stop cases every language runs, from one file (R152). The test answers
# as the server would from stop-cases.json: each receipt route serves only
# runs stopped by its own kind, so a fetch of the wrong kind fails the test.
stops <- jsonlite::fromJSON(file.path(REPO, "testdata", "stop-cases.json"), simplifyVector = FALSE)
stop_status <- setNames(lapply(stops$receipts, `[[`, "status"), vapply(stops$receipts, function(r) as.character(r$run_id), ""))
stop_status[[as.character(stops$not_stopped$run_id)]] <- stops$not_stopped$status
receipt_route <- function(kind) function(agent, run_id) {
  if (!identical(stop_status[[as.character(run_id)]], kind)) {
    stop(sprintf("the %s route was asked for run %d, which is %s", kind, run_id, stop_status[[as.character(run_id)]]))
  }
  r <- Filter(function(x) x$run_id == run_id, stops$receipts)[[1]]
  setNames(list(list(hash = r$hash)), r$field)
}
verdifax_get_run <- function(agent, run_id) list(status = stop_status[[as.character(run_id)]])
verdifax_export_runs <- function(agent, job_id) list(runs = if (identical(job_id, stops$job_id)) stops$export_rows else list())
verdifax_get_deny_receipt <- receipt_route("pepg_deny")
verdifax_get_aivp_t4_halt_receipt <- receipt_route("aivp_t4_halt")
verdifax_get_ccv_halt_receipt <- receipt_route("ccv_halt")
verdifax_get_macc_halt_receipt <- receipt_route("macc_halt")
verdifax_get_attestation_gap_receipt <- receipt_route("attestation_gap")
# fail() prints and quits; here it raises the message instead.
fail <- function(message) stop(structure(class = c("claims_review_fail", "error", "condition"), list(message = message, call = NULL)))
failure_of <- function(expr) tryCatch({ force(expr); "" }, claims_review_fail = function(e) conditionMessage(e))

constants <- parity$constants
check(identical(sort(names(STOP_KINDS)), sort(unlist(constants$stop_statuses))) &&
        identical(STOP_RECEIPT_FIELDS[unlist(constants$stop_statuses)], setNames(constants$stop_receipt_fields, unlist(constants$stop_statuses))),
      "the stop kinds match the parity file")

api_error <- structure(class = c("httr2_http_error", "error", "condition"),
                       list(message = sprintf("HTTP %d", stops$error$status), status = stops$error$status, body = stops$error$message))
found <- vapply(stops$stopped_run, function(c) {
  if (!is.null(c$failure)) return(identical(failure_of(stopped_run(NULL, stops$job_id, c$step, api_error)), c$failure))
  row <- stopped_run(NULL, stops$job_id, c$step, api_error)
  identical(as.numeric(row$run_id), as.numeric(c$run_id)) && identical(row$status, c$status)
}, logical(1))
check(all(found), "a stopped run is found by its job and step")

read <- vapply(stops$receipts, function(c) {
  identical(stop_receipt_hash(NULL, c$run_id), c$hash) &&
    identical(stop_receipt(NULL, c$status, c$run_id), setNames(list(list(hash = c$hash)), c$field))
}, logical(1))
check(all(read), "a stop's receipt hash is read from the receipt of its kind")

check(identical(failure_of(stop_receipt_hash(NULL, stops$not_stopped$run_id)), stops$not_stopped$failure), "a run that did not stop has no receipt")

# R170: the program reads every input as its exact bytes, so the hash it seals
# is the one shasum prints, and the parity file lists every sample and policy.
listed <- sort(names(parity$input_sha256))
on_disk <- sort(c(paste0("samples/", list.files(file.path(REPO, "samples"))), paste0("policy/", list.files(file.path(REPO, "policy")))))
check(identical(listed, on_disk), "the parity file lists every sample and policy file")
read_ok <- vapply(listed, function(f) identical(sha256_hex(read_text(file.path(REPO, f))), parity$input_sha256[[f]]), logical(1))
check(all(read_ok), "every input file reads as its exact bytes")

# R250: a deleted step's tombstone, from the cases every language runs.
deleted_cases <- jsonlite::fromJSON(file.path(REPO, "testdata", "deleted-step-cases.json"), simplifyVector = FALSE)
check(identical(deleted_steps(deleted_cases$events, deleted_cases$job_id), deleted_cases$want) &&
        length(deleted_steps(deleted_cases$events, "claims-review-no-such-job")) == 0 &&
        identical(DELETED_STEP_SUFFIX, constants$deleted_step_suffix),
      "a deleted step is read from its job's run.deleted events only")
# The SDK hands over the feed as parsed JSON, so the program writes every
# field again: each written tombstone must equal the served one field for
# field, job_links included (round 61): a list of one stays a list, final
# stays true, and every number keeps its digits.
tomb_dir <- tempfile("claims-review-")
write_deleted_steps(tomb_dir, deleted_steps(deleted_cases$events, deleted_cases$job_id))
kept_ok <- vapply(names(deleted_cases$want), function(step) {
  path <- file.path(tomb_dir, "steps", paste0(step, DELETED_STEP_SUFFIX))
  text <- paste(readLines(path, warn = FALSE), collapse = "")
  kept <- jsonlite::fromJSON(text, simplifyVector = FALSE)
  served <- deleted_cases$want[[step]]
  compact <- gsub("[[:space:]]", "", text)
  identical(kept, served) && is.list(kept$job_links$previous_record_hashes) &&
    grepl('"previous_record_hashes":[', compact, fixed = TRUE) &&
    all(vapply(c(sprintf('"run_id":%s', format(served$run_id, scientific = FALSE)), sprintf('"log_index":%s', format(served$anchor$log_index, scientific = FALSE)),
                 sprintf('"integrated_time":%s', format(served$anchor$integrated_time, scientific = FALSE))),
               function(literal) grepl(literal, compact, fixed = TRUE), logical(1)))
}, logical(1))
check(identical(sort(list.files(file.path(tomb_dir, "steps"))), c("rule-02.deleted.json", "verdict.deleted.json")) && all(kept_ok) &&
        identical(deleted_cases$want$verdict$job_links$final, TRUE),
      "a deleted step's tombstone is written beside the steps with every field, job_links included")
unlink(tomb_dir, recursive = TRUE)

# R246: every printed re-check of a saved history or feed states its head age
# limit, with the reason on the line after.
verify_dir <- tempfile("claims-review-")
dir.create(file.path(verify_dir, "records", "job-1"), recursive = TRUE)
old_wd <- setwd(verify_dir)
writeLines("[]", file.path("records", "job-1", "signoff-history.json"))
writeLines("[]", "signals.json")
printed <- capture.output(print_verify_commands("job-1"))
setwd(old_wd)
unlink(verify_dir, recursive = TRUE)
history_at <- grep("--signoff-history", printed, fixed = TRUE)[1]
feed_at <- grep("--signals", printed, fixed = TRUE)[1]
check(identical(MIN_VERIFIER, constants$min_verifier) && identical(printed[1], sprintf("verify with verdifax-verify %s or later:", MIN_VERIFIER)) &&
        identical(printed[history_at], sprintf("  verdifax-verify --signoff-history %s --signoff-head-max-age %s", file.path("records", "job-1", "signoff-history.json"), SAVED_HEAD_MAX_AGE)) &&
        identical(printed[history_at + 1], paste0("  ", HISTORY_AGE_REASON)) &&
        identical(printed[feed_at], sprintf("  verdifax-verify --signals signals.json --feed-head-max-age %s", SAVED_HEAD_MAX_AGE)) &&
        identical(printed[feed_at + 1], paste0("  ", FEED_AGE_REASON)) &&
        identical(c(SAVED_HEAD_MAX_AGE, HISTORY_AGE_REASON, FEED_AGE_REASON),
                  c(constants$saved_head_max_age, constants$history_age_reason, constants$feed_age_reason)),
      "every printed re-check of a saved history or feed states its head age limit")

# Round 62: the cases every language runs for the every-capability demo.
usage_cases <- jsonlite::fromJSON(file.path(REPO, "testdata", "usage-cases.json"), simplifyVector = FALSE)
for (c in usage_cases$cases) {
  got <- usage_substance(c$usage)
  check(if (is.null(c$want)) is.null(got) else identical(canonical(got), c$want_canonical), paste("usage declared:", c$name))
}
for (c in usage_cases$recorded) {
  check(identical(c(recorded_tokens(c$answer, "input_tokens"), recorded_tokens(c$answer, "output_tokens")),
                  c(as.integer(c$want$input_tokens), as.integer(c$want$output_tokens))), paste("recorded usage:", c$name))
}
quality_cases <- jsonlite::fromJSON(file.path(REPO, "testdata", "data-quality-cases.json"), simplifyVector = FALSE)
check(identical(RULE_SET_SHA256, quality_cases$rule_set_sha256) && identical(RULE_SET_SHA256, constants$rule_set_sha256) &&
        identical(DATA_QUALITY_CHECKS, unlist(constants$data_quality_checks)) && identical(DATA_QUALITY_SOURCE, constants$data_quality_source),
      "the data checks' names, source and pin match the parity file")
for (c in quality_cases$cases) {
  got <- data_quality(c$intake_draft_sha256, c$draft_sha256_now, c$rule_set_hash, quality_cases$as_of)
  check(identical(canonical(got), c$want_canonical), paste("data checks computed:", c$name))
}
check(identical(load_rules()$rule_set_hash, RULE_SET_SHA256), "the loaded rule set is the pinned one")
policy_v4 <- jsonlite::fromJSON(file.path(REPO, "policy", POLICY_FILE), simplifyVector = FALSE)
check(identical(POLICY_FILE, constants$policy_file) && identical(unlist(policy_v4$required_attested_fields), "actor_id"), "the policy is v4")

command_cases <- jsonlite::fromJSON(file.path(REPO, "testdata", "command-cases.json"), simplifyVector = FALSE)
check(all(vapply(command_cases$settings$valid, function(v) is.null(setting_error(v[[1]], v[[2]])), logical(1))) &&
        all(vapply(command_cases$settings$invalid, function(v) identical(setting_error(v$args[[1]], v$args[[2]]), v$error), logical(1))),
      "settings are checked before sending")
names_case <- command_cases$approver_names
check(all(vapply(names_case$valid, function(n) is.null(approver_name_error(n)), logical(1))) &&
        all(vapply(names_case$invalid, function(n) identical(approver_name_error(n), names_case$error), logical(1))) &&
        identical(approver_files(names_case$files$name), c(names_case$files$private_key, names_case$files$public_key)),
      "approver names and files")
registry_case <- command_cases$registry_record
check(identical(REGISTRY_RECORD_HASH, sha256_hex(registry_case$default_source)) &&
        identical(registry_record_hash("")$hash, REGISTRY_RECORD_HASH) && identical(registry_record_hash(NULL)$hash, REGISTRY_RECORD_HASH) &&
        identical(registry_record_hash(registry_case$given)$hash, registry_case$given) &&
        all(vapply(registry_case$invalid, function(bad) identical(registry_record_hash(bad)$error, registry_case$error), logical(1))),
      "the registry record a review claims")
local_case <- command_cases$local_server
with_base <- function(url, f) { old <- Sys.getenv("VERDIFAX_BASE_URL", NA); Sys.setenv(VERDIFAX_BASE_URL = url); on.exit(if (is.na(old)) Sys.unsetenv("VERDIFAX_BASE_URL") else Sys.setenv(VERDIFAX_BASE_URL = old)); f() }
check(all(vapply(local_case$local, function(u) with_base(u, is_local_server), logical(1))) &&
        !any(vapply(local_case$live, function(u) with_base(u, is_local_server), logical(1))) &&
        grepl(sprintf('fail("%s")', local_case$error), paste(readLines(file.path(REPO, "r", "claims-review.R")), collapse = "\n"), fixed = TRUE),
      "recorded answers only against a local server")
key_case <- command_cases$request_key
check(identical(canonical(with_agent_signing_key(key_case$keys, key_case$rotated_to)), canonical(key_case$after_rotate)) &&
        identical(canonical(with_agent_signing_key(key_case$keys, NULL)), canonical(key_case$after_clear)) &&
        !is.null(client_for(key_case$after_clear$agent)),
      "the agent's request key in the keys file")
holder_case <- key_case$holder
check(identical(canonical(with_signing_key(holder_case$keys, "holder", holder_case$rotated_to)), canonical(holder_case$after_rotate)) &&
        identical(canonical(with_signing_key(holder_case$keys, "holder", NULL)), canonical(holder_case$after_clear)) &&
        identical(key_for(holder_case$keys, TRUE)$role, "holder") && identical(key_for(holder_case$keys, TRUE)$key$id, 9L) &&
        identical(key_for(holder_case$keys, FALSE)$role, "agent"),
      "the holder's request key in the keys file, with --holder")
delete_case <- command_cases$delete_step
tomb <- deleted_steps(delete_case$feed, delete_case$job_id)[[delete_case$step]]
delete_dir <- tempfile("claims-review-")
dir.create(file.path(delete_dir, "steps"), recursive = TRUE)
writeLines(canonical(delete_case$record), file.path(delete_dir, "steps", paste0(delete_case$step, ".json")))
replace_with_tombstone(delete_dir, delete_case$step, tomb)
written <- jsonlite::fromJSON(file.path(delete_dir, delete_case$want_files[[1]]), simplifyVector = FALSE)
check(identical(canonical(tomb), canonical(delete_case$want_tombstone)) && identical(canonical(written), canonical(delete_case$want_tombstone)) &&
        !file.exists(file.path(delete_dir, delete_case$want_removed[[1]])),
      "a deleted step is replaced by its tombstone")
unlink(delete_dir, recursive = TRUE)
missing_dir <- tempfile("claims-review-")
dir.create(missing_dir)
old_wd <- setwd(missing_dir)
missing_out <- suppressWarnings(system2(file.path(R.home("bin"), "Rscript"), c(file.path(here, "claims-review.R"), "delete-step", delete_case$job_id, "rule-03"),
                                        stdout = TRUE, stderr = TRUE, env = "CLAIMS_REVIEW_NO_MAIN="))
setwd(old_wd)
unlink(missing_dir, recursive = TRUE)
check(identical(attr(missing_out, "status"), 1L) && any(grepl(delete_case$missing_step_error, missing_out, fixed = TRUE)),
      "delete-step needs the step's sealed record")

pub_case <- command_cases$publication
pub_dir <- tempfile("claims-review-")
dir.create(file.path(pub_dir, "steps"), recursive = TRUE)
for (step in pub_case$tombstones) writeLines("{}", file.path(pub_dir, "steps", paste0(step, DELETED_STEP_SUFFIX)))
check(identical(publishable(pub_dir, pub_case$targets), unlist(pub_case$want)), "publication skips a deleted step")
unlink(pub_dir, recursive = TRUE)

vc <- command_cases$verify_commands
policy_path <- file.path(REPO, "policy", POLICY_FILE)
vc_want <- function(lines) vapply(lines, function(l) sub("SIGNOFF_PENDING_REASON", SIGNOFF_PENDING_REASON, sub("POLICY", policy_path, l, fixed = TRUE), fixed = TRUE), "", USE.NAMES = FALSE)
vc_printed <- function() {
  out <- capture.output(print_verify_commands(vc$job_id))
  out[grepl("--account-policy ", out, fixed = TRUE) | startsWith(out, "  verdifax-verify --signoff ") | out == paste0("  ", SIGNOFF_PENDING_REASON)]
}
vc_dir <- tempfile("claims-review-")
dir.create(file.path(vc_dir, "records", vc$job_id, "steps"), recursive = TRUE)
old_wd <- setwd(vc_dir)
without <- vc_printed()
writeLines("[]", file.path("records", vc$job_id, "signoffs.json"))
with_signoffs <- vc_printed()
setwd(old_wd)
unlink(vc_dir, recursive = TRUE)
check(identical(SIGNOFF_PENDING_REASON, constants$signoff_pending_reason) &&
        identical(without, vc_want(vc$without_signoffs)) && identical(with_signoffs, vc_want(vc$with_signoffs)),
      "the account-policy check passes the sign-offs once given")
