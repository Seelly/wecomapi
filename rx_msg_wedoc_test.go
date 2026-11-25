package wecomapi

import "testing"

func TestRxEventDocChange(t *testing.T) {
	body := []byte(`<xml>
		<ToUserName><![CDATA[corp]]></ToUserName>
		<FromUserName><![CDATA[user]]></FromUserName>
		<CreateTime>1348831860</CreateTime>
		<MsgType><![CDATA[event]]></MsgType>
		<Event><![CDATA[doc_change]]></Event>
		<ChangeType><![CDATA[doc_member_change]]></ChangeType>
		<DocId><![CDATA[DOC1]]></DocId>
		<DocId><![CDATA[DOC2]]></DocId>
	</xml>`)

	msg, err := fromEnvelope(body)
	if err != nil {
		t.Fatal(err)
	}
	ev, ok := msg.EventDocChange()
	if !ok {
		t.Fatalf("EventDocChange not found: %s", msg)
	}
	if ev.GetChangeType() != ChangeTypeDocMemberChange {
		t.Errorf("change type = %q", ev.GetChangeType())
	}
	if got := ev.GetDocIDs(); len(got) != 2 || got[0] != "DOC1" || got[1] != "DOC2" {
		t.Errorf("doc ids = %#v", got)
	}
	if got := ev.GetFormIDs(); len(got) != 0 {
		t.Errorf("form ids = %#v", got)
	}
}

func TestRxEventDocChangeFormComplete(t *testing.T) {
	body := []byte(`<xml>
		<ToUserName><![CDATA[corp]]></ToUserName>
		<FromUserName><![CDATA[user]]></FromUserName>
		<CreateTime>1348831860</CreateTime>
		<MsgType><![CDATA[event]]></MsgType>
		<Event><![CDATA[doc_change]]></Event>
		<ChangeType><![CDATA[form_complete]]></ChangeType>
		<FormId><![CDATA[FORM1]]></FormId>
	</xml>`)

	msg, err := fromEnvelope(body)
	if err != nil {
		t.Fatal(err)
	}
	ev, ok := msg.EventDocChange()
	if !ok {
		t.Fatalf("EventDocChange not found: %s", msg)
	}
	if ev.GetChangeType() != ChangeTypeFormComplete {
		t.Errorf("change type = %q", ev.GetChangeType())
	}
	if got := ev.GetFormIDs(); len(got) != 1 || got[0] != "FORM1" {
		t.Errorf("form ids = %#v", got)
	}
}

func TestRxEventSmartSheetChange(t *testing.T) {
	body := []byte(`<xml>
		<ToUserName><![CDATA[corp]]></ToUserName>
		<FromUserName><![CDATA[user]]></FromUserName>
		<CreateTime>1348831860</CreateTime>
		<MsgType><![CDATA[event]]></MsgType>
		<Event><![CDATA[smart_sheet_change]]></Event>
		<ChangeType><![CDATA[update_record]]></ChangeType>
		<DocId><![CDATA[DOC1]]></DocId>
		<SheetId><![CDATA[SHEET1]]></SheetId>
		<RecordId><![CDATA[R1]]></RecordId>
		<RecordId><![CDATA[R2]]></RecordId>
	</xml>`)

	msg, err := fromEnvelope(body)
	if err != nil {
		t.Fatal(err)
	}
	ev, ok := msg.EventSmartSheetChange()
	if !ok {
		t.Fatalf("EventSmartSheetChange not found: %s", msg)
	}
	if ev.GetChangeType() != ChangeTypeSmartSheetUpdateRecord {
		t.Errorf("change type = %q", ev.GetChangeType())
	}
	if ev.GetDocID() != "DOC1" || ev.GetSheetID() != "SHEET1" {
		t.Errorf("doc/sheet = %q/%q", ev.GetDocID(), ev.GetSheetID())
	}
	if got := ev.GetRecordIDs(); len(got) != 2 || got[0] != "R1" || got[1] != "R2" {
		t.Errorf("record ids = %#v", got)
	}
	if got := ev.GetFieldIDs(); len(got) != 0 {
		t.Errorf("field ids = %#v", got)
	}
}
