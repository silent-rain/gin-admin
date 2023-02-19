/*DAO 数据库操作
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
