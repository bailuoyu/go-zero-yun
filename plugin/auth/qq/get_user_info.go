package qq

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/pkg/errors"
)

type RespUserInfo struct {
	RespError
	FigureurlQq2 string `json:"figureurl_qq_2"`
	Nickname     string `json:"nickname"`
}

// getAccessToken 根据用户ID获取用户信息
func (q *Qq) GetQqUserInfo(access_token string, openid string) (*RespUserInfo, error) {
	apiURL := "https://graph.qq.com/user/get_user_info"
	data := url.Values{
		"access_token":       {access_token},
		"openid":             {openid},
		"oauth_consumer_key": {q.cfg.ClientId},
		"fmt":                {"json"},
	}
	req, err := http.NewRequest(http.MethodGet, apiURL, nil)
	if err != nil {
		return nil, errors.Wrap(err, "qq access_token NewRequest error")
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.URL.RawQuery = data.Encode()
	resp, err := q.client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "qq user_info Do error")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "qq user_info ReadAll error")
	}
	r := &RespUserInfo{}
	if err := json.Unmarshal(body, r); err != nil {
		return nil, errors.Wrap(err, "qq user_info Unmarshal error:"+string(body))
	}

	if r.Error != "" && r.ErrorCode != 0 {
		return nil, errors.New("qq user_info resp error:" + r.Error)
	}
	return r, nil
}
