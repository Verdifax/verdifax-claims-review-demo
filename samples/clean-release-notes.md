# Quillstone Ledger 4.2 release notes

Quillstone Ledger 4.2 is available today for existing customers on the hosted plan.

What changed:

- Exports now include a column for the approving user's role. The column is empty for records created before 4.2, because the role was not stored then.
- Report generation is faster on large accounts. In our own test on November 3, 2026, a 50,000-row export took 41 seconds on 4.1 and 12 seconds on 4.2. Your results depend on your data.
- The audit view shows who changed a record and when. It does not show why, because the product does not collect a reason.

Planned, not yet available: a scheduled export that runs nightly. We expect to release it in the first quarter of 2027, and we will announce it in these notes when it ships.

Quillstone Ledger keeps a history of changes to records inside Quillstone. It does not watch changes made in other systems, and it is not an accounting system of record.

Questions go to support through the usual address in your account settings.
