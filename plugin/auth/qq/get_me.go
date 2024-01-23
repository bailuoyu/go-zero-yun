package qq

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/pkg/errors"
)

type RespMe struct {
	RespError
	Openid string `json:"openid"`
}

// UsersShow 根据用户ID获取用户信息
func (q *Qq) GetMe(access_token string) (*RespMe, error) {
	apiURL := "https://graph.qq.com/oauth2.0/me"
	data := url.Values{
		"access_token": {access_token},
		"fmt":          {"json"},
	}
	req, err := http.NewRequest(http.MethodGet, apiURL, nil)
	if err != nil {
		return nil, errors.Wrap(err, "qq access_token NewRequest error")
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.URL.RawQuery = data.Encode()
	resp, err := q.client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "qq Openid Do error")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "qq Openid ReadAll error")
	}
	r := &RespMe{}
	if err := json.Unmarshal(body, r); err != nil {
		return nil, errors.Wrap(err, "qq Openid Unmarshal error:"+string(body))
	}

	if r.Error != "" && r.ErrorCode != 0 {
		return nil, errors.New("qq Openid resp error:" + r.Error)
	}
	if r.Openid == "" {
		return nil, errors.New("qq Openid resp error:" + r.Error)
	}
	return r, nil
}
