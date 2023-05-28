// Package jwt 鉴权令牌
package jwt

import (
	"errors"
	"time"

	"github.com/silent-rain/gin-admin/internal/pkg/conf"
	"github.com/silent-rain/gin-admin/pkg/errcode"

	"github.com/golang-jwt/jwt/v5"
)

// Token 令牌
type Token struct {
	UserId               uint   // 用户ID
	Nickname             string // 用户昵称
	jwt.RegisteredClaims        // 内嵌标准的声明
}

// GenerateToken 生成 Token
func GenerateToken(userId uint, nickname string) (string, error) {
	cfgJWT := conf.Instance().JWT
	claims := Token{
		UserId:   userId,
		Nickname: nickname,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(cfgJWT.GetExpire())), // 过期时间
			Issuer:    cfgJWT.Issuer,                                          // 签发人
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 进行签名生成对应的token
	tokenString, err := token.SignedString([]byte(cfgJWT.Secret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// ParseToken 解析 Token
func ParseToken(tokenString string) (*Token, error) {
	cfgJWT := conf.Instance().JWT
	myToken := &Token{}
	token, err := jwt.ParseWithClaims(tokenString, myToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(cfgJWT.Secret), nil
	})

	if errors.Is(err, jwt.ErrTokenInvalidIssuer) {
		return nil, errcode.TokenIssuerError
	}
	if errors.Is(err, jwt.ErrTokenExpired) {
		return nil, errcode.TokenExpiredError
	}
	// 对token对象中的Claim进行类型断言
	// 校验token
	if !token.Valid {
		return nil, errcode.TokenInvalidError
	}
	return myToken, nil
}
