/*MD5 加密
 */
package utils

import (
	"crypto/md5"
	"encoding/hex"
	"time"

	"github.com/silent-rain/gin-admin/internal/pkg/constant"
)

// EncryptMd5 MD5 密匙加密
func EncryptMd5(v string) string {
	m := md5.New()
	m.Write([]byte(constant.Secret))
	m.Write([]byte(v))
	return hex.EncodeToString(m.Sum(nil))
}

// Md5 MD5 加密
func Md5(v string) string {
	m := md5.New()
	m.Write([]byte(v))
	return hex.EncodeToString(m.Sum(nil))
}

// GenerateTUserApiToken 生成用户API接口Token
func GenerateTUserApiToken() string {
	t := time.Now().UTC().Local().String()
	m := md5.New()
	m.Write([]byte(t))
	return hex.EncodeToString(m.Sum(nil))
}
