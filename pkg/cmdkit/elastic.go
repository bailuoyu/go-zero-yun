package cmdkit

import (
	"errors"
	"fmt"
	"github.com/olivere/elastic/v7"
	"go-zero-yun/pkg/logkit"
	pconf "go-zero-yun/public/config"
)

// initElastic 初始化elastic
func initElastic(check bool) {
	pconf.ClientCfg.Elastic = make(map[string]*pconf.ClientElasticConfig)
	for _, v := range pconf.Cfg.Client.Elastic {
		pconf.ClientCfg.Elastic[v.Name] = &pconf.ClientElasticConfig{
			ElasticCfg: v,
		}
		if check && !IsEnvLocal() {
			_, err := ElasticConnect(pconf.ClientCfg.Elastic[v.Name])
			if err != nil {
				panic(errors.New(fmt.Sprintf("elastic fatal %s, name: %s", err.Error(), v.Name)))
			}
		}
	}
}

// ElasticConnect 获取elastic连接
func ElasticConnect(v *pconf.ClientElasticConfig) (*elastic.Client, error) {
	//注入日志
	doer := logkit.GetEsDoer()
	//自定义请求方法
	client, err := elastic.NewClient(
		elastic.SetURL(v.Urls...),                    //连接地址
		elastic.SetBasicAuth(v.Username, v.Password), //账号密码
		elastic.SetSniff(v.Sniff),                    //是否开启嗅探
		elastic.SetHttpClient(doer),
	)
	if err != nil {
		return client, err
	}
	v.Client = client
	// ping一次连接确保能连接上
	//ctx := context.Background()
	//for _, url := range v.Urls {
	//	_, code, pErr := client.Ping(url).Do(ctx)
	//	if pErr != nil {
	//		return client, pErr
	//	}
	//	if code != http.StatusOK {
	//		return client, errors.New(fmt.Sprintf("ping expected status code = %d; got %d", http.StatusOK, code))
	//	}
	//}
	return client, err
}
