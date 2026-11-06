# Proposal

## Why

Russian speakers cannot use OpenTodo in their language or enter dates and repeats such as `5 декабря` and `каждый день`. Add Russian across the existing localization surfaces and the shared quick-add grammar, with a Russian-speaking contributor checking the wording before any issue or pull request is created.

## What Changes

- Add Russian (`ru`, `Русский`) to the interface, accessibility labels, date and count formatting, push notifications, browser extension, Android and iOS shell strings, outgoing email and delivery channel messages.
- Add Russian quick-add dates, times, repeats, reminders, deadlines, day parts, someday and time sector phrases. Cover natural inflections (`декабря`, `пятницу`, `каждую неделю`), English mixed with Russian, and spoken attribute cues. Keep stored repeat rules in canonical English.
- Keep browser and Go parsing identical, including Cyrillic normalization and original-text spans, through shared fixtures.
- Prepare a local review build and a checklist covering every translation surface and representative grammar examples. At the contributor’s request, commit the explicitly unreviewed draft and the known layout-problem screenshot to their fork for review through GitHub commit comments or follow-up edits. The contributor must read and correct all wording; an AI draft is not a reviewed translation. Do not open an issue or PR before the contributor explicitly finishes that review.
- Perform the contributor-requested exhaustive second language review: compare every Russian web message with all seven other interface languages and inspect code context where meaning is ambiguous. Review the remaining Russian surfaces against every available corresponding language. Record full coverage and actual corrections, while retaining the separate human-review requirement.
- Add the contributor-approved Russian grocery-aisle vocabulary to the existing device-side matching mechanism, with all nine current categories, English fallback, offline operation and no new classification service.
- No breaking changes, new dependencies, mandatory environment variables, API routes or op fields.

## Out of scope

Translation of the documentation, Help, CLI output, server error envelopes, saved filter syntax and stored recurrence syntax; arbitrary natural-language understanding; new recurrence semantics; Dutch and the contributor tooling/community-language tier in `more-languages`. English technical tokens such as `p1`, `#`, `@` and `+30m` remain portable.

## Roadmap

Extends the shipped `app-localization` and `nlp-quick-add` items, and advances the language coverage goal of `more-languages` without implementing its Dutch or contributor-tooling scope.

## Capabilities

### New Capabilities

None.

### Modified Capabilities

- `localization`: Russian as a complete, human-reviewed shipped language across existing client surfaces, with Russian plural forms.
- `natural-language-capture`: Russian grammar with deterministic offline behavior and identical Go/browser results.
- `voice-capture`: Russian spoken attribute cues and contextual numbers, using the existing recognition behavior.
- `outgoing-email`: Russian subject/body templates, roles, expiry and digest text.
- `notification-channels`: Russian generic and content-bearing delivery messages.
- `reusable-lists`: Russian product-name vocabulary for automatic aisle placement through the existing matcher.

## Impact

Changes the existing catalogs and language registries in `web/`, `extension/`, `mobile/`, `internal/mail`, `internal/channels`, the quick-add vocabulary and both parser implementations where required. Tests cover locale selection, plural categories, every surface, grammar, and Cyrillic spans. Documentation, generated help and an Unreleased changelog entry follow current repository conventions.

The existing `user_prefs.language` operation stores `ru`; older clients retain their existing unknown-language fallback. Tasks still store resolved values through operations and never reparse on sync. Catalogs stay lazy; measure first-load gzip and binary size before and after and retain the current bundle budget unless a separately justified change is necessary.
