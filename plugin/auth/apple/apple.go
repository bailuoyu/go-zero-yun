package apple

import (
	"crypto/rsa"
	"encoding/base64"
	"fmt"
	"math/big"
	"math/rand"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/pkg/errors"
)

// RespToken 获取 access token 接口的返回结果
type RespAuth struct {
	OpenId   string `json:"open_id"`
	NickName string `json:"nickname"`
	Face     string `json:"face"`
}
type JwtClaims struct {
	jwt.StandardClaims
}

func (a *Apple) GetUserInfo() (*RespAuth, error) {
	jKeys, err := a.GetJwt(a.code)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	//var pubKeys []rsa.PublicKey
	var pubKey rsa.PublicKey
	for _, val := range jKeys.Keys {
		n_bin, _ := base64.RawURLEncoding.DecodeString(val.N)
		n_data := new(big.Int).SetBytes(n_bin)

		e_bin, _ := base64.RawURLEncoding.DecodeString(val.E)
		e_data := new(big.Int).SetBytes(e_bin)
		pubKey.N = n_data
		pubKey.E = int(e_data.Uint64())
		break
	}
	accessTokenArr := strings.Split(a.code, "##")
	if len(accessTokenArr) == 0 {
		return nil, errors.New("token无效")
	}
	token, err := jwt.ParseWithClaims(accessTokenArr[0], &JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return pubKey, nil
	})
	if claims, ok := token.Claims.(*JwtClaims); ok {
		var nickname string
		if len(accessTokenArr) > 1 {
			nickname = accessTokenArr[1]
		} else {
			num := fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
			nickname = "喵星人" + num
		}
		face := "default.png"
		return &RespAuth{
			OpenId:   claims.Subject,
			NickName: nickname,
			Face:     face,
		}, nil
	} else {
		return nil, errors.New("token claims parse fail")
	}

}
