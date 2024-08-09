package adminDao

import (
	"AdminPro/common/driver"
	"AdminPro/common/model"
	"AdminPro/common/utils"
	"gorm.io/gorm"
)

type AccountPayeeCheckDao struct {
}

func (apd *AccountPayeeCheckDao) TableName() string {
	return "account_payee_check"
}

func (apd *AccountPayeeCheckDao) SelectByExample(uid *int, status *int, customizeSQL func(db *gorm.DB) *gorm.DB, page *model.Pagination) ([]AccountPayeeCheck, error) {
	var results []AccountPayeeCheck
	db := driver.GormDb

	query := db.Debug().Model(AccountPayeeCheck{}).Table(apd.TableName())

	if uid != nil {
		query = query.Where("uid = ?", *uid)
	}
	if status != nil {
		query = query.Where("status = ?", *status)
	}

	query.Scopes(utils.WithPagination(page.Page, page.Limit))

	query.Scopes(customizeSQL)

	err := query.Find(&results).Error
	if err != nil {
		return nil, err
	}

	return results, nil
}

func (apd *AccountPayeeCheckDao) SelectByExampleSelectGeneric(customizeSQL func(db *gorm.DB) *gorm.DB) ([]AccountPayeeCheck, error) {
	var results []AccountPayeeCheck

	err := utils.SelectGeneric(apd.TableName(), customizeSQL, &results)

	if err != nil {
		return nil, err
	}

	return results, nil
}

func (apd *AccountPayeeCheckDao) SelectByExampleEX(customizeSQL func(db *gorm.DB) *gorm.DB) ([]AccountPayeeCheck, error) {
	var results []AccountPayeeCheck

	db := driver.GormDb

	query := db.Debug().Model(AccountPayeeCheck{}).Table(apd.TableName())

	// 使用传入的自定义查询函数
	query = query.Scopes(customizeSQL)

	// 执行查询并填充结果集
	err := query.Find(&results).Error
	if err != nil {
		return nil, err
	}

	return results, nil
}

func (apd *AccountPayeeCheckDao) SelectByPrimaryKey(id int) (AccountPayeeCheck, error) {
	var result AccountPayeeCheck
	db := driver.GormDb

	err := db.Debug().Model(AccountPayeeCheck{}).Table(apd.TableName()).Where("id = ?", id).First(&result).Error
	if err != nil {
		return result, err
	}

	return result, nil
}

func (apd *AccountPayeeCheckDao) DeleteByExample(id int) error {
	db := driver.GormDb

	err := db.Debug().Table(apd.TableName()).Delete(&AccountPayeeCheck{}, "id = ?", id).Error
	if err != nil {
		return err
	}
	return nil
}

func (apd *AccountPayeeCheckDao) DeleteByCustomizeSQL(customizeSQL func(db *gorm.DB) *gorm.DB) error {
	db := driver.GormDb

	err := db.Debug().Table(apd.TableName()).Scopes(customizeSQL).Delete(&AccountPayeeCheck{}).Error
	if err != nil {
		return err
	}
	return nil
}

// Insert 插入,包含空值(没带的属性 "会" 添加至条件中在 DB NULL)
func (apd *AccountPayeeCheckDao) Insert(a AccountPayeeCheck) (int64, error) {

	// 创建记录时忽略 ID 字段
	a.ID = nil

	db := driver.GormDb

	// 开始一个事务
	tx := db.Begin()
	if tx.Error != nil {
		return 0, tx.Error
	}

	if err := tx.Debug().Table(apd.TableName()).Create(&a).Error; err != nil {
		tx.Rollback() // 发生错误回滚
		return 0, err
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		return 0, err
	}

	// 返回最后插入的自增 ID
	return int64(*a.ID), nil
}

// InsertSelective 插入 , 忽略空字段 (没带的属性 "不会" 添加到条件中)
func (apd *AccountPayeeCheckDao) InsertSelective(a AccountPayeeCheck) (int64, error) {
	// 创建记录时忽略 ID 字段
	a.ID = nil

	db := driver.GormDb

	// 开始一个事务
	tx := db.Begin()
	if tx.Error != nil {
		return 0, tx.Error
	}

	// 将非空字段插入数据库 因使用MAP GORM 不会映射 ID
	if err := tx.Debug().Table(apd.TableName()).Create(utils.BuildNotNullMap(&a)).Error; err != nil {
		tx.Rollback() // 发生错误回滚
		return 0, err
	}

	// 获取自增 ID
	var lastInsertID int64
	if err := tx.Raw("SELECT LAST_INSERT_ID()").Row().Scan(&lastInsertID); err != nil {
		tx.Rollback() // 发生错误回滚
		return 0, err
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		return 0, err
	}

	return lastInsertID, nil
}

