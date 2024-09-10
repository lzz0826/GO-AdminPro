package mysql

import (
	"gorm.io/gorm"
)

func WithTransaction(txFunc func(*gorm.DB) error) (err error) {
	tx := GormDb.Begin()
	defer func() {
		if p := recover(); p != nil {
			// 如果有 panic 發生，回滾事務
			tx.Rollback()
			panic(p) // 繼續傳播 panic
		} else if err != nil {
			// 如果有錯誤，回滾事務
			tx.Rollback()
		} else {
			// 提交事務
			err = tx.Commit().Error
		}
	}()
	return txFunc(tx)
}

//
//import (
//	"fmt"
//	"gorm.io/mysql/sqlite"
//	"gorm.io/gorm"
//)
//
//// 定義一個用於處理事務的函數
//func processTransaction(tx *gorm.DB) errors {
//	// 在這裡執行事務內的操作
//
//	// 查询操作
//	var user User
//	if err := tx.Where("name = ?", "John").First(&user).Error; err != nil {
//		return err
//	}
//
//	// 插入操作
//	newUser := User{Name: "Jane", Age: 25}
//	if err := tx.Create(&newUser).Error; err != nil {
//		return err
//	}
//
//	// 模擬一些操作，這裡假設操作成功
//	// 模擬成功的情况下返回 nil，模拟失败的情况下返回错误
//	return nil
//}
//
//func TestMain() {
//	// 初始化 Gorm 的 SQLite3 連接
//	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
//	if err != nil {
//		fmt.Println("Failed to connect to database")
//		return
//	}
//	// 關閉連接
//	defer db.Close()
//
//	// 開始一個事務
//	tx := db.Begin()
//
//	// 檢查事務是否成功開始
//	if tx.Error != nil {
//		// 處理錯誤
//		fmt.Println("Failed to begin transaction")
//		return
//	}
//
//	// 設定事務的隔離性和傳播性 隔离级别为 REPEATABLE READ
//	tx = tx.Set("gorm:query_option", "SET TRANSACTION ISOLATION LEVEL REPEATABLE READ")
//
//	// 執行事務內的操作
//	err = processTransaction(tx)
//
//	if err != nil {
//		// 發生錯誤，回滾事務
//		tx.Rollback()
//		fmt.Println("Transaction rolled back due to errors:", err)
//	} else {
//		// 沒有錯誤，提交事務
//		tx.Commit()
//		fmt.Println("Transaction committed successfully")
//	}
//}
