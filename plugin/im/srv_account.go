package im

type (
	QueryOnlineStatusReq struct {
		IsNeedDetail int      `json:"IsNeedDetail"`
		ToAccount    []string `json:"To_Account"`
	}
	QueryOnlineStatusRsp struct {
		Rsp
		QueryResult []struct {
			ToAccount string `json:"To_Account"`
			Status    string `json:"Status"`
			Detail    []struct {
				Platform string `json:"Platform"`
				Status   string `json:"Status"`
			} `json:"Detail,omitempty"`
		} `json:"QueryResult"`
		ErrorList []struct {
			ToAccount string `json:"To_Account"`
			ErrorCode int    `json:"ErrorCode"`
		} `json:"ErrorList"`
	}

	CreateAccountReq struct {
		UserID  string `json:"UserID"`
		Nick    string `json:"Nick"`
		FaceUrl string `json:"FaceUrl"`
	}
)

// AccountOnlineStatus 查询用户在线状态 https://cloud.tencent.com/document/product/269/2566
func (clm *ClientIM) AccountOnlineStatus(req QueryOnlineStatusReq) (rsp QueryOnlineStatusRsp, err error) {
	err = clm.post("openim/query_online_status", req, &rsp)
	return
}

func (clm *ClientIM) CreateAccount(req CreateAccountReq) (rsp Rsp, err error) {
	err = clm.post("im_open_login_svc/account_import", req, &rsp)
	return
}
