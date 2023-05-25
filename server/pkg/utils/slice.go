// Package utils 切片操作
package utils

// IndexOfArray 元素在字符串切片中的位置
func IndexOfArray[T comparable](arr []T, target T) int {
	for i, item := range arr {
		if item == target {
			return i
		}
	}
	return -1
}
