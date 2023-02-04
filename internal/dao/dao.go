/*
 * @Author: silent-rain
 * @Date: 2023-01-08 13:51:42
 * @LastEditors: silent-rain
 * @LastEditTime: 2023-01-08 21:27:04
 * @company:
 * @Mailbox: silent_rains@163.com
 * @FilePath: /gin-admin/internal/dao/dao.go
 * @Descripttion: Dao 数据库操作
 */
package dao

import (
	"gorm.io/gorm"
)

// Transaction 事务
type Transaction struct {
	tx *gorm.DB
	db *gorm.DB
}

// 创建事务对象
func NewTransaction(db *gorm.DB) *Transaction {
	return &Transaction{
		tx: new(gorm.DB),
		db: db,
	}
}

// Begin 开始事务
func (d *Transaction) Begin() {
	d.tx = d.db.Begin()
}

// 事务对象
func (d *Transaction) Tx() *gorm.DB {
	return d.tx
}

// 遇到错误时回滚事务
func (d *Transaction) Rollback() {
	d.tx.Rollback()
}

// 提交事务
func (d *Transaction) Commit() {
	d.tx.Commit()
}
