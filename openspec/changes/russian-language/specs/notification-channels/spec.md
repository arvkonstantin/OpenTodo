# Spec Delta

## ADDED Requirements

### Requirement: Russian delivery messages
A notification channel whose language resolves to Russian SHALL use Russian wording for every existing message kind, including reminders, assignments, mentions, task-page updates, rotation turns, routine notices and failures, and test messages. Names and task content SHALL remain as written when the existing content policy permits them; encrypted accounts SHALL receive only the existing generic form with Russian wording. Delivery SHALL still require a configured channel and a network connection, follow the existing durable retry behavior, and introduce no synchronized data changes.

#### Scenario: Generic encrypted reminder
- **WHEN** an encrypted account receives a reminder through a channel using Russian
- **THEN** the message uses Russian generic reminder wording and contains no task title or plaintext content

#### Scenario: Assignment with content
- **WHEN** a Russian channel permits titles and receives an assignment notification
- **THEN** its surrounding wording is Russian and the actor's name and task title remain as written

#### Scenario: Test a Russian channel
- **WHEN** the user sends an existing channel test with Russian selected
- **THEN** its title and body are in Russian and the delivery retains the existing failure and retry behavior
