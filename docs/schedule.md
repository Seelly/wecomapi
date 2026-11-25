# 日程接口

## Models

### `Schedule` 日程信息

> 创建日程时不得发出只读字段的空值（例如 `organizer: ""` 会被企业微信以 40058 拒绝），
> 因此下列只读/可选字段在 JSON 列直接声明了 `omitempty`。

Name|JSON|Type|Doc
:---|:---|:---|:--
`ScheduleID`|`schedule_id,omitempty`|`string`| 日程ID。创建日程时无需填写；更新、获取日程时使用
`Admins`|`admins,omitempty`|`[]string`| 日程的管理员userid列表，管理员必须在共享成员的列表中。最多指定3人
`Organizer`|`organizer,omitempty`|`string`| 组织者。仅获取日程时返回；更新日程时暂不支持变更组织者
`Attendees`|`attendees,omitempty`|`[]ScheduleAttendee`| 日程参与者列表。创建日程最多支持1000人，更新日程最多支持2000人
`Summary`|`summary`|`string`| 日程标题。0~128字符。不填会默认显示为“新建事件”
`Description`|`description`|`string`| 日程描述。创建日程不多于1000个字符，更新日程不多于512个字符
`Reminders`|`reminders`|`ScheduleReminders`| 提醒相关信息
`Location`|`location`|`string`| 日程地址。不多于128个字符
`StartTime`|`start_time`|`uint32`| 日程开始时间，Unix时间戳
`EndTime`|`end_time`|`uint32`| 日程结束时间，Unix时间戳
`Status`|`status,omitempty`|`uint32`| 日程状态。0-正常；1-已取消。仅获取日程时返回。注意，被取消的日程也可以拉取详情，调用者需要检查status
`IsWholeDay`|`is_whole_day`|`uint32`| 是否设置为全天日程。0-否；1-是
`CalID`|`cal_id,omitempty`|`string`| 日程所属日历ID。该日历必须是access_token所对应应用所创建的日历。如果不填，那么插入到access_token所对应应用的默认日历上。第三方应用必须指定cal_id
`Sequence`|`sequence,omitempty`|`uint32`| 日历下的日程序号。仅获取日历下的日程列表时返回

### `ScheduleAttendee` 日程参与者

Name|JSON|Type|Doc
:---|:---|:---|:--
`UserID`|`userid`|`string`| 日程参与者ID，不多于64字节
`ResponseStatus`|`response_status,omitempty`|`uint32`| 日程参与者的接受状态。0-未处理；1-待定；2-全部接受；3-仅接受一次；4-拒绝。仅获取日程时返回
`EventTime`|`event_time,omitempty`|`uint32`| 对重复日程，表示该参与者本次日程的开始时间。仅获取日程详情时返回

### `ScheduleReminders` 提醒相关信息

Name|JSON|Type|Doc
:---|:---|:---|:--
`IsRemind`|`is_remind`|`uint32`| 是否需要提醒。0-否；1-是
`RemindBeforeEventSecs`|`remind_before_event_secs`|`*uint32`| 日程开始前多少秒提醒，当is_remind为1时有效。目前仅支持以下数值：0-事件开始时；300-事件开始前5分钟；900-事件开始前15分钟；3600-事件开始前1小时；86400-事件开始前1天
`RemindTimeDiffs`|`remind_time_diffs`|`[]int32`| 提醒时间与日程开始时间的差值，当is_remind为1时有效，可以指定多个提醒时间。该字段与remind_before_event_secs仅一个字段会生效，当该字段有传值且列表不为空时，优先以该字段为准
`IsRepeat`|`is_repeat`|`uint32`| 是否重复日程。0-否；1-是
`RepeatType`|`repeat_type`|`*uint32`| 重复类型，当is_repeat为1时有效。0-每日；1-每周；2-每月；5-每年；7-工作日
`RepeatUntil`|`repeat_until`|`*uint32`| 重复结束时刻，Unix时间戳，当is_repeat为1时有效。不填或填0表示一直重复
`IsCustomRepeat`|`is_custom_repeat`|`*uint32`| 是否自定义重复。0-否；1-是。当is_repeat为1时有效
`RepeatInterval`|`repeat_interval`|`*uint32`| 重复间隔，仅当指定为自定义重复时有效。该字段随repeat_type不同而含义不同
`RepeatDayOfWeek`|`repeat_day_of_week`|`[]uint32`| 每周周几重复，取值范围1~7，分别表示周一至周日。仅当指定为自定义重复且重复类型为每周时有效
`RepeatDayOfMonth`|`repeat_day_of_month`|`[]uint32`| 每月哪几天重复，取值范围1~31。仅当指定为自定义重复且重复类型为每月时有效
`Timezone`|`timezone`|`*int32`| 时区。UTC偏移量表示（即偏离零时区的小时数），东区为正数，西区为负数，取值范围-12~+12。默认为北京时间东八区
`ExcludeTimeList`|`exclude_time_list`|`[]ScheduleExcludeTime`| 重复日程的例外时间列表。仅获取日程详情时返回

