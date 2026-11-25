package wecomapi

import (
	"encoding/json"
	"strings"
	"testing"
)

// 创建日程时不得发出只读字段的空值：企业微信对 organizer="" 返回 40058。
func TestScheduleCreateOmitsReadOnlyFields(t *testing.T) {
	before := uint32(900)
	body, err := reqScheduleAdd{Schedule: Schedule{
		Summary: "课", StartTime: 1, EndTime: 2,
		Attendees: []ScheduleAttendee{{UserID: "u1"}},
		Reminders: ScheduleReminders{IsRemind: 1, RemindBeforeEventSecs: &before},
	}}.intoBody()
	if err != nil {
		t.Fatal(err)
	}
	got := string(body)
	for _, key := range []string{`"organizer"`, `"schedule_id"`, `"status"`, `"sequence"`, `"admins"`, `"cal_id"`, `"response_status"`, `"event_time"`} {
		if strings.Contains(got, key) {
			t.Errorf("创建请求不应包含 %s: %s", key, got)
		}
	}
	var decoded struct {
		Schedule map[string]any `json:"schedule"`
	}
	_ = json.Unmarshal(body, &decoded)
	if decoded.Schedule["summary"] != "课" || len(decoded.Schedule["attendees"].([]any)) != 1 || decoded.Schedule["is_whole_day"] != float64(0) {
		t.Fatalf("必要字段缺失: %s", got)
	}
}
