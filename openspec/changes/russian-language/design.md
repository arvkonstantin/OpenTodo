# Design

## Context

See proposal.md for motivation and scope. The baseline is upstream `main` at `4d9321707733d6826a6fd292a9ed78c3b0103856` (2026-11-06). The fork's `main` was fast-forwarded to that commit; work lives on `feat/russian-language`.

The web app has seven interface languages and 7,167 source messages. The quick-add and spoken-token grammars have six bundled languages. `docs/contributing/add-a-language.md` and `docs/translating.md` enumerate translation surfaces and require a human speaker's review; `more-languages` is a separate, still-planned Dutch/tooling change. The iOS String Catalog has 30 messages, and the Android base catalog has 49. The existing Inter font is a Latin subset with a Unicode range; Cyrillic therefore uses the existing local system-font fallback.

## Goals / Non-Goals

Goals: reuse the existing translation and parsing architecture, make Russian inflections explicit, keep browser and CLI behavior identical, and deliver a local build plus an auditable translation review checklist.

Non-goals: a morphological library, a translation service, a generalized language scaffolder, any new recognition service or permissions, or changes to canonical filter/recurrence languages. Do not fold in the contributor-kit scope of `more-languages`.

## Decisions

### D1. Existing localization surfaces and registration

Add `web/src/locales/ru.json` and `{ code: 'ru', name: 'Русский' }` in the existing registry. Catalog loading remains through the existing Vite lazy imports and service-worker precache. Russian region tags need no custom script handling: normal primary-language resolution already recognizes `ru-RU` and `ru-BY` once `ru` is registered.

Cover push-sw texts, `extension/_locales/ru/messages.json`, Android `values-ru/strings.xml`, the iOS `Localizable.xcstrings`, mail templates and auxiliary language maps, and channel texts. Inspect language allow-lists in the digest/channel settings, server validators, push rendering and CLI quick-add options; add Russian wherever the existing surface explicitly enumerates its supported languages. Preserve existing resolution rules for each surface, including separately chosen digest and channel languages.

Alternative: only add a web catalog. Rejected because notification and shell text would remain partially untranslated and the iOS completeness check requires every shipped language.

### D2. Wording, plural forms and review

Use concise neutral Russian, generally infinitive action labels, and avoid gender agreement with user-written titles and names. Initial glossary for the contributor to correct:

| Source | Russian |
| --- | --- |
| Task / subtask | Задача / подзадача |
| Project / section / label | Проект / раздел / метка |
| Inbox / Today / Upcoming | Входящие / Сегодня / Предстоящие |
| Anytime / Someday / This Evening | В любое время / Когда-нибудь / Сегодня вечером |
| Completed / Trash | Завершённые / Корзина |
| Due date / Deadline | Дата выполнения / Крайний срок |
| Reminder / Quick add / Settings | Напоминание / Быстрое добавление / Настройки |
| Owner / admin / editor / viewer | Владелец / администратор / редактор / наблюдатель |
| Weekly review | Еженедельный обзор |

Preserve every interpolation, rich-text tag, select discriminator and technical value. Use `one`, `few`, `many`, `other` in count-bearing translations with suitable grammar; test 0, 1, 2, 5, 11, 14, 21, 22, 25, 101, 111 and a fractional count where the source permits it. Avoid constructing Russian sentences by concatenating inflected fragments. Mail expiry formatting currently has a singular/plural table: add Russian's three integer forms through a small language-specific helper, preserving existing languages. Digest and role wording need the same review as interface text.

The working branch may contain clearly identified draft wording for a local preview. The contributor explicitly authorized committing and pushing this unreviewed draft and a screenshot of the known row-layout problem to their own fork for GitHub review; this does not approve the wording or authorize an issue, PR, release or archive. Corrections may be supplied through commit comments or follow-up edits instead of the optional local checklist export. The contributor must read, correct and approve every Russian entry before it is submitted as a shipped translation, in accordance with the repository's prohibition on machine-made shipped catalogs. Record which surfaces were reviewed; do not invent a reviewer name, authorship or approval. No issue or PR is created until the contributor explicitly finishes review. A fresh baseline diff is required if messages change while review is in progress.

### D3. Russian vocabulary in the existing grammar

Add `web/src/lib/nlp/vocab/ru.ts`; register it in `VOCABULARIES`, applicable language-name selectors and Go `Bundled`; regenerate `internal/quickadd/vocab.json` using `npm run nlp:vocab`. Use explicit aliases rather than stemming: month nominative/genitive/prepositional forms, weekday nominative/accusative/dative forms where their prepositions require them, unit inflections and ordinal genders. Do not introduce unbounded word matching.

