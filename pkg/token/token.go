package token

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type (
	TokenOptions struct {
		AccessSecret string
		AccessExpire int64
		Fields       map[string]interface{}
	}

	Token struct {
		AccessToken  string `json:"access_token"`
		AccessExpire int64  `json:"access_expire"`
	}
)

func BuildToken(opt TokenOptions) (Token, error) {
	var token Token
	now := time.Now().Add(-time.Minute).Unix()
	// accessToken, err :=

	return token, nil
}

func GenToken(iat int64, secretKey string, payloads map[string]interface{}, seconds int64) (string, error) {
	claims := make(jwt.MapClaims)
}
