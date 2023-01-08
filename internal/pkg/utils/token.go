/*
 * @Author: silent-rain
 * @Date: 2023-01-08 17:34:33
 * @LastEditors: silent-rain
 * @LastEditTime: 2023-01-08 21:33:51
 * @company:
 * @Mailbox: silent_rains@163.com
 * @FilePath: /gin-admin/internal/pkg/utils/token.go
 * @Descripttion: 鉴权令牌
 */
package utils

import (
	"time"

	"gin-admin/internal/pkg/conf"
	statuscode "gin-admin/internal/pkg/status_code"

	"github.com/dgrijalva/jwt-go"
)

// 加密声明
type claims struct {
	UserId   uint
	Phone    string
	Email    string
	Password string
	jwt.StandardClaims
}

// GenerateToken 生成 Token
func GenerateToken(userId uint, phone, email, password string) (string, error) {
	cla := claims{
		UserId:   userId,
		Phone:    phone,
		Email:    email,
		Password: password,
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
func ParseToken(tokenString string) (*claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &claims{}, func(token *jwt.Token) (interface{}, error) {
		return conf.Secret, nil
	})
	if err != nil {
		return nil, statuscode.TokenParsingError.Error()
	}
	claims, ok := token.Claims.(*claims)
	if !ok && !token.Valid {
		return nil, statuscode.TokenInvalidError.Error()
	}
	return claims, nil
}
