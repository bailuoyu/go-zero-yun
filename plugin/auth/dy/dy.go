package dy

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/pkg/errors"
)

// RespToken 获取 access token 接口的返回结果
type RespAuth struct {
	OpenId   string `json:"open_id"`
	NickName string `json:"nickname"`
	Face     string `json:"face"`
}

func (d *Dy) GetUserInfo() (*RespAuth, error) {
	accessInfo, err := d.GetAccessToken(d.code)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	userInfo, err := d.GetDyUserInfo(accessInfo.Data.AccessToken, accessInfo.Data.OpenId)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	num := fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
	nickname := "喵星人" + num
	face := "default.png"
	if userInfo.UserInfo.Nickname != "" {
		nickname = userInfo.UserInfo.Nickname
	}
	if userInfo.UserInfo.Avatar != "" {
		face = userInfo.UserInfo.Avatar
	}
	return &RespAuth{
		OpenId:   accessInfo.Data.OpenId,
		NickName: nickname,
		Face:     face,
	}, nil
}
