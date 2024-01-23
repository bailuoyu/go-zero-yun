package im

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"

	"go-zero-yun/pkg/logkit"
)

type Rsp struct {
	ActionStatus string `json:"ActionStatus"`
	ErrorInfo    string `json:"ErrorInfo"`
	ErrorCode    int    `json:"ErrorCode"`
}

func (clm *ClientIM) getUrl(suf string) string {
	return fmt.Sprintf("https://%s/v4/%s", clm.Config.ApiDomain, suf)
}

func (clm *ClientIM) post(suf string, rq interface{}, rs interface{}) error {
	random, _ := rand.Int(rand.Reader, big.NewInt(4294967295))
	rawUrl := fmt.Sprintf("%s?sdkappid=%d&identifier=%s&usersig=%s&random=%d&contenttype=json",
		clm.getUrl(suf), clm.Config.SdkAppId, clm.Config.AdminId, clm.Sign, random)

	request := clm.ReqC.R().SetContext(clm.ctx)
	response, err := request.SetBodyJsonMarshal(rq).Post(rawUrl)
	//response, err := request.SetBodyJsonString(`{"SyncOtherMachine":2,"To_Account":"1573624","MsgLifeTime":60,"MsgSeq":93847636,"MsgRandom":1287657,"MsgBody":[],"CloudCustomData":"your cloud custom data","SupportMessageExtension":0}`).Post(rawUrl)
	//请求错误则直接报错
	if err != nil {
		return err
	}
	//解析通用返回,多一次判断，少一行代码
	var baseRsp Rsp
	err = response.UnmarshalJson(&baseRsp)
	if err != nil {
		return err
	}
	err = response.UnmarshalJson(rs)
	if err != nil {
		return err
	}
	if baseRsp.ErrorCode != 0 {
		logkit.WithType(pluginType).Errorf(clm.ctx, "%+v", baseRsp)
		err = errors.New(baseRsp.ErrorInfo)
		return err
	}
	return err
}

func (clm *ClientIM) getRequestUrl(suf string) string {
	random, _ := rand.Int(rand.Reader, big.NewInt(4294967295))
	rawUrl := fmt.Sprintf("%s?sdkappid=%d&identifier=%s&usersig=%s&random=%d&contenttype=json",
		clm.getUrl(suf), clm.Config.SdkAppId, clm.Config.AdminId, clm.Sign, random)
	return rawUrl
}
