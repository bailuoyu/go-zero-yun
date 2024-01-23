package im

type (
	TIMCustomElem struct {
		Data  string `json:"Data"` //自定义消息数据。 不作为 APNs 的 payload 字段下发，故从 payload 中无法获取 Data 字段。
		Desc  string `json:"Desc"` //自定义消息描述信息。当接收方为 iOS 或 Android 后台在线时，做离线推送文本展示。
		Ext   string `json:"Ext,omitempty"`
		Sound string `json:"Sound,omitempty"`
	}

	SendMsgBody struct {
		MsgType    string      `json:"MsgType"`
		MsgContent interface{} `json:"MsgContent"`
	}
	OfflinePushInfo struct {
		PushFlag int `json:"PushFlag"` //0表示推送  1：表示不离线推送
	}

	SendMsgReq struct {
		SyncOtherMachine        int             `json:"SyncOtherMachine"` // 1：把消息同步到 From_Account 在线终端和漫游上 2：消息不同步至 From_Account
		FromAccount             string          `json:"From_Account,omitempty"`
		ToAccount               string          `json:"To_Account"`
		MsgLifeTime             int             `json:"MsgLifeTime"`
		MsgSeq                  int             `json:"MsgSeq"`
		MsgRandom               int             `json:"MsgRandom"`
		MsgBody                 []SendMsgBody   `json:"MsgBody"`
		CloudCustomData         string          `json:"CloudCustomData"` //消息自定义数据（云端保存，会发送到对端，程序卸载重装后还能拉取到）
		SupportMessageExtension int             `json:"SupportMessageExtension"`
		SendMsgControl          []string        `json:"SendMsgControl"`
		ForbidCallbackControl   []string        `json:"ForbidCallbackControl"`
		IsNeedReadReceipt       int             `json:"IsNeedReadReceipt"` //该条消息是否需要已读回执，0为不需要，1为需要，默认为0
		OfflinePushInfo         OfflinePushInfo `json:"OfflinePushInfo,omitempty"`
	}
	SendMsgRsp struct {
		Rsp
		MsgTime int    `json:"MsgTime"`
		MsgKey  string `json:"MsgKey"`
	}

	BatchSendMsgReq struct {
		SyncOtherMachine        int             `json:"SyncOtherMachine"` // 1：把消息同步到 From_Account 在线终端和漫游上 2：消息不同步至 From_Account
		FromAccount             string          `json:"From_Account,omitempty"`
		ToAccount               []string        `json:"To_Account"`
		MsgLifeTime             int             `json:"MsgLifeTime"`
		MsgSeq                  int             `json:"MsgSeq"`
		MsgRandom               int             `json:"MsgRandom"`
		MsgBody                 []SendMsgBody   `json:"MsgBody"`
		CloudCustomData         string          `json:"CloudCustomData"` //消息自定义数据（云端保存，会发送到对端，程序卸载重装后还能拉取到）
		SupportMessageExtension int             `json:"SupportMessageExtension"`
		SendMsgControl          []string        `json:"SendMsgControl"`
		IsNeedReadReceipt       int             `json:"IsNeedReadReceipt"` //该条消息是否需要已读回执，0为不需要，1为需要，默认为0
		OfflinePushInfo         OfflinePushInfo `json:"OfflinePushInfo,omitempty"`
	}

	BatchSendMsgRsp struct {
		ActionStatus string `json:"ActionStatus"`
		ErrorCode    int    `json:"ErrorCode"`
		ErrorInfo    string `json:"ErrorInfo"`
		MsgKey       string `json:"MsgKey"`
		ErrorList    []struct {
			ToAccount string `json:"To_Account"`
			ErrorCode int    `json:"ErrorCode"`
		} `json:"ErrorList"`
	}
)

// SendMsg 单发单聊消息 https://cloud.tencent.com/document/product/269/2282
func (clm *ClientIM) SendMsg(req SendMsgReq) (rsp SendMsgRsp, err error) {
	err = clm.post("openim/sendmsg", req, &rsp)
	return
}

// BatchSendMsg 批量发单聊消息 https://cloud.tencent.com/document/product/269/1612
func (clm *ClientIM) BatchSendMsg(req BatchSendMsgReq) (rsp BatchSendMsgRsp, err error) {
	err = clm.post("openim/batchsendmsg", req, &rsp)
	return
}
