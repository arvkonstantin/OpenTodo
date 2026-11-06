package quickadd

import (
	"fmt"
	"time"
)

// Reminder is a reminder as typed: an offset from the due (negative = before)
// or a floating wall-clock moment.
type Reminder struct {
	Kind    string  `json:"kind"` // offset | moment
	Minutes int     `json:"minutes,omitempty"`
	Date    *string `json:"date,omitempty"`
	Time    string  `json:"time,omitempty"`
}

type remMatch struct {
	reminder Reminder
	from, to int
}

var unitMinutes = []struct {
	pick    func(*Vocab) []string
	minutes int
}{
	{func(v *Vocab) []string { return v.Units.Minute }, 1},
	{func(v *Vocab) []string { return v.Units.Hour }, 60},
	{func(v *Vocab) []string { return v.Units.Day }, 1440},
	{func(v *Vocab) []string { return v.Units.Week }, 10080},
}

func reminderAmountAt(tokens []token, i int, v *Vocab) (int, int) {
	if cn, count, ok := countAt(tokens, i, v); ok {
		for _, u := range unitMinutes {
			if n := phraseAt(tokens, i+cn, u.pick(v)); n > 0 {
				return cn + n, count * u.minutes
			}
		}
	}
	if t := tok(tokens, i); t != nil {
		if m := matchDuration(t.text, false); m > 0 {
			return 1, m
		}
	}
	return 0, 0
}

// Russian «через» counts from now. An explicit date keeps a later due independent.
func reminderFromNowAt(tokens []token, k, prefix int, v *Vocab, ctx *dateCtx) (Reminder, int, bool) {
	n, amount := reminderAmountAt(tokens, k+prefix, v)
	if ctx.now == "" || n == 0 || amount <= 0 {
		return Reminder{}, 0, false
	}
	total := minutesOf(ctx.now) + amount
	date := addDays(ctx.today, total/1440)
	at := fmt.Sprintf("%02d:%02d", total/60%24, total%60)
	return Reminder{Kind: "moment", Date: &date, Time: at}, k + prefix + n, true
}

func offsetAt(tokens []token, k int, ctx *dateCtx) (int, int, bool) {
	bestN, bestMin, found := 0, 0, false
	for _, v := range ctx.vocabs {
		preBefore, preAfter := 0, 0
		if v.Code == "ru" {
			preBefore = phraseAt(tokens, k, v.BeforeWords)
			if preBefore == 0 {
				preAfter = phraseAt(tokens, k, v.AfterWords)
			}
		}
		prefix := preBefore + preAfter
		j := k + prefix
		an, amount := reminderAmountAt(tokens, j, v)
		if an == 0 || amount <= 0 {
			continue
		}
		before, after := preBefore, preAfter
		if prefix == 0 {
			before = phraseAt(tokens, j+an, v.BeforeWords)
			if before == 0 {
				after = phraseAt(tokens, j+an, v.AfterWords)
			}
		}
		if before == 0 && after == 0 {
			continue
		}
		n := prefix + an
		if prefix == 0 {
			n += before + after
		}
		if !found || n > bestN {
			bestN, found = n, true
			bestMin = amount
			if before > 0 {
				bestMin = -amount
			}
		}
	}
	return bestN, bestMin, found
}

func momentAt(tokens []token, k int, ctx *dateCtx) (int, *string, string, bool) {
	if df := matchDateAt(tokens, k, ctx); df != nil {
		if t := matchTimeAt(tokens, df.to, ctx); t != nil {
			d := df.date
			return t.to - k, &d, t.time, true
		}
	}
	t := matchTimeAt(tokens, k, ctx)
	if t == nil {
		return 0, nil, "", false
	}
	if d := matchDateAt(tokens, t.to, ctx); d != nil {
		date := d.date
		return d.to - k, &date, t.time, true
	}
	if t.date != "" {
		date := t.date
		return t.to - k, &date, t.time, true
	}
	return t.to - k, nil, t.time, true
}

// parseReminderAt parses the reminder body (offset or moment) at token k.
func parseReminderAt(tokens []token, k int, ctx *dateCtx) (Reminder, int, bool) {
	for _, v := range ctx.vocabs {
		if v.Code == "ru" {
			if prefix := phraseAt(tokens, k, v.InWords); prefix > 0 {
				return reminderFromNowAt(tokens, k, prefix, v, ctx)
			}
		}
	}
	on, mins, okOff := offsetAt(tokens, k, ctx)
	mn, date, at, okMom := momentAt(tokens, k, ctx)
	if okOff && (!okMom || on >= mn) {
		return Reminder{Kind: "offset", Minutes: mins}, k + on, true
	}
	if okMom {
		return Reminder{Kind: "moment", Date: date, Time: at}, k + mn, true
	}
	return Reminder{}, 0, false
}

