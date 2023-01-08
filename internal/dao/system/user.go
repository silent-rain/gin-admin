/*
 * @Author: silent-rain
 * @Date: 2023-01-08 13:19:16
 * @LastEditors: silent-rain
 * @LastEditTime: 2023-01-08 14:36:29
 * @company:
 * @Mailbox: silent_rains@163.com
 * @FilePath: /gin-admin/internal/dao/system/user.go
 * @Descripttion: 用户 Dao
 */
package systemDao

// 用户对象
var UserImpl = new(user)

// 用户接口
type User interface {
}

// 用户结构
type user struct{}

// GetList 获取用户列表
func (d *user) GetList() {

}

// GetList 添加用户
func (d *user) Add() {

}
