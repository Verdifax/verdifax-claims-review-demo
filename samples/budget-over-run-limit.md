<!-- Written by Verdifax for the claims review demo's budget chapter. It is long on purpose: its first model step declares more tokens than the per-run limit of 100,000. Every product, version and figure in it is invented. -->

# Quillstone Ledger release notes archive

These are the release notes of every Quillstone Ledger release, oldest first.

## Quillstone Ledger 1.1

Quillstone Ledger 1.1 was released on 1 January 2019 for customers on the hosted plan.

What changed:

- The activity log: now includes a column for the reviewer's comment count. The column is empty for records created before 1.1, because the reviewer's comment count was not stored then.
- The import screen: now accepts the batch number as a filter. The filter applies to records created on or after 1.1.
- The audit view: now includes a column for the batch number. The column is empty for records created before 1.1, because the batch number was not stored then.
- The search index: now includes a column for the time zone of the entry. The column is empty for records created before 1.1, because the time zone of the entry was not stored then.
- The import screen: now shows the reviewer's comment count beside each entry. Entries made before 1.1 show a dash, because the value was not recorded.

- Fixed: the activity log did not keep the sort order after a page reload. It now keeps it until the user changes it.
- Fixed: totals in the tax summary were rounded before they were added, which could change the last digit. They are now added first and rounded once.

## Quillstone Ledger 1.2

Quillstone Ledger 1.2 was released on 8 January 2019 for customers on the hosted plan.

What changed:

- The search index: can now be turned off for each user. When it is turned off, the cost centre is not collected for that user from then on.
- The search index: now includes a column for the time zone of the entry. The column is empty for records created before 1.2, because the time zone of the entry was not stored then.
- Report generation: now shows the original currency beside each entry. Entries made before 1.2 show a dash, because the value was not recorded.
- Report generation: now shows the attachment count beside each entry. Entries made before 1.2 show a dash, because the value was not recorded.
- Invoice matching: now keeps the import job number for 180 days. After that the value is removed, and the entry shows when it was removed.

- Fixed: totals in invoice matching were rounded before they were added, which could change the last digit. They are now added first and rounded once.

## Quillstone Ledger 1.3

Quillstone Ledger 1.3 was released on 15 January 2019 for customers on the hosted plan.

What changed:

- User roles: now accepts the approving user's role as a filter. The filter applies to records created on or after 1.3.
- Currency conversion: can now be turned off for each user. When it is turned off, the time zone of the entry is not collected for that user from then on.
- Attachment storage: now keeps the source file name for 180 days. After that the value is removed, and the entry shows when it was removed.
- Account settings: now includes a column for the time zone of the entry. The column is empty for records created before 1.3, because the time zone of the entry was not stored then.

- Fixed: a date typed as day/month could be read as month/day in the audit view. The format now follows the account's settings.
- Fixed: a date typed as day/month could be read as month/day in the search index. The format now follows the account's settings.

Known issue: on slow connections the approval queue can take up to 30 seconds to load. We are working on it and will say here when it is fixed.

Planned, not yet available: the source file name in the mobile view. We expect to release it in the first quarter of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 1.4

Quillstone Ledger 1.4 was released on 22 January 2019 for customers on the hosted plan.

What changed:

- The approval queue: now keeps the cost centre for 180 days. After that the value is removed, and the entry shows when it was removed.
- The search index: now keeps the exchange rate used for 90 days. After that the value is removed, and the entry shows when it was removed.
- Scheduled reports: now accepts the batch number as a filter. The filter applies to records created on or after 1.4.
- The mobile view: now includes a column for the source file name. The column is empty for records created before 1.4, because the source file name was not stored then.

- Fixed: bank feeds did not keep the sort order after a page reload. It now keeps it until the user changes it.
- Fixed: in account settings, a month with no records showed a blank page. It now shows an empty table.

## Quillstone Ledger 1.5

Quillstone Ledger 1.5 was released on 1 February 2019 for customers on the hosted plan.

What changed:

- The audit view: now accepts the approving user's role as a filter. The filter applies to records created on or after 1.5.
- The mobile view: now shows the source file name beside each entry. Entries made before 1.5 show a dash, because the value was not recorded.
- Currency conversion: now keeps the original currency for 90 days. After that the value is removed, and the entry shows when it was removed.
- User roles: now names the user who last changed the time zone of the entry. It does not show why the change was made, because the product does not ask for a reason.

- Fixed: in user roles, a month with no records showed a blank page. It now shows an empty table.

Known issue: on slow connections the activity log can take up to 20 seconds to load. We are working on it and will say here when it is fixed.

Planned, not yet available: the cost centre in the search index. We expect to release it in the second quarter of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 1.6

Quillstone Ledger 1.6 was released on 8 February 2019 for customers on the hosted plan.

What changed:

- The approval queue: now keeps the attachment count for 365 days. After that the value is removed, and the entry shows when it was removed.
- Invoice matching: now keeps the attachment count for 365 days. After that the value is removed, and the entry shows when it was removed.
- The approval queue: now keeps the reviewer's comment count for 90 days. After that the value is removed, and the entry shows when it was removed.

- Fixed: totals in user roles were rounded before they were added, which could change the last digit. They are now added first and rounded once.

Known issue: the audit view does not yet support files larger than 50 MB. Larger files are refused with a message that names the limit.

Planned, not yet available: the original currency in the tax summary. We expect to release it in the first quarter of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 1.7

Quillstone Ledger 1.7 was released on 15 February 2019 for customers on the hosted plan.

What changed:

- Report generation: can now be turned off for each user. When it is turned off, the import job number is not collected for that user from then on.
- User roles: now shows the device type beside each entry. Entries made before 1.7 show a dash, because the value was not recorded.
- The mobile view: now names the user who last changed the approving user's role. It does not show why the change was made, because the product does not ask for a reason.
- Account settings: now names the user who last changed the device type. It does not show why the change was made, because the product does not ask for a reason.

- Fixed: a date typed as day/month could be read as month/day in attachment storage. The format now follows the account's settings.

Known issue: the mobile view does not yet support files larger than 50 MB. Larger files are refused with a message that names the limit.

## Quillstone Ledger 1.8

Quillstone Ledger 1.8 was released on 22 February 2019 for customers on the hosted plan.

What changed:

- Account settings: now accepts the import job number as a filter. The filter applies to records created on or after 1.8.
- The tax summary: now names the user who last changed the approving user's role. It does not show why the change was made, because the product does not ask for a reason.
- The audit view: now shows the reviewer's comment count beside each entry. Entries made before 1.8 show a dash, because the value was not recorded.
- Invoice matching: now includes a column for the approving user's role. The column is empty for records created before 1.8, because the approving user's role was not stored then.
- The approval queue: can now be turned off for each user. When it is turned off, the time zone of the entry is not collected for that user from then on.

- Fixed: in the audit view, a month with no records showed a blank page. It now shows an empty table.

Planned, not yet available: the device type in scheduled reports. We expect to release it in the second half of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 1.9

Quillstone Ledger 1.9 was released on 1 March 2019 for customers on the hosted plan.

What changed:

- Attachment storage: now keeps the cost centre for 180 days. After that the value is removed, and the entry shows when it was removed.
- The import screen: faster on accounts with many records. In our own test on 1 March 2019, a 160,000-row run took 55 seconds on 1.8 and 11 seconds on 1.9. Your results depend on your data.
- User roles: now accepts the cost centre as a filter. The filter applies to records created on or after 1.9.
- The audit view: can now be turned off for each user. When it is turned off, the payment terms is not collected for that user from then on.
- Currency conversion: now includes a column for the exchange rate used. The column is empty for records created before 1.9, because the exchange rate used was not stored then.

- Fixed: totals in scheduled reports were rounded before they were added, which could change the last digit. They are now added first and rounded once.

Known issue: user roles does not yet support files larger than 25 MB. Larger files are refused with a message that names the limit.

## Quillstone Ledger 2.0

Quillstone Ledger 2.0 was released on 8 March 2019 for customers on the hosted plan.

What changed:

- Report generation: can now be turned off for each user. When it is turned off, the batch number is not collected for that user from then on.
- User roles: now keeps the batch number for 365 days. After that the value is removed, and the entry shows when it was removed.
- Account settings: now names the user who last changed the source file name. It does not show why the change was made, because the product does not ask for a reason.
- The audit view: now accepts the cost centre as a filter. The filter applies to records created on or after 2.0.
- The mobile view: now includes a column for the cost centre. The column is empty for records created before 2.0, because the cost centre was not stored then.

- Fixed: in attachment storage, a month with no records showed a blank page. It now shows an empty table.
- Fixed: totals in the import screen were rounded before they were added, which could change the last digit. They are now added first and rounded once.

Planned, not yet available: the batch number in the mobile view. We expect to release it in the first quarter of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 2.1

Quillstone Ledger 2.1 was released on 15 March 2019 for customers on the hosted plan.

What changed:

- Invoice matching: now shows the source file name beside each entry. Entries made before 2.1 show a dash, because the value was not recorded.
- Attachment storage: can now be turned off for each user. When it is turned off, the cost centre is not collected for that user from then on.
- The search index: now accepts the attachment count as a filter. The filter applies to records created on or after 2.1.

- Fixed: invoice matching did not keep the sort order after a page reload. It now keeps it until the user changes it.
- Fixed: in scheduled reports, a month with no records showed a blank page. It now shows an empty table.

## Quillstone Ledger 2.2

Quillstone Ledger 2.2 was released on 22 March 2019 for customers on the hosted plan.

What changed:

- Invoice matching: can now be turned off for each user. When it is turned off, the device type is not collected for that user from then on.
- User roles: now names the user who last changed the cost centre. It does not show why the change was made, because the product does not ask for a reason.
- Account settings: now names the user who last changed the payment terms. It does not show why the change was made, because the product does not ask for a reason.
- Invoice matching: now shows the batch number beside each entry. Entries made before 2.2 show a dash, because the value was not recorded.

- Fixed: a date typed as day/month could be read as month/day in the import screen. The format now follows the account's settings.
- Fixed: totals in attachment storage were rounded before they were added, which could change the last digit. They are now added first and rounded once.

Known issue: on slow connections report generation can take up to 20 seconds to load. We are working on it and will say here when it is fixed.

## Quillstone Ledger 2.3

Quillstone Ledger 2.3 was released on 1 April 2019 for customers on the hosted plan.

What changed:

- Report generation: now accepts the original currency as a filter. The filter applies to records created on or after 2.3.
- The activity log: now keeps the batch number for 30 days. After that the value is removed, and the entry shows when it was removed.
- The mobile view: now includes a column for the time zone of the entry. The column is empty for records created before 2.3, because the time zone of the entry was not stored then.
- User roles: faster on accounts with many records. In our own test on 1 April 2019, a 28,000-row run took 55 seconds on 2.2 and 13 seconds on 2.3. Your results depend on your data.
- Invoice matching: now accepts the import job number as a filter. The filter applies to records created on or after 2.3.

- Fixed: totals in the approval queue were rounded before they were added, which could change the last digit. They are now added first and rounded once.
- Fixed: in report generation, a month with no records showed a blank page. It now shows an empty table.

Planned, not yet available: the import job number in the search index. We expect to release it in the second half of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 2.4

Quillstone Ledger 2.4 was released on 8 April 2019 for customers on the hosted plan.

What changed:

- Bank feeds: faster on accounts with many records. In our own test on 8 April 2019, a 110,000-row run took 48 seconds on 2.3 and 16 seconds on 2.4. Your results depend on your data.
- The mobile view: now includes a column for the payment terms. The column is empty for records created before 2.4, because the payment terms was not stored then.
- The search index: now shows the reviewer's comment count beside each entry. Entries made before 2.4 show a dash, because the value was not recorded.
- The tax summary: now keeps the time zone of the entry for 180 days. After that the value is removed, and the entry shows when it was removed.

- Fixed: currency conversion did not keep the sort order after a page reload. It now keeps it until the user changes it.
- Fixed: invoice matching did not keep the sort order after a page reload. It now keeps it until the user changes it.

Known issue: on slow connections the activity log can take up to 30 seconds to load. We are working on it and will say here when it is fixed.

## Quillstone Ledger 2.5

Quillstone Ledger 2.5 was released on 15 April 2019 for customers on the hosted plan.

What changed:

- Currency conversion: now includes a column for the approving user's role. The column is empty for records created before 2.5, because the approving user's role was not stored then.
- Invoice matching: now includes a column for the payment terms. The column is empty for records created before 2.5, because the payment terms was not stored then.
- Bank feeds: now includes a column for the import job number. The column is empty for records created before 2.5, because the import job number was not stored then.

- Fixed: attachment storage did not keep the sort order after a page reload. It now keeps it until the user changes it.
- Fixed: in the audit view, a month with no records showed a blank page. It now shows an empty table.

Known issue: the audit view does not yet support files larger than 25 MB. Larger files are refused with a message that names the limit.

## Quillstone Ledger 2.6

Quillstone Ledger 2.6 was released on 22 April 2019 for customers on the hosted plan.

What changed:

- Invoice matching: now accepts the source file name as a filter. The filter applies to records created on or after 2.6.
- The activity log: faster on accounts with many records. In our own test on 22 April 2019, a 133,000-row run took 64 seconds on 2.5 and 32 seconds on 2.6. Your results depend on your data.
- Report generation: now accepts the device type as a filter. The filter applies to records created on or after 2.6.
- Report generation: now keeps the import job number for 30 days. After that the value is removed, and the entry shows when it was removed.
- Report generation: now shows the cost centre beside each entry. Entries made before 2.6 show a dash, because the value was not recorded.

- Fixed: a date typed as day/month could be read as month/day in the approval queue. The format now follows the account's settings.

## Quillstone Ledger 2.7

Quillstone Ledger 2.7 was released on 1 May 2019 for customers on the hosted plan.

What changed:

- The approval queue: can now be turned off for each user. When it is turned off, the attachment count is not collected for that user from then on.
- Scheduled reports: faster on accounts with many records. In our own test on 1 May 2019, a 110,000-row run took 36 seconds on 2.6 and 18 seconds on 2.7. Your results depend on your data.
- The import screen: now includes a column for the import job number. The column is empty for records created before 2.7, because the import job number was not stored then.
- The search index: now names the user who last changed the reviewer's comment count. It does not show why the change was made, because the product does not ask for a reason.

- Fixed: totals in account settings were rounded before they were added, which could change the last digit. They are now added first and rounded once.

## Quillstone Ledger 2.8

Quillstone Ledger 2.8 was released on 8 May 2019 for customers on the hosted plan.

What changed:

- Currency conversion: now includes a column for the device type. The column is empty for records created before 2.8, because the device type was not stored then.
- The activity log: now names the user who last changed the attachment count. It does not show why the change was made, because the product does not ask for a reason.
- The audit view: now keeps the reviewer's comment count for 180 days. After that the value is removed, and the entry shows when it was removed.
- The approval queue: can now be turned off for each user. When it is turned off, the attachment count is not collected for that user from then on.
- The import screen: now includes a column for the exchange rate used. The column is empty for records created before 2.8, because the exchange rate used was not stored then.

- Fixed: a date typed as day/month could be read as month/day in invoice matching. The format now follows the account's settings.
- Fixed: invoice matching did not keep the sort order after a page reload. It now keeps it until the user changes it.

## Quillstone Ledger 2.9

Quillstone Ledger 2.9 was released on 15 May 2019 for customers on the hosted plan.

What changed:

- Report generation: faster on accounts with many records. In our own test on 15 May 2019, a 167,000-row run took 79 seconds on 2.8 and 39 seconds on 2.9. Your results depend on your data.
- The approval queue: now includes a column for the batch number. The column is empty for records created before 2.9, because the batch number was not stored then.
- Account settings: now includes a column for the original currency. The column is empty for records created before 2.9, because the original currency was not stored then.
- Attachment storage: can now be turned off for each user. When it is turned off, the payment terms is not collected for that user from then on.
- User roles: now keeps the batch number for 365 days. After that the value is removed, and the entry shows when it was removed.

- Fixed: totals in scheduled reports were rounded before they were added, which could change the last digit. They are now added first and rounded once.
- Fixed: a date typed as day/month could be read as month/day in the search index. The format now follows the account's settings.

## Quillstone Ledger 3.0

Quillstone Ledger 3.0 was released on 22 May 2019 for customers on the hosted plan.

What changed:

- The activity log: now includes a column for the source file name. The column is empty for records created before 3.0, because the source file name was not stored then.
- Attachment storage: faster on accounts with many records. In our own test on 22 May 2019, a 49,000-row run took 22 seconds on 2.9 and 11 seconds on 3.0. Your results depend on your data.
- Exports: can now be turned off for each user. When it is turned off, the exchange rate used is not collected for that user from then on.
- Report generation: now accepts the cost centre as a filter. The filter applies to records created on or after 3.0.
- Scheduled reports: can now be turned off for each user. When it is turned off, the payment terms is not collected for that user from then on.

- Fixed: invoice matching did not keep the sort order after a page reload. It now keeps it until the user changes it.

Known issue: the approval queue does not yet support files larger than 100 MB. Larger files are refused with a message that names the limit.

## Quillstone Ledger 3.1

Quillstone Ledger 3.1 was released on 1 June 2019 for customers on the hosted plan.

What changed:

- Currency conversion: now accepts the attachment count as a filter. The filter applies to records created on or after 3.1.
- The search index: can now be turned off for each user. When it is turned off, the attachment count is not collected for that user from then on.
- Account settings: now names the user who last changed the exchange rate used. It does not show why the change was made, because the product does not ask for a reason.
- The audit view: now accepts the original currency as a filter. The filter applies to records created on or after 3.1.

- Fixed: a date typed as day/month could be read as month/day in user roles. The format now follows the account's settings.

Planned, not yet available: the approving user's role in the mobile view. We expect to release it in the second quarter of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 3.2

Quillstone Ledger 3.2 was released on 8 June 2019 for customers on the hosted plan.

What changed:

- Scheduled reports: now accepts the exchange rate used as a filter. The filter applies to records created on or after 3.2.
- The approval queue: can now be turned off for each user. When it is turned off, the attachment count is not collected for that user from then on.
- Exports: now includes a column for the reviewer's comment count. The column is empty for records created before 3.2, because the reviewer's comment count was not stored then.
- Invoice matching: can now be turned off for each user. When it is turned off, the time zone of the entry is not collected for that user from then on.

- Fixed: bank feeds did not keep the sort order after a page reload. It now keeps it until the user changes it.

Known issue: on slow connections currency conversion can take up to 10 seconds to load. We are working on it and will say here when it is fixed.

Planned, not yet available: the original currency in the tax summary. We expect to release it in the second half of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 3.3

Quillstone Ledger 3.3 was released on 15 June 2019 for customers on the hosted plan.

What changed:

- Currency conversion: can now be turned off for each user. When it is turned off, the attachment count is not collected for that user from then on.
- Currency conversion: now includes a column for the device type. The column is empty for records created before 3.3, because the device type was not stored then.
- The audit view: now accepts the batch number as a filter. The filter applies to records created on or after 3.3.

- Fixed: the activity log did not keep the sort order after a page reload. It now keeps it until the user changes it.
- Fixed: in the search index, a month with no records showed a blank page. It now shows an empty table.

Known issue: the tax summary does not yet support files larger than 25 MB. Larger files are refused with a message that names the limit.

Planned, not yet available: the import job number in the approval queue. We expect to release it in the second half of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 3.4

Quillstone Ledger 3.4 was released on 22 June 2019 for customers on the hosted plan.

What changed:

- User roles: now shows the original currency beside each entry. Entries made before 3.4 show a dash, because the value was not recorded.
- Bank feeds: can now be turned off for each user. When it is turned off, the batch number is not collected for that user from then on.
- Invoice matching: now names the user who last changed the source file name. It does not show why the change was made, because the product does not ask for a reason.

- Fixed: account settings did not keep the sort order after a page reload. It now keeps it until the user changes it.
- Fixed: a date typed as day/month could be read as month/day in invoice matching. The format now follows the account's settings.

Known issue: on slow connections the audit view can take up to 20 seconds to load. We are working on it and will say here when it is fixed.

## Quillstone Ledger 3.5

Quillstone Ledger 3.5 was released on 1 July 2019 for customers on the hosted plan.

What changed:

- The tax summary: now shows the device type beside each entry. Entries made before 3.5 show a dash, because the value was not recorded.
- User roles: now includes a column for the payment terms. The column is empty for records created before 3.5, because the payment terms was not stored then.
- The approval queue: now names the user who last changed the payment terms. It does not show why the change was made, because the product does not ask for a reason.
- Currency conversion: now shows the cost centre beside each entry. Entries made before 3.5 show a dash, because the value was not recorded.
- Report generation: now accepts the device type as a filter. The filter applies to records created on or after 3.5.

- Fixed: totals in the activity log were rounded before they were added, which could change the last digit. They are now added first and rounded once.

Known issue: on slow connections the tax summary can take up to 20 seconds to load. We are working on it and will say here when it is fixed.

## Quillstone Ledger 3.6

Quillstone Ledger 3.6 was released on 8 July 2019 for customers on the hosted plan.

What changed:

- The search index: now keeps the import job number for 30 days. After that the value is removed, and the entry shows when it was removed.
- Scheduled reports: can now be turned off for each user. When it is turned off, the batch number is not collected for that user from then on.
- The search index: now accepts the cost centre as a filter. The filter applies to records created on or after 3.6.

- Fixed: account settings did not keep the sort order after a page reload. It now keeps it until the user changes it.

Planned, not yet available: the attachment count in the audit view. We expect to release it in the second quarter of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 3.7

Quillstone Ledger 3.7 was released on 15 July 2019 for customers on the hosted plan.

What changed:

- Exports: now names the user who last changed the payment terms. It does not show why the change was made, because the product does not ask for a reason.
- The import screen: now accepts the import job number as a filter. The filter applies to records created on or after 3.7.
- Account settings: now includes a column for the attachment count. The column is empty for records created before 3.7, because the attachment count was not stored then.

- Fixed: the audit view did not keep the sort order after a page reload. It now keeps it until the user changes it.

Planned, not yet available: the reviewer's comment count in exports. We expect to release it in the first quarter of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 3.8

Quillstone Ledger 3.8 was released on 22 July 2019 for customers on the hosted plan.

What changed:

- The audit view: now names the user who last changed the attachment count. It does not show why the change was made, because the product does not ask for a reason.
- Bank feeds: now keeps the original currency for 90 days. After that the value is removed, and the entry shows when it was removed.
- The search index: now keeps the time zone of the entry for 365 days. After that the value is removed, and the entry shows when it was removed.
- Scheduled reports: can now be turned off for each user. When it is turned off, the source file name is not collected for that user from then on.

- Fixed: report generation did not keep the sort order after a page reload. It now keeps it until the user changes it.
- Fixed: in the tax summary, a month with no records showed a blank page. It now shows an empty table.

## Quillstone Ledger 3.9

Quillstone Ledger 3.9 was released on 1 August 2019 for customers on the hosted plan.

What changed:

- Account settings: now shows the time zone of the entry beside each entry. Entries made before 3.9 show a dash, because the value was not recorded.
- Currency conversion: now names the user who last changed the original currency. It does not show why the change was made, because the product does not ask for a reason.
- Exports: now keeps the attachment count for 365 days. After that the value is removed, and the entry shows when it was removed.

- Fixed: user roles did not keep the sort order after a page reload. It now keeps it until the user changes it.
- Fixed: a date typed as day/month could be read as month/day in the approval queue. The format now follows the account's settings.

Known issue: currency conversion does not yet support files larger than 50 MB. Larger files are refused with a message that names the limit.

Planned, not yet available: the device type in the approval queue. We expect to release it in the second half of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 4.0

Quillstone Ledger 4.0 was released on 8 August 2019 for customers on the hosted plan.

What changed:

- Report generation: now names the user who last changed the reviewer's comment count. It does not show why the change was made, because the product does not ask for a reason.
- Report generation: now accepts the cost centre as a filter. The filter applies to records created on or after 4.0.
- The activity log: faster on accounts with many records. In our own test on 8 August 2019, a 39,000-row run took 64 seconds on 3.9 and 21 seconds on 4.0. Your results depend on your data.
- Currency conversion: now accepts the batch number as a filter. The filter applies to records created on or after 4.0.
- Account settings: now accepts the batch number as a filter. The filter applies to records created on or after 4.0.

- Fixed: in exports, a month with no records showed a blank page. It now shows an empty table.

Known issue: on slow connections user roles can take up to 20 seconds to load. We are working on it and will say here when it is fixed.

## Quillstone Ledger 4.1

Quillstone Ledger 4.1 was released on 15 August 2019 for customers on the hosted plan.

What changed:

- The mobile view: can now be turned off for each user. When it is turned off, the source file name is not collected for that user from then on.
- Scheduled reports: faster on accounts with many records. In our own test on 15 August 2019, a 139,000-row run took 75 seconds on 4.0 and 18 seconds on 4.1. Your results depend on your data.
- Currency conversion: now includes a column for the payment terms. The column is empty for records created before 4.1, because the payment terms was not stored then.

- Fixed: totals in currency conversion were rounded before they were added, which could change the last digit. They are now added first and rounded once.

Known issue: account settings does not yet support files larger than 50 MB. Larger files are refused with a message that names the limit.

## Quillstone Ledger 4.2

Quillstone Ledger 4.2 was released on 22 August 2019 for customers on the hosted plan.

What changed:

- The import screen: now includes a column for the batch number. The column is empty for records created before 4.2, because the batch number was not stored then.
- Currency conversion: now accepts the exchange rate used as a filter. The filter applies to records created on or after 4.2.
- Invoice matching: now shows the batch number beside each entry. Entries made before 4.2 show a dash, because the value was not recorded.
- The mobile view: now accepts the approving user's role as a filter. The filter applies to records created on or after 4.2.
- User roles: now names the user who last changed the time zone of the entry. It does not show why the change was made, because the product does not ask for a reason.

- Fixed: a date typed as day/month could be read as month/day in the tax summary. The format now follows the account's settings.
- Fixed: totals in account settings were rounded before they were added, which could change the last digit. They are now added first and rounded once.

Known issue: on slow connections scheduled reports can take up to 20 seconds to load. We are working on it and will say here when it is fixed.

## Quillstone Ledger 4.3

Quillstone Ledger 4.3 was released on 1 September 2019 for customers on the hosted plan.

What changed:

- Report generation: faster on accounts with many records. In our own test on 1 September 2019, a 94,000-row run took 25 seconds on 4.2 and 6 seconds on 4.3. Your results depend on your data.
- The import screen: faster on accounts with many records. In our own test on 1 September 2019, a 126,000-row run took 55 seconds on 4.2 and 18 seconds on 4.3. Your results depend on your data.
- Account settings: now includes a column for the payment terms. The column is empty for records created before 4.3, because the payment terms was not stored then.

- Fixed: a date typed as day/month could be read as month/day in user roles. The format now follows the account's settings.
- Fixed: in attachment storage, a month with no records showed a blank page. It now shows an empty table.

Planned, not yet available: the approving user's role in scheduled reports. We expect to release it in the second quarter of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 4.4

Quillstone Ledger 4.4 was released on 8 September 2019 for customers on the hosted plan.

What changed:

- The mobile view: can now be turned off for each user. When it is turned off, the device type is not collected for that user from then on.
- Invoice matching: now keeps the original currency for 180 days. After that the value is removed, and the entry shows when it was removed.
- The activity log: now shows the original currency beside each entry. Entries made before 4.4 show a dash, because the value was not recorded.
- Exports: now shows the payment terms beside each entry. Entries made before 4.4 show a dash, because the value was not recorded.

- Fixed: in the audit view, a month with no records showed a blank page. It now shows an empty table.
- Fixed: a date typed as day/month could be read as month/day in the tax summary. The format now follows the account's settings.

## Quillstone Ledger 4.5

Quillstone Ledger 4.5 was released on 15 September 2019 for customers on the hosted plan.

What changed:

- User roles: can now be turned off for each user. When it is turned off, the device type is not collected for that user from then on.
- The search index: can now be turned off for each user. When it is turned off, the import job number is not collected for that user from then on.
- The audit view: now accepts the exchange rate used as a filter. The filter applies to records created on or after 4.5.
- Currency conversion: now accepts the cost centre as a filter. The filter applies to records created on or after 4.5.

- Fixed: in the activity log, a month with no records showed a blank page. It now shows an empty table.
- Fixed: in the mobile view, a month with no records showed a blank page. It now shows an empty table.

Known issue: on slow connections user roles can take up to 10 seconds to load. We are working on it and will say here when it is fixed.

## Quillstone Ledger 4.6

Quillstone Ledger 4.6 was released on 22 September 2019 for customers on the hosted plan.

What changed:

- Report generation: can now be turned off for each user. When it is turned off, the original currency is not collected for that user from then on.
- The activity log: faster on accounts with many records. In our own test on 22 September 2019, a 126,000-row run took 62 seconds on 4.5 and 31 seconds on 4.6. Your results depend on your data.
- Scheduled reports: now names the user who last changed the time zone of the entry. It does not show why the change was made, because the product does not ask for a reason.
- The mobile view: now accepts the device type as a filter. The filter applies to records created on or after 4.6.

- Fixed: in user roles, a month with no records showed a blank page. It now shows an empty table.

## Quillstone Ledger 4.7

Quillstone Ledger 4.7 was released on 1 October 2019 for customers on the hosted plan.

What changed:

- Exports: now keeps the device type for 30 days. After that the value is removed, and the entry shows when it was removed.
- Attachment storage: now accepts the payment terms as a filter. The filter applies to records created on or after 4.7.
- Account settings: now names the user who last changed the approving user's role. It does not show why the change was made, because the product does not ask for a reason.

- Fixed: a date typed as day/month could be read as month/day in the mobile view. The format now follows the account's settings.
- Fixed: the activity log did not keep the sort order after a page reload. It now keeps it until the user changes it.

Planned, not yet available: the source file name in the search index. We expect to release it in the second quarter of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 4.8

Quillstone Ledger 4.8 was released on 8 October 2019 for customers on the hosted plan.

What changed:

- Bank feeds: can now be turned off for each user. When it is turned off, the source file name is not collected for that user from then on.
- The activity log: can now be turned off for each user. When it is turned off, the cost centre is not collected for that user from then on.
- The approval queue: faster on accounts with many records. In our own test on 8 October 2019, a 167,000-row run took 54 seconds on 4.7 and 13 seconds on 4.8. Your results depend on your data.

- Fixed: in the import screen, a month with no records showed a blank page. It now shows an empty table.

Known issue: on slow connections bank feeds can take up to 20 seconds to load. We are working on it and will say here when it is fixed.

Planned, not yet available: the cost centre in report generation. We expect to release it in the second quarter of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 4.9

Quillstone Ledger 4.9 was released on 15 October 2019 for customers on the hosted plan.

What changed:

- The mobile view: now includes a column for the exchange rate used. The column is empty for records created before 4.9, because the exchange rate used was not stored then.
- The search index: faster on accounts with many records. In our own test on 15 October 2019, a 151,000-row run took 51 seconds on 4.8 and 17 seconds on 4.9. Your results depend on your data.
- Bank feeds: now shows the time zone of the entry beside each entry. Entries made before 4.9 show a dash, because the value was not recorded.

- Fixed: scheduled reports did not keep the sort order after a page reload. It now keeps it until the user changes it.
- Fixed: in account settings, a month with no records showed a blank page. It now shows an empty table.

Known issue: on slow connections the import screen can take up to 10 seconds to load. We are working on it and will say here when it is fixed.

## Quillstone Ledger 5.0

Quillstone Ledger 5.0 was released on 22 October 2019 for customers on the hosted plan.

What changed:

- Attachment storage: now shows the attachment count beside each entry. Entries made before 5.0 show a dash, because the value was not recorded.
- Attachment storage: can now be turned off for each user. When it is turned off, the exchange rate used is not collected for that user from then on.
- The approval queue: now shows the import job number beside each entry. Entries made before 5.0 show a dash, because the value was not recorded.
- Invoice matching: now accepts the payment terms as a filter. The filter applies to records created on or after 5.0.

- Fixed: totals in bank feeds were rounded before they were added, which could change the last digit. They are now added first and rounded once.
- Fixed: totals in invoice matching were rounded before they were added, which could change the last digit. They are now added first and rounded once.

Known issue: on slow connections invoice matching can take up to 10 seconds to load. We are working on it and will say here when it is fixed.

## Quillstone Ledger 5.1

Quillstone Ledger 5.1 was released on 1 November 2019 for customers on the hosted plan.

What changed:

- User roles: now shows the cost centre beside each entry. Entries made before 5.1 show a dash, because the value was not recorded.
- Bank feeds: now shows the import job number beside each entry. Entries made before 5.1 show a dash, because the value was not recorded.
- Currency conversion: now accepts the exchange rate used as a filter. The filter applies to records created on or after 5.1.
- The mobile view: now keeps the device type for 180 days. After that the value is removed, and the entry shows when it was removed.
- The audit view: now accepts the approving user's role as a filter. The filter applies to records created on or after 5.1.

- Fixed: a date typed as day/month could be read as month/day in account settings. The format now follows the account's settings.
- Fixed: exports did not keep the sort order after a page reload. It now keeps it until the user changes it.

## Quillstone Ledger 5.2

Quillstone Ledger 5.2 was released on 8 November 2019 for customers on the hosted plan.

What changed:

- Currency conversion: now includes a column for the reviewer's comment count. The column is empty for records created before 5.2, because the reviewer's comment count was not stored then.
- Report generation: now includes a column for the device type. The column is empty for records created before 5.2, because the device type was not stored then.
- The tax summary: now keeps the import job number for 30 days. After that the value is removed, and the entry shows when it was removed.
- Currency conversion: now includes a column for the batch number. The column is empty for records created before 5.2, because the batch number was not stored then.
- Account settings: now includes a column for the source file name. The column is empty for records created before 5.2, because the source file name was not stored then.

- Fixed: in the tax summary, a month with no records showed a blank page. It now shows an empty table.
- Fixed: in report generation, a month with no records showed a blank page. It now shows an empty table.

## Quillstone Ledger 5.3

Quillstone Ledger 5.3 was released on 15 November 2019 for customers on the hosted plan.

What changed:

- Bank feeds: now keeps the exchange rate used for 30 days. After that the value is removed, and the entry shows when it was removed.
- The approval queue: can now be turned off for each user. When it is turned off, the exchange rate used is not collected for that user from then on.
- The audit view: faster on accounts with many records. In our own test on 15 November 2019, a 92,000-row run took 32 seconds on 5.2 and 8 seconds on 5.3. Your results depend on your data.
- The activity log: can now be turned off for each user. When it is turned off, the device type is not collected for that user from then on.

- Fixed: totals in user roles were rounded before they were added, which could change the last digit. They are now added first and rounded once.

Known issue: on slow connections the import screen can take up to 10 seconds to load. We are working on it and will say here when it is fixed.

## Quillstone Ledger 5.4

Quillstone Ledger 5.4 was released on 22 November 2019 for customers on the hosted plan.

What changed:

- Account settings: now keeps the exchange rate used for 90 days. After that the value is removed, and the entry shows when it was removed.
- The mobile view: now includes a column for the source file name. The column is empty for records created before 5.4, because the source file name was not stored then.
- User roles: can now be turned off for each user. When it is turned off, the import job number is not collected for that user from then on.
- Scheduled reports: now names the user who last changed the batch number. It does not show why the change was made, because the product does not ask for a reason.
- The audit view: now shows the cost centre beside each entry. Entries made before 5.4 show a dash, because the value was not recorded.

- Fixed: in account settings, a month with no records showed a blank page. It now shows an empty table.
- Fixed: totals in exports were rounded before they were added, which could change the last digit. They are now added first and rounded once.

## Quillstone Ledger 5.5

Quillstone Ledger 5.5 was released on 1 December 2019 for customers on the hosted plan.

What changed:

- The tax summary: now accepts the batch number as a filter. The filter applies to records created on or after 5.5.
- The activity log: now names the user who last changed the attachment count. It does not show why the change was made, because the product does not ask for a reason.
- Report generation: now includes a column for the device type. The column is empty for records created before 5.5, because the device type was not stored then.
- The audit view: can now be turned off for each user. When it is turned off, the source file name is not collected for that user from then on.

- Fixed: a date typed as day/month could be read as month/day in scheduled reports. The format now follows the account's settings.
- Fixed: a date typed as day/month could be read as month/day in the approval queue. The format now follows the account's settings.

Known issue: attachment storage does not yet support files larger than 25 MB. Larger files are refused with a message that names the limit.

## Quillstone Ledger 5.6

Quillstone Ledger 5.6 was released on 8 December 2019 for customers on the hosted plan.

What changed:

- Bank feeds: now accepts the device type as a filter. The filter applies to records created on or after 5.6.
- The tax summary: now accepts the reviewer's comment count as a filter. The filter applies to records created on or after 5.6.
- Attachment storage: can now be turned off for each user. When it is turned off, the reviewer's comment count is not collected for that user from then on.
- The audit view: faster on accounts with many records. In our own test on 8 December 2019, a 166,000-row run took 72 seconds on 5.5 and 18 seconds on 5.6. Your results depend on your data.
- Exports: can now be turned off for each user. When it is turned off, the batch number is not collected for that user from then on.

- Fixed: in account settings, a month with no records showed a blank page. It now shows an empty table.
- Fixed: in the approval queue, a month with no records showed a blank page. It now shows an empty table.

Known issue: on slow connections the tax summary can take up to 20 seconds to load. We are working on it and will say here when it is fixed.

Planned, not yet available: the approving user's role in invoice matching. We expect to release it in the first quarter of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 5.7

Quillstone Ledger 5.7 was released on 15 December 2019 for customers on the hosted plan.

What changed:

- The audit view: now shows the approving user's role beside each entry. Entries made before 5.7 show a dash, because the value was not recorded.
- Invoice matching: now keeps the exchange rate used for 30 days. After that the value is removed, and the entry shows when it was removed.
- Currency conversion: now keeps the reviewer's comment count for 180 days. After that the value is removed, and the entry shows when it was removed.
- Attachment storage: now shows the payment terms beside each entry. Entries made before 5.7 show a dash, because the value was not recorded.

- Fixed: user roles did not keep the sort order after a page reload. It now keeps it until the user changes it.
- Fixed: a date typed as day/month could be read as month/day in the audit view. The format now follows the account's settings.

Planned, not yet available: the time zone of the entry in user roles. We expect to release it in the first quarter of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 5.8

Quillstone Ledger 5.8 was released on 22 December 2019 for customers on the hosted plan.

What changed:

- Attachment storage: now accepts the source file name as a filter. The filter applies to records created on or after 5.8.
- The audit view: can now be turned off for each user. When it is turned off, the cost centre is not collected for that user from then on.
- The mobile view: now shows the cost centre beside each entry. Entries made before 5.8 show a dash, because the value was not recorded.
- The audit view: now includes a column for the cost centre. The column is empty for records created before 5.8, because the cost centre was not stored then.
- Invoice matching: now includes a column for the payment terms. The column is empty for records created before 5.8, because the payment terms was not stored then.

- Fixed: the tax summary did not keep the sort order after a page reload. It now keeps it until the user changes it.

Planned, not yet available: the original currency in the import screen. We expect to release it in the second quarter of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 5.9

Quillstone Ledger 5.9 was released on 1 January 2020 for customers on the hosted plan.

What changed:

- The approval queue: now accepts the payment terms as a filter. The filter applies to records created on or after 5.9.
- User roles: can now be turned off for each user. When it is turned off, the cost centre is not collected for that user from then on.
- The mobile view: now includes a column for the reviewer's comment count. The column is empty for records created before 5.9, because the reviewer's comment count was not stored then.
- Currency conversion: now accepts the original currency as a filter. The filter applies to records created on or after 5.9.
- Exports: now keeps the original currency for 365 days. After that the value is removed, and the entry shows when it was removed.

- Fixed: a date typed as day/month could be read as month/day in attachment storage. The format now follows the account's settings.

## Quillstone Ledger 6.0

Quillstone Ledger 6.0 was released on 8 January 2020 for customers on the hosted plan.

What changed:

- Attachment storage: now accepts the attachment count as a filter. The filter applies to records created on or after 6.0.
- The activity log: can now be turned off for each user. When it is turned off, the import job number is not collected for that user from then on.
- Bank feeds: now keeps the batch number for 365 days. After that the value is removed, and the entry shows when it was removed.

- Fixed: report generation did not keep the sort order after a page reload. It now keeps it until the user changes it.

## Quillstone Ledger 6.1

Quillstone Ledger 6.1 was released on 15 January 2020 for customers on the hosted plan.

What changed:

- Exports: now shows the exchange rate used beside each entry. Entries made before 6.1 show a dash, because the value was not recorded.
- Currency conversion: now keeps the import job number for 180 days. After that the value is removed, and the entry shows when it was removed.
- Exports: faster on accounts with many records. In our own test on 15 January 2020, a 126,000-row run took 78 seconds on 6.0 and 15 seconds on 6.1. Your results depend on your data.
- Attachment storage: now names the user who last changed the reviewer's comment count. It does not show why the change was made, because the product does not ask for a reason.

- Fixed: totals in the import screen were rounded before they were added, which could change the last digit. They are now added first and rounded once.

## Quillstone Ledger 6.2

Quillstone Ledger 6.2 was released on 22 January 2020 for customers on the hosted plan.

What changed:

- Exports: now keeps the reviewer's comment count for 30 days. After that the value is removed, and the entry shows when it was removed.
- Scheduled reports: faster on accounts with many records. In our own test on 22 January 2020, a 42,000-row run took 84 seconds on 6.1 and 28 seconds on 6.2. Your results depend on your data.
- The audit view: now includes a column for the attachment count. The column is empty for records created before 6.2, because the attachment count was not stored then.
- The approval queue: now shows the source file name beside each entry. Entries made before 6.2 show a dash, because the value was not recorded.

- Fixed: in the search index, a month with no records showed a blank page. It now shows an empty table.
- Fixed: a date typed as day/month could be read as month/day in exports. The format now follows the account's settings.

## Quillstone Ledger 6.3

Quillstone Ledger 6.3 was released on 1 February 2020 for customers on the hosted plan.

What changed:

- The approval queue: now keeps the source file name for 90 days. After that the value is removed, and the entry shows when it was removed.
- The approval queue: now keeps the exchange rate used for 365 days. After that the value is removed, and the entry shows when it was removed.
- Invoice matching: now includes a column for the approving user's role. The column is empty for records created before 6.3, because the approving user's role was not stored then.
- Invoice matching: faster on accounts with many records. In our own test on 1 February 2020, a 72,000-row run took 65 seconds on 6.2 and 13 seconds on 6.3. Your results depend on your data.

- Fixed: a date typed as day/month could be read as month/day in the search index. The format now follows the account's settings.
- Fixed: in report generation, a month with no records showed a blank page. It now shows an empty table.

## Quillstone Ledger 6.4

Quillstone Ledger 6.4 was released on 8 February 2020 for customers on the hosted plan.

What changed:

- Bank feeds: now keeps the exchange rate used for 180 days. After that the value is removed, and the entry shows when it was removed.
- The tax summary: now shows the source file name beside each entry. Entries made before 6.4 show a dash, because the value was not recorded.
- The import screen: now keeps the time zone of the entry for 90 days. After that the value is removed, and the entry shows when it was removed.
- Account settings: now names the user who last changed the attachment count. It does not show why the change was made, because the product does not ask for a reason.
- The approval queue: now shows the payment terms beside each entry. Entries made before 6.4 show a dash, because the value was not recorded.

- Fixed: a date typed as day/month could be read as month/day in scheduled reports. The format now follows the account's settings.
- Fixed: totals in exports were rounded before they were added, which could change the last digit. They are now added first and rounded once.

Planned, not yet available: the original currency in the mobile view. We expect to release it in the first quarter of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 6.5

Quillstone Ledger 6.5 was released on 15 February 2020 for customers on the hosted plan.

What changed:

- The approval queue: faster on accounts with many records. In our own test on 15 February 2020, a 125,000-row run took 29 seconds on 6.4 and 9 seconds on 6.5. Your results depend on your data.
- Account settings: now includes a column for the exchange rate used. The column is empty for records created before 6.5, because the exchange rate used was not stored then.
- Bank feeds: now names the user who last changed the approving user's role. It does not show why the change was made, because the product does not ask for a reason.
- Currency conversion: now shows the source file name beside each entry. Entries made before 6.5 show a dash, because the value was not recorded.

- Fixed: totals in invoice matching were rounded before they were added, which could change the last digit. They are now added first and rounded once.
- Fixed: a date typed as day/month could be read as month/day in exports. The format now follows the account's settings.

## Quillstone Ledger 6.6

Quillstone Ledger 6.6 was released on 22 February 2020 for customers on the hosted plan.

What changed:

- Report generation: now includes a column for the approving user's role. The column is empty for records created before 6.6, because the approving user's role was not stored then.
- The audit view: now includes a column for the reviewer's comment count. The column is empty for records created before 6.6, because the reviewer's comment count was not stored then.
- Report generation: now shows the import job number beside each entry. Entries made before 6.6 show a dash, because the value was not recorded.
- Invoice matching: now names the user who last changed the time zone of the entry. It does not show why the change was made, because the product does not ask for a reason.

- Fixed: totals in user roles were rounded before they were added, which could change the last digit. They are now added first and rounded once.

Known issue: on slow connections the mobile view can take up to 20 seconds to load. We are working on it and will say here when it is fixed.

## Quillstone Ledger 6.7

Quillstone Ledger 6.7 was released on 1 March 2020 for customers on the hosted plan.

What changed:

- The import screen: faster on accounts with many records. In our own test on 1 March 2020, a 161,000-row run took 30 seconds on 6.6 and 10 seconds on 6.7. Your results depend on your data.
- Account settings: now accepts the batch number as a filter. The filter applies to records created on or after 6.7.
- The tax summary: can now be turned off for each user. When it is turned off, the reviewer's comment count is not collected for that user from then on.
- Invoice matching: now keeps the exchange rate used for 30 days. After that the value is removed, and the entry shows when it was removed.

- Fixed: a date typed as day/month could be read as month/day in the import screen. The format now follows the account's settings.
- Fixed: a date typed as day/month could be read as month/day in currency conversion. The format now follows the account's settings.

Known issue: currency conversion does not yet support files larger than 50 MB. Larger files are refused with a message that names the limit.

## Quillstone Ledger 6.8

Quillstone Ledger 6.8 was released on 8 March 2020 for customers on the hosted plan.

What changed:

- Report generation: now keeps the exchange rate used for 180 days. After that the value is removed, and the entry shows when it was removed.
- Bank feeds: faster on accounts with many records. In our own test on 8 March 2020, a 47,000-row run took 24 seconds on 6.7 and 12 seconds on 6.8. Your results depend on your data.
- The approval queue: faster on accounts with many records. In our own test on 8 March 2020, a 75,000-row run took 71 seconds on 6.7 and 17 seconds on 6.8. Your results depend on your data.

- Fixed: in report generation, a month with no records showed a blank page. It now shows an empty table.
- Fixed: in attachment storage, a month with no records showed a blank page. It now shows an empty table.

## Quillstone Ledger 6.9

Quillstone Ledger 6.9 was released on 15 March 2020 for customers on the hosted plan.

What changed:

- The mobile view: can now be turned off for each user. When it is turned off, the attachment count is not collected for that user from then on.
- Account settings: now keeps the exchange rate used for 90 days. After that the value is removed, and the entry shows when it was removed.
- User roles: now keeps the import job number for 180 days. After that the value is removed, and the entry shows when it was removed.
- User roles: faster on accounts with many records. In our own test on 15 March 2020, a 52,000-row run took 72 seconds on 6.8 and 36 seconds on 6.9. Your results depend on your data.
- The tax summary: now shows the payment terms beside each entry. Entries made before 6.9 show a dash, because the value was not recorded.

- Fixed: totals in user roles were rounded before they were added, which could change the last digit. They are now added first and rounded once.
- Fixed: in the import screen, a month with no records showed a blank page. It now shows an empty table.

Known issue: the import screen does not yet support files larger than 25 MB. Larger files are refused with a message that names the limit.

Planned, not yet available: the exchange rate used in exports. We expect to release it in the first quarter of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 7.0

Quillstone Ledger 7.0 was released on 22 March 2020 for customers on the hosted plan.

What changed:

- Currency conversion: faster on accounts with many records. In our own test on 22 March 2020, a 164,000-row run took 25 seconds on 6.9 and 12 seconds on 7.0. Your results depend on your data.
- Account settings: can now be turned off for each user. When it is turned off, the batch number is not collected for that user from then on.
- The tax summary: can now be turned off for each user. When it is turned off, the payment terms is not collected for that user from then on.

- Fixed: in bank feeds, a month with no records showed a blank page. It now shows an empty table.
- Fixed: a date typed as day/month could be read as month/day in bank feeds. The format now follows the account's settings.

Known issue: the mobile view does not yet support files larger than 100 MB. Larger files are refused with a message that names the limit.

## Quillstone Ledger 7.1

Quillstone Ledger 7.1 was released on 1 April 2020 for customers on the hosted plan.

What changed:

- The activity log: faster on accounts with many records. In our own test on 1 April 2020, a 33,000-row run took 47 seconds on 7.0 and 9 seconds on 7.1. Your results depend on your data.
- The tax summary: now accepts the import job number as a filter. The filter applies to records created on or after 7.1.
- The approval queue: now names the user who last changed the time zone of the entry. It does not show why the change was made, because the product does not ask for a reason.
- Currency conversion: now names the user who last changed the source file name. It does not show why the change was made, because the product does not ask for a reason.
- Bank feeds: now names the user who last changed the exchange rate used. It does not show why the change was made, because the product does not ask for a reason.

- Fixed: bank feeds did not keep the sort order after a page reload. It now keeps it until the user changes it.
- Fixed: a date typed as day/month could be read as month/day in the audit view. The format now follows the account's settings.

Known issue: the activity log does not yet support files larger than 25 MB. Larger files are refused with a message that names the limit.

## Quillstone Ledger 7.2

Quillstone Ledger 7.2 was released on 8 April 2020 for customers on the hosted plan.

What changed:

- The search index: now accepts the exchange rate used as a filter. The filter applies to records created on or after 7.2.
- The tax summary: now names the user who last changed the import job number. It does not show why the change was made, because the product does not ask for a reason.
- The tax summary: can now be turned off for each user. When it is turned off, the time zone of the entry is not collected for that user from then on.

- Fixed: a date typed as day/month could be read as month/day in user roles. The format now follows the account's settings.

## Quillstone Ledger 7.3

Quillstone Ledger 7.3 was released on 15 April 2020 for customers on the hosted plan.

What changed:

- Bank feeds: faster on accounts with many records. In our own test on 15 April 2020, a 132,000-row run took 77 seconds on 7.2 and 38 seconds on 7.3. Your results depend on your data.
- The import screen: now names the user who last changed the exchange rate used. It does not show why the change was made, because the product does not ask for a reason.
- The mobile view: now names the user who last changed the payment terms. It does not show why the change was made, because the product does not ask for a reason.
- The tax summary: now keeps the time zone of the entry for 180 days. After that the value is removed, and the entry shows when it was removed.

- Fixed: a date typed as day/month could be read as month/day in report generation. The format now follows the account's settings.
- Fixed: totals in scheduled reports were rounded before they were added, which could change the last digit. They are now added first and rounded once.

Known issue: currency conversion does not yet support files larger than 100 MB. Larger files are refused with a message that names the limit.

Planned, not yet available: the payment terms in exports. We expect to release it in the first quarter of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 7.4

Quillstone Ledger 7.4 was released on 22 April 2020 for customers on the hosted plan.

What changed:

- Exports: now accepts the exchange rate used as a filter. The filter applies to records created on or after 7.4.
- The audit view: now includes a column for the reviewer's comment count. The column is empty for records created before 7.4, because the reviewer's comment count was not stored then.
- The tax summary: faster on accounts with many records. In our own test on 22 April 2020, a 186,000-row run took 77 seconds on 7.3 and 15 seconds on 7.4. Your results depend on your data.
- The mobile view: now names the user who last changed the device type. It does not show why the change was made, because the product does not ask for a reason.
- The mobile view: now shows the attachment count beside each entry. Entries made before 7.4 show a dash, because the value was not recorded.

- Fixed: in the activity log, a month with no records showed a blank page. It now shows an empty table.

Known issue: on slow connections invoice matching can take up to 10 seconds to load. We are working on it and will say here when it is fixed.

## Quillstone Ledger 7.5

Quillstone Ledger 7.5 was released on 1 May 2020 for customers on the hosted plan.

What changed:

- User roles: now includes a column for the exchange rate used. The column is empty for records created before 7.5, because the exchange rate used was not stored then.
- The mobile view: can now be turned off for each user. When it is turned off, the device type is not collected for that user from then on.
- The mobile view: now accepts the exchange rate used as a filter. The filter applies to records created on or after 7.5.

- Fixed: in user roles, a month with no records showed a blank page. It now shows an empty table.

Known issue: the activity log does not yet support files larger than 25 MB. Larger files are refused with a message that names the limit.

## Quillstone Ledger 7.6

Quillstone Ledger 7.6 was released on 8 May 2020 for customers on the hosted plan.

What changed:

- Invoice matching: can now be turned off for each user. When it is turned off, the batch number is not collected for that user from then on.
- Exports: now names the user who last changed the batch number. It does not show why the change was made, because the product does not ask for a reason.
- User roles: now accepts the batch number as a filter. The filter applies to records created on or after 7.6.
- Currency conversion: faster on accounts with many records. In our own test on 8 May 2020, a 28,000-row run took 45 seconds on 7.5 and 9 seconds on 7.6. Your results depend on your data.
- The activity log: now includes a column for the payment terms. The column is empty for records created before 7.6, because the payment terms was not stored then.

- Fixed: a date typed as day/month could be read as month/day in report generation. The format now follows the account's settings.

Known issue: on slow connections the approval queue can take up to 20 seconds to load. We are working on it and will say here when it is fixed.

## Quillstone Ledger 7.7

Quillstone Ledger 7.7 was released on 15 May 2020 for customers on the hosted plan.

What changed:

- User roles: now includes a column for the reviewer's comment count. The column is empty for records created before 7.7, because the reviewer's comment count was not stored then.
- Currency conversion: now includes a column for the attachment count. The column is empty for records created before 7.7, because the attachment count was not stored then.
- The activity log: now accepts the batch number as a filter. The filter applies to records created on or after 7.7.

- Fixed: a date typed as day/month could be read as month/day in the approval queue. The format now follows the account's settings.
- Fixed: in the search index, a month with no records showed a blank page. It now shows an empty table.

## Quillstone Ledger 7.8

Quillstone Ledger 7.8 was released on 22 May 2020 for customers on the hosted plan.

What changed:

- Account settings: now accepts the batch number as a filter. The filter applies to records created on or after 7.8.
- User roles: now includes a column for the cost centre. The column is empty for records created before 7.8, because the cost centre was not stored then.
- Invoice matching: now keeps the batch number for 365 days. After that the value is removed, and the entry shows when it was removed.
- Invoice matching: now accepts the exchange rate used as a filter. The filter applies to records created on or after 7.8.
- Bank feeds: now shows the time zone of the entry beside each entry. Entries made before 7.8 show a dash, because the value was not recorded.

- Fixed: report generation did not keep the sort order after a page reload. It now keeps it until the user changes it.
- Fixed: in attachment storage, a month with no records showed a blank page. It now shows an empty table.

Known issue: the import screen does not yet support files larger than 50 MB. Larger files are refused with a message that names the limit.

Planned, not yet available: the import job number in scheduled reports. We expect to release it in the first quarter of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 7.9

Quillstone Ledger 7.9 was released on 1 June 2020 for customers on the hosted plan.

What changed:

- The approval queue: now names the user who last changed the original currency. It does not show why the change was made, because the product does not ask for a reason.
- The approval queue: now accepts the batch number as a filter. The filter applies to records created on or after 7.9.
- The activity log: faster on accounts with many records. In our own test on 1 June 2020, a 26,000-row run took 20 seconds on 7.8 and 4 seconds on 7.9. Your results depend on your data.
- Exports: faster on accounts with many records. In our own test on 1 June 2020, a 152,000-row run took 86 seconds on 7.8 and 28 seconds on 7.9. Your results depend on your data.
- Account settings: can now be turned off for each user. When it is turned off, the payment terms is not collected for that user from then on.

- Fixed: in the tax summary, a month with no records showed a blank page. It now shows an empty table.

Known issue: the tax summary does not yet support files larger than 25 MB. Larger files are refused with a message that names the limit.

## Quillstone Ledger 8.0

Quillstone Ledger 8.0 was released on 8 June 2020 for customers on the hosted plan.

What changed:

- Invoice matching: now names the user who last changed the original currency. It does not show why the change was made, because the product does not ask for a reason.
- The approval queue: faster on accounts with many records. In our own test on 8 June 2020, a 91,000-row run took 60 seconds on 7.9 and 20 seconds on 8.0. Your results depend on your data.
- The approval queue: now includes a column for the exchange rate used. The column is empty for records created before 8.0, because the exchange rate used was not stored then.

- Fixed: a date typed as day/month could be read as month/day in the approval queue. The format now follows the account's settings.
- Fixed: totals in invoice matching were rounded before they were added, which could change the last digit. They are now added first and rounded once.

## Quillstone Ledger 8.1

Quillstone Ledger 8.1 was released on 15 June 2020 for customers on the hosted plan.

What changed:

- The search index: now names the user who last changed the import job number. It does not show why the change was made, because the product does not ask for a reason.
- The approval queue: now keeps the payment terms for 365 days. After that the value is removed, and the entry shows when it was removed.
- The import screen: now shows the time zone of the entry beside each entry. Entries made before 8.1 show a dash, because the value was not recorded.
- The activity log: faster on accounts with many records. In our own test on 15 June 2020, a 100,000-row run took 47 seconds on 8.0 and 11 seconds on 8.1. Your results depend on your data.

- Fixed: a date typed as day/month could be read as month/day in report generation. The format now follows the account's settings.

## Quillstone Ledger 8.2

Quillstone Ledger 8.2 was released on 22 June 2020 for customers on the hosted plan.

What changed:

- The tax summary: now names the user who last changed the reviewer's comment count. It does not show why the change was made, because the product does not ask for a reason.
- Currency conversion: now names the user who last changed the time zone of the entry. It does not show why the change was made, because the product does not ask for a reason.
- Invoice matching: faster on accounts with many records. In our own test on 22 June 2020, a 151,000-row run took 47 seconds on 8.1 and 23 seconds on 8.2. Your results depend on your data.

- Fixed: in bank feeds, a month with no records showed a blank page. It now shows an empty table.
- Fixed: totals in attachment storage were rounded before they were added, which could change the last digit. They are now added first and rounded once.

## Quillstone Ledger 8.3

Quillstone Ledger 8.3 was released on 1 July 2020 for customers on the hosted plan.

What changed:

- Attachment storage: now accepts the import job number as a filter. The filter applies to records created on or after 8.3.
- The activity log: can now be turned off for each user. When it is turned off, the payment terms is not collected for that user from then on.
- Bank feeds: now includes a column for the attachment count. The column is empty for records created before 8.3, because the attachment count was not stored then.

- Fixed: totals in the activity log were rounded before they were added, which could change the last digit. They are now added first and rounded once.

Planned, not yet available: the approving user's role in bank feeds. We expect to release it in the second quarter of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 8.4

Quillstone Ledger 8.4 was released on 8 July 2020 for customers on the hosted plan.

What changed:

- The mobile view: now accepts the import job number as a filter. The filter applies to records created on or after 8.4.
- The import screen: now includes a column for the batch number. The column is empty for records created before 8.4, because the batch number was not stored then.
- Account settings: now accepts the payment terms as a filter. The filter applies to records created on or after 8.4.
- Scheduled reports: can now be turned off for each user. When it is turned off, the cost centre is not collected for that user from then on.
- Account settings: now names the user who last changed the original currency. It does not show why the change was made, because the product does not ask for a reason.

- Fixed: totals in the tax summary were rounded before they were added, which could change the last digit. They are now added first and rounded once.
- Fixed: in attachment storage, a month with no records showed a blank page. It now shows an empty table.

## Quillstone Ledger 8.5

Quillstone Ledger 8.5 was released on 15 July 2020 for customers on the hosted plan.

What changed:

- Invoice matching: now includes a column for the exchange rate used. The column is empty for records created before 8.5, because the exchange rate used was not stored then.
- Account settings: faster on accounts with many records. In our own test on 15 July 2020, a 51,000-row run took 89 seconds on 8.4 and 22 seconds on 8.5. Your results depend on your data.
- Bank feeds: now includes a column for the batch number. The column is empty for records created before 8.5, because the batch number was not stored then.
- The approval queue: now shows the import job number beside each entry. Entries made before 8.5 show a dash, because the value was not recorded.
- The tax summary: now keeps the source file name for 365 days. After that the value is removed, and the entry shows when it was removed.

- Fixed: a date typed as day/month could be read as month/day in invoice matching. The format now follows the account's settings.
- Fixed: a date typed as day/month could be read as month/day in exports. The format now follows the account's settings.

Planned, not yet available: the source file name in the activity log. We expect to release it in the first quarter of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 8.6

Quillstone Ledger 8.6 was released on 22 July 2020 for customers on the hosted plan.

What changed:

- Attachment storage: now includes a column for the approving user's role. The column is empty for records created before 8.6, because the approving user's role was not stored then.
- The audit view: now names the user who last changed the batch number. It does not show why the change was made, because the product does not ask for a reason.
- The search index: now shows the payment terms beside each entry. Entries made before 8.6 show a dash, because the value was not recorded.

- Fixed: the mobile view did not keep the sort order after a page reload. It now keeps it until the user changes it.

Planned, not yet available: the reviewer's comment count in report generation. We expect to release it in the second half of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 8.7

Quillstone Ledger 8.7 was released on 1 August 2020 for customers on the hosted plan.

What changed:

- The search index: now includes a column for the exchange rate used. The column is empty for records created before 8.7, because the exchange rate used was not stored then.
- The activity log: now shows the approving user's role beside each entry. Entries made before 8.7 show a dash, because the value was not recorded.
- The tax summary: now names the user who last changed the original currency. It does not show why the change was made, because the product does not ask for a reason.

- Fixed: totals in invoice matching were rounded before they were added, which could change the last digit. They are now added first and rounded once.
- Fixed: the tax summary did not keep the sort order after a page reload. It now keeps it until the user changes it.

Planned, not yet available: the source file name in the tax summary. We expect to release it in the second quarter of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 8.8

Quillstone Ledger 8.8 was released on 8 August 2020 for customers on the hosted plan.

What changed:

- Invoice matching: now keeps the original currency for 180 days. After that the value is removed, and the entry shows when it was removed.
- Scheduled reports: faster on accounts with many records. In our own test on 8 August 2020, a 187,000-row run took 40 seconds on 8.7 and 10 seconds on 8.8. Your results depend on your data.
- Report generation: now shows the exchange rate used beside each entry. Entries made before 8.8 show a dash, because the value was not recorded.
- The activity log: now accepts the import job number as a filter. The filter applies to records created on or after 8.8.

- Fixed: a date typed as day/month could be read as month/day in the tax summary. The format now follows the account's settings.

Known issue: on slow connections the search index can take up to 30 seconds to load. We are working on it and will say here when it is fixed.

## Quillstone Ledger 8.9

Quillstone Ledger 8.9 was released on 15 August 2020 for customers on the hosted plan.

What changed:

- Currency conversion: now accepts the import job number as a filter. The filter applies to records created on or after 8.9.
- The import screen: now accepts the cost centre as a filter. The filter applies to records created on or after 8.9.
- Invoice matching: now keeps the reviewer's comment count for 30 days. After that the value is removed, and the entry shows when it was removed.

- Fixed: totals in user roles were rounded before they were added, which could change the last digit. They are now added first and rounded once.

Planned, not yet available: the device type in report generation. We expect to release it in the second half of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 9.0

Quillstone Ledger 9.0 was released on 22 August 2020 for customers on the hosted plan.

What changed:

- The approval queue: now includes a column for the reviewer's comment count. The column is empty for records created before 9.0, because the reviewer's comment count was not stored then.
- Exports: now names the user who last changed the payment terms. It does not show why the change was made, because the product does not ask for a reason.
- The mobile view: now shows the original currency beside each entry. Entries made before 9.0 show a dash, because the value was not recorded.
- Account settings: now keeps the import job number for 180 days. After that the value is removed, and the entry shows when it was removed.

- Fixed: a date typed as day/month could be read as month/day in the tax summary. The format now follows the account's settings.
- Fixed: a date typed as day/month could be read as month/day in the activity log. The format now follows the account's settings.

Known issue: the import screen does not yet support files larger than 50 MB. Larger files are refused with a message that names the limit.

## Quillstone Ledger 9.1

Quillstone Ledger 9.1 was released on 1 September 2020 for customers on the hosted plan.

What changed:

- The audit view: now keeps the original currency for 90 days. After that the value is removed, and the entry shows when it was removed.
- Scheduled reports: now accepts the time zone of the entry as a filter. The filter applies to records created on or after 9.1.
- The import screen: now names the user who last changed the reviewer's comment count. It does not show why the change was made, because the product does not ask for a reason.

- Fixed: in the tax summary, a month with no records showed a blank page. It now shows an empty table.
- Fixed: totals in report generation were rounded before they were added, which could change the last digit. They are now added first and rounded once.

Known issue: the mobile view does not yet support files larger than 50 MB. Larger files are refused with a message that names the limit.

## Quillstone Ledger 9.2

Quillstone Ledger 9.2 was released on 8 September 2020 for customers on the hosted plan.

What changed:

- The search index: can now be turned off for each user. When it is turned off, the batch number is not collected for that user from then on.
- The search index: now shows the payment terms beside each entry. Entries made before 9.2 show a dash, because the value was not recorded.
- Attachment storage: now keeps the original currency for 180 days. After that the value is removed, and the entry shows when it was removed.
- The tax summary: now keeps the source file name for 180 days. After that the value is removed, and the entry shows when it was removed.

- Fixed: in currency conversion, a month with no records showed a blank page. It now shows an empty table.
- Fixed: account settings did not keep the sort order after a page reload. It now keeps it until the user changes it.

Planned, not yet available: the exchange rate used in the search index. We expect to release it in the second quarter of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 9.3

Quillstone Ledger 9.3 was released on 15 September 2020 for customers on the hosted plan.

What changed:

- The audit view: can now be turned off for each user. When it is turned off, the cost centre is not collected for that user from then on.
- Report generation: now accepts the reviewer's comment count as a filter. The filter applies to records created on or after 9.3.
- Currency conversion: now shows the time zone of the entry beside each entry. Entries made before 9.3 show a dash, because the value was not recorded.
- Report generation: now names the user who last changed the payment terms. It does not show why the change was made, because the product does not ask for a reason.
- Attachment storage: now accepts the exchange rate used as a filter. The filter applies to records created on or after 9.3.

- Fixed: the mobile view did not keep the sort order after a page reload. It now keeps it until the user changes it.

## Quillstone Ledger 9.4

Quillstone Ledger 9.4 was released on 22 September 2020 for customers on the hosted plan.

What changed:

- User roles: now names the user who last changed the exchange rate used. It does not show why the change was made, because the product does not ask for a reason.
- Currency conversion: now shows the time zone of the entry beside each entry. Entries made before 9.4 show a dash, because the value was not recorded.
- Currency conversion: now keeps the reviewer's comment count for 365 days. After that the value is removed, and the entry shows when it was removed.
- Invoice matching: now names the user who last changed the original currency. It does not show why the change was made, because the product does not ask for a reason.
- User roles: now accepts the payment terms as a filter. The filter applies to records created on or after 9.4.

- Fixed: exports did not keep the sort order after a page reload. It now keeps it until the user changes it.
- Fixed: the approval queue did not keep the sort order after a page reload. It now keeps it until the user changes it.

Known issue: attachment storage does not yet support files larger than 25 MB. Larger files are refused with a message that names the limit.

## Quillstone Ledger 9.5

Quillstone Ledger 9.5 was released on 1 October 2020 for customers on the hosted plan.

What changed:

- The mobile view: can now be turned off for each user. When it is turned off, the original currency is not collected for that user from then on.
- The search index: now keeps the import job number for 30 days. After that the value is removed, and the entry shows when it was removed.
- Invoice matching: now includes a column for the import job number. The column is empty for records created before 9.5, because the import job number was not stored then.

- Fixed: a date typed as day/month could be read as month/day in scheduled reports. The format now follows the account's settings.

Known issue: exports does not yet support files larger than 100 MB. Larger files are refused with a message that names the limit.

## Quillstone Ledger 9.6

Quillstone Ledger 9.6 was released on 8 October 2020 for customers on the hosted plan.

What changed:

- Attachment storage: now includes a column for the batch number. The column is empty for records created before 9.6, because the batch number was not stored then.
- Account settings: now includes a column for the device type. The column is empty for records created before 9.6, because the device type was not stored then.
- Attachment storage: now shows the import job number beside each entry. Entries made before 9.6 show a dash, because the value was not recorded.

- Fixed: in currency conversion, a month with no records showed a blank page. It now shows an empty table.

## Quillstone Ledger 9.7

Quillstone Ledger 9.7 was released on 15 October 2020 for customers on the hosted plan.

What changed:

- Invoice matching: now includes a column for the original currency. The column is empty for records created before 9.7, because the original currency was not stored then.
- The search index: now keeps the attachment count for 90 days. After that the value is removed, and the entry shows when it was removed.
- Bank feeds: faster on accounts with many records. In our own test on 15 October 2020, a 54,000-row run took 88 seconds on 9.6 and 29 seconds on 9.7. Your results depend on your data.
- Exports: now accepts the reviewer's comment count as a filter. The filter applies to records created on or after 9.7.
- Attachment storage: faster on accounts with many records. In our own test on 15 October 2020, a 152,000-row run took 49 seconds on 9.6 and 12 seconds on 9.7. Your results depend on your data.

- Fixed: the activity log did not keep the sort order after a page reload. It now keeps it until the user changes it.

## Quillstone Ledger 9.8

Quillstone Ledger 9.8 was released on 22 October 2020 for customers on the hosted plan.

What changed:

- Scheduled reports: can now be turned off for each user. When it is turned off, the payment terms is not collected for that user from then on.
- Attachment storage: now shows the time zone of the entry beside each entry. Entries made before 9.8 show a dash, because the value was not recorded.
- The activity log: can now be turned off for each user. When it is turned off, the payment terms is not collected for that user from then on.

- Fixed: in the search index, a month with no records showed a blank page. It now shows an empty table.

Known issue: scheduled reports does not yet support files larger than 50 MB. Larger files are refused with a message that names the limit.

Planned, not yet available: the original currency in invoice matching. We expect to release it in the second quarter of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 9.9

Quillstone Ledger 9.9 was released on 1 November 2020 for customers on the hosted plan.

What changed:

- Exports: now names the user who last changed the batch number. It does not show why the change was made, because the product does not ask for a reason.
- The search index: now accepts the original currency as a filter. The filter applies to records created on or after 9.9.
- Exports: can now be turned off for each user. When it is turned off, the exchange rate used is not collected for that user from then on.
- The approval queue: now includes a column for the approving user's role. The column is empty for records created before 9.9, because the approving user's role was not stored then.
- The audit view: can now be turned off for each user. When it is turned off, the payment terms is not collected for that user from then on.

- Fixed: totals in scheduled reports were rounded before they were added, which could change the last digit. They are now added first and rounded once.
- Fixed: in bank feeds, a month with no records showed a blank page. It now shows an empty table.

Known issue: currency conversion does not yet support files larger than 25 MB. Larger files are refused with a message that names the limit.

## Quillstone Ledger 10.0

Quillstone Ledger 10.0 was released on 8 November 2020 for customers on the hosted plan.

What changed:

- User roles: now names the user who last changed the exchange rate used. It does not show why the change was made, because the product does not ask for a reason.
- The import screen: can now be turned off for each user. When it is turned off, the approving user's role is not collected for that user from then on.
- Exports: now shows the device type beside each entry. Entries made before 10.0 show a dash, because the value was not recorded.
- The search index: now shows the cost centre beside each entry. Entries made before 10.0 show a dash, because the value was not recorded.

- Fixed: in the tax summary, a month with no records showed a blank page. It now shows an empty table.

## Quillstone Ledger 10.1

Quillstone Ledger 10.1 was released on 15 November 2020 for customers on the hosted plan.

What changed:

- The audit view: now shows the reviewer's comment count beside each entry. Entries made before 10.1 show a dash, because the value was not recorded.
- The tax summary: now keeps the original currency for 180 days. After that the value is removed, and the entry shows when it was removed.
- Attachment storage: now keeps the source file name for 90 days. After that the value is removed, and the entry shows when it was removed.
- The tax summary: now keeps the batch number for 90 days. After that the value is removed, and the entry shows when it was removed.
- User roles: now shows the reviewer's comment count beside each entry. Entries made before 10.1 show a dash, because the value was not recorded.

- Fixed: a date typed as day/month could be read as month/day in the audit view. The format now follows the account's settings.
- Fixed: in the activity log, a month with no records showed a blank page. It now shows an empty table.

## Quillstone Ledger 10.2

Quillstone Ledger 10.2 was released on 22 November 2020 for customers on the hosted plan.

What changed:

- Scheduled reports: can now be turned off for each user. When it is turned off, the batch number is not collected for that user from then on.
- Scheduled reports: now names the user who last changed the reviewer's comment count. It does not show why the change was made, because the product does not ask for a reason.
- The import screen: now accepts the exchange rate used as a filter. The filter applies to records created on or after 10.2.

- Fixed: totals in invoice matching were rounded before they were added, which could change the last digit. They are now added first and rounded once.

## Quillstone Ledger 10.3

Quillstone Ledger 10.3 was released on 1 December 2020 for customers on the hosted plan.

What changed:

- The tax summary: now includes a column for the time zone of the entry. The column is empty for records created before 10.3, because the time zone of the entry was not stored then.
- Bank feeds: now names the user who last changed the original currency. It does not show why the change was made, because the product does not ask for a reason.
- The audit view: can now be turned off for each user. When it is turned off, the import job number is not collected for that user from then on.

- Fixed: account settings did not keep the sort order after a page reload. It now keeps it until the user changes it.
- Fixed: a date typed as day/month could be read as month/day in user roles. The format now follows the account's settings.

Known issue: the mobile view does not yet support files larger than 25 MB. Larger files are refused with a message that names the limit.

## Quillstone Ledger 10.4

Quillstone Ledger 10.4 was released on 8 December 2020 for customers on the hosted plan.

What changed:

- The activity log: faster on accounts with many records. In our own test on 8 December 2020, a 102,000-row run took 78 seconds on 10.3 and 39 seconds on 10.4. Your results depend on your data.
- The tax summary: faster on accounts with many records. In our own test on 8 December 2020, a 72,000-row run took 60 seconds on 10.3 and 20 seconds on 10.4. Your results depend on your data.
- Scheduled reports: now accepts the reviewer's comment count as a filter. The filter applies to records created on or after 10.4.
- Invoice matching: now shows the original currency beside each entry. Entries made before 10.4 show a dash, because the value was not recorded.

- Fixed: invoice matching did not keep the sort order after a page reload. It now keeps it until the user changes it.
- Fixed: totals in bank feeds were rounded before they were added, which could change the last digit. They are now added first and rounded once.

Known issue: invoice matching does not yet support files larger than 25 MB. Larger files are refused with a message that names the limit.

## Quillstone Ledger 10.5

Quillstone Ledger 10.5 was released on 15 December 2020 for customers on the hosted plan.

What changed:

- Currency conversion: now keeps the batch number for 90 days. After that the value is removed, and the entry shows when it was removed.
- The audit view: faster on accounts with many records. In our own test on 15 December 2020, a 59,000-row run took 68 seconds on 10.4 and 13 seconds on 10.5. Your results depend on your data.
- The import screen: now names the user who last changed the import job number. It does not show why the change was made, because the product does not ask for a reason.

- Fixed: attachment storage did not keep the sort order after a page reload. It now keeps it until the user changes it.
- Fixed: totals in the search index were rounded before they were added, which could change the last digit. They are now added first and rounded once.

Known issue: attachment storage does not yet support files larger than 25 MB. Larger files are refused with a message that names the limit.

Planned, not yet available: the source file name in account settings. We expect to release it in the second half of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 10.6

Quillstone Ledger 10.6 was released on 22 December 2020 for customers on the hosted plan.

What changed:

- Invoice matching: faster on accounts with many records. In our own test on 22 December 2020, a 116,000-row run took 60 seconds on 10.5 and 12 seconds on 10.6. Your results depend on your data.
- The tax summary: now keeps the source file name for 365 days. After that the value is removed, and the entry shows when it was removed.
- The mobile view: now accepts the original currency as a filter. The filter applies to records created on or after 10.6.
- The tax summary: now shows the attachment count beside each entry. Entries made before 10.6 show a dash, because the value was not recorded.

- Fixed: totals in the search index were rounded before they were added, which could change the last digit. They are now added first and rounded once.
- Fixed: in the import screen, a month with no records showed a blank page. It now shows an empty table.

## Quillstone Ledger 10.7

Quillstone Ledger 10.7 was released on 1 January 2021 for customers on the hosted plan.

What changed:

- The tax summary: now shows the original currency beside each entry. Entries made before 10.7 show a dash, because the value was not recorded.
- Scheduled reports: now keeps the time zone of the entry for 30 days. After that the value is removed, and the entry shows when it was removed.
- The approval queue: now accepts the attachment count as a filter. The filter applies to records created on or after 10.7.

- Fixed: totals in the mobile view were rounded before they were added, which could change the last digit. They are now added first and rounded once.

Known issue: on slow connections the import screen can take up to 10 seconds to load. We are working on it and will say here when it is fixed.

Planned, not yet available: the payment terms in attachment storage. We expect to release it in the second quarter of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 10.8

Quillstone Ledger 10.8 was released on 8 January 2021 for customers on the hosted plan.

What changed:

- The approval queue: faster on accounts with many records. In our own test on 8 January 2021, a 70,000-row run took 25 seconds on 10.7 and 8 seconds on 10.8. Your results depend on your data.
- The approval queue: now names the user who last changed the time zone of the entry. It does not show why the change was made, because the product does not ask for a reason.
- Account settings: now keeps the attachment count for 180 days. After that the value is removed, and the entry shows when it was removed.
- Exports: can now be turned off for each user. When it is turned off, the original currency is not collected for that user from then on.

- Fixed: totals in the import screen were rounded before they were added, which could change the last digit. They are now added first and rounded once.
- Fixed: a date typed as day/month could be read as month/day in invoice matching. The format now follows the account's settings.

## Quillstone Ledger 10.9

Quillstone Ledger 10.9 was released on 15 January 2021 for customers on the hosted plan.

What changed:

- Bank feeds: faster on accounts with many records. In our own test on 15 January 2021, a 103,000-row run took 88 seconds on 10.8 and 17 seconds on 10.9. Your results depend on your data.
- The approval queue: now includes a column for the approving user's role. The column is empty for records created before 10.9, because the approving user's role was not stored then.
- Currency conversion: now accepts the exchange rate used as a filter. The filter applies to records created on or after 10.9.
- Exports: now shows the batch number beside each entry. Entries made before 10.9 show a dash, because the value was not recorded.

- Fixed: in currency conversion, a month with no records showed a blank page. It now shows an empty table.
- Fixed: totals in the approval queue were rounded before they were added, which could change the last digit. They are now added first and rounded once.

Known issue: on slow connections the activity log can take up to 10 seconds to load. We are working on it and will say here when it is fixed.

Planned, not yet available: the batch number in scheduled reports. We expect to release it in the first quarter of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 11.0

Quillstone Ledger 11.0 was released on 22 January 2021 for customers on the hosted plan.

What changed:

- The audit view: now accepts the approving user's role as a filter. The filter applies to records created on or after 11.0.
- Scheduled reports: now includes a column for the import job number. The column is empty for records created before 11.0, because the import job number was not stored then.
- The import screen: now keeps the import job number for 30 days. After that the value is removed, and the entry shows when it was removed.

- Fixed: a date typed as day/month could be read as month/day in invoice matching. The format now follows the account's settings.
- Fixed: a date typed as day/month could be read as month/day in the approval queue. The format now follows the account's settings.

Known issue: account settings does not yet support files larger than 50 MB. Larger files are refused with a message that names the limit.

Planned, not yet available: the original currency in invoice matching. We expect to release it in the second half of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 11.1

Quillstone Ledger 11.1 was released on 1 February 2021 for customers on the hosted plan.

What changed:

- Attachment storage: can now be turned off for each user. When it is turned off, the import job number is not collected for that user from then on.
- The audit view: faster on accounts with many records. In our own test on 1 February 2021, a 172,000-row run took 30 seconds on 11.0 and 7 seconds on 11.1. Your results depend on your data.
- Exports: now includes a column for the batch number. The column is empty for records created before 11.1, because the batch number was not stored then.

- Fixed: totals in the search index were rounded before they were added, which could change the last digit. They are now added first and rounded once.

## Quillstone Ledger 11.2

Quillstone Ledger 11.2 was released on 8 February 2021 for customers on the hosted plan.

What changed:

- The search index: now includes a column for the payment terms. The column is empty for records created before 11.2, because the payment terms was not stored then.
- Bank feeds: now includes a column for the time zone of the entry. The column is empty for records created before 11.2, because the time zone of the entry was not stored then.
- The approval queue: now includes a column for the reviewer's comment count. The column is empty for records created before 11.2, because the reviewer's comment count was not stored then.

- Fixed: a date typed as day/month could be read as month/day in the mobile view. The format now follows the account's settings.

Known issue: on slow connections report generation can take up to 30 seconds to load. We are working on it and will say here when it is fixed.

## Quillstone Ledger 11.3

Quillstone Ledger 11.3 was released on 15 February 2021 for customers on the hosted plan.

What changed:

- Exports: now names the user who last changed the payment terms. It does not show why the change was made, because the product does not ask for a reason.
- Attachment storage: now names the user who last changed the approving user's role. It does not show why the change was made, because the product does not ask for a reason.
- Report generation: now shows the cost centre beside each entry. Entries made before 11.3 show a dash, because the value was not recorded.
- Account settings: now shows the original currency beside each entry. Entries made before 11.3 show a dash, because the value was not recorded.

- Fixed: the tax summary did not keep the sort order after a page reload. It now keeps it until the user changes it.

Planned, not yet available: the original currency in the search index. We expect to release it in the second half of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 11.4

Quillstone Ledger 11.4 was released on 22 February 2021 for customers on the hosted plan.

What changed:

- Attachment storage: faster on accounts with many records. In our own test on 22 February 2021, a 187,000-row run took 23 seconds on 11.3 and 5 seconds on 11.4. Your results depend on your data.
- Currency conversion: faster on accounts with many records. In our own test on 22 February 2021, a 32,000-row run took 74 seconds on 11.3 and 24 seconds on 11.4. Your results depend on your data.
- Currency conversion: now includes a column for the payment terms. The column is empty for records created before 11.4, because the payment terms was not stored then.
- Currency conversion: faster on accounts with many records. In our own test on 22 February 2021, a 39,000-row run took 32 seconds on 11.3 and 10 seconds on 11.4. Your results depend on your data.
- The search index: now includes a column for the time zone of the entry. The column is empty for records created before 11.4, because the time zone of the entry was not stored then.

- Fixed: totals in the tax summary were rounded before they were added, which could change the last digit. They are now added first and rounded once.

Known issue: on slow connections the tax summary can take up to 10 seconds to load. We are working on it and will say here when it is fixed.

Planned, not yet available: the time zone of the entry in bank feeds. We expect to release it in the second quarter of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 11.5

Quillstone Ledger 11.5 was released on 1 March 2021 for customers on the hosted plan.

What changed:

- Exports: can now be turned off for each user. When it is turned off, the import job number is not collected for that user from then on.
- Report generation: now names the user who last changed the attachment count. It does not show why the change was made, because the product does not ask for a reason.
- Scheduled reports: now shows the cost centre beside each entry. Entries made before 11.5 show a dash, because the value was not recorded.
- The tax summary: can now be turned off for each user. When it is turned off, the cost centre is not collected for that user from then on.
- The audit view: now includes a column for the exchange rate used. The column is empty for records created before 11.5, because the exchange rate used was not stored then.

- Fixed: the audit view did not keep the sort order after a page reload. It now keeps it until the user changes it.

Known issue: the activity log does not yet support files larger than 25 MB. Larger files are refused with a message that names the limit.

Planned, not yet available: the attachment count in attachment storage. We expect to release it in the second half of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 11.6

Quillstone Ledger 11.6 was released on 8 March 2021 for customers on the hosted plan.

What changed:

- Currency conversion: now shows the approving user's role beside each entry. Entries made before 11.6 show a dash, because the value was not recorded.
- The import screen: now names the user who last changed the device type. It does not show why the change was made, because the product does not ask for a reason.
- Currency conversion: now keeps the source file name for 90 days. After that the value is removed, and the entry shows when it was removed.

- Fixed: totals in bank feeds were rounded before they were added, which could change the last digit. They are now added first and rounded once.

Known issue: on slow connections attachment storage can take up to 10 seconds to load. We are working on it and will say here when it is fixed.

Planned, not yet available: the original currency in the search index. We expect to release it in the second quarter of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 11.7

Quillstone Ledger 11.7 was released on 15 March 2021 for customers on the hosted plan.

What changed:

- Attachment storage: now includes a column for the import job number. The column is empty for records created before 11.7, because the import job number was not stored then.
- The mobile view: now includes a column for the attachment count. The column is empty for records created before 11.7, because the attachment count was not stored then.
- Invoice matching: can now be turned off for each user. When it is turned off, the batch number is not collected for that user from then on.
- The search index: faster on accounts with many records. In our own test on 15 March 2021, a 11,000-row run took 20 seconds on 11.6 and 5 seconds on 11.7. Your results depend on your data.
- Report generation: now accepts the attachment count as a filter. The filter applies to records created on or after 11.7.

- Fixed: a date typed as day/month could be read as month/day in the search index. The format now follows the account's settings.

Known issue: on slow connections the mobile view can take up to 30 seconds to load. We are working on it and will say here when it is fixed.

## Quillstone Ledger 11.8

Quillstone Ledger 11.8 was released on 22 March 2021 for customers on the hosted plan.

What changed:

- The tax summary: can now be turned off for each user. When it is turned off, the time zone of the entry is not collected for that user from then on.
- Invoice matching: can now be turned off for each user. When it is turned off, the cost centre is not collected for that user from then on.
- Invoice matching: faster on accounts with many records. In our own test on 22 March 2021, a 40,000-row run took 52 seconds on 11.7 and 13 seconds on 11.8. Your results depend on your data.
- The import screen: now accepts the attachment count as a filter. The filter applies to records created on or after 11.8.
- Bank feeds: now shows the source file name beside each entry. Entries made before 11.8 show a dash, because the value was not recorded.

- Fixed: the tax summary did not keep the sort order after a page reload. It now keeps it until the user changes it.

## Quillstone Ledger 11.9

Quillstone Ledger 11.9 was released on 1 April 2021 for customers on the hosted plan.

What changed:

- Account settings: now accepts the original currency as a filter. The filter applies to records created on or after 11.9.
- Attachment storage: now shows the device type beside each entry. Entries made before 11.9 show a dash, because the value was not recorded.
- Bank feeds: now names the user who last changed the cost centre. It does not show why the change was made, because the product does not ask for a reason.
- The mobile view: can now be turned off for each user. When it is turned off, the import job number is not collected for that user from then on.
- User roles: can now be turned off for each user. When it is turned off, the reviewer's comment count is not collected for that user from then on.

- Fixed: in report generation, a month with no records showed a blank page. It now shows an empty table.

## Quillstone Ledger 12.0

Quillstone Ledger 12.0 was released on 8 April 2021 for customers on the hosted plan.

What changed:

- Currency conversion: can now be turned off for each user. When it is turned off, the approving user's role is not collected for that user from then on.
- The tax summary: now names the user who last changed the source file name. It does not show why the change was made, because the product does not ask for a reason.
- User roles: now accepts the device type as a filter. The filter applies to records created on or after 12.0.

- Fixed: in currency conversion, a month with no records showed a blank page. It now shows an empty table.
- Fixed: totals in the audit view were rounded before they were added, which could change the last digit. They are now added first and rounded once.

Known issue: scheduled reports does not yet support files larger than 25 MB. Larger files are refused with a message that names the limit.

Planned, not yet available: the cost centre in the activity log. We expect to release it in the second quarter of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 12.1

Quillstone Ledger 12.1 was released on 15 April 2021 for customers on the hosted plan.

What changed:

- User roles: now includes a column for the reviewer's comment count. The column is empty for records created before 12.1, because the reviewer's comment count was not stored then.
- The mobile view: now includes a column for the attachment count. The column is empty for records created before 12.1, because the attachment count was not stored then.
- Bank feeds: now names the user who last changed the approving user's role. It does not show why the change was made, because the product does not ask for a reason.
- Currency conversion: can now be turned off for each user. When it is turned off, the cost centre is not collected for that user from then on.
- Currency conversion: now shows the payment terms beside each entry. Entries made before 12.1 show a dash, because the value was not recorded.

- Fixed: exports did not keep the sort order after a page reload. It now keeps it until the user changes it.
- Fixed: totals in currency conversion were rounded before they were added, which could change the last digit. They are now added first and rounded once.

Planned, not yet available: the import job number in invoice matching. We expect to release it in the second half of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 12.2

Quillstone Ledger 12.2 was released on 22 April 2021 for customers on the hosted plan.

What changed:

- Invoice matching: now names the user who last changed the device type. It does not show why the change was made, because the product does not ask for a reason.
- User roles: now names the user who last changed the device type. It does not show why the change was made, because the product does not ask for a reason.
- The activity log: now accepts the reviewer's comment count as a filter. The filter applies to records created on or after 12.2.
- Report generation: faster on accounts with many records. In our own test on 22 April 2021, a 169,000-row run took 89 seconds on 12.1 and 22 seconds on 12.2. Your results depend on your data.

- Fixed: in attachment storage, a month with no records showed a blank page. It now shows an empty table.
- Fixed: a date typed as day/month could be read as month/day in the import screen. The format now follows the account's settings.

Known issue: the search index does not yet support files larger than 25 MB. Larger files are refused with a message that names the limit.

## Quillstone Ledger 12.3

Quillstone Ledger 12.3 was released on 1 May 2021 for customers on the hosted plan.

What changed:

- Attachment storage: now accepts the attachment count as a filter. The filter applies to records created on or after 12.3.
- The approval queue: can now be turned off for each user. When it is turned off, the batch number is not collected for that user from then on.
- Bank feeds: faster on accounts with many records. In our own test on 1 May 2021, a 72,000-row run took 84 seconds on 12.2 and 28 seconds on 12.3. Your results depend on your data.

- Fixed: totals in user roles were rounded before they were added, which could change the last digit. They are now added first and rounded once.
- Fixed: totals in account settings were rounded before they were added, which could change the last digit. They are now added first and rounded once.

Known issue: scheduled reports does not yet support files larger than 100 MB. Larger files are refused with a message that names the limit.

## Quillstone Ledger 12.4

Quillstone Ledger 12.4 was released on 8 May 2021 for customers on the hosted plan.

What changed:

- Attachment storage: now accepts the device type as a filter. The filter applies to records created on or after 12.4.
- Invoice matching: can now be turned off for each user. When it is turned off, the exchange rate used is not collected for that user from then on.
- Currency conversion: now includes a column for the device type. The column is empty for records created before 12.4, because the device type was not stored then.

- Fixed: a date typed as day/month could be read as month/day in exports. The format now follows the account's settings.
- Fixed: exports did not keep the sort order after a page reload. It now keeps it until the user changes it.

Planned, not yet available: the exchange rate used in report generation. We expect to release it in the second quarter of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 12.5

Quillstone Ledger 12.5 was released on 15 May 2021 for customers on the hosted plan.

What changed:

- The tax summary: now keeps the device type for 90 days. After that the value is removed, and the entry shows when it was removed.
- The activity log: now includes a column for the batch number. The column is empty for records created before 12.5, because the batch number was not stored then.
- Bank feeds: faster on accounts with many records. In our own test on 15 May 2021, a 134,000-row run took 24 seconds on 12.4 and 8 seconds on 12.5. Your results depend on your data.
- The search index: now shows the original currency beside each entry. Entries made before 12.5 show a dash, because the value was not recorded.

- Fixed: totals in exports were rounded before they were added, which could change the last digit. They are now added first and rounded once.
- Fixed: in account settings, a month with no records showed a blank page. It now shows an empty table.

## Quillstone Ledger 12.6

Quillstone Ledger 12.6 was released on 22 May 2021 for customers on the hosted plan.

What changed:

- Scheduled reports: now names the user who last changed the attachment count. It does not show why the change was made, because the product does not ask for a reason.
- The audit view: faster on accounts with many records. In our own test on 22 May 2021, a 160,000-row run took 71 seconds on 12.5 and 35 seconds on 12.6. Your results depend on your data.
- Attachment storage: can now be turned off for each user. When it is turned off, the approving user's role is not collected for that user from then on.

- Fixed: the activity log did not keep the sort order after a page reload. It now keeps it until the user changes it.
- Fixed: totals in the import screen were rounded before they were added, which could change the last digit. They are now added first and rounded once.

## Quillstone Ledger 12.7

Quillstone Ledger 12.7 was released on 1 June 2021 for customers on the hosted plan.

What changed:

- Exports: can now be turned off for each user. When it is turned off, the approving user's role is not collected for that user from then on.
- Exports: faster on accounts with many records. In our own test on 1 June 2021, a 127,000-row run took 20 seconds on 12.6 and 4 seconds on 12.7. Your results depend on your data.
- User roles: now shows the source file name beside each entry. Entries made before 12.7 show a dash, because the value was not recorded.
- The mobile view: now shows the attachment count beside each entry. Entries made before 12.7 show a dash, because the value was not recorded.
- Currency conversion: can now be turned off for each user. When it is turned off, the device type is not collected for that user from then on.

- Fixed: a date typed as day/month could be read as month/day in invoice matching. The format now follows the account's settings.
- Fixed: a date typed as day/month could be read as month/day in currency conversion. The format now follows the account's settings.

Known issue: on slow connections the mobile view can take up to 30 seconds to load. We are working on it and will say here when it is fixed.

## Quillstone Ledger 12.8

Quillstone Ledger 12.8 was released on 8 June 2021 for customers on the hosted plan.

What changed:

- The import screen: now names the user who last changed the time zone of the entry. It does not show why the change was made, because the product does not ask for a reason.
- Account settings: now shows the reviewer's comment count beside each entry. Entries made before 12.8 show a dash, because the value was not recorded.
- The mobile view: now names the user who last changed the reviewer's comment count. It does not show why the change was made, because the product does not ask for a reason.
- Account settings: now names the user who last changed the batch number. It does not show why the change was made, because the product does not ask for a reason.
- Attachment storage: now keeps the device type for 180 days. After that the value is removed, and the entry shows when it was removed.

- Fixed: a date typed as day/month could be read as month/day in bank feeds. The format now follows the account's settings.

Known issue: the mobile view does not yet support files larger than 25 MB. Larger files are refused with a message that names the limit.

## Quillstone Ledger 12.9

Quillstone Ledger 12.9 was released on 15 June 2021 for customers on the hosted plan.

What changed:

- The mobile view: faster on accounts with many records. In our own test on 15 June 2021, a 191,000-row run took 47 seconds on 12.8 and 9 seconds on 12.9. Your results depend on your data.
- Report generation: can now be turned off for each user. When it is turned off, the original currency is not collected for that user from then on.
- Invoice matching: now keeps the exchange rate used for 30 days. After that the value is removed, and the entry shows when it was removed.
- The audit view: now accepts the import job number as a filter. The filter applies to records created on or after 12.9.
- Account settings: now accepts the reviewer's comment count as a filter. The filter applies to records created on or after 12.9.

- Fixed: the import screen did not keep the sort order after a page reload. It now keeps it until the user changes it.

## Quillstone Ledger 13.0

Quillstone Ledger 13.0 was released on 22 June 2021 for customers on the hosted plan.

What changed:

- The import screen: now includes a column for the time zone of the entry. The column is empty for records created before 13.0, because the time zone of the entry was not stored then.
- The approval queue: now names the user who last changed the cost centre. It does not show why the change was made, because the product does not ask for a reason.
- User roles: can now be turned off for each user. When it is turned off, the time zone of the entry is not collected for that user from then on.
- The tax summary: now names the user who last changed the exchange rate used. It does not show why the change was made, because the product does not ask for a reason.

- Fixed: totals in currency conversion were rounded before they were added, which could change the last digit. They are now added first and rounded once.

Known issue: user roles does not yet support files larger than 50 MB. Larger files are refused with a message that names the limit.

Planned, not yet available: the batch number in the import screen. We expect to release it in the second half of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 13.1

Quillstone Ledger 13.1 was released on 1 July 2021 for customers on the hosted plan.

What changed:

- The search index: faster on accounts with many records. In our own test on 1 July 2021, a 29,000-row run took 49 seconds on 13.0 and 16 seconds on 13.1. Your results depend on your data.
- The activity log: faster on accounts with many records. In our own test on 1 July 2021, a 175,000-row run took 78 seconds on 13.0 and 39 seconds on 13.1. Your results depend on your data.
- The import screen: can now be turned off for each user. When it is turned off, the batch number is not collected for that user from then on.
- Invoice matching: faster on accounts with many records. In our own test on 1 July 2021, a 25,000-row run took 83 seconds on 13.0 and 27 seconds on 13.1. Your results depend on your data.

- Fixed: totals in the mobile view were rounded before they were added, which could change the last digit. They are now added first and rounded once.
- Fixed: a date typed as day/month could be read as month/day in currency conversion. The format now follows the account's settings.

Known issue: on slow connections user roles can take up to 30 seconds to load. We are working on it and will say here when it is fixed.

## Quillstone Ledger 13.2

Quillstone Ledger 13.2 was released on 8 July 2021 for customers on the hosted plan.

What changed:

- Scheduled reports: now keeps the source file name for 365 days. After that the value is removed, and the entry shows when it was removed.
- The activity log: faster on accounts with many records. In our own test on 8 July 2021, a 84,000-row run took 66 seconds on 13.1 and 22 seconds on 13.2. Your results depend on your data.
- The search index: now includes a column for the cost centre. The column is empty for records created before 13.2, because the cost centre was not stored then.

- Fixed: a date typed as day/month could be read as month/day in account settings. The format now follows the account's settings.
- Fixed: a date typed as day/month could be read as month/day in the tax summary. The format now follows the account's settings.

Known issue: the tax summary does not yet support files larger than 50 MB. Larger files are refused with a message that names the limit.

## Quillstone Ledger 13.3

Quillstone Ledger 13.3 was released on 15 July 2021 for customers on the hosted plan.

What changed:

- The approval queue: faster on accounts with many records. In our own test on 15 July 2021, a 134,000-row run took 62 seconds on 13.2 and 20 seconds on 13.3. Your results depend on your data.
- The mobile view: faster on accounts with many records. In our own test on 15 July 2021, a 38,000-row run took 57 seconds on 13.2 and 14 seconds on 13.3. Your results depend on your data.
- The tax summary: now names the user who last changed the approving user's role. It does not show why the change was made, because the product does not ask for a reason.

- Fixed: totals in scheduled reports were rounded before they were added, which could change the last digit. They are now added first and rounded once.
- Fixed: a date typed as day/month could be read as month/day in the tax summary. The format now follows the account's settings.

## Quillstone Ledger 13.4

Quillstone Ledger 13.4 was released on 22 July 2021 for customers on the hosted plan.

What changed:

- The audit view: now includes a column for the exchange rate used. The column is empty for records created before 13.4, because the exchange rate used was not stored then.
- The mobile view: now shows the exchange rate used beside each entry. Entries made before 13.4 show a dash, because the value was not recorded.
- The mobile view: now shows the source file name beside each entry. Entries made before 13.4 show a dash, because the value was not recorded.
- Invoice matching: now shows the payment terms beside each entry. Entries made before 13.4 show a dash, because the value was not recorded.
- The activity log: now shows the reviewer's comment count beside each entry. Entries made before 13.4 show a dash, because the value was not recorded.

- Fixed: in exports, a month with no records showed a blank page. It now shows an empty table.

Known issue: on slow connections invoice matching can take up to 30 seconds to load. We are working on it and will say here when it is fixed.

## Quillstone Ledger 13.5

Quillstone Ledger 13.5 was released on 1 August 2021 for customers on the hosted plan.

What changed:

- The tax summary: now accepts the batch number as a filter. The filter applies to records created on or after 13.5.
- The activity log: now names the user who last changed the reviewer's comment count. It does not show why the change was made, because the product does not ask for a reason.
- User roles: now includes a column for the device type. The column is empty for records created before 13.5, because the device type was not stored then.

- Fixed: in currency conversion, a month with no records showed a blank page. It now shows an empty table.

Planned, not yet available: the import job number in the tax summary. We expect to release it in the second quarter of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 13.6

Quillstone Ledger 13.6 was released on 8 August 2021 for customers on the hosted plan.

What changed:

- The activity log: faster on accounts with many records. In our own test on 8 August 2021, a 60,000-row run took 44 seconds on 13.5 and 11 seconds on 13.6. Your results depend on your data.
- Report generation: can now be turned off for each user. When it is turned off, the cost centre is not collected for that user from then on.
- Bank feeds: faster on accounts with many records. In our own test on 8 August 2021, a 142,000-row run took 24 seconds on 13.5 and 6 seconds on 13.6. Your results depend on your data.
- The activity log: now accepts the original currency as a filter. The filter applies to records created on or after 13.6.
- The search index: now includes a column for the time zone of the entry. The column is empty for records created before 13.6, because the time zone of the entry was not stored then.

- Fixed: a date typed as day/month could be read as month/day in scheduled reports. The format now follows the account's settings.
- Fixed: in user roles, a month with no records showed a blank page. It now shows an empty table.

## Quillstone Ledger 13.7

Quillstone Ledger 13.7 was released on 15 August 2021 for customers on the hosted plan.

What changed:

- Scheduled reports: faster on accounts with many records. In our own test on 15 August 2021, a 186,000-row run took 57 seconds on 13.6 and 14 seconds on 13.7. Your results depend on your data.
- Scheduled reports: now includes a column for the attachment count. The column is empty for records created before 13.7, because the attachment count was not stored then.
- The mobile view: now names the user who last changed the exchange rate used. It does not show why the change was made, because the product does not ask for a reason.

- Fixed: in exports, a month with no records showed a blank page. It now shows an empty table.
- Fixed: the mobile view did not keep the sort order after a page reload. It now keeps it until the user changes it.

Known issue: the search index does not yet support files larger than 25 MB. Larger files are refused with a message that names the limit.

## Quillstone Ledger 13.8

Quillstone Ledger 13.8 was released on 22 August 2021 for customers on the hosted plan.

What changed:

- The activity log: can now be turned off for each user. When it is turned off, the payment terms is not collected for that user from then on.
- The approval queue: faster on accounts with many records. In our own test on 22 August 2021, a 178,000-row run took 60 seconds on 13.7 and 15 seconds on 13.8. Your results depend on your data.
- Account settings: can now be turned off for each user. When it is turned off, the attachment count is not collected for that user from then on.
- Scheduled reports: faster on accounts with many records. In our own test on 22 August 2021, a 37,000-row run took 87 seconds on 13.7 and 43 seconds on 13.8. Your results depend on your data.
- Currency conversion: now keeps the attachment count for 90 days. After that the value is removed, and the entry shows when it was removed.

- Fixed: a date typed as day/month could be read as month/day in bank feeds. The format now follows the account's settings.
- Fixed: a date typed as day/month could be read as month/day in the approval queue. The format now follows the account's settings.

Planned, not yet available: the attachment count in report generation. We expect to release it in the second half of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 13.9

Quillstone Ledger 13.9 was released on 1 September 2021 for customers on the hosted plan.

What changed:

- Report generation: now keeps the time zone of the entry for 90 days. After that the value is removed, and the entry shows when it was removed.
- The audit view: now names the user who last changed the time zone of the entry. It does not show why the change was made, because the product does not ask for a reason.
- The tax summary: can now be turned off for each user. When it is turned off, the attachment count is not collected for that user from then on.

- Fixed: in the activity log, a month with no records showed a blank page. It now shows an empty table.
- Fixed: in the search index, a month with no records showed a blank page. It now shows an empty table.

Planned, not yet available: the time zone of the entry in the mobile view. We expect to release it in the first quarter of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 14.0

Quillstone Ledger 14.0 was released on 8 September 2021 for customers on the hosted plan.

What changed:

- Invoice matching: now accepts the source file name as a filter. The filter applies to records created on or after 14.0.
- Attachment storage: now names the user who last changed the reviewer's comment count. It does not show why the change was made, because the product does not ask for a reason.
- The import screen: now includes a column for the approving user's role. The column is empty for records created before 14.0, because the approving user's role was not stored then.

- Fixed: in user roles, a month with no records showed a blank page. It now shows an empty table.

## Quillstone Ledger 14.1

Quillstone Ledger 14.1 was released on 15 September 2021 for customers on the hosted plan.

What changed:

- The tax summary: now names the user who last changed the original currency. It does not show why the change was made, because the product does not ask for a reason.
- Account settings: now accepts the attachment count as a filter. The filter applies to records created on or after 14.1.
- Exports: faster on accounts with many records. In our own test on 15 September 2021, a 194,000-row run took 52 seconds on 14.0 and 26 seconds on 14.1. Your results depend on your data.
- Exports: can now be turned off for each user. When it is turned off, the device type is not collected for that user from then on.

- Fixed: invoice matching did not keep the sort order after a page reload. It now keeps it until the user changes it.

## Quillstone Ledger 14.2

Quillstone Ledger 14.2 was released on 22 September 2021 for customers on the hosted plan.

What changed:

- The mobile view: can now be turned off for each user. When it is turned off, the import job number is not collected for that user from then on.
- Exports: now includes a column for the batch number. The column is empty for records created before 14.2, because the batch number was not stored then.
- Invoice matching: now accepts the batch number as a filter. The filter applies to records created on or after 14.2.
- Invoice matching: now accepts the cost centre as a filter. The filter applies to records created on or after 14.2.

- Fixed: totals in bank feeds were rounded before they were added, which could change the last digit. They are now added first and rounded once.

## Quillstone Ledger 14.3

Quillstone Ledger 14.3 was released on 1 October 2021 for customers on the hosted plan.

What changed:

- Attachment storage: now shows the device type beside each entry. Entries made before 14.3 show a dash, because the value was not recorded.
- Invoice matching: now names the user who last changed the exchange rate used. It does not show why the change was made, because the product does not ask for a reason.
- Report generation: now shows the import job number beside each entry. Entries made before 14.3 show a dash, because the value was not recorded.
- The mobile view: faster on accounts with many records. In our own test on 1 October 2021, a 12,000-row run took 83 seconds on 14.2 and 16 seconds on 14.3. Your results depend on your data.
- Attachment storage: can now be turned off for each user. When it is turned off, the time zone of the entry is not collected for that user from then on.

- Fixed: the search index did not keep the sort order after a page reload. It now keeps it until the user changes it.
- Fixed: in exports, a month with no records showed a blank page. It now shows an empty table.

## Quillstone Ledger 14.4

Quillstone Ledger 14.4 was released on 8 October 2021 for customers on the hosted plan.

What changed:

- The mobile view: can now be turned off for each user. When it is turned off, the device type is not collected for that user from then on.
- Currency conversion: can now be turned off for each user. When it is turned off, the reviewer's comment count is not collected for that user from then on.
- Account settings: can now be turned off for each user. When it is turned off, the cost centre is not collected for that user from then on.
- The import screen: now accepts the approving user's role as a filter. The filter applies to records created on or after 14.4.

- Fixed: in the tax summary, a month with no records showed a blank page. It now shows an empty table.
- Fixed: invoice matching did not keep the sort order after a page reload. It now keeps it until the user changes it.

Known issue: on slow connections the audit view can take up to 20 seconds to load. We are working on it and will say here when it is fixed.

## Quillstone Ledger 14.5

Quillstone Ledger 14.5 was released on 15 October 2021 for customers on the hosted plan.

What changed:

- User roles: now includes a column for the attachment count. The column is empty for records created before 14.5, because the attachment count was not stored then.
- Exports: faster on accounts with many records. In our own test on 15 October 2021, a 113,000-row run took 55 seconds on 14.4 and 13 seconds on 14.5. Your results depend on your data.
- Exports: now keeps the import job number for 90 days. After that the value is removed, and the entry shows when it was removed.
- Report generation: now includes a column for the batch number. The column is empty for records created before 14.5, because the batch number was not stored then.

- Fixed: in account settings, a month with no records showed a blank page. It now shows an empty table.

Known issue: scheduled reports does not yet support files larger than 100 MB. Larger files are refused with a message that names the limit.

## Quillstone Ledger 14.6

Quillstone Ledger 14.6 was released on 22 October 2021 for customers on the hosted plan.

What changed:

- Report generation: now shows the import job number beside each entry. Entries made before 14.6 show a dash, because the value was not recorded.
- The audit view: now names the user who last changed the import job number. It does not show why the change was made, because the product does not ask for a reason.
- The activity log: now names the user who last changed the import job number. It does not show why the change was made, because the product does not ask for a reason.

- Fixed: exports did not keep the sort order after a page reload. It now keeps it until the user changes it.
- Fixed: totals in the import screen were rounded before they were added, which could change the last digit. They are now added first and rounded once.

Known issue: the approval queue does not yet support files larger than 50 MB. Larger files are refused with a message that names the limit.

## Quillstone Ledger 14.7

Quillstone Ledger 14.7 was released on 1 November 2021 for customers on the hosted plan.

What changed:

- Bank feeds: now includes a column for the exchange rate used. The column is empty for records created before 14.7, because the exchange rate used was not stored then.
- The import screen: now names the user who last changed the batch number. It does not show why the change was made, because the product does not ask for a reason.
- Report generation: now keeps the original currency for 365 days. After that the value is removed, and the entry shows when it was removed.

- Fixed: totals in the activity log were rounded before they were added, which could change the last digit. They are now added first and rounded once.

Known issue: on slow connections user roles can take up to 10 seconds to load. We are working on it and will say here when it is fixed.

## Quillstone Ledger 14.8

Quillstone Ledger 14.8 was released on 8 November 2021 for customers on the hosted plan.

What changed:

- The approval queue: now names the user who last changed the original currency. It does not show why the change was made, because the product does not ask for a reason.
- Exports: now keeps the time zone of the entry for 90 days. After that the value is removed, and the entry shows when it was removed.
- Bank feeds: now accepts the cost centre as a filter. The filter applies to records created on or after 14.8.
- Scheduled reports: now names the user who last changed the reviewer's comment count. It does not show why the change was made, because the product does not ask for a reason.

- Fixed: a date typed as day/month could be read as month/day in the search index. The format now follows the account's settings.

Known issue: on slow connections report generation can take up to 10 seconds to load. We are working on it and will say here when it is fixed.

Planned, not yet available: the time zone of the entry in invoice matching. We expect to release it in the first quarter of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 14.9

Quillstone Ledger 14.9 was released on 15 November 2021 for customers on the hosted plan.

What changed:

- Currency conversion: now accepts the approving user's role as a filter. The filter applies to records created on or after 14.9.
- Attachment storage: now names the user who last changed the reviewer's comment count. It does not show why the change was made, because the product does not ask for a reason.
- Bank feeds: now accepts the exchange rate used as a filter. The filter applies to records created on or after 14.9.
- Account settings: can now be turned off for each user. When it is turned off, the approving user's role is not collected for that user from then on.
- Scheduled reports: now names the user who last changed the reviewer's comment count. It does not show why the change was made, because the product does not ask for a reason.

- Fixed: a date typed as day/month could be read as month/day in currency conversion. The format now follows the account's settings.
- Fixed: scheduled reports did not keep the sort order after a page reload. It now keeps it until the user changes it.

## Quillstone Ledger 15.0

Quillstone Ledger 15.0 was released on 22 November 2021 for customers on the hosted plan.

What changed:

- The tax summary: now shows the source file name beside each entry. Entries made before 15.0 show a dash, because the value was not recorded.
- Currency conversion: now includes a column for the payment terms. The column is empty for records created before 15.0, because the payment terms was not stored then.
- Currency conversion: now names the user who last changed the original currency. It does not show why the change was made, because the product does not ask for a reason.
- User roles: now includes a column for the exchange rate used. The column is empty for records created before 15.0, because the exchange rate used was not stored then.

- Fixed: in currency conversion, a month with no records showed a blank page. It now shows an empty table.
- Fixed: scheduled reports did not keep the sort order after a page reload. It now keeps it until the user changes it.

Planned, not yet available: the source file name in the mobile view. We expect to release it in the second half of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 15.1

Quillstone Ledger 15.1 was released on 1 December 2021 for customers on the hosted plan.

What changed:

- The mobile view: can now be turned off for each user. When it is turned off, the time zone of the entry is not collected for that user from then on.
- The search index: can now be turned off for each user. When it is turned off, the device type is not collected for that user from then on.
- The mobile view: now shows the attachment count beside each entry. Entries made before 15.1 show a dash, because the value was not recorded.
- Account settings: now names the user who last changed the approving user's role. It does not show why the change was made, because the product does not ask for a reason.
- Attachment storage: now keeps the original currency for 180 days. After that the value is removed, and the entry shows when it was removed.

- Fixed: in user roles, a month with no records showed a blank page. It now shows an empty table.
- Fixed: a date typed as day/month could be read as month/day in attachment storage. The format now follows the account's settings.

Known issue: on slow connections bank feeds can take up to 10 seconds to load. We are working on it and will say here when it is fixed.

## Quillstone Ledger 15.2

Quillstone Ledger 15.2 was released on 8 December 2021 for customers on the hosted plan.

What changed:

- Scheduled reports: now names the user who last changed the cost centre. It does not show why the change was made, because the product does not ask for a reason.
- Account settings: can now be turned off for each user. When it is turned off, the original currency is not collected for that user from then on.
- The search index: now accepts the device type as a filter. The filter applies to records created on or after 15.2.
- The audit view: now accepts the time zone of the entry as a filter. The filter applies to records created on or after 15.2.

- Fixed: a date typed as day/month could be read as month/day in attachment storage. The format now follows the account's settings.

## Quillstone Ledger 15.3

Quillstone Ledger 15.3 was released on 15 December 2021 for customers on the hosted plan.

What changed:

- Account settings: faster on accounts with many records. In our own test on 15 December 2021, a 34,000-row run took 24 seconds on 15.2 and 12 seconds on 15.3. Your results depend on your data.
- The audit view: now keeps the batch number for 180 days. After that the value is removed, and the entry shows when it was removed.
- Bank feeds: faster on accounts with many records. In our own test on 15 December 2021, a 77,000-row run took 68 seconds on 15.2 and 13 seconds on 15.3. Your results depend on your data.

- Fixed: the mobile view did not keep the sort order after a page reload. It now keeps it until the user changes it.

Known issue: on slow connections exports can take up to 30 seconds to load. We are working on it and will say here when it is fixed.

## Quillstone Ledger 15.4

Quillstone Ledger 15.4 was released on 22 December 2021 for customers on the hosted plan.

What changed:

- Scheduled reports: can now be turned off for each user. When it is turned off, the attachment count is not collected for that user from then on.
- The mobile view: now shows the approving user's role beside each entry. Entries made before 15.4 show a dash, because the value was not recorded.
- Exports: now shows the device type beside each entry. Entries made before 15.4 show a dash, because the value was not recorded.
- User roles: now names the user who last changed the exchange rate used. It does not show why the change was made, because the product does not ask for a reason.
- The mobile view: can now be turned off for each user. When it is turned off, the device type is not collected for that user from then on.

- Fixed: in currency conversion, a month with no records showed a blank page. It now shows an empty table.
- Fixed: the import screen did not keep the sort order after a page reload. It now keeps it until the user changes it.

Known issue: the import screen does not yet support files larger than 25 MB. Larger files are refused with a message that names the limit.

## Quillstone Ledger 15.5

Quillstone Ledger 15.5 was released on 1 January 2022 for customers on the hosted plan.

What changed:

- Exports: can now be turned off for each user. When it is turned off, the attachment count is not collected for that user from then on.
- The activity log: now includes a column for the attachment count. The column is empty for records created before 15.5, because the attachment count was not stored then.
- Attachment storage: can now be turned off for each user. When it is turned off, the batch number is not collected for that user from then on.

- Fixed: bank feeds did not keep the sort order after a page reload. It now keeps it until the user changes it.
- Fixed: user roles did not keep the sort order after a page reload. It now keeps it until the user changes it.

Known issue: on slow connections invoice matching can take up to 20 seconds to load. We are working on it and will say here when it is fixed.

Planned, not yet available: the approving user's role in attachment storage. We expect to release it in the second quarter of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 15.6

Quillstone Ledger 15.6 was released on 8 January 2022 for customers on the hosted plan.

What changed:

- The tax summary: now keeps the reviewer's comment count for 365 days. After that the value is removed, and the entry shows when it was removed.
- Bank feeds: can now be turned off for each user. When it is turned off, the approving user's role is not collected for that user from then on.
- Account settings: now accepts the source file name as a filter. The filter applies to records created on or after 15.6.
- The approval queue: now names the user who last changed the approving user's role. It does not show why the change was made, because the product does not ask for a reason.
- The audit view: can now be turned off for each user. When it is turned off, the exchange rate used is not collected for that user from then on.

- Fixed: a date typed as day/month could be read as month/day in currency conversion. The format now follows the account's settings.

## Quillstone Ledger 15.7

Quillstone Ledger 15.7 was released on 15 January 2022 for customers on the hosted plan.

What changed:

- The import screen: now keeps the exchange rate used for 180 days. After that the value is removed, and the entry shows when it was removed.
- Currency conversion: faster on accounts with many records. In our own test on 15 January 2022, a 54,000-row run took 88 seconds on 15.6 and 29 seconds on 15.7. Your results depend on your data.
- The audit view: now names the user who last changed the original currency. It does not show why the change was made, because the product does not ask for a reason.
- Bank feeds: now names the user who last changed the exchange rate used. It does not show why the change was made, because the product does not ask for a reason.
- Invoice matching: now names the user who last changed the source file name. It does not show why the change was made, because the product does not ask for a reason.

- Fixed: the mobile view did not keep the sort order after a page reload. It now keeps it until the user changes it.
- Fixed: a date typed as day/month could be read as month/day in currency conversion. The format now follows the account's settings.

Known issue: account settings does not yet support files larger than 25 MB. Larger files are refused with a message that names the limit.

## Quillstone Ledger 15.8

Quillstone Ledger 15.8 was released on 22 January 2022 for customers on the hosted plan.

What changed:

- The approval queue: can now be turned off for each user. When it is turned off, the device type is not collected for that user from then on.
- The approval queue: now accepts the source file name as a filter. The filter applies to records created on or after 15.8.
- The mobile view: can now be turned off for each user. When it is turned off, the batch number is not collected for that user from then on.
- The mobile view: now accepts the batch number as a filter. The filter applies to records created on or after 15.8.
- Account settings: faster on accounts with many records. In our own test on 22 January 2022, a 71,000-row run took 40 seconds on 15.7 and 8 seconds on 15.8. Your results depend on your data.

- Fixed: scheduled reports did not keep the sort order after a page reload. It now keeps it until the user changes it.
- Fixed: a date typed as day/month could be read as month/day in bank feeds. The format now follows the account's settings.

Known issue: on slow connections the import screen can take up to 10 seconds to load. We are working on it and will say here when it is fixed.

Planned, not yet available: the time zone of the entry in the tax summary. We expect to release it in the second half of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 15.9

Quillstone Ledger 15.9 was released on 1 February 2022 for customers on the hosted plan.

What changed:

- The audit view: faster on accounts with many records. In our own test on 1 February 2022, a 12,000-row run took 72 seconds on 15.8 and 36 seconds on 15.9. Your results depend on your data.
- Bank feeds: now accepts the reviewer's comment count as a filter. The filter applies to records created on or after 15.9.
- The tax summary: now names the user who last changed the device type. It does not show why the change was made, because the product does not ask for a reason.
- The mobile view: now keeps the attachment count for 365 days. After that the value is removed, and the entry shows when it was removed.
- User roles: now shows the reviewer's comment count beside each entry. Entries made before 15.9 show a dash, because the value was not recorded.

- Fixed: in the approval queue, a month with no records showed a blank page. It now shows an empty table.
- Fixed: attachment storage did not keep the sort order after a page reload. It now keeps it until the user changes it.

Known issue: invoice matching does not yet support files larger than 25 MB. Larger files are refused with a message that names the limit.

## Quillstone Ledger 16.0

Quillstone Ledger 16.0 was released on 8 February 2022 for customers on the hosted plan.

What changed:

- User roles: can now be turned off for each user. When it is turned off, the import job number is not collected for that user from then on.
- Invoice matching: now names the user who last changed the source file name. It does not show why the change was made, because the product does not ask for a reason.
- The import screen: now names the user who last changed the approving user's role. It does not show why the change was made, because the product does not ask for a reason.
- The activity log: now keeps the payment terms for 365 days. After that the value is removed, and the entry shows when it was removed.
- The mobile view: now shows the source file name beside each entry. Entries made before 16.0 show a dash, because the value was not recorded.

- Fixed: totals in the import screen were rounded before they were added, which could change the last digit. They are now added first and rounded once.
- Fixed: totals in the tax summary were rounded before they were added, which could change the last digit. They are now added first and rounded once.

Known issue: on slow connections user roles can take up to 20 seconds to load. We are working on it and will say here when it is fixed.

Planned, not yet available: the device type in the mobile view. We expect to release it in the first quarter of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 16.1

Quillstone Ledger 16.1 was released on 15 February 2022 for customers on the hosted plan.

What changed:

- Bank feeds: now shows the approving user's role beside each entry. Entries made before 16.1 show a dash, because the value was not recorded.
- Report generation: now accepts the payment terms as a filter. The filter applies to records created on or after 16.1.
- Bank feeds: now names the user who last changed the source file name. It does not show why the change was made, because the product does not ask for a reason.
- Scheduled reports: now includes a column for the batch number. The column is empty for records created before 16.1, because the batch number was not stored then.

- Fixed: account settings did not keep the sort order after a page reload. It now keeps it until the user changes it.

Known issue: on slow connections attachment storage can take up to 30 seconds to load. We are working on it and will say here when it is fixed.

Planned, not yet available: the original currency in scheduled reports. We expect to release it in the second quarter of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 16.2

Quillstone Ledger 16.2 was released on 22 February 2022 for customers on the hosted plan.

What changed:

- The audit view: now keeps the approving user's role for 30 days. After that the value is removed, and the entry shows when it was removed.
- Invoice matching: now keeps the approving user's role for 90 days. After that the value is removed, and the entry shows when it was removed.
- Attachment storage: can now be turned off for each user. When it is turned off, the cost centre is not collected for that user from then on.
- Bank feeds: now keeps the payment terms for 30 days. After that the value is removed, and the entry shows when it was removed.

- Fixed: totals in scheduled reports were rounded before they were added, which could change the last digit. They are now added first and rounded once.

## Quillstone Ledger 16.3

Quillstone Ledger 16.3 was released on 1 March 2022 for customers on the hosted plan.

What changed:

- Currency conversion: now shows the import job number beside each entry. Entries made before 16.3 show a dash, because the value was not recorded.
- The activity log: now names the user who last changed the cost centre. It does not show why the change was made, because the product does not ask for a reason.
- Report generation: now keeps the approving user's role for 180 days. After that the value is removed, and the entry shows when it was removed.

- Fixed: totals in the import screen were rounded before they were added, which could change the last digit. They are now added first and rounded once.
- Fixed: totals in currency conversion were rounded before they were added, which could change the last digit. They are now added first and rounded once.

Known issue: on slow connections the tax summary can take up to 30 seconds to load. We are working on it and will say here when it is fixed.

Planned, not yet available: the approving user's role in currency conversion. We expect to release it in the second half of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 16.4

Quillstone Ledger 16.4 was released on 8 March 2022 for customers on the hosted plan.

What changed:

- Report generation: now names the user who last changed the original currency. It does not show why the change was made, because the product does not ask for a reason.
- User roles: now names the user who last changed the payment terms. It does not show why the change was made, because the product does not ask for a reason.
- The search index: faster on accounts with many records. In our own test on 8 March 2022, a 42,000-row run took 75 seconds on 16.3 and 25 seconds on 16.4. Your results depend on your data.
- The import screen: now includes a column for the reviewer's comment count. The column is empty for records created before 16.4, because the reviewer's comment count was not stored then.

- Fixed: report generation did not keep the sort order after a page reload. It now keeps it until the user changes it.
- Fixed: the search index did not keep the sort order after a page reload. It now keeps it until the user changes it.

Known issue: exports does not yet support files larger than 25 MB. Larger files are refused with a message that names the limit.

## Quillstone Ledger 16.5

Quillstone Ledger 16.5 was released on 15 March 2022 for customers on the hosted plan.

What changed:

- Attachment storage: now includes a column for the exchange rate used. The column is empty for records created before 16.5, because the exchange rate used was not stored then.
- The approval queue: faster on accounts with many records. In our own test on 15 March 2022, a 32,000-row run took 74 seconds on 16.4 and 24 seconds on 16.5. Your results depend on your data.
- The mobile view: now includes a column for the source file name. The column is empty for records created before 16.5, because the source file name was not stored then.

- Fixed: totals in invoice matching were rounded before they were added, which could change the last digit. They are now added first and rounded once.

Known issue: the approval queue does not yet support files larger than 100 MB. Larger files are refused with a message that names the limit.

Planned, not yet available: the approving user's role in report generation. We expect to release it in the second quarter of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 16.6

Quillstone Ledger 16.6 was released on 22 March 2022 for customers on the hosted plan.

What changed:

- Attachment storage: now shows the import job number beside each entry. Entries made before 16.6 show a dash, because the value was not recorded.
- Report generation: now shows the exchange rate used beside each entry. Entries made before 16.6 show a dash, because the value was not recorded.
- Attachment storage: now names the user who last changed the time zone of the entry. It does not show why the change was made, because the product does not ask for a reason.

- Fixed: totals in currency conversion were rounded before they were added, which could change the last digit. They are now added first and rounded once.
- Fixed: in invoice matching, a month with no records showed a blank page. It now shows an empty table.

Known issue: on slow connections bank feeds can take up to 20 seconds to load. We are working on it and will say here when it is fixed.

## Quillstone Ledger 16.7

Quillstone Ledger 16.7 was released on 1 April 2022 for customers on the hosted plan.

What changed:

- Scheduled reports: now keeps the approving user's role for 365 days. After that the value is removed, and the entry shows when it was removed.
- The tax summary: faster on accounts with many records. In our own test on 1 April 2022, a 33,000-row run took 65 seconds on 16.6 and 16 seconds on 16.7. Your results depend on your data.
- The import screen: now names the user who last changed the batch number. It does not show why the change was made, because the product does not ask for a reason.
- User roles: now names the user who last changed the import job number. It does not show why the change was made, because the product does not ask for a reason.

- Fixed: in report generation, a month with no records showed a blank page. It now shows an empty table.

Known issue: report generation does not yet support files larger than 25 MB. Larger files are refused with a message that names the limit.

## Quillstone Ledger 16.8

Quillstone Ledger 16.8 was released on 8 April 2022 for customers on the hosted plan.

What changed:

- Report generation: can now be turned off for each user. When it is turned off, the cost centre is not collected for that user from then on.
- Currency conversion: faster on accounts with many records. In our own test on 8 April 2022, a 19,000-row run took 44 seconds on 16.7 and 22 seconds on 16.8. Your results depend on your data.
- The activity log: can now be turned off for each user. When it is turned off, the payment terms is not collected for that user from then on.
- Bank feeds: now keeps the batch number for 30 days. After that the value is removed, and the entry shows when it was removed.

- Fixed: a date typed as day/month could be read as month/day in the import screen. The format now follows the account's settings.
- Fixed: totals in the activity log were rounded before they were added, which could change the last digit. They are now added first and rounded once.

Known issue: attachment storage does not yet support files larger than 50 MB. Larger files are refused with a message that names the limit.

Planned, not yet available: the batch number in the import screen. We expect to release it in the first quarter of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 16.9

Quillstone Ledger 16.9 was released on 15 April 2022 for customers on the hosted plan.

What changed:

- Report generation: now shows the cost centre beside each entry. Entries made before 16.9 show a dash, because the value was not recorded.
- Report generation: now names the user who last changed the batch number. It does not show why the change was made, because the product does not ask for a reason.
- The tax summary: now keeps the exchange rate used for 180 days. After that the value is removed, and the entry shows when it was removed.
- Currency conversion: faster on accounts with many records. In our own test on 15 April 2022, a 142,000-row run took 51 seconds on 16.8 and 10 seconds on 16.9. Your results depend on your data.

- Fixed: totals in the approval queue were rounded before they were added, which could change the last digit. They are now added first and rounded once.

## Quillstone Ledger 17.0

Quillstone Ledger 17.0 was released on 22 April 2022 for customers on the hosted plan.

What changed:

- User roles: now accepts the device type as a filter. The filter applies to records created on or after 17.0.
- The audit view: now names the user who last changed the batch number. It does not show why the change was made, because the product does not ask for a reason.
- Bank feeds: now shows the cost centre beside each entry. Entries made before 17.0 show a dash, because the value was not recorded.

- Fixed: totals in account settings were rounded before they were added, which could change the last digit. They are now added first and rounded once.

Known issue: the search index does not yet support files larger than 50 MB. Larger files are refused with a message that names the limit.

## Quillstone Ledger 17.1

Quillstone Ledger 17.1 was released on 1 May 2022 for customers on the hosted plan.

What changed:

- The tax summary: faster on accounts with many records. In our own test on 1 May 2022, a 47,000-row run took 86 seconds on 17.0 and 21 seconds on 17.1. Your results depend on your data.
- Currency conversion: faster on accounts with many records. In our own test on 1 May 2022, a 68,000-row run took 78 seconds on 17.0 and 19 seconds on 17.1. Your results depend on your data.
- Report generation: now shows the exchange rate used beside each entry. Entries made before 17.1 show a dash, because the value was not recorded.
- The search index: now includes a column for the batch number. The column is empty for records created before 17.1, because the batch number was not stored then.

- Fixed: a date typed as day/month could be read as month/day in scheduled reports. The format now follows the account's settings.
- Fixed: scheduled reports did not keep the sort order after a page reload. It now keeps it until the user changes it.

Known issue: currency conversion does not yet support files larger than 100 MB. Larger files are refused with a message that names the limit.

## Quillstone Ledger 17.2

Quillstone Ledger 17.2 was released on 8 May 2022 for customers on the hosted plan.

What changed:

- Currency conversion: now keeps the import job number for 180 days. After that the value is removed, and the entry shows when it was removed.
- The approval queue: now keeps the batch number for 365 days. After that the value is removed, and the entry shows when it was removed.
- The approval queue: now includes a column for the device type. The column is empty for records created before 17.2, because the device type was not stored then.
- Bank feeds: now accepts the reviewer's comment count as a filter. The filter applies to records created on or after 17.2.

- Fixed: user roles did not keep the sort order after a page reload. It now keeps it until the user changes it.

Known issue: bank feeds does not yet support files larger than 25 MB. Larger files are refused with a message that names the limit.

Planned, not yet available: the approving user's role in account settings. We expect to release it in the second half of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 17.3

Quillstone Ledger 17.3 was released on 15 May 2022 for customers on the hosted plan.

What changed:

- The search index: now includes a column for the device type. The column is empty for records created before 17.3, because the device type was not stored then.
- The tax summary: now keeps the batch number for 180 days. After that the value is removed, and the entry shows when it was removed.
- The mobile view: now shows the original currency beside each entry. Entries made before 17.3 show a dash, because the value was not recorded.

- Fixed: totals in exports were rounded before they were added, which could change the last digit. They are now added first and rounded once.

Planned, not yet available: the device type in bank feeds. We expect to release it in the second half of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 17.4

Quillstone Ledger 17.4 was released on 22 May 2022 for customers on the hosted plan.

What changed:

- The approval queue: now accepts the device type as a filter. The filter applies to records created on or after 17.4.
- The activity log: now names the user who last changed the device type. It does not show why the change was made, because the product does not ask for a reason.
- The search index: faster on accounts with many records. In our own test on 22 May 2022, a 187,000-row run took 84 seconds on 17.3 and 28 seconds on 17.4. Your results depend on your data.
- The activity log: can now be turned off for each user. When it is turned off, the original currency is not collected for that user from then on.
- Account settings: now includes a column for the payment terms. The column is empty for records created before 17.4, because the payment terms was not stored then.

- Fixed: attachment storage did not keep the sort order after a page reload. It now keeps it until the user changes it.
- Fixed: scheduled reports did not keep the sort order after a page reload. It now keeps it until the user changes it.

Known issue: account settings does not yet support files larger than 100 MB. Larger files are refused with a message that names the limit.

## Quillstone Ledger 17.5

Quillstone Ledger 17.5 was released on 1 June 2022 for customers on the hosted plan.

What changed:

- Currency conversion: now accepts the cost centre as a filter. The filter applies to records created on or after 17.5.
- Invoice matching: now keeps the batch number for 90 days. After that the value is removed, and the entry shows when it was removed.
- Bank feeds: now includes a column for the exchange rate used. The column is empty for records created before 17.5, because the exchange rate used was not stored then.

- Fixed: a date typed as day/month could be read as month/day in exports. The format now follows the account's settings.
- Fixed: in the mobile view, a month with no records showed a blank page. It now shows an empty table.

## Quillstone Ledger 17.6

Quillstone Ledger 17.6 was released on 8 June 2022 for customers on the hosted plan.

What changed:

- Invoice matching: now accepts the import job number as a filter. The filter applies to records created on or after 17.6.
- The mobile view: now shows the payment terms beside each entry. Entries made before 17.6 show a dash, because the value was not recorded.
- Bank feeds: can now be turned off for each user. When it is turned off, the original currency is not collected for that user from then on.
- Invoice matching: now includes a column for the attachment count. The column is empty for records created before 17.6, because the attachment count was not stored then.

- Fixed: in scheduled reports, a month with no records showed a blank page. It now shows an empty table.

## Quillstone Ledger 17.7

Quillstone Ledger 17.7 was released on 15 June 2022 for customers on the hosted plan.

What changed:

- Invoice matching: now names the user who last changed the source file name. It does not show why the change was made, because the product does not ask for a reason.
- The approval queue: now names the user who last changed the cost centre. It does not show why the change was made, because the product does not ask for a reason.
- Invoice matching: now names the user who last changed the import job number. It does not show why the change was made, because the product does not ask for a reason.
- The activity log: now keeps the payment terms for 365 days. After that the value is removed, and the entry shows when it was removed.

- Fixed: user roles did not keep the sort order after a page reload. It now keeps it until the user changes it.

Known issue: the approval queue does not yet support files larger than 50 MB. Larger files are refused with a message that names the limit.

Planned, not yet available: the batch number in the import screen. We expect to release it in the second half of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 17.8

Quillstone Ledger 17.8 was released on 22 June 2022 for customers on the hosted plan.

What changed:

- The search index: now includes a column for the reviewer's comment count. The column is empty for records created before 17.8, because the reviewer's comment count was not stored then.
- The import screen: now keeps the exchange rate used for 365 days. After that the value is removed, and the entry shows when it was removed.
- Attachment storage: now names the user who last changed the batch number. It does not show why the change was made, because the product does not ask for a reason.
- Attachment storage: now keeps the payment terms for 90 days. After that the value is removed, and the entry shows when it was removed.

- Fixed: a date typed as day/month could be read as month/day in invoice matching. The format now follows the account's settings.
- Fixed: a date typed as day/month could be read as month/day in attachment storage. The format now follows the account's settings.

Planned, not yet available: the original currency in user roles. We expect to release it in the second half of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 17.9

Quillstone Ledger 17.9 was released on 1 July 2022 for customers on the hosted plan.

What changed:

- Scheduled reports: now names the user who last changed the exchange rate used. It does not show why the change was made, because the product does not ask for a reason.
- Exports: now accepts the original currency as a filter. The filter applies to records created on or after 17.9.
- The import screen: now shows the reviewer's comment count beside each entry. Entries made before 17.9 show a dash, because the value was not recorded.
- User roles: now includes a column for the time zone of the entry. The column is empty for records created before 17.9, because the time zone of the entry was not stored then.

- Fixed: a date typed as day/month could be read as month/day in the audit view. The format now follows the account's settings.
- Fixed: a date typed as day/month could be read as month/day in currency conversion. The format now follows the account's settings.

Planned, not yet available: the batch number in the search index. We expect to release it in the first quarter of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 18.0

Quillstone Ledger 18.0 was released on 8 July 2022 for customers on the hosted plan.

What changed:

- The import screen: now accepts the approving user's role as a filter. The filter applies to records created on or after 18.0.
- The import screen: now includes a column for the import job number. The column is empty for records created before 18.0, because the import job number was not stored then.
- The import screen: now names the user who last changed the original currency. It does not show why the change was made, because the product does not ask for a reason.

- Fixed: totals in invoice matching were rounded before they were added, which could change the last digit. They are now added first and rounded once.

Planned, not yet available: the cost centre in the activity log. We expect to release it in the second half of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 18.1

Quillstone Ledger 18.1 was released on 15 July 2022 for customers on the hosted plan.

What changed:

- Currency conversion: can now be turned off for each user. When it is turned off, the payment terms is not collected for that user from then on.
- Attachment storage: now keeps the payment terms for 30 days. After that the value is removed, and the entry shows when it was removed.
- The search index: now shows the batch number beside each entry. Entries made before 18.1 show a dash, because the value was not recorded.
- The tax summary: now accepts the time zone of the entry as a filter. The filter applies to records created on or after 18.1.
- The search index: can now be turned off for each user. When it is turned off, the attachment count is not collected for that user from then on.

- Fixed: user roles did not keep the sort order after a page reload. It now keeps it until the user changes it.

Planned, not yet available: the exchange rate used in invoice matching. We expect to release it in the second half of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 18.2

Quillstone Ledger 18.2 was released on 22 July 2022 for customers on the hosted plan.

What changed:

- The tax summary: now accepts the payment terms as a filter. The filter applies to records created on or after 18.2.
- Exports: now names the user who last changed the attachment count. It does not show why the change was made, because the product does not ask for a reason.
- Invoice matching: now shows the exchange rate used beside each entry. Entries made before 18.2 show a dash, because the value was not recorded.
- Account settings: now includes a column for the original currency. The column is empty for records created before 18.2, because the original currency was not stored then.
- Invoice matching: now shows the payment terms beside each entry. Entries made before 18.2 show a dash, because the value was not recorded.

- Fixed: a date typed as day/month could be read as month/day in invoice matching. The format now follows the account's settings.
- Fixed: a date typed as day/month could be read as month/day in the activity log. The format now follows the account's settings.

## Quillstone Ledger 18.3

Quillstone Ledger 18.3 was released on 1 August 2022 for customers on the hosted plan.

What changed:

- Scheduled reports: can now be turned off for each user. When it is turned off, the time zone of the entry is not collected for that user from then on.
- The approval queue: now includes a column for the reviewer's comment count. The column is empty for records created before 18.3, because the reviewer's comment count was not stored then.
- The import screen: now includes a column for the original currency. The column is empty for records created before 18.3, because the original currency was not stored then.
- Exports: can now be turned off for each user. When it is turned off, the cost centre is not collected for that user from then on.
- Attachment storage: faster on accounts with many records. In our own test on 1 August 2022, a 41,000-row run took 78 seconds on 18.2 and 15 seconds on 18.3. Your results depend on your data.

- Fixed: attachment storage did not keep the sort order after a page reload. It now keeps it until the user changes it.

Known issue: the approval queue does not yet support files larger than 100 MB. Larger files are refused with a message that names the limit.

## Quillstone Ledger 18.4

Quillstone Ledger 18.4 was released on 8 August 2022 for customers on the hosted plan.

What changed:

- Invoice matching: now keeps the import job number for 90 days. After that the value is removed, and the entry shows when it was removed.
- The mobile view: can now be turned off for each user. When it is turned off, the cost centre is not collected for that user from then on.
- Report generation: now names the user who last changed the cost centre. It does not show why the change was made, because the product does not ask for a reason.

- Fixed: in user roles, a month with no records showed a blank page. It now shows an empty table.
- Fixed: totals in the search index were rounded before they were added, which could change the last digit. They are now added first and rounded once.

Planned, not yet available: the attachment count in user roles. We expect to release it in the second half of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 18.5

Quillstone Ledger 18.5 was released on 15 August 2022 for customers on the hosted plan.

What changed:

- Report generation: now accepts the import job number as a filter. The filter applies to records created on or after 18.5.
- The mobile view: now shows the cost centre beside each entry. Entries made before 18.5 show a dash, because the value was not recorded.
- The search index: now names the user who last changed the reviewer's comment count. It does not show why the change was made, because the product does not ask for a reason.
- Account settings: faster on accounts with many records. In our own test on 15 August 2022, a 104,000-row run took 58 seconds on 18.4 and 14 seconds on 18.5. Your results depend on your data.
- The audit view: now names the user who last changed the original currency. It does not show why the change was made, because the product does not ask for a reason.

- Fixed: in the audit view, a month with no records showed a blank page. It now shows an empty table.
- Fixed: invoice matching did not keep the sort order after a page reload. It now keeps it until the user changes it.

Planned, not yet available: the batch number in invoice matching. We expect to release it in the second half of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 18.6

Quillstone Ledger 18.6 was released on 22 August 2022 for customers on the hosted plan.

What changed:

- Currency conversion: now names the user who last changed the device type. It does not show why the change was made, because the product does not ask for a reason.
- Report generation: now keeps the attachment count for 365 days. After that the value is removed, and the entry shows when it was removed.
- The mobile view: now names the user who last changed the reviewer's comment count. It does not show why the change was made, because the product does not ask for a reason.
- User roles: now names the user who last changed the payment terms. It does not show why the change was made, because the product does not ask for a reason.

- Fixed: totals in the search index were rounded before they were added, which could change the last digit. They are now added first and rounded once.

## Quillstone Ledger 18.7

Quillstone Ledger 18.7 was released on 1 September 2022 for customers on the hosted plan.

What changed:

- The approval queue: faster on accounts with many records. In our own test on 1 September 2022, a 161,000-row run took 79 seconds on 18.6 and 19 seconds on 18.7. Your results depend on your data.
- Account settings: can now be turned off for each user. When it is turned off, the exchange rate used is not collected for that user from then on.
- Bank feeds: can now be turned off for each user. When it is turned off, the payment terms is not collected for that user from then on.
- The approval queue: now accepts the exchange rate used as a filter. The filter applies to records created on or after 18.7.
- Exports: now names the user who last changed the payment terms. It does not show why the change was made, because the product does not ask for a reason.

- Fixed: totals in bank feeds were rounded before they were added, which could change the last digit. They are now added first and rounded once.

Known issue: user roles does not yet support files larger than 25 MB. Larger files are refused with a message that names the limit.

## Quillstone Ledger 18.8

Quillstone Ledger 18.8 was released on 8 September 2022 for customers on the hosted plan.

What changed:

- The activity log: now keeps the exchange rate used for 90 days. After that the value is removed, and the entry shows when it was removed.
- The approval queue: now shows the exchange rate used beside each entry. Entries made before 18.8 show a dash, because the value was not recorded.
- The approval queue: now shows the original currency beside each entry. Entries made before 18.8 show a dash, because the value was not recorded.

- Fixed: in account settings, a month with no records showed a blank page. It now shows an empty table.

Known issue: on slow connections the approval queue can take up to 30 seconds to load. We are working on it and will say here when it is fixed.

## Quillstone Ledger 18.9

Quillstone Ledger 18.9 was released on 15 September 2022 for customers on the hosted plan.

What changed:

- The import screen: now includes a column for the approving user's role. The column is empty for records created before 18.9, because the approving user's role was not stored then.
- The approval queue: now accepts the batch number as a filter. The filter applies to records created on or after 18.9.
- Invoice matching: now accepts the batch number as a filter. The filter applies to records created on or after 18.9.
- Exports: now names the user who last changed the import job number. It does not show why the change was made, because the product does not ask for a reason.

- Fixed: a date typed as day/month could be read as month/day in the tax summary. The format now follows the account's settings.
- Fixed: in account settings, a month with no records showed a blank page. It now shows an empty table.

Known issue: invoice matching does not yet support files larger than 50 MB. Larger files are refused with a message that names the limit.

## Quillstone Ledger 19.0

Quillstone Ledger 19.0 was released on 22 September 2022 for customers on the hosted plan.

What changed:

- The import screen: now shows the time zone of the entry beside each entry. Entries made before 19.0 show a dash, because the value was not recorded.
- Scheduled reports: now names the user who last changed the original currency. It does not show why the change was made, because the product does not ask for a reason.
- User roles: now names the user who last changed the import job number. It does not show why the change was made, because the product does not ask for a reason.

- Fixed: a date typed as day/month could be read as month/day in attachment storage. The format now follows the account's settings.

Known issue: the activity log does not yet support files larger than 25 MB. Larger files are refused with a message that names the limit.

## Quillstone Ledger 19.1

Quillstone Ledger 19.1 was released on 1 October 2022 for customers on the hosted plan.

What changed:

- The audit view: now includes a column for the source file name. The column is empty for records created before 19.1, because the source file name was not stored then.
- Attachment storage: now keeps the import job number for 30 days. After that the value is removed, and the entry shows when it was removed.
- Account settings: now keeps the cost centre for 180 days. After that the value is removed, and the entry shows when it was removed.
- User roles: now shows the approving user's role beside each entry. Entries made before 19.1 show a dash, because the value was not recorded.

- Fixed: bank feeds did not keep the sort order after a page reload. It now keeps it until the user changes it.
- Fixed: scheduled reports did not keep the sort order after a page reload. It now keeps it until the user changes it.

Planned, not yet available: the device type in scheduled reports. We expect to release it in the second quarter of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 19.2

Quillstone Ledger 19.2 was released on 8 October 2022 for customers on the hosted plan.

What changed:

- Invoice matching: now keeps the original currency for 180 days. After that the value is removed, and the entry shows when it was removed.
- Bank feeds: now names the user who last changed the reviewer's comment count. It does not show why the change was made, because the product does not ask for a reason.
- The tax summary: now includes a column for the import job number. The column is empty for records created before 19.2, because the import job number was not stored then.
- The approval queue: now names the user who last changed the time zone of the entry. It does not show why the change was made, because the product does not ask for a reason.
- The mobile view: now names the user who last changed the import job number. It does not show why the change was made, because the product does not ask for a reason.

- Fixed: totals in the audit view were rounded before they were added, which could change the last digit. They are now added first and rounded once.
- Fixed: totals in user roles were rounded before they were added, which could change the last digit. They are now added first and rounded once.

Known issue: on slow connections the search index can take up to 10 seconds to load. We are working on it and will say here when it is fixed.

Planned, not yet available: the device type in the mobile view. We expect to release it in the second quarter of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 19.3

Quillstone Ledger 19.3 was released on 15 October 2022 for customers on the hosted plan.

What changed:

- Exports: now accepts the payment terms as a filter. The filter applies to records created on or after 19.3.
- Exports: can now be turned off for each user. When it is turned off, the approving user's role is not collected for that user from then on.
- Invoice matching: can now be turned off for each user. When it is turned off, the batch number is not collected for that user from then on.
- Attachment storage: now includes a column for the import job number. The column is empty for records created before 19.3, because the import job number was not stored then.
- Currency conversion: now includes a column for the exchange rate used. The column is empty for records created before 19.3, because the exchange rate used was not stored then.

- Fixed: totals in the approval queue were rounded before they were added, which could change the last digit. They are now added first and rounded once.
- Fixed: a date typed as day/month could be read as month/day in the mobile view. The format now follows the account's settings.

## Quillstone Ledger 19.4

Quillstone Ledger 19.4 was released on 22 October 2022 for customers on the hosted plan.

What changed:

- The mobile view: now accepts the cost centre as a filter. The filter applies to records created on or after 19.4.
- Exports: now shows the reviewer's comment count beside each entry. Entries made before 19.4 show a dash, because the value was not recorded.
- Account settings: now keeps the cost centre for 365 days. After that the value is removed, and the entry shows when it was removed.
- Invoice matching: now includes a column for the import job number. The column is empty for records created before 19.4, because the import job number was not stored then.
- Invoice matching: now includes a column for the device type. The column is empty for records created before 19.4, because the device type was not stored then.

- Fixed: the tax summary did not keep the sort order after a page reload. It now keeps it until the user changes it.
- Fixed: in account settings, a month with no records showed a blank page. It now shows an empty table.

Known issue: on slow connections currency conversion can take up to 20 seconds to load. We are working on it and will say here when it is fixed.

Planned, not yet available: the payment terms in user roles. We expect to release it in the first quarter of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 19.5

Quillstone Ledger 19.5 was released on 1 November 2022 for customers on the hosted plan.

What changed:

- Exports: can now be turned off for each user. When it is turned off, the payment terms is not collected for that user from then on.
- Account settings: can now be turned off for each user. When it is turned off, the time zone of the entry is not collected for that user from then on.
- The search index: now includes a column for the import job number. The column is empty for records created before 19.5, because the import job number was not stored then.

- Fixed: a date typed as day/month could be read as month/day in attachment storage. The format now follows the account's settings.
- Fixed: a date typed as day/month could be read as month/day in the approval queue. The format now follows the account's settings.

Known issue: invoice matching does not yet support files larger than 25 MB. Larger files are refused with a message that names the limit.

## Quillstone Ledger 19.6

Quillstone Ledger 19.6 was released on 8 November 2022 for customers on the hosted plan.

What changed:

- The search index: now accepts the reviewer's comment count as a filter. The filter applies to records created on or after 19.6.
- Attachment storage: now shows the time zone of the entry beside each entry. Entries made before 19.6 show a dash, because the value was not recorded.
- Report generation: faster on accounts with many records. In our own test on 8 November 2022, a 70,000-row run took 67 seconds on 19.5 and 13 seconds on 19.6. Your results depend on your data.
- The search index: now accepts the import job number as a filter. The filter applies to records created on or after 19.6.
- Account settings: now accepts the batch number as a filter. The filter applies to records created on or after 19.6.

- Fixed: in exports, a month with no records showed a blank page. It now shows an empty table.

Known issue: on slow connections the activity log can take up to 10 seconds to load. We are working on it and will say here when it is fixed.

## Quillstone Ledger 19.7

Quillstone Ledger 19.7 was released on 15 November 2022 for customers on the hosted plan.

What changed:

- The audit view: now shows the attachment count beside each entry. Entries made before 19.7 show a dash, because the value was not recorded.
- Invoice matching: now accepts the exchange rate used as a filter. The filter applies to records created on or after 19.7.
- The approval queue: can now be turned off for each user. When it is turned off, the approving user's role is not collected for that user from then on.
- Exports: now shows the reviewer's comment count beside each entry. Entries made before 19.7 show a dash, because the value was not recorded.
- Invoice matching: now names the user who last changed the batch number. It does not show why the change was made, because the product does not ask for a reason.

- Fixed: in report generation, a month with no records showed a blank page. It now shows an empty table.
- Fixed: currency conversion did not keep the sort order after a page reload. It now keeps it until the user changes it.

Planned, not yet available: the reviewer's comment count in account settings. We expect to release it in the second half of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 19.8

Quillstone Ledger 19.8 was released on 22 November 2022 for customers on the hosted plan.

What changed:

- The search index: now includes a column for the original currency. The column is empty for records created before 19.8, because the original currency was not stored then.
- Report generation: now includes a column for the attachment count. The column is empty for records created before 19.8, because the attachment count was not stored then.
- Report generation: now keeps the device type for 90 days. After that the value is removed, and the entry shows when it was removed.
- Report generation: now shows the source file name beside each entry. Entries made before 19.8 show a dash, because the value was not recorded.
- Account settings: can now be turned off for each user. When it is turned off, the payment terms is not collected for that user from then on.

- Fixed: the activity log did not keep the sort order after a page reload. It now keeps it until the user changes it.
- Fixed: in invoice matching, a month with no records showed a blank page. It now shows an empty table.

## Quillstone Ledger 19.9

Quillstone Ledger 19.9 was released on 1 December 2022 for customers on the hosted plan.

What changed:

- Report generation: now keeps the exchange rate used for 30 days. After that the value is removed, and the entry shows when it was removed.
- User roles: now includes a column for the original currency. The column is empty for records created before 19.9, because the original currency was not stored then.
- The approval queue: can now be turned off for each user. When it is turned off, the payment terms is not collected for that user from then on.

- Fixed: a date typed as day/month could be read as month/day in invoice matching. The format now follows the account's settings.

Known issue: currency conversion does not yet support files larger than 25 MB. Larger files are refused with a message that names the limit.

## Quillstone Ledger 20.0

Quillstone Ledger 20.0 was released on 8 December 2022 for customers on the hosted plan.

What changed:

- Bank feeds: now accepts the approving user's role as a filter. The filter applies to records created on or after 20.0.
- The mobile view: now includes a column for the import job number. The column is empty for records created before 20.0, because the import job number was not stored then.
- The mobile view: now names the user who last changed the attachment count. It does not show why the change was made, because the product does not ask for a reason.
- Currency conversion: now accepts the cost centre as a filter. The filter applies to records created on or after 20.0.
- Account settings: now keeps the original currency for 30 days. After that the value is removed, and the entry shows when it was removed.

- Fixed: a date typed as day/month could be read as month/day in the activity log. The format now follows the account's settings.
- Fixed: the audit view did not keep the sort order after a page reload. It now keeps it until the user changes it.

Known issue: currency conversion does not yet support files larger than 50 MB. Larger files are refused with a message that names the limit.

## Quillstone Ledger 20.1

Quillstone Ledger 20.1 was released on 15 December 2022 for customers on the hosted plan.

What changed:

- Account settings: now keeps the source file name for 30 days. After that the value is removed, and the entry shows when it was removed.
- The search index: now names the user who last changed the reviewer's comment count. It does not show why the change was made, because the product does not ask for a reason.
- Scheduled reports: now accepts the attachment count as a filter. The filter applies to records created on or after 20.1.

- Fixed: the import screen did not keep the sort order after a page reload. It now keeps it until the user changes it.

## Quillstone Ledger 20.2

Quillstone Ledger 20.2 was released on 22 December 2022 for customers on the hosted plan.

What changed:

- User roles: now includes a column for the approving user's role. The column is empty for records created before 20.2, because the approving user's role was not stored then.
- Currency conversion: now accepts the time zone of the entry as a filter. The filter applies to records created on or after 20.2.
- Attachment storage: now shows the exchange rate used beside each entry. Entries made before 20.2 show a dash, because the value was not recorded.
- The mobile view: now shows the exchange rate used beside each entry. Entries made before 20.2 show a dash, because the value was not recorded.
- The search index: now shows the import job number beside each entry. Entries made before 20.2 show a dash, because the value was not recorded.

- Fixed: in the search index, a month with no records showed a blank page. It now shows an empty table.

Known issue: on slow connections attachment storage can take up to 20 seconds to load. We are working on it and will say here when it is fixed.

## Quillstone Ledger 20.3

Quillstone Ledger 20.3 was released on 1 January 2023 for customers on the hosted plan.

What changed:

- Account settings: can now be turned off for each user. When it is turned off, the exchange rate used is not collected for that user from then on.
- User roles: now keeps the import job number for 180 days. After that the value is removed, and the entry shows when it was removed.
- The tax summary: now includes a column for the payment terms. The column is empty for records created before 20.3, because the payment terms was not stored then.
- Attachment storage: now keeps the device type for 30 days. After that the value is removed, and the entry shows when it was removed.
- Currency conversion: now shows the payment terms beside each entry. Entries made before 20.3 show a dash, because the value was not recorded.

- Fixed: the tax summary did not keep the sort order after a page reload. It now keeps it until the user changes it.

Known issue: the mobile view does not yet support files larger than 25 MB. Larger files are refused with a message that names the limit.

Planned, not yet available: the source file name in bank feeds. We expect to release it in the second quarter of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 20.4

Quillstone Ledger 20.4 was released on 8 January 2023 for customers on the hosted plan.

What changed:

- The search index: now shows the device type beside each entry. Entries made before 20.4 show a dash, because the value was not recorded.
- The approval queue: now accepts the attachment count as a filter. The filter applies to records created on or after 20.4.
- The mobile view: faster on accounts with many records. In our own test on 8 January 2023, a 40,000-row run took 66 seconds on 20.3 and 22 seconds on 20.4. Your results depend on your data.
- The tax summary: now shows the cost centre beside each entry. Entries made before 20.4 show a dash, because the value was not recorded.
- The activity log: faster on accounts with many records. In our own test on 8 January 2023, a 139,000-row run took 50 seconds on 20.3 and 16 seconds on 20.4. Your results depend on your data.

- Fixed: totals in currency conversion were rounded before they were added, which could change the last digit. They are now added first and rounded once.
- Fixed: the tax summary did not keep the sort order after a page reload. It now keeps it until the user changes it.

Known issue: invoice matching does not yet support files larger than 25 MB. Larger files are refused with a message that names the limit.

## Quillstone Ledger 20.5

Quillstone Ledger 20.5 was released on 15 January 2023 for customers on the hosted plan.

What changed:

- Exports: can now be turned off for each user. When it is turned off, the import job number is not collected for that user from then on.
- Attachment storage: now keeps the source file name for 30 days. After that the value is removed, and the entry shows when it was removed.
- The audit view: now accepts the device type as a filter. The filter applies to records created on or after 20.5.
- Currency conversion: faster on accounts with many records. In our own test on 15 January 2023, a 12,000-row run took 74 seconds on 20.4 and 14 seconds on 20.5. Your results depend on your data.
- Exports: now shows the batch number beside each entry. Entries made before 20.5 show a dash, because the value was not recorded.

- Fixed: totals in report generation were rounded before they were added, which could change the last digit. They are now added first and rounded once.

Known issue: on slow connections bank feeds can take up to 10 seconds to load. We are working on it and will say here when it is fixed.

## Quillstone Ledger 20.6

Quillstone Ledger 20.6 was released on 22 January 2023 for customers on the hosted plan.

What changed:

- The approval queue: now shows the payment terms beside each entry. Entries made before 20.6 show a dash, because the value was not recorded.
- Exports: now keeps the payment terms for 365 days. After that the value is removed, and the entry shows when it was removed.
- Invoice matching: now keeps the attachment count for 365 days. After that the value is removed, and the entry shows when it was removed.
- Invoice matching: now accepts the original currency as a filter. The filter applies to records created on or after 20.6.
- Attachment storage: now includes a column for the original currency. The column is empty for records created before 20.6, because the original currency was not stored then.

- Fixed: in invoice matching, a month with no records showed a blank page. It now shows an empty table.
- Fixed: a date typed as day/month could be read as month/day in invoice matching. The format now follows the account's settings.

## Quillstone Ledger 20.7

Quillstone Ledger 20.7 was released on 1 February 2023 for customers on the hosted plan.

What changed:

- The import screen: can now be turned off for each user. When it is turned off, the reviewer's comment count is not collected for that user from then on.
- The tax summary: faster on accounts with many records. In our own test on 1 February 2023, a 126,000-row run took 22 seconds on 20.6 and 5 seconds on 20.7. Your results depend on your data.
- Currency conversion: now accepts the approving user's role as a filter. The filter applies to records created on or after 20.7.
- User roles: can now be turned off for each user. When it is turned off, the approving user's role is not collected for that user from then on.
- Invoice matching: now keeps the import job number for 365 days. After that the value is removed, and the entry shows when it was removed.

- Fixed: report generation did not keep the sort order after a page reload. It now keeps it until the user changes it.
- Fixed: a date typed as day/month could be read as month/day in exports. The format now follows the account's settings.

Known issue: on slow connections the mobile view can take up to 10 seconds to load. We are working on it and will say here when it is fixed.

## Quillstone Ledger 20.8

Quillstone Ledger 20.8 was released on 8 February 2023 for customers on the hosted plan.

What changed:

- The activity log: now accepts the batch number as a filter. The filter applies to records created on or after 20.8.
- The mobile view: can now be turned off for each user. When it is turned off, the time zone of the entry is not collected for that user from then on.
- The search index: now includes a column for the device type. The column is empty for records created before 20.8, because the device type was not stored then.
- Exports: now names the user who last changed the payment terms. It does not show why the change was made, because the product does not ask for a reason.

- Fixed: invoice matching did not keep the sort order after a page reload. It now keeps it until the user changes it.

## Quillstone Ledger 20.9

Quillstone Ledger 20.9 was released on 15 February 2023 for customers on the hosted plan.

What changed:

- The mobile view: now keeps the approving user's role for 180 days. After that the value is removed, and the entry shows when it was removed.
- The import screen: now shows the device type beside each entry. Entries made before 20.9 show a dash, because the value was not recorded.
- The tax summary: now includes a column for the exchange rate used. The column is empty for records created before 20.9, because the exchange rate used was not stored then.

- Fixed: the search index did not keep the sort order after a page reload. It now keeps it until the user changes it.
- Fixed: totals in exports were rounded before they were added, which could change the last digit. They are now added first and rounded once.

Planned, not yet available: the approving user's role in invoice matching. We expect to release it in the second quarter of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 21.0

Quillstone Ledger 21.0 was released on 22 February 2023 for customers on the hosted plan.

What changed:

- Account settings: faster on accounts with many records. In our own test on 22 February 2023, a 175,000-row run took 38 seconds on 20.9 and 12 seconds on 21.0. Your results depend on your data.
- Currency conversion: now names the user who last changed the batch number. It does not show why the change was made, because the product does not ask for a reason.
- Account settings: now accepts the reviewer's comment count as a filter. The filter applies to records created on or after 21.0.
- The import screen: now includes a column for the original currency. The column is empty for records created before 21.0, because the original currency was not stored then.

- Fixed: the audit view did not keep the sort order after a page reload. It now keeps it until the user changes it.
- Fixed: totals in report generation were rounded before they were added, which could change the last digit. They are now added first and rounded once.

## Quillstone Ledger 21.1

Quillstone Ledger 21.1 was released on 1 March 2023 for customers on the hosted plan.

What changed:

- The approval queue: faster on accounts with many records. In our own test on 1 March 2023, a 86,000-row run took 69 seconds on 21.0 and 13 seconds on 21.1. Your results depend on your data.
- Attachment storage: faster on accounts with many records. In our own test on 1 March 2023, a 111,000-row run took 73 seconds on 21.0 and 18 seconds on 21.1. Your results depend on your data.
- The activity log: faster on accounts with many records. In our own test on 1 March 2023, a 156,000-row run took 62 seconds on 21.0 and 15 seconds on 21.1. Your results depend on your data.

- Fixed: in the tax summary, a month with no records showed a blank page. It now shows an empty table.

Known issue: user roles does not yet support files larger than 100 MB. Larger files are refused with a message that names the limit.

## Quillstone Ledger 21.2

Quillstone Ledger 21.2 was released on 8 March 2023 for customers on the hosted plan.

What changed:

- Account settings: can now be turned off for each user. When it is turned off, the exchange rate used is not collected for that user from then on.
- The import screen: can now be turned off for each user. When it is turned off, the import job number is not collected for that user from then on.
- User roles: now keeps the import job number for 90 days. After that the value is removed, and the entry shows when it was removed.

- Fixed: a date typed as day/month could be read as month/day in the import screen. The format now follows the account's settings.
- Fixed: in the mobile view, a month with no records showed a blank page. It now shows an empty table.

Known issue: on slow connections invoice matching can take up to 20 seconds to load. We are working on it and will say here when it is fixed.

## Quillstone Ledger 21.3

Quillstone Ledger 21.3 was released on 15 March 2023 for customers on the hosted plan.

What changed:

- Currency conversion: now names the user who last changed the device type. It does not show why the change was made, because the product does not ask for a reason.
- The import screen: now accepts the device type as a filter. The filter applies to records created on or after 21.3.
- The audit view: faster on accounts with many records. In our own test on 15 March 2023, a 106,000-row run took 66 seconds on 21.2 and 33 seconds on 21.3. Your results depend on your data.
- The audit view: now names the user who last changed the source file name. It does not show why the change was made, because the product does not ask for a reason.

- Fixed: a date typed as day/month could be read as month/day in the import screen. The format now follows the account's settings.
- Fixed: in scheduled reports, a month with no records showed a blank page. It now shows an empty table.

Known issue: on slow connections invoice matching can take up to 20 seconds to load. We are working on it and will say here when it is fixed.

Planned, not yet available: the batch number in attachment storage. We expect to release it in the first quarter of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 21.4

Quillstone Ledger 21.4 was released on 22 March 2023 for customers on the hosted plan.

What changed:

- Report generation: now includes a column for the time zone of the entry. The column is empty for records created before 21.4, because the time zone of the entry was not stored then.
- The activity log: now accepts the time zone of the entry as a filter. The filter applies to records created on or after 21.4.
- Bank feeds: now keeps the device type for 365 days. After that the value is removed, and the entry shows when it was removed.
- The approval queue: now includes a column for the source file name. The column is empty for records created before 21.4, because the source file name was not stored then.

- Fixed: totals in the activity log were rounded before they were added, which could change the last digit. They are now added first and rounded once.

## Quillstone Ledger 21.5

Quillstone Ledger 21.5 was released on 1 April 2023 for customers on the hosted plan.

What changed:

- The audit view: can now be turned off for each user. When it is turned off, the original currency is not collected for that user from then on.
- Report generation: now keeps the source file name for 180 days. After that the value is removed, and the entry shows when it was removed.
- Attachment storage: now names the user who last changed the device type. It does not show why the change was made, because the product does not ask for a reason.
- Scheduled reports: now accepts the device type as a filter. The filter applies to records created on or after 21.5.

- Fixed: the audit view did not keep the sort order after a page reload. It now keeps it until the user changes it.
- Fixed: in the approval queue, a month with no records showed a blank page. It now shows an empty table.

## Quillstone Ledger 21.6

Quillstone Ledger 21.6 was released on 8 April 2023 for customers on the hosted plan.

What changed:

- Attachment storage: now shows the original currency beside each entry. Entries made before 21.6 show a dash, because the value was not recorded.
- Currency conversion: faster on accounts with many records. In our own test on 8 April 2023, a 149,000-row run took 78 seconds on 21.5 and 15 seconds on 21.6. Your results depend on your data.
- Exports: can now be turned off for each user. When it is turned off, the time zone of the entry is not collected for that user from then on.
- The audit view: now accepts the original currency as a filter. The filter applies to records created on or after 21.6.

- Fixed: totals in user roles were rounded before they were added, which could change the last digit. They are now added first and rounded once.

## Quillstone Ledger 21.7

Quillstone Ledger 21.7 was released on 15 April 2023 for customers on the hosted plan.

What changed:

- Bank feeds: faster on accounts with many records. In our own test on 15 April 2023, a 172,000-row run took 44 seconds on 21.6 and 14 seconds on 21.7. Your results depend on your data.
- The activity log: faster on accounts with many records. In our own test on 15 April 2023, a 67,000-row run took 26 seconds on 21.6 and 13 seconds on 21.7. Your results depend on your data.
- Report generation: faster on accounts with many records. In our own test on 15 April 2023, a 96,000-row run took 46 seconds on 21.6 and 23 seconds on 21.7. Your results depend on your data.
- The search index: now names the user who last changed the batch number. It does not show why the change was made, because the product does not ask for a reason.

- Fixed: a date typed as day/month could be read as month/day in the audit view. The format now follows the account's settings.
- Fixed: in the approval queue, a month with no records showed a blank page. It now shows an empty table.

Known issue: attachment storage does not yet support files larger than 50 MB. Larger files are refused with a message that names the limit.

## Quillstone Ledger 21.8

Quillstone Ledger 21.8 was released on 22 April 2023 for customers on the hosted plan.

What changed:

- The import screen: now includes a column for the reviewer's comment count. The column is empty for records created before 21.8, because the reviewer's comment count was not stored then.
- The activity log: now accepts the payment terms as a filter. The filter applies to records created on or after 21.8.
- Invoice matching: now names the user who last changed the cost centre. It does not show why the change was made, because the product does not ask for a reason.
- Invoice matching: now keeps the time zone of the entry for 365 days. After that the value is removed, and the entry shows when it was removed.

- Fixed: invoice matching did not keep the sort order after a page reload. It now keeps it until the user changes it.

Known issue: on slow connections the tax summary can take up to 30 seconds to load. We are working on it and will say here when it is fixed.

## Quillstone Ledger 21.9

Quillstone Ledger 21.9 was released on 1 May 2023 for customers on the hosted plan.

What changed:

- The tax summary: now names the user who last changed the original currency. It does not show why the change was made, because the product does not ask for a reason.
- Attachment storage: now shows the attachment count beside each entry. Entries made before 21.9 show a dash, because the value was not recorded.
- Exports: can now be turned off for each user. When it is turned off, the payment terms is not collected for that user from then on.

- Fixed: a date typed as day/month could be read as month/day in report generation. The format now follows the account's settings.

Planned, not yet available: the payment terms in the approval queue. We expect to release it in the second quarter of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 22.0

Quillstone Ledger 22.0 was released on 8 May 2023 for customers on the hosted plan.

What changed:

- The search index: now includes a column for the device type. The column is empty for records created before 22.0, because the device type was not stored then.
- Scheduled reports: now keeps the batch number for 180 days. After that the value is removed, and the entry shows when it was removed.
- The mobile view: now shows the import job number beside each entry. Entries made before 22.0 show a dash, because the value was not recorded.

- Fixed: totals in bank feeds were rounded before they were added, which could change the last digit. They are now added first and rounded once.

Known issue: scheduled reports does not yet support files larger than 25 MB. Larger files are refused with a message that names the limit.

## Quillstone Ledger 22.1

Quillstone Ledger 22.1 was released on 15 May 2023 for customers on the hosted plan.

What changed:

- Invoice matching: now keeps the cost centre for 30 days. After that the value is removed, and the entry shows when it was removed.
- Scheduled reports: now keeps the original currency for 365 days. After that the value is removed, and the entry shows when it was removed.
- The search index: now shows the batch number beside each entry. Entries made before 22.1 show a dash, because the value was not recorded.
- Report generation: now keeps the cost centre for 365 days. After that the value is removed, and the entry shows when it was removed.

- Fixed: attachment storage did not keep the sort order after a page reload. It now keeps it until the user changes it.

## Quillstone Ledger 22.2

Quillstone Ledger 22.2 was released on 22 May 2023 for customers on the hosted plan.

What changed:

- Currency conversion: now includes a column for the time zone of the entry. The column is empty for records created before 22.2, because the time zone of the entry was not stored then.
- Report generation: now names the user who last changed the reviewer's comment count. It does not show why the change was made, because the product does not ask for a reason.
- Currency conversion: now accepts the reviewer's comment count as a filter. The filter applies to records created on or after 22.2.
- User roles: can now be turned off for each user. When it is turned off, the device type is not collected for that user from then on.
- Scheduled reports: now keeps the cost centre for 30 days. After that the value is removed, and the entry shows when it was removed.

- Fixed: totals in scheduled reports were rounded before they were added, which could change the last digit. They are now added first and rounded once.

Known issue: user roles does not yet support files larger than 50 MB. Larger files are refused with a message that names the limit.

## Quillstone Ledger 22.3

Quillstone Ledger 22.3 was released on 1 June 2023 for customers on the hosted plan.

What changed:

- The search index: now accepts the device type as a filter. The filter applies to records created on or after 22.3.
- User roles: now accepts the attachment count as a filter. The filter applies to records created on or after 22.3.
- The search index: now keeps the source file name for 180 days. After that the value is removed, and the entry shows when it was removed.
- The mobile view: faster on accounts with many records. In our own test on 1 June 2023, a 75,000-row run took 84 seconds on 22.2 and 28 seconds on 22.3. Your results depend on your data.
- The search index: now names the user who last changed the source file name. It does not show why the change was made, because the product does not ask for a reason.

- Fixed: in the activity log, a month with no records showed a blank page. It now shows an empty table.
- Fixed: in the audit view, a month with no records showed a blank page. It now shows an empty table.

## Quillstone Ledger 22.4

Quillstone Ledger 22.4 was released on 8 June 2023 for customers on the hosted plan.

What changed:

- Report generation: now accepts the original currency as a filter. The filter applies to records created on or after 22.4.
- The activity log: now accepts the exchange rate used as a filter. The filter applies to records created on or after 22.4.
- Currency conversion: now names the user who last changed the import job number. It does not show why the change was made, because the product does not ask for a reason.

- Fixed: in the mobile view, a month with no records showed a blank page. It now shows an empty table.

Known issue: invoice matching does not yet support files larger than 50 MB. Larger files are refused with a message that names the limit.

## Quillstone Ledger 22.5

Quillstone Ledger 22.5 was released on 15 June 2023 for customers on the hosted plan.

What changed:

- User roles: now accepts the original currency as a filter. The filter applies to records created on or after 22.5.
- Exports: now accepts the batch number as a filter. The filter applies to records created on or after 22.5.
- Scheduled reports: can now be turned off for each user. When it is turned off, the exchange rate used is not collected for that user from then on.

- Fixed: the activity log did not keep the sort order after a page reload. It now keeps it until the user changes it.
- Fixed: a date typed as day/month could be read as month/day in user roles. The format now follows the account's settings.

## Quillstone Ledger 22.6

Quillstone Ledger 22.6 was released on 22 June 2023 for customers on the hosted plan.

What changed:

- Invoice matching: now keeps the attachment count for 180 days. After that the value is removed, and the entry shows when it was removed.
- The audit view: can now be turned off for each user. When it is turned off, the payment terms is not collected for that user from then on.
- Account settings: now accepts the time zone of the entry as a filter. The filter applies to records created on or after 22.6.
- The mobile view: now accepts the attachment count as a filter. The filter applies to records created on or after 22.6.
- User roles: now accepts the reviewer's comment count as a filter. The filter applies to records created on or after 22.6.

- Fixed: totals in account settings were rounded before they were added, which could change the last digit. They are now added first and rounded once.

Known issue: on slow connections the tax summary can take up to 20 seconds to load. We are working on it and will say here when it is fixed.

## Quillstone Ledger 22.7

Quillstone Ledger 22.7 was released on 1 July 2023 for customers on the hosted plan.

What changed:

- Attachment storage: can now be turned off for each user. When it is turned off, the approving user's role is not collected for that user from then on.
- The tax summary: now shows the batch number beside each entry. Entries made before 22.7 show a dash, because the value was not recorded.
- The import screen: now keeps the exchange rate used for 30 days. After that the value is removed, and the entry shows when it was removed.

- Fixed: a date typed as day/month could be read as month/day in exports. The format now follows the account's settings.
- Fixed: a date typed as day/month could be read as month/day in scheduled reports. The format now follows the account's settings.

## Quillstone Ledger 22.8

Quillstone Ledger 22.8 was released on 8 July 2023 for customers on the hosted plan.

What changed:

- The activity log: now shows the source file name beside each entry. Entries made before 22.8 show a dash, because the value was not recorded.
- Account settings: now includes a column for the batch number. The column is empty for records created before 22.8, because the batch number was not stored then.
- The approval queue: now includes a column for the import job number. The column is empty for records created before 22.8, because the import job number was not stored then.
- Exports: now accepts the payment terms as a filter. The filter applies to records created on or after 22.8.
- The search index: now includes a column for the attachment count. The column is empty for records created before 22.8, because the attachment count was not stored then.

- Fixed: a date typed as day/month could be read as month/day in user roles. The format now follows the account's settings.

Known issue: the search index does not yet support files larger than 25 MB. Larger files are refused with a message that names the limit.

## Quillstone Ledger 22.9

Quillstone Ledger 22.9 was released on 15 July 2023 for customers on the hosted plan.

What changed:

- Exports: now accepts the exchange rate used as a filter. The filter applies to records created on or after 22.9.
- Attachment storage: now accepts the payment terms as a filter. The filter applies to records created on or after 22.9.
- The import screen: now keeps the reviewer's comment count for 365 days. After that the value is removed, and the entry shows when it was removed.

- Fixed: in account settings, a month with no records showed a blank page. It now shows an empty table.
- Fixed: scheduled reports did not keep the sort order after a page reload. It now keeps it until the user changes it.

## Quillstone Ledger 23.0

Quillstone Ledger 23.0 was released on 22 July 2023 for customers on the hosted plan.

What changed:

- The audit view: faster on accounts with many records. In our own test on 22 July 2023, a 157,000-row run took 53 seconds on 22.9 and 13 seconds on 23.0. Your results depend on your data.
- Scheduled reports: can now be turned off for each user. When it is turned off, the batch number is not collected for that user from then on.
- The audit view: can now be turned off for each user. When it is turned off, the exchange rate used is not collected for that user from then on.
- User roles: now names the user who last changed the time zone of the entry. It does not show why the change was made, because the product does not ask for a reason.
- The audit view: now accepts the import job number as a filter. The filter applies to records created on or after 23.0.

- Fixed: the activity log did not keep the sort order after a page reload. It now keeps it until the user changes it.

## Quillstone Ledger 23.1

Quillstone Ledger 23.1 was released on 1 August 2023 for customers on the hosted plan.

What changed:

- Invoice matching: now includes a column for the attachment count. The column is empty for records created before 23.1, because the attachment count was not stored then.
- The tax summary: now names the user who last changed the reviewer's comment count. It does not show why the change was made, because the product does not ask for a reason.
- The search index: now names the user who last changed the time zone of the entry. It does not show why the change was made, because the product does not ask for a reason.

- Fixed: in user roles, a month with no records showed a blank page. It now shows an empty table.
- Fixed: the mobile view did not keep the sort order after a page reload. It now keeps it until the user changes it.

## Quillstone Ledger 23.2

Quillstone Ledger 23.2 was released on 8 August 2023 for customers on the hosted plan.

What changed:

- Report generation: now includes a column for the batch number. The column is empty for records created before 23.2, because the batch number was not stored then.
- Attachment storage: now shows the payment terms beside each entry. Entries made before 23.2 show a dash, because the value was not recorded.
- Bank feeds: now shows the batch number beside each entry. Entries made before 23.2 show a dash, because the value was not recorded.

- Fixed: totals in invoice matching were rounded before they were added, which could change the last digit. They are now added first and rounded once.
- Fixed: account settings did not keep the sort order after a page reload. It now keeps it until the user changes it.

## Quillstone Ledger 23.3

Quillstone Ledger 23.3 was released on 15 August 2023 for customers on the hosted plan.

What changed:

- Exports: now accepts the exchange rate used as a filter. The filter applies to records created on or after 23.3.
- Account settings: now keeps the cost centre for 365 days. After that the value is removed, and the entry shows when it was removed.
- The activity log: now includes a column for the batch number. The column is empty for records created before 23.3, because the batch number was not stored then.

- Fixed: in report generation, a month with no records showed a blank page. It now shows an empty table.
- Fixed: totals in the tax summary were rounded before they were added, which could change the last digit. They are now added first and rounded once.

Known issue: on slow connections bank feeds can take up to 10 seconds to load. We are working on it and will say here when it is fixed.

## Quillstone Ledger 23.4

Quillstone Ledger 23.4 was released on 22 August 2023 for customers on the hosted plan.

What changed:

- Currency conversion: now shows the approving user's role beside each entry. Entries made before 23.4 show a dash, because the value was not recorded.
- The search index: faster on accounts with many records. In our own test on 22 August 2023, a 125,000-row run took 51 seconds on 23.3 and 10 seconds on 23.4. Your results depend on your data.
- Attachment storage: now shows the original currency beside each entry. Entries made before 23.4 show a dash, because the value was not recorded.

- Fixed: the activity log did not keep the sort order after a page reload. It now keeps it until the user changes it.
- Fixed: a date typed as day/month could be read as month/day in the mobile view. The format now follows the account's settings.

Planned, not yet available: the cost centre in bank feeds. We expect to release it in the first quarter of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 23.5

Quillstone Ledger 23.5 was released on 1 September 2023 for customers on the hosted plan.

What changed:

- Exports: now names the user who last changed the reviewer's comment count. It does not show why the change was made, because the product does not ask for a reason.
- Scheduled reports: now includes a column for the device type. The column is empty for records created before 23.5, because the device type was not stored then.
- Scheduled reports: faster on accounts with many records. In our own test on 1 September 2023, a 140,000-row run took 85 seconds on 23.4 and 28 seconds on 23.5. Your results depend on your data.
- Exports: now includes a column for the approving user's role. The column is empty for records created before 23.5, because the approving user's role was not stored then.
- Attachment storage: now keeps the batch number for 180 days. After that the value is removed, and the entry shows when it was removed.

- Fixed: totals in user roles were rounded before they were added, which could change the last digit. They are now added first and rounded once.

Known issue: exports does not yet support files larger than 100 MB. Larger files are refused with a message that names the limit.

## Quillstone Ledger 23.6

Quillstone Ledger 23.6 was released on 8 September 2023 for customers on the hosted plan.

What changed:

- The import screen: now names the user who last changed the original currency. It does not show why the change was made, because the product does not ask for a reason.
- The approval queue: can now be turned off for each user. When it is turned off, the original currency is not collected for that user from then on.
- Account settings: now names the user who last changed the source file name. It does not show why the change was made, because the product does not ask for a reason.

- Fixed: totals in the activity log were rounded before they were added, which could change the last digit. They are now added first and rounded once.
- Fixed: in the activity log, a month with no records showed a blank page. It now shows an empty table.

## Quillstone Ledger 23.7

Quillstone Ledger 23.7 was released on 15 September 2023 for customers on the hosted plan.

What changed:

- The audit view: now shows the attachment count beside each entry. Entries made before 23.7 show a dash, because the value was not recorded.
- Bank feeds: now accepts the cost centre as a filter. The filter applies to records created on or after 23.7.
- The activity log: now includes a column for the cost centre. The column is empty for records created before 23.7, because the cost centre was not stored then.

- Fixed: totals in account settings were rounded before they were added, which could change the last digit. They are now added first and rounded once.

Known issue: on slow connections bank feeds can take up to 20 seconds to load. We are working on it and will say here when it is fixed.

Planned, not yet available: the source file name in the search index. We expect to release it in the first quarter of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 23.8

Quillstone Ledger 23.8 was released on 22 September 2023 for customers on the hosted plan.

What changed:

- The import screen: faster on accounts with many records. In our own test on 22 September 2023, a 124,000-row run took 40 seconds on 23.7 and 20 seconds on 23.8. Your results depend on your data.
- Invoice matching: now keeps the time zone of the entry for 365 days. After that the value is removed, and the entry shows when it was removed.
- The activity log: now keeps the import job number for 180 days. After that the value is removed, and the entry shows when it was removed.
- The import screen: now keeps the original currency for 30 days. After that the value is removed, and the entry shows when it was removed.

- Fixed: the mobile view did not keep the sort order after a page reload. It now keeps it until the user changes it.

Known issue: on slow connections exports can take up to 30 seconds to load. We are working on it and will say here when it is fixed.

## Quillstone Ledger 23.9

Quillstone Ledger 23.9 was released on 1 October 2023 for customers on the hosted plan.

What changed:

- The import screen: now accepts the approving user's role as a filter. The filter applies to records created on or after 23.9.
- The activity log: can now be turned off for each user. When it is turned off, the exchange rate used is not collected for that user from then on.
- The search index: now shows the payment terms beside each entry. Entries made before 23.9 show a dash, because the value was not recorded.
- Invoice matching: now shows the reviewer's comment count beside each entry. Entries made before 23.9 show a dash, because the value was not recorded.
- The audit view: now includes a column for the cost centre. The column is empty for records created before 23.9, because the cost centre was not stored then.

- Fixed: a date typed as day/month could be read as month/day in bank feeds. The format now follows the account's settings.
- Fixed: currency conversion did not keep the sort order after a page reload. It now keeps it until the user changes it.

## Quillstone Ledger 24.0

Quillstone Ledger 24.0 was released on 8 October 2023 for customers on the hosted plan.

What changed:

- The search index: now accepts the time zone of the entry as a filter. The filter applies to records created on or after 24.0.
- Exports: now includes a column for the approving user's role. The column is empty for records created before 24.0, because the approving user's role was not stored then.
- The approval queue: now includes a column for the exchange rate used. The column is empty for records created before 24.0, because the exchange rate used was not stored then.
- Report generation: now includes a column for the cost centre. The column is empty for records created before 24.0, because the cost centre was not stored then.

- Fixed: the tax summary did not keep the sort order after a page reload. It now keeps it until the user changes it.

Known issue: the audit view does not yet support files larger than 100 MB. Larger files are refused with a message that names the limit.

## Quillstone Ledger 24.1

Quillstone Ledger 24.1 was released on 15 October 2023 for customers on the hosted plan.

What changed:

- Currency conversion: now names the user who last changed the exchange rate used. It does not show why the change was made, because the product does not ask for a reason.
- Invoice matching: now includes a column for the payment terms. The column is empty for records created before 24.1, because the payment terms was not stored then.
- Exports: now includes a column for the attachment count. The column is empty for records created before 24.1, because the attachment count was not stored then.
- Report generation: now includes a column for the original currency. The column is empty for records created before 24.1, because the original currency was not stored then.

- Fixed: a date typed as day/month could be read as month/day in the tax summary. The format now follows the account's settings.
- Fixed: totals in scheduled reports were rounded before they were added, which could change the last digit. They are now added first and rounded once.

Known issue: invoice matching does not yet support files larger than 100 MB. Larger files are refused with a message that names the limit.

## Quillstone Ledger 24.2

Quillstone Ledger 24.2 was released on 22 October 2023 for customers on the hosted plan.

What changed:

- User roles: faster on accounts with many records. In our own test on 22 October 2023, a 80,000-row run took 73 seconds on 24.1 and 24 seconds on 24.2. Your results depend on your data.
- The search index: now names the user who last changed the time zone of the entry. It does not show why the change was made, because the product does not ask for a reason.
- Currency conversion: now includes a column for the device type. The column is empty for records created before 24.2, because the device type was not stored then.

- Fixed: in the activity log, a month with no records showed a blank page. It now shows an empty table.
- Fixed: exports did not keep the sort order after a page reload. It now keeps it until the user changes it.

Known issue: on slow connections currency conversion can take up to 10 seconds to load. We are working on it and will say here when it is fixed.

## Quillstone Ledger 24.3

Quillstone Ledger 24.3 was released on 1 November 2023 for customers on the hosted plan.

What changed:

- Report generation: now includes a column for the exchange rate used. The column is empty for records created before 24.3, because the exchange rate used was not stored then.
- Currency conversion: can now be turned off for each user. When it is turned off, the payment terms is not collected for that user from then on.
- Account settings: now includes a column for the attachment count. The column is empty for records created before 24.3, because the attachment count was not stored then.
- Report generation: now keeps the source file name for 365 days. After that the value is removed, and the entry shows when it was removed.

- Fixed: the import screen did not keep the sort order after a page reload. It now keeps it until the user changes it.

Planned, not yet available: the batch number in user roles. We expect to release it in the first quarter of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 24.4

Quillstone Ledger 24.4 was released on 8 November 2023 for customers on the hosted plan.

What changed:

- The activity log: now shows the attachment count beside each entry. Entries made before 24.4 show a dash, because the value was not recorded.
- Bank feeds: now names the user who last changed the cost centre. It does not show why the change was made, because the product does not ask for a reason.
- The approval queue: now includes a column for the exchange rate used. The column is empty for records created before 24.4, because the exchange rate used was not stored then.

- Fixed: totals in report generation were rounded before they were added, which could change the last digit. They are now added first and rounded once.

## Quillstone Ledger 24.5

Quillstone Ledger 24.5 was released on 15 November 2023 for customers on the hosted plan.

What changed:

- The activity log: now keeps the batch number for 365 days. After that the value is removed, and the entry shows when it was removed.
- Exports: can now be turned off for each user. When it is turned off, the approving user's role is not collected for that user from then on.
- Currency conversion: now accepts the source file name as a filter. The filter applies to records created on or after 24.5.

- Fixed: totals in the search index were rounded before they were added, which could change the last digit. They are now added first and rounded once.

## Quillstone Ledger 24.6

Quillstone Ledger 24.6 was released on 22 November 2023 for customers on the hosted plan.

What changed:

- Attachment storage: now accepts the device type as a filter. The filter applies to records created on or after 24.6.
- The audit view: now shows the batch number beside each entry. Entries made before 24.6 show a dash, because the value was not recorded.
- Bank feeds: faster on accounts with many records. In our own test on 22 November 2023, a 70,000-row run took 49 seconds on 24.5 and 16 seconds on 24.6. Your results depend on your data.

- Fixed: a date typed as day/month could be read as month/day in bank feeds. The format now follows the account's settings.

## Quillstone Ledger 24.7

Quillstone Ledger 24.7 was released on 1 December 2023 for customers on the hosted plan.

What changed:

- Exports: can now be turned off for each user. When it is turned off, the import job number is not collected for that user from then on.
- Attachment storage: can now be turned off for each user. When it is turned off, the exchange rate used is not collected for that user from then on.
- The activity log: now accepts the source file name as a filter. The filter applies to records created on or after 24.7.

- Fixed: totals in the mobile view were rounded before they were added, which could change the last digit. They are now added first and rounded once.
- Fixed: a date typed as day/month could be read as month/day in exports. The format now follows the account's settings.

Known issue: on slow connections attachment storage can take up to 30 seconds to load. We are working on it and will say here when it is fixed.

Planned, not yet available: the attachment count in bank feeds. We expect to release it in the first quarter of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 24.8

Quillstone Ledger 24.8 was released on 8 December 2023 for customers on the hosted plan.

What changed:

- Currency conversion: now accepts the batch number as a filter. The filter applies to records created on or after 24.8.
- Invoice matching: now accepts the original currency as a filter. The filter applies to records created on or after 24.8.
- Currency conversion: faster on accounts with many records. In our own test on 8 December 2023, a 65,000-row run took 65 seconds on 24.7 and 13 seconds on 24.8. Your results depend on your data.
- Bank feeds: now includes a column for the device type. The column is empty for records created before 24.8, because the device type was not stored then.

- Fixed: the search index did not keep the sort order after a page reload. It now keeps it until the user changes it.

Known issue: currency conversion does not yet support files larger than 50 MB. Larger files are refused with a message that names the limit.

## Quillstone Ledger 24.9

Quillstone Ledger 24.9 was released on 15 December 2023 for customers on the hosted plan.

What changed:

- Account settings: now keeps the source file name for 365 days. After that the value is removed, and the entry shows when it was removed.
- The search index: now keeps the reviewer's comment count for 90 days. After that the value is removed, and the entry shows when it was removed.
- The tax summary: now names the user who last changed the exchange rate used. It does not show why the change was made, because the product does not ask for a reason.
- The activity log: now keeps the reviewer's comment count for 90 days. After that the value is removed, and the entry shows when it was removed.
- Exports: now keeps the time zone of the entry for 30 days. After that the value is removed, and the entry shows when it was removed.

- Fixed: totals in the activity log were rounded before they were added, which could change the last digit. They are now added first and rounded once.
- Fixed: the mobile view did not keep the sort order after a page reload. It now keeps it until the user changes it.

Known issue: on slow connections the tax summary can take up to 30 seconds to load. We are working on it and will say here when it is fixed.

## Quillstone Ledger 25.0

Quillstone Ledger 25.0 was released on 22 December 2023 for customers on the hosted plan.

What changed:

- The mobile view: now accepts the time zone of the entry as a filter. The filter applies to records created on or after 25.0.
- The tax summary: can now be turned off for each user. When it is turned off, the device type is not collected for that user from then on.
- Invoice matching: now shows the original currency beside each entry. Entries made before 25.0 show a dash, because the value was not recorded.

- Fixed: a date typed as day/month could be read as month/day in the approval queue. The format now follows the account's settings.

Known issue: on slow connections the audit view can take up to 20 seconds to load. We are working on it and will say here when it is fixed.

## Quillstone Ledger 25.1

Quillstone Ledger 25.1 was released on 1 January 2024 for customers on the hosted plan.

What changed:

- Bank feeds: now keeps the payment terms for 30 days. After that the value is removed, and the entry shows when it was removed.
- Exports: now includes a column for the source file name. The column is empty for records created before 25.1, because the source file name was not stored then.
- The audit view: now includes a column for the cost centre. The column is empty for records created before 25.1, because the cost centre was not stored then.
- The tax summary: now keeps the time zone of the entry for 365 days. After that the value is removed, and the entry shows when it was removed.
- The audit view: now accepts the approving user's role as a filter. The filter applies to records created on or after 25.1.

- Fixed: totals in exports were rounded before they were added, which could change the last digit. They are now added first and rounded once.
- Fixed: in the approval queue, a month with no records showed a blank page. It now shows an empty table.

## Quillstone Ledger 25.2

Quillstone Ledger 25.2 was released on 8 January 2024 for customers on the hosted plan.

What changed:

- Attachment storage: now includes a column for the device type. The column is empty for records created before 25.2, because the device type was not stored then.
- Scheduled reports: can now be turned off for each user. When it is turned off, the original currency is not collected for that user from then on.
- Scheduled reports: now keeps the attachment count for 180 days. After that the value is removed, and the entry shows when it was removed.
- Currency conversion: now keeps the reviewer's comment count for 30 days. After that the value is removed, and the entry shows when it was removed.
- The search index: now keeps the cost centre for 180 days. After that the value is removed, and the entry shows when it was removed.

- Fixed: in bank feeds, a month with no records showed a blank page. It now shows an empty table.
- Fixed: totals in exports were rounded before they were added, which could change the last digit. They are now added first and rounded once.

Known issue: bank feeds does not yet support files larger than 100 MB. Larger files are refused with a message that names the limit.

## Quillstone Ledger 25.3

Quillstone Ledger 25.3 was released on 15 January 2024 for customers on the hosted plan.

What changed:

- Bank feeds: now shows the source file name beside each entry. Entries made before 25.3 show a dash, because the value was not recorded.
- The approval queue: now includes a column for the time zone of the entry. The column is empty for records created before 25.3, because the time zone of the entry was not stored then.
- Scheduled reports: now accepts the attachment count as a filter. The filter applies to records created on or after 25.3.
- Invoice matching: can now be turned off for each user. When it is turned off, the attachment count is not collected for that user from then on.

- Fixed: the mobile view did not keep the sort order after a page reload. It now keeps it until the user changes it.

## Quillstone Ledger 25.4

Quillstone Ledger 25.4 was released on 22 January 2024 for customers on the hosted plan.

What changed:

- Bank feeds: faster on accounts with many records. In our own test on 22 January 2024, a 151,000-row run took 89 seconds on 25.3 and 17 seconds on 25.4. Your results depend on your data.
- Invoice matching: now accepts the reviewer's comment count as a filter. The filter applies to records created on or after 25.4.
- Attachment storage: now shows the payment terms beside each entry. Entries made before 25.4 show a dash, because the value was not recorded.
- The activity log: can now be turned off for each user. When it is turned off, the exchange rate used is not collected for that user from then on.
- The mobile view: now accepts the cost centre as a filter. The filter applies to records created on or after 25.4.

- Fixed: a date typed as day/month could be read as month/day in attachment storage. The format now follows the account's settings.
- Fixed: totals in the import screen were rounded before they were added, which could change the last digit. They are now added first and rounded once.

## Quillstone Ledger 25.5

Quillstone Ledger 25.5 was released on 1 February 2024 for customers on the hosted plan.

What changed:

- The approval queue: now shows the original currency beside each entry. Entries made before 25.5 show a dash, because the value was not recorded.
- The search index: now accepts the original currency as a filter. The filter applies to records created on or after 25.5.
- The mobile view: now keeps the attachment count for 365 days. After that the value is removed, and the entry shows when it was removed.
- The activity log: now keeps the cost centre for 90 days. After that the value is removed, and the entry shows when it was removed.
- Account settings: faster on accounts with many records. In our own test on 1 February 2024, a 40,000-row run took 32 seconds on 25.4 and 16 seconds on 25.5. Your results depend on your data.

- Fixed: a date typed as day/month could be read as month/day in bank feeds. The format now follows the account's settings.

Known issue: report generation does not yet support files larger than 100 MB. Larger files are refused with a message that names the limit.

Planned, not yet available: the attachment count in exports. We expect to release it in the second half of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 25.6

Quillstone Ledger 25.6 was released on 8 February 2024 for customers on the hosted plan.

What changed:

- The search index: now keeps the original currency for 180 days. After that the value is removed, and the entry shows when it was removed.
- User roles: now keeps the reviewer's comment count for 90 days. After that the value is removed, and the entry shows when it was removed.
- The import screen: now names the user who last changed the time zone of the entry. It does not show why the change was made, because the product does not ask for a reason.
- The search index: now keeps the attachment count for 90 days. After that the value is removed, and the entry shows when it was removed.
- Currency conversion: now accepts the original currency as a filter. The filter applies to records created on or after 25.6.

- Fixed: a date typed as day/month could be read as month/day in the mobile view. The format now follows the account's settings.

Known issue: on slow connections exports can take up to 20 seconds to load. We are working on it and will say here when it is fixed.

Planned, not yet available: the original currency in scheduled reports. We expect to release it in the first quarter of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 25.7

Quillstone Ledger 25.7 was released on 15 February 2024 for customers on the hosted plan.

What changed:

- Currency conversion: now names the user who last changed the import job number. It does not show why the change was made, because the product does not ask for a reason.
- The mobile view: now accepts the original currency as a filter. The filter applies to records created on or after 25.7.
- The approval queue: now names the user who last changed the time zone of the entry. It does not show why the change was made, because the product does not ask for a reason.
- Scheduled reports: now shows the batch number beside each entry. Entries made before 25.7 show a dash, because the value was not recorded.

- Fixed: the tax summary did not keep the sort order after a page reload. It now keeps it until the user changes it.

## Quillstone Ledger 25.8

Quillstone Ledger 25.8 was released on 22 February 2024 for customers on the hosted plan.

What changed:

- Attachment storage: now shows the time zone of the entry beside each entry. Entries made before 25.8 show a dash, because the value was not recorded.
- User roles: faster on accounts with many records. In our own test on 22 February 2024, a 91,000-row run took 90 seconds on 25.7 and 22 seconds on 25.8. Your results depend on your data.
- Attachment storage: faster on accounts with many records. In our own test on 22 February 2024, a 29,000-row run took 60 seconds on 25.7 and 15 seconds on 25.8. Your results depend on your data.
- Account settings: now shows the source file name beside each entry. Entries made before 25.8 show a dash, because the value was not recorded.

- Fixed: in the audit view, a month with no records showed a blank page. It now shows an empty table.
- Fixed: the audit view did not keep the sort order after a page reload. It now keeps it until the user changes it.

## Quillstone Ledger 25.9

Quillstone Ledger 25.9 was released on 1 March 2024 for customers on the hosted plan.

What changed:

- Exports: now shows the exchange rate used beside each entry. Entries made before 25.9 show a dash, because the value was not recorded.
- Account settings: faster on accounts with many records. In our own test on 1 March 2024, a 179,000-row run took 71 seconds on 25.8 and 23 seconds on 25.9. Your results depend on your data.
- Bank feeds: now includes a column for the approving user's role. The column is empty for records created before 25.9, because the approving user's role was not stored then.
- The search index: now includes a column for the exchange rate used. The column is empty for records created before 25.9, because the exchange rate used was not stored then.

- Fixed: account settings did not keep the sort order after a page reload. It now keeps it until the user changes it.
- Fixed: totals in invoice matching were rounded before they were added, which could change the last digit. They are now added first and rounded once.

Known issue: the mobile view does not yet support files larger than 25 MB. Larger files are refused with a message that names the limit.

## Quillstone Ledger 26.0

Quillstone Ledger 26.0 was released on 8 March 2024 for customers on the hosted plan.

What changed:

- The approval queue: faster on accounts with many records. In our own test on 8 March 2024, a 83,000-row run took 45 seconds on 25.9 and 9 seconds on 26.0. Your results depend on your data.
- Scheduled reports: faster on accounts with many records. In our own test on 8 March 2024, a 17,000-row run took 23 seconds on 25.9 and 11 seconds on 26.0. Your results depend on your data.
- The mobile view: now names the user who last changed the reviewer's comment count. It does not show why the change was made, because the product does not ask for a reason.
- User roles: faster on accounts with many records. In our own test on 8 March 2024, a 150,000-row run took 47 seconds on 25.9 and 15 seconds on 26.0. Your results depend on your data.
- The import screen: can now be turned off for each user. When it is turned off, the time zone of the entry is not collected for that user from then on.

- Fixed: in exports, a month with no records showed a blank page. It now shows an empty table.

## Quillstone Ledger 26.1

Quillstone Ledger 26.1 was released on 15 March 2024 for customers on the hosted plan.

What changed:

- Attachment storage: now keeps the source file name for 365 days. After that the value is removed, and the entry shows when it was removed.
- Invoice matching: now includes a column for the attachment count. The column is empty for records created before 26.1, because the attachment count was not stored then.
- The tax summary: now keeps the approving user's role for 365 days. After that the value is removed, and the entry shows when it was removed.
- The search index: now shows the source file name beside each entry. Entries made before 26.1 show a dash, because the value was not recorded.

- Fixed: in the import screen, a month with no records showed a blank page. It now shows an empty table.
- Fixed: the import screen did not keep the sort order after a page reload. It now keeps it until the user changes it.

Known issue: on slow connections currency conversion can take up to 10 seconds to load. We are working on it and will say here when it is fixed.

## Quillstone Ledger 26.2

Quillstone Ledger 26.2 was released on 22 March 2024 for customers on the hosted plan.

What changed:

- The tax summary: now shows the device type beside each entry. Entries made before 26.2 show a dash, because the value was not recorded.
- The audit view: now includes a column for the approving user's role. The column is empty for records created before 26.2, because the approving user's role was not stored then.
- The mobile view: can now be turned off for each user. When it is turned off, the reviewer's comment count is not collected for that user from then on.

- Fixed: in the search index, a month with no records showed a blank page. It now shows an empty table.
- Fixed: account settings did not keep the sort order after a page reload. It now keeps it until the user changes it.

Planned, not yet available: the source file name in exports. We expect to release it in the second half of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 26.3

Quillstone Ledger 26.3 was released on 1 April 2024 for customers on the hosted plan.

What changed:

- The mobile view: now accepts the time zone of the entry as a filter. The filter applies to records created on or after 26.3.
- The mobile view: now keeps the batch number for 30 days. After that the value is removed, and the entry shows when it was removed.
- Exports: now names the user who last changed the reviewer's comment count. It does not show why the change was made, because the product does not ask for a reason.

- Fixed: a date typed as day/month could be read as month/day in the activity log. The format now follows the account's settings.

## Quillstone Ledger 26.4

Quillstone Ledger 26.4 was released on 8 April 2024 for customers on the hosted plan.

What changed:

- Attachment storage: can now be turned off for each user. When it is turned off, the cost centre is not collected for that user from then on.
- Bank feeds: faster on accounts with many records. In our own test on 8 April 2024, a 100,000-row run took 51 seconds on 26.3 and 17 seconds on 26.4. Your results depend on your data.
- Bank feeds: now accepts the import job number as a filter. The filter applies to records created on or after 26.4.
- The approval queue: now accepts the payment terms as a filter. The filter applies to records created on or after 26.4.
- User roles: now includes a column for the original currency. The column is empty for records created before 26.4, because the original currency was not stored then.

- Fixed: totals in exports were rounded before they were added, which could change the last digit. They are now added first and rounded once.
- Fixed: scheduled reports did not keep the sort order after a page reload. It now keeps it until the user changes it.

Known issue: exports does not yet support files larger than 100 MB. Larger files are refused with a message that names the limit.

## Quillstone Ledger 26.5

Quillstone Ledger 26.5 was released on 15 April 2024 for customers on the hosted plan.

What changed:

- The import screen: now accepts the source file name as a filter. The filter applies to records created on or after 26.5.
- The activity log: now names the user who last changed the source file name. It does not show why the change was made, because the product does not ask for a reason.
- The approval queue: now keeps the time zone of the entry for 90 days. After that the value is removed, and the entry shows when it was removed.
- Account settings: faster on accounts with many records. In our own test on 15 April 2024, a 122,000-row run took 31 seconds on 26.4 and 7 seconds on 26.5. Your results depend on your data.

- Fixed: in user roles, a month with no records showed a blank page. It now shows an empty table.

Known issue: on slow connections the search index can take up to 30 seconds to load. We are working on it and will say here when it is fixed.

## Quillstone Ledger 26.6

Quillstone Ledger 26.6 was released on 22 April 2024 for customers on the hosted plan.

What changed:

- Account settings: now includes a column for the source file name. The column is empty for records created before 26.6, because the source file name was not stored then.
- Report generation: now accepts the source file name as a filter. The filter applies to records created on or after 26.6.
- Bank feeds: now accepts the original currency as a filter. The filter applies to records created on or after 26.6.
- Exports: now includes a column for the exchange rate used. The column is empty for records created before 26.6, because the exchange rate used was not stored then.

- Fixed: in the import screen, a month with no records showed a blank page. It now shows an empty table.

Known issue: on slow connections the search index can take up to 30 seconds to load. We are working on it and will say here when it is fixed.

## Quillstone Ledger 26.7

Quillstone Ledger 26.7 was released on 1 May 2024 for customers on the hosted plan.

What changed:

- Exports: now accepts the payment terms as a filter. The filter applies to records created on or after 26.7.
- Invoice matching: now accepts the time zone of the entry as a filter. The filter applies to records created on or after 26.7.
- The search index: now shows the approving user's role beside each entry. Entries made before 26.7 show a dash, because the value was not recorded.

- Fixed: in the import screen, a month with no records showed a blank page. It now shows an empty table.
- Fixed: the import screen did not keep the sort order after a page reload. It now keeps it until the user changes it.

Planned, not yet available: the batch number in attachment storage. We expect to release it in the second quarter of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 26.8

Quillstone Ledger 26.8 was released on 8 May 2024 for customers on the hosted plan.

What changed:

- Attachment storage: now keeps the batch number for 30 days. After that the value is removed, and the entry shows when it was removed.
- Report generation: faster on accounts with many records. In our own test on 8 May 2024, a 89,000-row run took 49 seconds on 26.7 and 24 seconds on 26.8. Your results depend on your data.
- The tax summary: now includes a column for the reviewer's comment count. The column is empty for records created before 26.8, because the reviewer's comment count was not stored then.
- Report generation: now keeps the source file name for 90 days. After that the value is removed, and the entry shows when it was removed.

- Fixed: report generation did not keep the sort order after a page reload. It now keeps it until the user changes it.
- Fixed: a date typed as day/month could be read as month/day in exports. The format now follows the account's settings.

Known issue: on slow connections bank feeds can take up to 20 seconds to load. We are working on it and will say here when it is fixed.

## Quillstone Ledger 26.9

Quillstone Ledger 26.9 was released on 15 May 2024 for customers on the hosted plan.

What changed:

- Attachment storage: now shows the cost centre beside each entry. Entries made before 26.9 show a dash, because the value was not recorded.
- The tax summary: now shows the source file name beside each entry. Entries made before 26.9 show a dash, because the value was not recorded.
- The audit view: can now be turned off for each user. When it is turned off, the import job number is not collected for that user from then on.
- Exports: now includes a column for the batch number. The column is empty for records created before 26.9, because the batch number was not stored then.

- Fixed: in exports, a month with no records showed a blank page. It now shows an empty table.
- Fixed: in scheduled reports, a month with no records showed a blank page. It now shows an empty table.

## Quillstone Ledger 27.0

Quillstone Ledger 27.0 was released on 22 May 2024 for customers on the hosted plan.

What changed:

- Scheduled reports: now names the user who last changed the reviewer's comment count. It does not show why the change was made, because the product does not ask for a reason.
- User roles: now accepts the attachment count as a filter. The filter applies to records created on or after 27.0.
- Currency conversion: now accepts the import job number as a filter. The filter applies to records created on or after 27.0.
- Currency conversion: now keeps the original currency for 30 days. After that the value is removed, and the entry shows when it was removed.
- The mobile view: faster on accounts with many records. In our own test on 22 May 2024, a 99,000-row run took 67 seconds on 26.9 and 13 seconds on 27.0. Your results depend on your data.

- Fixed: totals in the tax summary were rounded before they were added, which could change the last digit. They are now added first and rounded once.
- Fixed: in currency conversion, a month with no records showed a blank page. It now shows an empty table.

## Quillstone Ledger 27.1

Quillstone Ledger 27.1 was released on 1 June 2024 for customers on the hosted plan.

What changed:

- Account settings: now names the user who last changed the source file name. It does not show why the change was made, because the product does not ask for a reason.
- Report generation: now accepts the payment terms as a filter. The filter applies to records created on or after 27.1.
- Bank feeds: now names the user who last changed the import job number. It does not show why the change was made, because the product does not ask for a reason.
- Exports: now shows the time zone of the entry beside each entry. Entries made before 27.1 show a dash, because the value was not recorded.

- Fixed: totals in attachment storage were rounded before they were added, which could change the last digit. They are now added first and rounded once.

Known issue: invoice matching does not yet support files larger than 25 MB. Larger files are refused with a message that names the limit.

## Quillstone Ledger 27.2

Quillstone Ledger 27.2 was released on 8 June 2024 for customers on the hosted plan.

What changed:

- Scheduled reports: now keeps the payment terms for 365 days. After that the value is removed, and the entry shows when it was removed.
- The search index: now names the user who last changed the time zone of the entry. It does not show why the change was made, because the product does not ask for a reason.
- Bank feeds: now shows the attachment count beside each entry. Entries made before 27.2 show a dash, because the value was not recorded.
- Invoice matching: can now be turned off for each user. When it is turned off, the cost centre is not collected for that user from then on.

- Fixed: in scheduled reports, a month with no records showed a blank page. It now shows an empty table.
- Fixed: totals in exports were rounded before they were added, which could change the last digit. They are now added first and rounded once.

## Quillstone Ledger 27.3

Quillstone Ledger 27.3 was released on 15 June 2024 for customers on the hosted plan.

What changed:

- The activity log: faster on accounts with many records. In our own test on 15 June 2024, a 128,000-row run took 78 seconds on 27.2 and 26 seconds on 27.3. Your results depend on your data.
- The mobile view: now accepts the import job number as a filter. The filter applies to records created on or after 27.3.
- Invoice matching: now accepts the approving user's role as a filter. The filter applies to records created on or after 27.3.
- User roles: can now be turned off for each user. When it is turned off, the batch number is not collected for that user from then on.
- The tax summary: now accepts the attachment count as a filter. The filter applies to records created on or after 27.3.

- Fixed: the import screen did not keep the sort order after a page reload. It now keeps it until the user changes it.

## Quillstone Ledger 27.4

Quillstone Ledger 27.4 was released on 22 June 2024 for customers on the hosted plan.

What changed:

- Report generation: now includes a column for the reviewer's comment count. The column is empty for records created before 27.4, because the reviewer's comment count was not stored then.
- Currency conversion: can now be turned off for each user. When it is turned off, the batch number is not collected for that user from then on.
- Scheduled reports: now includes a column for the device type. The column is empty for records created before 27.4, because the device type was not stored then.

- Fixed: the approval queue did not keep the sort order after a page reload. It now keeps it until the user changes it.

## Quillstone Ledger 27.5

Quillstone Ledger 27.5 was released on 1 July 2024 for customers on the hosted plan.

What changed:

- Exports: faster on accounts with many records. In our own test on 1 July 2024, a 34,000-row run took 57 seconds on 27.4 and 11 seconds on 27.5. Your results depend on your data.
- Report generation: now includes a column for the time zone of the entry. The column is empty for records created before 27.5, because the time zone of the entry was not stored then.
- Report generation: now keeps the approving user's role for 180 days. After that the value is removed, and the entry shows when it was removed.
- The search index: now accepts the time zone of the entry as a filter. The filter applies to records created on or after 27.5.

- Fixed: the tax summary did not keep the sort order after a page reload. It now keeps it until the user changes it.
- Fixed: the activity log did not keep the sort order after a page reload. It now keeps it until the user changes it.

## Quillstone Ledger 27.6

Quillstone Ledger 27.6 was released on 8 July 2024 for customers on the hosted plan.

What changed:

- The audit view: can now be turned off for each user. When it is turned off, the payment terms is not collected for that user from then on.
- Account settings: now shows the exchange rate used beside each entry. Entries made before 27.6 show a dash, because the value was not recorded.
- The mobile view: can now be turned off for each user. When it is turned off, the batch number is not collected for that user from then on.

- Fixed: a date typed as day/month could be read as month/day in the approval queue. The format now follows the account's settings.
- Fixed: in the search index, a month with no records showed a blank page. It now shows an empty table.

## Quillstone Ledger 27.7

Quillstone Ledger 27.7 was released on 15 July 2024 for customers on the hosted plan.

What changed:

- The tax summary: now shows the reviewer's comment count beside each entry. Entries made before 27.7 show a dash, because the value was not recorded.
- The mobile view: now includes a column for the exchange rate used. The column is empty for records created before 27.7, because the exchange rate used was not stored then.
- Bank feeds: can now be turned off for each user. When it is turned off, the attachment count is not collected for that user from then on.

- Fixed: report generation did not keep the sort order after a page reload. It now keeps it until the user changes it.
- Fixed: totals in the mobile view were rounded before they were added, which could change the last digit. They are now added first and rounded once.

Known issue: on slow connections account settings can take up to 10 seconds to load. We are working on it and will say here when it is fixed.

## Quillstone Ledger 27.8

Quillstone Ledger 27.8 was released on 22 July 2024 for customers on the hosted plan.

What changed:

- The tax summary: now accepts the attachment count as a filter. The filter applies to records created on or after 27.8.
- The approval queue: now accepts the approving user's role as a filter. The filter applies to records created on or after 27.8.
- Exports: can now be turned off for each user. When it is turned off, the import job number is not collected for that user from then on.

- Fixed: totals in the audit view were rounded before they were added, which could change the last digit. They are now added first and rounded once.
- Fixed: a date typed as day/month could be read as month/day in currency conversion. The format now follows the account's settings.

## Quillstone Ledger 27.9

Quillstone Ledger 27.9 was released on 1 August 2024 for customers on the hosted plan.

What changed:

- Exports: faster on accounts with many records. In our own test on 1 August 2024, a 29,000-row run took 55 seconds on 27.8 and 11 seconds on 27.9. Your results depend on your data.
- User roles: now names the user who last changed the payment terms. It does not show why the change was made, because the product does not ask for a reason.
- The import screen: now includes a column for the cost centre. The column is empty for records created before 27.9, because the cost centre was not stored then.
- Account settings: now includes a column for the payment terms. The column is empty for records created before 27.9, because the payment terms was not stored then.

- Fixed: a date typed as day/month could be read as month/day in the audit view. The format now follows the account's settings.

Known issue: on slow connections exports can take up to 10 seconds to load. We are working on it and will say here when it is fixed.

## Quillstone Ledger 28.0

Quillstone Ledger 28.0 was released on 8 August 2024 for customers on the hosted plan.

What changed:

- Bank feeds: faster on accounts with many records. In our own test on 8 August 2024, a 63,000-row run took 41 seconds on 27.9 and 8 seconds on 28.0. Your results depend on your data.
- Report generation: faster on accounts with many records. In our own test on 8 August 2024, a 158,000-row run took 58 seconds on 27.9 and 14 seconds on 28.0. Your results depend on your data.
- The mobile view: now includes a column for the batch number. The column is empty for records created before 28.0, because the batch number was not stored then.
- The approval queue: now includes a column for the source file name. The column is empty for records created before 28.0, because the source file name was not stored then.
- User roles: faster on accounts with many records. In our own test on 8 August 2024, a 15,000-row run took 43 seconds on 27.9 and 21 seconds on 28.0. Your results depend on your data.

- Fixed: a date typed as day/month could be read as month/day in the activity log. The format now follows the account's settings.

Planned, not yet available: the source file name in the activity log. We expect to release it in the first quarter of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 28.1

Quillstone Ledger 28.1 was released on 15 August 2024 for customers on the hosted plan.

What changed:

- User roles: now names the user who last changed the batch number. It does not show why the change was made, because the product does not ask for a reason.
- The tax summary: can now be turned off for each user. When it is turned off, the attachment count is not collected for that user from then on.
- Account settings: can now be turned off for each user. When it is turned off, the approving user's role is not collected for that user from then on.
- Account settings: can now be turned off for each user. When it is turned off, the original currency is not collected for that user from then on.

- Fixed: the approval queue did not keep the sort order after a page reload. It now keeps it until the user changes it.

Known issue: on slow connections attachment storage can take up to 30 seconds to load. We are working on it and will say here when it is fixed.

## Quillstone Ledger 28.2

Quillstone Ledger 28.2 was released on 22 August 2024 for customers on the hosted plan.

What changed:

- Bank feeds: can now be turned off for each user. When it is turned off, the time zone of the entry is not collected for that user from then on.
- The mobile view: now keeps the batch number for 30 days. After that the value is removed, and the entry shows when it was removed.
- The search index: faster on accounts with many records. In our own test on 22 August 2024, a 75,000-row run took 76 seconds on 28.1 and 19 seconds on 28.2. Your results depend on your data.
- The import screen: can now be turned off for each user. When it is turned off, the attachment count is not collected for that user from then on.

- Fixed: the activity log did not keep the sort order after a page reload. It now keeps it until the user changes it.

Known issue: invoice matching does not yet support files larger than 100 MB. Larger files are refused with a message that names the limit.

## Quillstone Ledger 28.3

Quillstone Ledger 28.3 was released on 1 September 2024 for customers on the hosted plan.

What changed:

- The tax summary: now shows the source file name beside each entry. Entries made before 28.3 show a dash, because the value was not recorded.
- User roles: now includes a column for the exchange rate used. The column is empty for records created before 28.3, because the exchange rate used was not stored then.
- Invoice matching: now shows the source file name beside each entry. Entries made before 28.3 show a dash, because the value was not recorded.
- Currency conversion: can now be turned off for each user. When it is turned off, the payment terms is not collected for that user from then on.
- The tax summary: now shows the cost centre beside each entry. Entries made before 28.3 show a dash, because the value was not recorded.

- Fixed: in report generation, a month with no records showed a blank page. It now shows an empty table.

Known issue: on slow connections report generation can take up to 20 seconds to load. We are working on it and will say here when it is fixed.

## Quillstone Ledger 28.4

Quillstone Ledger 28.4 was released on 8 September 2024 for customers on the hosted plan.

What changed:

- Exports: now includes a column for the batch number. The column is empty for records created before 28.4, because the batch number was not stored then.
- The activity log: now shows the attachment count beside each entry. Entries made before 28.4 show a dash, because the value was not recorded.
- The tax summary: now includes a column for the import job number. The column is empty for records created before 28.4, because the import job number was not stored then.

- Fixed: totals in scheduled reports were rounded before they were added, which could change the last digit. They are now added first and rounded once.

Known issue: on slow connections exports can take up to 30 seconds to load. We are working on it and will say here when it is fixed.

## Quillstone Ledger 28.5

Quillstone Ledger 28.5 was released on 15 September 2024 for customers on the hosted plan.

What changed:

- The audit view: now names the user who last changed the attachment count. It does not show why the change was made, because the product does not ask for a reason.
- Report generation: now names the user who last changed the cost centre. It does not show why the change was made, because the product does not ask for a reason.
- Report generation: now shows the payment terms beside each entry. Entries made before 28.5 show a dash, because the value was not recorded.

- Fixed: a date typed as day/month could be read as month/day in the mobile view. The format now follows the account's settings.

Known issue: on slow connections the search index can take up to 10 seconds to load. We are working on it and will say here when it is fixed.

## Quillstone Ledger 28.6

Quillstone Ledger 28.6 was released on 22 September 2024 for customers on the hosted plan.

What changed:

- The tax summary: now includes a column for the batch number. The column is empty for records created before 28.6, because the batch number was not stored then.
- Scheduled reports: now names the user who last changed the exchange rate used. It does not show why the change was made, because the product does not ask for a reason.
- The audit view: now names the user who last changed the reviewer's comment count. It does not show why the change was made, because the product does not ask for a reason.
- The mobile view: now names the user who last changed the import job number. It does not show why the change was made, because the product does not ask for a reason.
- Account settings: faster on accounts with many records. In our own test on 22 September 2024, a 121,000-row run took 55 seconds on 28.5 and 11 seconds on 28.6. Your results depend on your data.

- Fixed: in exports, a month with no records showed a blank page. It now shows an empty table.
- Fixed: in the search index, a month with no records showed a blank page. It now shows an empty table.

Known issue: on slow connections bank feeds can take up to 20 seconds to load. We are working on it and will say here when it is fixed.

Planned, not yet available: the payment terms in currency conversion. We expect to release it in the first quarter of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 28.7

Quillstone Ledger 28.7 was released on 1 October 2024 for customers on the hosted plan.

What changed:

- Report generation: now shows the approving user's role beside each entry. Entries made before 28.7 show a dash, because the value was not recorded.
- Account settings: faster on accounts with many records. In our own test on 1 October 2024, a 150,000-row run took 59 seconds on 28.6 and 11 seconds on 28.7. Your results depend on your data.
- The tax summary: can now be turned off for each user. When it is turned off, the source file name is not collected for that user from then on.

- Fixed: a date typed as day/month could be read as month/day in invoice matching. The format now follows the account's settings.
- Fixed: totals in scheduled reports were rounded before they were added, which could change the last digit. They are now added first and rounded once.

Planned, not yet available: the device type in account settings. We expect to release it in the second quarter of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 28.8

Quillstone Ledger 28.8 was released on 8 October 2024 for customers on the hosted plan.

What changed:

- Attachment storage: now shows the batch number beside each entry. Entries made before 28.8 show a dash, because the value was not recorded.
- Report generation: can now be turned off for each user. When it is turned off, the exchange rate used is not collected for that user from then on.
- The audit view: faster on accounts with many records. In our own test on 8 October 2024, a 44,000-row run took 74 seconds on 28.7 and 14 seconds on 28.8. Your results depend on your data.
- The mobile view: can now be turned off for each user. When it is turned off, the attachment count is not collected for that user from then on.
- Account settings: now includes a column for the payment terms. The column is empty for records created before 28.8, because the payment terms was not stored then.

- Fixed: exports did not keep the sort order after a page reload. It now keeps it until the user changes it.

Planned, not yet available: the attachment count in exports. We expect to release it in the second quarter of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 28.9

Quillstone Ledger 28.9 was released on 15 October 2024 for customers on the hosted plan.

What changed:

- Invoice matching: now shows the time zone of the entry beside each entry. Entries made before 28.9 show a dash, because the value was not recorded.
- The audit view: now keeps the import job number for 90 days. After that the value is removed, and the entry shows when it was removed.
- User roles: now accepts the source file name as a filter. The filter applies to records created on or after 28.9.

- Fixed: the search index did not keep the sort order after a page reload. It now keeps it until the user changes it.
- Fixed: in bank feeds, a month with no records showed a blank page. It now shows an empty table.

Known issue: bank feeds does not yet support files larger than 100 MB. Larger files are refused with a message that names the limit.

Planned, not yet available: the attachment count in user roles. We expect to release it in the first quarter of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 29.0

Quillstone Ledger 29.0 was released on 22 October 2024 for customers on the hosted plan.

What changed:

- Currency conversion: now includes a column for the device type. The column is empty for records created before 29.0, because the device type was not stored then.
- Invoice matching: now includes a column for the time zone of the entry. The column is empty for records created before 29.0, because the time zone of the entry was not stored then.
- The audit view: now accepts the attachment count as a filter. The filter applies to records created on or after 29.0.
- Scheduled reports: now names the user who last changed the device type. It does not show why the change was made, because the product does not ask for a reason.

- Fixed: totals in invoice matching were rounded before they were added, which could change the last digit. They are now added first and rounded once.
- Fixed: a date typed as day/month could be read as month/day in the import screen. The format now follows the account's settings.

Known issue: the activity log does not yet support files larger than 25 MB. Larger files are refused with a message that names the limit.

Planned, not yet available: the batch number in the import screen. We expect to release it in the second half of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 29.1

Quillstone Ledger 29.1 was released on 1 November 2024 for customers on the hosted plan.

What changed:

- Invoice matching: now keeps the cost centre for 180 days. After that the value is removed, and the entry shows when it was removed.
- Report generation: now keeps the source file name for 30 days. After that the value is removed, and the entry shows when it was removed.
- Exports: can now be turned off for each user. When it is turned off, the device type is not collected for that user from then on.
- The audit view: can now be turned off for each user. When it is turned off, the import job number is not collected for that user from then on.
- The audit view: now names the user who last changed the exchange rate used. It does not show why the change was made, because the product does not ask for a reason.

- Fixed: in attachment storage, a month with no records showed a blank page. It now shows an empty table.

Known issue: the audit view does not yet support files larger than 50 MB. Larger files are refused with a message that names the limit.

## Quillstone Ledger 29.2

Quillstone Ledger 29.2 was released on 8 November 2024 for customers on the hosted plan.

What changed:

- Scheduled reports: can now be turned off for each user. When it is turned off, the approving user's role is not collected for that user from then on.
- The tax summary: can now be turned off for each user. When it is turned off, the cost centre is not collected for that user from then on.
- Bank feeds: now keeps the batch number for 180 days. After that the value is removed, and the entry shows when it was removed.
- Attachment storage: now keeps the approving user's role for 180 days. After that the value is removed, and the entry shows when it was removed.
- The activity log: now shows the device type beside each entry. Entries made before 29.2 show a dash, because the value was not recorded.

- Fixed: in the audit view, a month with no records showed a blank page. It now shows an empty table.
- Fixed: user roles did not keep the sort order after a page reload. It now keeps it until the user changes it.

## Quillstone Ledger 29.3

Quillstone Ledger 29.3 was released on 15 November 2024 for customers on the hosted plan.

What changed:

- The mobile view: faster on accounts with many records. In our own test on 15 November 2024, a 96,000-row run took 55 seconds on 29.2 and 13 seconds on 29.3. Your results depend on your data.
- Bank feeds: now includes a column for the cost centre. The column is empty for records created before 29.3, because the cost centre was not stored then.
- Report generation: can now be turned off for each user. When it is turned off, the import job number is not collected for that user from then on.
- The tax summary: now includes a column for the exchange rate used. The column is empty for records created before 29.3, because the exchange rate used was not stored then.
- The activity log: now keeps the attachment count for 365 days. After that the value is removed, and the entry shows when it was removed.

- Fixed: the mobile view did not keep the sort order after a page reload. It now keeps it until the user changes it.

Planned, not yet available: the attachment count in bank feeds. We expect to release it in the first quarter of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 29.4

Quillstone Ledger 29.4 was released on 22 November 2024 for customers on the hosted plan.

What changed:

- The activity log: now shows the attachment count beside each entry. Entries made before 29.4 show a dash, because the value was not recorded.
- Scheduled reports: faster on accounts with many records. In our own test on 22 November 2024, a 130,000-row run took 77 seconds on 29.3 and 19 seconds on 29.4. Your results depend on your data.
- Report generation: now accepts the attachment count as a filter. The filter applies to records created on or after 29.4.
- The search index: now keeps the import job number for 90 days. After that the value is removed, and the entry shows when it was removed.

- Fixed: totals in the search index were rounded before they were added, which could change the last digit. They are now added first and rounded once.
- Fixed: a date typed as day/month could be read as month/day in currency conversion. The format now follows the account's settings.

Planned, not yet available: the reviewer's comment count in scheduled reports. We expect to release it in the second half of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 29.5

Quillstone Ledger 29.5 was released on 1 December 2024 for customers on the hosted plan.

What changed:

- The approval queue: now keeps the import job number for 180 days. After that the value is removed, and the entry shows when it was removed.
- User roles: now keeps the approving user's role for 90 days. After that the value is removed, and the entry shows when it was removed.
- The import screen: now names the user who last changed the reviewer's comment count. It does not show why the change was made, because the product does not ask for a reason.

- Fixed: totals in the audit view were rounded before they were added, which could change the last digit. They are now added first and rounded once.
- Fixed: in the tax summary, a month with no records showed a blank page. It now shows an empty table.

## Quillstone Ledger 29.6

Quillstone Ledger 29.6 was released on 8 December 2024 for customers on the hosted plan.

What changed:

- Account settings: now includes a column for the import job number. The column is empty for records created before 29.6, because the import job number was not stored then.
- Currency conversion: now keeps the device type for 180 days. After that the value is removed, and the entry shows when it was removed.
- Exports: can now be turned off for each user. When it is turned off, the cost centre is not collected for that user from then on.

- Fixed: totals in user roles were rounded before they were added, which could change the last digit. They are now added first and rounded once.
- Fixed: a date typed as day/month could be read as month/day in invoice matching. The format now follows the account's settings.

Known issue: on slow connections account settings can take up to 30 seconds to load. We are working on it and will say here when it is fixed.

Planned, not yet available: the attachment count in currency conversion. We expect to release it in the second quarter of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 29.7

Quillstone Ledger 29.7 was released on 15 December 2024 for customers on the hosted plan.

What changed:

- Currency conversion: now names the user who last changed the import job number. It does not show why the change was made, because the product does not ask for a reason.
- The mobile view: now shows the approving user's role beside each entry. Entries made before 29.7 show a dash, because the value was not recorded.
- Exports: now includes a column for the cost centre. The column is empty for records created before 29.7, because the cost centre was not stored then.

- Fixed: in the mobile view, a month with no records showed a blank page. It now shows an empty table.

## Quillstone Ledger 29.8

Quillstone Ledger 29.8 was released on 22 December 2024 for customers on the hosted plan.

What changed:

- The mobile view: can now be turned off for each user. When it is turned off, the approving user's role is not collected for that user from then on.
- User roles: now keeps the payment terms for 365 days. After that the value is removed, and the entry shows when it was removed.
- Invoice matching: now accepts the original currency as a filter. The filter applies to records created on or after 29.8.

- Fixed: a date typed as day/month could be read as month/day in bank feeds. The format now follows the account's settings.

Known issue: exports does not yet support files larger than 100 MB. Larger files are refused with a message that names the limit.

## Quillstone Ledger 29.9

Quillstone Ledger 29.9 was released on 1 January 2025 for customers on the hosted plan.

What changed:

- Scheduled reports: faster on accounts with many records. In our own test on 1 January 2025, a 137,000-row run took 49 seconds on 29.8 and 16 seconds on 29.9. Your results depend on your data.
- The mobile view: now shows the batch number beside each entry. Entries made before 29.9 show a dash, because the value was not recorded.
- User roles: now names the user who last changed the reviewer's comment count. It does not show why the change was made, because the product does not ask for a reason.
- Bank feeds: now keeps the exchange rate used for 365 days. After that the value is removed, and the entry shows when it was removed.
- Report generation: now includes a column for the exchange rate used. The column is empty for records created before 29.9, because the exchange rate used was not stored then.

- Fixed: a date typed as day/month could be read as month/day in user roles. The format now follows the account's settings.

## Quillstone Ledger 30.0

Quillstone Ledger 30.0 was released on 8 January 2025 for customers on the hosted plan.

What changed:

- The approval queue: now shows the source file name beside each entry. Entries made before 30.0 show a dash, because the value was not recorded.
- The mobile view: faster on accounts with many records. In our own test on 8 January 2025, a 98,000-row run took 81 seconds on 29.9 and 27 seconds on 30.0. Your results depend on your data.
- Report generation: now names the user who last changed the batch number. It does not show why the change was made, because the product does not ask for a reason.

- Fixed: account settings did not keep the sort order after a page reload. It now keeps it until the user changes it.
- Fixed: totals in report generation were rounded before they were added, which could change the last digit. They are now added first and rounded once.

Known issue: on slow connections invoice matching can take up to 30 seconds to load. We are working on it and will say here when it is fixed.

## Quillstone Ledger 30.1

Quillstone Ledger 30.1 was released on 15 January 2025 for customers on the hosted plan.

What changed:

- The search index: now includes a column for the device type. The column is empty for records created before 30.1, because the device type was not stored then.
- Exports: now keeps the import job number for 365 days. After that the value is removed, and the entry shows when it was removed.
- Exports: now includes a column for the reviewer's comment count. The column is empty for records created before 30.1, because the reviewer's comment count was not stored then.

- Fixed: report generation did not keep the sort order after a page reload. It now keeps it until the user changes it.

## Quillstone Ledger 30.2

Quillstone Ledger 30.2 was released on 22 January 2025 for customers on the hosted plan.

What changed:

- Invoice matching: now accepts the import job number as a filter. The filter applies to records created on or after 30.2.
- The import screen: faster on accounts with many records. In our own test on 22 January 2025, a 132,000-row run took 74 seconds on 30.1 and 18 seconds on 30.2. Your results depend on your data.
- The import screen: now names the user who last changed the payment terms. It does not show why the change was made, because the product does not ask for a reason.
- Account settings: now accepts the payment terms as a filter. The filter applies to records created on or after 30.2.
- The search index: can now be turned off for each user. When it is turned off, the original currency is not collected for that user from then on.

- Fixed: in invoice matching, a month with no records showed a blank page. It now shows an empty table.
- Fixed: in bank feeds, a month with no records showed a blank page. It now shows an empty table.

## Quillstone Ledger 30.3

Quillstone Ledger 30.3 was released on 1 February 2025 for customers on the hosted plan.

What changed:

- Scheduled reports: now shows the device type beside each entry. Entries made before 30.3 show a dash, because the value was not recorded.
- The mobile view: now includes a column for the reviewer's comment count. The column is empty for records created before 30.3, because the reviewer's comment count was not stored then.
- Invoice matching: now includes a column for the approving user's role. The column is empty for records created before 30.3, because the approving user's role was not stored then.
- The search index: faster on accounts with many records. In our own test on 1 February 2025, a 168,000-row run took 51 seconds on 30.2 and 12 seconds on 30.3. Your results depend on your data.
- Exports: now accepts the payment terms as a filter. The filter applies to records created on or after 30.3.

- Fixed: totals in the import screen were rounded before they were added, which could change the last digit. They are now added first and rounded once.

## Quillstone Ledger 30.4

Quillstone Ledger 30.4 was released on 8 February 2025 for customers on the hosted plan.

What changed:

- The tax summary: now includes a column for the time zone of the entry. The column is empty for records created before 30.4, because the time zone of the entry was not stored then.
- Report generation: now keeps the payment terms for 365 days. After that the value is removed, and the entry shows when it was removed.
- Scheduled reports: now includes a column for the exchange rate used. The column is empty for records created before 30.4, because the exchange rate used was not stored then.
- Invoice matching: now keeps the import job number for 365 days. After that the value is removed, and the entry shows when it was removed.
- The activity log: now keeps the exchange rate used for 30 days. After that the value is removed, and the entry shows when it was removed.

- Fixed: totals in scheduled reports were rounded before they were added, which could change the last digit. They are now added first and rounded once.

Planned, not yet available: the reviewer's comment count in user roles. We expect to release it in the first quarter of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 30.5

Quillstone Ledger 30.5 was released on 15 February 2025 for customers on the hosted plan.

What changed:

- Currency conversion: now names the user who last changed the exchange rate used. It does not show why the change was made, because the product does not ask for a reason.
- Exports: now names the user who last changed the approving user's role. It does not show why the change was made, because the product does not ask for a reason.
- The search index: now names the user who last changed the exchange rate used. It does not show why the change was made, because the product does not ask for a reason.

- Fixed: report generation did not keep the sort order after a page reload. It now keeps it until the user changes it.

## Quillstone Ledger 30.6

Quillstone Ledger 30.6 was released on 22 February 2025 for customers on the hosted plan.

What changed:

- The mobile view: faster on accounts with many records. In our own test on 22 February 2025, a 179,000-row run took 31 seconds on 30.5 and 7 seconds on 30.6. Your results depend on your data.
- The audit view: now shows the import job number beside each entry. Entries made before 30.6 show a dash, because the value was not recorded.
- Bank feeds: now accepts the exchange rate used as a filter. The filter applies to records created on or after 30.6.
- The approval queue: faster on accounts with many records. In our own test on 22 February 2025, a 79,000-row run took 47 seconds on 30.5 and 23 seconds on 30.6. Your results depend on your data.

- Fixed: a date typed as day/month could be read as month/day in the tax summary. The format now follows the account's settings.
- Fixed: a date typed as day/month could be read as month/day in the approval queue. The format now follows the account's settings.

## Quillstone Ledger 30.7

Quillstone Ledger 30.7 was released on 1 March 2025 for customers on the hosted plan.

What changed:

- The mobile view: now shows the source file name beside each entry. Entries made before 30.7 show a dash, because the value was not recorded.
- Invoice matching: now names the user who last changed the reviewer's comment count. It does not show why the change was made, because the product does not ask for a reason.
- The approval queue: faster on accounts with many records. In our own test on 1 March 2025, a 35,000-row run took 20 seconds on 30.6 and 10 seconds on 30.7. Your results depend on your data.
- Scheduled reports: now names the user who last changed the approving user's role. It does not show why the change was made, because the product does not ask for a reason.
- The audit view: now shows the exchange rate used beside each entry. Entries made before 30.7 show a dash, because the value was not recorded.

- Fixed: in account settings, a month with no records showed a blank page. It now shows an empty table.
- Fixed: a date typed as day/month could be read as month/day in scheduled reports. The format now follows the account's settings.

Planned, not yet available: the payment terms in currency conversion. We expect to release it in the first quarter of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 30.8

Quillstone Ledger 30.8 was released on 8 March 2025 for customers on the hosted plan.

What changed:

- Exports: now keeps the payment terms for 30 days. After that the value is removed, and the entry shows when it was removed.
- Exports: can now be turned off for each user. When it is turned off, the import job number is not collected for that user from then on.
- Report generation: can now be turned off for each user. When it is turned off, the payment terms is not collected for that user from then on.
- Scheduled reports: now keeps the device type for 180 days. After that the value is removed, and the entry shows when it was removed.
- Currency conversion: can now be turned off for each user. When it is turned off, the source file name is not collected for that user from then on.

- Fixed: in currency conversion, a month with no records showed a blank page. It now shows an empty table.

Planned, not yet available: the time zone of the entry in the audit view. We expect to release it in the second half of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 30.9

Quillstone Ledger 30.9 was released on 15 March 2025 for customers on the hosted plan.

What changed:

- Account settings: now accepts the device type as a filter. The filter applies to records created on or after 30.9.
- Account settings: now includes a column for the device type. The column is empty for records created before 30.9, because the device type was not stored then.
- The tax summary: now includes a column for the import job number. The column is empty for records created before 30.9, because the import job number was not stored then.
- The import screen: now keeps the reviewer's comment count for 365 days. After that the value is removed, and the entry shows when it was removed.
- Bank feeds: faster on accounts with many records. In our own test on 15 March 2025, a 67,000-row run took 51 seconds on 30.8 and 10 seconds on 30.9. Your results depend on your data.

- Fixed: a date typed as day/month could be read as month/day in invoice matching. The format now follows the account's settings.

## Quillstone Ledger 31.0

Quillstone Ledger 31.0 was released on 22 March 2025 for customers on the hosted plan.

What changed:

- The tax summary: now keeps the batch number for 365 days. After that the value is removed, and the entry shows when it was removed.
- Invoice matching: faster on accounts with many records. In our own test on 22 March 2025, a 165,000-row run took 70 seconds on 30.9 and 35 seconds on 31.0. Your results depend on your data.
- The activity log: can now be turned off for each user. When it is turned off, the cost centre is not collected for that user from then on.
- Exports: can now be turned off for each user. When it is turned off, the source file name is not collected for that user from then on.

- Fixed: a date typed as day/month could be read as month/day in scheduled reports. The format now follows the account's settings.
- Fixed: a date typed as day/month could be read as month/day in account settings. The format now follows the account's settings.

Planned, not yet available: the reviewer's comment count in account settings. We expect to release it in the second half of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 31.1

Quillstone Ledger 31.1 was released on 1 April 2025 for customers on the hosted plan.

What changed:

- Account settings: now shows the approving user's role beside each entry. Entries made before 31.1 show a dash, because the value was not recorded.
- The audit view: can now be turned off for each user. When it is turned off, the source file name is not collected for that user from then on.
- User roles: now names the user who last changed the attachment count. It does not show why the change was made, because the product does not ask for a reason.
- The search index: can now be turned off for each user. When it is turned off, the device type is not collected for that user from then on.

- Fixed: in the tax summary, a month with no records showed a blank page. It now shows an empty table.
- Fixed: in invoice matching, a month with no records showed a blank page. It now shows an empty table.

## Quillstone Ledger 31.2

Quillstone Ledger 31.2 was released on 8 April 2025 for customers on the hosted plan.

What changed:

- The audit view: now accepts the exchange rate used as a filter. The filter applies to records created on or after 31.2.
- The import screen: can now be turned off for each user. When it is turned off, the original currency is not collected for that user from then on.
- The audit view: can now be turned off for each user. When it is turned off, the device type is not collected for that user from then on.
- Bank feeds: now includes a column for the device type. The column is empty for records created before 31.2, because the device type was not stored then.

- Fixed: a date typed as day/month could be read as month/day in attachment storage. The format now follows the account's settings.
- Fixed: totals in invoice matching were rounded before they were added, which could change the last digit. They are now added first and rounded once.

Known issue: on slow connections user roles can take up to 10 seconds to load. We are working on it and will say here when it is fixed.

Planned, not yet available: the import job number in the import screen. We expect to release it in the second quarter of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 31.3

Quillstone Ledger 31.3 was released on 15 April 2025 for customers on the hosted plan.

What changed:

- Account settings: now includes a column for the exchange rate used. The column is empty for records created before 31.3, because the exchange rate used was not stored then.
- Invoice matching: now names the user who last changed the device type. It does not show why the change was made, because the product does not ask for a reason.
- Scheduled reports: can now be turned off for each user. When it is turned off, the batch number is not collected for that user from then on.
- The tax summary: faster on accounts with many records. In our own test on 15 April 2025, a 80,000-row run took 55 seconds on 31.2 and 27 seconds on 31.3. Your results depend on your data.
- The audit view: can now be turned off for each user. When it is turned off, the import job number is not collected for that user from then on.

- Fixed: totals in attachment storage were rounded before they were added, which could change the last digit. They are now added first and rounded once.

Known issue: the audit view does not yet support files larger than 50 MB. Larger files are refused with a message that names the limit.

## Quillstone Ledger 31.4

Quillstone Ledger 31.4 was released on 22 April 2025 for customers on the hosted plan.

What changed:

- The tax summary: now names the user who last changed the source file name. It does not show why the change was made, because the product does not ask for a reason.
- Invoice matching: now keeps the exchange rate used for 90 days. After that the value is removed, and the entry shows when it was removed.
- The audit view: now accepts the device type as a filter. The filter applies to records created on or after 31.4.

- Fixed: a date typed as day/month could be read as month/day in attachment storage. The format now follows the account's settings.
- Fixed: the approval queue did not keep the sort order after a page reload. It now keeps it until the user changes it.

Known issue: on slow connections the tax summary can take up to 20 seconds to load. We are working on it and will say here when it is fixed.

## Quillstone Ledger 31.5

Quillstone Ledger 31.5 was released on 1 May 2025 for customers on the hosted plan.

What changed:

- Exports: now accepts the import job number as a filter. The filter applies to records created on or after 31.5.
- The import screen: now names the user who last changed the reviewer's comment count. It does not show why the change was made, because the product does not ask for a reason.
- The search index: now shows the time zone of the entry beside each entry. Entries made before 31.5 show a dash, because the value was not recorded.
- Scheduled reports: now shows the exchange rate used beside each entry. Entries made before 31.5 show a dash, because the value was not recorded.

- Fixed: a date typed as day/month could be read as month/day in the activity log. The format now follows the account's settings.
- Fixed: report generation did not keep the sort order after a page reload. It now keeps it until the user changes it.

Planned, not yet available: the exchange rate used in the audit view. We expect to release it in the second half of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 31.6

Quillstone Ledger 31.6 was released on 8 May 2025 for customers on the hosted plan.

What changed:

- The import screen: now accepts the approving user's role as a filter. The filter applies to records created on or after 31.6.
- User roles: can now be turned off for each user. When it is turned off, the device type is not collected for that user from then on.
- User roles: now shows the cost centre beside each entry. Entries made before 31.6 show a dash, because the value was not recorded.
- Currency conversion: can now be turned off for each user. When it is turned off, the cost centre is not collected for that user from then on.

- Fixed: scheduled reports did not keep the sort order after a page reload. It now keeps it until the user changes it.
- Fixed: totals in exports were rounded before they were added, which could change the last digit. They are now added first and rounded once.

Known issue: currency conversion does not yet support files larger than 25 MB. Larger files are refused with a message that names the limit.

Planned, not yet available: the attachment count in account settings. We expect to release it in the second half of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 31.7

Quillstone Ledger 31.7 was released on 15 May 2025 for customers on the hosted plan.

What changed:

- Account settings: now names the user who last changed the attachment count. It does not show why the change was made, because the product does not ask for a reason.
- The approval queue: now accepts the import job number as a filter. The filter applies to records created on or after 31.7.
- The approval queue: now shows the import job number beside each entry. Entries made before 31.7 show a dash, because the value was not recorded.
- The mobile view: now names the user who last changed the reviewer's comment count. It does not show why the change was made, because the product does not ask for a reason.

- Fixed: the audit view did not keep the sort order after a page reload. It now keeps it until the user changes it.

Known issue: the search index does not yet support files larger than 25 MB. Larger files are refused with a message that names the limit.

Planned, not yet available: the source file name in scheduled reports. We expect to release it in the second quarter of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 31.8

Quillstone Ledger 31.8 was released on 22 May 2025 for customers on the hosted plan.

What changed:

- Bank feeds: now accepts the time zone of the entry as a filter. The filter applies to records created on or after 31.8.
- Scheduled reports: now names the user who last changed the import job number. It does not show why the change was made, because the product does not ask for a reason.
- Invoice matching: can now be turned off for each user. When it is turned off, the device type is not collected for that user from then on.
- Exports: faster on accounts with many records. In our own test on 22 May 2025, a 24,000-row run took 77 seconds on 31.7 and 25 seconds on 31.8. Your results depend on your data.

- Fixed: in the activity log, a month with no records showed a blank page. It now shows an empty table.

Known issue: the import screen does not yet support files larger than 25 MB. Larger files are refused with a message that names the limit.

## Quillstone Ledger 31.9

Quillstone Ledger 31.9 was released on 1 June 2025 for customers on the hosted plan.

What changed:

- The mobile view: faster on accounts with many records. In our own test on 1 June 2025, a 36,000-row run took 74 seconds on 31.8 and 24 seconds on 31.9. Your results depend on your data.
- Report generation: faster on accounts with many records. In our own test on 1 June 2025, a 51,000-row run took 46 seconds on 31.8 and 11 seconds on 31.9. Your results depend on your data.
- Attachment storage: now keeps the time zone of the entry for 30 days. After that the value is removed, and the entry shows when it was removed.
- Bank feeds: now shows the import job number beside each entry. Entries made before 31.9 show a dash, because the value was not recorded.
- Scheduled reports: now names the user who last changed the attachment count. It does not show why the change was made, because the product does not ask for a reason.

- Fixed: totals in account settings were rounded before they were added, which could change the last digit. They are now added first and rounded once.

Known issue: on slow connections the mobile view can take up to 20 seconds to load. We are working on it and will say here when it is fixed.

Planned, not yet available: the source file name in invoice matching. We expect to release it in the second half of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 32.0

Quillstone Ledger 32.0 was released on 8 June 2025 for customers on the hosted plan.

What changed:

- The tax summary: can now be turned off for each user. When it is turned off, the device type is not collected for that user from then on.
- Exports: can now be turned off for each user. When it is turned off, the batch number is not collected for that user from then on.
- The import screen: now names the user who last changed the import job number. It does not show why the change was made, because the product does not ask for a reason.
- Scheduled reports: now names the user who last changed the payment terms. It does not show why the change was made, because the product does not ask for a reason.
- Bank feeds: now shows the time zone of the entry beside each entry. Entries made before 32.0 show a dash, because the value was not recorded.

- Fixed: totals in the search index were rounded before they were added, which could change the last digit. They are now added first and rounded once.

Known issue: the approval queue does not yet support files larger than 100 MB. Larger files are refused with a message that names the limit.

## Quillstone Ledger 32.1

Quillstone Ledger 32.1 was released on 15 June 2025 for customers on the hosted plan.

What changed:

- Attachment storage: faster on accounts with many records. In our own test on 15 June 2025, a 40,000-row run took 20 seconds on 32.0 and 5 seconds on 32.1. Your results depend on your data.
- The import screen: now accepts the cost centre as a filter. The filter applies to records created on or after 32.1.
- Currency conversion: faster on accounts with many records. In our own test on 15 June 2025, a 183,000-row run took 78 seconds on 32.0 and 15 seconds on 32.1. Your results depend on your data.

- Fixed: in the activity log, a month with no records showed a blank page. It now shows an empty table.
- Fixed: a date typed as day/month could be read as month/day in scheduled reports. The format now follows the account's settings.

Known issue: exports does not yet support files larger than 100 MB. Larger files are refused with a message that names the limit.

## Quillstone Ledger 32.2

Quillstone Ledger 32.2 was released on 22 June 2025 for customers on the hosted plan.

What changed:

- Currency conversion: now accepts the source file name as a filter. The filter applies to records created on or after 32.2.
- The tax summary: now keeps the cost centre for 90 days. After that the value is removed, and the entry shows when it was removed.
- The mobile view: now includes a column for the exchange rate used. The column is empty for records created before 32.2, because the exchange rate used was not stored then.
- Attachment storage: now accepts the exchange rate used as a filter. The filter applies to records created on or after 32.2.
- Account settings: now names the user who last changed the time zone of the entry. It does not show why the change was made, because the product does not ask for a reason.

- Fixed: in the search index, a month with no records showed a blank page. It now shows an empty table.
- Fixed: the activity log did not keep the sort order after a page reload. It now keeps it until the user changes it.

Known issue: on slow connections scheduled reports can take up to 10 seconds to load. We are working on it and will say here when it is fixed.

Planned, not yet available: the attachment count in bank feeds. We expect to release it in the first quarter of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 32.3

Quillstone Ledger 32.3 was released on 1 July 2025 for customers on the hosted plan.

What changed:

- The search index: can now be turned off for each user. When it is turned off, the import job number is not collected for that user from then on.
- The approval queue: faster on accounts with many records. In our own test on 1 July 2025, a 163,000-row run took 66 seconds on 32.2 and 33 seconds on 32.3. Your results depend on your data.
- The mobile view: faster on accounts with many records. In our own test on 1 July 2025, a 121,000-row run took 61 seconds on 32.2 and 20 seconds on 32.3. Your results depend on your data.
- The search index: now accepts the attachment count as a filter. The filter applies to records created on or after 32.3.
- Scheduled reports: now names the user who last changed the payment terms. It does not show why the change was made, because the product does not ask for a reason.

- Fixed: in the activity log, a month with no records showed a blank page. It now shows an empty table.

## Quillstone Ledger 32.4

Quillstone Ledger 32.4 was released on 8 July 2025 for customers on the hosted plan.

What changed:

- The audit view: now shows the time zone of the entry beside each entry. Entries made before 32.4 show a dash, because the value was not recorded.
- Invoice matching: now keeps the approving user's role for 30 days. After that the value is removed, and the entry shows when it was removed.
- Scheduled reports: faster on accounts with many records. In our own test on 8 July 2025, a 34,000-row run took 36 seconds on 32.3 and 12 seconds on 32.4. Your results depend on your data.
- Account settings: now keeps the source file name for 30 days. After that the value is removed, and the entry shows when it was removed.

- Fixed: totals in the search index were rounded before they were added, which could change the last digit. They are now added first and rounded once.

Known issue: on slow connections the search index can take up to 10 seconds to load. We are working on it and will say here when it is fixed.

## Quillstone Ledger 32.5

Quillstone Ledger 32.5 was released on 15 July 2025 for customers on the hosted plan.

What changed:

- The mobile view: now keeps the cost centre for 30 days. After that the value is removed, and the entry shows when it was removed.
- The activity log: now names the user who last changed the import job number. It does not show why the change was made, because the product does not ask for a reason.
- Currency conversion: faster on accounts with many records. In our own test on 15 July 2025, a 95,000-row run took 58 seconds on 32.4 and 14 seconds on 32.5. Your results depend on your data.

- Fixed: in scheduled reports, a month with no records showed a blank page. It now shows an empty table.

Known issue: currency conversion does not yet support files larger than 50 MB. Larger files are refused with a message that names the limit.

## Quillstone Ledger 32.6

Quillstone Ledger 32.6 was released on 22 July 2025 for customers on the hosted plan.

What changed:

- Invoice matching: now keeps the exchange rate used for 180 days. After that the value is removed, and the entry shows when it was removed.
- Attachment storage: now keeps the attachment count for 365 days. After that the value is removed, and the entry shows when it was removed.
- Attachment storage: can now be turned off for each user. When it is turned off, the original currency is not collected for that user from then on.
- Report generation: now includes a column for the source file name. The column is empty for records created before 32.6, because the source file name was not stored then.
- The import screen: now accepts the device type as a filter. The filter applies to records created on or after 32.6.

- Fixed: a date typed as day/month could be read as month/day in the import screen. The format now follows the account's settings.
- Fixed: totals in exports were rounded before they were added, which could change the last digit. They are now added first and rounded once.

## Quillstone Ledger 32.7

Quillstone Ledger 32.7 was released on 1 August 2025 for customers on the hosted plan.

What changed:

- Currency conversion: now names the user who last changed the source file name. It does not show why the change was made, because the product does not ask for a reason.
- Account settings: now names the user who last changed the import job number. It does not show why the change was made, because the product does not ask for a reason.
- The activity log: now keeps the batch number for 90 days. After that the value is removed, and the entry shows when it was removed.

- Fixed: totals in the mobile view were rounded before they were added, which could change the last digit. They are now added first and rounded once.
- Fixed: totals in bank feeds were rounded before they were added, which could change the last digit. They are now added first and rounded once.

## Quillstone Ledger 32.8

Quillstone Ledger 32.8 was released on 8 August 2025 for customers on the hosted plan.

What changed:

- Scheduled reports: now shows the time zone of the entry beside each entry. Entries made before 32.8 show a dash, because the value was not recorded.
- The approval queue: now names the user who last changed the payment terms. It does not show why the change was made, because the product does not ask for a reason.
- Scheduled reports: now names the user who last changed the device type. It does not show why the change was made, because the product does not ask for a reason.

- Fixed: a date typed as day/month could be read as month/day in report generation. The format now follows the account's settings.

## Quillstone Ledger 32.9

Quillstone Ledger 32.9 was released on 15 August 2025 for customers on the hosted plan.

What changed:

- The activity log: now names the user who last changed the attachment count. It does not show why the change was made, because the product does not ask for a reason.
- The audit view: now includes a column for the attachment count. The column is empty for records created before 32.9, because the attachment count was not stored then.
- Scheduled reports: now names the user who last changed the source file name. It does not show why the change was made, because the product does not ask for a reason.

- Fixed: totals in the audit view were rounded before they were added, which could change the last digit. They are now added first and rounded once.
- Fixed: totals in currency conversion were rounded before they were added, which could change the last digit. They are now added first and rounded once.

Known issue: user roles does not yet support files larger than 50 MB. Larger files are refused with a message that names the limit.

## Quillstone Ledger 33.0

Quillstone Ledger 33.0 was released on 22 August 2025 for customers on the hosted plan.

What changed:

- Report generation: faster on accounts with many records. In our own test on 22 August 2025, a 34,000-row run took 84 seconds on 32.9 and 16 seconds on 33.0. Your results depend on your data.
- The tax summary: now shows the exchange rate used beside each entry. Entries made before 33.0 show a dash, because the value was not recorded.
- Report generation: can now be turned off for each user. When it is turned off, the approving user's role is not collected for that user from then on.
- The tax summary: now names the user who last changed the original currency. It does not show why the change was made, because the product does not ask for a reason.

- Fixed: in the mobile view, a month with no records showed a blank page. It now shows an empty table.
- Fixed: totals in the tax summary were rounded before they were added, which could change the last digit. They are now added first and rounded once.

Known issue: the mobile view does not yet support files larger than 25 MB. Larger files are refused with a message that names the limit.

## Quillstone Ledger 33.1

Quillstone Ledger 33.1 was released on 1 September 2025 for customers on the hosted plan.

What changed:

- Invoice matching: faster on accounts with many records. In our own test on 1 September 2025, a 161,000-row run took 47 seconds on 33.0 and 9 seconds on 33.1. Your results depend on your data.
- Currency conversion: faster on accounts with many records. In our own test on 1 September 2025, a 184,000-row run took 88 seconds on 33.0 and 22 seconds on 33.1. Your results depend on your data.
- The import screen: faster on accounts with many records. In our own test on 1 September 2025, a 91,000-row run took 46 seconds on 33.0 and 9 seconds on 33.1. Your results depend on your data.
- Attachment storage: now names the user who last changed the time zone of the entry. It does not show why the change was made, because the product does not ask for a reason.

- Fixed: in the mobile view, a month with no records showed a blank page. It now shows an empty table.
- Fixed: totals in the mobile view were rounded before they were added, which could change the last digit. They are now added first and rounded once.

## Quillstone Ledger 33.2

Quillstone Ledger 33.2 was released on 8 September 2025 for customers on the hosted plan.

What changed:

- Bank feeds: can now be turned off for each user. When it is turned off, the import job number is not collected for that user from then on.
- The approval queue: can now be turned off for each user. When it is turned off, the exchange rate used is not collected for that user from then on.
- Currency conversion: now names the user who last changed the approving user's role. It does not show why the change was made, because the product does not ask for a reason.
- The activity log: now names the user who last changed the payment terms. It does not show why the change was made, because the product does not ask for a reason.
- Currency conversion: now shows the source file name beside each entry. Entries made before 33.2 show a dash, because the value was not recorded.

- Fixed: in the search index, a month with no records showed a blank page. It now shows an empty table.
- Fixed: report generation did not keep the sort order after a page reload. It now keeps it until the user changes it.

## Quillstone Ledger 33.3

Quillstone Ledger 33.3 was released on 15 September 2025 for customers on the hosted plan.

What changed:

- The import screen: faster on accounts with many records. In our own test on 15 September 2025, a 136,000-row run took 43 seconds on 33.2 and 21 seconds on 33.3. Your results depend on your data.
- Invoice matching: now names the user who last changed the approving user's role. It does not show why the change was made, because the product does not ask for a reason.
- Bank feeds: now accepts the reviewer's comment count as a filter. The filter applies to records created on or after 33.3.
- Exports: now names the user who last changed the exchange rate used. It does not show why the change was made, because the product does not ask for a reason.

- Fixed: the import screen did not keep the sort order after a page reload. It now keeps it until the user changes it.

Planned, not yet available: the cost centre in the search index. We expect to release it in the second half of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 33.4

Quillstone Ledger 33.4 was released on 22 September 2025 for customers on the hosted plan.

What changed:

- Invoice matching: faster on accounts with many records. In our own test on 22 September 2025, a 170,000-row run took 64 seconds on 33.3 and 21 seconds on 33.4. Your results depend on your data.
- The approval queue: now names the user who last changed the original currency. It does not show why the change was made, because the product does not ask for a reason.
- Currency conversion: now shows the approving user's role beside each entry. Entries made before 33.4 show a dash, because the value was not recorded.
- Currency conversion: faster on accounts with many records. In our own test on 22 September 2025, a 121,000-row run took 63 seconds on 33.3 and 21 seconds on 33.4. Your results depend on your data.

- Fixed: in exports, a month with no records showed a blank page. It now shows an empty table.
- Fixed: in account settings, a month with no records showed a blank page. It now shows an empty table.

## Quillstone Ledger 33.5

Quillstone Ledger 33.5 was released on 1 October 2025 for customers on the hosted plan.

What changed:

- The mobile view: now keeps the import job number for 365 days. After that the value is removed, and the entry shows when it was removed.
- The mobile view: can now be turned off for each user. When it is turned off, the device type is not collected for that user from then on.
- Invoice matching: now shows the reviewer's comment count beside each entry. Entries made before 33.5 show a dash, because the value was not recorded.
- Report generation: now accepts the device type as a filter. The filter applies to records created on or after 33.5.
- The audit view: now shows the device type beside each entry. Entries made before 33.5 show a dash, because the value was not recorded.

- Fixed: the activity log did not keep the sort order after a page reload. It now keeps it until the user changes it.

## Quillstone Ledger 33.6

Quillstone Ledger 33.6 was released on 8 October 2025 for customers on the hosted plan.

What changed:

- The activity log: now includes a column for the exchange rate used. The column is empty for records created before 33.6, because the exchange rate used was not stored then.
- The search index: faster on accounts with many records. In our own test on 8 October 2025, a 126,000-row run took 44 seconds on 33.5 and 22 seconds on 33.6. Your results depend on your data.
- Scheduled reports: now names the user who last changed the device type. It does not show why the change was made, because the product does not ask for a reason.
- The activity log: can now be turned off for each user. When it is turned off, the reviewer's comment count is not collected for that user from then on.
- Attachment storage: now keeps the source file name for 90 days. After that the value is removed, and the entry shows when it was removed.

- Fixed: the import screen did not keep the sort order after a page reload. It now keeps it until the user changes it.

## Quillstone Ledger 33.7

Quillstone Ledger 33.7 was released on 15 October 2025 for customers on the hosted plan.

What changed:

- The search index: now names the user who last changed the cost centre. It does not show why the change was made, because the product does not ask for a reason.
- Attachment storage: can now be turned off for each user. When it is turned off, the reviewer's comment count is not collected for that user from then on.
- The mobile view: can now be turned off for each user. When it is turned off, the attachment count is not collected for that user from then on.
- The import screen: now accepts the source file name as a filter. The filter applies to records created on or after 33.7.

- Fixed: bank feeds did not keep the sort order after a page reload. It now keeps it until the user changes it.
- Fixed: in currency conversion, a month with no records showed a blank page. It now shows an empty table.

## Quillstone Ledger 33.8

Quillstone Ledger 33.8 was released on 22 October 2025 for customers on the hosted plan.

What changed:

- Invoice matching: now keeps the reviewer's comment count for 180 days. After that the value is removed, and the entry shows when it was removed.
- Scheduled reports: faster on accounts with many records. In our own test on 22 October 2025, a 122,000-row run took 84 seconds on 33.7 and 28 seconds on 33.8. Your results depend on your data.
- The mobile view: can now be turned off for each user. When it is turned off, the original currency is not collected for that user from then on.
- Exports: faster on accounts with many records. In our own test on 22 October 2025, a 133,000-row run took 72 seconds on 33.7 and 36 seconds on 33.8. Your results depend on your data.
- Account settings: now names the user who last changed the exchange rate used. It does not show why the change was made, because the product does not ask for a reason.

- Fixed: in the search index, a month with no records showed a blank page. It now shows an empty table.
- Fixed: a date typed as day/month could be read as month/day in the import screen. The format now follows the account's settings.

## Quillstone Ledger 33.9

Quillstone Ledger 33.9 was released on 1 November 2025 for customers on the hosted plan.

What changed:

- The search index: now accepts the device type as a filter. The filter applies to records created on or after 33.9.
- Attachment storage: now shows the time zone of the entry beside each entry. Entries made before 33.9 show a dash, because the value was not recorded.
- The mobile view: now shows the approving user's role beside each entry. Entries made before 33.9 show a dash, because the value was not recorded.
- The tax summary: now shows the source file name beside each entry. Entries made before 33.9 show a dash, because the value was not recorded.

- Fixed: a date typed as day/month could be read as month/day in account settings. The format now follows the account's settings.
- Fixed: the audit view did not keep the sort order after a page reload. It now keeps it until the user changes it.

Known issue: on slow connections currency conversion can take up to 30 seconds to load. We are working on it and will say here when it is fixed.

## Quillstone Ledger 34.0

Quillstone Ledger 34.0 was released on 8 November 2025 for customers on the hosted plan.

What changed:

- The approval queue: now keeps the device type for 365 days. After that the value is removed, and the entry shows when it was removed.
- The import screen: now accepts the device type as a filter. The filter applies to records created on or after 34.0.
- Invoice matching: now shows the import job number beside each entry. Entries made before 34.0 show a dash, because the value was not recorded.

- Fixed: a date typed as day/month could be read as month/day in the approval queue. The format now follows the account's settings.

Known issue: on slow connections the audit view can take up to 10 seconds to load. We are working on it and will say here when it is fixed.

## Quillstone Ledger 34.1

Quillstone Ledger 34.1 was released on 15 November 2025 for customers on the hosted plan.

What changed:

- Invoice matching: now accepts the cost centre as a filter. The filter applies to records created on or after 34.1.
- The approval queue: now shows the time zone of the entry beside each entry. Entries made before 34.1 show a dash, because the value was not recorded.
- The search index: now accepts the batch number as a filter. The filter applies to records created on or after 34.1.

- Fixed: a date typed as day/month could be read as month/day in the mobile view. The format now follows the account's settings.
- Fixed: the activity log did not keep the sort order after a page reload. It now keeps it until the user changes it.

Known issue: the mobile view does not yet support files larger than 100 MB. Larger files are refused with a message that names the limit.

## Quillstone Ledger 34.2

Quillstone Ledger 34.2 was released on 22 November 2025 for customers on the hosted plan.

What changed:

- The mobile view: now includes a column for the time zone of the entry. The column is empty for records created before 34.2, because the time zone of the entry was not stored then.
- Report generation: now includes a column for the original currency. The column is empty for records created before 34.2, because the original currency was not stored then.
- Exports: now keeps the original currency for 90 days. After that the value is removed, and the entry shows when it was removed.
- Report generation: now names the user who last changed the device type. It does not show why the change was made, because the product does not ask for a reason.
- Bank feeds: now names the user who last changed the source file name. It does not show why the change was made, because the product does not ask for a reason.

- Fixed: in currency conversion, a month with no records showed a blank page. It now shows an empty table.
- Fixed: in bank feeds, a month with no records showed a blank page. It now shows an empty table.

Known issue: on slow connections exports can take up to 30 seconds to load. We are working on it and will say here when it is fixed.

Planned, not yet available: the device type in the import screen. We expect to release it in the second quarter of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 34.3

Quillstone Ledger 34.3 was released on 1 December 2025 for customers on the hosted plan.

What changed:

- The import screen: now accepts the source file name as a filter. The filter applies to records created on or after 34.3.
- Invoice matching: now shows the device type beside each entry. Entries made before 34.3 show a dash, because the value was not recorded.
- The import screen: now keeps the exchange rate used for 30 days. After that the value is removed, and the entry shows when it was removed.
- The import screen: now shows the original currency beside each entry. Entries made before 34.3 show a dash, because the value was not recorded.

- Fixed: a date typed as day/month could be read as month/day in currency conversion. The format now follows the account's settings.

## Quillstone Ledger 34.4

Quillstone Ledger 34.4 was released on 8 December 2025 for customers on the hosted plan.

What changed:

- The audit view: now keeps the approving user's role for 30 days. After that the value is removed, and the entry shows when it was removed.
- Exports: now keeps the device type for 90 days. After that the value is removed, and the entry shows when it was removed.
- Report generation: now names the user who last changed the cost centre. It does not show why the change was made, because the product does not ask for a reason.
- Bank feeds: now shows the payment terms beside each entry. Entries made before 34.4 show a dash, because the value was not recorded.

- Fixed: in the activity log, a month with no records showed a blank page. It now shows an empty table.

## Quillstone Ledger 34.5

Quillstone Ledger 34.5 was released on 15 December 2025 for customers on the hosted plan.

What changed:

- User roles: now keeps the payment terms for 30 days. After that the value is removed, and the entry shows when it was removed.
- The approval queue: now keeps the time zone of the entry for 30 days. After that the value is removed, and the entry shows when it was removed.
- The activity log: now keeps the payment terms for 365 days. After that the value is removed, and the entry shows when it was removed.
- The audit view: can now be turned off for each user. When it is turned off, the time zone of the entry is not collected for that user from then on.

- Fixed: a date typed as day/month could be read as month/day in invoice matching. The format now follows the account's settings.

Planned, not yet available: the device type in report generation. We expect to release it in the first quarter of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 34.6

Quillstone Ledger 34.6 was released on 22 December 2025 for customers on the hosted plan.

What changed:

- Report generation: faster on accounts with many records. In our own test on 22 December 2025, a 182,000-row run took 60 seconds on 34.5 and 15 seconds on 34.6. Your results depend on your data.
- Currency conversion: faster on accounts with many records. In our own test on 22 December 2025, a 135,000-row run took 31 seconds on 34.5 and 10 seconds on 34.6. Your results depend on your data.
- The mobile view: now keeps the device type for 30 days. After that the value is removed, and the entry shows when it was removed.
- Report generation: now includes a column for the source file name. The column is empty for records created before 34.6, because the source file name was not stored then.

- Fixed: in the activity log, a month with no records showed a blank page. It now shows an empty table.

Known issue: account settings does not yet support files larger than 100 MB. Larger files are refused with a message that names the limit.

## Quillstone Ledger 34.7

Quillstone Ledger 34.7 was released on 1 January 2026 for customers on the hosted plan.

What changed:

- Attachment storage: now accepts the exchange rate used as a filter. The filter applies to records created on or after 34.7.
- The mobile view: now accepts the exchange rate used as a filter. The filter applies to records created on or after 34.7.
- The approval queue: now shows the import job number beside each entry. Entries made before 34.7 show a dash, because the value was not recorded.

- Fixed: exports did not keep the sort order after a page reload. It now keeps it until the user changes it.
- Fixed: the audit view did not keep the sort order after a page reload. It now keeps it until the user changes it.

Planned, not yet available: the attachment count in invoice matching. We expect to release it in the second quarter of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 34.8

Quillstone Ledger 34.8 was released on 8 January 2026 for customers on the hosted plan.

What changed:

- The mobile view: faster on accounts with many records. In our own test on 8 January 2026, a 101,000-row run took 20 seconds on 34.7 and 10 seconds on 34.8. Your results depend on your data.
- The approval queue: now accepts the exchange rate used as a filter. The filter applies to records created on or after 34.8.
- Currency conversion: now names the user who last changed the source file name. It does not show why the change was made, because the product does not ask for a reason.

- Fixed: in the approval queue, a month with no records showed a blank page. It now shows an empty table.

Known issue: exports does not yet support files larger than 50 MB. Larger files are refused with a message that names the limit.

## Quillstone Ledger 34.9

Quillstone Ledger 34.9 was released on 15 January 2026 for customers on the hosted plan.

What changed:

- Scheduled reports: now accepts the cost centre as a filter. The filter applies to records created on or after 34.9.
- Currency conversion: now accepts the cost centre as a filter. The filter applies to records created on or after 34.9.
- Attachment storage: now accepts the time zone of the entry as a filter. The filter applies to records created on or after 34.9.
- The search index: faster on accounts with many records. In our own test on 15 January 2026, a 138,000-row run took 34 seconds on 34.8 and 6 seconds on 34.9. Your results depend on your data.

- Fixed: user roles did not keep the sort order after a page reload. It now keeps it until the user changes it.

Planned, not yet available: the cost centre in exports. We expect to release it in the second quarter of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 35.0

Quillstone Ledger 35.0 was released on 22 January 2026 for customers on the hosted plan.

What changed:

- The import screen: now shows the import job number beside each entry. Entries made before 35.0 show a dash, because the value was not recorded.
- The mobile view: now keeps the device type for 180 days. After that the value is removed, and the entry shows when it was removed.
- Exports: can now be turned off for each user. When it is turned off, the original currency is not collected for that user from then on.
- User roles: now accepts the batch number as a filter. The filter applies to records created on or after 35.0.

- Fixed: bank feeds did not keep the sort order after a page reload. It now keeps it until the user changes it.

## Quillstone Ledger 35.1

Quillstone Ledger 35.1 was released on 1 February 2026 for customers on the hosted plan.

What changed:

- User roles: now names the user who last changed the device type. It does not show why the change was made, because the product does not ask for a reason.
- The mobile view: now names the user who last changed the payment terms. It does not show why the change was made, because the product does not ask for a reason.
- Exports: faster on accounts with many records. In our own test on 1 February 2026, a 189,000-row run took 40 seconds on 35.0 and 8 seconds on 35.1. Your results depend on your data.

- Fixed: totals in attachment storage were rounded before they were added, which could change the last digit. They are now added first and rounded once.

Known issue: the import screen does not yet support files larger than 100 MB. Larger files are refused with a message that names the limit.

## Quillstone Ledger 35.2

Quillstone Ledger 35.2 was released on 8 February 2026 for customers on the hosted plan.

What changed:

- Bank feeds: now names the user who last changed the batch number. It does not show why the change was made, because the product does not ask for a reason.
- The audit view: faster on accounts with many records. In our own test on 8 February 2026, a 86,000-row run took 31 seconds on 35.1 and 15 seconds on 35.2. Your results depend on your data.
- Report generation: now accepts the batch number as a filter. The filter applies to records created on or after 35.2.

- Fixed: a date typed as day/month could be read as month/day in invoice matching. The format now follows the account's settings.

Known issue: on slow connections user roles can take up to 30 seconds to load. We are working on it and will say here when it is fixed.

## Quillstone Ledger 35.3

Quillstone Ledger 35.3 was released on 15 February 2026 for customers on the hosted plan.

What changed:

- Attachment storage: faster on accounts with many records. In our own test on 15 February 2026, a 99,000-row run took 48 seconds on 35.2 and 24 seconds on 35.3. Your results depend on your data.
- Bank feeds: can now be turned off for each user. When it is turned off, the device type is not collected for that user from then on.
- The mobile view: now keeps the cost centre for 365 days. After that the value is removed, and the entry shows when it was removed.
- User roles: can now be turned off for each user. When it is turned off, the payment terms is not collected for that user from then on.

- Fixed: a date typed as day/month could be read as month/day in invoice matching. The format now follows the account's settings.

Known issue: the tax summary does not yet support files larger than 50 MB. Larger files are refused with a message that names the limit.

## Quillstone Ledger 35.4

Quillstone Ledger 35.4 was released on 22 February 2026 for customers on the hosted plan.

What changed:

- The search index: now keeps the cost centre for 90 days. After that the value is removed, and the entry shows when it was removed.
- Account settings: now accepts the time zone of the entry as a filter. The filter applies to records created on or after 35.4.
- The search index: can now be turned off for each user. When it is turned off, the import job number is not collected for that user from then on.
- Exports: now includes a column for the source file name. The column is empty for records created before 35.4, because the source file name was not stored then.

- Fixed: in account settings, a month with no records showed a blank page. It now shows an empty table.

## Quillstone Ledger 35.5

Quillstone Ledger 35.5 was released on 1 March 2026 for customers on the hosted plan.

What changed:

- The approval queue: can now be turned off for each user. When it is turned off, the import job number is not collected for that user from then on.
- User roles: now shows the cost centre beside each entry. Entries made before 35.5 show a dash, because the value was not recorded.
- The search index: faster on accounts with many records. In our own test on 1 March 2026, a 123,000-row run took 69 seconds on 35.4 and 13 seconds on 35.5. Your results depend on your data.

- Fixed: totals in exports were rounded before they were added, which could change the last digit. They are now added first and rounded once.
- Fixed: a date typed as day/month could be read as month/day in the search index. The format now follows the account's settings.

Known issue: on slow connections currency conversion can take up to 10 seconds to load. We are working on it and will say here when it is fixed.

## Quillstone Ledger 35.6

Quillstone Ledger 35.6 was released on 8 March 2026 for customers on the hosted plan.

What changed:

- The tax summary: now keeps the cost centre for 180 days. After that the value is removed, and the entry shows when it was removed.
- The audit view: can now be turned off for each user. When it is turned off, the reviewer's comment count is not collected for that user from then on.
- Currency conversion: can now be turned off for each user. When it is turned off, the attachment count is not collected for that user from then on.
- The tax summary: now includes a column for the source file name. The column is empty for records created before 35.6, because the source file name was not stored then.

- Fixed: a date typed as day/month could be read as month/day in report generation. The format now follows the account's settings.
- Fixed: a date typed as day/month could be read as month/day in the activity log. The format now follows the account's settings.

Known issue: the audit view does not yet support files larger than 50 MB. Larger files are refused with a message that names the limit.

## Quillstone Ledger 35.7

Quillstone Ledger 35.7 was released on 15 March 2026 for customers on the hosted plan.

What changed:

- Attachment storage: now shows the import job number beside each entry. Entries made before 35.7 show a dash, because the value was not recorded.
- The import screen: faster on accounts with many records. In our own test on 15 March 2026, a 184,000-row run took 71 seconds on 35.6 and 14 seconds on 35.7. Your results depend on your data.
- Attachment storage: now names the user who last changed the cost centre. It does not show why the change was made, because the product does not ask for a reason.
- The audit view: now keeps the import job number for 90 days. After that the value is removed, and the entry shows when it was removed.
- The audit view: now shows the original currency beside each entry. Entries made before 35.7 show a dash, because the value was not recorded.

- Fixed: in bank feeds, a month with no records showed a blank page. It now shows an empty table.

Planned, not yet available: the reviewer's comment count in account settings. We expect to release it in the first quarter of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 35.8

Quillstone Ledger 35.8 was released on 22 March 2026 for customers on the hosted plan.

What changed:

- Account settings: now accepts the exchange rate used as a filter. The filter applies to records created on or after 35.8.
- The audit view: now keeps the payment terms for 180 days. After that the value is removed, and the entry shows when it was removed.
- Currency conversion: can now be turned off for each user. When it is turned off, the cost centre is not collected for that user from then on.

- Fixed: in scheduled reports, a month with no records showed a blank page. It now shows an empty table.

Known issue: on slow connections attachment storage can take up to 10 seconds to load. We are working on it and will say here when it is fixed.

## Quillstone Ledger 35.9

Quillstone Ledger 35.9 was released on 1 April 2026 for customers on the hosted plan.

What changed:

- The tax summary: now shows the import job number beside each entry. Entries made before 35.9 show a dash, because the value was not recorded.
- Scheduled reports: now accepts the time zone of the entry as a filter. The filter applies to records created on or after 35.9.
- The activity log: faster on accounts with many records. In our own test on 1 April 2026, a 84,000-row run took 78 seconds on 35.8 and 19 seconds on 35.9. Your results depend on your data.
- Exports: now keeps the attachment count for 30 days. After that the value is removed, and the entry shows when it was removed.

- Fixed: a date typed as day/month could be read as month/day in user roles. The format now follows the account's settings.
- Fixed: in user roles, a month with no records showed a blank page. It now shows an empty table.

Known issue: on slow connections user roles can take up to 20 seconds to load. We are working on it and will say here when it is fixed.

Planned, not yet available: the exchange rate used in the import screen. We expect to release it in the first quarter of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 36.0

Quillstone Ledger 36.0 was released on 8 April 2026 for customers on the hosted plan.

What changed:

- Currency conversion: now shows the exchange rate used beside each entry. Entries made before 36.0 show a dash, because the value was not recorded.
- The tax summary: now accepts the payment terms as a filter. The filter applies to records created on or after 36.0.
- Invoice matching: now names the user who last changed the time zone of the entry. It does not show why the change was made, because the product does not ask for a reason.

- Fixed: a date typed as day/month could be read as month/day in the audit view. The format now follows the account's settings.
- Fixed: the import screen did not keep the sort order after a page reload. It now keeps it until the user changes it.

Known issue: the activity log does not yet support files larger than 100 MB. Larger files are refused with a message that names the limit.

Planned, not yet available: the source file name in scheduled reports. We expect to release it in the first quarter of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 36.1

Quillstone Ledger 36.1 was released on 15 April 2026 for customers on the hosted plan.

What changed:

- Exports: faster on accounts with many records. In our own test on 15 April 2026, a 178,000-row run took 31 seconds on 36.0 and 7 seconds on 36.1. Your results depend on your data.
- The import screen: now accepts the payment terms as a filter. The filter applies to records created on or after 36.1.
- The approval queue: now names the user who last changed the attachment count. It does not show why the change was made, because the product does not ask for a reason.

- Fixed: report generation did not keep the sort order after a page reload. It now keeps it until the user changes it.

Known issue: on slow connections exports can take up to 20 seconds to load. We are working on it and will say here when it is fixed.

Planned, not yet available: the approving user's role in the audit view. We expect to release it in the first quarter of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 36.2

Quillstone Ledger 36.2 was released on 22 April 2026 for customers on the hosted plan.

What changed:

- Attachment storage: now includes a column for the reviewer's comment count. The column is empty for records created before 36.2, because the reviewer's comment count was not stored then.
- The tax summary: now includes a column for the device type. The column is empty for records created before 36.2, because the device type was not stored then.
- The import screen: now accepts the original currency as a filter. The filter applies to records created on or after 36.2.
- The audit view: now shows the cost centre beside each entry. Entries made before 36.2 show a dash, because the value was not recorded.
- Currency conversion: now includes a column for the time zone of the entry. The column is empty for records created before 36.2, because the time zone of the entry was not stored then.

- Fixed: attachment storage did not keep the sort order after a page reload. It now keeps it until the user changes it.

## Quillstone Ledger 36.3

Quillstone Ledger 36.3 was released on 1 May 2026 for customers on the hosted plan.

What changed:

- The activity log: faster on accounts with many records. In our own test on 1 May 2026, a 103,000-row run took 37 seconds on 36.2 and 12 seconds on 36.3. Your results depend on your data.
- The activity log: now keeps the import job number for 365 days. After that the value is removed, and the entry shows when it was removed.
- The mobile view: now keeps the attachment count for 90 days. After that the value is removed, and the entry shows when it was removed.
- Exports: now shows the import job number beside each entry. Entries made before 36.3 show a dash, because the value was not recorded.
- Scheduled reports: faster on accounts with many records. In our own test on 1 May 2026, a 175,000-row run took 20 seconds on 36.2 and 4 seconds on 36.3. Your results depend on your data.

- Fixed: in the audit view, a month with no records showed a blank page. It now shows an empty table.
- Fixed: report generation did not keep the sort order after a page reload. It now keeps it until the user changes it.

## Quillstone Ledger 36.4

Quillstone Ledger 36.4 was released on 8 May 2026 for customers on the hosted plan.

What changed:

- Report generation: now keeps the original currency for 365 days. After that the value is removed, and the entry shows when it was removed.
- The mobile view: now accepts the payment terms as a filter. The filter applies to records created on or after 36.4.
- Invoice matching: faster on accounts with many records. In our own test on 8 May 2026, a 135,000-row run took 30 seconds on 36.3 and 15 seconds on 36.4. Your results depend on your data.
- Account settings: now keeps the attachment count for 180 days. After that the value is removed, and the entry shows when it was removed.
- Invoice matching: now includes a column for the time zone of the entry. The column is empty for records created before 36.4, because the time zone of the entry was not stored then.

- Fixed: in report generation, a month with no records showed a blank page. It now shows an empty table.
- Fixed: in the activity log, a month with no records showed a blank page. It now shows an empty table.

Known issue: on slow connections report generation can take up to 20 seconds to load. We are working on it and will say here when it is fixed.

Planned, not yet available: the source file name in the tax summary. We expect to release it in the first quarter of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 36.5

Quillstone Ledger 36.5 was released on 15 May 2026 for customers on the hosted plan.

What changed:

- Report generation: now shows the original currency beside each entry. Entries made before 36.5 show a dash, because the value was not recorded.
- The import screen: now keeps the batch number for 180 days. After that the value is removed, and the entry shows when it was removed.
- The audit view: can now be turned off for each user. When it is turned off, the reviewer's comment count is not collected for that user from then on.
- The approval queue: can now be turned off for each user. When it is turned off, the reviewer's comment count is not collected for that user from then on.
- Attachment storage: now names the user who last changed the source file name. It does not show why the change was made, because the product does not ask for a reason.

- Fixed: the tax summary did not keep the sort order after a page reload. It now keeps it until the user changes it.

Known issue: on slow connections the import screen can take up to 20 seconds to load. We are working on it and will say here when it is fixed.

## Quillstone Ledger 36.6

Quillstone Ledger 36.6 was released on 22 May 2026 for customers on the hosted plan.

What changed:

- The import screen: now includes a column for the device type. The column is empty for records created before 36.6, because the device type was not stored then.
- The import screen: now shows the reviewer's comment count beside each entry. Entries made before 36.6 show a dash, because the value was not recorded.
- Attachment storage: can now be turned off for each user. When it is turned off, the batch number is not collected for that user from then on.
- Account settings: now keeps the device type for 180 days. After that the value is removed, and the entry shows when it was removed.
- The search index: can now be turned off for each user. When it is turned off, the exchange rate used is not collected for that user from then on.

- Fixed: a date typed as day/month could be read as month/day in the tax summary. The format now follows the account's settings.
- Fixed: in account settings, a month with no records showed a blank page. It now shows an empty table.

Known issue: bank feeds does not yet support files larger than 25 MB. Larger files are refused with a message that names the limit.

## Quillstone Ledger 36.7

Quillstone Ledger 36.7 was released on 1 June 2026 for customers on the hosted plan.

What changed:

- The tax summary: now names the user who last changed the cost centre. It does not show why the change was made, because the product does not ask for a reason.
- The audit view: now includes a column for the batch number. The column is empty for records created before 36.7, because the batch number was not stored then.
- Invoice matching: now shows the source file name beside each entry. Entries made before 36.7 show a dash, because the value was not recorded.
- Currency conversion: now keeps the device type for 365 days. After that the value is removed, and the entry shows when it was removed.
- The mobile view: now names the user who last changed the attachment count. It does not show why the change was made, because the product does not ask for a reason.

- Fixed: report generation did not keep the sort order after a page reload. It now keeps it until the user changes it.
- Fixed: in the search index, a month with no records showed a blank page. It now shows an empty table.

Known issue: on slow connections the audit view can take up to 30 seconds to load. We are working on it and will say here when it is fixed.

## Quillstone Ledger 36.8

Quillstone Ledger 36.8 was released on 8 June 2026 for customers on the hosted plan.

What changed:

- Currency conversion: faster on accounts with many records. In our own test on 8 June 2026, a 118,000-row run took 32 seconds on 36.7 and 8 seconds on 36.8. Your results depend on your data.
- Exports: now accepts the time zone of the entry as a filter. The filter applies to records created on or after 36.8.
- Exports: can now be turned off for each user. When it is turned off, the cost centre is not collected for that user from then on.

- Fixed: totals in scheduled reports were rounded before they were added, which could change the last digit. They are now added first and rounded once.

Planned, not yet available: the import job number in invoice matching. We expect to release it in the first quarter of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 36.9

Quillstone Ledger 36.9 was released on 15 June 2026 for customers on the hosted plan.

What changed:

- Report generation: now keeps the import job number for 180 days. After that the value is removed, and the entry shows when it was removed.
- The mobile view: now shows the attachment count beside each entry. Entries made before 36.9 show a dash, because the value was not recorded.
- Report generation: can now be turned off for each user. When it is turned off, the payment terms is not collected for that user from then on.

- Fixed: totals in invoice matching were rounded before they were added, which could change the last digit. They are now added first and rounded once.

Known issue: on slow connections the approval queue can take up to 30 seconds to load. We are working on it and will say here when it is fixed.

## Quillstone Ledger 37.0

Quillstone Ledger 37.0 was released on 22 June 2026 for customers on the hosted plan.

What changed:

- Scheduled reports: faster on accounts with many records. In our own test on 22 June 2026, a 23,000-row run took 68 seconds on 36.9 and 34 seconds on 37.0. Your results depend on your data.
- Scheduled reports: now keeps the original currency for 30 days. After that the value is removed, and the entry shows when it was removed.
- The tax summary: now accepts the device type as a filter. The filter applies to records created on or after 37.0.
- Report generation: now accepts the exchange rate used as a filter. The filter applies to records created on or after 37.0.
- The import screen: now shows the time zone of the entry beside each entry. Entries made before 37.0 show a dash, because the value was not recorded.

- Fixed: exports did not keep the sort order after a page reload. It now keeps it until the user changes it.
- Fixed: in report generation, a month with no records showed a blank page. It now shows an empty table.

## Quillstone Ledger 37.1

Quillstone Ledger 37.1 was released on 1 July 2026 for customers on the hosted plan.

What changed:

- Scheduled reports: now names the user who last changed the payment terms. It does not show why the change was made, because the product does not ask for a reason.
- The tax summary: now includes a column for the time zone of the entry. The column is empty for records created before 37.1, because the time zone of the entry was not stored then.
- The mobile view: faster on accounts with many records. In our own test on 1 July 2026, a 70,000-row run took 74 seconds on 37.0 and 24 seconds on 37.1. Your results depend on your data.
- The audit view: now shows the attachment count beside each entry. Entries made before 37.1 show a dash, because the value was not recorded.

- Fixed: totals in scheduled reports were rounded before they were added, which could change the last digit. They are now added first and rounded once.

## Quillstone Ledger 37.2

Quillstone Ledger 37.2 was released on 8 July 2026 for customers on the hosted plan.

What changed:

- The mobile view: can now be turned off for each user. When it is turned off, the device type is not collected for that user from then on.
- The mobile view: can now be turned off for each user. When it is turned off, the original currency is not collected for that user from then on.
- Invoice matching: now names the user who last changed the batch number. It does not show why the change was made, because the product does not ask for a reason.
- The import screen: now accepts the source file name as a filter. The filter applies to records created on or after 37.2.

- Fixed: the activity log did not keep the sort order after a page reload. It now keeps it until the user changes it.

## Quillstone Ledger 37.3

Quillstone Ledger 37.3 was released on 15 July 2026 for customers on the hosted plan.

What changed:

- Scheduled reports: faster on accounts with many records. In our own test on 15 July 2026, a 171,000-row run took 71 seconds on 37.2 and 35 seconds on 37.3. Your results depend on your data.
- User roles: now includes a column for the batch number. The column is empty for records created before 37.3, because the batch number was not stored then.
- The tax summary: now shows the cost centre beside each entry. Entries made before 37.3 show a dash, because the value was not recorded.

- Fixed: in the tax summary, a month with no records showed a blank page. It now shows an empty table.
- Fixed: in the approval queue, a month with no records showed a blank page. It now shows an empty table.

## Quillstone Ledger 37.4

Quillstone Ledger 37.4 was released on 22 July 2026 for customers on the hosted plan.

What changed:

- Scheduled reports: now includes a column for the batch number. The column is empty for records created before 37.4, because the batch number was not stored then.
- The audit view: now keeps the exchange rate used for 90 days. After that the value is removed, and the entry shows when it was removed.
- Report generation: faster on accounts with many records. In our own test on 22 July 2026, a 23,000-row run took 68 seconds on 37.3 and 13 seconds on 37.4. Your results depend on your data.
- The mobile view: faster on accounts with many records. In our own test on 22 July 2026, a 121,000-row run took 75 seconds on 37.3 and 37 seconds on 37.4. Your results depend on your data.

- Fixed: the mobile view did not keep the sort order after a page reload. It now keeps it until the user changes it.

## Quillstone Ledger 37.5

Quillstone Ledger 37.5 was released on 1 August 2026 for customers on the hosted plan.

What changed:

- Exports: can now be turned off for each user. When it is turned off, the payment terms is not collected for that user from then on.
- Account settings: faster on accounts with many records. In our own test on 1 August 2026, a 40,000-row run took 60 seconds on 37.4 and 30 seconds on 37.5. Your results depend on your data.
- The audit view: now shows the time zone of the entry beside each entry. Entries made before 37.5 show a dash, because the value was not recorded.

- Fixed: totals in user roles were rounded before they were added, which could change the last digit. They are now added first and rounded once.
- Fixed: in the audit view, a month with no records showed a blank page. It now shows an empty table.

Known issue: the search index does not yet support files larger than 25 MB. Larger files are refused with a message that names the limit.

## Quillstone Ledger 37.6

Quillstone Ledger 37.6 was released on 8 August 2026 for customers on the hosted plan.

What changed:

- The activity log: now names the user who last changed the source file name. It does not show why the change was made, because the product does not ask for a reason.
- The approval queue: now accepts the attachment count as a filter. The filter applies to records created on or after 37.6.
- The activity log: can now be turned off for each user. When it is turned off, the payment terms is not collected for that user from then on.
- The import screen: faster on accounts with many records. In our own test on 8 August 2026, a 119,000-row run took 67 seconds on 37.5 and 22 seconds on 37.6. Your results depend on your data.

- Fixed: a date typed as day/month could be read as month/day in the search index. The format now follows the account's settings.

Planned, not yet available: the device type in exports. We expect to release it in the first quarter of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 37.7

Quillstone Ledger 37.7 was released on 15 August 2026 for customers on the hosted plan.

What changed:

- Invoice matching: now names the user who last changed the original currency. It does not show why the change was made, because the product does not ask for a reason.
- Bank feeds: faster on accounts with many records. In our own test on 15 August 2026, a 136,000-row run took 43 seconds on 37.6 and 8 seconds on 37.7. Your results depend on your data.
- The tax summary: faster on accounts with many records. In our own test on 15 August 2026, a 69,000-row run took 62 seconds on 37.6 and 15 seconds on 37.7. Your results depend on your data.

- Fixed: totals in exports were rounded before they were added, which could change the last digit. They are now added first and rounded once.

## Quillstone Ledger 37.8

Quillstone Ledger 37.8 was released on 22 August 2026 for customers on the hosted plan.

What changed:

- The activity log: can now be turned off for each user. When it is turned off, the approving user's role is not collected for that user from then on.
- User roles: now names the user who last changed the approving user's role. It does not show why the change was made, because the product does not ask for a reason.
- Scheduled reports: now accepts the batch number as a filter. The filter applies to records created on or after 37.8.
- User roles: now accepts the attachment count as a filter. The filter applies to records created on or after 37.8.

- Fixed: a date typed as day/month could be read as month/day in exports. The format now follows the account's settings.
- Fixed: totals in the import screen were rounded before they were added, which could change the last digit. They are now added first and rounded once.

## Quillstone Ledger 37.9

Quillstone Ledger 37.9 was released on 1 September 2026 for customers on the hosted plan.

What changed:

- The import screen: now accepts the time zone of the entry as a filter. The filter applies to records created on or after 37.9.
- Account settings: faster on accounts with many records. In our own test on 1 September 2026, a 174,000-row run took 35 seconds on 37.8 and 11 seconds on 37.9. Your results depend on your data.
- The audit view: can now be turned off for each user. When it is turned off, the attachment count is not collected for that user from then on.
- The mobile view: now includes a column for the attachment count. The column is empty for records created before 37.9, because the attachment count was not stored then.

- Fixed: a date typed as day/month could be read as month/day in bank feeds. The format now follows the account's settings.
- Fixed: a date typed as day/month could be read as month/day in bank feeds. The format now follows the account's settings.

## Quillstone Ledger 38.0

Quillstone Ledger 38.0 was released on 8 September 2026 for customers on the hosted plan.

What changed:

- The search index: faster on accounts with many records. In our own test on 8 September 2026, a 175,000-row run took 33 seconds on 37.9 and 11 seconds on 38.0. Your results depend on your data.
- Account settings: now shows the device type beside each entry. Entries made before 38.0 show a dash, because the value was not recorded.
- The search index: now keeps the batch number for 90 days. After that the value is removed, and the entry shows when it was removed.
- The tax summary: now keeps the batch number for 90 days. After that the value is removed, and the entry shows when it was removed.

- Fixed: exports did not keep the sort order after a page reload. It now keeps it until the user changes it.
- Fixed: totals in account settings were rounded before they were added, which could change the last digit. They are now added first and rounded once.

Known issue: bank feeds does not yet support files larger than 25 MB. Larger files are refused with a message that names the limit.

## Quillstone Ledger 38.1

Quillstone Ledger 38.1 was released on 15 September 2026 for customers on the hosted plan.

What changed:

- Currency conversion: now names the user who last changed the source file name. It does not show why the change was made, because the product does not ask for a reason.
- The approval queue: faster on accounts with many records. In our own test on 15 September 2026, a 160,000-row run took 34 seconds on 38.0 and 17 seconds on 38.1. Your results depend on your data.
- The import screen: now names the user who last changed the cost centre. It does not show why the change was made, because the product does not ask for a reason.

- Fixed: a date typed as day/month could be read as month/day in the activity log. The format now follows the account's settings.

Known issue: the audit view does not yet support files larger than 25 MB. Larger files are refused with a message that names the limit.

Planned, not yet available: the payment terms in currency conversion. We expect to release it in the first quarter of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 38.2

Quillstone Ledger 38.2 was released on 22 September 2026 for customers on the hosted plan.

What changed:

- The activity log: now accepts the reviewer's comment count as a filter. The filter applies to records created on or after 38.2.
- The approval queue: now keeps the import job number for 365 days. After that the value is removed, and the entry shows when it was removed.
- Attachment storage: can now be turned off for each user. When it is turned off, the source file name is not collected for that user from then on.
- The activity log: now shows the payment terms beside each entry. Entries made before 38.2 show a dash, because the value was not recorded.
- Bank feeds: now shows the approving user's role beside each entry. Entries made before 38.2 show a dash, because the value was not recorded.

- Fixed: totals in the tax summary were rounded before they were added, which could change the last digit. They are now added first and rounded once.

Planned, not yet available: the import job number in report generation. We expect to release it in the second quarter of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 38.3

Quillstone Ledger 38.3 was released on 1 October 2026 for customers on the hosted plan.

What changed:

- The audit view: can now be turned off for each user. When it is turned off, the time zone of the entry is not collected for that user from then on.
- User roles: can now be turned off for each user. When it is turned off, the original currency is not collected for that user from then on.
- The import screen: now accepts the cost centre as a filter. The filter applies to records created on or after 38.3.
- Currency conversion: now keeps the cost centre for 180 days. After that the value is removed, and the entry shows when it was removed.

- Fixed: the import screen did not keep the sort order after a page reload. It now keeps it until the user changes it.

Known issue: on slow connections account settings can take up to 30 seconds to load. We are working on it and will say here when it is fixed.

Planned, not yet available: the exchange rate used in user roles. We expect to release it in the second quarter of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 38.4

Quillstone Ledger 38.4 was released on 8 October 2026 for customers on the hosted plan.

What changed:

- User roles: now names the user who last changed the cost centre. It does not show why the change was made, because the product does not ask for a reason.
- The approval queue: now names the user who last changed the device type. It does not show why the change was made, because the product does not ask for a reason.
- Bank feeds: now names the user who last changed the exchange rate used. It does not show why the change was made, because the product does not ask for a reason.

- Fixed: totals in report generation were rounded before they were added, which could change the last digit. They are now added first and rounded once.

Known issue: on slow connections the import screen can take up to 30 seconds to load. We are working on it and will say here when it is fixed.

Planned, not yet available: the batch number in scheduled reports. We expect to release it in the second quarter of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 38.5

Quillstone Ledger 38.5 was released on 15 October 2026 for customers on the hosted plan.

What changed:

- User roles: now accepts the attachment count as a filter. The filter applies to records created on or after 38.5.
- Attachment storage: now shows the time zone of the entry beside each entry. Entries made before 38.5 show a dash, because the value was not recorded.
- The search index: can now be turned off for each user. When it is turned off, the time zone of the entry is not collected for that user from then on.

- Fixed: a date typed as day/month could be read as month/day in the mobile view. The format now follows the account's settings.

## Quillstone Ledger 38.6

Quillstone Ledger 38.6 was released on 22 October 2026 for customers on the hosted plan.

What changed:

- The mobile view: now includes a column for the import job number. The column is empty for records created before 38.6, because the import job number was not stored then.
- The search index: now shows the reviewer's comment count beside each entry. Entries made before 38.6 show a dash, because the value was not recorded.
- Attachment storage: now names the user who last changed the original currency. It does not show why the change was made, because the product does not ask for a reason.

- Fixed: totals in the search index were rounded before they were added, which could change the last digit. They are now added first and rounded once.

Planned, not yet available: the time zone of the entry in exports. We expect to release it in the second half of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 38.7

Quillstone Ledger 38.7 was released on 1 November 2026 for customers on the hosted plan.

What changed:

- The approval queue: now includes a column for the time zone of the entry. The column is empty for records created before 38.7, because the time zone of the entry was not stored then.
- Report generation: can now be turned off for each user. When it is turned off, the approving user's role is not collected for that user from then on.
- The mobile view: faster on accounts with many records. In our own test on 1 November 2026, a 83,000-row run took 32 seconds on 38.6 and 10 seconds on 38.7. Your results depend on your data.

- Fixed: in attachment storage, a month with no records showed a blank page. It now shows an empty table.

Known issue: account settings does not yet support files larger than 50 MB. Larger files are refused with a message that names the limit.

## Quillstone Ledger 38.8

Quillstone Ledger 38.8 was released on 8 November 2026 for customers on the hosted plan.

What changed:

- Invoice matching: now includes a column for the exchange rate used. The column is empty for records created before 38.8, because the exchange rate used was not stored then.
- Scheduled reports: now names the user who last changed the original currency. It does not show why the change was made, because the product does not ask for a reason.
- The import screen: can now be turned off for each user. When it is turned off, the attachment count is not collected for that user from then on.

- Fixed: totals in report generation were rounded before they were added, which could change the last digit. They are now added first and rounded once.
- Fixed: totals in the approval queue were rounded before they were added, which could change the last digit. They are now added first and rounded once.

## Quillstone Ledger 38.9

Quillstone Ledger 38.9 was released on 15 November 2026 for customers on the hosted plan.

What changed:

- The search index: now names the user who last changed the batch number. It does not show why the change was made, because the product does not ask for a reason.
- Bank feeds: can now be turned off for each user. When it is turned off, the attachment count is not collected for that user from then on.
- User roles: faster on accounts with many records. In our own test on 15 November 2026, a 76,000-row run took 90 seconds on 38.8 and 22 seconds on 38.9. Your results depend on your data.

- Fixed: in account settings, a month with no records showed a blank page. It now shows an empty table.
- Fixed: the audit view did not keep the sort order after a page reload. It now keeps it until the user changes it.

Known issue: bank feeds does not yet support files larger than 100 MB. Larger files are refused with a message that names the limit.

Planned, not yet available: the original currency in bank feeds. We expect to release it in the first quarter of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 39.0

Quillstone Ledger 39.0 was released on 22 November 2026 for customers on the hosted plan.

What changed:

- Scheduled reports: now names the user who last changed the payment terms. It does not show why the change was made, because the product does not ask for a reason.
- The activity log: can now be turned off for each user. When it is turned off, the cost centre is not collected for that user from then on.
- The import screen: faster on accounts with many records. In our own test on 22 November 2026, a 120,000-row run took 64 seconds on 38.9 and 12 seconds on 39.0. Your results depend on your data.

- Fixed: totals in invoice matching were rounded before they were added, which could change the last digit. They are now added first and rounded once.

Known issue: the activity log does not yet support files larger than 25 MB. Larger files are refused with a message that names the limit.

Planned, not yet available: the original currency in the tax summary. We expect to release it in the first quarter of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 39.1

Quillstone Ledger 39.1 was released on 1 December 2026 for customers on the hosted plan.

What changed:

- The tax summary: now shows the import job number beside each entry. Entries made before 39.1 show a dash, because the value was not recorded.
- The tax summary: now keeps the time zone of the entry for 30 days. After that the value is removed, and the entry shows when it was removed.
- The tax summary: now includes a column for the original currency. The column is empty for records created before 39.1, because the original currency was not stored then.
- The tax summary: can now be turned off for each user. When it is turned off, the reviewer's comment count is not collected for that user from then on.

- Fixed: in scheduled reports, a month with no records showed a blank page. It now shows an empty table.

## Quillstone Ledger 39.2

Quillstone Ledger 39.2 was released on 8 December 2026 for customers on the hosted plan.

What changed:

- Exports: now keeps the cost centre for 365 days. After that the value is removed, and the entry shows when it was removed.
- The tax summary: now keeps the cost centre for 365 days. After that the value is removed, and the entry shows when it was removed.
- The activity log: now shows the attachment count beside each entry. Entries made before 39.2 show a dash, because the value was not recorded.
- The audit view: now names the user who last changed the time zone of the entry. It does not show why the change was made, because the product does not ask for a reason.
- The mobile view: now includes a column for the attachment count. The column is empty for records created before 39.2, because the attachment count was not stored then.

- Fixed: in the tax summary, a month with no records showed a blank page. It now shows an empty table.

## Quillstone Ledger 39.3

Quillstone Ledger 39.3 was released on 15 December 2026 for customers on the hosted plan.

What changed:

- The activity log: now keeps the time zone of the entry for 90 days. After that the value is removed, and the entry shows when it was removed.
- Account settings: faster on accounts with many records. In our own test on 15 December 2026, a 116,000-row run took 77 seconds on 39.2 and 38 seconds on 39.3. Your results depend on your data.
- The tax summary: faster on accounts with many records. In our own test on 15 December 2026, a 61,000-row run took 26 seconds on 39.2 and 13 seconds on 39.3. Your results depend on your data.

- Fixed: totals in the activity log were rounded before they were added, which could change the last digit. They are now added first and rounded once.
- Fixed: a date typed as day/month could be read as month/day in scheduled reports. The format now follows the account's settings.

Planned, not yet available: the approving user's role in the audit view. We expect to release it in the second quarter of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 39.4

Quillstone Ledger 39.4 was released on 22 December 2026 for customers on the hosted plan.

What changed:

- Report generation: now accepts the attachment count as a filter. The filter applies to records created on or after 39.4.
- Attachment storage: can now be turned off for each user. When it is turned off, the payment terms is not collected for that user from then on.
- Account settings: now shows the reviewer's comment count beside each entry. Entries made before 39.4 show a dash, because the value was not recorded.
- The approval queue: now includes a column for the import job number. The column is empty for records created before 39.4, because the import job number was not stored then.
- User roles: now names the user who last changed the reviewer's comment count. It does not show why the change was made, because the product does not ask for a reason.

- Fixed: a date typed as day/month could be read as month/day in currency conversion. The format now follows the account's settings.

Planned, not yet available: the reviewer's comment count in the search index. We expect to release it in the first quarter of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 39.5

Quillstone Ledger 39.5 was released on 1 January 2027 for customers on the hosted plan.

What changed:

- Bank feeds: now accepts the reviewer's comment count as a filter. The filter applies to records created on or after 39.5.
- The search index: faster on accounts with many records. In our own test on 1 January 2027, a 49,000-row run took 64 seconds on 39.4 and 21 seconds on 39.5. Your results depend on your data.
- The tax summary: now names the user who last changed the cost centre. It does not show why the change was made, because the product does not ask for a reason.

- Fixed: totals in currency conversion were rounded before they were added, which could change the last digit. They are now added first and rounded once.
- Fixed: the approval queue did not keep the sort order after a page reload. It now keeps it until the user changes it.

Known issue: on slow connections the activity log can take up to 30 seconds to load. We are working on it and will say here when it is fixed.

## Quillstone Ledger 39.6

Quillstone Ledger 39.6 was released on 8 January 2027 for customers on the hosted plan.

What changed:

- User roles: can now be turned off for each user. When it is turned off, the payment terms is not collected for that user from then on.
- Report generation: now includes a column for the cost centre. The column is empty for records created before 39.6, because the cost centre was not stored then.
- The activity log: now includes a column for the approving user's role. The column is empty for records created before 39.6, because the approving user's role was not stored then.

- Fixed: in the tax summary, a month with no records showed a blank page. It now shows an empty table.

Known issue: attachment storage does not yet support files larger than 100 MB. Larger files are refused with a message that names the limit.

## Quillstone Ledger 39.7

Quillstone Ledger 39.7 was released on 15 January 2027 for customers on the hosted plan.

What changed:

- Account settings: can now be turned off for each user. When it is turned off, the import job number is not collected for that user from then on.
- The activity log: now accepts the time zone of the entry as a filter. The filter applies to records created on or after 39.7.
- Invoice matching: now accepts the reviewer's comment count as a filter. The filter applies to records created on or after 39.7.
- Currency conversion: faster on accounts with many records. In our own test on 15 January 2027, a 64,000-row run took 68 seconds on 39.6 and 34 seconds on 39.7. Your results depend on your data.
- Invoice matching: now keeps the cost centre for 90 days. After that the value is removed, and the entry shows when it was removed.

- Fixed: a date typed as day/month could be read as month/day in the search index. The format now follows the account's settings.

## Quillstone Ledger 39.8

Quillstone Ledger 39.8 was released on 22 January 2027 for customers on the hosted plan.

What changed:

- The mobile view: now keeps the original currency for 30 days. After that the value is removed, and the entry shows when it was removed.
- Attachment storage: now names the user who last changed the exchange rate used. It does not show why the change was made, because the product does not ask for a reason.
- User roles: faster on accounts with many records. In our own test on 22 January 2027, a 30,000-row run took 72 seconds on 39.7 and 24 seconds on 39.8. Your results depend on your data.

- Fixed: the search index did not keep the sort order after a page reload. It now keeps it until the user changes it.

Known issue: account settings does not yet support files larger than 50 MB. Larger files are refused with a message that names the limit.

## Quillstone Ledger 39.9

Quillstone Ledger 39.9 was released on 1 February 2027 for customers on the hosted plan.

What changed:

- The mobile view: now shows the attachment count beside each entry. Entries made before 39.9 show a dash, because the value was not recorded.
- The activity log: now includes a column for the device type. The column is empty for records created before 39.9, because the device type was not stored then.
- Report generation: now includes a column for the exchange rate used. The column is empty for records created before 39.9, because the exchange rate used was not stored then.

- Fixed: a date typed as day/month could be read as month/day in invoice matching. The format now follows the account's settings.

Known issue: account settings does not yet support files larger than 100 MB. Larger files are refused with a message that names the limit.

Planned, not yet available: the reviewer's comment count in account settings. We expect to release it in the first quarter of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 40.0

Quillstone Ledger 40.0 was released on 8 February 2027 for customers on the hosted plan.

What changed:

- User roles: now accepts the source file name as a filter. The filter applies to records created on or after 40.0.
- The mobile view: faster on accounts with many records. In our own test on 8 February 2027, a 173,000-row run took 74 seconds on 39.9 and 14 seconds on 40.0. Your results depend on your data.
- Report generation: now names the user who last changed the reviewer's comment count. It does not show why the change was made, because the product does not ask for a reason.
- The search index: can now be turned off for each user. When it is turned off, the payment terms is not collected for that user from then on.

- Fixed: a date typed as day/month could be read as month/day in invoice matching. The format now follows the account's settings.

Known issue: on slow connections user roles can take up to 10 seconds to load. We are working on it and will say here when it is fixed.

## Quillstone Ledger 40.1

Quillstone Ledger 40.1 was released on 15 February 2027 for customers on the hosted plan.

What changed:

- Invoice matching: now accepts the source file name as a filter. The filter applies to records created on or after 40.1.
- The audit view: now names the user who last changed the device type. It does not show why the change was made, because the product does not ask for a reason.
- The mobile view: now names the user who last changed the attachment count. It does not show why the change was made, because the product does not ask for a reason.
- The activity log: now accepts the payment terms as a filter. The filter applies to records created on or after 40.1.

- Fixed: currency conversion did not keep the sort order after a page reload. It now keeps it until the user changes it.
- Fixed: in report generation, a month with no records showed a blank page. It now shows an empty table.

## Quillstone Ledger 40.2

Quillstone Ledger 40.2 was released on 22 February 2027 for customers on the hosted plan.

What changed:

- Scheduled reports: can now be turned off for each user. When it is turned off, the payment terms is not collected for that user from then on.
- Attachment storage: now names the user who last changed the device type. It does not show why the change was made, because the product does not ask for a reason.
- Report generation: now names the user who last changed the original currency. It does not show why the change was made, because the product does not ask for a reason.
- Currency conversion: now accepts the original currency as a filter. The filter applies to records created on or after 40.2.

- Fixed: invoice matching did not keep the sort order after a page reload. It now keeps it until the user changes it.

## Quillstone Ledger 40.3

Quillstone Ledger 40.3 was released on 1 March 2027 for customers on the hosted plan.

What changed:

- The approval queue: now names the user who last changed the batch number. It does not show why the change was made, because the product does not ask for a reason.
- User roles: can now be turned off for each user. When it is turned off, the batch number is not collected for that user from then on.
- The tax summary: now shows the source file name beside each entry. Entries made before 40.3 show a dash, because the value was not recorded.
- Report generation: now names the user who last changed the device type. It does not show why the change was made, because the product does not ask for a reason.
- The mobile view: now names the user who last changed the approving user's role. It does not show why the change was made, because the product does not ask for a reason.

- Fixed: the search index did not keep the sort order after a page reload. It now keeps it until the user changes it.

Known issue: the mobile view does not yet support files larger than 50 MB. Larger files are refused with a message that names the limit.

## Quillstone Ledger 40.4

Quillstone Ledger 40.4 was released on 8 March 2027 for customers on the hosted plan.

What changed:

- Attachment storage: faster on accounts with many records. In our own test on 8 March 2027, a 76,000-row run took 29 seconds on 40.3 and 9 seconds on 40.4. Your results depend on your data.
- Invoice matching: now shows the approving user's role beside each entry. Entries made before 40.4 show a dash, because the value was not recorded.
- Report generation: now names the user who last changed the device type. It does not show why the change was made, because the product does not ask for a reason.
- The import screen: faster on accounts with many records. In our own test on 8 March 2027, a 27,000-row run took 30 seconds on 40.3 and 6 seconds on 40.4. Your results depend on your data.
- The search index: now names the user who last changed the batch number. It does not show why the change was made, because the product does not ask for a reason.

- Fixed: in report generation, a month with no records showed a blank page. It now shows an empty table.

Known issue: on slow connections the mobile view can take up to 10 seconds to load. We are working on it and will say here when it is fixed.

## Quillstone Ledger 40.5

Quillstone Ledger 40.5 was released on 15 March 2027 for customers on the hosted plan.

What changed:

- Report generation: now includes a column for the batch number. The column is empty for records created before 40.5, because the batch number was not stored then.
- The mobile view: now keeps the exchange rate used for 365 days. After that the value is removed, and the entry shows when it was removed.
- The tax summary: now includes a column for the original currency. The column is empty for records created before 40.5, because the original currency was not stored then.
- Attachment storage: now shows the approving user's role beside each entry. Entries made before 40.5 show a dash, because the value was not recorded.

- Fixed: in the mobile view, a month with no records showed a blank page. It now shows an empty table.

Known issue: bank feeds does not yet support files larger than 50 MB. Larger files are refused with a message that names the limit.

## Quillstone Ledger 40.6

Quillstone Ledger 40.6 was released on 22 March 2027 for customers on the hosted plan.

What changed:

- Exports: can now be turned off for each user. When it is turned off, the payment terms is not collected for that user from then on.
- Scheduled reports: faster on accounts with many records. In our own test on 22 March 2027, a 191,000-row run took 77 seconds on 40.5 and 38 seconds on 40.6. Your results depend on your data.
- User roles: can now be turned off for each user. When it is turned off, the original currency is not collected for that user from then on.

- Fixed: totals in scheduled reports were rounded before they were added, which could change the last digit. They are now added first and rounded once.

## Quillstone Ledger 40.7

Quillstone Ledger 40.7 was released on 1 April 2027 for customers on the hosted plan.

What changed:

- Report generation: now accepts the payment terms as a filter. The filter applies to records created on or after 40.7.
- Exports: now accepts the device type as a filter. The filter applies to records created on or after 40.7.
- The import screen: now names the user who last changed the device type. It does not show why the change was made, because the product does not ask for a reason.
- Report generation: now accepts the device type as a filter. The filter applies to records created on or after 40.7.

- Fixed: a date typed as day/month could be read as month/day in bank feeds. The format now follows the account's settings.
- Fixed: in the approval queue, a month with no records showed a blank page. It now shows an empty table.

## Quillstone Ledger 40.8

Quillstone Ledger 40.8 was released on 8 April 2027 for customers on the hosted plan.

What changed:

- Exports: now names the user who last changed the cost centre. It does not show why the change was made, because the product does not ask for a reason.
- Attachment storage: faster on accounts with many records. In our own test on 8 April 2027, a 16,000-row run took 63 seconds on 40.7 and 21 seconds on 40.8. Your results depend on your data.
- Attachment storage: faster on accounts with many records. In our own test on 8 April 2027, a 163,000-row run took 58 seconds on 40.7 and 29 seconds on 40.8. Your results depend on your data.
- The activity log: faster on accounts with many records. In our own test on 8 April 2027, a 155,000-row run took 27 seconds on 40.7 and 6 seconds on 40.8. Your results depend on your data.
- The search index: faster on accounts with many records. In our own test on 8 April 2027, a 70,000-row run took 32 seconds on 40.7 and 10 seconds on 40.8. Your results depend on your data.

- Fixed: totals in user roles were rounded before they were added, which could change the last digit. They are now added first and rounded once.
- Fixed: a date typed as day/month could be read as month/day in the approval queue. The format now follows the account's settings.

## Quillstone Ledger 40.9

Quillstone Ledger 40.9 was released on 15 April 2027 for customers on the hosted plan.

What changed:

- The mobile view: faster on accounts with many records. In our own test on 15 April 2027, a 171,000-row run took 41 seconds on 40.8 and 10 seconds on 40.9. Your results depend on your data.
- The import screen: now shows the attachment count beside each entry. Entries made before 40.9 show a dash, because the value was not recorded.
- Account settings: now includes a column for the attachment count. The column is empty for records created before 40.9, because the attachment count was not stored then.

- Fixed: totals in the mobile view were rounded before they were added, which could change the last digit. They are now added first and rounded once.
- Fixed: in currency conversion, a month with no records showed a blank page. It now shows an empty table.

## Quillstone Ledger 41.0

Quillstone Ledger 41.0 was released on 22 April 2027 for customers on the hosted plan.

What changed:

- Invoice matching: now accepts the time zone of the entry as a filter. The filter applies to records created on or after 41.0.
- The tax summary: now includes a column for the import job number. The column is empty for records created before 41.0, because the import job number was not stored then.
- Invoice matching: now includes a column for the payment terms. The column is empty for records created before 41.0, because the payment terms was not stored then.
- The tax summary: now shows the approving user's role beside each entry. Entries made before 41.0 show a dash, because the value was not recorded.

- Fixed: totals in invoice matching were rounded before they were added, which could change the last digit. They are now added first and rounded once.

## Quillstone Ledger 41.1

Quillstone Ledger 41.1 was released on 1 May 2027 for customers on the hosted plan.

What changed:

- Currency conversion: can now be turned off for each user. When it is turned off, the batch number is not collected for that user from then on.
- Exports: now shows the exchange rate used beside each entry. Entries made before 41.1 show a dash, because the value was not recorded.
- User roles: now keeps the reviewer's comment count for 365 days. After that the value is removed, and the entry shows when it was removed.

- Fixed: totals in the approval queue were rounded before they were added, which could change the last digit. They are now added first and rounded once.

## Quillstone Ledger 41.2

Quillstone Ledger 41.2 was released on 8 May 2027 for customers on the hosted plan.

What changed:

- The activity log: now names the user who last changed the device type. It does not show why the change was made, because the product does not ask for a reason.
- Attachment storage: now keeps the original currency for 30 days. After that the value is removed, and the entry shows when it was removed.
- The import screen: now accepts the attachment count as a filter. The filter applies to records created on or after 41.2.

- Fixed: a date typed as day/month could be read as month/day in currency conversion. The format now follows the account's settings.
- Fixed: invoice matching did not keep the sort order after a page reload. It now keeps it until the user changes it.

Planned, not yet available: the attachment count in bank feeds. We expect to release it in the second half of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 41.3

Quillstone Ledger 41.3 was released on 15 May 2027 for customers on the hosted plan.

What changed:

- Attachment storage: now shows the source file name beside each entry. Entries made before 41.3 show a dash, because the value was not recorded.
- The mobile view: now keeps the source file name for 365 days. After that the value is removed, and the entry shows when it was removed.
- The activity log: now names the user who last changed the batch number. It does not show why the change was made, because the product does not ask for a reason.
- The mobile view: can now be turned off for each user. When it is turned off, the device type is not collected for that user from then on.
- Invoice matching: now accepts the approving user's role as a filter. The filter applies to records created on or after 41.3.

- Fixed: totals in user roles were rounded before they were added, which could change the last digit. They are now added first and rounded once.
- Fixed: totals in invoice matching were rounded before they were added, which could change the last digit. They are now added first and rounded once.

## Quillstone Ledger 41.4

Quillstone Ledger 41.4 was released on 22 May 2027 for customers on the hosted plan.

What changed:

- Invoice matching: now keeps the device type for 90 days. After that the value is removed, and the entry shows when it was removed.
- The approval queue: now names the user who last changed the exchange rate used. It does not show why the change was made, because the product does not ask for a reason.
- The approval queue: now accepts the cost centre as a filter. The filter applies to records created on or after 41.4.
- Exports: now shows the cost centre beside each entry. Entries made before 41.4 show a dash, because the value was not recorded.

- Fixed: totals in the import screen were rounded before they were added, which could change the last digit. They are now added first and rounded once.

## Quillstone Ledger 41.5

Quillstone Ledger 41.5 was released on 1 June 2027 for customers on the hosted plan.

What changed:

- Invoice matching: now names the user who last changed the exchange rate used. It does not show why the change was made, because the product does not ask for a reason.
- The search index: now names the user who last changed the import job number. It does not show why the change was made, because the product does not ask for a reason.
- The audit view: now names the user who last changed the attachment count. It does not show why the change was made, because the product does not ask for a reason.

- Fixed: a date typed as day/month could be read as month/day in the activity log. The format now follows the account's settings.
- Fixed: totals in report generation were rounded before they were added, which could change the last digit. They are now added first and rounded once.

Known issue: the import screen does not yet support files larger than 25 MB. Larger files are refused with a message that names the limit.

## Quillstone Ledger 41.6

Quillstone Ledger 41.6 was released on 8 June 2027 for customers on the hosted plan.

What changed:

- Attachment storage: now names the user who last changed the cost centre. It does not show why the change was made, because the product does not ask for a reason.
- Scheduled reports: can now be turned off for each user. When it is turned off, the original currency is not collected for that user from then on.
- Report generation: can now be turned off for each user. When it is turned off, the original currency is not collected for that user from then on.
- Scheduled reports: now names the user who last changed the source file name. It does not show why the change was made, because the product does not ask for a reason.

- Fixed: a date typed as day/month could be read as month/day in currency conversion. The format now follows the account's settings.
- Fixed: a date typed as day/month could be read as month/day in invoice matching. The format now follows the account's settings.

Known issue: on slow connections bank feeds can take up to 10 seconds to load. We are working on it and will say here when it is fixed.

## Quillstone Ledger 41.7

Quillstone Ledger 41.7 was released on 15 June 2027 for customers on the hosted plan.

What changed:

- The import screen: now accepts the cost centre as a filter. The filter applies to records created on or after 41.7.
- Exports: now shows the import job number beside each entry. Entries made before 41.7 show a dash, because the value was not recorded.
- Report generation: faster on accounts with many records. In our own test on 15 June 2027, a 177,000-row run took 64 seconds on 41.6 and 21 seconds on 41.7. Your results depend on your data.

- Fixed: in the approval queue, a month with no records showed a blank page. It now shows an empty table.

Known issue: on slow connections the search index can take up to 20 seconds to load. We are working on it and will say here when it is fixed.

## Quillstone Ledger 41.8

Quillstone Ledger 41.8 was released on 22 June 2027 for customers on the hosted plan.

What changed:

- Account settings: now shows the time zone of the entry beside each entry. Entries made before 41.8 show a dash, because the value was not recorded.
- The search index: now names the user who last changed the approving user's role. It does not show why the change was made, because the product does not ask for a reason.
- Currency conversion: now includes a column for the time zone of the entry. The column is empty for records created before 41.8, because the time zone of the entry was not stored then.
- The tax summary: now shows the original currency beside each entry. Entries made before 41.8 show a dash, because the value was not recorded.

- Fixed: in the activity log, a month with no records showed a blank page. It now shows an empty table.
- Fixed: in scheduled reports, a month with no records showed a blank page. It now shows an empty table.

Known issue: on slow connections account settings can take up to 20 seconds to load. We are working on it and will say here when it is fixed.

## Quillstone Ledger 41.9

Quillstone Ledger 41.9 was released on 1 July 2027 for customers on the hosted plan.

What changed:

- Exports: now includes a column for the attachment count. The column is empty for records created before 41.9, because the attachment count was not stored then.
- The import screen: now keeps the approving user's role for 90 days. After that the value is removed, and the entry shows when it was removed.
- The approval queue: faster on accounts with many records. In our own test on 1 July 2027, a 194,000-row run took 31 seconds on 41.8 and 6 seconds on 41.9. Your results depend on your data.
- The audit view: faster on accounts with many records. In our own test on 1 July 2027, a 54,000-row run took 40 seconds on 41.8 and 10 seconds on 41.9. Your results depend on your data.
- Report generation: now includes a column for the reviewer's comment count. The column is empty for records created before 41.9, because the reviewer's comment count was not stored then.

- Fixed: currency conversion did not keep the sort order after a page reload. It now keeps it until the user changes it.

## Quillstone Ledger 42.0

Quillstone Ledger 42.0 was released on 8 July 2027 for customers on the hosted plan.

What changed:

- The audit view: now names the user who last changed the original currency. It does not show why the change was made, because the product does not ask for a reason.
- Report generation: now accepts the attachment count as a filter. The filter applies to records created on or after 42.0.
- The approval queue: faster on accounts with many records. In our own test on 8 July 2027, a 55,000-row run took 87 seconds on 41.9 and 21 seconds on 42.0. Your results depend on your data.
- The search index: now keeps the exchange rate used for 30 days. After that the value is removed, and the entry shows when it was removed.

- Fixed: totals in scheduled reports were rounded before they were added, which could change the last digit. They are now added first and rounded once.
- Fixed: in the import screen, a month with no records showed a blank page. It now shows an empty table.

## Quillstone Ledger 42.1

Quillstone Ledger 42.1 was released on 15 July 2027 for customers on the hosted plan.

What changed:

- The mobile view: now keeps the approving user's role for 180 days. After that the value is removed, and the entry shows when it was removed.
- Currency conversion: faster on accounts with many records. In our own test on 15 July 2027, a 166,000-row run took 90 seconds on 42.0 and 30 seconds on 42.1. Your results depend on your data.
- The search index: now shows the original currency beside each entry. Entries made before 42.1 show a dash, because the value was not recorded.
- The approval queue: now keeps the attachment count for 365 days. After that the value is removed, and the entry shows when it was removed.
- User roles: now keeps the reviewer's comment count for 365 days. After that the value is removed, and the entry shows when it was removed.

- Fixed: in the approval queue, a month with no records showed a blank page. It now shows an empty table.

Planned, not yet available: the import job number in invoice matching. We expect to release it in the second quarter of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 42.2

Quillstone Ledger 42.2 was released on 22 July 2027 for customers on the hosted plan.

What changed:

- The activity log: now includes a column for the approving user's role. The column is empty for records created before 42.2, because the approving user's role was not stored then.
- The search index: now keeps the payment terms for 365 days. After that the value is removed, and the entry shows when it was removed.
- Bank feeds: now shows the exchange rate used beside each entry. Entries made before 42.2 show a dash, because the value was not recorded.
- Report generation: faster on accounts with many records. In our own test on 22 July 2027, a 196,000-row run took 80 seconds on 42.1 and 20 seconds on 42.2. Your results depend on your data.

- Fixed: a date typed as day/month could be read as month/day in the tax summary. The format now follows the account's settings.

## Quillstone Ledger 42.3

Quillstone Ledger 42.3 was released on 1 August 2027 for customers on the hosted plan.

What changed:

- Bank feeds: now shows the time zone of the entry beside each entry. Entries made before 42.3 show a dash, because the value was not recorded.
- Currency conversion: now keeps the approving user's role for 365 days. After that the value is removed, and the entry shows when it was removed.
- User roles: now accepts the approving user's role as a filter. The filter applies to records created on or after 42.3.

- Fixed: a date typed as day/month could be read as month/day in the audit view. The format now follows the account's settings.
- Fixed: a date typed as day/month could be read as month/day in user roles. The format now follows the account's settings.

## Quillstone Ledger 42.4

Quillstone Ledger 42.4 was released on 8 August 2027 for customers on the hosted plan.

What changed:

- Invoice matching: now accepts the cost centre as a filter. The filter applies to records created on or after 42.4.
- Scheduled reports: now accepts the approving user's role as a filter. The filter applies to records created on or after 42.4.
- The audit view: faster on accounts with many records. In our own test on 8 August 2027, a 72,000-row run took 27 seconds on 42.3 and 5 seconds on 42.4. Your results depend on your data.

- Fixed: totals in invoice matching were rounded before they were added, which could change the last digit. They are now added first and rounded once.

Known issue: the audit view does not yet support files larger than 25 MB. Larger files are refused with a message that names the limit.

## Quillstone Ledger 42.5

Quillstone Ledger 42.5 was released on 15 August 2027 for customers on the hosted plan.

What changed:

- Currency conversion: now names the user who last changed the reviewer's comment count. It does not show why the change was made, because the product does not ask for a reason.
- The activity log: now names the user who last changed the import job number. It does not show why the change was made, because the product does not ask for a reason.
- The approval queue: now includes a column for the source file name. The column is empty for records created before 42.5, because the source file name was not stored then.
- The audit view: now shows the payment terms beside each entry. Entries made before 42.5 show a dash, because the value was not recorded.
- Bank feeds: now shows the cost centre beside each entry. Entries made before 42.5 show a dash, because the value was not recorded.

- Fixed: totals in attachment storage were rounded before they were added, which could change the last digit. They are now added first and rounded once.

## Quillstone Ledger 42.6

Quillstone Ledger 42.6 was released on 22 August 2027 for customers on the hosted plan.

What changed:

- Invoice matching: now keeps the exchange rate used for 180 days. After that the value is removed, and the entry shows when it was removed.
- The audit view: now shows the exchange rate used beside each entry. Entries made before 42.6 show a dash, because the value was not recorded.
- Report generation: now keeps the exchange rate used for 365 days. After that the value is removed, and the entry shows when it was removed.
- Currency conversion: now keeps the time zone of the entry for 365 days. After that the value is removed, and the entry shows when it was removed.
- Invoice matching: can now be turned off for each user. When it is turned off, the source file name is not collected for that user from then on.

- Fixed: in the import screen, a month with no records showed a blank page. It now shows an empty table.

## Quillstone Ledger 42.7

Quillstone Ledger 42.7 was released on 1 September 2027 for customers on the hosted plan.

What changed:

- Bank feeds: now names the user who last changed the attachment count. It does not show why the change was made, because the product does not ask for a reason.
- The approval queue: now includes a column for the batch number. The column is empty for records created before 42.7, because the batch number was not stored then.
- The import screen: now shows the cost centre beside each entry. Entries made before 42.7 show a dash, because the value was not recorded.

- Fixed: in user roles, a month with no records showed a blank page. It now shows an empty table.

Known issue: on slow connections bank feeds can take up to 30 seconds to load. We are working on it and will say here when it is fixed.

## Quillstone Ledger 42.8

Quillstone Ledger 42.8 was released on 8 September 2027 for customers on the hosted plan.

What changed:

- The audit view: now shows the cost centre beside each entry. Entries made before 42.8 show a dash, because the value was not recorded.
- Account settings: now keeps the exchange rate used for 90 days. After that the value is removed, and the entry shows when it was removed.
- Invoice matching: now shows the cost centre beside each entry. Entries made before 42.8 show a dash, because the value was not recorded.
- Currency conversion: now names the user who last changed the reviewer's comment count. It does not show why the change was made, because the product does not ask for a reason.
- The mobile view: can now be turned off for each user. When it is turned off, the original currency is not collected for that user from then on.

- Fixed: totals in the import screen were rounded before they were added, which could change the last digit. They are now added first and rounded once.

Planned, not yet available: the attachment count in exports. We expect to release it in the second quarter of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 42.9

Quillstone Ledger 42.9 was released on 15 September 2027 for customers on the hosted plan.

What changed:

- The approval queue: can now be turned off for each user. When it is turned off, the exchange rate used is not collected for that user from then on.
- The import screen: now shows the reviewer's comment count beside each entry. Entries made before 42.9 show a dash, because the value was not recorded.
- Report generation: now includes a column for the import job number. The column is empty for records created before 42.9, because the import job number was not stored then.

- Fixed: totals in account settings were rounded before they were added, which could change the last digit. They are now added first and rounded once.

Known issue: on slow connections scheduled reports can take up to 10 seconds to load. We are working on it and will say here when it is fixed.

## Quillstone Ledger 43.0

Quillstone Ledger 43.0 was released on 22 September 2027 for customers on the hosted plan.

What changed:

- Scheduled reports: faster on accounts with many records. In our own test on 22 September 2027, a 16,000-row run took 58 seconds on 42.9 and 14 seconds on 43.0. Your results depend on your data.
- Currency conversion: now shows the reviewer's comment count beside each entry. Entries made before 43.0 show a dash, because the value was not recorded.
- The mobile view: now shows the import job number beside each entry. Entries made before 43.0 show a dash, because the value was not recorded.
- Exports: faster on accounts with many records. In our own test on 22 September 2027, a 79,000-row run took 38 seconds on 42.9 and 12 seconds on 43.0. Your results depend on your data.
- Bank feeds: now keeps the attachment count for 180 days. After that the value is removed, and the entry shows when it was removed.

- Fixed: in the activity log, a month with no records showed a blank page. It now shows an empty table.
- Fixed: a date typed as day/month could be read as month/day in user roles. The format now follows the account's settings.

## Quillstone Ledger 43.1

Quillstone Ledger 43.1 was released on 1 October 2027 for customers on the hosted plan.

What changed:

- The activity log: now includes a column for the batch number. The column is empty for records created before 43.1, because the batch number was not stored then.
- Account settings: now keeps the import job number for 30 days. After that the value is removed, and the entry shows when it was removed.
- The tax summary: can now be turned off for each user. When it is turned off, the batch number is not collected for that user from then on.
- User roles: now shows the source file name beside each entry. Entries made before 43.1 show a dash, because the value was not recorded.

- Fixed: in the import screen, a month with no records showed a blank page. It now shows an empty table.

## Quillstone Ledger 43.2

Quillstone Ledger 43.2 was released on 8 October 2027 for customers on the hosted plan.

What changed:

- Account settings: now accepts the batch number as a filter. The filter applies to records created on or after 43.2.
- The import screen: now accepts the original currency as a filter. The filter applies to records created on or after 43.2.
- The approval queue: can now be turned off for each user. When it is turned off, the approving user's role is not collected for that user from then on.
- The search index: now accepts the cost centre as a filter. The filter applies to records created on or after 43.2.

- Fixed: attachment storage did not keep the sort order after a page reload. It now keeps it until the user changes it.

Planned, not yet available: the cost centre in the audit view. We expect to release it in the second quarter of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 43.3

Quillstone Ledger 43.3 was released on 15 October 2027 for customers on the hosted plan.

What changed:

- The approval queue: now accepts the original currency as a filter. The filter applies to records created on or after 43.3.
- Report generation: now keeps the original currency for 180 days. After that the value is removed, and the entry shows when it was removed.
- Account settings: can now be turned off for each user. When it is turned off, the exchange rate used is not collected for that user from then on.
- The activity log: now names the user who last changed the reviewer's comment count. It does not show why the change was made, because the product does not ask for a reason.

- Fixed: the search index did not keep the sort order after a page reload. It now keeps it until the user changes it.

Planned, not yet available: the batch number in invoice matching. We expect to release it in the first quarter of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 43.4

Quillstone Ledger 43.4 was released on 22 October 2027 for customers on the hosted plan.

What changed:

- Account settings: faster on accounts with many records. In our own test on 22 October 2027, a 141,000-row run took 49 seconds on 43.3 and 16 seconds on 43.4. Your results depend on your data.
- The search index: now includes a column for the exchange rate used. The column is empty for records created before 43.4, because the exchange rate used was not stored then.
- The import screen: now keeps the device type for 30 days. After that the value is removed, and the entry shows when it was removed.
- The approval queue: faster on accounts with many records. In our own test on 22 October 2027, a 13,000-row run took 90 seconds on 43.3 and 30 seconds on 43.4. Your results depend on your data.
- Report generation: now includes a column for the batch number. The column is empty for records created before 43.4, because the batch number was not stored then.

- Fixed: in account settings, a month with no records showed a blank page. It now shows an empty table.
- Fixed: a date typed as day/month could be read as month/day in the mobile view. The format now follows the account's settings.

Known issue: on slow connections the approval queue can take up to 30 seconds to load. We are working on it and will say here when it is fixed.

## Quillstone Ledger 43.5

Quillstone Ledger 43.5 was released on 1 November 2027 for customers on the hosted plan.

What changed:

- Attachment storage: now names the user who last changed the import job number. It does not show why the change was made, because the product does not ask for a reason.
- The activity log: now names the user who last changed the payment terms. It does not show why the change was made, because the product does not ask for a reason.
- Bank feeds: now accepts the reviewer's comment count as a filter. The filter applies to records created on or after 43.5.
- Exports: now names the user who last changed the import job number. It does not show why the change was made, because the product does not ask for a reason.

- Fixed: totals in account settings were rounded before they were added, which could change the last digit. They are now added first and rounded once.

Known issue: the import screen does not yet support files larger than 100 MB. Larger files are refused with a message that names the limit.

## Quillstone Ledger 43.6

Quillstone Ledger 43.6 was released on 8 November 2027 for customers on the hosted plan.

What changed:

- Account settings: now names the user who last changed the import job number. It does not show why the change was made, because the product does not ask for a reason.
- The activity log: now keeps the device type for 365 days. After that the value is removed, and the entry shows when it was removed.
- Exports: can now be turned off for each user. When it is turned off, the import job number is not collected for that user from then on.
- Report generation: now names the user who last changed the source file name. It does not show why the change was made, because the product does not ask for a reason.

- Fixed: report generation did not keep the sort order after a page reload. It now keeps it until the user changes it.
- Fixed: a date typed as day/month could be read as month/day in the import screen. The format now follows the account's settings.

Known issue: account settings does not yet support files larger than 25 MB. Larger files are refused with a message that names the limit.

Planned, not yet available: the original currency in report generation. We expect to release it in the second half of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 43.7

Quillstone Ledger 43.7 was released on 15 November 2027 for customers on the hosted plan.

What changed:

- Invoice matching: faster on accounts with many records. In our own test on 15 November 2027, a 105,000-row run took 81 seconds on 43.6 and 40 seconds on 43.7. Your results depend on your data.
- The activity log: now shows the attachment count beside each entry. Entries made before 43.7 show a dash, because the value was not recorded.
- Exports: now accepts the payment terms as a filter. The filter applies to records created on or after 43.7.
- The activity log: now includes a column for the time zone of the entry. The column is empty for records created before 43.7, because the time zone of the entry was not stored then.
- User roles: now shows the original currency beside each entry. Entries made before 43.7 show a dash, because the value was not recorded.

- Fixed: bank feeds did not keep the sort order after a page reload. It now keeps it until the user changes it.

Known issue: the search index does not yet support files larger than 25 MB. Larger files are refused with a message that names the limit.

Planned, not yet available: the source file name in the activity log. We expect to release it in the second quarter of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 43.8

Quillstone Ledger 43.8 was released on 22 November 2027 for customers on the hosted plan.

What changed:

- Exports: now keeps the original currency for 180 days. After that the value is removed, and the entry shows when it was removed.
- The activity log: now shows the attachment count beside each entry. Entries made before 43.8 show a dash, because the value was not recorded.
- Account settings: now accepts the exchange rate used as a filter. The filter applies to records created on or after 43.8.

- Fixed: in bank feeds, a month with no records showed a blank page. It now shows an empty table.

## Quillstone Ledger 43.9

Quillstone Ledger 43.9 was released on 1 December 2027 for customers on the hosted plan.

What changed:

- Currency conversion: now accepts the reviewer's comment count as a filter. The filter applies to records created on or after 43.9.
- The import screen: now accepts the approving user's role as a filter. The filter applies to records created on or after 43.9.
- The audit view: now keeps the payment terms for 90 days. After that the value is removed, and the entry shows when it was removed.

- Fixed: in report generation, a month with no records showed a blank page. It now shows an empty table.

Planned, not yet available: the device type in the search index. We expect to release it in the first quarter of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 44.0

Quillstone Ledger 44.0 was released on 8 December 2027 for customers on the hosted plan.

What changed:

- Report generation: now keeps the device type for 90 days. After that the value is removed, and the entry shows when it was removed.
- Invoice matching: faster on accounts with many records. In our own test on 8 December 2027, a 45,000-row run took 28 seconds on 43.9 and 14 seconds on 44.0. Your results depend on your data.
- Account settings: now shows the original currency beside each entry. Entries made before 44.0 show a dash, because the value was not recorded.
- Attachment storage: can now be turned off for each user. When it is turned off, the approving user's role is not collected for that user from then on.
- The approval queue: now shows the device type beside each entry. Entries made before 44.0 show a dash, because the value was not recorded.

- Fixed: a date typed as day/month could be read as month/day in report generation. The format now follows the account's settings.

## Quillstone Ledger 44.1

Quillstone Ledger 44.1 was released on 15 December 2027 for customers on the hosted plan.

What changed:

- The search index: can now be turned off for each user. When it is turned off, the time zone of the entry is not collected for that user from then on.
- Account settings: now accepts the batch number as a filter. The filter applies to records created on or after 44.1.
- The search index: now keeps the exchange rate used for 30 days. After that the value is removed, and the entry shows when it was removed.
- Report generation: now names the user who last changed the time zone of the entry. It does not show why the change was made, because the product does not ask for a reason.

- Fixed: a date typed as day/month could be read as month/day in attachment storage. The format now follows the account's settings.
- Fixed: a date typed as day/month could be read as month/day in bank feeds. The format now follows the account's settings.

## Quillstone Ledger 44.2

Quillstone Ledger 44.2 was released on 22 December 2027 for customers on the hosted plan.

What changed:

- Invoice matching: now includes a column for the import job number. The column is empty for records created before 44.2, because the import job number was not stored then.
- The audit view: now accepts the original currency as a filter. The filter applies to records created on or after 44.2.
- The search index: faster on accounts with many records. In our own test on 22 December 2027, a 94,000-row run took 37 seconds on 44.1 and 12 seconds on 44.2. Your results depend on your data.
- Report generation: now keeps the cost centre for 30 days. After that the value is removed, and the entry shows when it was removed.
- Currency conversion: now includes a column for the approving user's role. The column is empty for records created before 44.2, because the approving user's role was not stored then.

- Fixed: totals in the mobile view were rounded before they were added, which could change the last digit. They are now added first and rounded once.
- Fixed: in exports, a month with no records showed a blank page. It now shows an empty table.

Known issue: user roles does not yet support files larger than 25 MB. Larger files are refused with a message that names the limit.

Planned, not yet available: the time zone of the entry in the approval queue. We expect to release it in the second quarter of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 44.3

Quillstone Ledger 44.3 was released on 1 January 2028 for customers on the hosted plan.

What changed:

- The mobile view: now shows the reviewer's comment count beside each entry. Entries made before 44.3 show a dash, because the value was not recorded.
- Account settings: now accepts the payment terms as a filter. The filter applies to records created on or after 44.3.
- Bank feeds: now shows the source file name beside each entry. Entries made before 44.3 show a dash, because the value was not recorded.

- Fixed: in attachment storage, a month with no records showed a blank page. It now shows an empty table.
- Fixed: the search index did not keep the sort order after a page reload. It now keeps it until the user changes it.

Known issue: on slow connections the mobile view can take up to 30 seconds to load. We are working on it and will say here when it is fixed.

Planned, not yet available: the reviewer's comment count in report generation. We expect to release it in the second half of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 44.4

Quillstone Ledger 44.4 was released on 8 January 2028 for customers on the hosted plan.

What changed:

- The tax summary: now accepts the attachment count as a filter. The filter applies to records created on or after 44.4.
- The audit view: now includes a column for the approving user's role. The column is empty for records created before 44.4, because the approving user's role was not stored then.
- The activity log: can now be turned off for each user. When it is turned off, the original currency is not collected for that user from then on.

- Fixed: the audit view did not keep the sort order after a page reload. It now keeps it until the user changes it.
- Fixed: totals in bank feeds were rounded before they were added, which could change the last digit. They are now added first and rounded once.

## Quillstone Ledger 44.5

Quillstone Ledger 44.5 was released on 15 January 2028 for customers on the hosted plan.

What changed:

- The import screen: can now be turned off for each user. When it is turned off, the device type is not collected for that user from then on.
- Attachment storage: faster on accounts with many records. In our own test on 15 January 2028, a 163,000-row run took 45 seconds on 44.4 and 22 seconds on 44.5. Your results depend on your data.
- The approval queue: now keeps the device type for 365 days. After that the value is removed, and the entry shows when it was removed.
- The tax summary: now shows the payment terms beside each entry. Entries made before 44.5 show a dash, because the value was not recorded.

- Fixed: in scheduled reports, a month with no records showed a blank page. It now shows an empty table.
- Fixed: in attachment storage, a month with no records showed a blank page. It now shows an empty table.

Known issue: on slow connections attachment storage can take up to 10 seconds to load. We are working on it and will say here when it is fixed.

## Quillstone Ledger 44.6

Quillstone Ledger 44.6 was released on 22 January 2028 for customers on the hosted plan.

What changed:

- The tax summary: now names the user who last changed the payment terms. It does not show why the change was made, because the product does not ask for a reason.
- Currency conversion: now keeps the time zone of the entry for 365 days. After that the value is removed, and the entry shows when it was removed.
- Bank feeds: now shows the source file name beside each entry. Entries made before 44.6 show a dash, because the value was not recorded.
- User roles: now accepts the batch number as a filter. The filter applies to records created on or after 44.6.
- The activity log: now accepts the exchange rate used as a filter. The filter applies to records created on or after 44.6.

- Fixed: in bank feeds, a month with no records showed a blank page. It now shows an empty table.
- Fixed: totals in the search index were rounded before they were added, which could change the last digit. They are now added first and rounded once.

Known issue: report generation does not yet support files larger than 25 MB. Larger files are refused with a message that names the limit.

Planned, not yet available: the approving user's role in exports. We expect to release it in the second quarter of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 44.7

Quillstone Ledger 44.7 was released on 1 February 2028 for customers on the hosted plan.

What changed:

- Exports: now accepts the attachment count as a filter. The filter applies to records created on or after 44.7.
- The activity log: now includes a column for the approving user's role. The column is empty for records created before 44.7, because the approving user's role was not stored then.
- Scheduled reports: now includes a column for the exchange rate used. The column is empty for records created before 44.7, because the exchange rate used was not stored then.
- Attachment storage: can now be turned off for each user. When it is turned off, the source file name is not collected for that user from then on.

- Fixed: exports did not keep the sort order after a page reload. It now keeps it until the user changes it.

Known issue: on slow connections account settings can take up to 20 seconds to load. We are working on it and will say here when it is fixed.

## Quillstone Ledger 44.8

Quillstone Ledger 44.8 was released on 8 February 2028 for customers on the hosted plan.

What changed:

- Account settings: now shows the source file name beside each entry. Entries made before 44.8 show a dash, because the value was not recorded.
- Attachment storage: now shows the payment terms beside each entry. Entries made before 44.8 show a dash, because the value was not recorded.
- The audit view: now includes a column for the source file name. The column is empty for records created before 44.8, because the source file name was not stored then.

- Fixed: in scheduled reports, a month with no records showed a blank page. It now shows an empty table.
- Fixed: attachment storage did not keep the sort order after a page reload. It now keeps it until the user changes it.

Known issue: on slow connections report generation can take up to 20 seconds to load. We are working on it and will say here when it is fixed.

## Quillstone Ledger 44.9

Quillstone Ledger 44.9 was released on 15 February 2028 for customers on the hosted plan.

What changed:

- Report generation: now shows the batch number beside each entry. Entries made before 44.9 show a dash, because the value was not recorded.
- Account settings: now includes a column for the time zone of the entry. The column is empty for records created before 44.9, because the time zone of the entry was not stored then.
- Exports: now includes a column for the exchange rate used. The column is empty for records created before 44.9, because the exchange rate used was not stored then.
- Invoice matching: now keeps the exchange rate used for 180 days. After that the value is removed, and the entry shows when it was removed.
- The tax summary: now includes a column for the device type. The column is empty for records created before 44.9, because the device type was not stored then.

- Fixed: currency conversion did not keep the sort order after a page reload. It now keeps it until the user changes it.

Known issue: the audit view does not yet support files larger than 100 MB. Larger files are refused with a message that names the limit.

## Quillstone Ledger 45.0

Quillstone Ledger 45.0 was released on 22 February 2028 for customers on the hosted plan.

What changed:

- The tax summary: now names the user who last changed the device type. It does not show why the change was made, because the product does not ask for a reason.
- Account settings: now includes a column for the time zone of the entry. The column is empty for records created before 45.0, because the time zone of the entry was not stored then.
- Account settings: now includes a column for the reviewer's comment count. The column is empty for records created before 45.0, because the reviewer's comment count was not stored then.
- Currency conversion: now shows the cost centre beside each entry. Entries made before 45.0 show a dash, because the value was not recorded.
- User roles: now keeps the exchange rate used for 180 days. After that the value is removed, and the entry shows when it was removed.

- Fixed: totals in the activity log were rounded before they were added, which could change the last digit. They are now added first and rounded once.
- Fixed: totals in invoice matching were rounded before they were added, which could change the last digit. They are now added first and rounded once.

Known issue: currency conversion does not yet support files larger than 25 MB. Larger files are refused with a message that names the limit.

## Quillstone Ledger 45.1

Quillstone Ledger 45.1 was released on 1 March 2028 for customers on the hosted plan.

What changed:

- Account settings: can now be turned off for each user. When it is turned off, the exchange rate used is not collected for that user from then on.
- Invoice matching: faster on accounts with many records. In our own test on 1 March 2028, a 151,000-row run took 34 seconds on 45.0 and 6 seconds on 45.1. Your results depend on your data.
- Invoice matching: now keeps the reviewer's comment count for 30 days. After that the value is removed, and the entry shows when it was removed.
- Invoice matching: faster on accounts with many records. In our own test on 1 March 2028, a 94,000-row run took 31 seconds on 45.0 and 10 seconds on 45.1. Your results depend on your data.
- Report generation: now keeps the import job number for 90 days. After that the value is removed, and the entry shows when it was removed.

- Fixed: a date typed as day/month could be read as month/day in exports. The format now follows the account's settings.

## Quillstone Ledger 45.2

Quillstone Ledger 45.2 was released on 8 March 2028 for customers on the hosted plan.

What changed:

- The approval queue: now shows the exchange rate used beside each entry. Entries made before 45.2 show a dash, because the value was not recorded.
- Exports: now shows the attachment count beside each entry. Entries made before 45.2 show a dash, because the value was not recorded.
- Invoice matching: now shows the exchange rate used beside each entry. Entries made before 45.2 show a dash, because the value was not recorded.
- Report generation: now names the user who last changed the attachment count. It does not show why the change was made, because the product does not ask for a reason.

- Fixed: in the activity log, a month with no records showed a blank page. It now shows an empty table.

Planned, not yet available: the attachment count in the activity log. We expect to release it in the second half of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 45.3

Quillstone Ledger 45.3 was released on 15 March 2028 for customers on the hosted plan.

What changed:

- Scheduled reports: now accepts the time zone of the entry as a filter. The filter applies to records created on or after 45.3.
- Attachment storage: now keeps the cost centre for 365 days. After that the value is removed, and the entry shows when it was removed.
- Invoice matching: now keeps the cost centre for 180 days. After that the value is removed, and the entry shows when it was removed.
- Report generation: now includes a column for the batch number. The column is empty for records created before 45.3, because the batch number was not stored then.

- Fixed: totals in the mobile view were rounded before they were added, which could change the last digit. They are now added first and rounded once.

## Quillstone Ledger 45.4

Quillstone Ledger 45.4 was released on 22 March 2028 for customers on the hosted plan.

What changed:

- Currency conversion: now includes a column for the original currency. The column is empty for records created before 45.4, because the original currency was not stored then.
- The mobile view: can now be turned off for each user. When it is turned off, the import job number is not collected for that user from then on.
- Scheduled reports: can now be turned off for each user. When it is turned off, the attachment count is not collected for that user from then on.
- The search index: now keeps the device type for 365 days. After that the value is removed, and the entry shows when it was removed.

- Fixed: bank feeds did not keep the sort order after a page reload. It now keeps it until the user changes it.
- Fixed: a date typed as day/month could be read as month/day in invoice matching. The format now follows the account's settings.

## Quillstone Ledger 45.5

Quillstone Ledger 45.5 was released on 1 April 2028 for customers on the hosted plan.

What changed:

- Bank feeds: now shows the original currency beside each entry. Entries made before 45.5 show a dash, because the value was not recorded.
- Currency conversion: faster on accounts with many records. In our own test on 1 April 2028, a 40,000-row run took 65 seconds on 45.4 and 16 seconds on 45.5. Your results depend on your data.
- Scheduled reports: now keeps the batch number for 90 days. After that the value is removed, and the entry shows when it was removed.
- Report generation: now names the user who last changed the import job number. It does not show why the change was made, because the product does not ask for a reason.

- Fixed: in account settings, a month with no records showed a blank page. It now shows an empty table.

Known issue: the mobile view does not yet support files larger than 25 MB. Larger files are refused with a message that names the limit.

## Quillstone Ledger 45.6

Quillstone Ledger 45.6 was released on 8 April 2028 for customers on the hosted plan.

What changed:

- The import screen: now shows the attachment count beside each entry. Entries made before 45.6 show a dash, because the value was not recorded.
- The audit view: now shows the exchange rate used beside each entry. Entries made before 45.6 show a dash, because the value was not recorded.
- Invoice matching: now names the user who last changed the payment terms. It does not show why the change was made, because the product does not ask for a reason.
- Account settings: can now be turned off for each user. When it is turned off, the exchange rate used is not collected for that user from then on.

- Fixed: user roles did not keep the sort order after a page reload. It now keeps it until the user changes it.

Planned, not yet available: the original currency in the import screen. We expect to release it in the second half of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 45.7

Quillstone Ledger 45.7 was released on 15 April 2028 for customers on the hosted plan.

What changed:

- The tax summary: now accepts the reviewer's comment count as a filter. The filter applies to records created on or after 45.7.
- User roles: now keeps the attachment count for 30 days. After that the value is removed, and the entry shows when it was removed.
- Invoice matching: now includes a column for the batch number. The column is empty for records created before 45.7, because the batch number was not stored then.
- Attachment storage: now accepts the reviewer's comment count as a filter. The filter applies to records created on or after 45.7.

- Fixed: a date typed as day/month could be read as month/day in invoice matching. The format now follows the account's settings.
- Fixed: in scheduled reports, a month with no records showed a blank page. It now shows an empty table.

Known issue: the audit view does not yet support files larger than 25 MB. Larger files are refused with a message that names the limit.

## Quillstone Ledger 45.8

Quillstone Ledger 45.8 was released on 22 April 2028 for customers on the hosted plan.

What changed:

- The tax summary: now accepts the reviewer's comment count as a filter. The filter applies to records created on or after 45.8.
- The search index: now accepts the approving user's role as a filter. The filter applies to records created on or after 45.8.
- Currency conversion: faster on accounts with many records. In our own test on 22 April 2028, a 149,000-row run took 63 seconds on 45.7 and 15 seconds on 45.8. Your results depend on your data.
- Bank feeds: now shows the approving user's role beside each entry. Entries made before 45.8 show a dash, because the value was not recorded.

- Fixed: totals in bank feeds were rounded before they were added, which could change the last digit. They are now added first and rounded once.

Known issue: currency conversion does not yet support files larger than 100 MB. Larger files are refused with a message that names the limit.

## Quillstone Ledger 45.9

Quillstone Ledger 45.9 was released on 1 May 2028 for customers on the hosted plan.

What changed:

- The activity log: faster on accounts with many records. In our own test on 1 May 2028, a 194,000-row run took 55 seconds on 45.8 and 11 seconds on 45.9. Your results depend on your data.
- Currency conversion: faster on accounts with many records. In our own test on 1 May 2028, a 112,000-row run took 85 seconds on 45.8 and 28 seconds on 45.9. Your results depend on your data.
- Currency conversion: faster on accounts with many records. In our own test on 1 May 2028, a 59,000-row run took 25 seconds on 45.8 and 5 seconds on 45.9. Your results depend on your data.
- User roles: now includes a column for the reviewer's comment count. The column is empty for records created before 45.9, because the reviewer's comment count was not stored then.

- Fixed: report generation did not keep the sort order after a page reload. It now keeps it until the user changes it.
- Fixed: report generation did not keep the sort order after a page reload. It now keeps it until the user changes it.

Known issue: on slow connections the tax summary can take up to 30 seconds to load. We are working on it and will say here when it is fixed.

Planned, not yet available: the import job number in report generation. We expect to release it in the second quarter of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 46.0

Quillstone Ledger 46.0 was released on 8 May 2028 for customers on the hosted plan.

What changed:

- Bank feeds: faster on accounts with many records. In our own test on 8 May 2028, a 89,000-row run took 76 seconds on 45.9 and 19 seconds on 46.0. Your results depend on your data.
- Account settings: now names the user who last changed the reviewer's comment count. It does not show why the change was made, because the product does not ask for a reason.
- Scheduled reports: faster on accounts with many records. In our own test on 8 May 2028, a 142,000-row run took 69 seconds on 45.9 and 13 seconds on 46.0. Your results depend on your data.
- Bank feeds: now includes a column for the time zone of the entry. The column is empty for records created before 46.0, because the time zone of the entry was not stored then.
- Scheduled reports: now includes a column for the device type. The column is empty for records created before 46.0, because the device type was not stored then.

- Fixed: a date typed as day/month could be read as month/day in account settings. The format now follows the account's settings.
- Fixed: a date typed as day/month could be read as month/day in the audit view. The format now follows the account's settings.

Planned, not yet available: the time zone of the entry in scheduled reports. We expect to release it in the second half of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 46.1

Quillstone Ledger 46.1 was released on 15 May 2028 for customers on the hosted plan.

What changed:

- Currency conversion: faster on accounts with many records. In our own test on 15 May 2028, a 108,000-row run took 70 seconds on 46.0 and 14 seconds on 46.1. Your results depend on your data.
- Currency conversion: can now be turned off for each user. When it is turned off, the cost centre is not collected for that user from then on.
- Currency conversion: now keeps the cost centre for 30 days. After that the value is removed, and the entry shows when it was removed.
- The approval queue: now accepts the reviewer's comment count as a filter. The filter applies to records created on or after 46.1.

- Fixed: a date typed as day/month could be read as month/day in attachment storage. The format now follows the account's settings.
- Fixed: a date typed as day/month could be read as month/day in attachment storage. The format now follows the account's settings.

Known issue: the activity log does not yet support files larger than 25 MB. Larger files are refused with a message that names the limit.

## Quillstone Ledger 46.2

Quillstone Ledger 46.2 was released on 22 May 2028 for customers on the hosted plan.

What changed:

- Currency conversion: faster on accounts with many records. In our own test on 22 May 2028, a 33,000-row run took 52 seconds on 46.1 and 17 seconds on 46.2. Your results depend on your data.
- Bank feeds: now keeps the cost centre for 365 days. After that the value is removed, and the entry shows when it was removed.
- Report generation: can now be turned off for each user. When it is turned off, the exchange rate used is not collected for that user from then on.
- Account settings: now accepts the exchange rate used as a filter. The filter applies to records created on or after 46.2.
- The tax summary: now shows the batch number beside each entry. Entries made before 46.2 show a dash, because the value was not recorded.

- Fixed: a date typed as day/month could be read as month/day in invoice matching. The format now follows the account's settings.

## Quillstone Ledger 46.3

Quillstone Ledger 46.3 was released on 1 June 2028 for customers on the hosted plan.

What changed:

- Report generation: faster on accounts with many records. In our own test on 1 June 2028, a 145,000-row run took 53 seconds on 46.2 and 17 seconds on 46.3. Your results depend on your data.
- Report generation: faster on accounts with many records. In our own test on 1 June 2028, a 75,000-row run took 83 seconds on 46.2 and 16 seconds on 46.3. Your results depend on your data.
- Exports: now accepts the approving user's role as a filter. The filter applies to records created on or after 46.3.
- Report generation: faster on accounts with many records. In our own test on 1 June 2028, a 107,000-row run took 30 seconds on 46.2 and 7 seconds on 46.3. Your results depend on your data.
- Invoice matching: now names the user who last changed the import job number. It does not show why the change was made, because the product does not ask for a reason.

- Fixed: totals in invoice matching were rounded before they were added, which could change the last digit. They are now added first and rounded once.

Planned, not yet available: the source file name in report generation. We expect to release it in the first quarter of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 46.4

Quillstone Ledger 46.4 was released on 8 June 2028 for customers on the hosted plan.

What changed:

- The tax summary: now shows the import job number beside each entry. Entries made before 46.4 show a dash, because the value was not recorded.
- Attachment storage: faster on accounts with many records. In our own test on 8 June 2028, a 91,000-row run took 38 seconds on 46.3 and 19 seconds on 46.4. Your results depend on your data.
- The mobile view: now names the user who last changed the original currency. It does not show why the change was made, because the product does not ask for a reason.
- Report generation: can now be turned off for each user. When it is turned off, the payment terms is not collected for that user from then on.

- Fixed: the import screen did not keep the sort order after a page reload. It now keeps it until the user changes it.
- Fixed: in scheduled reports, a month with no records showed a blank page. It now shows an empty table.

Known issue: on slow connections bank feeds can take up to 20 seconds to load. We are working on it and will say here when it is fixed.

## Quillstone Ledger 46.5

Quillstone Ledger 46.5 was released on 15 June 2028 for customers on the hosted plan.

What changed:

- Invoice matching: now shows the import job number beside each entry. Entries made before 46.5 show a dash, because the value was not recorded.
- The import screen: now keeps the cost centre for 180 days. After that the value is removed, and the entry shows when it was removed.
- Report generation: now shows the payment terms beside each entry. Entries made before 46.5 show a dash, because the value was not recorded.

- Fixed: attachment storage did not keep the sort order after a page reload. It now keeps it until the user changes it.
- Fixed: a date typed as day/month could be read as month/day in the approval queue. The format now follows the account's settings.

## Quillstone Ledger 46.6

Quillstone Ledger 46.6 was released on 22 June 2028 for customers on the hosted plan.

What changed:

- The import screen: can now be turned off for each user. When it is turned off, the exchange rate used is not collected for that user from then on.
- The audit view: now shows the attachment count beside each entry. Entries made before 46.6 show a dash, because the value was not recorded.
- Invoice matching: now names the user who last changed the time zone of the entry. It does not show why the change was made, because the product does not ask for a reason.
- Currency conversion: now includes a column for the device type. The column is empty for records created before 46.6, because the device type was not stored then.

- Fixed: a date typed as day/month could be read as month/day in the approval queue. The format now follows the account's settings.

## Quillstone Ledger 46.7

Quillstone Ledger 46.7 was released on 1 July 2028 for customers on the hosted plan.

What changed:

- Invoice matching: now accepts the approving user's role as a filter. The filter applies to records created on or after 46.7.
- The import screen: now names the user who last changed the exchange rate used. It does not show why the change was made, because the product does not ask for a reason.
- Attachment storage: now includes a column for the cost centre. The column is empty for records created before 46.7, because the cost centre was not stored then.
- User roles: now names the user who last changed the batch number. It does not show why the change was made, because the product does not ask for a reason.

- Fixed: totals in the mobile view were rounded before they were added, which could change the last digit. They are now added first and rounded once.

Known issue: exports does not yet support files larger than 50 MB. Larger files are refused with a message that names the limit.

Planned, not yet available: the batch number in currency conversion. We expect to release it in the second half of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 46.8

Quillstone Ledger 46.8 was released on 8 July 2028 for customers on the hosted plan.

What changed:

- The mobile view: now accepts the exchange rate used as a filter. The filter applies to records created on or after 46.8.
- Exports: now names the user who last changed the batch number. It does not show why the change was made, because the product does not ask for a reason.
- The audit view: can now be turned off for each user. When it is turned off, the approving user's role is not collected for that user from then on.
- The search index: now names the user who last changed the device type. It does not show why the change was made, because the product does not ask for a reason.

- Fixed: totals in invoice matching were rounded before they were added, which could change the last digit. They are now added first and rounded once.
- Fixed: in bank feeds, a month with no records showed a blank page. It now shows an empty table.

Known issue: report generation does not yet support files larger than 50 MB. Larger files are refused with a message that names the limit.

## Quillstone Ledger 46.9

Quillstone Ledger 46.9 was released on 15 July 2028 for customers on the hosted plan.

What changed:

- The search index: can now be turned off for each user. When it is turned off, the reviewer's comment count is not collected for that user from then on.
- The approval queue: can now be turned off for each user. When it is turned off, the import job number is not collected for that user from then on.
- The tax summary: now includes a column for the time zone of the entry. The column is empty for records created before 46.9, because the time zone of the entry was not stored then.
- The import screen: now shows the reviewer's comment count beside each entry. Entries made before 46.9 show a dash, because the value was not recorded.
- Report generation: now shows the device type beside each entry. Entries made before 46.9 show a dash, because the value was not recorded.

- Fixed: exports did not keep the sort order after a page reload. It now keeps it until the user changes it.
- Fixed: the import screen did not keep the sort order after a page reload. It now keeps it until the user changes it.

## Quillstone Ledger 47.0

Quillstone Ledger 47.0 was released on 22 July 2028 for customers on the hosted plan.

What changed:

- The activity log: now keeps the batch number for 30 days. After that the value is removed, and the entry shows when it was removed.
- The approval queue: now shows the cost centre beside each entry. Entries made before 47.0 show a dash, because the value was not recorded.
- Exports: faster on accounts with many records. In our own test on 22 July 2028, a 10,000-row run took 50 seconds on 46.9 and 16 seconds on 47.0. Your results depend on your data.
- Account settings: now keeps the import job number for 90 days. After that the value is removed, and the entry shows when it was removed.
- The search index: now includes a column for the exchange rate used. The column is empty for records created before 47.0, because the exchange rate used was not stored then.

- Fixed: in the tax summary, a month with no records showed a blank page. It now shows an empty table.

Planned, not yet available: the source file name in the mobile view. We expect to release it in the second quarter of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 47.1

Quillstone Ledger 47.1 was released on 1 August 2028 for customers on the hosted plan.

What changed:

- The tax summary: now shows the approving user's role beside each entry. Entries made before 47.1 show a dash, because the value was not recorded.
- Invoice matching: now shows the payment terms beside each entry. Entries made before 47.1 show a dash, because the value was not recorded.
- The search index: now includes a column for the cost centre. The column is empty for records created before 47.1, because the cost centre was not stored then.
- Scheduled reports: now includes a column for the time zone of the entry. The column is empty for records created before 47.1, because the time zone of the entry was not stored then.
- Invoice matching: can now be turned off for each user. When it is turned off, the approving user's role is not collected for that user from then on.

- Fixed: a date typed as day/month could be read as month/day in bank feeds. The format now follows the account's settings.

Known issue: on slow connections report generation can take up to 10 seconds to load. We are working on it and will say here when it is fixed.

## Quillstone Ledger 47.2

Quillstone Ledger 47.2 was released on 8 August 2028 for customers on the hosted plan.

What changed:

- Scheduled reports: now keeps the source file name for 90 days. After that the value is removed, and the entry shows when it was removed.
- Attachment storage: now names the user who last changed the cost centre. It does not show why the change was made, because the product does not ask for a reason.
- The tax summary: now names the user who last changed the source file name. It does not show why the change was made, because the product does not ask for a reason.
- The search index: can now be turned off for each user. When it is turned off, the import job number is not collected for that user from then on.
- The approval queue: now accepts the cost centre as a filter. The filter applies to records created on or after 47.2.

- Fixed: a date typed as day/month could be read as month/day in attachment storage. The format now follows the account's settings.
- Fixed: totals in the mobile view were rounded before they were added, which could change the last digit. They are now added first and rounded once.

## Quillstone Ledger 47.3

Quillstone Ledger 47.3 was released on 15 August 2028 for customers on the hosted plan.

What changed:

- The audit view: now names the user who last changed the batch number. It does not show why the change was made, because the product does not ask for a reason.
- Bank feeds: can now be turned off for each user. When it is turned off, the import job number is not collected for that user from then on.
- User roles: now keeps the attachment count for 180 days. After that the value is removed, and the entry shows when it was removed.
- The tax summary: faster on accounts with many records. In our own test on 15 August 2028, a 23,000-row run took 43 seconds on 47.2 and 8 seconds on 47.3. Your results depend on your data.

- Fixed: a date typed as day/month could be read as month/day in bank feeds. The format now follows the account's settings.
- Fixed: totals in scheduled reports were rounded before they were added, which could change the last digit. They are now added first and rounded once.

## Quillstone Ledger 47.4

Quillstone Ledger 47.4 was released on 22 August 2028 for customers on the hosted plan.

What changed:

- The tax summary: now includes a column for the attachment count. The column is empty for records created before 47.4, because the attachment count was not stored then.
- Report generation: now names the user who last changed the original currency. It does not show why the change was made, because the product does not ask for a reason.
- Exports: now accepts the import job number as a filter. The filter applies to records created on or after 47.4.
- The search index: now names the user who last changed the source file name. It does not show why the change was made, because the product does not ask for a reason.
- Invoice matching: now shows the import job number beside each entry. Entries made before 47.4 show a dash, because the value was not recorded.

- Fixed: user roles did not keep the sort order after a page reload. It now keeps it until the user changes it.

Planned, not yet available: the original currency in the audit view. We expect to release it in the second half of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 47.5

Quillstone Ledger 47.5 was released on 1 September 2028 for customers on the hosted plan.

What changed:

- Attachment storage: now keeps the exchange rate used for 30 days. After that the value is removed, and the entry shows when it was removed.
- Exports: can now be turned off for each user. When it is turned off, the approving user's role is not collected for that user from then on.
- Scheduled reports: now includes a column for the approving user's role. The column is empty for records created before 47.5, because the approving user's role was not stored then.
- Invoice matching: now names the user who last changed the cost centre. It does not show why the change was made, because the product does not ask for a reason.
- Bank feeds: now keeps the attachment count for 365 days. After that the value is removed, and the entry shows when it was removed.

- Fixed: in bank feeds, a month with no records showed a blank page. It now shows an empty table.

Planned, not yet available: the original currency in account settings. We expect to release it in the first quarter of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 47.6

Quillstone Ledger 47.6 was released on 8 September 2028 for customers on the hosted plan.

What changed:

- The import screen: now accepts the reviewer's comment count as a filter. The filter applies to records created on or after 47.6.
- Bank feeds: now accepts the import job number as a filter. The filter applies to records created on or after 47.6.
- Report generation: now accepts the import job number as a filter. The filter applies to records created on or after 47.6.

- Fixed: in the approval queue, a month with no records showed a blank page. It now shows an empty table.

## Quillstone Ledger 47.7

Quillstone Ledger 47.7 was released on 15 September 2028 for customers on the hosted plan.

What changed:

- The audit view: now names the user who last changed the reviewer's comment count. It does not show why the change was made, because the product does not ask for a reason.
- Invoice matching: faster on accounts with many records. In our own test on 15 September 2028, a 68,000-row run took 34 seconds on 47.6 and 17 seconds on 47.7. Your results depend on your data.
- Invoice matching: now shows the attachment count beside each entry. Entries made before 47.7 show a dash, because the value was not recorded.
- Report generation: now names the user who last changed the time zone of the entry. It does not show why the change was made, because the product does not ask for a reason.
- Invoice matching: now keeps the attachment count for 365 days. After that the value is removed, and the entry shows when it was removed.

- Fixed: in currency conversion, a month with no records showed a blank page. It now shows an empty table.
- Fixed: bank feeds did not keep the sort order after a page reload. It now keeps it until the user changes it.

Planned, not yet available: the original currency in the audit view. We expect to release it in the second half of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 47.8

Quillstone Ledger 47.8 was released on 22 September 2028 for customers on the hosted plan.

What changed:

- The activity log: can now be turned off for each user. When it is turned off, the cost centre is not collected for that user from then on.
- Exports: can now be turned off for each user. When it is turned off, the approving user's role is not collected for that user from then on.
- The mobile view: now names the user who last changed the original currency. It does not show why the change was made, because the product does not ask for a reason.
- The tax summary: now keeps the payment terms for 90 days. After that the value is removed, and the entry shows when it was removed.

- Fixed: totals in user roles were rounded before they were added, which could change the last digit. They are now added first and rounded once.
- Fixed: in currency conversion, a month with no records showed a blank page. It now shows an empty table.

Known issue: the activity log does not yet support files larger than 100 MB. Larger files are refused with a message that names the limit.

Planned, not yet available: the reviewer's comment count in the approval queue. We expect to release it in the second half of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 47.9

Quillstone Ledger 47.9 was released on 1 October 2028 for customers on the hosted plan.

What changed:

- Account settings: now includes a column for the reviewer's comment count. The column is empty for records created before 47.9, because the reviewer's comment count was not stored then.
- The search index: now includes a column for the batch number. The column is empty for records created before 47.9, because the batch number was not stored then.
- The activity log: now accepts the device type as a filter. The filter applies to records created on or after 47.9.
- The activity log: now includes a column for the cost centre. The column is empty for records created before 47.9, because the cost centre was not stored then.

- Fixed: in invoice matching, a month with no records showed a blank page. It now shows an empty table.
- Fixed: a date typed as day/month could be read as month/day in account settings. The format now follows the account's settings.

Known issue: on slow connections the mobile view can take up to 20 seconds to load. We are working on it and will say here when it is fixed.

## Quillstone Ledger 48.0

Quillstone Ledger 48.0 was released on 8 October 2028 for customers on the hosted plan.

What changed:

- The audit view: now keeps the reviewer's comment count for 180 days. After that the value is removed, and the entry shows when it was removed.
- Report generation: can now be turned off for each user. When it is turned off, the approving user's role is not collected for that user from then on.
- Bank feeds: can now be turned off for each user. When it is turned off, the attachment count is not collected for that user from then on.

- Fixed: currency conversion did not keep the sort order after a page reload. It now keeps it until the user changes it.

## Quillstone Ledger 48.1

Quillstone Ledger 48.1 was released on 15 October 2028 for customers on the hosted plan.

What changed:

- The import screen: now shows the batch number beside each entry. Entries made before 48.1 show a dash, because the value was not recorded.
- The tax summary: now includes a column for the source file name. The column is empty for records created before 48.1, because the source file name was not stored then.
- The activity log: now accepts the time zone of the entry as a filter. The filter applies to records created on or after 48.1.

- Fixed: currency conversion did not keep the sort order after a page reload. It now keeps it until the user changes it.

Planned, not yet available: the import job number in the audit view. We expect to release it in the second half of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 48.2

Quillstone Ledger 48.2 was released on 22 October 2028 for customers on the hosted plan.

What changed:

- Exports: now names the user who last changed the original currency. It does not show why the change was made, because the product does not ask for a reason.
- The tax summary: now shows the attachment count beside each entry. Entries made before 48.2 show a dash, because the value was not recorded.
- User roles: now includes a column for the batch number. The column is empty for records created before 48.2, because the batch number was not stored then.
- The tax summary: now shows the source file name beside each entry. Entries made before 48.2 show a dash, because the value was not recorded.

- Fixed: bank feeds did not keep the sort order after a page reload. It now keeps it until the user changes it.

## Quillstone Ledger 48.3

Quillstone Ledger 48.3 was released on 1 November 2028 for customers on the hosted plan.

What changed:

- The search index: now accepts the attachment count as a filter. The filter applies to records created on or after 48.3.
- Exports: now shows the source file name beside each entry. Entries made before 48.3 show a dash, because the value was not recorded.
- Report generation: now names the user who last changed the original currency. It does not show why the change was made, because the product does not ask for a reason.
- Currency conversion: now includes a column for the original currency. The column is empty for records created before 48.3, because the original currency was not stored then.

- Fixed: attachment storage did not keep the sort order after a page reload. It now keeps it until the user changes it.

Known issue: on slow connections the tax summary can take up to 20 seconds to load. We are working on it and will say here when it is fixed.

## Quillstone Ledger 48.4

Quillstone Ledger 48.4 was released on 8 November 2028 for customers on the hosted plan.

What changed:

- The audit view: now includes a column for the reviewer's comment count. The column is empty for records created before 48.4, because the reviewer's comment count was not stored then.
- Account settings: now shows the import job number beside each entry. Entries made before 48.4 show a dash, because the value was not recorded.
- The activity log: can now be turned off for each user. When it is turned off, the original currency is not collected for that user from then on.
- Scheduled reports: now shows the batch number beside each entry. Entries made before 48.4 show a dash, because the value was not recorded.
- The mobile view: can now be turned off for each user. When it is turned off, the device type is not collected for that user from then on.

- Fixed: in exports, a month with no records showed a blank page. It now shows an empty table.
- Fixed: the approval queue did not keep the sort order after a page reload. It now keeps it until the user changes it.

Known issue: on slow connections user roles can take up to 30 seconds to load. We are working on it and will say here when it is fixed.

## Quillstone Ledger 48.5

Quillstone Ledger 48.5 was released on 15 November 2028 for customers on the hosted plan.

What changed:

- Exports: now includes a column for the device type. The column is empty for records created before 48.5, because the device type was not stored then.
- Bank feeds: now includes a column for the attachment count. The column is empty for records created before 48.5, because the attachment count was not stored then.
- The audit view: faster on accounts with many records. In our own test on 15 November 2028, a 65,000-row run took 58 seconds on 48.4 and 29 seconds on 48.5. Your results depend on your data.
- Scheduled reports: now shows the original currency beside each entry. Entries made before 48.5 show a dash, because the value was not recorded.
- Account settings: now names the user who last changed the device type. It does not show why the change was made, because the product does not ask for a reason.

- Fixed: in scheduled reports, a month with no records showed a blank page. It now shows an empty table.
- Fixed: the audit view did not keep the sort order after a page reload. It now keeps it until the user changes it.

## Quillstone Ledger 48.6

Quillstone Ledger 48.6 was released on 22 November 2028 for customers on the hosted plan.

What changed:

- Attachment storage: now names the user who last changed the source file name. It does not show why the change was made, because the product does not ask for a reason.
- Account settings: now names the user who last changed the time zone of the entry. It does not show why the change was made, because the product does not ask for a reason.
- The import screen: now keeps the attachment count for 365 days. After that the value is removed, and the entry shows when it was removed.
- Account settings: now accepts the batch number as a filter. The filter applies to records created on or after 48.6.
- User roles: faster on accounts with many records. In our own test on 22 November 2028, a 177,000-row run took 30 seconds on 48.5 and 7 seconds on 48.6. Your results depend on your data.

- Fixed: in report generation, a month with no records showed a blank page. It now shows an empty table.

## Quillstone Ledger 48.7

Quillstone Ledger 48.7 was released on 1 December 2028 for customers on the hosted plan.

What changed:

- The audit view: now includes a column for the import job number. The column is empty for records created before 48.7, because the import job number was not stored then.
- The search index: can now be turned off for each user. When it is turned off, the time zone of the entry is not collected for that user from then on.
- The audit view: now names the user who last changed the reviewer's comment count. It does not show why the change was made, because the product does not ask for a reason.
- Scheduled reports: can now be turned off for each user. When it is turned off, the device type is not collected for that user from then on.
- The tax summary: now accepts the batch number as a filter. The filter applies to records created on or after 48.7.

- Fixed: invoice matching did not keep the sort order after a page reload. It now keeps it until the user changes it.

Known issue: on slow connections the import screen can take up to 20 seconds to load. We are working on it and will say here when it is fixed.

## Quillstone Ledger 48.8

Quillstone Ledger 48.8 was released on 8 December 2028 for customers on the hosted plan.

What changed:

- The tax summary: now includes a column for the cost centre. The column is empty for records created before 48.8, because the cost centre was not stored then.
- The audit view: can now be turned off for each user. When it is turned off, the source file name is not collected for that user from then on.
- Account settings: now keeps the original currency for 90 days. After that the value is removed, and the entry shows when it was removed.
- Exports: now includes a column for the original currency. The column is empty for records created before 48.8, because the original currency was not stored then.
- Currency conversion: now accepts the device type as a filter. The filter applies to records created on or after 48.8.

- Fixed: attachment storage did not keep the sort order after a page reload. It now keeps it until the user changes it.

## Quillstone Ledger 48.9

Quillstone Ledger 48.9 was released on 15 December 2028 for customers on the hosted plan.

What changed:

- User roles: can now be turned off for each user. When it is turned off, the payment terms is not collected for that user from then on.
- The tax summary: faster on accounts with many records. In our own test on 15 December 2028, a 198,000-row run took 73 seconds on 48.8 and 18 seconds on 48.9. Your results depend on your data.
- Report generation: now includes a column for the import job number. The column is empty for records created before 48.9, because the import job number was not stored then.

- Fixed: a date typed as day/month could be read as month/day in the mobile view. The format now follows the account's settings.

Known issue: account settings does not yet support files larger than 50 MB. Larger files are refused with a message that names the limit.

Planned, not yet available: the source file name in scheduled reports. We expect to release it in the second quarter of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 49.0

Quillstone Ledger 49.0 was released on 22 December 2028 for customers on the hosted plan.

What changed:

- Bank feeds: can now be turned off for each user. When it is turned off, the source file name is not collected for that user from then on.
- Account settings: now names the user who last changed the time zone of the entry. It does not show why the change was made, because the product does not ask for a reason.
- The search index: can now be turned off for each user. When it is turned off, the source file name is not collected for that user from then on.

- Fixed: invoice matching did not keep the sort order after a page reload. It now keeps it until the user changes it.

Known issue: attachment storage does not yet support files larger than 100 MB. Larger files are refused with a message that names the limit.

## Quillstone Ledger 49.1

Quillstone Ledger 49.1 was released on 1 January 2029 for customers on the hosted plan.

What changed:

- Bank feeds: now keeps the original currency for 30 days. After that the value is removed, and the entry shows when it was removed.
- The tax summary: faster on accounts with many records. In our own test on 1 January 2029, a 53,000-row run took 65 seconds on 49.0 and 13 seconds on 49.1. Your results depend on your data.
- Report generation: can now be turned off for each user. When it is turned off, the original currency is not collected for that user from then on.
- Bank feeds: faster on accounts with many records. In our own test on 1 January 2029, a 166,000-row run took 30 seconds on 49.0 and 7 seconds on 49.1. Your results depend on your data.
- The search index: now names the user who last changed the approving user's role. It does not show why the change was made, because the product does not ask for a reason.

- Fixed: totals in bank feeds were rounded before they were added, which could change the last digit. They are now added first and rounded once.
- Fixed: totals in user roles were rounded before they were added, which could change the last digit. They are now added first and rounded once.

Known issue: on slow connections scheduled reports can take up to 20 seconds to load. We are working on it and will say here when it is fixed.

## Quillstone Ledger 49.2

Quillstone Ledger 49.2 was released on 8 January 2029 for customers on the hosted plan.

What changed:

- Invoice matching: faster on accounts with many records. In our own test on 8 January 2029, a 14,000-row run took 65 seconds on 49.1 and 16 seconds on 49.2. Your results depend on your data.
- Invoice matching: now names the user who last changed the time zone of the entry. It does not show why the change was made, because the product does not ask for a reason.
- Exports: faster on accounts with many records. In our own test on 8 January 2029, a 94,000-row run took 46 seconds on 49.1 and 11 seconds on 49.2. Your results depend on your data.
- The mobile view: faster on accounts with many records. In our own test on 8 January 2029, a 130,000-row run took 58 seconds on 49.1 and 19 seconds on 49.2. Your results depend on your data.
- The audit view: now includes a column for the exchange rate used. The column is empty for records created before 49.2, because the exchange rate used was not stored then.

- Fixed: totals in invoice matching were rounded before they were added, which could change the last digit. They are now added first and rounded once.
- Fixed: in report generation, a month with no records showed a blank page. It now shows an empty table.

Known issue: on slow connections the mobile view can take up to 30 seconds to load. We are working on it and will say here when it is fixed.

## Quillstone Ledger 49.3

Quillstone Ledger 49.3 was released on 15 January 2029 for customers on the hosted plan.

What changed:

- The tax summary: faster on accounts with many records. In our own test on 15 January 2029, a 192,000-row run took 71 seconds on 49.2 and 23 seconds on 49.3. Your results depend on your data.
- The approval queue: now accepts the reviewer's comment count as a filter. The filter applies to records created on or after 49.3.
- Exports: now keeps the import job number for 90 days. After that the value is removed, and the entry shows when it was removed.
- The import screen: now includes a column for the payment terms. The column is empty for records created before 49.3, because the payment terms was not stored then.

- Fixed: a date typed as day/month could be read as month/day in the import screen. The format now follows the account's settings.

## Quillstone Ledger 49.4

Quillstone Ledger 49.4 was released on 22 January 2029 for customers on the hosted plan.

What changed:

- The activity log: now includes a column for the approving user's role. The column is empty for records created before 49.4, because the approving user's role was not stored then.
- The audit view: now names the user who last changed the device type. It does not show why the change was made, because the product does not ask for a reason.
- The audit view: can now be turned off for each user. When it is turned off, the source file name is not collected for that user from then on.
- Currency conversion: now names the user who last changed the exchange rate used. It does not show why the change was made, because the product does not ask for a reason.

- Fixed: in invoice matching, a month with no records showed a blank page. It now shows an empty table.

Planned, not yet available: the original currency in account settings. We expect to release it in the second quarter of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 49.5

Quillstone Ledger 49.5 was released on 1 February 2029 for customers on the hosted plan.

What changed:

- Invoice matching: now names the user who last changed the time zone of the entry. It does not show why the change was made, because the product does not ask for a reason.
- The search index: now shows the exchange rate used beside each entry. Entries made before 49.5 show a dash, because the value was not recorded.
- Report generation: now accepts the attachment count as a filter. The filter applies to records created on or after 49.5.
- The approval queue: now accepts the approving user's role as a filter. The filter applies to records created on or after 49.5.

- Fixed: account settings did not keep the sort order after a page reload. It now keeps it until the user changes it.

Known issue: exports does not yet support files larger than 50 MB. Larger files are refused with a message that names the limit.

Planned, not yet available: the source file name in the mobile view. We expect to release it in the second half of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 49.6

Quillstone Ledger 49.6 was released on 8 February 2029 for customers on the hosted plan.

What changed:

- Currency conversion: faster on accounts with many records. In our own test on 8 February 2029, a 11,000-row run took 34 seconds on 49.5 and 11 seconds on 49.6. Your results depend on your data.
- The search index: faster on accounts with many records. In our own test on 8 February 2029, a 17,000-row run took 69 seconds on 49.5 and 23 seconds on 49.6. Your results depend on your data.
- User roles: now keeps the device type for 30 days. After that the value is removed, and the entry shows when it was removed.
- The audit view: now names the user who last changed the batch number. It does not show why the change was made, because the product does not ask for a reason.
- User roles: now names the user who last changed the payment terms. It does not show why the change was made, because the product does not ask for a reason.

- Fixed: scheduled reports did not keep the sort order after a page reload. It now keeps it until the user changes it.

Known issue: the tax summary does not yet support files larger than 25 MB. Larger files are refused with a message that names the limit.

Planned, not yet available: the reviewer's comment count in the approval queue. We expect to release it in the first quarter of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 49.7

Quillstone Ledger 49.7 was released on 15 February 2029 for customers on the hosted plan.

What changed:

- User roles: can now be turned off for each user. When it is turned off, the batch number is not collected for that user from then on.
- The tax summary: now shows the import job number beside each entry. Entries made before 49.7 show a dash, because the value was not recorded.
- The search index: now shows the approving user's role beside each entry. Entries made before 49.7 show a dash, because the value was not recorded.
- The mobile view: now keeps the time zone of the entry for 365 days. After that the value is removed, and the entry shows when it was removed.

- Fixed: in account settings, a month with no records showed a blank page. It now shows an empty table.
- Fixed: invoice matching did not keep the sort order after a page reload. It now keeps it until the user changes it.

## Quillstone Ledger 49.8

Quillstone Ledger 49.8 was released on 22 February 2029 for customers on the hosted plan.

What changed:

- Attachment storage: now includes a column for the reviewer's comment count. The column is empty for records created before 49.8, because the reviewer's comment count was not stored then.
- User roles: now accepts the source file name as a filter. The filter applies to records created on or after 49.8.
- Account settings: faster on accounts with many records. In our own test on 22 February 2029, a 102,000-row run took 84 seconds on 49.7 and 21 seconds on 49.8. Your results depend on your data.
- Bank feeds: now names the user who last changed the batch number. It does not show why the change was made, because the product does not ask for a reason.
- Scheduled reports: now keeps the source file name for 365 days. After that the value is removed, and the entry shows when it was removed.

- Fixed: a date typed as day/month could be read as month/day in report generation. The format now follows the account's settings.
- Fixed: a date typed as day/month could be read as month/day in the search index. The format now follows the account's settings.

## Quillstone Ledger 49.9

Quillstone Ledger 49.9 was released on 1 March 2029 for customers on the hosted plan.

What changed:

- The search index: now names the user who last changed the time zone of the entry. It does not show why the change was made, because the product does not ask for a reason.
- Exports: now accepts the approving user's role as a filter. The filter applies to records created on or after 49.9.
- Currency conversion: now keeps the approving user's role for 365 days. After that the value is removed, and the entry shows when it was removed.
- Attachment storage: now accepts the approving user's role as a filter. The filter applies to records created on or after 49.9.
- Report generation: now accepts the device type as a filter. The filter applies to records created on or after 49.9.

- Fixed: in scheduled reports, a month with no records showed a blank page. It now shows an empty table.

Planned, not yet available: the source file name in the activity log. We expect to release it in the second quarter of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 50.0

Quillstone Ledger 50.0 was released on 8 March 2029 for customers on the hosted plan.

What changed:

- Account settings: now shows the reviewer's comment count beside each entry. Entries made before 50.0 show a dash, because the value was not recorded.
- User roles: now accepts the reviewer's comment count as a filter. The filter applies to records created on or after 50.0.
- The audit view: now keeps the batch number for 180 days. After that the value is removed, and the entry shows when it was removed.

- Fixed: scheduled reports did not keep the sort order after a page reload. It now keeps it until the user changes it.
- Fixed: in user roles, a month with no records showed a blank page. It now shows an empty table.

Known issue: on slow connections the activity log can take up to 10 seconds to load. We are working on it and will say here when it is fixed.

Planned, not yet available: the reviewer's comment count in the search index. We expect to release it in the first quarter of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 50.1

Quillstone Ledger 50.1 was released on 15 March 2029 for customers on the hosted plan.

What changed:

- The mobile view: now shows the batch number beside each entry. Entries made before 50.1 show a dash, because the value was not recorded.
- Account settings: can now be turned off for each user. When it is turned off, the exchange rate used is not collected for that user from then on.
- Scheduled reports: faster on accounts with many records. In our own test on 15 March 2029, a 156,000-row run took 79 seconds on 50.0 and 19 seconds on 50.1. Your results depend on your data.

- Fixed: the activity log did not keep the sort order after a page reload. It now keeps it until the user changes it.
- Fixed: in user roles, a month with no records showed a blank page. It now shows an empty table.

Known issue: on slow connections attachment storage can take up to 20 seconds to load. We are working on it and will say here when it is fixed.

## Quillstone Ledger 50.2

Quillstone Ledger 50.2 was released on 22 March 2029 for customers on the hosted plan.

What changed:

- The search index: now names the user who last changed the device type. It does not show why the change was made, because the product does not ask for a reason.
- User roles: now accepts the exchange rate used as a filter. The filter applies to records created on or after 50.2.
- User roles: now keeps the cost centre for 30 days. After that the value is removed, and the entry shows when it was removed.

- Fixed: currency conversion did not keep the sort order after a page reload. It now keeps it until the user changes it.
- Fixed: a date typed as day/month could be read as month/day in the audit view. The format now follows the account's settings.

## Quillstone Ledger 50.3

Quillstone Ledger 50.3 was released on 1 April 2029 for customers on the hosted plan.

What changed:

- User roles: now keeps the device type for 30 days. After that the value is removed, and the entry shows when it was removed.
- The approval queue: can now be turned off for each user. When it is turned off, the reviewer's comment count is not collected for that user from then on.
- Exports: now includes a column for the exchange rate used. The column is empty for records created before 50.3, because the exchange rate used was not stored then.
- Exports: now includes a column for the import job number. The column is empty for records created before 50.3, because the import job number was not stored then.

- Fixed: in the tax summary, a month with no records showed a blank page. It now shows an empty table.
- Fixed: totals in the audit view were rounded before they were added, which could change the last digit. They are now added first and rounded once.

Known issue: on slow connections the tax summary can take up to 20 seconds to load. We are working on it and will say here when it is fixed.

Planned, not yet available: the source file name in the audit view. We expect to release it in the first quarter of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 50.4

Quillstone Ledger 50.4 was released on 8 April 2029 for customers on the hosted plan.

What changed:

- The search index: now shows the reviewer's comment count beside each entry. Entries made before 50.4 show a dash, because the value was not recorded.
- Invoice matching: now names the user who last changed the attachment count. It does not show why the change was made, because the product does not ask for a reason.
- User roles: now accepts the import job number as a filter. The filter applies to records created on or after 50.4.

- Fixed: totals in report generation were rounded before they were added, which could change the last digit. They are now added first and rounded once.
- Fixed: a date typed as day/month could be read as month/day in the audit view. The format now follows the account's settings.

## Quillstone Ledger 50.5

Quillstone Ledger 50.5 was released on 15 April 2029 for customers on the hosted plan.

What changed:

- Bank feeds: now keeps the approving user's role for 30 days. After that the value is removed, and the entry shows when it was removed.
- The search index: now includes a column for the import job number. The column is empty for records created before 50.5, because the import job number was not stored then.
- Exports: now accepts the original currency as a filter. The filter applies to records created on or after 50.5.
- Attachment storage: faster on accounts with many records. In our own test on 15 April 2029, a 121,000-row run took 31 seconds on 50.4 and 10 seconds on 50.5. Your results depend on your data.
- Scheduled reports: now includes a column for the original currency. The column is empty for records created before 50.5, because the original currency was not stored then.

- Fixed: scheduled reports did not keep the sort order after a page reload. It now keeps it until the user changes it.

Known issue: account settings does not yet support files larger than 25 MB. Larger files are refused with a message that names the limit.

Planned, not yet available: the batch number in user roles. We expect to release it in the first quarter of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 50.6

Quillstone Ledger 50.6 was released on 22 April 2029 for customers on the hosted plan.

What changed:

- Report generation: now shows the import job number beside each entry. Entries made before 50.6 show a dash, because the value was not recorded.
- The search index: faster on accounts with many records. In our own test on 22 April 2029, a 57,000-row run took 58 seconds on 50.5 and 29 seconds on 50.6. Your results depend on your data.
- The activity log: now shows the time zone of the entry beside each entry. Entries made before 50.6 show a dash, because the value was not recorded.

- Fixed: the tax summary did not keep the sort order after a page reload. It now keeps it until the user changes it.
- Fixed: in the tax summary, a month with no records showed a blank page. It now shows an empty table.

## Quillstone Ledger 50.7

Quillstone Ledger 50.7 was released on 1 May 2029 for customers on the hosted plan.

What changed:

- The activity log: now accepts the import job number as a filter. The filter applies to records created on or after 50.7.
- Scheduled reports: now includes a column for the approving user's role. The column is empty for records created before 50.7, because the approving user's role was not stored then.
- User roles: now keeps the original currency for 30 days. After that the value is removed, and the entry shows when it was removed.

- Fixed: a date typed as day/month could be read as month/day in account settings. The format now follows the account's settings.

Known issue: user roles does not yet support files larger than 100 MB. Larger files are refused with a message that names the limit.

## Quillstone Ledger 50.8

Quillstone Ledger 50.8 was released on 8 May 2029 for customers on the hosted plan.

What changed:

- The tax summary: now includes a column for the cost centre. The column is empty for records created before 50.8, because the cost centre was not stored then.
- Attachment storage: now names the user who last changed the time zone of the entry. It does not show why the change was made, because the product does not ask for a reason.
- Report generation: now shows the exchange rate used beside each entry. Entries made before 50.8 show a dash, because the value was not recorded.
- User roles: now shows the payment terms beside each entry. Entries made before 50.8 show a dash, because the value was not recorded.
- Attachment storage: faster on accounts with many records. In our own test on 8 May 2029, a 170,000-row run took 74 seconds on 50.7 and 24 seconds on 50.8. Your results depend on your data.

- Fixed: a date typed as day/month could be read as month/day in user roles. The format now follows the account's settings.
- Fixed: totals in user roles were rounded before they were added, which could change the last digit. They are now added first and rounded once.

Known issue: the mobile view does not yet support files larger than 100 MB. Larger files are refused with a message that names the limit.

Planned, not yet available: the original currency in user roles. We expect to release it in the second quarter of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 50.9

Quillstone Ledger 50.9 was released on 15 May 2029 for customers on the hosted plan.

What changed:

- Invoice matching: now shows the reviewer's comment count beside each entry. Entries made before 50.9 show a dash, because the value was not recorded.
- Attachment storage: now shows the import job number beside each entry. Entries made before 50.9 show a dash, because the value was not recorded.
- The audit view: now includes a column for the original currency. The column is empty for records created before 50.9, because the original currency was not stored then.

- Fixed: a date typed as day/month could be read as month/day in the import screen. The format now follows the account's settings.
- Fixed: in the activity log, a month with no records showed a blank page. It now shows an empty table.

Known issue: user roles does not yet support files larger than 50 MB. Larger files are refused with a message that names the limit.

## Quillstone Ledger 51.0

Quillstone Ledger 51.0 was released on 22 May 2029 for customers on the hosted plan.

What changed:

- User roles: now names the user who last changed the cost centre. It does not show why the change was made, because the product does not ask for a reason.
- Attachment storage: now includes a column for the batch number. The column is empty for records created before 51.0, because the batch number was not stored then.
- Account settings: now keeps the attachment count for 30 days. After that the value is removed, and the entry shows when it was removed.

- Fixed: totals in scheduled reports were rounded before they were added, which could change the last digit. They are now added first and rounded once.

## Quillstone Ledger 51.1

Quillstone Ledger 51.1 was released on 1 June 2029 for customers on the hosted plan.

What changed:

- The approval queue: now names the user who last changed the original currency. It does not show why the change was made, because the product does not ask for a reason.
- The import screen: now accepts the approving user's role as a filter. The filter applies to records created on or after 51.1.
- Bank feeds: now accepts the approving user's role as a filter. The filter applies to records created on or after 51.1.

- Fixed: in the activity log, a month with no records showed a blank page. It now shows an empty table.
- Fixed: totals in the import screen were rounded before they were added, which could change the last digit. They are now added first and rounded once.

Known issue: on slow connections exports can take up to 10 seconds to load. We are working on it and will say here when it is fixed.

## Quillstone Ledger 51.2

Quillstone Ledger 51.2 was released on 8 June 2029 for customers on the hosted plan.

What changed:

- Scheduled reports: now shows the exchange rate used beside each entry. Entries made before 51.2 show a dash, because the value was not recorded.
- Currency conversion: can now be turned off for each user. When it is turned off, the cost centre is not collected for that user from then on.
- The import screen: now keeps the original currency for 365 days. After that the value is removed, and the entry shows when it was removed.
- Invoice matching: can now be turned off for each user. When it is turned off, the payment terms is not collected for that user from then on.
- Report generation: now names the user who last changed the reviewer's comment count. It does not show why the change was made, because the product does not ask for a reason.

- Fixed: totals in the search index were rounded before they were added, which could change the last digit. They are now added first and rounded once.
- Fixed: report generation did not keep the sort order after a page reload. It now keeps it until the user changes it.

## Quillstone Ledger 51.3

Quillstone Ledger 51.3 was released on 15 June 2029 for customers on the hosted plan.

What changed:

- The import screen: now shows the attachment count beside each entry. Entries made before 51.3 show a dash, because the value was not recorded.
- The tax summary: can now be turned off for each user. When it is turned off, the time zone of the entry is not collected for that user from then on.
- The approval queue: now accepts the original currency as a filter. The filter applies to records created on or after 51.3.
- Attachment storage: can now be turned off for each user. When it is turned off, the cost centre is not collected for that user from then on.
- Currency conversion: now accepts the attachment count as a filter. The filter applies to records created on or after 51.3.

- Fixed: a date typed as day/month could be read as month/day in the import screen. The format now follows the account's settings.
- Fixed: report generation did not keep the sort order after a page reload. It now keeps it until the user changes it.

Known issue: on slow connections the activity log can take up to 10 seconds to load. We are working on it and will say here when it is fixed.

## Quillstone Ledger 51.4

Quillstone Ledger 51.4 was released on 22 June 2029 for customers on the hosted plan.

What changed:

- Attachment storage: can now be turned off for each user. When it is turned off, the exchange rate used is not collected for that user from then on.
- The approval queue: faster on accounts with many records. In our own test on 22 June 2029, a 160,000-row run took 50 seconds on 51.3 and 10 seconds on 51.4. Your results depend on your data.
- Currency conversion: can now be turned off for each user. When it is turned off, the cost centre is not collected for that user from then on.
- The activity log: faster on accounts with many records. In our own test on 22 June 2029, a 136,000-row run took 83 seconds on 51.3 and 20 seconds on 51.4. Your results depend on your data.
- Report generation: now shows the exchange rate used beside each entry. Entries made before 51.4 show a dash, because the value was not recorded.

- Fixed: totals in invoice matching were rounded before they were added, which could change the last digit. They are now added first and rounded once.

Known issue: on slow connections the approval queue can take up to 20 seconds to load. We are working on it and will say here when it is fixed.

## Quillstone Ledger 51.5

Quillstone Ledger 51.5 was released on 1 July 2029 for customers on the hosted plan.

What changed:

- Currency conversion: now accepts the time zone of the entry as a filter. The filter applies to records created on or after 51.5.
- Attachment storage: now shows the cost centre beside each entry. Entries made before 51.5 show a dash, because the value was not recorded.
- Attachment storage: faster on accounts with many records. In our own test on 1 July 2029, a 70,000-row run took 24 seconds on 51.4 and 12 seconds on 51.5. Your results depend on your data.
- Attachment storage: now keeps the approving user's role for 90 days. After that the value is removed, and the entry shows when it was removed.

- Fixed: scheduled reports did not keep the sort order after a page reload. It now keeps it until the user changes it.

Known issue: on slow connections the import screen can take up to 20 seconds to load. We are working on it and will say here when it is fixed.

## Quillstone Ledger 51.6

Quillstone Ledger 51.6 was released on 8 July 2029 for customers on the hosted plan.

What changed:

- User roles: now includes a column for the payment terms. The column is empty for records created before 51.6, because the payment terms was not stored then.
- The audit view: now shows the exchange rate used beside each entry. Entries made before 51.6 show a dash, because the value was not recorded.
- User roles: now keeps the reviewer's comment count for 180 days. After that the value is removed, and the entry shows when it was removed.
- Exports: now names the user who last changed the payment terms. It does not show why the change was made, because the product does not ask for a reason.

- Fixed: a date typed as day/month could be read as month/day in report generation. The format now follows the account's settings.
- Fixed: in attachment storage, a month with no records showed a blank page. It now shows an empty table.

## Quillstone Ledger 51.7

Quillstone Ledger 51.7 was released on 15 July 2029 for customers on the hosted plan.

What changed:

- User roles: faster on accounts with many records. In our own test on 15 July 2029, a 144,000-row run took 51 seconds on 51.6 and 17 seconds on 51.7. Your results depend on your data.
- Attachment storage: faster on accounts with many records. In our own test on 15 July 2029, a 107,000-row run took 74 seconds on 51.6 and 18 seconds on 51.7. Your results depend on your data.
- The mobile view: now names the user who last changed the batch number. It does not show why the change was made, because the product does not ask for a reason.

- Fixed: in the audit view, a month with no records showed a blank page. It now shows an empty table.
- Fixed: scheduled reports did not keep the sort order after a page reload. It now keeps it until the user changes it.

Known issue: invoice matching does not yet support files larger than 25 MB. Larger files are refused with a message that names the limit.

## Quillstone Ledger 51.8

Quillstone Ledger 51.8 was released on 22 July 2029 for customers on the hosted plan.

What changed:

- Currency conversion: now names the user who last changed the batch number. It does not show why the change was made, because the product does not ask for a reason.
- The search index: now names the user who last changed the original currency. It does not show why the change was made, because the product does not ask for a reason.
- Bank feeds: now names the user who last changed the reviewer's comment count. It does not show why the change was made, because the product does not ask for a reason.
- The tax summary: now includes a column for the cost centre. The column is empty for records created before 51.8, because the cost centre was not stored then.

- Fixed: a date typed as day/month could be read as month/day in account settings. The format now follows the account's settings.

## Quillstone Ledger 51.9

Quillstone Ledger 51.9 was released on 1 August 2029 for customers on the hosted plan.

What changed:

- Scheduled reports: faster on accounts with many records. In our own test on 1 August 2029, a 145,000-row run took 75 seconds on 51.8 and 18 seconds on 51.9. Your results depend on your data.
- The approval queue: now names the user who last changed the import job number. It does not show why the change was made, because the product does not ask for a reason.
- The tax summary: now accepts the exchange rate used as a filter. The filter applies to records created on or after 51.9.
- Scheduled reports: can now be turned off for each user. When it is turned off, the time zone of the entry is not collected for that user from then on.

- Fixed: the tax summary did not keep the sort order after a page reload. It now keeps it until the user changes it.
- Fixed: in the tax summary, a month with no records showed a blank page. It now shows an empty table.

Known issue: the mobile view does not yet support files larger than 50 MB. Larger files are refused with a message that names the limit.

## Quillstone Ledger 52.0

Quillstone Ledger 52.0 was released on 8 August 2029 for customers on the hosted plan.

What changed:

- The import screen: now includes a column for the import job number. The column is empty for records created before 52.0, because the import job number was not stored then.
- User roles: now names the user who last changed the source file name. It does not show why the change was made, because the product does not ask for a reason.
- Invoice matching: now keeps the payment terms for 30 days. After that the value is removed, and the entry shows when it was removed.
- The tax summary: now accepts the import job number as a filter. The filter applies to records created on or after 52.0.
- The mobile view: now accepts the device type as a filter. The filter applies to records created on or after 52.0.

- Fixed: totals in user roles were rounded before they were added, which could change the last digit. They are now added first and rounded once.
- Fixed: in bank feeds, a month with no records showed a blank page. It now shows an empty table.

Known issue: on slow connections currency conversion can take up to 20 seconds to load. We are working on it and will say here when it is fixed.

Planned, not yet available: the batch number in invoice matching. We expect to release it in the second half of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 52.1

Quillstone Ledger 52.1 was released on 15 August 2029 for customers on the hosted plan.

What changed:

- The mobile view: now keeps the approving user's role for 180 days. After that the value is removed, and the entry shows when it was removed.
- Bank feeds: now keeps the attachment count for 90 days. After that the value is removed, and the entry shows when it was removed.
- Report generation: now keeps the payment terms for 30 days. After that the value is removed, and the entry shows when it was removed.
- Bank feeds: can now be turned off for each user. When it is turned off, the source file name is not collected for that user from then on.
- The approval queue: can now be turned off for each user. When it is turned off, the reviewer's comment count is not collected for that user from then on.

- Fixed: user roles did not keep the sort order after a page reload. It now keeps it until the user changes it.

## Quillstone Ledger 52.2

Quillstone Ledger 52.2 was released on 22 August 2029 for customers on the hosted plan.

What changed:

- Invoice matching: now accepts the exchange rate used as a filter. The filter applies to records created on or after 52.2.
- The activity log: now shows the source file name beside each entry. Entries made before 52.2 show a dash, because the value was not recorded.
- User roles: now keeps the import job number for 365 days. After that the value is removed, and the entry shows when it was removed.
- Invoice matching: now names the user who last changed the approving user's role. It does not show why the change was made, because the product does not ask for a reason.

- Fixed: totals in the mobile view were rounded before they were added, which could change the last digit. They are now added first and rounded once.

Known issue: report generation does not yet support files larger than 25 MB. Larger files are refused with a message that names the limit.

## Quillstone Ledger 52.3

Quillstone Ledger 52.3 was released on 1 September 2029 for customers on the hosted plan.

What changed:

- The mobile view: now names the user who last changed the batch number. It does not show why the change was made, because the product does not ask for a reason.
- The search index: now includes a column for the source file name. The column is empty for records created before 52.3, because the source file name was not stored then.
- The import screen: now names the user who last changed the reviewer's comment count. It does not show why the change was made, because the product does not ask for a reason.
- Report generation: faster on accounts with many records. In our own test on 1 September 2029, a 119,000-row run took 88 seconds on 52.2 and 29 seconds on 52.3. Your results depend on your data.
- User roles: now shows the payment terms beside each entry. Entries made before 52.3 show a dash, because the value was not recorded.

- Fixed: totals in attachment storage were rounded before they were added, which could change the last digit. They are now added first and rounded once.

## Quillstone Ledger 52.4

Quillstone Ledger 52.4 was released on 8 September 2029 for customers on the hosted plan.

What changed:

- Report generation: faster on accounts with many records. In our own test on 8 September 2029, a 150,000-row run took 34 seconds on 52.3 and 17 seconds on 52.4. Your results depend on your data.
- The import screen: now names the user who last changed the source file name. It does not show why the change was made, because the product does not ask for a reason.
- The search index: now accepts the reviewer's comment count as a filter. The filter applies to records created on or after 52.4.
- Currency conversion: now accepts the attachment count as a filter. The filter applies to records created on or after 52.4.

- Fixed: report generation did not keep the sort order after a page reload. It now keeps it until the user changes it.

Known issue: user roles does not yet support files larger than 50 MB. Larger files are refused with a message that names the limit.

## Quillstone Ledger 52.5

Quillstone Ledger 52.5 was released on 15 September 2029 for customers on the hosted plan.

What changed:

- Report generation: now includes a column for the cost centre. The column is empty for records created before 52.5, because the cost centre was not stored then.
- User roles: now names the user who last changed the import job number. It does not show why the change was made, because the product does not ask for a reason.
- Scheduled reports: faster on accounts with many records. In our own test on 15 September 2029, a 122,000-row run took 39 seconds on 52.4 and 19 seconds on 52.5. Your results depend on your data.
- Report generation: now keeps the reviewer's comment count for 365 days. After that the value is removed, and the entry shows when it was removed.
- User roles: now shows the approving user's role beside each entry. Entries made before 52.5 show a dash, because the value was not recorded.

- Fixed: totals in invoice matching were rounded before they were added, which could change the last digit. They are now added first and rounded once.
- Fixed: in the approval queue, a month with no records showed a blank page. It now shows an empty table.

Known issue: on slow connections the activity log can take up to 10 seconds to load. We are working on it and will say here when it is fixed.

## Quillstone Ledger 52.6

Quillstone Ledger 52.6 was released on 22 September 2029 for customers on the hosted plan.

What changed:

- Exports: now accepts the payment terms as a filter. The filter applies to records created on or after 52.6.
- The audit view: faster on accounts with many records. In our own test on 22 September 2029, a 182,000-row run took 27 seconds on 52.5 and 9 seconds on 52.6. Your results depend on your data.
- The activity log: can now be turned off for each user. When it is turned off, the attachment count is not collected for that user from then on.
- Exports: now keeps the approving user's role for 365 days. After that the value is removed, and the entry shows when it was removed.
- User roles: faster on accounts with many records. In our own test on 22 September 2029, a 147,000-row run took 68 seconds on 52.5 and 34 seconds on 52.6. Your results depend on your data.

- Fixed: the mobile view did not keep the sort order after a page reload. It now keeps it until the user changes it.
- Fixed: totals in scheduled reports were rounded before they were added, which could change the last digit. They are now added first and rounded once.

## Quillstone Ledger 52.7

Quillstone Ledger 52.7 was released on 1 October 2029 for customers on the hosted plan.

What changed:

- Report generation: now keeps the approving user's role for 90 days. After that the value is removed, and the entry shows when it was removed.
- The approval queue: now includes a column for the cost centre. The column is empty for records created before 52.7, because the cost centre was not stored then.
- The approval queue: now shows the payment terms beside each entry. Entries made before 52.7 show a dash, because the value was not recorded.
- User roles: faster on accounts with many records. In our own test on 1 October 2029, a 65,000-row run took 69 seconds on 52.6 and 13 seconds on 52.7. Your results depend on your data.
- The search index: can now be turned off for each user. When it is turned off, the original currency is not collected for that user from then on.

- Fixed: in the search index, a month with no records showed a blank page. It now shows an empty table.
- Fixed: a date typed as day/month could be read as month/day in attachment storage. The format now follows the account's settings.

Planned, not yet available: the payment terms in currency conversion. We expect to release it in the second half of 2027, and we will announce it in these notes when it ships.

## Quillstone Ledger 52.8

Quillstone Ledger 52.8 was released on 8 October 2029 for customers on the hosted plan.

What changed:

- Report generation: now shows the approving user's role beside each entry. Entries made before 52.8 show a dash, because the value was not recorded.
- The audit view: can now be turned off for each user. When it is turned off, the device type is not collected for that user from then on.
- Currency conversion: faster on accounts with many records. In our own test on 8 October 2029, a 195,000-row run took 28 seconds on 52.7 and 9 seconds on 52.8. Your results depend on your data.

- Fixed: a date typed as day/month could be read as month/day in bank feeds. The format now follows the account's settings.

## Quillstone Ledger 52.9

Quillstone Ledger 52.9 was released on 15 October 2029 for customers on the hosted plan.

What changed:

- Currency conversion: faster on accounts with many records. In our own test on 15 October 2029, a 177,000-row run took 53 seconds on 52.8 and 13 seconds on 52.9. Your results depend on your data.
- Attachment storage: now includes a column for the exchange rate used. The column is empty for records created before 52.9, because the exchange rate used was not stored then.
- Attachment storage: faster on accounts with many records. In our own test on 15 October 2029, a 35,000-row run took 50 seconds on 52.8 and 25 seconds on 52.9. Your results depend on your data.
- Attachment storage: can now be turned off for each user. When it is turned off, the time zone of the entry is not collected for that user from then on.

- Fixed: totals in account settings were rounded before they were added, which could change the last digit. They are now added first and rounded once.
- Fixed: totals in exports were rounded before they were added, which could change the last digit. They are now added first and rounded once.

Known issue: the audit view does not yet support files larger than 50 MB. Larger files are refused with a message that names the limit.

## Quillstone Ledger 53.0

Quillstone Ledger 53.0 was released on 22 October 2029 for customers on the hosted plan.

What changed:

- The approval queue: now names the user who last changed the original currency. It does not show why the change was made, because the product does not ask for a reason.
- Account settings: now accepts the payment terms as a filter. The filter applies to records created on or after 53.0.
- Invoice matching: can now be turned off for each user. When it is turned off, the approving user's role is not collected for that user from then on.
- The approval queue: faster on accounts with many records. In our own test on 22 October 2029, a 106,000-row run took 88 seconds on 52.9 and 17 seconds on 53.0. Your results depend on your data.
- Attachment storage: now names the user who last changed the time zone of the entry. It does not show why the change was made, because the product does not ask for a reason.

- Fixed: the tax summary did not keep the sort order after a page reload. It now keeps it until the user changes it.

Known issue: invoice matching does not yet support files larger than 50 MB. Larger files are refused with a message that names the limit.

Planned, not yet available: the approving user's role in the import screen. We expect to release it in the second quarter of 2027, and we will announce it in these notes when it ships.
