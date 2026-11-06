package mail

import (
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestRussianExpiry(t *testing.T) {
	for _, tc := range []struct {
		d    time.Duration
		want string
	}{
		{time.Hour, "1 час"}, {2 * time.Hour, "2 часа"}, {5 * time.Hour, "5 часов"},
		{11 * time.Hour, "11 часов"}, {21 * time.Hour, "21 час"}, {22 * time.Hour, "22 часа"},
		{25 * time.Hour, "25 часов"}, {49 * time.Hour, "2 дня"}, {5 * 24 * time.Hour, "5 дней"},
		{11 * 24 * time.Hour, "11 дней"}, {21 * 24 * time.Hour, "21 день"},
		{22 * 24 * time.Hour, "22 дня"}, {25 * 24 * time.Hour, "25 дней"},
		{111 * 24 * time.Hour, "111 дней"},
	} {
		if got := formatExpiry(tc.d, "ru"); got != tc.want {
			t.Errorf("%v: got %q, want %q", tc.d, got, tc.want)
		}
	}
}

func TestRussianInvitationRoles(t *testing.T) {
	for role, want := range map[string]string{"admin": "администратор", "editor": "редактор", "viewer": "наблюдатель", "member": "участник"} {
		for _, kind := range []Kind{KindProjectInvite, KindWorkspaceInvite, KindMemberInvitation} {
			_, body, err := Render(Message{Kind: kind, Lang: "ru", To: "anna@example.com",
				Vars: Vars{Role: role, Expiry: 22 * time.Hour, Link: "https://todo.example.com/join"}}, "https://todo.example.com")
			if err != nil {
				t.Fatal(err)
			}
			for _, text := range []string{"Кто-то", "ролью «" + want + "»", "через 22 часа", "https://todo.example.com/join"} {
				if !strings.Contains(body, text) {
					t.Errorf("%s/%s: missing %q in %s", kind, role, text, body)
				}
			}
		}
	}
}

func TestRussianDigestCounts(t *testing.T) {
	for _, n := range []int{1, 2, 5, 11, 21, 22, 25} {
		d := &DigestVars{WeekStart: "2026-11-02", WeekEnd: "2026-11-08", Completed: n, Created: n, Late: n, Slipped: n,
			Overdue: []DigestItem{{Title: "Отчёт", Due: "2026-11-01"}}, OverdueMore: n,
			Ahead: []DigestItem{{Title: "Встреча", Due: "2026-11-09"}}, AheadMore: n, Idle: []string{"Дом"}, IdleMore: n}
		subj, body, err := Render(Message{Kind: KindDigest, Lang: "ru", Vars: Vars{Digest: d}}, "https://todo.example.com")
		if err != nil {
			t.Fatal(err)
		}
		count := strconv.Itoa(n)
		if subj != "Ваша неделя в OpenTodo: завершено "+count {
			t.Errorf("count %d: subject %q", n, subj)
		}
		for _, label := range []string{"Завершено: ", "Создано: ", "Завершено с опозданием: ", "Не завершено к концу недели: "} {
			if !strings.Contains(body, label+count+"\n") {
				t.Errorf("count %d: missing %q", n, label+count)
			}
		}
		if strings.Count(body, "и ещё "+count+"\n") != 3 || !strings.Contains(body, "Входящие") || !strings.Contains(body, "Последний обзор: никогда") {
			t.Errorf("count %d: incorrect Russian digest lists: %s", n, body)
		}
	}
}
