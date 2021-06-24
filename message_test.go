package tim

import (
	"testing"
	"time"
)

func TestBatchSendTextMsg(t *testing.T) {
	toaccount := []string{}
	toaccount = append(toaccount, "10")
	toaccount = append(toaccount, "go1")
	resp, err := testTimAPP.BatchSendTextMsg("qa_admin", toaccount, "hello", SyncToFrom)
	if err != nil {
		t.Error(err)
	}
	if resp.ActionStatus == ResponseFail {
		t.Error(resp.ErrorInfo)
	}
}
func TestSendMsg(t *testing.T) {
	textMsg := MsgElement{
		MsgType: TIMText,
		MsgContent: MsgText{
			Text: "1",
		},
	}
	faceMsg := MsgElement{
		MsgType: TIMFace,
		MsgContent: MsgFace{
			Index: 1,
			Data:  "表情",
		},
	}
	locMsg := MsgElement{
		MsgType: TIMLocation,
		MsgContent: MsgLocation{
			Latitude:  31.9894418379,
			Longitude: 110.6542968750,
		},
	}
	customMsg := MsgElement{
		MsgType: TIMCustom,
		MsgContent: MsgCustom{
			Desc: "自定义消息",
			Data: map[string]string{
				"Key": "Value",
			},
		},
	}

	param := SendMsgData{
		FromAccount: "goTest",
		ToAccount:   "10",
		MsgBody:     []MsgElement{locMsg, textMsg, faceMsg, customMsg},
	}
	resp, err := testTimAPP.SendMsg(param)
	if err != nil {
		t.Error(err)
	}
	if resp.ActionStatus == ResponseFail {
		t.Error(resp.ErrorInfo, resp.ErrorCode)
	}
}

func TestImportMsg(t *testing.T) {
	textMsg := MsgElement{
		MsgType: TIMText,
		MsgContent: MsgText{
			Text: "历史消息",
		},
	}
	param := SendMsgData{
		SyncFromOldSystem: SyncNotCount,
		FromAccount:       "goTest",
		ToAccount:         "10",
		MsgTimeStamp:      time.Date(2018, 7, 20, 12, 3, 35, 0, time.UTC).Unix(),
		MsgBody:           []MsgElement{textMsg},
	}
	resp, err := testTimAPP.ImportMsg(param)
	if err != nil {
		t.Error(err)
	}
	if resp.ActionStatus == ResponseFail {
		t.Error(resp.ErrorInfo)
	}
}
