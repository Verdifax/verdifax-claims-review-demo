#!/usr/bin/env python3
"""Counts the two budget drafts' tokens with the model's own token counter (a
free call) and exits 1 unless each would trigger the stop it is written for:

  samples/budget-over-run-limit.md    its first model step is over the per-run
                                      limit of 100,000 tokens (ccv_halt)
  samples/budget-over-daily-limit.md  every step stays under the per-run limit,
                                      with room for its answer, while the steps
                                      of one review together pass the account's
                                      daily limit of 1,000,000 (macc_halt)

The demo driver (scripts/demo-samples.sh) runs it before sending either draft,
and it is run before the paid run is handed over. Round 62: drafts sized by an
estimate came out about 35% larger in the model's own count, so only this
count decides. Needs ANTHROPIC_API_KEY; sends no draft to the model.

Usage: python3 tools/check-budget-drafts.py [DEMO_FOLDER]
"""
import json
import sys
from pathlib import Path

import anthropic

MODEL = "claude-fable-5-1"
RUN_LIMIT, DAY_LIMIT, ANSWER_ROOM = 100_000, 1_000_000, 4_000


def main() -> int:
    demo = Path(sys.argv[1] if len(sys.argv) > 1 else Path(__file__).resolve().parent.parent)
    client = anthropic.Anthropic()

    def count(system: str, user: str) -> int:
        return client.messages.count_tokens(model=MODEL, system=system, messages=[{"role": "user", "content": user}]).input_tokens

    claims, rule_prompt = (demo / "rules" / "claims-prompt-v1.txt").read_text(), (demo / "rules" / "prompt-v1.txt").read_text()
    rules = json.loads((demo / "rules" / "v3.json").read_text())["rules"]
    over = (demo / "samples" / "budget-over-run-limit.md").read_bytes().decode("utf-8")
    day = (demo / "samples" / "budget-over-daily-limit.md").read_bytes().decode("utf-8")
    first = count(claims, "Draft:\n" + over)
    steps = [count(claims, "Draft:\n" + day)] + [count(rule_prompt, "Rule %s: %s\n\nDraft:\n%s" % (r["id"], r["text"], day)) for r in rules]
    running, halt_at = 0, None
    for i, n in enumerate(steps):
        running += n
        if running > DAY_LIMIT and halt_at is None:
            halt_at = i
    names = ["claims"] + [r["id"] for r in rules]
    ok = first > RUN_LIMIT and max(steps) + ANSWER_ROOM < RUN_LIMIT and halt_at is not None and halt_at > 0
    print("over-run draft, first step: %d input tokens (needs more than %d)" % (first, RUN_LIMIT))
    print("over-day draft, each step: %d to %d input tokens (needs each under %d), together %d (needs more than %d)"
          % (min(steps), max(steps), RUN_LIMIT - ANSWER_ROOM, sum(steps), DAY_LIMIT))
    if halt_at is not None:
        print("over-day draft: the daily limit is passed at step %s, number %d of %d" % (names[halt_at], halt_at + 1, len(steps)))
    return 0 if ok else 1


if __name__ == "__main__":
    sys.exit(main())
