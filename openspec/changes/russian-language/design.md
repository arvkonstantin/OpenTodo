# Design

## Context

See proposal.md for motivation and scope. The baseline is upstream `main` at `4d9321707733d6826a6fd292a9ed78c3b0103856` (2026-11-06). The fork's `main` was fast-forwarded to that commit; work lives on `feat/russian-language`.

The web app has seven interface languages and 7,167 source messages. The quick-add and spoken-token grammars have six bundled languages. `docs/contributing/add-a-language.md` and `docs/translating.md` enumerate translation surfaces and require a human speaker's review; `more-languages` is a separate, still-planned Dutch/tooling change. The iOS String Catalog has 30 messages, and the Android base catalog has 49. The existing Inter font is a Latin subset with a Unicode range; Cyrillic therefore uses the existing local system-font fallback.

## Goals / Non-Goals

Goals: reuse the existing translation and parsing architecture, make Russian inflections explicit, keep browser and CLI behavior identical, and deliver a local build plus an auditable translation review checklist.

Non-goals: a morphological library, a translation service, a generalized language scaffolder, any new recognition service or permissions, or changes to canonical filter/recurrence languages. Do not fold in the contributor-kit scope of `more-languages`. The contributor explicitly excludes fixing the shared enlarged-text task-row defect: it also reproduces with English and all other shipped interface languages. Keep its screenshot and comparison evidence for reference; it does not block Russian localization verification.

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

### D2a. Exhaustive second language review

Pick cards use contextual translations for their sentence and mode words, shared by the interactive words and plain accessible sentence. Russian names a generic list element for singular/all/matching selections and expresses counted selections as a number taken from the start or end, without a noun that would require agreement with that number. The referenced list name is quoted without inflection, so both task and line outputs remain grammatical. Other catalogs receive the same contexts with singular/count-aware mode wording where needed; visible English, inputs, mode switching and stored steps remain unchanged. This adds no UI element, dependency or eager bundle payload.

The contributor requested a fresh agent review after finding literal Russian wording. Compare each of the 7,167 web messages with all seven other catalogs: English source, German, French, Spanish, Italian, Portuguese and Simplified Chinese. Do not sample or compare other languages only for doubtful entries. Consult referenced application code to resolve actual product behavior; other translations are supporting evidence and can themselves be wrong. Review all Russian non-web surfaces against every counterpart available for that surface, recording any absent language rather than inventing a comparison.

Freeze the input inventory, assign disjoint ranges to reviewers, and retain a coverage ledger listing every reviewed key, comparison languages and relevant input/output hashes. Review grammar, meaning, natural phrasing, terminology, count agreement, interpolated names and permission/security promises. Preserve ICU selectors, placeholders, rich-text tags, technical literals and the separate due-date/deadline semantics. Check accepted NLP aliases separately from recommended example wording. The approved extension corrections remove literal “other addresses” wording, clarify how to name and store a token, and refer to permission to access the server.

Integrate proposed corrections centrally to avoid conflicting catalog writes. Recheck meaning-sensitive changes against source context and verify Russian settings paths match actual labels. Update all matching surfaces, especially push/catalog copies and rendered server examples. Run required repository gates and production builds, remeasure changed asset sizes and refresh contributor-review material. Agent coverage is not native-speaker approval; tasks 6.1/6.2 stay pending until the contributor completes review. The translation review remains editorial except for the separately confirmed current-time reminder correction below; it must not silently add other parser behavior or repair the excluded shared row-layout defect. Where whole-phrase grammar is currently bypassed, route activity lines through the existing `{actor} {action}` message (rich values preserve styling), using a colon in Russian. Separate the existing ambiguous Schedule action from its noun, and the daily-check Start action from start-date field labels, through explicit action translation contexts in every catalog; preserve visible English wording and existing behavior.

### D3. Russian vocabulary in the existing grammar

Add `web/src/lib/nlp/vocab/ru.ts`; register it in `VOCABULARIES`, applicable language-name selectors and Go `Bundled`; regenerate `internal/quickadd/vocab.json` using `npm run nlp:vocab`. Use explicit aliases rather than stemming: month nominative/genitive/prepositional forms, weekday nominative/accusative/dative forms where their prepositions require them, unit inflections and ordinal genders. Do not introduce unbounded word matching.

Acceptance examples, with `today=2026-11-06`, `now=10:00` and Monday-start weeks:

