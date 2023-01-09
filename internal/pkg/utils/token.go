/*
 * @Author: silent-rain
 * @Date: 2023-01-08 17:34:33
 * @LastEditors: silent-rain
 * @LastEditTime: 2023-01-09 23:26:56
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
	"github.com/gin-gonic/gin"
)

// Token 令牌
type Token struct {
	userId   uint
	phone    string
	email    string
	password string
	jwt.StandardClaims
}

// GenerateToken 生成 Token
func GenerateToken(userId uint, phone, email, password string) (string, error) {
	cla := Token{
		userId:   userId,
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
		return nil, statuscode.TokenInvalidError.Error()
	} else if !claims.VerifyIssuer(conf.TokenIssuer, false) {
		return nil, statuscode.TokenInvalidError.Error()
	} else if !claims.VerifyExpiresAt(time.Now().Unix(), false) {
		return nil, statuscode.TokenExpiredError.Error()
	} else if !token.Valid {
		return nil, statuscode.TokenInvalidError.Error()
	}
	return claims, nil
}

// GetUserId 获取用户 ID
func GetUserId(ctx *gin.Context) uint {
	v, ok := ctx.Get(GinContextToken)
	if !ok {
		return 0
	}
	token := v.(Token)
	return token.userId
}
