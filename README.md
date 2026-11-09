# Verdifax claims review demo

A small agent reads a draft (a web page, a press release, a sales email) and checks it against 15 general writing rules, such as "don't call anything unbreakable" and "don't promise a legal result". For each rule it returns pass or fail with a reason. A person decides. Verdifax seals every step, so anyone can check afterwards, offline, which rules were checked, what the model answered, which policy admitted each step, and who signed off.

The same program is here in Python, Node, Go and R, built on the Verdifax SDK in each language. Verdifax's production readiness check runs all four against the live service.

## What this shows, and what it does not

- It is a demo, not customer proof. No customer is named.
- A sealed record shows the checks ran and what the model returned. It does not show the model's judgment was right.
- The 15 rules in [rules/v3.json](rules/v3.json) are general writing rules. They are not legal advice and not a compliance standard. Version 2 reworded rules 04, 05, 06 and 12, and version 3 rewords rules 01 and 05, so each covers only what it was written for; [rules/v1.json](rules/v1.json) and [rules/v2.json](rules/v2.json) are kept as published for records sealed under them, and each record names the hash of the rule set it used.
- The account policy judges what each step declared. It cannot show the declared facts are true.
- A sign-off shows a registered approver key signed that decision on that exact result. It does not show who holds the key or that the review was careful, and Verdifax did not review the result.

## What happens in one review

