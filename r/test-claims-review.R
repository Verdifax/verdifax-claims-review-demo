# Unit tests for the R program: the same cases every language tests.
Sys.setenv(CLAIMS_REVIEW_NO_MAIN = "1")
here <- normalizePath(dirname(sub("^--file=", "", grep("^--file=", commandArgs(trailingOnly = FALSE), value = TRUE)[1])))
source(file.path(here, "claims-review.R"))
check <- function(ok, what) if (!isTRUE(ok)) stop("FAIL: ", what, call. = FALSE) else cat("ok   ", what, "\n")

fixed <- list(b = 1L, a = "x", c = list(1L, 2L), d = list(z = TRUE, y = NULL), e = "ü \"q\" \\ /")
expected <- '{"a":"x","b":1,"c":[1,2],"d":{"y":null,"z":true},"e":"ü \\"q\\" \\\\ /"}'
check(identical(canonical(fixed), expected), "canonical JSON matches every language")

loaded <- load_rules()
check(length(loaded$rules) == 15 && loaded$rule_set_hash == "61110740e99a53d6fe48378a52a450f053ef2d887ef491be55801f75b4e4e1ad", "the rule set hash")

check(draft_classification("# A draft\n\nText.") == "INTERNAL" &&
        draft_classification("<!-- classification: CONFIDENTIAL -->\n# Memo") == "CONFIDENTIAL", "the classification line")

source_text <- readChar(file.path(REPO, "demo-parity.json"), file.info(file.path(REPO, "demo-parity.json"))$size, useBytes = TRUE)
copy_text <- readChar(file.path(here, "testdata", "demo-parity.json"), file.info(file.path(here, "testdata", "demo-parity.json"))$size, useBytes = TRUE)
check(identical(source_text, copy_text), "the parity copy matches the source")
