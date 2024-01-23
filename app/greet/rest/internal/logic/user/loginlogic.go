package user

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-yun/app/greet/rest/internal/svc"
	types "go-zero-yun/app/greet/rest/internal/types"
	conf "go-zero-yun/public/config"
	model "go-zero-yun/public/model/mongo/core"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.UserLoginReq) (resp *types.UserLoginRsp, err error) {
	// todo: add your logic here and delete this line
	payload := types.JwtPayload{
		UserId:   1,
		Username: "cat",
		IsAdmin:  true,
	}
	token, err := payload.GetToken()
	if err != nil {
		return
	}
	now := time.Now()
	MonAccessToken := model.NewAccessTokenModel()
	expireAt := now.Add(time.Duration(conf.Cfg.Pkg.Jwt.AccessExpire) * time.Second)

	act := model.AccessToken{
		ID:       primitive.ObjectID{},
		UserId:   payload.UserId,
		Token:    token,
		ExpireAt: expireAt,
	}
	if err = MonAccessToken.Insert(l.ctx, &act); err != nil {
		return
	}
	//删除其他token
	_, err = MonAccessToken.Conn().DeleteMany(l.ctx,
		bson.D{
			{"user_id", payload.UserId},
			{"_id", bson.M{"$lt": act.ID}},
		},
	)
	if err != nil {
		return
	}
	resp = &types.UserLoginRsp{
		AccessToken:  token,
		AccessExpire: int(expireAt.Unix()),
		RefreshAfter: int(conf.Cfg.Pkg.Jwt.AccessExpire / 2),
	}
	return
}
