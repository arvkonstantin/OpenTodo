# Calendar completed-task review

Feature comparison for [OpenTodoApp/OpenTodo#626](https://github.com/OpenTodoApp/OpenTodo/pull/626), captured from implementation commit `c53c1ea7ad8eec2a0fdbd528fc9c042245b43633`.

## Before: Show completed off → After: Show completed on

**Both columns use the same build and the same three tasks.** The only change is the Show completed setting. The two completed tasks exist in both states: they are hidden in Before and visible with struck-through titles in After.

- Open: **Review the launch checklist**, scheduled for 10:00, estimated at 1 hour.
- Completed: **Send the project update**, scheduled for 09:00, and **Book the train tickets**, with a date but no time.
- Remaining capacity stays **1h / 8h** in both states. Task data is not changed.
- Same account, date, viewport, theme and fixed clock within each pair. Desktop: 1280 × 800; phone: 390 × 844 with touch input and device scale factor 2.
- On phones, the setting is in **More actions**; the menu is closed for each screenshot.

There are **16 off/on pairs (32 screenshots)** below.

## Day

| View | Before — completed tasks hidden | After — completed tasks shown, struck through |
| --- | --- | --- |
| Desktop, light | ![Show completed off: day, desktop, light](comparison-before/calendar-day-desktop-light.png) | ![Show completed on: day, desktop, light](comparison-after/calendar-day-desktop-light.png) |
| Desktop, dark | ![Show completed off: day, desktop, dark](comparison-before/calendar-day-desktop-dark.png) | ![Show completed on: day, desktop, dark](comparison-after/calendar-day-desktop-dark.png) |
| Phone, light | ![Show completed off: day, phone, light](comparison-before/calendar-day-phone-light.png) | ![Show completed on: day, phone, light](comparison-after/calendar-day-phone-light.png) |
| Phone, dark | ![Show completed off: day, phone, dark](comparison-before/calendar-day-phone-dark.png) | ![Show completed on: day, phone, dark](comparison-after/calendar-day-phone-dark.png) |

## Week

| View | Before — completed tasks hidden | After — completed tasks shown, struck through |
| --- | --- | --- |
| Desktop, light | ![Show completed off: week, desktop, light](comparison-before/calendar-week-desktop-light.png) | ![Show completed on: week, desktop, light](comparison-after/calendar-week-desktop-light.png) |
| Desktop, dark | ![Show completed off: week, desktop, dark](comparison-before/calendar-week-desktop-dark.png) | ![Show completed on: week, desktop, dark](comparison-after/calendar-week-desktop-dark.png) |
| Phone, light | ![Show completed off: week, phone, light](comparison-before/calendar-week-phone-light.png) | ![Show completed on: week, phone, light](comparison-after/calendar-week-phone-light.png) |
| Phone, dark | ![Show completed off: week, phone, dark](comparison-before/calendar-week-phone-dark.png) | ![Show completed on: week, phone, dark](comparison-after/calendar-week-phone-dark.png) |

## Month

| View | Before — completed tasks hidden | After — completed tasks shown, struck through |
| --- | --- | --- |
| Desktop, light | ![Show completed off: month, desktop, light](comparison-before/calendar-month-desktop-light.png) | ![Show completed on: month, desktop, light](comparison-after/calendar-month-desktop-light.png) |
| Desktop, dark | ![Show completed off: month, desktop, dark](comparison-before/calendar-month-desktop-dark.png) | ![Show completed on: month, desktop, dark](comparison-after/calendar-month-desktop-dark.png) |
| Phone, light | ![Show completed off: month, phone, light](comparison-before/calendar-month-phone-light.png) | ![Show completed on: month, phone, light](comparison-after/calendar-month-phone-light.png) |
| Phone, dark | ![Show completed off: month, phone, dark](comparison-before/calendar-month-phone-dark.png) | ![Show completed on: month, phone, dark](comparison-after/calendar-month-phone-dark.png) |

## Plan

| View | Before — completed tasks hidden | After — completed tasks shown, struck through |
| --- | --- | --- |
| Desktop, light | ![Show completed off: plan, desktop, light](comparison-before/calendar-plan-desktop-light.png) | ![Show completed on: plan, desktop, light](comparison-after/calendar-plan-desktop-light.png) |
| Desktop, dark | ![Show completed off: plan, desktop, dark](comparison-before/calendar-plan-desktop-dark.png) | ![Show completed on: plan, desktop, dark](comparison-after/calendar-plan-desktop-dark.png) |
| Phone, light | ![Show completed off: plan, phone, light](comparison-before/calendar-plan-phone-light.png) | ![Show completed on: plan, phone, light](comparison-after/calendar-plan-phone-light.png) |
| Phone, dark | ![Show completed off: plan, phone, dark](comparison-before/calendar-plan-phone-dark.png) | ![Show completed on: plan, phone, dark](comparison-after/calendar-plan-phone-dark.png) |

## Separate layout regression screenshots

[Upstream default vs feature default](layout-comparison.md) preserves the original code-before/code-after comparison required by CONTRIBUTING.md. **Show completed is off in both builds** in those pairs, so they primarily demonstrate the new header control and preserved layout. Use the off/on pairs above to review the completed-task behavior.
