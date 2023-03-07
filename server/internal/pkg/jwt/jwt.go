/*鉴权令牌
 */
package jwt

import (
	"time"

	"gin-admin/internal/pkg/conf"
	"gin-admin/pkg/errcode"

	"github.com/dgrijalva/jwt-go"
)

// Token 令牌
type Token struct {
	UserId   uint
	Nickname string
	jwt.StandardClaims
}

// GenerateToken 生成 Token
func GenerateToken(userId uint, nickname string) (string, error) {
	cfgJWT := conf.Instance().JWT
	cla := Token{
		UserId:   userId,
		Nickname: nickname,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(cfgJWT.GetExpire()).Unix(), // 过期时间
			Issuer:    cfgJWT.Issuer,                             // 签发人
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, cla)
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
	token, err := jwt.ParseWithClaims(tokenString, &Token{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(cfgJWT.Secret), nil
	})
	if err != nil {
		return nil, errcode.New(errcode.TokenParsingError)
	}
	claims, ok := token.Claims.(*Token)
	if !ok {
		return nil, errcode.New(errcode.TokenInvalidError)
	} else if !claims.VerifyIssuer(cfgJWT.Issuer, true) {
		return nil, errcode.New(errcode.TokenInvalidError)
	} else if !claims.VerifyExpiresAt(time.Now().Unix(), true) {
		return nil, errcode.New(errcode.TokenExpiredError)
	} else if !token.Valid {
		return nil, errcode.New(errcode.TokenInvalidError)
	}
	return claims, nil
}
