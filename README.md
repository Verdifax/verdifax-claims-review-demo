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

1. **Intake.** One sealed step records the draft's SHA-256, its size and the rule set's hash. No model is called at intake. Every step, intake included, names the agent that sent it, which the demo's version 4 policy requires: a request that leaves the name out is refused before it runs, with a sealed receipt for the missing field (`review --without-actor` sends intake that way on purpose). Every step also declares two data checks the program computes before it: `draft_unchanged`, the draft re-read and re-hashed against the hash sealed at intake, and `rule_set_pinned`, the rule set's hash against the pinned one; the policy admits a step after the gate only when both pass. Every step that called a model declares the model's own token counts for that call, which the server's per-run and per-account budgets count; intake calls no model and declares none. A draft marked `CONFIDENTIAL` or above is refused here by the account policy, so it never reaches an outside model; the refusal is sealed as a receipt. A stop at any later step stops the job the same way, whether the account policy or the server's content screen refuses the step or the server's review of the model's answer halts it: the stop is sealed as a receipt of its kind, the steps sealed before it are kept beside it, and the job never gets a verdict.
2. **Claims.** The gate passed, so the draft may reach a model. The model lists the factual claims the draft makes, without judging them. The step seals the list, the model that answered, the prompt's hash and the SHA-256 of the model's exact answer (the program keeps the answer's text in the job's `answers/` folder), and Verdifax's own AI review reads that answer: it lists the statements the answer makes and checks them against each other for a contradiction, and the record seals what it found, the SHA-256 of the text it checked, and what the account policy chose a contradiction does.
3. **Fifteen rule steps**, one per rule, each a branch after the claims step. The model is sent the draft and one rule and answers in a fixed form, `{verdict, reason}`. Each step seals the rule, the parsed verdict and reason, the model that answered (as the API's response named it), the prompt's hash and the SHA-256 of the model's exact answer, whose text the program keeps in `answers/` and Verdifax's AI review checks the same way. The demo's policy, [policy/claims-review-v4.json](policy/claims-review-v4.json), writes the review's setting out on its rules for every step after the gate as `"ai_review": {"contradiction": "halt", "required": true}`: a contradiction stops the job, and the server refuses any of these steps that carries no AI answer, so none can finish without a sealed review. Change `halt` to `record` and a contradiction is sealed while the job continues to its verdict and a person's sign-off.
4. **Verdict.** A final step joins all fifteen branches. The program decides it from the rule answers: pass only when every rule passed, so anyone can decide it again from the sealed records. The model then explains that verdict in plain words for an executive, from the verdict and the rule results only (it is not sent the draft), and its explanation is reviewed and sealed like every other answer. A missing, extra or reordered step breaks the job.
5. **Sign-off.** A person holding the approver key signs approve or reject on the verdict, or on the receipt of a refused or halted step. The version 4 policy requires a sign-off on each verdict: the record seals the requirement, and `verdifax-verify` reads a verdict given with no sign-off as `ANCHORED, SIGN-OFF NOT SHOWN`, exit `4`. A decision means the same on every result: approve agrees with the result signed, the verdict or the stop the receipt records, and reject judges it wrong. A sign-off releases or blocks nothing on its own.
6. **Publish.** The account's owner key makes the job's records public on random links, or withdraws them. Each decision is sealed by Verdifax as a run publication with its place in the feed, so an auditor sees who made what public and when.
7. **Feed.** The refusals, the halts, the sign-offs, the publications, the deletions and every change to the account's keys, approvers and settings appear in the account's signal feed, as a GRC system would receive them; a webhook can receive each one as it happens.

Every record and its place in the job is anchored on the public log, and `verdifax-verify` checks it all offline. Records sealed under the demo's version 4 policy, with the agent name, the data checks and the verdict's sign-off requirement, are read by the release the commands below name. Records sealed under the version 3 policy are read in full since v0.28.0, which hashes its review requirement and runs it again, and records sealed under the version 2 policy, before the claims step, verify since v0.27.0. The verify commands below use what the verifier has done since v0.34.0: it limits the age of a saved history's or feed's head statement, which the commands set explicitly, and reads a deleted step's tombstone in the step's place.

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
# Go 1.27.4 or later
(cd /path/to/verdifax-claims-review-demo/go && go build -o claims-review .)
CLAIMS_REVIEW_REPO=/path/to/verdifax-claims-review-demo /path/to/verdifax-claims-review-demo/go/claims-review setup
CLAIMS_REVIEW_REPO=/path/to/verdifax-claims-review-demo /path/to/verdifax-claims-review-demo/go/claims-review review draft.md
```

```bash
# R 4.3 or later
Rscript -e 'install.packages(c("remotes", "sodium")); remotes::install_github("Verdifax/verdifax-sdk-r@v1.16.0")'
Rscript /path/to/verdifax-claims-review-demo/r/claims-review.R setup
Rscript /path/to/verdifax-claims-review-demo/r/claims-review.R review draft.md
```

The commands and the files are the same in every language: a folder set up by one program is used by any other, so `setup` runs once in a folder however many of the programs you use there.

| Command | Does |
|---|---|
| `setup` | Registers the owner's request key, creates the agent's key and a feed key, registers an approver key generated on this machine, and activates the account policy in [policy/claims-review-v4.json](policy/claims-review-v4.json) ([v1](policy/claims-review-v1.json), [v2](policy/claims-review-v2.json) and [v3](policy/claims-review-v3.json) are kept for records sealed under them). Runs once in a folder. Running it in a second folder for the same account adds another agent key, feed key and approver key. |
| `review DRAFT.md [--model ID] [--without-actor]` | Runs one job on the draft and writes its sealed records to `records/<job_id>/steps/` (the folder `verdifax-verify --job` reads), with the exact payload bytes beside them, so a reader recomputes every sealed payload hash. Prints the verdict, the measured model use and the verify commands. The model's exact answers go to `records/<job_id>/answers/`, one file per rule, so a reader recomputes each sealed `ai_output_hash`, and the text Verdifax's AI review checked goes to `records/<job_id>/traces/`, as the server returns it once, so a reader re-runs the review's check. A refused or halted step, at intake or later, stops the job: its receipt, the records sealed before it, every payload, every answer and every review trace go to `records/stopped-<run_id>/`, and `job.json` there names the step, the kind of stop and the model use spent before it. Any other error from the server is printed as the server gave it, and nothing is recorded as a stop. A step whose record was deleted from the account is written as its tombstone, `steps/<step>.deleted.json`, the `run.deleted` entry from the account's feed, which `verdifax-verify --job` reads in the step's place: the job then reads complete except the deleted steps, never complete. |
| `sign-off JOB_ID approve\|reject [--note TEXT] [--approver NAME]` | Signs the decision on the job's verdict with the approver key, or with an approver added by `approver-add`. Only the note's hash is sent. |
| `sign-off --receipt RUN_ID approve\|reject [--note TEXT]` | Signs a decision on the receipt of a refused or halted step: approve agrees that the step should have stopped, reject judges the stop wrong. Only the note's hash is sent. |
| `publish JOB_ID` | Makes every record of the job public on a random link with the owner key, and writes `records/<job_id>/published.json` with each link and the sealed publication. `publish --receipt RUN_ID` makes a stopped step's receipt public with every step the job sealed before it, and writes `records/stopped-<run_id>/published.json`. A repeated publish keeps the sealed publication saved the first time. |
| `unpublish JOB_ID` | Withdraws the job's records from their links; a later `publish` restores the same links. `--receipt RUN_ID` for a receipt. |
| `approver-add NAME` | Registers a second approver key, generated on this machine, as the owner: `approver-NAME-key.json` (owner-only permissions, with the id the server gave it) and `approver-NAME.txt`, and the account's approvers in `approvers.json`. NAME is 1 to 40 lowercase letters, digits or dashes. |
| `approver-revoke NAME` | Revokes that approver as the owner. The sign-offs it gave stay as they were; the server refuses any later one. |
| `settings SETTING VALUE` | The owner's blocks on Verdifax: `owner-replacement` or `verdifax-deletion`, `allowed` or `blocked`, signed with the owner's request key. |
| `request-key-rotate` | Replaces the agent key's request key, approved by the key it replaces, and keeps the new one in `demo-keys.json`. |
| `request-key-clear` | Clears the agent key's request key, approved by that key; the agent then sends unsigned requests. |
| `key-revoke` | Revokes the agent key as the owner, when its work is done, and writes the account's keys to `account-keys.json`. The records it sealed stay as they were. |
| `delete-step JOB_ID STEP` | Deletes one step's record as the owner, signed over the record's hash, and replaces `records/<job_id>/steps/STEP.json` with `STEP.deleted.json`, its tombstone from the feed; `verdifax-verify --job` then reads the job as complete except the deleted step. A public record must be withdrawn first. |
| `webhook-subscribe URL` | Subscribes an `https://` address to every kind of event as the owner, and writes the subscription and its secret, shown once, to `webhook.json` with owner-only permissions. |
| `webhook-list` | Writes the account's webhooks to `webhooks.json` and each one's delivery log to `webhook-deliveries.json`. |
| `feed` | Reads the account's whole signal feed with the feed key into `signals.json`. Every change the commands above make is in it, each webhook added or removed included from verifier v0.35.0's server release on. |
| `verify-commands JOB_ID` | Prints the `verdifax-verify` commands for the job, the policy, the sign-offs and the feed. |

`demo-keys.json`, `approver-key.json`, each `approver-NAME-key.json` and `webhook.json` hold secrets; the program writes them with owner-only permissions. Keep them out of version control.

## Check the results offline

Install `verdifax-verify` from [github.com/Verdifax/verdifax-verify](https://github.com/Verdifax/verdifax-verify). It reads every part of these records since v0.28.0. The commands below use the flags it has had since v0.34.0, which limit the age of a saved history's or feed's head statement (24 hours unless the flag says otherwise), and since that release it reads a deleted step's tombstone. Then, in the folder the program used:

```bash
verdifax-verify --job records/<job_id>/steps --aivp-traces records/<job_id>/traces
verdifax-verify --account-policy /path/to/verdifax-claims-review-demo/policy/claims-review-v4.json records/<job_id>/steps/verdict.json
verdifax-verify --signoff records/<job_id>/signoffs.json --approver-key approver.txt records/<job_id>/steps/verdict.json
# 168h: a saved history is checked as of when it was fetched; raise the limit for an older copy, or run sign-off again to fetch it fresh
verdifax-verify --signoff-history records/<job_id>/signoff-history.json --signoff-head-max-age 168h
# 168h: a saved feed is checked as of when it was fetched; raise the limit for an older copy, or run feed again to fetch it fresh
verdifax-verify --signals signals.json --feed-head-max-age 168h
```

## The sample drafts

[samples/](samples/) holds six drafts written by Verdifax for this demo. Every company, person and product named in them is invented, and no quotation is real: two written to pass, three written to fail one rule each (rule 2, rule 3 and rule 11), and one marked confidential, which the policy refuses at intake. Verdifax's own sealed runs of these samples, and their verify commands, are at [verdifax.com/demo](https://verdifax.com/demo).

## Tests

`testdata/answers.json` holds recorded model answers for tests against a local Verdifax server, and `answers-refused.json` and `answers-halted.json` hold two variants whose rule 3 answer the server's content screen refuses or its review of the model's answer halts; the program refuses recorded answers against production. `deleted-step-cases.json` holds feed events every language reads to find a job's deleted steps and write each one's tombstone as `<step>.deleted.json`. `demo-parity.json` lists every command, flag, file and sealed field, and Verdifax's checks hold the four programs equal to its commands, flags, environment variables, constants and sealed fields (the files each command writes are listed for readers and are not compared).

## License

MIT, as the Verdifax SDKs.
