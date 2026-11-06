import { describe, expect, it } from 'vitest'
import { parseEdit } from './edit'
import { parse } from './parse'
import { parseReminderAt, resolveReminder } from './reminders'
import { tokenize } from './tokenize'
import { activeLanguages, normalize, vocabulariesFor } from './vocab'

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

  it('resolves a current-time reminder independently of a future task due', () => {
    const result = parse('Встреча завтра в 15:00 напомни через 10 минут', ctx)
    expect(resolveReminder(result.reminders[0], result.due, ctx.today, ctx.now)).toEqual({ kind: 'relative', offsetMinutes: -1730 })
    expect(resolveReminder(result.reminders[0], null, ctx.today, ctx.now)).toEqual({ kind: 'absolute', date: '2026-11-06', time: '10:10' })
  })

  it('keeps the explicit next-day moment when an undated reminder crosses midnight', () => {
    const result = parse('Позвонить напомни через 10 минут', { ...ctx, now: '23:55' })
    expect(resolveReminder(result.reminders[0], null, ctx.today, '23:55')).toEqual({ kind: 'absolute', date: '2026-11-07', time: '00:05' })
  })

  it('keeps the complete current-time reminder literal when its chip is dismissed', () => {
    const line = 'Позвонить напомни через 10 минут'
    const first = parse(line, ctx)
    const chip = first.chips.find((c) => c.kind === 'reminder')!
    expect(chip.spans).toEqual([{ start: line.indexOf('напомни'), end: line.length, text: 'напомни через 10 минут' }])
    expect(parse(line, { ...ctx, dismissed: chip.spans })).toMatchObject({ title: line, due: null, reminders: [] })
  })

  it.each(['через 0 минут', 'через 0 часов', 'через 0 дней', 'через 1000 недель', 'через 999999999999999999999min', 'через месяц', 'через год', 'через "10 минут"'])('rejects an unsupported or nonpositive reminder body: %s', (phrase) => {
    const dateCtx = { ...ctx, vocabs: vocabulariesFor(ctx.languages), dayMonthOrder: 'DM' as const }
    expect(parseReminderAt(tokenize(phrase).tokens, 0, dateCtx)).toBeNull()
  })

  it('requires a current clock for a current-time reminder', () => {
    const dateCtx = { today: ctx.today, vocabs: vocabulariesFor(ctx.languages), dayMonthOrder: 'DM' as const }
    expect(parseReminderAt(tokenize('через час').tokens, 0, dateCtx)).toBeNull()
  })

  it('keeps the original night phrase when its scheduling chip is dismissed', () => {
    const line = '📝 Позвонить завтра в 11 часов ночи'
    const first = parse(line, ctx)
    const chip = first.chips.find((c) => c.kind === 'due')!
    expect(chip.spans.map((s) => s.text)).toEqual(['завтра', 'в 11 часов ночи'])
    expect(parse(line, { ...ctx, dismissed: chip.spans })).toMatchObject({ title: line, due: null, chips: [] })
  })

  it('keeps a dismissed night reminder literal, including its own date', () => {
    const line = 'Позвонить напомни завтра в 11 ночи'
    const first = parse(line, ctx)
    const chip = first.chips.find((c) => c.kind === 'reminder')!
    expect(parse(line, { ...ctx, dismissed: chip.spans })).toMatchObject({ title: line, due: null, reminders: [] })
  })

  it('updates the result as an ordinary clock gains a complete night qualifier', () => {
    expect(parse('Позвонить в 8', ctx)).toMatchObject({ title: 'Позвонить', due: { time: '08:00' } })
    expect(parse('Позвонить в 8 ночи', ctx)).toMatchObject({ title: 'Позвонить в 8 ночи', due: null })
    expect(parse('Позвонить в 11:30', ctx)).toMatchObject({ title: 'Позвонить', due: { time: '11:30' } })
    expect(parse('Позвонить в 11:30 ночи', ctx)).toMatchObject({ title: 'Позвонить', due: { time: '23:30' } })
  })

  it.each(['в 7 ночи', 'завтра в 8 часов ночи', 'в 9:15 ночи завтра', '8.11 ночи'])('rejects an unsupported night reminder body without a partial clock: %s', (phrase) => {
    const dateCtx = { ...ctx, vocabs: vocabulariesFor(ctx.languages), dayMonthOrder: 'DM' as const }
    expect(parseReminderAt(tokenize(phrase).tokens, 0, dateCtx)).toBeNull()
  })

  it('keeps the existing next-day resolution for a dateless midnight reminder', () => {
    const result = parse('Позвонить напомни в 12 ночи', ctx)
    expect(resolveReminder(result.reminders[0], null, ctx.today, ctx.now)).toEqual({ kind: 'absolute', date: '2026-11-07', time: '00:00' })
  })
})
