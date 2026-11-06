// @vitest-environment happy-dom
// Language-review regressions: whole-phrase agreement, exact-count branches and units.
import { afterEach, describe, expect, it } from 'vitest'
import { applyLanguagePreference, t } from '.'

const counts = [0, 1, 2, 5, 11, 14, 21, 22, 25, 101, 111, 1.5]
const planned = [
  'На завтра запланировано 0 задач', 'На завтра запланирована 1 задача',
  'На завтра запланированы 2 задачи', 'На завтра запланировано 5 задач',
  'На завтра запланировано 11 задач', 'На завтра запланировано 14 задач',
  'На завтра запланирована 21 задача', 'На завтра запланированы 22 задачи',
  'На завтра запланировано 25 задач', 'На завтра запланирована 101 задача',
  'На завтра запланировано 111 задач', 'На завтра запланировано 1,5 задачи',
]

afterEach(async () => { await applyLanguagePreference('en') })

describe('Russian review corrections', () => {
  it('agrees the full tomorrow sentence, including its verb, for every required count', async () => {
    await applyLanguagePreference('ru')
    const untimed = 'Tomorrow starts with {n, plural, one {one task} other {# tasks}}; the first is <task>{title}</task>.'
    const timed = 'Tomorrow starts with {n, plural, one {one task} other {# tasks}}; the first is <task>{title}</task> at {time}.'
    counts.forEach((n, index) => {
      expect(t(untimed, { n, title: 'Моя задача' })).toBe(`${planned[index]}; первая — Моя задача.`)
      expect(t(timed, { n, title: 'Моя задача', time: '15:00' })).toBe(`${planned[index]}; первая — Моя задача в 15:00.`)
    })
  })

  it('keeps counts 21 and 101 instead of treating Russian one as exactly one item', async () => {
    await applyLanguagePreference('ru')
    const source = 'Select a task to decide where it goes, or press {key} to process {count, plural, one {it} other {all #, one at a time}}.'
    expect(t(source, { key: 'Enter', count: 1 })).toContain('разобрать её')
    for (const count of [21, 101]) {
      const text = t(source, { key: 'Enter', count })
      expect(text).toContain(`${count} задачу`)
      expect(text).not.toContain('разобрать её')
    }
  })

  it('keeps the consecutive-failure count for 1, 21 and 101', async () => {
    await applyLanguagePreference('ru')
    const source = '{n, plural, one {The last send failed: {error}.} other {The last # sends failed: {error}.}}'
    for (const n of [1, 21, 101]) {
      expect(t(source, { n, error: 'timeout' })).toBe(`Неудачных отправок подряд: ${n}. Ошибка: timeout.`)
    }
  })

  it('labels a location radius in metres, not minutes', async () => {
    await applyLanguagePreference('ru')
    expect(t('{m} m', { m: 500 })).toBe('500 м')
  })
})
