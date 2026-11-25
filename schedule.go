package wecomapi

// CreateSchedule 创建日程
//
// 返回创建成功的日程ID。日程会插入到 Schedule.CalID 指定的日历上；不填则插入到
// access_token 所对应应用的默认日历上。
func (c *App) CreateSchedule(schedule Schedule) (scheduleID string, err error) {
	resp, err := c.execScheduleAdd(reqScheduleAdd{
		Schedule: schedule,
	})
	if err != nil {
		return "", err
	}
	return resp.ScheduleID, nil
}

// UpdateScheduleReq 更新日程的请求参数
type UpdateScheduleReq struct {
	// Schedule 日程信息，ScheduleID 必填。注意更新操作是覆盖式的，而不是增量式；
	// 不可更新组织者和日程所属日历ID
	Schedule Schedule
	// SkipAttendees 是否不更新参与人
	SkipAttendees bool
	// OpMode 操作模式。是重复日程时有效。0-默认更新所有日程；1-仅更新此日程；2-更新本次及后续日程。
	// nil 表示使用默认值 0
	OpMode *uint32
	// OpStartTime 操作起始时间（Unix时间戳）。仅当 OpMode 为 1 或 2 时有效，
	// 该时间必须是重复日程的某一次开始时间
	OpStartTime *uint32
}

// UpdateSchedule 更新日程
//
// 注意：更新操作是覆盖式的，而不是增量式。不可更新组织者和日程所属日历ID。
// 若需增量式更新成员，请使用 AddScheduleAttendees 与 DelScheduleAttendees。
func (c *App) UpdateSchedule(req UpdateScheduleReq) error {
	body := reqScheduleUpdate{
		OpMode:      req.OpMode,
		OpStartTime: req.OpStartTime,
		Schedule:    req.Schedule,
	}
	if req.SkipAttendees {
		v := uint32(1)
		body.SkipAttendees = &v
	}

	_, err := c.execScheduleUpdate(body)
	return err
}

// DeleteScheduleReq 取消日程的请求参数
type DeleteScheduleReq struct {
	// ScheduleID 日程ID
	ScheduleID string
	// OpMode 操作模式。是重复日程时有效。0-默认取消所有日程；1-仅取消此日程；2-取消本次及后续日程。
	// nil 表示使用默认值 0
	OpMode *uint32
	// OpStartTime 操作起始时间（Unix时间戳）。仅当 OpMode 为 1 或 2 时有效，
	// 该时间必须是重复日程的某一次开始时间
	OpStartTime *uint32
}

// DeleteSchedule 取消日程
func (c *App) DeleteSchedule(req DeleteScheduleReq) error {
	_, err := c.execScheduleDel(reqScheduleDel{
		ScheduleID:  req.ScheduleID,
		OpMode:      req.OpMode,
		OpStartTime: req.OpStartTime,
	})
	return err
}

// GetSchedules 批量获取日程详情
//
// 一次最多拉取1000条。注意，被取消的日程也可以拉取详情，调用者需要检查 Status 字段。
func (c *App) GetSchedules(scheduleIDs []string) ([]Schedule, error) {
	resp, err := c.execScheduleGet(reqScheduleGet{
		ScheduleIDList: scheduleIDs,
	})
	if err != nil {
		return nil, err
	}
	return resp.ScheduleList, nil
}

// AddScheduleAttendees 新增日程参与者（增量式）
func (c *App) AddScheduleAttendees(scheduleID string, attendees []ScheduleAttendee) error {
	_, err := c.execScheduleAddAttendees(reqScheduleAttendees{
		ScheduleID: scheduleID,
		Attendees:  attendees,
	})
	return err
}

// DelScheduleAttendees 删除日程参与者（增量式）
func (c *App) DelScheduleAttendees(scheduleID string, attendees []ScheduleAttendee) error {
	_, err := c.execScheduleDelAttendees(reqScheduleAttendees{
		ScheduleID: scheduleID,
		Attendees:  attendees,
	})
	return err
}

// ListSchedulesByCalendar 获取指定的日历下的日程列表（分页）
//
// 仅可获取应用自己创建的日历下的日程。offset 以 0 为起点；当返回的列表为空时，
// 表示 offset 已经过大，此时应终止获取。
func (c *App) ListSchedulesByCalendar(calID string, offset, limit uint32) ([]Schedule, error) {
	resp, err := c.execScheduleGetByCalendar(reqScheduleGetByCalendar{
		CalID:  calID,
		Offset: offset,
		Limit:  limit,
	})
	if err != nil {
		return nil, err
	}
	return resp.ScheduleList, nil
}
