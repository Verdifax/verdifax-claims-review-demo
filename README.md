# Verdifax claims review demo

A small agent reads a draft (a web page, a press release, a sales email) and checks it against 15 general writing rules, such as "don't call anything unbreakable" and "don't promise a legal result". For each rule it returns pass or fail with a reason. A person decides. Verdifax seals every step, so anyone can check afterwards, offline, which rules were checked, what the model answered, which policy admitted each step, and who signed off.

The same program is here in Python, Node, Go and R, built on the Verdifax SDK in each language. Verdifax's production readiness check runs all four against the live service.

## What this shows, and what it does not

- It is a demo, not customer proof. No customer is named.
- A sealed record shows the checks ran and what the model returned. It does not show the model's judgment was right.
- The 15 rules in [rules/v1.json](rules/v1.json) are general writing rules. They are not legal advice and not a compliance standard.
- The account policy judges what each step declared. It cannot show the declared facts are true.
- A sign-off shows a registered approver key signed that decision on that exact result. It does not show who holds the key or that the review was careful, and Verdifax did not review the result.

## What happens in one review

1. **Intake.** One sealed step records the draft's SHA-256, its size and the rule set's hash. A draft marked `CONFIDENTIAL` or above is refused here by the account policy, so it never reaches an outside model; the refusal is sealed as a receipt. A refusal at any later step, by the account policy or by the server's own content screen, stops the job the same way: the refusal is sealed as a receipt, the steps sealed before it are kept beside it, and the job never gets a verdict.
2. **Fifteen rule steps**, one per rule, each a branch after intake. The model is sent the draft and one rule and answers in a fixed form, `{verdict, reason}`. Each step seals the rule, the answer, the model that answered (as the API's response named it), the prompt's hash and the model's exact answer, which Verdifax's own AI review also reads.
3. **Verdict.** A final step joins all fifteen branches: pass only when every rule passed. No model call; the verdict is computed from the answers. A missing, extra or reordered step breaks the job.
4. **Sign-off.** A person holding the approver key signs approve or reject on the verdict, or rejects a refusal's receipt.
5. **Feed.** The refusals, the sign-offs and the key changes appear in the account's signal feed, as a GRC system would receive them.

Every record and its place in the job is anchored on the public log, and `verdifax-verify` checks it all offline; the verifier reads these records since v0.23.0, and an earlier release refuses them.

## Run it on your own drafts

You need a Verdifax account's owner key (ask for the free pilot account at verdifax.com) and your own Anthropic API key. The model defaults to `claude-fable-5-1`; `--model` takes another id, and each step seals the model that actually answered. Model use is yours and is printed after each review at the published prices in [rules/model-prices-v1.json](rules/model-prices-v1.json).

Pick a language and work in an empty folder; the program writes its keys and records there.

```bash
# Python
python3 -m venv .venv && .venv/bin/pip install "verdifax[signing]" anthropic pydantic
export VERDIFAX_OWNER_KEY=...   # the account's owner key
.venv/bin/python /path/to/verdifax-claims-review-demo/python/claims_review.py setup
export ANTHROPIC_API_KEY=...
.venv/bin/python /path/to/verdifax-claims-review-demo/python/claims_review.py review draft.md
```

```bash
# Node 22 or later
cd /path/to/verdifax-claims-review-demo/node && npm install && cd -
node /path/to/verdifax-claims-review-demo/node/claims-review.mjs setup
node /path/to/verdifax-claims-review-demo/node/claims-review.mjs review draft.md
```

```bash
# Go 1.23 or later
(cd /path/to/verdifax-claims-review-demo/go && go build -o claims-review .)
CLAIMS_REVIEW_REPO=/path/to/verdifax-claims-review-demo /path/to/verdifax-claims-review-demo/go/claims-review setup
CLAIMS_REVIEW_REPO=/path/to/verdifax-claims-review-demo /path/to/verdifax-claims-review-demo/go/claims-review review draft.md
```

```bash
# R 4.3 or later, with the verdifax package installed from github.com/Verdifax/verdifax-sdk-r
Rscript /path/to/verdifax-claims-review-demo/r/claims-review.R setup
Rscript /path/to/verdifax-claims-review-demo/r/claims-review.R review draft.md
```

The commands are the same in every language, and so are the files: a folder set up by one program is used by any other, so `setup` runs once per account however many of the programs you use.

The commands are the same in every language:

| Command | Does |
|---|---|
| `setup` | Registers the owner's request key, creates the agent's key and a feed key, registers an approver key generated on this machine, and activates the account policy in [policy/claims-review-v1.json](policy/claims-review-v1.json). Runs once per account and folder. |
| `review DRAFT.md [--model ID]` | Runs one job on the draft and writes its records to `records/<job_id>/`, with the exact payload bytes beside them, so a reader recomputes every sealed payload hash. Prints the verdict, the measured model use and the verify commands. A refused step, at intake or later, stops the job: its receipt, the records sealed before it and every payload go to `records/refused-<run_id>/`, and `job.json` there names the step. |
| `sign-off JOB_ID approve\|reject [--note TEXT]` | Signs the decision on the job's verdict with the approver key. Only the note's hash is sent. |
| `sign-off --receipt RUN_ID reject` | Signs a decision on a refusal's receipt. |
| `feed` | Reads the account's whole signal feed with the feed key into `signals.json`. |
| `verify-commands JOB_ID` | Prints the `verdifax-verify` commands for the job, the policy, the sign-offs and the feed. |

`demo-keys.json` and `approver-key.json` hold secrets; the program writes them with owner-only permissions. Keep them out of version control.

## Check the results offline

Install `verdifax-verify` from github.com/Verdifax/verdifax-verify (the verifier reads these records since v0.23.0), then:

```bash
verdifax-verify --job records/<job_id>
verdifax-verify --account-policy policy/claims-review-v1.json records/<job_id>/verdict.json
verdifax-verify --signoff records/<job_id>/signoffs.json --approver-key approver.txt records/<job_id>/verdict.json
verdifax-verify --signoff-history records/<job_id>/signoff-history.json
verdifax-verify --signals signals.json
```

## The sample drafts

[samples/](samples/) holds six drafts written by Verdifax for this demo. No real company, person, product or quotation appears in them: two written to pass, three written to fail one rule each (rule 2, rule 3 and rule 11), and one marked confidential, which the policy refuses at intake. Verdifax's own sealed runs of these samples, and their verify commands, are at verdifax.com/demo.

## Tests

`testdata/answers.json` holds recorded model answers for tests against a local Verdifax server; the program refuses them against production. `demo-parity.json` lists every command, flag, file and sealed field, and Verdifax's checks hold the four programs equal to it.

## License

MIT, as the Verdifax SDKs.
