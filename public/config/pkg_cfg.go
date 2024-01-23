package config

type PkgCfg struct {
	Jwt    JwtCfg    `json:"Jwt,optional""`
	Sts    StsCfg    `json:"Sts,optional"`
	WxWork WxWorkCfg `json:"WxWork,optional"`
}

type JwtCfg struct {
	AccessSecret string `json:"AccessSecret"`
	AccessExpire int64  `json:"AccessExpire"`
}

type StsCfg struct {
	Host   string `json:"Host,optional"`
	Scheme string `json:"Scheme,optional"`
}

type WxWorkCfg struct {
	AppChat struct {
		WarningId string `json:"WarningId"`
	} `json:"AppChat"`
}
