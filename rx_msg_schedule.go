package wecomapi

import (
	"fmt"
	"io"
)

// NOTE: 本文件中的事件常量与结构体按 rx_msg.md.go 的既有生成模式手写维护。
// rx 层的 sdkcodegen 规格（docs/rx_msg.md）在仓库重组时被剥离，恢复后应迁回规格文件。

// 事件类型：日历与日程相关回调（MsgType 为 event）
const (
	// EventTypeModifyCalendar 修改日历事件
	EventTypeModifyCalendar EventType = "modify_calendar"
	// EventTypeDeleteCalendar 删除日历事件
	EventTypeDeleteCalendar EventType = "delete_calendar"
	// EventTypeAddSchedule 添加日程事件
	EventTypeAddSchedule EventType = "add_schedule"
	// EventTypeModifySchedule 修改日程事件
	EventTypeModifySchedule EventType = "modify_schedule"
	// EventTypeDeleteSchedule 删除日程事件
	EventTypeDeleteSchedule EventType = "delete_schedule"
	// EventTypeRespondSchedule 日程回执事件
	EventTypeRespondSchedule EventType = "respond_schedule"
)

// rxCalendarEvent 日历变更事件（修改日历、删除日历）的事件体
type rxCalendarEvent struct {
	// CalID 日历ID
	CalID string `xml:"CalId"`
}

// rxScheduleEvent 日程变更/回执事件（添加日程、修改日程、删除日程、日程回执）的事件体
type rxScheduleEvent struct {
	// CalID 日历ID
	CalID string `xml:"CalId"`
	// ScheduleID 日程ID
	ScheduleID string `xml:"ScheduleId"`
}

// CalendarEventExtras 日历变更事件的参数。
type CalendarEventExtras interface {
	messageKind

	// GetCalID 返回日历ID
	GetCalID() string
}

var _ CalendarEventExtras = (*rxCalendarEvent)(nil)

func (r *rxCalendarEvent) formatInto(w io.Writer) {
	_, _ = fmt.Fprintf(w, "CalID: %#v", r.CalID)
}

func (r *rxCalendarEvent) GetCalID() string {
	return r.CalID
}

// ScheduleEventExtras 日程变更/回执事件的参数。
type ScheduleEventExtras interface {
	messageKind

	// GetCalID 返回日历ID
	GetCalID() string
	// GetScheduleID 返回日程ID
	GetScheduleID() string
}

var _ ScheduleEventExtras = (*rxScheduleEvent)(nil)

func (r *rxScheduleEvent) formatInto(w io.Writer) {
	_, _ = fmt.Fprintf(w, "CalID: %#v, ScheduleID: %#v", r.CalID, r.ScheduleID)
}

func (r *rxScheduleEvent) GetCalID() string {
	return r.CalID
}

func (r *rxScheduleEvent) GetScheduleID() string {
	return r.ScheduleID
}
