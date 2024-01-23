package qq

import (
	"github.com/pkg/errors"
)

// RespToken 获取 access token 接口的返回结果
type RespAuth struct {
	OpenId   string `json:"open_id"`
	NickName string `json:"nickname"`
	Face     string `json:"face"`
}

func (q *Qq) GetUserInfo() (*RespAuth, error) {
	accessTokenInfo, err := q.GetAccessToken(q.code)
	if err != nil || accessTokenInfo.AccessToken == "" {
		return nil, errors.New(err.Error())
	}
	meInfo, err := q.GetMe(accessTokenInfo.AccessToken)
	if err != nil || meInfo.Openid == "" {
		return nil, errors.New(err.Error())
	}
	userInfo, err := q.GetQqUserInfo(accessTokenInfo.AccessToken, meInfo.Openid)
	if err != nil || userInfo.Nickname == "" {
		return nil, errors.New(err.Error())
	}
	return &RespAuth{
		OpenId:   meInfo.Openid,
		NickName: userInfo.Nickname,
		Face:     userInfo.FigureurlQq2,
	}, nil
}