1. **Intake.** One sealed step records the draft's SHA-256, its size and the rule set's hash. A draft marked `CONFIDENTIAL` or above is refused here by the account policy, so it never reaches an outside model; the refusal is sealed as a receipt. A stop at any later step stops the job the same way, whether the account policy or the server's content screen refuses the step or the server's review of the model's answer halts it: the stop is sealed as a receipt of its kind, the steps sealed before it are kept beside it, and the job never gets a verdict.
2. **Fifteen rule steps**, one per rule, each a branch after intake. The model is sent the draft and one rule and answers in a fixed form, `{verdict, reason}`. Each step seals the rule, the answer, the model that answered (as the API's response named it), the prompt's hash and the model's exact answer, which Verdifax's own AI review also reads. The review lists the statements the answer makes and checks them against each other for a contradiction; each step's record seals what it found, the SHA-256 of the text it checked, and what the account policy chose a contradiction does. The demo's policy, [policy/claims-review-v2.json](policy/claims-review-v2.json), writes that choice out on its rule-check rule as `"ai_review": {"contradiction": "halt"}`, the default: a contradiction stops the job. Change `halt` to `record` and a contradiction is sealed while the job continues to its verdict and a person's sign-off.
3. **Verdict.** A final step joins all fifteen branches: pass only when every rule passed. No model call; the verdict is computed from the answers. A missing, extra or reordered step breaks the job.
4. **Sign-off.** A person holding the approver key signs approve or reject on the verdict, or on the receipt of a refused or halted step. A decision means the same on every result: approve agrees with the result signed, the verdict or the stop the receipt records, and reject judges it wrong. A sign-off releases or blocks nothing on its own.
5. **Publish.** The account's owner key makes the job's records public on random links, or withdraws them. Each decision is sealed by Verdifax as a run publication with its place in the feed, so an auditor sees who made what public and when.
6. **Feed.** The refusals, the sign-offs, the publications and the key changes appear in the account's signal feed, as a GRC system would receive them.

Every record and its place in the job is anchored on the public log, and `verdifax-verify` checks it all offline; it reads every part of these records since v0.27.0. An earlier release refuses the AI review's record in each step and the version 2 publications in the feed, which name who decided each one.

## Run it on your own drafts

You need a Verdifax account's owner key (ask for the free pilot account at [verdifax.com](https://verdifax.com)) and your own Anthropic API key. The model defaults to `claude-fable-5-1`; `--model` takes another id, and each step seals the model that actually answered. Model use is yours and is printed after each review at the published prices in [rules/model-prices-v1.json](rules/model-prices-v1.json).

Pick a language and work in an empty folder; the program writes its keys and records there. Each language installs the exact versions this repository pins. Every program reads the owner key from `VERDIFAX_OWNER_KEY` at setup and the model key from `ANTHROPIC_API_KEY` at review; the R program also needs the `sodium` package, which signs its requests.

```bash
# Python
python3 -m venv .venv && .venv/bin/pip install -r /path/to/verdifax-claims-review-demo/python/requirements.txt
export VERDIFAX_OWNER_KEY=...   # the account's owner key
.venv/bin/python /path/to/verdifax-claims-review-demo/python/claims_review.py setup
export ANTHROPIC_API_KEY=...
.venv/bin/python /path/to/verdifax-claims-review-demo/python/claims_review.py review draft.md
```

```bash
# Node 22 or later
cd /path/to/verdifax-claims-review-demo/node && npm ci && cd -
node /path/to/verdifax-claims-review-demo/node/claims-review.mjs setup
node /path/to/verdifax-claims-review-demo/node/claims-review.mjs review draft.md
```

```bash
# Go 1.27.1 or later
(cd /path/to/verdifax-claims-review-demo/go && go build -o claims-review .)
CLAIMS_REVIEW_REPO=/path/to/verdifax-claims-review-demo /path/to/verdifax-claims-review-demo/go/claims-review setup
CLAIMS_REVIEW_REPO=/path/to/verdifax-claims-review-demo /path/to/verdifax-claims-review-demo/go/claims-review review draft.md
```

```bash
# R 4.3 or later
Rscript -e 'install.packages(c("remotes", "sodium")); remotes::install_github("Verdifax/verdifax-sdk-r@v1.9.0")'
Rscript /path/to/verdifax-claims-review-demo/r/claims-review.R setup
Rscript /path/to/verdifax-claims-review-demo/r/claims-review.R review draft.md
```

The commands and the files are the same in every language: a folder set up by one program is used by any other, so `setup` runs once in a folder however many of the programs you use there.

| Command | Does |
|---|---|
| `setup` | Registers the owner's request key, creates the agent's key and a feed key, registers an approver key generated on this machine, and activates the account policy in [policy/claims-review-v2.json](policy/claims-review-v2.json) ([v1](policy/claims-review-v1.json) is kept for records sealed under it). Runs once in a folder. Running it in a second folder for the same account adds another agent key, feed key and approver key. |
| `review DRAFT.md [--model ID]` | Runs one job on the draft and writes its sealed records to `records/<job_id>/steps/` (the folder `verdifax-verify --job` reads), with the exact payload bytes beside them, so a reader recomputes every sealed payload hash. Prints the verdict, the measured model use and the verify commands. The model's exact answers go to `records/<job_id>/answers/`, one file per rule, so a reader recomputes each sealed `ai_output_hash`, and the text Verdifax's AI review checked goes to `records/<job_id>/traces/`, as the server returns it once, so a reader re-runs the review's check. A refused or halted step, at intake or later, stops the job: its receipt, the records sealed before it, every payload, every answer and every review trace go to `records/stopped-<run_id>/`, and `job.json` there names the step, the kind of stop and the model use spent before it. Any other error from the server is printed as the server gave it, and nothing is recorded as a stop. |
| `sign-off JOB_ID approve\|reject [--note TEXT]` | Signs the decision on the job's verdict with the approver key. Only the note's hash is sent. |
| `sign-off --receipt RUN_ID approve\|reject [--note TEXT]` | Signs a decision on the receipt of a refused or halted step: approve agrees that the step should have stopped, reject judges the stop wrong. Only the note's hash is sent. |
| `publish JOB_ID` | Makes every record of the job public on a random link with the owner key, and writes `records/<job_id>/published.json` with each link and the sealed publication. `publish --receipt RUN_ID` makes a stopped step's receipt public with every step the job sealed before it, and writes `records/stopped-<run_id>/published.json`. A repeated publish keeps the sealed publication saved the first time. |
| `unpublish JOB_ID` | Withdraws the job's records from their links; a later `publish` restores the same links. `--receipt RUN_ID` for a receipt. |
| `feed` | Reads the account's whole signal feed with the feed key into `signals.json`. |
| `verify-commands JOB_ID` | Prints the `verdifax-verify` commands for the job, the policy, the sign-offs and the feed. |

`demo-keys.json` and `approver-key.json` hold secrets; the program writes them with owner-only permissions. Keep them out of version control.

## Check the results offline

Install `verdifax-verify` from [github.com/Verdifax/verdifax-verify](https://github.com/Verdifax/verdifax-verify); it reads every part of these records since v0.27.0. Then, in the folder the program used:

```bash
verdifax-verify --job records/<job_id>/steps --aivp-traces records/<job_id>/traces
verdifax-verify --account-policy /path/to/verdifax-claims-review-demo/policy/claims-review-v2.json records/<job_id>/steps/verdict.json
verdifax-verify --signoff records/<job_id>/signoffs.json --approver-key approver.txt records/<job_id>/steps/verdict.json
verdifax-verify --signoff-history records/<job_id>/signoff-history.json
verdifax-verify --signals signals.json
```

## The sample drafts

[samples/](samples/) holds six drafts written by Verdifax for this demo. Every company, person and product named in them is invented, and no quotation is real: two written to pass, three written to fail one rule each (rule 2, rule 3 and rule 11), and one marked confidential, which the policy refuses at intake. Verdifax's own sealed runs of these samples, and their verify commands, are at [verdifax.com/demo](https://verdifax.com/demo).

## Tests

`testdata/answers.json` holds recorded model answers for tests against a local Verdifax server, and `answers-refused.json` and `answers-halted.json` hold two variants whose rule 3 answer the server's content screen refuses or its review of the model's answer halts; the program refuses recorded answers against production. `demo-parity.json` lists every command, flag, file and sealed field, and Verdifax's checks hold the four programs equal to it.

## License

MIT, as the Verdifax SDKs.
