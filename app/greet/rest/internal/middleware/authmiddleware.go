package middleware

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-yun/app/greet/rest/internal/types"
	"go-zero-yun/pkg/jwtkit"
	"go-zero-yun/pkg/logkit"
	conf "go-zero-yun/public/config"
	"go-zero-yun/public/handler"
	model "go-zero-yun/public/model/mongo/core"
	"net/http"
	"time"
)

type AuthMiddleware struct {
}

func NewAuthMiddleware() *AuthMiddleware {
	return &AuthMiddleware{}
}

func (m *AuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO generate middleware implement function, delete after code implementation
		var payload types.JwtPayload
		err := payload.ValueFromCtx(r.Context())
		if err != nil {
			handler.JwtFailRsp(w, r, err)
			return
		}
		// 在ctx及日志链路中注入属性
		ctx := context.WithValue(r.Context(), logkit.UserIdName, payload.UserId)
		ctx = logx.WithFields(ctx, logx.LogField{
			Key:   logkit.UserIdName,
			Value: payload.UserId,
		})
		r = r.WithContext(ctx)
		// 验证token是否过期
		token := r.Header.Get(jwtkit.AuthorizationKey)
		if token == "" {
			handler.JwtFailRsp(w, r, errors.New("empty token"))
			return
		}
		act, err := model.NewAccessTokenModel().CheckToken(r.Context(), payload.UserId, token)
		if err != nil {
			handler.JwtFailRsp(w, r, err)
			return
		}
		//验证是否需要刷新
		if m.needRefresh(act) {
			w.Header().Set(jwtkit.TokenKey, jwtkit.TokenValue)
		}
		// Passthrough to next handler if need
		next(w, r)
	}
}

func (m *AuthMiddleware) needRefresh(act model.AccessToken) bool {
	expire := act.ExpireAt.Unix() - time.Now().Unix()
	if expire < (conf.Cfg.Pkg.Jwt.AccessExpire / 2) {
		return true
	}
	return false
}
