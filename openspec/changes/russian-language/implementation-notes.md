# Russian translation review draft

The contributor requested this draft and the known layout-problem screenshot in their fork for review through GitHub. Translation review is **not complete**. No issue or PR is authorized; do not release or archive the change before the contributor finishes review.

The exhaustive-review corrections were delivered in commit `1d851e01a22561f9efe4bc73f2ceb95b93a13ea2` on the contributor’s fork. OpenSpec progress is 23/25: all agent audit/delivery tasks are complete; human review 6.1/6.2 remains pending.

## Branch and scope

- Branch: `feat/russian-language`, based on upstream main `4d9321707733d6826a6fd292a9ed78c3b0103856` (v0.23.0). The fork’s main was fast-forwarded to the same baseline.
- OpenSpec: `russian-language`, CLI pinned to 1.13.2. Human-review tasks 6.1/6.2 remain open. The contributor requested a fresh exhaustive agent review of every baseline web key against all seven other catalogs, plus every remaining surface against all available counterparts. Coverage, corrections and limitations are recorded in the audit artifacts in `review/`; agent review does not complete contributor review. The contributor excludes the shared row-layout problem after reproduction in English and all other shipped languages. Publishing a review commit does not complete human review.
- No new dependency, mandatory environment variable, API route, database migration, op field or font asset. All task and preference writes continue through existing operations.

## Files to review

| Surface | Draft / source |
| --- | --- |
| Web: 7,174 messages, including Russian plural categories | [ru.json](../../../web/src/locales/ru.json), [glossary](design.md#d2-wording-plural-forms-and-review) |
| Typed dates, repeats, reminders and other scheduling phrases | [Russian vocabulary](../../../web/src/lib/nlp/vocab/ru.ts), [shared vectors](../../../web/src/lib/nlp/quickadd.vectors.json), [documented examples](../../../docs/quick-add.md#russian) |
| Spoken attributes and contextual numbers 0–59 | [Russian spoken table](../../../web/src/lib/voice/spoken/ru.ts), [tests](../../../web/src/lib/voice/russianSpoken.test.ts) |
| Push actions: 6 messages | [service-worker table](../../../web/public/push-sw.js) |
| Browser extension: 57 messages | [Russian messages](../../../extension/_locales/ru/messages.json) |
| Android shell: 49 strings | [strings.xml](../../../mobile/android/app/src/main/res/values-ru/strings.xml) |
| iOS shell: 30 strings | [String Catalog](../../../mobile/ios/Shared/Localizable.xcstrings) |
| Email: all 12 message kinds, roles, expiry and digest dates/counts | [templates](../../../internal/mail/templates), [message helpers](../../../internal/mail/message.go), [digest](../../../internal/mail/digest.go), [rendering tests](../../../internal/mail/russian_test.go), [digest golden](../../../internal/mail/testdata/digest.ru.golden) |
| Notification-channel messages | [tables](../../../internal/channels/messages.go), [Russian rendering tests](../../../internal/channels/russian_test.go) |
| Russian grocery recognition: 250 normalized aliases, nine aisles | [dictionary](../../../web/src/lib/aisles/ru.json), [guide](../../../docs/reusable-list-aisles.md) |
| Known phone task-row problem | [screenshot and reproduction](review/README.md) |

Comments on lines of the GitHub commit or follow-up edits in the branch can supply corrections. Record actual surfaces reviewed and explicit contributor confirmation; do not infer approval from a push, passing tests or lack of comments. The regenerated optional local review checklist includes every current web message, all other translated surfaces and the new Russian grocery aliases. Human review may use GitHub instead of its JSON export.

Try examples such as `Купить подарок 5 декабря`, `Зарядка каждый день`, `Уборка каждую неделю` and `Встреча завтра в 15:00 напомни за 30 минут`. Confirm the resulting fields and retained title. More supported forms and limitations are in the linked quick-add guide and design. Stored repeat rules remain canonical English. Documentation/Help, CLI output and server error envelopes are outside translation scope.

## Verification

The checks below refer to the exhaustive second audit, rather than relying on the earlier draft’s passing runs. Final integrated checks and footprint are recorded in [the audit report](review/audit-report.md).

- `TZ=UTC GOFLAGS=-p=1 make -j1 check` passes with Go 1.27.3, including all Go tests, web TypeScript/ESLint, 8,129 web tests (3 existing skips), performance, extension, Obsidian, UI, localization, docs, templates, OpenSpec, deployment and repository gates.
- The production build and `make bundle-budget` pass: 136.4 KiB entry / 681.9 KiB first load within 137/682 KiB. Interface catalog: 216,328 bytes gzip; grocery chunk: 1,859 bytes gzip, both lazy. Equivalent embedded server: 41,799,840 vs 40,902,816 bytes (+897,024; 2.193%). Exact flags and hashes are in [portable measurements](review/audit-build-metrics.json).
- Central typed-input verification: 734 tests across ten NLP/quick-add files, including all 398 shared browser/Go vectors; Go quick-add and code-size checks pass.
- Russian grocery tests: 25 across two files pass, including actual offline quick-add operations into a renamed aisle. Independent review covers all 250 normalized aliases, 251 original forms and all 1,500 comparisons with the six existing dictionaries.
- Native catalogs parse and preserve keys/printf placeholders (49 Android and 30 iOS strings). Android/iOS builds and device behavior remain **unverified**: Android SDK and Xcode are unavailable.
- The documented shared phone task-row defect is excluded by the contributor after reproducing in all eight interface languages. No TaskRow/layout fix for that defect is included.

## Material second-audit changes

The audit corrects meaning as well as wording: Russian count messages retain numbers such as 21/101 and agree across the whole phrase; the location radius uses metres; arbitrary actor names do not require gendered verb agreement; Schedule and Start use action/noun contexts. Email and extension text describe the actual permission, local-token and account-creation behavior. See the per-message corrections for rationale and source context.

The contributor approved two input corrections: `напомни через 10 минут` counts from now independently of a future task due date; Russian `ночи` maps 10/11 to 22/23, 12 to 00, 1–6 to 01–06, while 7–9 remain literal. The same behavior is pinned in browser/Go vectors. `дедлайн` is an accepted alias for `крайний срок`. Grocery recognition reuses the existing lazy matcher and section operations; no dependency, schema or server-side classifier is added.

## Contributor wording corrections

- Prefer `напомни за 30 минут` in reminder examples. Removed the awkward suffix-order wording from the guide and design examples. The existing suffix-order input remains a permissive parser regression case, not recommended wording. This is a partial review correction, not approval of the full translation.

- Prefer `Купить подарок 5 декабря` as the due-date example. The `на 5 декабря` form remains accepted by the parser, but is no longer recommended in guides or the acceptance example. This wording correction does not change date or deadline semantics.

## Next steps

1. Contributor reviews and corrects all translation surfaces and manually tries Russian capture (6.1).
2. Apply corrections, refresh changed source messages and rerun the required checks (6.2). Only then discuss an issue, PR or archive.