| Input after a title | Meaning |
| --- | --- |
| `сегодня`, `завтра`, `послезавтра` | Nov 6, 7, 8 |
| `5 декабря`, `5 декабря 2026` | Dec 5, 2026 |
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
| `напомни за 30 минут` | Reminder offset -30 |
| `напомни через 10 минут` | Reminder moment ten minutes after the parser’s current date/time, independently of the task’s due date |
| `напомни завтра в 9:00` | Existing reminder moment |
| `крайний срок 9 ноября`, `дедлайн 9 ноября`, `к пятнице` | Deadline without a following time; `к` follows the existing by-word time exception |
| `когда-нибудь`, `сегодня вечером` | Someday / Evening |
| `завтра утром`, `завтра днём` | Day parts, only when enabled |
| `на этой неделе`, `в этом месяце` | Time sector when enabled |

For Russian sector phrases use natural `на этой неделе`, `в этом месяце`, `в этом квартале`, `в следующем квартале` and `когда-нибудь`-qualified next week/month forms. Cover all six existing sector values and switches on/off. Preserve the distinction between next-period dates and indefinite sectors. Use unambiguous monthly ordinal forms with `месяца`; do not classify every ordinal weekday as an alternating-week rule.

Extend the two reminder parsers to recognize a before/after offset cue before a count as well as the existing suffix order. Use the existing per-language cue lists; test Russian prefix cases and existing suffix cases, and keep incomplete cues literal. Add Russian numeric ordinal suffixes only through matching language-scoped logic mirrored in Go, with regression vectors for all existing suffixes. All supported examples become shared vectors; add negative examples for invalid dates, unsupported hourly recurrence, quoted words, inactive Russian, ordinary words containing keyword fragments and dismissed chips.

### D3a. Contributor-approved current-time reminders

The contributor explicitly chose the natural Russian meaning of `напомни через 10 минут`: count from the parser’s current `today`/`now`, not from the task’s due time. The same existing duration units (minutes, hours, days and weeks) follow this rule after the Russian relative-time cue; month/year duration semantics are not added. Resolve the phrase to the existing `moment` representation with an explicit floating date and time, including date rollover. A future task due date must not move the requested reminder into the future task day. Existing reminder resolution may convert the explicit moment to an offset from a dated task, and produces an absolute reminder for an undated task; synchronized fields retain those existing rules and require no new op or storage schema.

Retain `за` and suffix-before forms for reminders before the due time, and explicit `после`/`позже` forms for after-due offsets. Match the relative-from-now form before it can be consumed as an after-due offset. Keep quoted/disabled/inactive Russian literal, reject incomplete/nonpositive/out-of-range amounts safely, and preserve all existing non-Russian grammar. Pin the current-time semantics, midnight rollover, future/no due and invalid/literal cases in shared browser/Go vectors and relevant reminder resolution tests. This is a user-approved correction of previously misleading Russian grammar, not a new interpretation for other languages.

### D3b. Contributor-approved Russian night qualifiers

The contributor explicitly approved this conservative rule for `ночи` after a clock hour: 10/11 mean 22/23, 12 means 00, and 1–6 mean 01–06. The ambiguous 7–9 forms stay literal; recommend a 24-hour time instead. This is a product parsing rule chosen by the contributor, not a claim that colloquial Russian forbids other usages. Do not automatically interpret every `ночи` qualifier as AM. Handle bare/colon clock forms and optional hour words consistently in browser and Go, preserving other meridiem and non-Russian behavior. Invalid/unsupported night phrases must not leak a partial due/time match through the general scanner or a reminder body; an independent valid date/time outside the protected phrase remains eligible. No public schema or operation fields change.

### D3c. Contributor-approved Russian grocery vocabulary

The contributor approved adding Russian product-name recognition after the audit found translated aisle labels but an English-only word fallback. Reuse the existing lazy device-side aisle matcher, fixed nine category keys, longest whole-phrase matching and English fallback. Add a reviewed Russian JSON dictionary comparable to the six existing dictionaries, about 200–250 explicit product aliases with useful inflections and phrases. Keys follow the existing title normalization (including ё→е and й→и), not a new morphology algorithm; original task titles remain unchanged. Compare every Russian alias/category against the six available English/German/French/Spanish/Italian/Portuguese dictionaries; Chinese has no aisle vocabulary and is not added by this change.

