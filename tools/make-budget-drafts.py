#!/usr/bin/env python3
"""Writes the two long drafts of the demo's budget chapter, byte for byte the
same on every run (a fixed seed, no clock).

Both are release notes for the invented product Quillstone Ledger, written by
Verdifax in the plain style of the other samples, long on purpose:

  samples/budget-over-run-limit.md    176,949 tokens in its first model step,
                                      more than the per-run limit of 100,000,
                                      so that step is halted (ccv_halt)
  samples/budget-over-daily-limit.md  84,905 to 85,018 tokens a step, so each
                                      stays under the per-run limit while the
                                      steps of one review together pass the
                                      account's daily limit of 1,000,000 at
                                      step 12, rule-11 (macc_halt)

The counts are Fable 5.1's own, measured on November 14, 2026 with
tools/check-budget-drafts.py, which the demo driver runs again before sending
either draft. An estimate is never used: the first over-day draft, sized by
one, counted about 35% larger and would have been halted by the per-run limit.

Usage: python3 tools/make-budget-drafts.py
"""

import random
from pathlib import Path

HERE = Path(__file__).resolve().parent.parent

AREAS = ["exports", "the audit view", "report generation", "user roles", "the import screen", "account settings",
         "scheduled reports", "the approval queue", "invoice matching", "the search index", "currency conversion",
         "the activity log", "bank feeds", "the tax summary", "attachment storage", "the mobile view"]
CHANGES = [
    "now includes a column for {thing}. The column is empty for records created before {ver}, because {thing} was not stored then.",
    "now shows {thing} beside each entry. Entries made before {ver} show a dash, because the value was not recorded.",
    "now accepts {thing} as a filter. The filter applies to records created on or after {ver}.",
    "faster on accounts with many records. In our own test on {date}, a {rows:,}-row run took {old} seconds on {prev} and {new} seconds on {ver}. Your results depend on your data.",
    "now keeps {thing} for {days} days. After that the value is removed, and the entry shows when it was removed.",
    "can now be turned off for each user. When it is turned off, {thing} is not collected for that user from then on.",
    "now names the user who last changed {thing}. It does not show why the change was made, because the product does not ask for a reason.",
]
THINGS = ["the approving user's role", "the source file name", "the original currency", "the reviewer's comment count",
          "the device type", "the exchange rate used", "the batch number", "the cost centre", "the payment terms",
          "the import job number", "the time zone of the entry", "the attachment count"]
FIXES = [
    "Fixed: in {area}, a month with no records showed a blank page. It now shows an empty table.",
    "Fixed: totals in {area} were rounded before they were added, which could change the last digit. They are now added first and rounded once.",
    "Fixed: a date typed as day/month could be read as month/day in {area}. The format now follows the account's settings.",
    "Fixed: {area} did not keep the sort order after a page reload. It now keeps it until the user changes it.",
]
KNOWN = [
    "Known issue: {area} does not yet support files larger than {mb} MB. Larger files are refused with a message that names the limit.",
    "Known issue: on slow connections {area} can take up to {secs} seconds to load. We are working on it and will say here when it is fixed.",
]
PLANNED = "Planned, not yet available: {thing} in {area}. We expect to release it in {when}, and we will announce it in these notes when it ships."
WHEN = ["the first quarter of 2027", "the second quarter of 2027", "the second half of 2027"]


def notes(seed: int, releases: int) -> str:
    rnd = random.Random(seed)
    out = []
    major, minor = 1, 0
    for i in range(releases):
        prev = "%d.%d" % (major, minor)
        minor += 1
        if minor == 10:
            major, minor = major + 1, 0
        ver = "%d.%d" % (major, minor)
        day = 1 + (i * 7) % 28
        month = 1 + (i // 4) % 12
        date = "%d %s %d" % (day, ["January", "February", "March", "April", "May", "June", "July", "August",
                                    "September", "October", "November", "December"][month - 1], 2019 + i // 48)
        out.append("## Quillstone Ledger %s\n" % ver)
        out.append("Quillstone Ledger %s was released on %s for customers on the hosted plan.\n" % (ver, date))
        out.append("What changed:\n")
        for _ in range(rnd.randint(3, 5)):
            area, thing = rnd.choice(AREAS), rnd.choice(THINGS)
            old = rnd.randint(20, 90)
            line = rnd.choice(CHANGES).format(thing=thing, ver=ver, prev=prev, date=date, rows=rnd.randint(10, 200) * 1000,
                                              old=old, new=max(2, old // rnd.randint(2, 5)), days=rnd.choice([30, 90, 180, 365]))
            out.append("- %s: %s" % (area[0].upper() + area[1:], line))
        out.append("")
        for _ in range(rnd.randint(1, 2)):
            out.append("- " + rnd.choice(FIXES).format(area=rnd.choice(AREAS)))
        out.append("")
        if rnd.random() < 0.5:
            out.append(rnd.choice(KNOWN).format(area=rnd.choice(AREAS), mb=rnd.choice([25, 50, 100]), secs=rnd.choice([10, 20, 30])) + "\n")
        if rnd.random() < 0.3:
            out.append(PLANNED.format(thing=rnd.choice(THINGS), area=rnd.choice(AREAS), when=rnd.choice(WHEN)) + "\n")
    return "\n".join(out)


def write(name: str, purpose: str, seed: int, target_chars: int) -> None:
    head = ("<!-- Written by Verdifax for the claims review demo's budget chapter. It is long on purpose: %s. "
            "Every product, version and figure in it is invented. -->\n\n# Quillstone Ledger release notes archive\n\n"
            "These are the release notes of every Quillstone Ledger release, oldest first.\n\n" % purpose)
    releases = 10
    while len(head + notes(seed, releases)) < target_chars:
        releases += 10
    (HERE / "samples" / name).write_bytes((head + notes(seed, releases)).encode("utf-8"))
    print("%s: %d releases, %d characters" % (name, releases, len(head + notes(seed, releases))))


if __name__ == "__main__":
    write("budget-over-run-limit.md", "its first model step declares more tokens than the per-run limit of 100,000", 62001, 546_000)
    write("budget-over-daily-limit.md", "each step stays under the per-run limit while one review's steps together pass the daily limit of 1,000,000", 62002, 258_000)
