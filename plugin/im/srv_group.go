package im

type (
	CreateGroupReq struct {
		GroupId string `json:"GroupId"`
		Type    string `json:"Type"` // 群组类型：Private/Public/ChatRoom/AVChatRoom/Community
		Name    string `json:"Name"`
	}

	CreateGroupRsp struct {
		Rsp
		GroupId string `json:"GroupId"`
	}

	DestroyGroupReq struct {
		GroupId string `json:"GroupId"`
	}
	DestroyGroupRsp struct {
		Rsp
		GroupId string `json:"GroupId"`
	}

	GroupNotifyReq struct {
		GroupId string `json:"GroupId"`
		Content string `json:"Content"`
	}
)

// CreateGroup 创建组群  https://cloud.tencent.com/document/product/269/1615
func (clm *ClientIM) CreateGroup(req CreateGroupReq) (rsp DestroyGroupRsp, err error) {
	err = clm.post("group_open_http_svc/create_group", req, &rsp)
	return
}

// DestroyGroup 解散组群 https://cloud.tencent.com/document/product/269/1624
func (clm *ClientIM) DestroyGroup(req DestroyGroupReq) (rsp Rsp, err error) {
	err = clm.post("group_open_http_svc/destroy_group", req, &rsp)
	return
}

// SendGroupNotify 在群组中发送系统通知 https://cloud.tencent.com/document/product/269/1630
func (clm *ClientIM) SendGroupNotify(req GroupNotifyReq) (rsp Rsp, err error) {
	err = clm.post("group_open_http_svc/send_group_system_notification", req, &rsp)
	return
}

func (clm *ClientIM) GetRequestUrl(suf string) string {
	return clm.getRequestUrl(suf)
}
