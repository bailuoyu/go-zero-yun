package wx

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/pkg/errors"
)

type RespAccessToken struct {
	RespError
	AcceccToken string `json:"access_token"`
	Openid      string `json:"openid"`
}

// getAccessToken 根据用户ID获取用户信息
func (w *Wx) GetAccessToken(code string) (*RespAccessToken, error) {
	fmt.Printf("%+v", w.cfg)
	apiURL := "https://api.weixin.qq.com/sns/oauth2/access_token"
	data := url.Values{
		"code":       {code},
		"grant_type": {"authorization_code"},
		"appid":      {w.cfg.Appid},
		"secret":     {w.cfg.Secret},
		"fmt":        {"json"},
	}
	req, err := http.NewRequest(http.MethodGet, apiURL, nil)
	if err != nil {
		return nil, errors.Wrap(err, "wx access_token NewRequest error")
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.URL.RawQuery = data.Encode()
	resp, err := w.client.Do(req)

	if err != nil {
		return nil, errors.Wrap(err, "wx access_token Do error")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "wx access_token ReadAll error")
	}
	r := &RespAccessToken{}
	if err := json.Unmarshal(body, r); err != nil {
		return nil, errors.Wrap(err, "wx access_token Unmarshal error:"+string(body))
	}
	if r.Errmsg != "" && r.ErrorCode != 0 {
		return nil, errors.New("wx access_token resp error:" + r.Errmsg)
	}
	if r.AcceccToken == "" || r.Openid == "" {
		return nil, errors.New("获取token失败")
	}
	return r, nil
}