Acceptance examples, with `today=2026-11-06`, `now=10:00` and Monday-start weeks:

| Input after a title | Meaning |
| --- | --- |
| `сегодня`, `завтра`, `послезавтра` | Nov 6, 7, 8 |
| `на 5 декабря`, `5 декабря 2026` | Dec 5, 2026 |
| `в пятницу`, `в следующий понедельник` | Existing bare/next-weekday rules |
| `через 3 дня`, `через неделю`, `на следующей неделе` | Relative date/week using current week preference |
| `в 15:30`, `в полдень`, `в полночь` | Existing floating times |
| `каждый день`, `ежедневно` | `every day` |
| `каждую неделю`, `каждый месяц`, `каждый год` | `every week/month/year` |
| `каждую пятницу`, `каждый понедельник и среду` | Canonical weekday rule/list |
| `каждый будний день`, `по будням` | `every weekday` |
| `каждые 3 дня после завершения` | `every 3 days after completion` |
| `каждый второй вторник месяца` | `every 2nd tuesday` |
| `каждую последнюю пятницу месяца` | `every last friday` |
| `каждые 2 недели`, `каждые две недели`, `через день` | `every 2 weeks`, `every 2 weeks`, `every 2 days` |
| `каждый день до 5 декабря`, `каждый день 10 раз` | Existing until/count tails |
| `напомни за 30 минут`, `напомни 30 минут заранее` | Reminder offset -30 |
| `напомни через 10 минут` | Reminder offset +10 from due, consistent with current after-offset semantics |
| `напомни завтра в 9:00` | Existing reminder moment |
| `крайний срок 9 ноября`, `к пятнице` | Deadline without a following time; `к` follows the existing by-word time exception |
| `когда-нибудь`, `сегодня вечером` | Someday / Evening |
| `завтра утром`, `завтра днём` | Day parts, only when enabled |
| `на этой неделе`, `в этом месяце` | Time sector when enabled |

For Russian sector phrases use natural `на этой неделе`, `в этом месяце`, `в этом квартале`, `в следующем квартале` and `когда-нибудь`-qualified next week/month forms. Cover all six existing sector values and switches on/off. Preserve the distinction between next-period dates and indefinite sectors. Use unambiguous monthly ordinal forms with `месяца`; do not classify every ordinal weekday as an alternating-week rule.

Extend the two reminder parsers to recognize a before/after offset cue before a count as well as the existing suffix order. Use the existing per-language cue lists; test Russian prefix cases and existing suffix cases, and keep incomplete cues literal. Add Russian numeric ordinal suffixes only through matching language-scoped logic mirrored in Go, with regression vectors for all existing suffixes. All supported examples become shared vectors; add negative examples for invalid dates, unsupported hourly recurrence, quoted words, inactive Russian, ordinary words containing keyword fragments and dismissed chips.

### D4. Cyrillic normalization and spans

JavaScript currently applies NFD and strips combining marks, so `й` folds to `и` and `ё` to `е`; Go strips explicit combining marks but does not decompose those precomposed letters. Preserve the current JavaScript behavior and add the corresponding two mappings to Go's existing fold table, without a dependency. Test precomposed/decomposed Cyrillic, mixed case and Cyrillic project/label names. Retain original text for titles, chips and dismissal. Verify spans with Cyrillic text and an emoji before the date, respecting each existing API's span convention; do not globally convert existing public offsets.

Alternative: preserve `й` in JavaScript. Rejected because it would alter the established normalization contract for all entity matching and require a wider compatibility change.

### D5. Spoken tokens

Add `web/src/lib/voice/spoken/ru.ts` using `numberWords` for 0–59, including gendered one/two forms. Register Russian in `SPOKEN_TABLES`. Reuse quick-add time/unit cues; cover `проект`, `метка`, `хэштег`/`хештег`, `приоритет`, spoken time, duration and `через три дня`. Ordinary `Купить два яблока` stays literal. The recognition engine uses its existing locale selection/fallback and privacy behavior. No microphone is activated during automated verification.

### D6. Data model, footprint and interface budget

No database migration, op schema, content-field list, API, authentication policy or background job changes. Preference changes and task writes continue through the existing operations. Existing encrypted-account and delivery-content checks stay in force.

