# Tasks

## 1. Translation draft and review material

- [x] 1.1 Prepare the complete Russian web catalog in reviewable batches using the D2 glossary, retaining all placeholders/tags and adding appropriate plural forms; verify `checkCatalog('ru', ...)` from `scripts/i18n-check.mjs` against the current source inventory, and deliver an English/Russian review list with source locations and counts. Keep draft status explicit until contributor review.
- [x] 1.2 Translate service-worker notification texts, extension messages, Android strings and the iOS String Catalog; verify notification/catalog parity tests, `make test-extension`, `node --test scripts/i18n/shell.test.mjs`, XML parsing and printf placeholders, and include each surface in the review list. Record Android/iOS build results if their toolchains are available.
- [x] 1.3 Integrate `ru` / `Русский` in language selection and relevant existing client allow-lists for the local review build; add locale-selection, Russian plural/count and date-formatting tests in the existing i18n/date suites, verify `cd web && npm test && npm run build`, and update `docs/language.md`, `docs/translating.md` and `docs/contributing/add-a-language.md` with Russian coverage and glossary. Do not claim human approval or submit the catalog before section 6.

## 2. Shared Russian quick-add grammar

- [x] 2.1 Add Russian vocabulary, inflected aliases and language registrations following D3, and regenerate the Go vocabulary with `cd web && npm run nlp:vocab`; add shared vectors for dates, numeric DM dates, weekdays, months/year rollover, times, repeat intervals/lists/ordinals/tails, deadline/by semantics, someday, day parts and all sector phrases with gates on/off. Verify `cd web && npm test -- nlp` and `go test ./internal/quickadd/...`.
- [x] 2.2 Mirror Cyrillic normalization in Go and add prefix reminder offsets plus language-scoped ordinal suffix handling to both parsers where needed; add shared cases for `напомни за 30 минут`, `напомни через 10 минут`, precomposed/decomposed `й` and `ё`, mixed case and Cyrillic/emoji spans. Verify existing-language regressions with `cd web && npm test -- nlp` and `go test ./internal/quickadd/...`.
- [x] 2.3 Cover literal/inactive Russian, quoted and dismissed chips, invalid dates, unsupported recurrence and ordinary text containing keyword fragments in parser/edit tests; verify the same Cyrillic project/label resolves in the browser and CLI with shared vectors, then run `go test ./...` and `cd web && npm test && npm run build`.
- [x] 2.4 Add Russian examples and supported-form limits to `docs/quick-add.md`, preserving canonical English stored/filter syntax; verify documented examples against shared vectors, run `make docs` and `make docs-check`.

## 3. Russian spoken attributes

- [x] 3.1 Add and register Russian spoken-token tables with contextual 0–59 numbers and D5 cues; extend `voice/spoken.test.ts` and locale tests for project/label, priority, times, durations, relative periods and literal ordinary numbers. Verify `cd web && npm test && npm run build`; document the Russian trigger examples in `docs/quick-add.md` and regenerate docs.

## 4. Server-written Russian text

- [x] 4.1 Add all Russian email templates, language negotiation, role/expiry/digest tables and applicable digest-language validators; add Russian rendering tests and a digest golden with plural cases 1/2/5/11/21/22/25, preserving content restrictions. Verify `go test ./internal/mail/... ./internal/digest/...`, then `go test ./...`, and include rendered examples in the contributor's review material.
- [x] 4.2 Add Russian delivery-channel text and applicable language validators; cover every existing text key, named/generic messages, ciphertext content restrictions and test messages with `go test ./internal/channels/...` and `go test ./...`. Document Russian email/channel availability in the existing relevant user guides and regenerate/check docs.

## 5. Integrated verification and local handoff

- [x] 5.1 Add the Unreleased changelog entry, update the relevant roadmap localization coverage without changing the Dutch/tooling plan, and run `make docs`; verify `make docs-check`, `make openspec-validate`, `make i18n-check` and `make ui-lint`, with every Russian surface represented in the review checklist.
- [x] 5.2 Measure equivalent baseline/changed production web and Go builds, record actual gzip/embedded binary deltas in design.md, verify the Russian catalog is lazy, and run `make bundle-budget` with the explicitly documented budget. Verify no dependency, mandatory environment variable, API route or op-schema change was introduced.
- [ ] 5.3 Verify Russian Today, task details, quick add, repeat/reminder controls and language settings at 360/1280 px in light/dark and with 20/24 px text; deliver screenshots and a running local preview. Check offline language switching and capture/reload, then synchronization with the existing test harness; record exact checks run and any unavailable native-device validation.
- [x] 5.4 Run `TZ=UTC make check`, `go test ./...`, `cd web && npm test && npm run build` and `make bundle-budget`; fix attributable failures, rerun affected checks and deliver their results with local preview instructions. Do not mark this task done if required checks fail.

## 6. Contributor review before publication

- [ ] 6.1 The contributor reads and corrects every Russian entry across all surfaces, confirms glossary/plural choices, and manually tries the documented Russian quick-add examples in the local build. Review may use GitHub commit comments or follow-up edits on the contributor-authorized draft branch in their fork. Verify explicit contributor confirmation and record factual review coverage; no issue, PR, release or archive before this review. This task cannot be completed by the agent on the contributor's behalf.
- [ ] 6.2 Apply contributor corrections and refresh any source-message delta; verify `make i18n-check`, relevant grammar/rendering tests, `make docs`, `make openspec-validate`, `make bundle-budget` and `TZ=UTC make check` before declaring implementation and translation review finished. Creating an issue or PR remains a separate explicitly authorized action.
