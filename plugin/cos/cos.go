package cos

import (
	"bytes"
	"context"
	"fmt"
	"go-zero-yun/plugin"
	conf "go-zero-yun/public/config"

	cos5 "github.com/tencentyun/cos-go-sdk-v5"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
)

const (
	publicTmp    = "public/tmp/"    //公有tmp目录
	protectedTmp = "protected/tmp/" //私有tmp目录
	TmpTagKey    = "tmp"
)

// ClientCos cos客户端
type ClientCos struct {
	Config   Config
	Client   *cos5.Client
	CosTmp   string //cos临时目录
	LocalTmp string //本地临时目录
	ctx      context.Context
}

// GetClient 获取Client
func GetClient(ctx context.Context) *ClientCos {
	return GetClientByName(ctx, plugin.DefaultName)
}

// GetClientByName 根据名称获取Client
func GetClientByName(ctx context.Context, name string) *ClientCos {
	cfg := GetCfgByName(name)
	urlStr := fmt.Sprintf(
		//"https://%s.cos-internal.%s.myqcloud.com",	//内网专用,腾讯云会自动解析,使用外网即可
		"https://%s.cos.%s.myqcloud.com",
		cfg.Bucket, cfg.Region)
	u, _ := url.Parse(urlStr)
	b := &cos5.BaseURL{BucketURL: u}
	client := cos5.NewClient(b, &http.Client{
		Transport: &cos5.AuthorizationTransport{
			SecretID:  cfg.SecretId,
			SecretKey: cfg.SecretKey,
		},
	})
	clc := &ClientCos{
		Config: cfg,
		Client: client,
	}
	clc.GetCosTmp()
	clc.GetLocalTmp()
	clc.ctx = ctx
	return clc
}

// GetDomain 获取访问域名
func (clc *ClientCos) GetDomain() string {
	return clc.Config.CdnDomain
}

// GetCosTmp 获取Cos临时目录
func (clc *ClientCos) GetCosTmp() string {
	if clc.CosTmp == "" {
		clc.CosTmp = fmt.Sprintf("%s/%s", conf.Cfg.Server.App, publicTmp)
	}
	return clc.CosTmp
}

// GetLocalTmp 获取本地临时目录
func (clc *ClientCos) GetLocalTmp() string {
	if clc.LocalTmp == "" {
		clc.LocalTmp = os.TempDir()
	}
	return clc.LocalTmp
}

// GetCosTmpRandomDir 获取cos临时随机目录
func (clc *ClientCos) GetCosTmpRandomDir() string {
	dir := fmt.Sprintf("%s/%s", clc.GetCosTmp(), randomStr(32))
	return dir
}

// GetLocalTmpRandomDir 获取本地临时随机目录
func (clc *ClientCos) GetLocalTmpRandomDir() (string, error) {
	dir := fmt.Sprintf("%s/cos/%s", clc.GetLocalTmp(), randomStr(32))
	err := os.MkdirAll(dir, 644)
	return dir, err
}

// UploadLocal 上传本地文件，key为cos路径，localPath为本地路径
func (clc *ClientCos) UploadLocal(key string, localPath string, opt *cos5.MultiUploadOptions) (
	*cos5.CompleteMultipartUploadResult, error) {
	res, _, err := clc.Client.Object.Upload(clc.ctx, key, localPath, opt)
	return res, err
}

// UploadFile 上传文件，key为cos路径
func (clc *ClientCos) UploadFile(key string, r *bytes.Reader) (err error) {
	_, err = clc.Client.Object.Put(clc.ctx, key, r, nil)
	return err
}

// UploadLocalToTmp 上传本地文件至临时目录，localPath为本地路径，文件名和本地路径名一样
func (clc *ClientCos) UploadLocalToTmp(localPath string, opt *cos5.MultiUploadOptions) (
	*cos5.CompleteMultipartUploadResult, error) {
	fileName := path.Base(filepath.ToSlash(localPath))
	key := clc.GetCosTmp() + fileName
	res, err := clc.UploadLocal(key, localPath, opt)
	return res, err
}

// DownloadLocal 下载到本地，key为cos路径，localPath为本地路径
func (clc *ClientCos) DownloadLocal(key string, localPath string,
	opt *cos5.MultiDownloadOptions, id ...string) error {
	_, err := clc.Client.Object.Download(clc.ctx, key, localPath, opt, id...)
	return err
}

// DownloadLocalToTmp 下载到本地临时目录，key为cos路径(支持路径和url)，文件名和本地路径名一样
func (clc *ClientCos) DownloadLocalToTmp(key string, opt *cos5.MultiDownloadOptions) (
	string, error) {
	fileName := path.Base(key)
	key, err := GetKey(key)
	if err != nil {
		return "", err
	}
	localPath := clc.GetLocalTmp() + fileName
	err = clc.DownloadLocal(key, localPath, opt)
	return localPath, err
}

// GetFileByte 获取文件byte
func (clc *ClientCos) GetFileByte(key string, opt *cos5.ObjectGetOptions) ([]byte, error) {
	res, err := clc.Client.Object.Get(clc.ctx, key, opt)
	if err != nil {
		return []byte{}, err
	}
	body, err := io.ReadAll(res.Body)
	return body, err
}

// Delete 删除文件
func (clc *ClientCos) Delete(key string) error {
	key, err := GetKey(key)
	if err != nil {
		return err
	}
	_, err = clc.Client.Object.Delete(clc.ctx, key)
	return err
}

// MultiDelete 批量删除文件
func (clc *ClientCos) MultiDelete(keys []string, opt *cos5.ObjectDeleteMultiOptions) (
	*cos5.ObjectDeleteMultiResult, error) {
	var obs []cos5.Object
	for _, v := range keys {
		key, err := GetKey(v)
		if err != nil {
			return nil, err
		}
		obs = append(obs, cos5.Object{Key: key})
	}
	opt.Objects = obs
	res, _, err := clc.Client.Object.DeleteMulti(clc.ctx, opt)
	return res, err
}
