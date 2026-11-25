package wecomapi

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type wedocStaticTokenProvider struct{}

func (wedocStaticTokenProvider) GetToken(context.Context) (string, error) {
	return "test-token", nil
}

func newWedocTestApp(t *testing.T, handler http.HandlerFunc) (*App, *httptest.Server) {
	t.Helper()

	srv := httptest.NewServer(handler)
	app := New(
		"corp",
		WithQYAPIHost(srv.URL),
		WithAccessTokenProvider(wedocStaticTokenProvider{}),
	).WithApp("secret", 1)

	return app, srv
}

func TestWedocCreateDoc(t *testing.T) {
	app, srv := newWedocTestApp(t, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("method = %s, want POST", r.Method)
		}
		if r.URL.Path != "/cgi-bin/wedoc/create_doc" {
			t.Errorf("path = %s", r.URL.Path)
		}
		if got := r.URL.Query().Get("access_token"); got != "test-token" {
			t.Errorf("access_token = %q", got)
		}

		var body map[string]any
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			t.Fatal(err)
		}
		if body["doc_type"] != float64(3) || body["doc_name"] != "n" {
			t.Errorf("body = %#v", body)
		}
		if _, ok := body["spaceid"]; ok {
			t.Errorf("omitempty spaceid should be absent: %#v", body)
		}
		if _, ok := body["admin_users"]; ok {
			t.Errorf("omitempty admin_users should be absent: %#v", body)
		}

		_, _ = io.WriteString(w, `{"errcode":0,"errmsg":"ok","url":"https://doc.example/x","docid":"DOC1"}`)
	})
	defer srv.Close()

	resp, err := app.WedocCreateDoc(WedocCreateDocReq{DocType: 3, DocName: "n"})
	if err != nil {
		t.Fatal(err)
	}
	if resp.DocID != "DOC1" || resp.URL != "https://doc.example/x" {
		t.Fatalf("resp = %#v", resp)
	}
}

func TestWedocWebhookAddRecords(t *testing.T) {
	app, srv := newWedocTestApp(t, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/cgi-bin/wedoc/smartsheet/webhook" {
			t.Errorf("path = %s", r.URL.Path)
		}
		if got := r.URL.Query().Get("key"); got != "WHKEY" {
			t.Errorf("key = %q", got)
		}
		if got := r.URL.Query().Get("access_token"); got != "" {
			t.Errorf("access_token should be absent, got %q", got)
		}

		var body struct {
			AddRecords []struct {
				Values map[string]any `json:"values"`
			} `json:"add_records"`
		}
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			t.Fatal(err)
		}
		if len(body.AddRecords) != 1 || body.AddRecords[0].Values["f"] == nil {
			t.Errorf("body = %#v", body)
		}

		_, _ = io.WriteString(w, `{"errcode":0,"errmsg":"ok","add_records":[{"record_id":"R1","values":{"f":[{"type":"text","text":"x"}]}}]}`)
	})
	defer srv.Close()

	resp, err := app.WedocWebhookAddRecords("WHKEY", WedocWebhookAddRecordsReq{
		AddRecords: []WedocAddRecord{{
			Values: map[string]any{
				"f": []map[string]any{{"type": "text", "text": "x"}},
			},
		}},
	})
	if err != nil {
		t.Fatal(err)
	}
	if len(resp.AddRecords) != 1 || resp.AddRecords[0].RecordID != "R1" {
		t.Fatalf("resp = %#v", resp)
	}
}