// CountByExample Count(没带的属性 "不会" 添加到条件中)
func (apd *AccountPayeeCheckDao) CountByExample(a AccountPayeeCheck) (int64, error) {
	db := driver.GormDb
	var count int64
	result := db.Debug().Table(apd.TableName()).Where(utils.BuildNotNullMap(a)).Count(&count)
	if result.Error != nil {
		return 0, result.Error
	}
	return count, nil
}

func (apd *AccountPayeeCheckDao) CountByCustomizeSQL(customizeSQL func(db *gorm.DB) *gorm.DB) (int64, error) {
	db := driver.GormDb
	var count int64
	result := db.Debug().Table(apd.TableName()).Scopes(customizeSQL).Count(&count)
	if result.Error != nil {
		return 0, result.Error
	}
	return count, nil
}

// UpdateByExampleSelective 更新 (没带的属性 "不会" 添加到条件中)
func (apd *AccountPayeeCheckDao) UpdateByExampleSelective(updatesReq AccountPayeeCheck, whereReq AccountPayeeCheck) (int64, error) {
	db := driver.GormDb
	result := db.Debug().Table(apd.TableName()).Where(utils.BuildNotNullMap(whereReq)).Updates(updatesReq)

	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, nil
}

// UpdateByCustomizeSQL 更新 (没带的属性 "不会" 添加到条件中)
func (apd *AccountPayeeCheckDao) UpdateByCustomizeSQL(updatesReq AccountPayeeCheck, whereReq AccountPayeeCheck, customizeSQL func(db *gorm.DB) *gorm.DB) (int64, error) {
	db := driver.GormDb
	result := db.Debug().Table(apd.TableName()).Where(utils.BuildNotNullMap(whereReq)).Scopes(customizeSQL).Updates(updatesReq)
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}

// UpdateByExample 更新 (没带的属性 "会" 添加至条件中在 DB NULL)
func (apd *AccountPayeeCheckDao) UpdateByExample(updatesReq AccountPayeeCheck, whereReq AccountPayeeCheck) (int64, error) {
	db := driver.GormDb
	upReq := utils.BuildNullMap(updatesReq)
	//更新条件去掉id
	delete(upReq, "id")
	result := db.Debug().Table(apd.TableName()).Where(utils.BuildNotNullMap(whereReq)).Updates(upReq)
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}

func (apd *AccountPayeeCheckDao) UpdateByExampleCustomizeSQL(updatesReq AccountPayeeCheck, whereReq AccountPayeeCheck, customizeSQL func(db *gorm.DB) *gorm.DB) (int64, error) {
	db := driver.GormDb
	upReq := utils.BuildNullMap(updatesReq)
	//更新条件去掉id
	delete(upReq, "id")
	result := db.Debug().Table(apd.TableName()).Where(utils.BuildNotNullMap(whereReq)).Scopes(customizeSQL).Updates(upReq)
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}

// UpdateByPrimaryKeySelective 更新 (没带的属性 "不会" 添加到条件中)
func (apd *AccountPayeeCheckDao) UpdateByPrimaryKeySelective(id int, updatesReq AccountPayeeCheck) error {
	db := driver.GormDb
	err := db.Debug().Table(apd.TableName()).Where("id = ?", id).Updates(updatesReq).Error
	if err != nil {
		return err
	}
	return nil
}

// UpdateByPrimaryKey 更新  (没带的属性 "会" 添加至条件中在 DB NULL)
func (apd *AccountPayeeCheckDao) UpdateByPrimaryKey(id int, updatesReq AccountPayeeCheck) (int64, error) {
	db := driver.GormDb
	upReq := utils.BuildNullMap(updatesReq)
	//更新条件去掉id
	delete(upReq, "id")
	result := db.Table(apd.TableName()).Where("id = ?", id).Updates(upReq)
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}
