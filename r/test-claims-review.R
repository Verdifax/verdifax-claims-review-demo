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
