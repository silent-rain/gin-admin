// Package utils 文件处理工具
package utils

import "os"

// DirNotExistCreate 文件夹不存在则创建
func DirNotExistCreate(path string) error {
	ok, err := PathExists(path)
	if err != nil {
		return err
	}
	if ok {
		return nil
	}
	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		return err
	}
	return nil
}

// PathExists 判断一个文件或文件夹是否存在
// 输入文件路径，根据返回的bool值来判断文件或文件夹是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