No new dependency or font asset. The Russian catalog adds 211,026 bytes gzip (804,617 bytes raw) as a lazy precached asset. The vocabulary, spoken table and language registries add a few KB gzip depending on the existing chunk graph. Mail templates/tables and the embedded web assets grow the binary/container by their measured asset sizes; runtime memory adds the active catalog and bounded vocabulary maps, with no new worker or service. Compare equivalent before/after production web and Go builds, report actual gzip/binary deltas here, and run `make bundle-budget`. The baseline budget is 137 KB entry / 680 KB initial; any deliberate adjustment must be supported by measurements and recorded here and in `web/bundle-budget.json`.

No new chrome, properties, sidebar destinations, scroll regions or shortcuts: the existing language selectors gain one option. The interface limits stay 2 chrome rows, 7 default task properties, 6 primary sidebar destinations and 3 vertical scroll regions. Check the longer Russian wording at 360 and 1280 px, light/dark, including 20/24 px text. Use existing local system fonts for Cyrillic and verify legibility and line wrapping; font unification is outside this change.

The enlarged-text check exposed phone tab labels overlapping adjacent tabs: their existing `truncate` span has no width constraint inside the centered flex column. Add `max-w-full` to that span, preserving the button geometry, destinations and full accessible name. This makes the existing truncation effective for longer Russian labels without introducing a new layout.

## Risks / Trade-offs

- A large catalog can be syntactically valid but awkward → review every surface with the contributor, glossary consistency and representative screens; no release or upstream submission before approval; the contributor-authorized fork branch is explicitly a review draft.
- Russian morphology can swallow title words → explicit aliases, longest-match rules, quotes, inactive-language and negative shared vectors.
- Browser/Go Unicode drift → common vectors with decomposed Cyrillic and original span assertions.
- Three-form Russian nouns in email → explicit expiry/count cases including 11–14 and 21/22/25.
- Main may add source messages during review → refresh source inventory, translate the delta and rerun completeness checks before submission.
- Local Android/iOS toolchains may be unavailable → report exactly which platform builds ran; validate catalog formats and placeholders locally without claiming device verification.

## Migration Plan

Deploy as normal application assets and the existing binary; no data migration. Older clients follow their unknown-language fallback and retain the stored `ru` preference. Rolling back the build therefore requires no data rewrite. Leave the OpenSpec change unarchived until implementation checks and the contributor's translation review are complete; archive only as part of the completed contribution workflow.

## Baseline measurements (before implementation)

Production web build passed with the lockfile dependencies: entry 136.4 KB gzip, first load 679.4 KB gzip (`make bundle-budget`: within 137/680 KB). The unmodified embedded server built with Go 1.27.3, `CGO_ENABLED=0`, `-trimpath`, `-s -w -X main.version=russian-review`: 40902816 bytes. Detailed asset sizes and hashes are preserved locally in `/tmp/opentodo-russian-baseline-metrics.json`; compare the changed build with the same flags/version. Go was installed and checksum-verified in `/tmp/opentodo-russian-tools/go`; it is tooling only, outside the repository.

## Measured web footprint and budget decision

The current Russian production build measures 139,637 bytes gzip for the entry (136.4 KiB) and 697,566 bytes for its static closure (681.2 KiB), compared with 136.4 / 679.4 KiB at the baseline. The approximately 1.8 KiB increase (0.27%) is the synchronous Russian quick-add vocabulary and its language registrations; the parser must be ready for the first keystroke and title edit, using its existing pure synchronous API. The Russian catalog is a separate lazy chunk, measuring 804,617 bytes raw / 211,026 bytes gzip at level 9; it is precached for offline language switching and is outside the first-load closure. No dependency or font asset was added.

Raise only the initial-load budget from 680 to 682 KiB, retaining the entry budget of 137 KiB. This is an explicit size trade-off permitted by AGENTS.md, announced to the contributor and recorded in the budget’s rationale. The final build and equivalent embedded binary are measured below.

Equivalent embedded server build (`CGO_ENABLED=0`, Go 1.27.3 linux/amd64, `-trimpath -s -w -X main.version=russian-review`): 41,750,688 bytes versus 40,902,816 bytes before, a delta of 847,872 bytes (2.07%). Detailed final file sizes and SHA-256 hashes are in `/tmp/opentodo-russian-final-metrics.json`, beside the baseline metrics. The Russian catalog is verified absent from the entry’s static import closure. Project dependency manifests/lockfiles, API routes, required environment variables and the op schema are unchanged.
