import { describe, expect, it } from 'vitest'
import { parseEdit } from './edit'
import { parse } from './parse'
import { activeLanguages, normalize } from './vocab'

const ctx = { today: '2026-11-06', now: '10:00', languages: ['ru-RU'], projects: [] }

describe('Russian quick add', () => {
  it('uses regional Russian with English available', () => {
    expect(activeLanguages(['ru-RU', 'ru-BY'])).toEqual(['ru', 'en'])
    expect(parse('Купить хлеб tomorrow', ctx)).toMatchObject({ title: 'Купить хлеб', due: { date: '2026-11-07' } })
  })

  it('keeps original Cyrillic and emoji spans when a chip is dismissed', () => {
    const line = '📝 Отчёт в следующий понедельник'
    const first = parse(line, ctx)
    const chip = first.chips.find((c) => c.kind === 'due')!
    expect(chip.spans).toEqual([{ start: line.indexOf('в '), end: line.length, text: 'в следующий понедельник' }])
    expect(parse(line, { ...ctx, dismissed: chip.spans })).toMatchObject({ title: line, due: null, chips: [] })
    expect(normalize('СЛЕДУЮЩИЙ')).toBe(normalize('следующий'))
    expect(normalize('ОТЧЁТ')).toBe(normalize('отчёт'))
  })

  it('parses only the newly typed Russian phrase when editing', () => {
    expect(parseEdit('Обсудить завтра', 'Обсудить завтра в пятницу', ctx)).toMatchObject({
      title: 'Обсудить завтра', due: { date: '2026-11-13' },
    })
  })

  it('parks a someday task without changing the title', () => {
    expect(parse('Перебрать шкаф когда-нибудь', ctx)).toMatchObject({ title: 'Перебрать шкаф', start: 'someday', due: null })
    expect(parse('Перебрать "когда-нибудь" шкаф', ctx)).toMatchObject({ title: 'Перебрать когда-нибудь шкаф', start: null, due: null })
  })

  it.each([
    ['5 января', '2027-01-05'], ['5 февраля', '2027-02-05'], ['5 марта', '2027-03-05'],
    ['5 апреля', '2027-04-05'], ['5 мая', '2027-05-05'], ['5 июня', '2027-06-05'],
    ['5 июля', '2027-07-05'], ['5 августа', '2027-08-05'], ['5 сентября', '2027-09-05'],
    ['5 октября', '2027-10-05'], ['5 ноября', '2027-11-05'], ['5 декабря', '2026-12-05'],
  ])('understands the genitive month in %s and rolls over the year', (phrase, date) => {
    expect(parse(`Оплатить ${phrase}`, ctx)).toMatchObject({ title: 'Оплатить', due: { date } })
  })

  it.each(['2026-02-30', '32/13', '31 февраля', 'февраля 31', '32 декабря', 'декабря 32', '29 февраля 2027', 'напомни за', 'каждыйзавтра', 'пятничный отчёт'])('keeps invalid or non-keyword text literal: %s', (phrase) => {
    const line = `Проверить ${phrase}`
    expect(parse(line, ctx)).toMatchObject({ title: line, due: null, recurrence: null, reminders: [] })
  })
})
