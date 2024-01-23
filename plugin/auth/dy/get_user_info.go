package dy

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/pkg/errors"
)

type UserInfo struct {
	Nickname    string `json:"nickname"`
	Avatar      string `json:"avatar"`
	ErrCode     int    `json:"error_code"`
	Description string `json:"description"`
}
type RespUserInfo struct {
	UserInfo UserInfo `json:"data"`
}

// getAccessToken 根据用户ID获取用户信息
func (d *Dy) GetDyUserInfo(access_token string, openid string) (*RespUserInfo, error) {
	apiURL := "https://open.douyin.com/oauth/userinfo/"
	data := url.Values{
		"access_token": {access_token},
		"open_id":      {openid},
	}
	req, err := http.NewRequest(http.MethodPost, apiURL, nil)
	if err != nil {
		return nil, errors.Wrap(err, "dy user_info NewRequest error")
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.URL.RawQuery = data.Encode()
	resp, err := d.client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "dy user_info Do error")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "dy user_info ReadAll error")
	}
	r := &RespUserInfo{}
	if err := json.Unmarshal(body, r); err != nil {
		return nil, errors.Wrap(err, "dy user_info Unmarshal error:"+string(body))
	}
	if r.UserInfo.Description != "" && r.UserInfo.ErrCode != 0 {
		return nil, errors.New("dy user_info resp error:" + r.UserInfo.Description)
	}
	return r, nil
}
