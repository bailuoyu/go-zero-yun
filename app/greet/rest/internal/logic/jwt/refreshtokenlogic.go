package jwt

import (
	"context"
	"go-zero-yun/app/greet/rest/internal/svc"
	"go-zero-yun/app/greet/rest/internal/types"
	conf "go-zero-yun/public/config"
	model "go-zero-yun/public/model/mongo/core"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type RefreshTokenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRefreshTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RefreshTokenLogic {
	return &RefreshTokenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RefreshTokenLogic) RefreshToken(req *types.Empty) (resp *types.RefreshTokenRsp, err error) {
	// todo: add your logic here and delete this line
	var payload types.JwtPayload
	err = payload.ValueFromCtx(l.ctx)
	if err != nil {
		return
	}
	token, err := payload.GetToken()
	if err != nil {
		return
	}
	now := time.Now().Unix()
	MonAccessToken := model.NewAccessTokenModel()
	act := model.AccessToken{
		ID:     primitive.ObjectID{},
		UserId: payload.UserId,
		Token:  token,
	}
	if err = MonAccessToken.Insert(l.ctx, &act); err != nil {
		return
	}
	//限制最多5个
	MonAccessToken.TokenLimit(l.ctx, payload.UserId, 5)
	resp = &types.RefreshTokenRsp{
		AccessToken:  token,
		AccessExpire: now + conf.Cfg.Pkg.Jwt.AccessExpire,
		RefreshAfter: now + conf.Cfg.Pkg.Jwt.AccessExpire/2,
	}
	return
}
