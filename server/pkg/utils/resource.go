/*
静态内嵌资源处理
主要用于路由访问路径重置
*/
package utils

import (
	"embed"
	"errors"
	"io/fs"
	"path"
	"path/filepath"
	"strings"

	"gin-admin/assets"
)

// Resource 静态内嵌资源
type Resource struct {
	fs   embed.FS
	path string
}

// NewResource 获取静态内嵌资源对象
func NewResource() *Resource {
	return &Resource{
		fs:   assets.WebStaticAssets,
		path: "dist",
	}
}

// Open 打开静态资源
func (r *Resource) Open(name string) (fs.File, error) {
	if filepath.Separator != '/' && strings.ContainsRune(name, filepath.Separator) {
		return nil, errors.New("http: invalid character in file path")
	}
	fullName := filepath.Join(r.path, filepath.FromSlash(path.Clean("/static/"+name)))
	file, err := r.fs.Open(fullName)
	return file, err
}

// DocsResource Docs 静态内嵌资源
type DocsResource struct {
	fs   embed.FS
	path string
}

// NewDocsResource 获取静态内嵌资源对象
func NewDocsResource() *DocsResource {
	return &DocsResource{
		fs:   assets.WebDocsAssets,
		path: "docs",
	}
}

// Open 打开静态资源
func (r *DocsResource) Open(name string) (fs.File, error) {
	if filepath.Separator != '/' && strings.ContainsRune(name, filepath.Separator) {
		return nil, errors.New("http: invalid character in file path")
	}
	fullName := filepath.Join(r.path, filepath.FromSlash(""+path.Clean(name)))
	file, err := r.fs.Open(fullName)
	return file, err
}
