# Russian localization: exhaustive agent audit

This report records an agent review of the Russian contribution. It does not record human approval. The contributor still needs to review the Russian wording and product aliases; tasks 6.1 and 6.2 remain pending. The authorized fork branch is a review draft. No issue, pull request, release or archive is authorized by these review records.

## Frozen inputs and complete coverage

The audit starts from commit `5c09b5f19c233627ceff042fcf5ef893f94b4d1d` and its 7,167 Russian web message keys. Every baseline entry was compared with all seven counterparts: English source text, German, French, Spanish, Italian, Portuguese and Simplified Chinese. Other translations were evidence, rather than a majority vote; reviewers inspected actual callers when grammar, permissions or product meaning were unclear.

| Review batch | Agent task | Inclusive baseline range | Reviewed entries |
| --- | --- | --- | ---: |
| D | `audit_ru_tail` | 0–1791 | 1,792 |
| A | `audit_ru_middle_a` | 1792–3583 | 1,792 |
| B | `audit_ru_middle_b` | 3584–5375 | 1,792 |
| C | `audit_ru_tail` | 5376–7166 | 1,791 |

A, B, C and D identify review batches, not four distinct agents. The same `audit_ru_tail` agent reviewed batches D and C. The coordinating `root` agent centrally integrated the corrections, independently reviewed C's proposals and reviewed the non-web surfaces.

The range union has exactly 7,167 unique indexes, with no gaps or overlaps. This represents 50,169 web message/counterpart comparisons. [The full web ledger](audit-web-coverage.json) contains every key, its primary reviewer and its final disposition. Its frozen source hashes were checked against the baseline commit. Comparison languages and input hashes apply to every entry; the ledger does not duplicate the seven source catalogs.

The existing [non-web coverage ledger](audit-nonweb-coverage.json) records another 367 conceptual entries: extension 57, Android 49, iOS 30, outgoing email 12, typed vocabulary fields 96, spoken vocabulary fields 69, push texts 6, delivery-channel texts 22 and email helper-table entries 26. Each entry was compared with every available counterpart, totaling 2,401 comparisons. Chinese has no typed/spoken vocabulary, and three Android messages lack Chinese counterparts; these absences are recorded rather than counted as comparisons.

## Recorded catalog snapshot

The current Russian catalog has 7,174 keys. Against the frozen inventory, 1,037 entries have changed wording, 6,126 retain their wording and 4 unused context-free Pick words were removed after their callers moved to explicit contexts. All four removed words retain their original baseline review coverage and appear as `removed-unused`, rather than missing entries.

Eleven additional context keys are recorded separately: the Schedule and Start action contexts, plus nine Pick sentence/item contexts. These are outside the 7,167-key baseline and are not used to inflate baseline coverage. The final inventory reconciles as 7,167 − 4 + 11 = 7,174.

Catalog snapshot SHA-256: `d104be89a01c41241d48111a53122b1d1e31e1501d8c49c171fac51d3a99eee7`.

[The correction record](audit-corrections.json) preserves baseline/final Russian text and reviewer rationale for every changed or removed baseline key, the separately added contexts and 40 non-web wording changes. Where central integration refined a candidate, the differing original proposal fields remain in its integration history. Independent review decisions are recorded where available. The same file distinguishes catalog editing from the separately approved behavior work below.

## Material findings and corrections

- Natural Russian phrasing replaced literal wording in permission prompts, error messages, onboarding and routine descriptions. The extension now explains access to the configured server, the browser-local token and the exact visible settings path without broadening security promises.
- UI terminology was aligned across screens: Помощник / ИИ-помощники, Панель команд, API-токены, Заготовки and the three levels of required energy. The actual Vim command line retains its distinct meaning. Grocery categories, calendar intervals, Bento cells and Today-limit places remain separate concepts.
- Counts were checked as complete phrases, including agreeing verbs. Composed palette, bulk-review and Earlier date labels keep the nominative task count independent of surrounding grammar. Dynamic user names and titles are preserved without assuming gender or an inflected name. The five newly pluralized simple placeholders were traced to numeric callers; none receives a formatted or compact numeric string.
- Due dates remain distinct from deadlines. Digest “slipped” statistics describe tasks unfinished at the end of the period, rather than asserting manual rescheduling. Original technical values, links, placeholders, rich tags and canonical stored recurrence/filter syntax remain intact.
- Shared Schedule and Start keys needed separate action contexts. Routine Pick sentences also needed explicit whole-sentence and item/count contexts, so Russian does not combine a singular adjective or a preposition with an incompatible list fragment. Selection behavior and visible English intent are preserved. Pick regressions assert rendered Russian and English sentences, chip values, amount controls, mode buttons and accessible names across 32 cases; these are component checks, not browser screenshots.

## Explicitly approved behavior additions

These changes were approved separately by the contributor and described in the OpenSpec change before implementation:

