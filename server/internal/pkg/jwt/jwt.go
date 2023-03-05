/*鉴权令牌
 */
package jwt

import (
	"time"

	"gin-admin/internal/pkg/constant"
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
	cla := Token{
		UserId:   userId,
		Nickname: nickname,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(constant.TokenExpireDuration).Unix(), // 过期时间
			Issuer:    constant.TokenIssuer,                                // 签发人
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, cla)
	// 进行签名生成对应的token
	tokenString, err := token.SignedString([]byte(constant.Secret))
	if err != nil {
		return "", err
	}
	tokenString = constant.TokenPrefix + tokenString
	return tokenString, nil
}

// ParseToken 解析 Token
func ParseToken(tokenString string) (*Token, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Token{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(constant.Secret), nil
	})
	if err != nil {
		return nil, errcode.New(errcode.TokenParsingError)
	}
	claims, ok := token.Claims.(*Token)
	if !ok {
		return nil, errcode.New(errcode.TokenInvalidError)
	} else if !claims.VerifyIssuer(constant.TokenIssuer, true) {
		return nil, errcode.New(errcode.TokenInvalidError)
	} else if !claims.VerifyExpiresAt(time.Now().Unix(), true) {
		return nil, errcode.New(errcode.TokenExpiredError)
	} else if !token.Valid {
		return nil, errcode.New(errcode.TokenInvalidError)
	}
	return claims, nil
}
