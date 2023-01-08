/*
 * @Author: silent-rain
 * @Date: 2023-01-08 15:52:19
 * @LastEditors: silent-rain
 * @LastEditTime: 2023-01-08 21:33:33
 * @company:
 * @Mailbox: silent_rains@163.com
 * @FilePath: /gin-admin/internal/pkg/utils/md5.go
 * @Descripttion:
 */
/*
 * @Author: silent-rain
 * @Date: 2023-01-08 15:52:19
 * @LastEditors: silent-rain
 * @LastEditTime: 2023-01-08 15:56:26
 * @company:
 * @Mailbox: silent_rains@163.com
 * @FilePath: /gin-admin/internal/pkg/utils/md5.go
 * @Descripttion: MD5 加密
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
