package imsvr

import (
	"context"

	"go-zero-yun/plugin/im"
)

// AccountOnlineStatus 查询用户在线状态 https://cloud.tencent.com/document/product/269/2566
func AccountOnlineStatus(ctx context.Context, userIds []string) (rsp im.QueryOnlineStatusRsp, err error) {
	req := im.QueryOnlineStatusReq{
		IsNeedDetail: 0,
		ToAccount:    userIds,
	}

	client, err := im.GetClient(ctx)
	if err != nil {
		return
	}

	rsp, err = client.AccountOnlineStatus(req)
	return
}

func CreateAccount(ctx context.Context, uid, nick, faceUrl string) (rsp im.Rsp, err error) {
	req := im.CreateAccountReq{
		UserID:  uid,
		Nick:    nick,
		FaceUrl: faceUrl,
	}

	client, err := im.GetClient(ctx)
	if err != nil {
		return
	}

	rsp, err = client.CreateAccount(req)
	return
}