func TestWedocCreateFormNesting(t *testing.T) {
	app, srv := newWedocTestApp(t, func(w http.ResponseWriter, r *http.Request) {
		raw, err := io.ReadAll(r.Body)
		if err != nil {
			t.Fatal(err)
		}
		var body map[string]any
		if err := json.Unmarshal(raw, &body); err != nil {
			t.Fatal(err)
		}

		info, ok := body["form_info"].(map[string]any)
		if !ok {
			t.Fatalf("form_info missing: %s", raw)
		}
		question := info["form_question"].(map[string]any)
		items := question["items"].([]any)
		first := items[0].(map[string]any)
		if first["question_id"] != float64(1) || first["title"] != "q" {
			t.Errorf("question = %#v", first)
		}
		if _, ok := info["formid"]; ok {
			t.Errorf("empty formid should be omitted: %#v", info)
		}
		if _, ok := first["note"]; ok {
			t.Errorf("empty note should be omitted: %#v", first)
		}

		_, _ = io.WriteString(w, `{"errcode":0,"errmsg":"ok","formid":"FORM1"}`)
	})
	defer srv.Close()

	resp, err := app.WedocCreateForm(WedocCreateFormReq{
		FormInfo: WedocFormInfo{
			FormTitle: "t",
			FormQuestion: WedocFormQuestion{Items: []WedocFormItem{{
				QuestionID: 1,
				Title:      "q",
				Pos:        1,
				Status:     1,
				ReplyType:  1,
				MustReply:  true,
				OptionItem: []WedocFormOptionItem{{Key: 1, Value: "A", Status: 1}},
			}}},
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	if resp.FormID != "FORM1" {
		t.Fatalf("resp = %#v", resp)
	}
}

func TestWedocGetDocAuthResponse(t *testing.T) {
	app, srv := newWedocTestApp(t, func(w http.ResponseWriter, r *http.Request) {
		_, _ = io.WriteString(w, `{
			"errcode":0,"errmsg":"ok",
			"access_rule":{"enable_corp_internal":true,"corp_internal_auth":1,"ban_share_external":false},
			"secure_setting":{"enable_readonly_copy":false,"watermark":{"margin_type":2,"show_text":true,"text":"x"},"enable_readonly_comment":false},
			"doc_member_list":[{"type":1,"userid":"u1","auth":7},{"type":1,"tmp_external_userid":"e1","auth":1}],
			"co_auth_list":[{"type":2,"departmentid":1,"auth":1}]
		}`)
	})
	defer srv.Close()

	resp, err := app.WedocGetDocAuth(WedocGetDocAuthReq{DocID: "DOC1"})
	if err != nil {
		t.Fatal(err)
	}
	if !resp.AccessRule.EnableCorpInternal || resp.AccessRule.CorpInternalAuth != 1 {
		t.Errorf("access_rule = %#v", resp.AccessRule)
	}
	if resp.SecureSetting.Watermark == nil || resp.SecureSetting.Watermark.Text != "x" {
		t.Errorf("secure_setting = %#v", resp.SecureSetting)
	}
	if len(resp.DocMemberList) != 2 || resp.DocMemberList[1].TmpExternalUserID != "e1" {
		t.Errorf("doc_member_list = %#v", resp.DocMemberList)
	}
	if len(resp.CoAuthList) != 1 || resp.CoAuthList[0].DepartmentID != 1 {
		t.Errorf("co_auth_list = %#v", resp.CoAuthList)
	}
}

func TestWedocGetDocumentResponse(t *testing.T) {
	app, srv := newWedocTestApp(t, func(w http.ResponseWriter, r *http.Request) {
		_, _ = io.WriteString(w, `{
			"errcode":0,"errmsg":"ok","version":3,
			"document":{"begin":0,"end":10,"type":"Document","children":[
				{"begin":0,"end":5,"type":"Paragraph","text":"hello"}
			]}
		}`)
	})
	defer srv.Close()

	resp, err := app.WedocGetDocument(WedocGetDocumentReq{DocID: "DOC1"})
	if err != nil {
		t.Fatal(err)
	}
	if resp.Version != 3 || resp.Document.Type != "Document" || len(resp.Document.Children) != 1 {
		t.Fatalf("resp = %#v", resp)
	}
	if resp.Document.Children[0].Text != "hello" {
		t.Errorf("child = %#v", resp.Document.Children[0])
	}
}

func TestWedocGetDocBaseInfoNested(t *testing.T) {
	app, srv := newWedocTestApp(t, func(w http.ResponseWriter, r *http.Request) {
		_, _ = io.WriteString(w, `{"errcode":0,"errmsg":"ok","doc_base_info":{"docid":"D1","doc_name":"N1","create_time":1,"modify_time":2,"doc_type":3}}`)
	})
	defer srv.Close()

	resp, err := app.WedocGetDocBaseInfo(WedocGetDocBaseInfoReq{DocID: "D1"})
	if err != nil {
		t.Fatal(err)
	}
	if resp.DocBaseInfo == nil || resp.DocBaseInfo.DocName != "N1" || resp.DocBaseInfo.DocType != 3 {
		t.Fatalf("doc_base_info = %#v", resp.DocBaseInfo)
	}
}

func TestWedocUpdateSheetProperties(t *testing.T) {
	app, srv := newWedocTestApp(t, func(w http.ResponseWriter, r *http.Request) {
		var body struct {
			Properties struct {
				SheetID string `json:"sheet_id"`
				Title   string `json:"title"`
			} `json:"properties"`
		}
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			t.Fatal(err)
		}
		if body.Properties.SheetID != "S1" || body.Properties.Title != "T1" {
			t.Errorf("properties = %#v", body.Properties)
		}
		_, _ = io.WriteString(w, `{"errcode":0,"errmsg":"ok"}`)
	})
	defer srv.Close()

	err := app.WedocUpdateSheet(WedocUpdateSheetReq{
		DocID: "D1",
		Properties: WedocSheetProperties{
			SheetID: "S1",
			Title:   "T1",
		},
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestWedocBatchUpdateDocument(t *testing.T) {
	app, srv := newWedocTestApp(t, func(w http.ResponseWriter, r *http.Request) {
		var body struct {
			Requests []struct {
				InsertText *struct {
					Text     string `json:"text"`
					Location *struct {
						Index uint32 `json:"index"`
					} `json:"location"`
				} `json:"insert_text"`
			} `json:"requests"`
		}
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			t.Fatal(err)
		}
		if len(body.Requests) != 1 || body.Requests[0].InsertText == nil {
			t.Fatalf("requests = %#v", body.Requests)
		}
		if body.Requests[0].InsertText.Text != "hello" || body.Requests[0].InsertText.Location.Index != 3 {
			t.Errorf("insert_text = %#v", body.Requests[0].InsertText)
		}
		_, _ = io.WriteString(w, `{"errcode":0,"errmsg":"ok"}`)
	})
	defer srv.Close()

	err := app.WedocBatchUpdateDocument(WedocBatchUpdateDocumentReq{
		DocID: "D1",
		Requests: []WedocDocumentUpdateRequest{{
			InsertText: &WedocInsertText{
				Text:     "hello",
				Location: &WedocLocation{Index: 3},
			},
		}},
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestWedocAddBlocks(t *testing.T) {
	app, srv := newWedocTestApp(t, func(w http.ResponseWriter, r *http.Request) {
		var body struct {
			Blocks []struct {
				Type  string `json:"type"`
				Title string `json:"title"`
				Props struct {
					TextProps struct {
						Align uint32 `json:"align"`
					} `json:"text_props"`
				} `json:"props"`
			} `json:"blocks"`
		}
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			t.Fatal(err)
		}
		if len(body.Blocks) != 1 || body.Blocks[0].Type != "BLOCK_TYPE_TEXT" {
			t.Fatalf("blocks = %#v", body.Blocks)
		}
		if body.Blocks[0].Props.TextProps.Align != 1 {
			t.Errorf("props = %#v", body.Blocks[0].Props)
		}
		_, _ = io.WriteString(w, `{"errcode":0,"errmsg":"ok","blocks":[{"id":"B1","type":"BLOCK_TYPE_TEXT","title":"hello"}]}`)
	})
	defer srv.Close()

	resp, err := app.WedocAddBlocks(WedocAddBlocksReq{
		DocID:  "D1",
		PageID: "P1",
		Blocks: []WedocBlock{{
			Type:  "BLOCK_TYPE_TEXT",
			Title: "hello",
			Props: &WedocBlockProps{TextProps: &WedocTextBlockProps{Align: 1}},
		}},
	})
	if err != nil {
		t.Fatal(err)
	}
	if len(resp.Blocks) != 1 || resp.Blocks[0].ID != "B1" {
		t.Fatalf("resp = %#v", resp)
	}
}

func TestWedocGetSheetPriv(t *testing.T) {
	app, srv := newWedocTestApp(t, func(w http.ResponseWriter, r *http.Request) {
		var body struct {
			Type       uint32   `json:"type"`
			RuleIDList []uint32 `json:"rule_id_list"`
		}
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			t.Fatal(err)
		}
		if body.Type != 2 || len(body.RuleIDList) != 2 || body.RuleIDList[1] != 7 {
			t.Errorf("body = %#v", body)
		}
		_, _ = io.WriteString(w, `{"errcode":0,"errmsg":"ok","rule_list":[{"rule_id":2,"type":2,"name":"n","priv_list":[{"sheet_id":"S1","priv":"2","can_insert_record":true,"field_priv":{"field_range_type":2,"field_rule_list":[{"field_id":"F1","field_type":"FIELD_TYPE_TEXT","can_edit":true,"can_insert":true,"can_view":true}]}}]}]}`)
	})
	defer srv.Close()

	resp, err := app.WedocGetSheetPriv(WedocGetSheetPrivReq{
		DocID:      "D1",
		Type:       2,
		RuleIDList: []uint32{2, 7},
	})
	if err != nil {
		t.Fatal(err)
	}
	if len(resp.RuleList) != 1 || resp.RuleList[0].RuleID != 2 {
		t.Fatalf("resp = %#v", resp)
	}
	priv := resp.RuleList[0].PrivList[0]
	if priv.SheetID != "S1" || priv.FieldPriv == nil || len(priv.FieldPriv.FieldRuleList) != 1 {
		t.Fatalf("priv = %#v", priv)
	}
	if !priv.FieldPriv.FieldRuleList[0].CanEdit {
		t.Errorf("field rule = %#v", priv.FieldPriv.FieldRuleList[0])
	}
}

func TestWedocErrorResponse(t *testing.T) {
	app, srv := newWedocTestApp(t, func(w http.ResponseWriter, r *http.Request) {
		_, _ = io.WriteString(w, `{"errcode":640008,"errmsg":"permission denied"}`)
	})
	defer srv.Close()

	_, err := app.WedocPublishSmartDoc(WedocPublishSmartDocReq{DocID: "DOC1"})
	if err == nil {
		t.Fatal("expected error")
	}
	var clientErr *ClientError
	if !strings.Contains(err.Error(), "640008") {
		t.Fatalf("err = %v", err)
	}
	if ok := errors.As(err, &clientErr); !ok || clientErr.Code != 640008 {
		t.Fatalf("err = %#v", err)
	}
}
