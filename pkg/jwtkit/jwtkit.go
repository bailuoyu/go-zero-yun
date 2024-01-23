package jwtkit

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt/v4"
	conf "go-zero-yun/public/config"
)

const (
	AuthorizationKey = "Authorization"
	PayloadName      = "jwt_payload"
	TokenKey         = "token"
	TokenValue       = "refresh"
)

type Payload interface {
	GetToken() (string, error)
	ValueFromCtx(ctx context.Context) error
}

func GetToken(payload Payload) (string, error) {
	jwtCfg := conf.Cfg.Pkg.Jwt
	claims := make(jwt.MapClaims)
	iat := time.Now().Unix()
	claims["exp"] = iat + jwtCfg.AccessExpire
	claims["iat"] = iat
	claims[PayloadName] = payload
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims

	return token.SignedString([]byte(jwtCfg.AccessSecret))
}
