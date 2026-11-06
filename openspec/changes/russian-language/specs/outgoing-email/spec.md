# Spec Delta

## MODIFIED Requirements

### Requirement: Message language
A message SHALL be written in the first language of the `Accept-Language` header of the request that caused it that the server supports (`en`, `de`, `fr`, `es`, `it`, `pt`, `zh`, `ru`), else in English. Subject and body SHALL be translated together; the link and the server origin are not translated. For a verification message the language SHALL be the one captured at registration, so a resend uses the same language.

#### Scenario: German browser
- **WHEN** a visitor whose browser sends `Accept-Language: de-CH, fr;q=0.8` requests a password reset
- **THEN** subject and body are in German and the link is unchanged

#### Scenario: Unsupported language
- **WHEN** the header is `ja` or absent
- **THEN** the message is in English

#### Scenario: Invitation language follows the creator
- **WHEN** an admin whose browser sends `Accept-Language: it` sends an invitation
- **THEN** the invitee receives an Italian message

## ADDED Requirements

### Requirement: Russian message wording
Every existing outgoing email kind, including the weekly digest, SHALL have Russian subject and body wording when Russian is selected by its existing language rules. Expiry durations, roles, digest date labels and counts SHALL be grammatically correct in Russian. Names, addresses, links and permitted user content SHALL remain unchanged. No email SHALL be sent offline or without the configured SMTP service; existing encryption restrictions, queues and retries SHALL apply without changing synchronized user data.

#### Scenario: Russian password reset
- **WHEN** a visitor requests a password reset with `Accept-Language: ru-RU, en;q=0.8` on a configured server
- **THEN** the existing reset message has Russian subject, instructions and expiry wording, with its action link unchanged

#### Scenario: Russian weekly digest
- **WHEN** a user with the Russian digest language receives the scheduled digest
- **THEN** its headings and date labels are Russian, while task and project names are inserted as written under the existing content policy
