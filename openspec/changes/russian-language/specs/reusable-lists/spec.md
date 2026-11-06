# Spec Delta

## MODIFIED Requirements

### Requirement: Aisle sections
While a reusable list has "Sort into aisles" on, the user SHALL be offered to add the nine standard aisle sections (produce, bakery, dairy and eggs, meat and fish, frozen, pantry, drinks, household, personal care), named in the app's language and each carrying its fixed aisle key. A new item added from quick add without a section token SHALL be placed in the first live section of the project whose aisle key matches a word or a sequence of up to four words of its title (the longest wins) in the bundled vocabulary of the app's language, falling back to English. Without a matching section or word it SHALL go where quick add places it today. A reopened item SHALL NOT be moved. The vocabulary SHALL be bundled and evaluated on the device, with no network request. Russian product names and reviewed aliases SHALL be bundled alongside English, German, French, Spanish, Italian and Portuguese, covering the same nine aisle keys and preserving original task titles. Russian placement SHALL work offline, write only existing task operations and converge through normal synchronization; a missing vocabulary SHALL retain the existing English fallback without claiming native product-name support.

#### Scenario: Yoghurt to Dairy
- **WHEN** a reusable list with aisle sections and an app in English receives the new item "Greek yoghurt"
- **THEN** the item is created in the section keyed `dairy`, even if it was renamed "Fridge"

#### Scenario: Other language, same list
- **WHEN** a member using the app in Italian adds "latte" to the same list
- **THEN** the item lands in the section keyed `dairy`, whatever its name

#### Scenario: Learned placement
- **WHEN** the user moves "Oat drink" from Drinks to Dairy, ticks it off and later adds "Oat drink" again
- **THEN** it is reopened in Dairy

#### Scenario: Unknown word
- **WHEN** the user adds "Birthday candles" and no vocabulary entry matches
- **THEN** the item is created in the current section or in "No section"

#### Scenario: Russian product in a renamed aisle
- **WHEN** a member using Russian adds `Молоко` offline to a reusable list whose section keyed `dairy` was renamed `Холодильник`
- **THEN** the new task keeps the title `Молоко`, is placed in that keyed section through existing operations, and the assignment converges on other devices after sync

#### Scenario: Longer Russian product phrase
- **WHEN** Russian aisle matching receives `замороженные ягоды` with both a frozen phrase and a produce word in the dictionary
- **THEN** the longest phrase places it in `frozen`, while `ягоды` alone still matches `produce`

#### Scenario: Russian lookup keeps English fallback
- **WHEN** Russian is selected and a new item has an English title such as `Milk`
- **THEN** the existing English fallback can match `dairy`, and explicit section choices and reopened-item placement remain unchanged
