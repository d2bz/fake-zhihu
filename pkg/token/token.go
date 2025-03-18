package token

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type (
	TokenOptions struct {
		AccessSecret string
		AccessExpire int64
		UserId       int64
	}

	claims struct {
		jwt.RegisteredClaims
		UserId int64
	}

	Token struct {
		AccessToken  string `json:"access_token"`
		AccessExpire int64  `json:"access_expire"`
	}
)

func BuildToken(opt TokenOptions) (Token, error) {
	var token Token
	accessToken, err := GenToken(opt.AccessSecret, opt.UserId, opt.AccessExpire)
	if err != nil {
		return token, err
	}
	token.AccessToken = accessToken
	token.AccessExpire = time.Now().Unix() + opt.AccessExpire

	return token, nil
}

func GenToken(secretKey string, uid int64, seconds int64) (string, error) {
	claims := &claims{
		UserId: uid,
		RegisteredClaims: jwt.RegisteredClaims{
			// time.Duration的单位是纳秒，需要转化为秒
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(seconds) * time.Second)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	return token.SignedString([]byte(secretKey))
}
