package wx

import (
	"github.com/pkg/errors"
)

// RespToken 获取 access token 接口的返回结果
type RespAuth struct {
	OpenId   string `json:"open_id"`
	NickName string `json:"nickname"`
	Face     string `json:"face"`
}

func (w *Wx) GetUserInfo() (*RespAuth, error) {

	accessInfo, err := w.GetAccessToken(w.code)
	if err != nil || accessInfo.Openid == "" {
		return nil, errors.New(err.Error())
	}
	userInfo, err := w.GetWxUserInfo(accessInfo.AcceccToken, accessInfo.Openid)
	if err != nil || userInfo.Nickname == "" {
		return nil, errors.New(err.Error())
	}
	return &RespAuth{
		OpenId:   accessInfo.Openid,
		NickName: userInfo.Nickname,
		Face:     userInfo.Headimgurl,
	}, nil
}
