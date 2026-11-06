package quickadd

import "testing"

func TestRussianCurrentTimeReminderResolution(t *testing.T) {
	ctx := Context{Today: "2026-11-06", Now: "10:00", Languages: []string{"ru"}}
	r, ok := ParseReminder("через 10 минут", ctx)
	if !ok {
		t.Fatal("current-time reminder not parsed")
	}
	if got := Resolve(r, "2026-11-07", "15:00", ctx.Today, ctx.Now); got != (Resolved{Relative: true, OffsetMinutes: -1730}) {
		t.Fatalf("future task due: %+v", got)
	}
	if got := Resolve(r, "", "", ctx.Today, ctx.Now); got != (Resolved{Date: "2026-11-06", Time: "10:10"}) {
		t.Fatalf("undated task: %+v", got)
	}
	ctx.Now = "23:55"
	r, ok = ParseReminder("через 10 минут", ctx)
	if !ok {
		t.Fatal("midnight reminder not parsed")
	}
	if got := Resolve(r, "", "", ctx.Today, ctx.Now); got != (Resolved{Date: "2026-11-07", Time: "00:05"}) {
		t.Fatalf("midnight: %+v", got)
	}
}

func TestRussianInvalidCurrentTimeReminder(t *testing.T) {
	ctx := Context{Today: "2026-11-06", Now: "10:00", Languages: []string{"ru"}}
	for _, phrase := range []string{"через 0 минут", "через 0 часов", "через 0 дней", "через 1000 недель", "через 999999999999999999999min", "через месяц", "через год", `через "10 минут"`} {
		t.Run(phrase, func(t *testing.T) {
			if r, ok := ParseReminder(phrase, ctx); ok {
				t.Fatalf("invalid reminder accepted: %+v", r)
			}
		})
	}
	ctx.Now = ""
	if r, ok := ParseReminder("через час", ctx); ok {
		t.Fatalf("missing clock: %+v", r)
	}
}

func TestRussianNightReminderBody(t *testing.T) {
	ctx := Context{Today: "2026-11-06", Now: "10:00", Languages: []string{"ru"}}
	r, ok := ParseReminder("завтра в 11 часов ночи", ctx)
	if !ok || r.Kind != "moment" || r.Date == nil || *r.Date != "2026-11-07" || r.Time != "23:00" {
		t.Fatalf("night reminder body: %+v, ok=%v", r, ok)
	}
	r, ok = ParseReminder("в 12 ночи", ctx)
	if !ok || Resolve(r, "", "", ctx.Today, ctx.Now) != (Resolved{Date: "2026-11-07", Time: "00:00"}) {
		t.Fatalf("dateless midnight reminder: %+v, ok=%v", r, ok)
	}
	for _, phrase := range []string{"в 7 ночи", "завтра в 8 часов ночи", "в 9:15 ночи завтра", "8.11 ночи"} {
		if r, ok := ParseReminder(phrase, ctx); ok {
			t.Fatalf("unsupported night body %q accepted: %+v", phrase, r)
		}
	}
}
