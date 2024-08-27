package adminDao

import (
	"AdminPro/common/driver"
	"AdminPro/common/enum"
	"AdminPro/common/model"
	"fmt"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type AccountPayeeCheckBasicDao struct {
}

//Raw: 用于执行原生 SQL 查询并返回结果。
//Exec: 用于执行非查询操作（如插入、更新、删除）。
//First、Last、Find: 用于获取单条或多条记录。
//Pluck: 用于提取单个字段的值。
//Scan: 用于将查询结果映射到自定义结构体。
//Count: 用于统计记录数。
//Scopes: 用于复用查询条件。
//ScanRows: 用于手动处理查询结果集的每一行。

// 使用原始SQL查询
func SelectTypeLast(typeValue int) (*AccountPayeeCheck, error) {
	var result AccountPayeeCheck
	db := driver.GormDb
	sql := `
        SELECT * 
        FROM account_payee_check
        WHERE type > ?
        ORDER BY id desc
        LIMIT 1
    `
	err := db.Debug().Raw(sql, typeValue).Scan(&result).Error
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// 使用原始SQL 修改
func SetMAXType(id int) error {
	db := driver.GormDb
	sql := `
		UPDATE account_payee_check, 
		(SELECT MAX(type) AS m FROM account_payee_check) b 
		SET type = b.m + 1 
		WHERE id = ?
	`
	err := db.Debug().Exec(sql, id).Error
	return err
}

// ListAccountPayeeChecks
func (dao *AccountPayeeCheckBasicDao) ListAccountPayeeChecks(userRandomId *string, status *enum.EAccountPayeeCheckStatusEnum, page *model.Pagination) ([]AccountPayeeCheck, error) {
	var results []AccountPayeeCheck
	customizeSQL := func(db *gorm.DB) *gorm.DB {
		if userRandomId != nil {
			db = db.Where("uid = ?", userRandomId)
		}
		if status != nil {
			db = db.Where("status = ?", status)
		}
		db = db.Order("case when status = 0 then 0 else 1 end asc, created_time desc, id desc")
		return db
	}
	err := SelectByExample(customizeSQL, &results, &AccountPayeeCheck{})

	if err != nil {
		return results, err
	}

	return results, nil
}

func (dao *AccountPayeeCheckBasicDao) ListAccountPayeeChecksPage(userRandomId *string, status *enum.EAccountPayeeCheckStatusEnum, page *model.Pagination) (int64, []AccountPayeeCheck, error) {
	var results []AccountPayeeCheck
	customizeSQL := func(db *gorm.DB) *gorm.DB {
		if userRandomId != nil {
			//db = db.Select("description")

			db = db.Where("uid = ?", userRandomId)
		}
		if status != nil {
			db = db.Where("status = ?", status)
		}
		db = db.Order("case when status = 0 then 0 else 1 end asc, created_time desc, id desc")
		return db
	}
	total, err := SelectByExamplePage(customizeSQL, &results, page, &AccountPayeeCheck{})

	if err != nil {
		return 0, results, err
	}
	return total, results, nil
}

func (dao *AccountPayeeCheckBasicDao) SumTotalStatusSUM(customizeSQL func(db *gorm.DB) *gorm.DB) (*decimal.Decimal, error) {
	var totalAmount decimal.Decimal
	db := driver.GormDb.Debug()
	query := db.Table(AccountPayeeCheck{}.GetTableName()).
		Select("IFNULL(SUM(status), 0) AS total_amount").
		Scopes(customizeSQL)

	var totalAmountStr string
	if err := query.Scan(&totalAmountStr).Error; err != nil {
		return nil, err
	}

	totalAmount, err := decimal.NewFromString(totalAmountStr)
	if err != nil {
		return nil, fmt.Errorf("failed to convert total amount to decimal: %w", err)
	}

	return &totalAmount, nil
}
