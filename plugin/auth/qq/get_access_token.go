package qq

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/pkg/errors"
)

type RespAccessToken struct {
	RespError
	AccessToken string `json:"access_token"`
}

// getAccessToken 根据用户ID获取用户信息
func (q *Qq) GetAccessToken(code string) (*RespAccessToken, error) {
	apiURL := "https://graph.qq.com/oauth2.0/token"
	data := url.Values{
		"code":          {code},
		"grant_type":    {"authorization_code"},
		"client_id":     {q.cfg.ClientId},
		"client_secret": {q.cfg.ClientSecret},
		"redirect_uri":  {q.cfg.RedirectUri},
		"fmt":           {"json"},
	}
	req, err := http.NewRequest(http.MethodGet, apiURL, nil)
	if err != nil {
		return nil, errors.Wrap(err, "qq access_token NewRequest error")
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.URL.RawQuery = data.Encode()
	resp, err := q.client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "qq access_token Do error")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "qq access_token ReadAll error")
	}
	r := &RespAccessToken{}
	if err := json.Unmarshal(body, r); err != nil {
		return nil, errors.Wrap(err, "qq access_token Unmarshal error:"+string(body))
	}
	if r.Error != "" && r.ErrorCode != 0 {
		return nil, errors.New("qq access_token resp error:" + r.Error)
	}
	return r, nil
}
