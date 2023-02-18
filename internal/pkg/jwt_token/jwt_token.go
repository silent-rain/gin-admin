/*鉴权令牌
 */
package jwt_token

import (
	"time"

	"gin-admin/internal/pkg/conf"
	"gin-admin/pkg/errcode"

	"github.com/dgrijalva/jwt-go"
)

// Token 令牌
type Token struct {
	UserId   uint
	phone    string
	email    string
	password string
	jwt.StandardClaims
}

// GenerateToken 生成 Token
func GenerateToken(userId uint, phone, email, password string) (string, error) {
	cla := Token{
		UserId:   userId,
		phone:    phone,
		email:    email,
		password: password,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(conf.TokenExpireDuration).Unix(), // 过期时间
			Issuer:    conf.TokenIssuer,                                // 签发人
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, cla)
	// 进行签名生成对应的token
	tokenString, err := token.SignedString([]byte(conf.Secret))
	if err != nil {
		return "", err
	}
	tokenString = conf.TokenPrefix + tokenString
	return tokenString, nil
}

// ParseToken 解析 Token
func ParseToken(tokenString string) (*Token, error) {
	token, _ := jwt.ParseWithClaims(tokenString, &Token{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(conf.Secret), nil
	})
	claims, ok := token.Claims.(*Token)
	if !ok {
		return nil, errcode.New(errcode.TokenInvalidError)
	} else if !claims.VerifyIssuer(conf.TokenIssuer, true) {
		return nil, errcode.New(errcode.TokenInvalidError)
	} else if !claims.VerifyExpiresAt(time.Now().Unix(), true) {
		return nil, errcode.New(errcode.TokenExpiredError)
	} else if !token.Valid {
		return nil, errcode.New(errcode.TokenInvalidError)
	}
	return claims, nil
}
