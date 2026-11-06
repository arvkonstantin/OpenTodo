# Spec Delta

## ADDED Requirements

### Requirement: Russian spoken attributes
With Russian active, spoken attribute normalization SHALL recognize Russian project, label, hashtag and priority cues, contextual numbers from zero to fifty-nine, times and durations using the existing token grammar. Numbers without an attribute, time or duration cue SHALL remain literal. Russian date phrases in the transcript SHALL be passed to quick add. Recognition availability, consent and fallback behavior SHALL stay as specified by the existing voice-capture requirements. Text normalization SHALL be deterministic and offline; it SHALL save no data until the user confirms capture, after which the existing task operations and convergence rules apply.

#### Scenario: Spoken priority and project
- **WHEN** Russian is active, an existing project is named `Дом`, and the transcript is `Купить краску проект Дом приоритет два`
- **THEN** the normalized line is `Купить краску #Дом p2`, the preview resolves that project and priority, and nothing is saved before confirmation

#### Scenario: Spoken time
- **WHEN** Russian is active and the transcript is `Позвонить завтра в девять тридцать`
- **THEN** quick add receives a time of 09:30 and resolves tomorrow using the same rules as typed Russian

#### Scenario: An ordinary number stays text
- **WHEN** Russian is active and the transcript is `Купить два яблока`
- **THEN** the text stays `Купить два яблока` and no attribute is inferred
