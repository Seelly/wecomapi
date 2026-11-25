package wecomapi

import (
	"testing"

	c "github.com/smartystreets/goconvey/convey"
)

func TestRxMessageScheduleCalendarEvents(t *testing.T) {
	c.Convey("解析日历与日程相关的事件回调", t, func() {
		c.Convey("修改日历事件", func() {
			body := []byte("<xml><ToUserName><![CDATA[toUser]]></ToUserName><FromUserName><![CDATA[sys]]></FromUserName><CreateTime>1348831860</CreateTime><MsgType><![CDATA[event]]></MsgType><Event><![CDATA[modify_calendar]]></Event><CalId><![CDATA[wcjgewCwAAqeJcPI1d8Pwbjt7nttzAAA]]></CalId></xml>")

			msg, err := fromEnvelope(body)
			c.So(err, c.ShouldBeNil)
			c.So(msg.MsgType, c.ShouldEqual, MessageTypeEvent)
			c.So(msg.Event, c.ShouldEqual, EventTypeModifyCalendar)

			e, ok := msg.EventModifyCalendar()
			c.So(ok, c.ShouldBeTrue)
			c.So(e.GetCalID(), c.ShouldEqual, "wcjgewCwAAqeJcPI1d8Pwbjt7nttzAAA")
		})

		c.Convey("删除日历事件", func() {
			body := []byte("<xml><ToUserName><![CDATA[toUser]]></ToUserName><FromUserName><![CDATA[sys]]></FromUserName><CreateTime>1348831860</CreateTime><MsgType><![CDATA[event]]></MsgType><Event><![CDATA[delete_calendar]]></Event><CalId><![CDATA[wcjgewCwAAqeJcPI1d8Pwbjt7nttzAAA]]></CalId></xml>")

			msg, err := fromEnvelope(body)
			c.So(err, c.ShouldBeNil)
			c.So(msg.Event, c.ShouldEqual, EventTypeDeleteCalendar)

			e, ok := msg.EventDeleteCalendar()
			c.So(ok, c.ShouldBeTrue)
			c.So(e.GetCalID(), c.ShouldEqual, "wcjgewCwAAqeJcPI1d8Pwbjt7nttzAAA")
		})

		c.Convey("日程回执事件", func() {
			body := []byte("<xml><ToUserName><![CDATA[toUser]]></ToUserName><FromUserName><![CDATA[foobar]]></FromUserName><CreateTime>1348831860</CreateTime><MsgType><![CDATA[event]]></MsgType><Event><![CDATA[respond_schedule]]></Event><CalId><![CDATA[wcjgewCwAAqeJcPI1d8Pwbjt7nttzAAA]]></CalId><ScheduleId><![CDATA[17c7d2bd9f20d652840f72f59e796AAA]]></ScheduleId></xml>")

			msg, err := fromEnvelope(body)
			c.So(err, c.ShouldBeNil)
			c.So(msg.Event, c.ShouldEqual, EventTypeRespondSchedule)

			e, ok := msg.EventRespondSchedule()
			c.So(ok, c.ShouldBeTrue)
			c.So(e.GetCalID(), c.ShouldEqual, "wcjgewCwAAqeJcPI1d8Pwbjt7nttzAAA")
			c.So(e.GetScheduleID(), c.ShouldEqual, "17c7d2bd9f20d652840f72f59e796AAA")
		})

		c.Convey("添加/修改/删除日程事件共用日程事件体", func() {
			for _, ev := range []struct {
				event EventType
				get   func(*RxMessage) (ScheduleEventExtras, bool)
			}{
				{EventTypeAddSchedule, (*RxMessage).EventAddSchedule},
				{EventTypeModifySchedule, (*RxMessage).EventModifySchedule},
				{EventTypeDeleteSchedule, (*RxMessage).EventDeleteSchedule},
			} {
				body := []byte("<xml><ToUserName><![CDATA[toUser]]></ToUserName><FromUserName><![CDATA[sys]]></FromUserName><CreateTime>1348831860</CreateTime><MsgType><![CDATA[event]]></MsgType><Event><![CDATA[" + string(ev.event) + "]]></Event><CalId><![CDATA[wcjgewCwAAqeJcPI1d8Pwbjt7nttzAAA]]></CalId><ScheduleId><![CDATA[17c7d2bd9f20d652840f72f59e796AAA]]></ScheduleId></xml>")

				msg, err := fromEnvelope(body)
				c.So(err, c.ShouldBeNil)
				c.So(msg.Event, c.ShouldEqual, ev.event)

				e, ok := ev.get(msg)
				c.So(ok, c.ShouldBeTrue)
				c.So(e.GetCalID(), c.ShouldEqual, "wcjgewCwAAqeJcPI1d8Pwbjt7nttzAAA")
				c.So(e.GetScheduleID(), c.ShouldEqual, "17c7d2bd9f20d652840f72f59e796AAA")
			}
		})

		c.Convey("事件类型不匹配时取值返回 false", func() {
			body := []byte("<xml><ToUserName><![CDATA[toUser]]></ToUserName><FromUserName><![CDATA[sys]]></FromUserName><CreateTime>1348831860</CreateTime><MsgType><![CDATA[event]]></MsgType><Event><![CDATA[delete_calendar]]></Event><CalId><![CDATA[wcjgewCwAAqeJcPI1d8Pwbjt7nttzAAA]]></CalId></xml>")

			msg, err := fromEnvelope(body)
			c.So(err, c.ShouldBeNil)

			_, ok := msg.EventAddSchedule()
			c.So(ok, c.ShouldBeFalse)

			_, ok = msg.EventRespondSchedule()
			c.So(ok, c.ShouldBeFalse)
		})
	})
}
