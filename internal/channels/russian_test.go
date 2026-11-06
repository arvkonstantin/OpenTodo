package channels

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestRussianChannelTextAndContentPolicy(t *testing.T) {
	if len(texts["ru"]) != len(texts["en"]) {
		t.Fatal("Russian channel text keys differ from English")
	}
	for key, source := range texts["en"] {
		value := texts["ru"][key]
		if value == "" || strings.Count(value, "{name}") != strings.Count(source, "{name}") {
			t.Errorf("%s: missing text or actor placeholder", key)
		}
	}
	for _, tag := range []string{"ru", "ru-RU", "ru_BY"} {
		if normalizeLang(tag) != "ru" {
			t.Errorf("%s did not select Russian", tag)
		}
	}
	for _, keys := range [][2]string{
		{txtReminder, txtReminderWith}, {txtAssigned, txtAssignedBy}, {txtMentioned, txtMentionedBy},
		{txtUnblocked, txtUnblockedWith}, {txtPageReplied, txtPageRepliedBy},
		{txtPageDone, txtPageDoneBy}, {txtPageReopened, txtPageReopenedBy}, {txtTurn, txtTurnWith},
		{txtRoutineFailed, txtRoutineFailedWith},
	} {
		e := event{kind: EventReminder, key: keys[0], withKey: keys[1], taskID: "t1",
			taskTitle: "Секретная задача", actorName: "Алексей", excerpt: "Закрытый текст"}
		generic := e.render("ru", "https://todo.example.com", false)
		if generic.Title != "OpenTodo" || generic.Body != texts["ru"][keys[0]] {
			t.Errorf("%s: wrong generic message: %+v", keys[0], generic)
		}
		raw, _ := json.Marshal(generic)
		for _, secret := range []string{e.taskTitle, e.actorName, e.excerpt} {
			if strings.Contains(string(raw), secret) {
				t.Errorf("%s: private content left in generic message", keys[0])
			}
		}
		named := e.render("ru", "", true)
		want := strings.ReplaceAll(texts["ru"][keys[1]], "{name}", "Алексей") + ": Закрытый текст"
		if named.Title != "Секретная задача" || named.Body != want {
			t.Errorf("%s: wrong named message: %+v", keys[0], named)
		}
		e.taskTitle, e.actorName, e.excerpt = "enc:v1:abc", "enc:v1:def", "enc:v1:ghi"
		cipher := e.render("ru", "", true)
		if cipher.Title != "OpenTodo" || cipher.Body != texts["ru"][keys[0]] || cipher.ActorName != "" || cipher.TaskTitle != "" {
			t.Errorf("%s: ciphertext content policy: %+v", keys[0], cipher)
		}
	}
	m := testMessage("ru", "https://todo.example.com", 10)
	if m.Title != "Тест от OpenTodo" || m.Body != "Этот канал работает." {
		t.Fatalf("test message: %+v", m)
	}
}
