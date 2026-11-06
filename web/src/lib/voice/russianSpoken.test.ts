import { describe, expect, it } from 'vitest'
import { normalizeSpoken } from './spoken'
import { parse } from '../nlp/parse'

const languages = ['ru-RU']
const ctx = { today: '2026-11-06', now: '10:00', languages, projects: [{ id: 'p-home', name: 'Дом' }] }

describe('Russian spoken attributes', () => {
  it('normalizes project, label and priority cues', () => {
    const line = normalizeSpoken('Купить краску проект Дом метка покупки приоритет два', languages, ctx)
    expect(line).toBe('Купить краску #Дом @покупки p2')
    expect(parse(line, ctx)).toMatchObject({ title: 'Купить краску', projectId: 'p-home', labels: ['покупки'], priority: 2 })
  })

  it.each([
    ['Позвонить завтра в девять тридцать', 'Позвонить завтра в 9:30'],
    ['Позвонить в девять ноль пять', 'Позвонить в 9:05'],
    ['Позвонить в три часа дня', 'Позвонить в 3pm'],
    ['Сделать через три дня', 'Сделать через 3 дня'],
    ['Сделать через две недели', 'Сделать через 2 недели'],
    ['Читать на тридцать минут', 'Читать +30m'],
    ['Читать один час и десять минут', 'Читать +70m'],
    ['Читать двадцать одну минуту', 'Читать +21m'],
    ['Купить два яблока', 'Купить два яблока'],
    ['Купить хэштег Дом', 'Купить #Дом'],
    ['Купить решётка Дом', 'Купить #Дом'],
  ])('normalizes %s with contextual numbers only', (line, expected) => {
    expect(normalizeSpoken(line, languages, ctx)).toBe(expected)
  })

  it('keeps Russian words literal when the language is inactive', () => {
    const line = 'Купить краску проект Дом приоритет два'
    expect(normalizeSpoken(line, ['en'], ctx)).toBe(line)
  })
})