func matchReminderAt(tokens []token, i int, ctx *dateCtx) *remMatch {
	if tok(tokens, i) == nil {
		return nil
	}
	word := 0
	for _, v := range ctx.vocabs {
		word = max(word, phraseAt(tokens, i, v.RemindWords))
	}
	if word == 0 {
		return nil
	}
	r, to, ok := parseReminderAt(tokens, i+word, ctx)
	if !ok {
		return nil
	}
	return &remMatch{r, i, to}
}

func literalNightReminderEndAt(tokens []token, k int, ctx *dateCtx) int {
	night := russianNightAt(tokens, k, ctx)
	if night == nil {
		if date := matchDateAt(tokens, k, ctx); date != nil {
			night = russianNightAt(tokens, date.to, ctx)
		}
	}
	if night == nil || night.time != "" {
		return 0
	}
	if date := matchDateAt(tokens, night.to, ctx); date != nil {
		return date.to
	}
	return night.to
}

// literalReminderEndAt prevents an invalid Russian reminder body
// from being reinterpreted as a separate task date/time by the general scanner.
func literalReminderEndAt(tokens []token, i int, ctx *dateCtx) int {
	for _, v := range ctx.vocabs {
		if v.Code != "ru" {
			continue
		}
		word := 0
		for _, r := range ctx.vocabs {
			word = max(word, phraseAt(tokens, i, r.RemindWords))
		}
		if word == 0 {
			continue
		}
		k := i + word
		if to := literalNightReminderEndAt(tokens, k, ctx); to > 0 {
			return to
		}
		prefix := phraseAt(tokens, k, v.InWords)
		if prefix == 0 {
			continue
		}
		if _, _, ok := reminderFromNowAt(tokens, k, prefix, v, ctx); ok {
			continue
		}
		j := k + prefix
		cn, _, ok := countAt(tokens, j, v)
		if !ok {
			if t := tok(tokens, j); t != nil {
				body := t.norm
				if len(body) > 0 && (body[0] == '+' || body[0] == '-') {
					body = body[1:]
				}
				if len(body) > 0 && body[0] >= '0' && body[0] <= '9' {
					cn = 1
				}
			}
		}
		unitN := 0
		for _, words := range [][]string{v.Units.Minute, v.Units.Hour, v.Units.Day, v.Units.Week, v.Units.Month, v.Units.Year} {
			unitN = max(unitN, phraseAt(tokens, j+cn, words))
		}
		return j + cn + unitN
	}
	return 0
}

// Resolved is how a typed reminder is stored once the due is known.
type Resolved struct {
	Relative      bool
	OffsetMinutes int
	Date, Time    string // absolute moments, in the caller's zone
}

func dayNum(iso string) int {
	t, _ := time.Parse("2006-01-02", iso)
	return int(t.Unix() / 86400)
}

func minutesOf(hhmm string) int {
	t, _ := time.Parse("15:04", hhmm)
	return t.Hour()*60 + t.Minute()
}

// Resolve mirrors resolveReminder in web/src/lib/nlp/reminders.ts: an offset is
// relative; a moment is relative to the due time (09:00 for a date-only due)
// in wall-clock minutes; without a due a moment is absolute, today or tomorrow.
func Resolve(r Reminder, dueDate, dueTime, today, now string) Resolved {
	if r.Kind == "offset" {
		return Resolved{Relative: true, OffsetMinutes: r.Minutes}
	}
	if dueDate != "" {
		if dueTime == "" {
			dueTime = "09:00"
		}
		date := dueDate
		if r.Date != nil {
			date = *r.Date
		}
		return Resolved{Relative: true, OffsetMinutes: dayNum(date)*1440 + minutesOf(r.Time) - (dayNum(dueDate)*1440 + minutesOf(dueTime))}
	}
	if r.Date != nil {
		return Resolved{Date: *r.Date, Time: r.Time}
	}
	if minutesOf(r.Time) > minutesOf(now) {
		return Resolved{Date: today, Time: r.Time}
	}
	return Resolved{Date: addDays(today, 1), Time: r.Time}
}
