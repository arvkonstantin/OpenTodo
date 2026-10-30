# Layout regression: completed tasks hidden in both builds

Review screenshots for OpenTodoApp/OpenTodo#624 and the `calendar-completed-tasks` OpenSpec change. The feature implementation is on `feat/calendar-completed-tasks` in this fork. Review PNGs are kept on this separate review-assets branch so the gallery does not add to the app bundle.

## Upstream default vs feature default: Show completed tasks is off in both

**This is a layout comparison, not a demonstration of enabling the feature.** Completed tasks are hidden in both columns. The same seeded demo fixture is used. Before: upstream `b0a069a0`; after: feature implementation `74be282e7032c70e57ef0a374a68dc5afa69e291`. Desktop: 1280 × 800; phone: 390 × 844 (device scale factor 2). Both themes. The moving now line and cursor position reflect capture time.

### Day

| View | Before | After |
| --- | --- | --- |
| Desktop, light | ![Before day, desktop, light](before/calendar-day-desktop-light.png) | ![After day, desktop, light](after/calendar-day-desktop-light.png) |
| Desktop, dark | ![Before day, desktop, dark](before/calendar-day-desktop-dark.png) | ![After day, desktop, dark](after/calendar-day-desktop-dark.png) |
| Phone, light | ![Before day, phone, light](before/calendar-day-phone-light.png) | ![After day, phone, light](after/calendar-day-phone-light.png) |
| Phone, dark | ![Before day, phone, dark](before/calendar-day-phone-dark.png) | ![After day, phone, dark](after/calendar-day-phone-dark.png) |

### Week

| View | Before | After |
| --- | --- | --- |
| Desktop, light | ![Before week, desktop, light](before/calendar-week-desktop-light.png) | ![After week, desktop, light](after/calendar-week-desktop-light.png) |
| Desktop, dark | ![Before week, desktop, dark](before/calendar-week-desktop-dark.png) | ![After week, desktop, dark](after/calendar-week-desktop-dark.png) |
| Phone, light | ![Before week, phone, light](before/calendar-week-phone-light.png) | ![After week, phone, light](after/calendar-week-phone-light.png) |
| Phone, dark | ![Before week, phone, dark](before/calendar-week-phone-dark.png) | ![After week, phone, dark](after/calendar-week-phone-dark.png) |

### Month

| View | Before | After |
| --- | --- | --- |
| Desktop, light | ![Before month, desktop, light](before/calendar-month-desktop-light.png) | ![After month, desktop, light](after/calendar-month-desktop-light.png) |
| Desktop, dark | ![Before month, desktop, dark](before/calendar-month-desktop-dark.png) | ![After month, desktop, dark](after/calendar-month-desktop-dark.png) |
| Phone, light | ![Before month, phone, light](before/calendar-month-phone-light.png) | ![After month, phone, light](after/calendar-month-phone-light.png) |
| Phone, dark | ![Before month, phone, dark](before/calendar-month-phone-dark.png) | ![After month, phone, dark](after/calendar-month-phone-dark.png) |

### Plan

| View | Before | After |
| --- | --- | --- |
| Desktop, light | ![Before plan, desktop, light](before/calendar-plan-desktop-light.png) | ![After plan, desktop, light](after/calendar-plan-desktop-light.png) |
| Desktop, dark | ![Before plan, desktop, dark](before/calendar-plan-desktop-dark.png) | ![After plan, desktop, dark](after/calendar-plan-desktop-dark.png) |
| Phone, light | ![Before plan, phone, light](before/calendar-plan-phone-light.png) | ![After plan, phone, light](after/calendar-plan-phone-light.png) |
| Phone, dark | ![Before plan, phone, dark](before/calendar-plan-phone-dark.png) | ![After plan, phone, dark](after/calendar-plan-phone-dark.png) |

## Show completed tasks enabled

A separate small fixture with two completed tasks (one timed and one all-day) and two open tasks. The remaining daily capacity stays **1h 45m / 8h**. On smaller screens the checked control lives in **More actions**.

### Day with completed tasks

| View | Screenshot |
| --- | --- |
| Desktop, light | ![day with completed tasks, desktop, light](completed/calendar-day-desktop-light.png) |
| Desktop, dark | ![day with completed tasks, desktop, dark](completed/calendar-day-desktop-dark.png) |
| Phone, light | ![day with completed tasks, phone, light](completed/calendar-day-phone-light.png) |
| Phone, dark | ![day with completed tasks, phone, dark](completed/calendar-day-phone-dark.png) |

### Week with completed tasks

| View | Screenshot |
| --- | --- |
| Desktop, light | ![week with completed tasks, desktop, light](completed/calendar-week-desktop-light.png) |
| Desktop, dark | ![week with completed tasks, desktop, dark](completed/calendar-week-desktop-dark.png) |
| Phone, light | ![week with completed tasks, phone, light](completed/calendar-week-phone-light.png) |
| Phone, dark | ![week with completed tasks, phone, dark](completed/calendar-week-phone-dark.png) |

### Month with completed tasks

| View | Screenshot |
| --- | --- |
| Desktop, light | ![month with completed tasks, desktop, light](completed/calendar-month-desktop-light.png) |
| Desktop, dark | ![month with completed tasks, desktop, dark](completed/calendar-month-desktop-dark.png) |
| Phone, light | ![month with completed tasks, phone, light](completed/calendar-month-phone-light.png) |
| Phone, dark | ![month with completed tasks, phone, dark](completed/calendar-month-phone-dark.png) |

### Plan with completed tasks

| View | Screenshot |
| --- | --- |
| Desktop, light | ![plan with completed tasks, desktop, light](completed/calendar-plan-desktop-light.png) |
| Desktop, dark | ![plan with completed tasks, desktop, dark](completed/calendar-plan-desktop-dark.png) |
| Phone, light | ![plan with completed tasks, phone, light](completed/calendar-plan-phone-light.png) |
| Phone, dark | ![plan with completed tasks, phone, dark](completed/calendar-plan-phone-dark.png) |
