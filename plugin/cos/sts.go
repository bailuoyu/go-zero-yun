package cos

import (
	"fmt"
	"github.com/tencentyun/qcloud-cos-sts-sdk/go"
	"go-zero-yun/plugin"
	conf "go-zero-yun/public/config"
	"time"
)

// GetSts 获取sts授权
func GetSts(opt *sts.CredentialOptions) (*sts.CredentialResult, Config, error) {
	return GetStsByName(plugin.DefaultName, opt)
}

// GetStsByName 获取sts授权
func GetStsByName(name string, opt *sts.CredentialOptions) (*sts.CredentialResult, Config, error) {
	cosCfg := GetCfgByName(name)
	client := sts.NewClient(
		// 通过环境变量获取密钥
		cosCfg.SecretId,
		cosCfg.SecretKey,
		nil,
		// sts.Host("sts.internal.tencentcloudapi.com"), // 设置域名, 默认域名sts.tencentcloudapi.com
		// sts.Scheme("http"),      // 设置协议, 默认为https，公有云sts获取临时密钥不允许走http，特殊场景才需要设置http
	)
	if conf.Cfg.Pkg.Sts.Host != "" {
		client.SetHost(conf.Cfg.Pkg.Sts.Host)
	}
	if conf.Cfg.Pkg.Sts.Scheme != "" {
		client.SetScheme(conf.Cfg.Pkg.Sts.Scheme)
	}
	res, err := client.GetCredential(opt)
	return res, cosCfg, err
}

// SimpleSts 简单获取sts授权
func SimpleSts(uid string) (*sts.CredentialResult, Config, error) {
	cosCfg := GetCfgByName(plugin.DefaultName)
	// 资源授权前缀，分app目录
	resourcePre := fmt.Sprintf("qcs::cos:%s:uid/%d:%s/%s", cosCfg.Region, cosCfg.AppId, cosCfg.Bucket, conf.Cfg.Server.App)
	// 策略概述 https://cloud.tencent.com/document/product/436/18023
	// 策略生产器 https://console.cloud.tencent.com/cam/policy
	opt := &sts.CredentialOptions{
		DurationSeconds: int64(2 * time.Hour.Seconds()), //有效时间两个小时
		Region:          cosCfg.Region,
		Policy: &sts.CredentialPolicy{
			Statement: []sts.CredentialPolicyStatement{
				{
					// 密钥的权限列表。简单上传和分片需要以下的权限，其他权限列表请看 https://cloud.tencent.com/document/product/436/31923
					Action: []string{
						// 简单上传
						"name/cos:PostObject",
						// 表单上传
						"name/cos:PutObject",
						// 分片上传
						"name/cos:InitiateMultipartUpload",
						"name/cos:ListMultipartUploads",
						"name/cos:ListParts",
						"name/cos:UploadPart",
						"name/cos:CompleteMultipartUpload",
						// 下载操作
						"name/cos:GetObject",
						//"name/cos:GetObjectUrl",	// 获取下载链接,已废弃
						// 标签操作
						"name/cos:PutObjectTagging", //添加/修改标签
						"name/cos:GetObjectTagging", //获取标签
						//"name/cos:DeleteObjectTagging",
					},
					Effect: "allow",
					Resource: []string{
						// 这里改成允许的路径前缀，可以根据自己网站的用户登录态判断允许上传的具体路径，例子： a.jpg 或者 a/* 或者 * (使用通配符*存在重大安全风险, 请谨慎评估使用)
						// 存储桶的命名格式为 BucketName-APPID，此处填写的 bucket 必须为此格式
						resourcePre + "/upload/" + uid + "/*",
						resourcePre + "/upload/tmp/" + uid + "/*",
					},
				},
			},
		},
	}
	return GetSts(opt)
}
