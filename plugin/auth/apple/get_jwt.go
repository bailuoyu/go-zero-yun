package apple

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/pkg/errors"
)

type Keys struct {
	Kty string `json:"kty"`
	Kid string `json:"kid"`
	Kse string `json:"use"`
	Alg string `json:"alg"`
	N   string `json:"n"`
	E   string `json:"e"`
}
type RespJwt struct {
	Message string `json:"message"`
	Keys    []Keys `json:"keys"`
}

// getAccessToken 根据用户ID获取用户信息
func (a *Apple) GetJwt(code string) (*RespJwt, error) {
	data := url.Values{
		"code": {code},
	}

	req, err := http.NewRequest(http.MethodGet, a.cfg.KeysUrl, nil)
	if err != nil {
		return nil, errors.Wrap(err, "apple jwt NewRequest error")
	}
	req.Header.Set("Content-Type", "aapplication/json")
	req.URL.RawQuery = data.Encode()
	resp, err := a.client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "apple jwt Do error")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, errors.Wrap(err, "apple jwt ReadAll error")
	}
	r := &RespJwt{}
	if err := json.Unmarshal(body, r); err != nil {
		return nil, errors.Wrap(err, "apple jwt Unmarshal error:"+string(body))
	}
	return r, nil
}