No runtime classifier, Go dictionary, data schema, permission, service or dependency is added. Existing operation paths write only normal section assignments; no explicit section/reopened-item behavior changes. Validate normalization, all nine categories, longest-phrase precedence (such as frozen berries), English fallback, negative whole-word matching and an actual Russian quick-add placement into a renamed keyed section. Clarify actual supported grocery languages in user/contributor docs instead of implying Chinese product recognition. Estimated extra lazy gzip is 1.5–2.5 KiB before measurement; retain lazy loading and measure its final asset and equivalent binary impact.

### D4. Cyrillic normalization and spans

JavaScript currently applies NFD and strips combining marks, so `й` folds to `и` and `ё` to `е`; Go strips explicit combining marks but does not decompose those precomposed letters. Preserve the current JavaScript behavior and add the corresponding two mappings to Go's existing fold table, without a dependency. Test precomposed/decomposed Cyrillic, mixed case and Cyrillic project/label names. Retain original text for titles, chips and dismissal. Verify spans with Cyrillic text and an emoji before the date, respecting each existing API's span convention; do not globally convert existing public offsets.

Alternative: preserve `й` in JavaScript. Rejected because it would alter the established normalization contract for all entity matching and require a wider compatibility change.

### D5. Spoken tokens

Add `web/src/lib/voice/spoken/ru.ts` using `numberWords` for 0–59, including gendered one/two forms. Register Russian in `SPOKEN_TABLES`. Reuse quick-add time/unit cues; cover `проект`, `метка`, `хэштег`/`хештег`, `приоритет`, spoken time, duration and `через три дня`. Ordinary `Купить два яблока` stays literal. The recognition engine uses its existing locale selection/fallback and privacy behavior. No microphone is activated during automated verification.

### D6. Data model, footprint and interface budget

No database migration, op schema, content-field list, API, authentication policy or background job changes. Preference changes and task writes continue through the existing operations. Existing encrypted-account and delivery-content checks stay in force.

No new dependency or font asset. The final Russian catalog adds 216,328 bytes gzip (827,441 bytes raw) as a lazy precached asset. The vocabulary, spoken table and language registries add a few KB gzip depending on the existing chunk graph. Mail templates/tables and the embedded web assets grow the binary/container by their measured asset sizes; runtime memory adds the active catalog and bounded vocabulary maps, with no new worker or service. Compare equivalent before/after production web and Go builds, report actual gzip/binary deltas here, and run `make bundle-budget`. The baseline budget is 137 KB entry / 680 KB initial; any deliberate adjustment must be supported by measurements and recorded here and in `web/bundle-budget.json`.

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

The exhaustive-audit production build measures 139,679 bytes gzip for the entry (136.4 KiB) and 698,245 bytes for its static closure (681.9 KiB), compared with 136.4 / 679.4 KiB at the upstream baseline. The approximately 2.5 KiB increase includes Russian synchronous quick-add vocabulary, language registrations and the contributor-approved reminder/night handling. The parser is ready for the first keystroke and title edit through its existing synchronous API. The initial budget remains the explicitly approved 682 KiB (raised from the upstream 680 KiB during the initial Russian implementation); the entry budget remains 137 KiB. The final build passes both budgets without another increase.

The Russian interface catalog is a lazy asset outside the static import closure: 827,441 bytes raw / 216,328 bytes gzip at level 9. The new Russian grocery dictionary is a separate lazy asset: 6,938 bytes raw / 1,859 bytes gzip at level 9. Both are precached for offline use. The grocery source JSON has 250 normalized aliases (251 original forms), 8,292 bytes raw / 1,794 bytes gzip; source-file sizes are distinct from the built chunk. No dependency or font asset was added.

Equivalent embedded server (`CGO_ENABLED=0`, Go 1.27.3 linux/amd64, `-trimpath -s -w -X main.version=russian-review`): 41,799,840 bytes versus 40,902,816 bytes at upstream baseline, a delta of 897,024 bytes (2.193%). [Portable measurements](review/audit-build-metrics.json) record exact assets, hashes and flags; detailed per-asset measurements are also kept locally in `/tmp/opentodo-russian-audit-final-metrics.json`. Both Russian assets are absent from the entry’s static import closure. Dependency manifests/lockfiles, API routes, required environment variables and the op schema are unchanged.
