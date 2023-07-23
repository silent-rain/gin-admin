// Package md5 MD5 加密
package md5

import (
	"crypto/md5"
	"encoding/hex"
	"mime/multipart"
	"path"
	"strconv"
	"time"

	"github.com/silent-rain/gin-admin/pkg/constant"
)

// EncryptMd5 MD5 密匙加密
func EncryptMd5(v string) string {
	m := md5.New()
	m.Write([]byte(constant.Secret))
	m.Write([]byte(v))
	return hex.EncodeToString(m.Sum(nil))
}

// SimpleMd5 简单 MD5 加密
func SimpleMd5(v string) string {
	m := md5.New()
	m.Write([]byte(v))
	return hex.EncodeToString(m.Sum(nil))
}

// GenerateUserApiToken 生成用户API接口Token
func GenerateUserApiToken() string {
	t := time.Now().UTC().Local().String()
	m := md5.New()
	m.Write([]byte(t))
	return hex.EncodeToString(m.Sum(nil))
}

// FIleNameHash 文件名称进行 Hash
func FIleNameHash(file *multipart.FileHeader) string {
	ext := path.Ext(file.Filename)
	filename := file.Filename
	size := strconv.Itoa(int(file.Size))
	timestamp := time.Now().Local().String()
	newFilename := SimpleMd5(filename+size+timestamp) + ext
	return newFilename
}
