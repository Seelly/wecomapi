package wecomapi

import (
	"fmt"
	"io"
)

// 文档 API 回调事件类型。
const (
	// EventTypeDocChange 文档变更事件
	EventTypeDocChange EventType = "doc_change"
	// EventTypeSmartSheetChange 智能表格变更事件
	EventTypeSmartSheetChange EventType = "smart_sheet_change"
)

// 文档变更事件的 ChangeType。
const (
	// ChangeTypeDocMemberChange 修改文档成员事件
	ChangeTypeDocMemberChange ChangeType = "doc_member_change"
	// ChangeTypeDeleteDoc 删除文档事件
	ChangeTypeDeleteDoc ChangeType = "delete_doc"
	// ChangeTypeFormComplete 收集表完成事件
	ChangeTypeFormComplete ChangeType = "form_complete"
	// ChangeTypeDeleteForm 删除收集表事件
	ChangeTypeDeleteForm ChangeType = "delete_form"
	// ChangeTypeFormSettingsChange 修改收集表设置事件
	ChangeTypeFormSettingsChange ChangeType = "form_settings_change"
)

// 智能表格变更事件的 ChangeType。
const (
	// ChangeTypeSmartSheetAddField 新增字段事件
	ChangeTypeSmartSheetAddField ChangeType = "add_filed"
	// ChangeTypeSmartSheetUpdateField 更新字段事件
	ChangeTypeSmartSheetUpdateField ChangeType = "update_filed"
	// ChangeTypeSmartSheetDeleteField 删除字段事件
	ChangeTypeSmartSheetDeleteField ChangeType = "delete_filed"
	// ChangeTypeSmartSheetAddRecord 新增记录事件
	ChangeTypeSmartSheetAddRecord ChangeType = "add_record"
	// ChangeTypeSmartSheetUpdateRecord 更新记录事件
	ChangeTypeSmartSheetUpdateRecord ChangeType = "update_record"
	// ChangeTypeSmartSheetDeleteRecord 删除记录事件
	ChangeTypeSmartSheetDeleteRecord ChangeType = "delete_record"
)

// EventDocChange 文档变更事件（doc_change）。
//
// 根据 ChangeType 区分修改文档成员、删除文档、收集表完成、删除收集表、修改收集表设置。
type EventDocChange interface {
	messageKind

	// GetChangeType 变更类型
	GetChangeType() ChangeType
	// GetDocIDs 文档 ID 列表
	GetDocIDs() []string
	// GetFormIDs 收集表 ID 列表
	GetFormIDs() []string
}

type rxEventDocChange struct {
	ChangeType ChangeType `xml:"ChangeType"`
	DocIDs     []string   `xml:"DocId"`
	FormIDs    []string   `xml:"FormId"`
}

var _ EventDocChange = (*rxEventDocChange)(nil)

func (r *rxEventDocChange) formatInto(w io.Writer) {
	_, _ = fmt.Fprintf(
		w,
		"ChangeType: %#v, DocIDs: %#v, FormIDs: %#v",
		r.ChangeType,
		r.DocIDs,
		r.FormIDs,
	)
}

func (r *rxEventDocChange) GetChangeType() ChangeType { return r.ChangeType }
func (r *rxEventDocChange) GetDocIDs() []string       { return r.DocIDs }
func (r *rxEventDocChange) GetFormIDs() []string      { return r.FormIDs }

// EventSmartSheetChange 智能表格变更事件（smart_sheet_change）。
//
// 根据 ChangeType 区分字段和记录的新增、更新、删除。
type EventSmartSheetChange interface {
	messageKind

	// GetChangeType 变更类型
	GetChangeType() ChangeType
	// GetDocID 文档 ID
	GetDocID() string
	// GetSheetID 子表 ID
	GetSheetID() string
	// GetFieldIDs 字段 ID 列表
	GetFieldIDs() []string
	// GetRecordIDs 记录 ID 列表
	GetRecordIDs() []string
}

type rxEventSmartSheetChange struct {
	ChangeType ChangeType `xml:"ChangeType"`
	DocID      string     `xml:"DocId"`
	SheetID    string     `xml:"SheetId"`
	FieldIDs   []string   `xml:"FieldId"`
	RecordIDs  []string   `xml:"RecordId"`
}

var _ EventSmartSheetChange = (*rxEventSmartSheetChange)(nil)

func (r *rxEventSmartSheetChange) formatInto(w io.Writer) {
	_, _ = fmt.Fprintf(
		w,
		"ChangeType: %#v, DocID: %#v, SheetID: %#v, FieldIDs: %#v, RecordIDs: %#v",
		r.ChangeType,
		r.DocID,
		r.SheetID,
		r.FieldIDs,
		r.RecordIDs,
	)
}

func (r *rxEventSmartSheetChange) GetChangeType() ChangeType { return r.ChangeType }
func (r *rxEventSmartSheetChange) GetDocID() string          { return r.DocID }
func (r *rxEventSmartSheetChange) GetSheetID() string        { return r.SheetID }
func (r *rxEventSmartSheetChange) GetFieldIDs() []string     { return r.FieldIDs }
func (r *rxEventSmartSheetChange) GetRecordIDs() []string    { return r.RecordIDs }

// EventDocChange 如果消息为文档变更事件，则拿出相应消息参数，否则返回 nil, false
func (m *RxMessage) EventDocChange() (EventDocChange, bool) {
	y, ok := m.extras.(EventDocChange)
	return y, ok
}

// EventSmartSheetChange 如果消息为智能表格变更事件，则拿出相应消息参数，否则返回 nil, false
func (m *RxMessage) EventSmartSheetChange() (EventSmartSheetChange, bool) {
	y, ok := m.extras.(EventSmartSheetChange)
	return y, ok
}
