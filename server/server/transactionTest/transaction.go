package transactionTest

import (
	"AdminPro/common/mysql"
	"AdminPro/dao/model/adminDao"
	"context"
	"database/sql"
	"gorm.io/gorm"
	"log"
)

func SomeNeedTransactionFun(ctx context.Context) {
	// 使用 TransactionOptions 设置事务隔离级别Transaction
	err := mysql.GormDb.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 在事务中执行一些数据库操作
		if err := SomeDbFun(tx); err != nil {
			return err // 如果出现错误，则回滚事务
		}

		// 更新字段
		if err := tx.Model(&adminDao.AdminDAO{}).Where("id = ?", 1).Update("field2", "value2").Error; err != nil {
			return err // 如果出现错误，则回滚事务
		}
		// 如果没有错误，则提交事务
		return nil
	}, &sql.TxOptions{
		Isolation: sql.LevelReadCommitted, // 设置隔离级别
		ReadOnly:  false,                  // 设置是否只读
	})

	// 处理事务中可能出现的错误
	if err != nil {
		// 处理错误，例如记录日志
		log.Printf("Transaction failed: %v", err)
	}
}

func SomeDbFun(tx *gorm.DB) error {
	// 插入新记录
	if err := tx.Create(&adminDao.AdminDAO{ID: "value1"}).Error; err != nil {
		return err // 如果出现错误，则回滚事务
	}
	return nil
}
