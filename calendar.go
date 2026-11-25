package wecomapi

// CreateCalendar 创建日历
//
// 返回创建成功的日历ID。若 Shares 中存在无效的通知范围成员，errcode 仍为成功，
// 无效成员会通过 failResult 返回，调用者应检查其 Shares 字段。
func (c *App) CreateCalendar(calendar Calendar) (calID string, failResult CalendarFailResult, err error) {
	resp, err := c.execCalendarAdd(reqCalendarAdd{
		Calendar: calendar,
	})
	if err != nil {
		return "", CalendarFailResult{}, err
	}
	return resp.CalID, resp.FailResult, nil
}

// UpdateCalendarReq 更新日历的请求参数
type UpdateCalendarReq struct {
	// Calendar 日历信息，CalID、Summary、Color 必填。
	// 注意更新操作是覆盖式的，而不是增量式；is_public、is_corp_calendar 等属性不可更新
	Calendar Calendar
	// SkipPublicRange 是否不更新可订阅范围
	SkipPublicRange bool
}

// UpdateCalendar 更新日历
//
// 注意：更新操作是覆盖式的，而不是增量式。若 Shares 中存在无效的通知范围成员，
// errcode 仍为成功，无效成员会通过 failResult 返回，调用者应检查其 Shares 字段。
func (c *App) UpdateCalendar(req UpdateCalendarReq) (failResult CalendarFailResult, err error) {
	body := reqCalendarUpdate{
		Calendar: req.Calendar,
	}
	if req.SkipPublicRange {
		v := uint32(1)
		body.SkipPublicRange = &v
	}

	resp, err := c.execCalendarUpdate(body)
	if err != nil {
		return CalendarFailResult{}, err
	}
	return resp.FailResult, nil
}

// DeleteCalendar 删除日历
func (c *App) DeleteCalendar(calID string) error {
	_, err := c.execCalendarDel(reqCalendarDel{
		CalID: calID,
	})
	return err
}

// GetCalendars 批量获取日历详情
//
// calIDList 为日历ID列表（调用创建日历接口后获得），一次最多可获取1000条。
func (c *App) GetCalendars(calIDList []string) ([]Calendar, error) {
	resp, err := c.execCalendarGet(reqCalendarGet{
		CalIDList: calIDList,
	})
	if err != nil {
		return nil, err
	}
	return resp.CalendarList, nil
}
