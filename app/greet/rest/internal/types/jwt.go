package types

import (
	"context"
	"github.com/mitchellh/mapstructure"
	"go-zero-yun/pkg/jwtkit"
)

type JwtPayload struct {
	UserId   int    `json:"user_id" mapstructure:"user_id"`
	Username string `json:"username" mapstructure:"username"`
	IsAdmin  bool   `json:"is_admin"  mapstructure:"is_admin"`
}

// GetToken 获取token
func (jp *JwtPayload) GetToken() (string, error) {
	return jwtkit.GetToken(jp)
}

// ValueFromCtx 从ctx中解析指
func (jp *JwtPayload) ValueFromCtx(ctx context.Context) error {
	payloadAny := ctx.Value(jwtkit.PayloadName)
	return mapstructure.Decode(payloadAny, jp)
}
