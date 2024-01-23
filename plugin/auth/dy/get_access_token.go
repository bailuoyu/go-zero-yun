package dy

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/pkg/errors"
)

type Data struct {
	AccessToken string `json:"access_token"`
	OpenId      string `json:"open_id"`
	Description string `json:"description"`
	ErrCode     int    `json:"error_code"`
}
type RespAccessToken struct {
	Message string `json:"message"`
	Data    Data   `json:"data"`
}

// getAccessToken 根据用户ID获取用户信息
func (d *Dy) GetAccessToken(code string) (*RespAccessToken, error) {
	apiURL := "https://open.douyin.com/oauth/access_token/"
	data := url.Values{
		"code":          {code},
		"client_secret": {d.cfg.ClientSecret},
		"client_key":    {d.cfg.ClientKey},
		"grant_type":    {"authorization_code"},
	}

	req, err := http.NewRequest(http.MethodPost, apiURL, nil)
	if err != nil {
		return nil, errors.Wrap(err, "dy access_token NewRequest error")
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.URL.RawQuery = data.Encode()
	resp, err := d.client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "dy access_token Do error")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, errors.Wrap(err, "dy access_token ReadAll error")
	}
	r := &RespAccessToken{}
	if err := json.Unmarshal(body, r); err != nil {
		return nil, errors.Wrap(err, "dy access_token Unmarshal error:"+string(body))
	}
	if r.Data.Description != "" && r.Data.ErrCode != 0 {
		return nil, errors.New("dy access_token resp error:" + r.Data.Description)
	}
	return r, nil
}
