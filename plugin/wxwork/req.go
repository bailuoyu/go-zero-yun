package wxwork

import (
	"fmt"
	"github.com/imroc/req/v3"
)

const apiHost = "https://qyapi.weixin.qq.com"

type Rsp struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

func (agent *Agent) getUrl(suf string) string {
	return fmt.Sprintf("%s/cgi-bin/%s", apiHost, suf)
}

func (agent *Agent) reqGet(suf string, mp map[string]string, rs interface{}) error {
	url := agent.getUrl(suf)
	request := agent.ReqC.Get(url)
	if agent.Token.AccessToken != "" {
		agent.setReqToken(request)
	}
	for k, v := range mp {
		request.SetQueryParam(k, v)
	}
	response := request.Do(agent.ctx)
	//解析通用返回,多一次判断，少一行代码
	err := agent.processRsp(request, response)
	if err != nil {
		return err
	}
	err = response.UnmarshalJson(rs)
	return err
}

func (agent *Agent) reqPost(suf string, rq interface{}, rs interface{}) error {
	url := agent.getUrl(suf)
	request := agent.ReqC.Post(url)
	if agent.Token.AccessToken != "" {
		agent.setReqToken(request)
	}
	request = request.SetBodyJsonMarshal(rq)
	response := request.Do(agent.ctx)
	if response.Err != nil {
		return response.Err
	}
	//解析通用返回,多一次判断，少一行代码
	err := agent.processRsp(request, response)
	if err != nil {
		return err
	}
	err = response.UnmarshalJson(rs)
	return err
}

func (agent *Agent) setReqToken(request *req.Request) {
	request.AddQueryParam("access_token", agent.Token.AccessToken)
}

func (agent *Agent) processRsp(request *req.Request, response *req.Response) error {
	var rs Rsp
	err := response.UnmarshalJson(&rs)
	if err != nil {
		return err
	}
	if rs.Errcode == 0 {
		return nil
	}
	// 如果token失效,而且未重试过
	//if (rs.Errcode == 40014 || rs.Errcode == 42001) {
	//}
	err = fmt.Errorf("wx_work %s response error:%s , code:%d", request.RawURL, rs.Errmsg, rs.Errcode)
	return err
}
