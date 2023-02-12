/*MD5 加密
 */
package utils

import (
	"crypto/md5"
	"encoding/hex"

	"gin-admin/internal/pkg/conf"
)

// Md5 加密
func Md5(v string) string {
	m := md5.New()
	m.Write([]byte(conf.Secret))
	m.Write([]byte(v))
	return hex.EncodeToString(m.Sum(nil))
}
