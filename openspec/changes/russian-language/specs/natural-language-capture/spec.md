# Spec Delta

## MODIFIED Requirements

### Requirement: Multi-language vocabulary
Date, time, recurrence and reminder vocabulary SHALL be bundled for English, German, French, Italian, Spanish, Portuguese and Russian, available with no network. The active set SHALL be the interface language (OpenSpec `localization`) when not `system`, then the browser's bundled preferred languages, with English always active, overridable in Settings on the device. Attribute tokens (`#`, `@`, `p1`–`p4`, `+`) SHALL be language-independent. The stored rule SHALL always be the canonical English form.

#### Scenario: German phrase
- **WHEN** the browser language is `de` and the user enters `Zahnarzt nächsten Mittwoch um 15 Uhr`
- **THEN** the task is due next Wednesday at `15:00` with title `Zahnarzt`

#### Scenario: English always works
- **WHEN** the browser language is `it` and the user enters `Pay rent tomorrow`
- **THEN** the task is due tomorrow with title `Pay rent`

#### Scenario: Unsupported language falls back
- **WHEN** the browser language is one with no bundled vocabulary
- **THEN** English vocabulary is active, attribute tokens still work, and no error is shown

#### Scenario: Interface language comes first
- **WHEN** the browser language is `en-US`, the interface language is Spanish and the device has no language override
- **THEN** the active languages are Spanish then English, `Pagar alquiler mañana` is due tomorrow, numeric dates read day-month while the date format is `system`, and dictation listens in Spanish

#### Scenario: Device override still wins
- **WHEN** the interface language is Spanish and the device's quick-add languages are set to German only
- **THEN** the active languages are German then English

#### Scenario: Recurrence in each language
- **WHEN** the user enters `jeden Montag`, `tous les lundis`, `ogni lunedì`, `cada lunes` or `toda segunda-feira` after a title, with that language active
- **THEN** the rule is `every monday` in every case

#### Scenario: French interval with completion
- **WHEN** the browser language is `fr` and the user enters `Changer le filtre tous les 3 jours après achèvement`
- **THEN** the rule is `every 3 days after completion` and the title is `Changer le filtre`

#### Scenario: Reminder in each language
- **WHEN** the user enters `rappelle-moi 30 min avant`, `ricordamelo 30 minuti prima`, `recuérdame 30 minutos antes` or `lembre-me 30 minutos antes` after a timed task, with that language active
- **THEN** a relative reminder with offset −30 minutes is created in every case

## ADDED Requirements

### Requirement: Russian date and scheduling phrases
With Russian active, quick add SHALL understand Russian relative days, weekday and month names with their date inflections, relative periods, times, recurrence, reminders, deadlines, someday, day parts and time sectors. English SHALL remain available. Matching SHALL ignore case and consistently fold Cyrillic combining marks, `ё` and `й` in both parsers, while preserving unclaimed original text. Quotes, chip dismissal and disabling parsing SHALL keep their existing literal behavior. Unsupported phrases SHALL stay in the title. Browser and server/CLI results SHALL agree for the same line, date, clock, languages and preferences. Parsing SHALL require no network; created tasks and reminders SHALL use existing operations, store resolved fields and canonical English repeat rules, and converge without reparsing after synchronization.

#### Scenario: Month in the genitive
- **WHEN** Russian is active, today is 2026-11-06 and the user enters `Купить подарок 5 декабря`
- **THEN** the title is `Купить подарок` and the due date is 2026-12-05

#### Scenario: Relative days and time
- **WHEN** Russian is active, today is 2026-11-06 and the user enters `Позвонить завтра в 15:30`
- **THEN** the title is `Позвонить` and the due value is 2026-11-07 at 15:30

#### Scenario: Russian intervals
- **WHEN** Russian is active, today is 2026-11-06 and the user enters `Поливать каждые 3 дня после завершения`
- **THEN** the title is `Поливать`, the first occurrence is due today and the stored rule is `every 3 days after completion`

#### Scenario: Daily and weekly forms
- **WHEN** Russian is active and `Зарядка каждый день`, `Уборка каждую неделю` and `Отчёт каждую пятницу` are parsed
- **THEN** their rules are `every day`, `every week` and `every friday`, their first occurrences follow the existing recurrence rules, and only the scheduling phrases leave the titles

#### Scenario: Reminder before the due time
- **WHEN** Russian is active, today is 2026-11-06 and the user enters `Встреча завтра в 15:00 напомни за 30 минут`
- **THEN** the task is titled `Встреча`, due 2026-11-07 at 15:00, with a relative reminder of minus 30 minutes

#### Scenario: Deadline and evening
- **WHEN** Russian is active, today is 2026-11-06 and the user enters `Отчёт завтра вечером крайний срок 9 ноября`
- **THEN** the title is `Отчёт`, the task is due 2026-11-07 with Evening, and its deadline is 2026-11-09

#### Scenario: Feature-gated scheduling
- **WHEN** Russian is active, today is 2026-11-06 and `Разминка завтра утром` is parsed with Day parts on or `Бюджет на этой неделе` with Time sectors on and Monday weeks
- **THEN** the first task is due 2026-11-07 with Morning and the second has sector `2026-W45` without a due date; turning off each feature keeps its phrase literal according to the existing feature rules

#### Scenario: Cyrillic normalization and original spans
- **WHEN** Russian is active and the parsers receive `Сделать отчёт в следующий понедельник` with precomposed or decomposed Cyrillic letters
- **THEN** both resolve the same date and remove exactly the scheduling phrase, preserving `Сделать отчёт` in its original spelling

#### Scenario: Literal and inactive Russian
- **WHEN** parsing is disabled, Russian is inactive, a phrase is quoted, or its chip is dismissed
- **THEN** the affected Russian words stay as typed under the existing literal rules

#### Scenario: Mixed language and offline synchronization
- **WHEN** Russian is active and an offline device creates `Купить хлеб every day`, then reconnects
- **THEN** the task stores `every day` with the resolved first due date locally, and other devices converge on those fields without parsing the line again
