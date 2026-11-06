---
title: Add a language
summary: Every file a new interface language touches, from the catalog to emails, push texts, the extension and the Android app, and who reviews it.
audience: developer
section: Contributing
specs: []
screens: []
order: 5
---
# Add a language

The app ships in eight languages. A new one is a catalog plus a handful of small tables, and a speaker of the language who reviews it. **Machine translation is not accepted.** How messages are written is in [Translating OpenTodo](../translating.md).

## Before you start

Open an issue or a discussion naming the language and who will review it. Until it is listed, a catalog can sit in the repository unfinished: the check only prints how far it is ("pl: 412 of 3550").

## The files

| What | Where |
| --- | --- |
| The interface catalog | `web/src/locales/<code>.json`: every key of `web/src/locales/messages.json`, translated |
| The language list | `{ code, name }` in `SHIPPED`, `web/src/locales/index.ts`, with the name in the language itself |
| Stored language values | `LANGUAGES` in `web/src/lib/userPrefs.ts` and `LanguagePref` in `web/src/lib/types.ts`, so a synced choice is recognized |
| Browser tags that pick it | `web/src/lib/i18n/locale.ts`, only when the language has script or region variants (as Chinese does) |
| Notification texts of the service worker | the table in `web/public/push-sw.js` |
| The browser extension | `extension/_locales/<code>/messages.json` |
| The iOS app | `mobile/ios/Shared/Localizable.xcstrings` and the project’s `knownRegions` |
| The Android app | `mobile/android/app/src/main/res/values-<code>/strings.xml` |
| Emails | `internal/mail/templates/*.<code>.txt`, and `<code>` in `mail.Languages` |
| Delivery channel messages | the rows in `internal/channels/messages.go` |
| Product names for grocery aisles | `web/src/lib/aisles/<code>.json`, when adding product-name recognition for the language; see below |
| The docs | the glossary column in [Translating OpenTodo](../translating.md#glossary), and the list in [Language](../language.md) |

Quick add understands dates and priorities in a language only with a vocabulary in `web/src/lib/nlp/vocab/`. That is a separate contribution; without it, quick add keeps English.

Russian also has quick-add and spoken-token vocabularies. Keep its explicit case forms and Cyrillic normalization identical in the browser and Go parser; cover dates, repeats and reminders with shared vectors.

Grocery-aisle dictionaries are separate from both interface catalogs and quick-add date grammars. They currently cover English, German, French, Spanish, Italian, Portuguese and Russian; other interface languages use English product names only. A dictionary maps normalized whole words or phrases of up to four words to the existing nine aisle keys. Keep explicit useful forms and synonyms, compare category choices with the existing dictionaries, and have a speaker review every entry. The Russian dictionary has 250 normalized entries, including explicit case forms; title normalization folds `ё` to `е` and `й` to `и`, while stored item titles stay unchanged.

Each aisle dictionary loads lazily through the existing matcher. The Russian JSON is 8,292 bytes raw and 1,794 bytes gzip at level 9; these are source-file measurements, not the production chunk or binary delta. Measure the built asset and embedded binary as part of the change's footprint checks. No Go vocabulary or new classifier is needed: automatic aisle placement runs on the device. Cover all nine categories, phrase precedence, English fallback, negative whole-word matches and actual placement into a renamed keyed section. Run `cd web && npm test -- src/lib/aisles/aisles.test.ts src/app/quickAdd.reusable.test.tsx --no-fsModuleCache` after adding a new dictionary; disabling the transformed-module cache ensures its new lazy import is discovered.

## Check it

```sh
node scripts/i18n/extract.mjs --write   # after any change to English messages
make i18n-check                          # complete catalogs, placeholders, plurals, Settings paths
cd web && npm test -- i18n
go test ./internal/mail/... ./internal/channels/...
```

Use the plural categories of your language (`one`/`other`, Russian `one`/`few`/`many`/`other`, or only `other` for languages without plurals), keep every placeholder and tag, and use `’` rather than `'` in translations.

## In the pull request

- Name the reviewer: a speaker of the language who reads every entry, now and for later catalog changes.
- Add the language to `CHANGELOG.md` under Unreleased.
