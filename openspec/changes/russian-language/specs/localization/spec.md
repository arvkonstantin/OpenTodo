# Spec Delta

## MODIFIED Requirements

### Requirement: Shipped languages and their catalogs
English, German, French, Spanish, Italian, Portuguese, Simplified Chinese and Russian SHALL be maintained in the repository and SHALL be complete in every release. A language SHALL be offered in Settings only when its catalog is complete. Shipped catalogs SHALL be written and reviewed by people; machine translation SHALL NOT be used for a shipped catalog. A community language follows the same rule.

#### Scenario: Complete languages offered
- **WHEN** the user opens the Language panel
- **THEN** the eight shipped languages are listed, each in its own name

#### Scenario: Work-in-progress catalog
- **WHEN** a catalog in the repository is incomplete
- **THEN** it is not offered in Settings and the release check reports its missing messages

## ADDED Requirements

### Requirement: Russian across client surfaces
Russian SHALL be offered as `Русский` with code `ru`, selected for Russian browser tags such as `ru-RU` under System and through the existing personal language preference. Russian SHALL cover interface messages including accessible names, service-worker notification actions, browser-extension messages, and Android and iOS shell strings. Counts SHALL use Russian plural categories, including one, few, many and other where appropriate. User content SHALL stay as written. Language changes SHALL work offline with the installed application, use the existing preference operations and converge by last-writer-wins across the account's devices.

#### Scenario: Russian browser
- **WHEN** a browser prefers `ru-RU` and the account has the System language preference
- **THEN** the interface is in Russian, the language choice names `Русский`, and the document language is `ru`

#### Scenario: Russian counts
- **WHEN** the Russian interface displays task counts of 1, 2, 5, 11, 21, 22 and 25
- **THEN** the noun forms are respectively `задача`, `задачи`, `задач`, `задач`, `задача`, `задачи` and `задач`

#### Scenario: Switch offline and sync
- **WHEN** the installed app switches to Russian offline and reconnects later
- **THEN** its Russian catalog works immediately, the preference reaches the account's other devices on sync, and concurrent language choices follow the existing last-writer-wins rule

#### Scenario: Notification and shell text
- **WHEN** Russian is selected by a surface's existing language resolution and a reminder or a shell setup screen is shown
- **THEN** its application-provided text and accessible names are Russian and the task title or server address is unchanged

#### Scenario: Long phone navigation labels
- **WHEN** Russian navigation labels do not fit their existing phone tab buttons, including with enlarged text
- **THEN** each visible label is truncated within its own button without overlapping adjacent tabs, while its accessible name retains the full translated label

#### Scenario: Actor names and Russian activity grammar
- **WHEN** a Russian task history, project history, shared activity feed or grouped run combines an arbitrary actor name with an event description
- **THEN** the full line uses the localized actor/action message, preserving the actor name and allowing Russian punctuation and neutral event wording without assuming the actor’s gender

#### Scenario: Scheduling action and schedule noun
- **WHEN** the Russian interface offers a task-scheduling action and displays a backup schedule or a task’s schedule facts
- **THEN** action controls use `Запланировать` while the noun label uses `Расписание`, through separate translation contexts with the English visible text unchanged

#### Scenario: Start date and starting an action
- **WHEN** Russian task/date fields display Start and the daily-check prompt offers a Start button
- **THEN** the field label uses `Начало` and the action uses `Начать`, through separate translation contexts with the English visible text unchanged

#### Scenario: Picking from different list types
- **WHEN** Russian displays a routine Pick card for the first, last, first N, last N, all or matching items of a task or line list, including quantities 1, 2 and 21
- **THEN** its interactive sentence and accessible name use the same grammatical, mode-aware wording, retain list-reference names unchanged and keep the existing English display, editable inputs and stored step schema
