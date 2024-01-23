// https://open.weibo.com/wiki/Oauth2/access_token
// 请求参数：
//    client_id	true	string	申请应用时分配的AppKey。
//    client_secret	true	string	申请应用时分配的AppSecret。
//    grant_type	true	string	请求的类型，填写authorization_code
// grant_type为authorization_code时:
//    code	true	string	调用authorize获得的code值。
//    redirect_uri	true	string	回调地址，需需与注册应用里的回调地址一致。
// 返回数据：
//    access_token	string	用户授权的唯一票据，用于调用微博的开放接口，同时也是第三方应用验证微博用户登录的唯一票据，第三方应用应该用该票据和自己应用内的用户建立唯一影射关系，来识别登录状态，不能使用本返回值里的UID字段来做登录识别。
//    expires_in	string	access_token的生命周期，单位是秒数。
//    remind_in	string	access_token的生命周期（该参数即将废弃，开发者请使用expires_in）。
//    uid	string	授权用户的UID，本字段只是为了方便开发者，减少一次user/show接口调用而返回的，第三方应用不能用此字段作为用户登录状态的识别，只有access_token才是用户授权的唯一票据。

package weibo

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/pkg/errors"
)

// RespToken 获取 access token 接口的返回结果
type RespAuth struct {
	OpenId   string `json:"open_id"`
	NickName string `json:"nickname"`
	Face     string `json:"face"`
}

func (w *Weibo) GetUserInfo() (*RespAuth, error) {
	tokenInfo, err := w.TokenInfo(w.access_token)
	if err != nil {
		return nil, errors.New("获取token失败")
	}
	// if tokenInfo.Appkey != "11111" {
	// 	return nil, errors.New("appkey错误")
	// }
	userShow, err := w.UsersShow(w.access_token, tokenInfo.UID, "")
	if err != nil || userShow.ID == 0 {
		return nil, errors.New("获取基本信息失败")
	}
	num := fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
	nickname := "喵星人" + num
	face := "default.png"
	if userShow.Name != "" {
		nickname = userShow.Name
	}
	if userShow.AvatarLarge != "" {
		face = userShow.AvatarLarge
	}
	openId := strconv.Itoa(userShow.ID)
	return &RespAuth{
		OpenId:   openId,
		NickName: nickname,
		Face:     face,
	}, nil
}