### `ScheduleExcludeTime` 重复日程的例外时间

Name|JSON|Type|Doc
:---|:---|:---|:--
`StartTime`|`start_time`|`uint32`| 例外日程的开始时间，Unix时间戳

## API calls

Name|Request type|Response type|URL|Access token|Doc
:---|:---|:---|:---|:---|:---
`execScheduleAdd`|`reqScheduleAdd`|`respScheduleAdd`|`POST /cgi-bin/oa/schedule/add`|yes|创建日程
`execScheduleUpdate`|`reqScheduleUpdate`|`respScheduleUpdate`|`POST /cgi-bin/oa/schedule/update`|yes|更新日程
`execScheduleDel`|`reqScheduleDel`|`respScheduleDel`|`POST /cgi-bin/oa/schedule/del`|yes|取消日程
`execScheduleGet`|`reqScheduleGet`|`respScheduleGet`|`POST /cgi-bin/oa/schedule/get`|yes|获取日程详情
`execScheduleAddAttendees`|`reqScheduleAttendees`|`respScheduleAddAttendees`|`POST /cgi-bin/oa/schedule/add_attendees`|yes|新增日程参与者
`execScheduleDelAttendees`|`reqScheduleAttendees`|`respScheduleDelAttendees`|`POST /cgi-bin/oa/schedule/del_attendees`|yes|删除日程参与者
`execScheduleGetByCalendar`|`reqScheduleGetByCalendar`|`respScheduleGetByCalendar`|`POST /cgi-bin/oa/schedule/get_by_calendar`|yes|获取日历下的日程列表

# 日历接口

## Models

### `Calendar` 日历信息

Name|JSON|Type|Doc
:---|:---|:---|:--
`CalID`|`cal_id`|`string`| 日历ID。创建日历时无需填写；更新、获取日历时使用
`Admins`|`admins`|`[]string`| 日历的管理员userid列表，管理员必须在通知范围成员的列表中。最多指定3人
`SetAsDefault`|`set_as_default`|`*uint32`| 是否将该日历设置为access_token所对应应用的默认日历。0-否；1-是。默认为0，即不设为默认日历。第三方应用不支持使用该参数。仅创建日历时有效
`Summary`|`summary`|`string`| 日历标题。1~128字符
`Color`|`color`|`string`| 日历在终端上显示的颜色，RGB颜色编码16进制表示，例如：“#0000FF”表示纯蓝色
`Description`|`description`|`string`| 日历描述。0~512字符
`IsPublic`|`is_public`|`*uint32`| 是否公共日历。0-否；1-是。注意：每个人最多可创建或订阅100个公共日历。该属性不可更新
`PublicRange`|`public_range`|`CalendarPublicRange`| 公开范围。仅当是公共日历时有效
`IsCorpCalendar`|`is_corp_calendar`|`*uint32`| 是否全员日历。0-否；1-是。注意：每个企业最多可创建20个全员日历；全员日历也是公共日历的一种，需要指定public_range；全员日历不支持指定颜色、默认日历、只读权限。该属性不可更新
`Shares`|`shares`|`[]CalendarShare`| 日历通知范围成员列表。最多2000人

### `CalendarShare` 日历通知范围成员

Name|JSON|Type|Doc
:---|:---|:---|:--
`UserID`|`userid`|`string`| 日历通知范围成员的id
`Permission`|`permission`|`*uint32`| 日历通知范围成员权限（不填则默认为「可查看」）。1-可查看；3-仅查看闲忙状态

### `CalendarPublicRange` 公共日历的公开范围

Name|JSON|Type|Doc
:---|:---|:---|:--
`UserIDs`|`userids`|`[]string`| 公开的成员列表范围。最多指定1000个成员
`PartyIDs`|`partyids`|`[]uint32`| 公开的部门列表范围。最多指定100个部门

### `CalendarFailShare` 无效的日历通知范围成员

Name|JSON|Type|Doc
:---|:---|:---|:--
`UserID`|`userid`|`string`| 日历通知范围成员的id
`ErrCode`|`errcode`|`int32`| 错误码
`ErrMsg`|`errmsg`|`string`| 错误码说明

### `CalendarFailResult` 无效的输入内容

Name|JSON|Type|Doc
:---|:---|:---|:--
`Shares`|`shares`|`[]CalendarFailShare`| 无效的日历通知范围成员列表

## API calls

Name|Request type|Response type|URL|Access token|Doc
:---|:---|:---|:---|:---|:---
`execCalendarAdd`|`reqCalendarAdd`|`respCalendarAdd`|`POST /cgi-bin/oa/calendar/add`|yes|创建日历
`execCalendarUpdate`|`reqCalendarUpdate`|`respCalendarUpdate`|`POST /cgi-bin/oa/calendar/update`|yes|更新日历
`execCalendarDel`|`reqCalendarDel`|`respCalendarDel`|`POST /cgi-bin/oa/calendar/del`|yes|删除日历
`execCalendarGet`|`reqCalendarGet`|`respCalendarGet`|`POST /cgi-bin/oa/calendar/get`|yes|获取日历详情