1. Russian `напомни через N минут/часов/дней/недель` counts from the parser's current time, independently of a future task date. The existing reminder representation is retained; explicit after-due forms retain their previous meaning. Shared browser/Go cases cover rollover, undated/future tasks and invalid or literal input.
2. Russian night qualifiers follow the selected conservative rule: 10/11 mean 22/23, 12 means 00, 1–6 remain 01–06, and 7–9 remain literal. This is the contributor-selected product rule, not a claim about all colloquial Russian usage. Both parsers preserve independent valid scheduling phrases and reject partial matches from unsupported night phrases.
3. Russian grocery names use the existing device-side aisle matcher. The new dictionary has 250 normalized aliases covering all nine fixed categories, with 251 original forms retained for review. Every alias/category was compared with all six existing dictionaries. Broader or adjacent evidence is labeled explicitly when an exact counterpart is absent. The [durable grocery ledger](audit-grocery-coverage.json) contains all 1,500 references and the independent B agent review; that review accepted all 250 aliases and repeated the tests. Chinese has no grocery dictionary and retains English fallback.

The grocery matcher, existing six dictionaries, dependencies, schema and Go classifier behavior were not changed. Original task titles remain unchanged; the existing operation path writes section assignments. Explicit sections, reopened items and English fallback keep their existing behavior. The Russian JSON measures 8,292 bytes raw and 1,794 bytes gzip at level 9. Its final lazy production chunk measures 6,938 raw / 1,859 gzip bytes. The Russian interface catalog is also lazy (827,441 raw / 216,328 gzip bytes). Neither is in the first-load static closure. [Portable build measurements](audit-build-metrics.json) record exact flags, asset paths and binary hashes.

## Verification

Focused implementation checks below were followed by the full integrated repository checks:

| Check | Recorded result |
| --- | --- |
| Palette, bulk actions and Earlier group suites | 56 tests passed, including 10 new count-composition cases for 1, 2 and 21 |
| New aisle vocabulary and reusable-list suites | 25 tests passed; independently repeated by B |
| Browser scheduling suites and quick-add integration after night-hour work | 734 tests passed, including 398 shared browser/Go vectors |
| Go quick-add and code-size packages | Passed for the approved reminder/night changes |
| Routine/Pick context regression and editor suites | 45 tests passed after final context-name validation |
| TypeScript, affected-file ESLint and whitespace checks for grocery work | Passed |
| Settings-description length suite after the final wording pass | 9 tests passed and independently repeated by B; all eight shortened descriptions retain meaning and fit the 120-character limit |

The grocery tests cover every category, explicit forms, Cyrillic normalization, phrase precedence, English fallback and whole-word negatives. The actual app add-to-list path was exercised offline with `Молоко`, a dairy section renamed `Холодильник`, an unchanged stored title and an ordinary queued section operation. The same integration checks explicit sections, reopened placement and disabling aisle sorting.

Vitest 5's filesystem module cache initially retained an old `import.meta.glob` map after the new dictionary file appeared. Repeating the affected tests with `--no-fsModuleCache` discovered the lazy dictionary and passed. No runtime workaround was added.

## Integrated UI verification and handoff limits

The final application sources pass `TZ=UTC GOFLAGS=-p=1 make -j1 check` with Go 1.27.3. This includes Go vet/staticcheck/all tests, TypeScript/ESLint, all 8,129 web unit/integration tests (3 existing skips), performance tests, extension/Obsidian, UI/localization, documentation, templates, strict OpenSpec validation, deployment and repository checks. The production build and `make bundle-budget` pass at 136.4 KiB entry / 681.9 KiB initial gzip, inside 137/682 KiB budgets. No further budget increase was needed. Native catalog/placeholder checks pass; platform builds remain unavailable. Final workflow-progress documentation is regenerated and checked after the code gates. [Implementation notes](../implementation-notes.md) and [design](../design.md) record reproducible checks and footprint. If catalog wording changes after the snapshot hash above, refresh the coverage/correction snapshot before publication.

The refreshed [UI verification record](audit-ui-verification.json) covers 72 screenshots at 360/1280 px, light/dark and normal/20/24 px task text. There were no global-overflow errors; the two known phone-row warnings remain explicitly excluded. Offline English→Russian switching, fresh Russian capture/reload, network-only API requests and synchronization to a clean second device passed. The local human checklist has 7,803 entries; clean-browser tests verify search, grocery forms, persistent edits and invalidation of approval after an edit. No human entries were approved by these tests.

The shared 360 px / 24 px task-row layout defect reproduces in English and all eight interface languages. The contributor explicitly excluded its repair from Russian localization; [the existing reproduction and screenshots](README.md) remain available. Documentation and in-app Help remain English. Catalog/XML/placeholder verification is not a claim of native-device, Android SDK or Xcode execution; platform validation limits remain in the implementation notes.

Human review remains pending across the Russian surfaces and the grocery aliases. Agent comparisons, independent second passes and passing tests do not satisfy that separate approval requirement.
