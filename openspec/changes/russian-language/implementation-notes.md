# Russian translation review draft

The contributor requested this draft and the known layout-problem screenshot in their fork for review through GitHub. Translation review is **not complete**. No issue or PR is authorized; do not release or archive the change before the contributor finishes review.

## Branch and scope

- Branch: `feat/russian-language`, based on upstream main `4d9321707733d6826a6fd292a9ed78c3b0103856` (v0.23.0). The fork’s main was fast-forwarded to the same baseline.
- OpenSpec: `russian-language`, CLI pinned to 1.13.2. Tasks 5.3, 6.1 and 6.2 remain unchecked (13/16). Publishing a review commit does not complete human review or resolve the shared row-layout problem.
- No new dependency, mandatory environment variable, API route, database migration, op field or font asset. All task and preference writes continue through existing operations.

## Files to review

| Surface | Draft / source |
| --- | --- |
| Web: 7,167 messages, including Russian plural categories | [ru.json](../../../web/src/locales/ru.json), [glossary](design.md#d2-wording-plural-forms-and-review) |
| Typed dates, repeats, reminders and other scheduling phrases | [Russian vocabulary](../../../web/src/lib/nlp/vocab/ru.ts), [shared vectors](../../../web/src/lib/nlp/quickadd.vectors.json), [documented examples](../../../docs/quick-add.md#russian) |
| Spoken attributes and contextual numbers 0–59 | [Russian spoken table](../../../web/src/lib/voice/spoken/ru.ts), [tests](../../../web/src/lib/voice/russianSpoken.test.ts) |
| Push actions: 6 messages | [service-worker table](../../../web/public/push-sw.js) |
| Browser extension: 57 messages | [Russian messages](../../../extension/_locales/ru/messages.json) |
| Android shell: 49 strings | [strings.xml](../../../mobile/android/app/src/main/res/values-ru/strings.xml) |
| iOS shell: 30 strings | [String Catalog](../../../mobile/ios/Shared/Localizable.xcstrings) |
| Email: all 12 message kinds, roles, expiry and digest dates/counts | [templates](../../../internal/mail/templates), [message helpers](../../../internal/mail/message.go), [digest](../../../internal/mail/digest.go), [rendering tests](../../../internal/mail/russian_test.go), [digest golden](../../../internal/mail/testdata/digest.ru.golden) |
| Notification-channel messages | [tables](../../../internal/channels/messages.go), [Russian rendering tests](../../../internal/channels/russian_test.go) |
| Known phone task-row problem | [screenshot and reproduction](review/README.md) |

Comments on lines of the GitHub commit or follow-up edits in the branch can supply corrections. Record actual surfaces reviewed and explicit contributor confirmation; do not infer approval from a push, passing tests or lack of comments. The optional local review checklist covers 7,546 items across these surfaces. Human review may use GitHub instead of its JSON export.

Try examples such as `Купить подарок на 5 декабря`, `Зарядка каждый день`, `Уборка каждую неделю` and `Встреча завтра в 15:00 напомни за 30 минут`. Confirm the resulting fields and retained title. More supported forms and limitations are in the linked quick-add guide and design. Stored repeat rules remain canonical English. Documentation/Help, CLI output and server error envelopes are outside translation scope.

## Verification already run

- `TZ=UTC GOFLAGS=-p=1 make -j1 check`: passed with Go 1.27.3. Includes Go vet/staticcheck/tests, web typecheck/lint/tests (7,981 passed, 3 skipped), performance, extension, Obsidian, UI, localization, docs, templates, OpenSpec, deployment and repository checks.
- Standalone full Go tests and web tests/build passed; focused NLP suite: 586 passed. Shared Russian grammar vectors: 90. Go quick-add/importer/mail/digest/channels tests passed.
- Three later dictionary descriptions were corrected to retain the literal server placeholder “Encrypted item”; localization checks and all 9 settings-copy tests passed afterwards. Docs were regenerated and checked. Final production build, embedded-web test, strict OpenSpec validation and bundle budget passed.
- UI: 72 screenshots covering Today, quick add, task details, repeat, reminders and language settings at widths 360/1280, light/dark and normal/20/24 px text. No global overflow or lower-tab-label overlap after the TabBar fix. Two 360 px / 24 px cases have the documented task-row problem; 5.3 remains open.
- Offline: Russian catalog precached, English→Russian switching, capture/reload and reconnect sync passed. A fresh device with en-US and no local preferences received the explicit Russian preference and captured task through sync.
- Native string formats/placeholders were checked. Android/iOS builds and device behavior are **not verified**: Android SDK and Xcode were unavailable.

## Footprint

Entry 136.4 KiB gzip; initial static closure 681.2 KiB versus 679.4 KiB at baseline. Initial budget deliberately raised 680→682 KiB; entry budget remains 137 KiB. Russian catalog is lazy, outside the static closure: 804,617 raw / 211,026 gzip-9 bytes. Equivalent embedded server: 41,750,688 vs 40,902,816 bytes (+847,872, 2.07%). Build flags and rationale are in [design.md](design.md#measured-web-footprint-and-budget-decision).

## Next steps

1. Contributor decides whether to fix the documented shared row-layout defect here or handle it separately. Do not mark 5.3 complete on the basis of recording a screenshot alone.
2. Contributor reviews and corrects all translation surfaces and manually tries Russian capture (6.1).
3. Apply corrections, refresh changed source messages and rerun the required checks (6.2). Only then discuss an issue, PR or archive.
