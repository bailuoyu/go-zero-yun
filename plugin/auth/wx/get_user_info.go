package wx

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/pkg/errors"
)

type RespUserInfo struct {
	RespError
	Nickname   string `json:"Nickname"`
	Headimgurl string `json:"headimgurl"`
}

// getAccessToken 根据用户ID获取用户信息
func (w *Wx) GetWxUserInfo(access_token string, openid string) (*RespUserInfo, error) {
	apiURL := "https://api.weixin.qq.com/sns/userinfo"
	data := url.Values{
		"access_token": {access_token},
		"openid":       {openid},
	}
	req, err := http.NewRequest(http.MethodGet, apiURL, nil)
	if err != nil {
		return nil, errors.Wrap(err, "wx userInfo NewRequest error")
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.URL.RawQuery = data.Encode()
	resp, err := w.client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "wx userInfo Do error")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "wx userInfo ReadAll error")
	}
	r := &RespUserInfo{}
	if err := json.Unmarshal(body, r); err != nil {
		return nil, errors.Wrap(err, "wx userInfo Unmarshal error:"+string(body))
	}
	if r.Errmsg != "" && r.ErrorCode != 0 {
		return nil, errors.New("wx userInfo resp error:" + r.Errmsg)
	}
	return r, nil
}
