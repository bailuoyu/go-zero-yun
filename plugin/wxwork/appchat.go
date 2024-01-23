package wxwork

const MsgTypeText = "text"
const MsgTypeImage = "image"
const MsgTypeVoice = "voice"
const MsgTypeVideo = "video"
const MsgTypeFile = "file"
const MsgTypeTextCard = "textcard"
const MsgTypeNews = "news"
const MsgTypeMpNews = "mpnews"
const MsgTypeMarkdown = "markdown"

// AppChatCreateReq 创建群聊会话
type AppChatCreateReq struct {
	Name     string   `json:"name,omitempty"`
	Owner    string   `json:"owner,omitempty"`
	Userlist []string `json:"userlist"`
	Chatid   string   `json:"chatid,omitempty"`
}

// AppChatCreateRsp 创建群聊会话
type AppChatCreateRsp struct {
	Rsp
	Chatid string `json:"chatid"`
}

// AppChatCreate 创建群聊会话 https://developer.work.weixin.qq.com/document/path/90245
func (agent *Agent) AppChatCreate(rq AppChatCreateReq) (AppChatCreateRsp, error) {
	var rs AppChatCreateRsp
	err := agent.reqPost("appchat/create", rq, &rs)
	return rs, err
}

// AppChatUpdateReq 修改群聊会话
type AppChatUpdateReq struct {
	Chatid      string   `json:"chatid"`
	Name        string   `json:"name,omitempty"`
	Owner       string   `json:"owner,omitempty"`
	AddUserList []string `json:"add_user_list,omitempty"`
	DelUserList []string `json:"del_user_list,omitempty"`
}

// AppChatUpdate 修改群聊会话 https://developer.work.weixin.qq.com/document/path/90246
func (agent *Agent) AppChatUpdate(rq AppChatCreateReq) (Rsp, error) {
	var rs Rsp
	err := agent.reqPost("appchat/update", rq, &rs)
	return rs, err
}

// AppChatGetRsp 获取群聊会话
type AppChatGetRsp struct {
	Rsp
	ChatInfo struct {
		Chatid   string   `json:"chatid"`
		Name     string   `json:"name"`
		Owner    string   `json:"owner"`
		Userlist []string `json:"userlist"`
	} `json:"chat_info"`
}

// AppChatGet 获取群聊会话 https://developer.work.weixin.qq.com/document/path/90247
func (agent *Agent) AppChatGet(chatid string) (AppChatGetRsp, error) {
	mp := map[string]string{
		"chatid": chatid,
	}
	var rs AppChatGetRsp
	err := agent.reqGet("appchat/get", mp, &rs)
	return rs, err
}

/**
 * 应用推送消息 https://developer.work.weixin.qq.com/document/path/90248
 */

// AppChatSendReq 应用推送消息
type AppChatSendReq struct {
	Chatid  string                 `json:"chatid"`
	Msgtype string                 `json:"msgtype"`
	Msg     map[string]interface{} `json:"msg"`
}

// AppChatSend 应用推送消息 https://developer.work.weixin.qq.com/document/path/90248
func (agent *Agent) AppChatSend(rq AppChatSendReq) (Rsp, error) {
	var rs Rsp
	rqMp := map[string]interface{}{
		"chatid":   rq.Chatid,
		"msgtype":  rq.Msgtype,
		rq.Msgtype: rq.Msg,
	}
	err := agent.reqPost("appchat/send", rqMp, &rs)
	return rs, err
}

// AppChatSendText 应用推送消息text https://developer.work.weixin.qq.com/document/path/90248
func (agent *Agent) AppChatSendText(chatid string, content string) (Rsp, error) {
	rq := AppChatSendReq{
		Chatid:  chatid,
		Msgtype: MsgTypeText,
		Msg: map[string]interface{}{
			"content": content,
		},
	}
	rs, err := agent.AppChatSend(rq)
	return rs, err
}

// AppChatSendMarkdown 应用推送消息markdown https://developer.work.weixin.qq.com/document/path/90248
func (agent *Agent) AppChatSendMarkdown(chatid string, content string) (Rsp, error) {
	rq := AppChatSendReq{
		Chatid:  chatid,
		Msgtype: MsgTypeMarkdown,
		Msg: map[string]interface{}{
			"content": content,
		},
	}
	rs, err := agent.AppChatSend(rq)
	return rs, err
}
