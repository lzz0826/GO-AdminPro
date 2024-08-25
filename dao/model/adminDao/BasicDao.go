package adminDao

import (
	"AdminPro/common/driver"
	"AdminPro/common/utils"
	"gorm.io/gorm"
)

type Model interface {
	GetTableName() string
}

func SelectByExample(customizeSQL func(db *gorm.DB) *gorm.DB, out interface{}, table Model) error {
	db := driver.GormDb
	query := db.Debug().Table(table.GetTableName()).Scopes(customizeSQL)
	err := query.Find(out).Error
	if err != nil {
		return err
	}
	return nil
}

func SelectByPrimaryKey(id int, out interface{}, table Model) error {
	db := driver.GormDb
	err := db.Debug().Table(table.GetTableName()).Where("id = ?", id).First(out).Error
	if err != nil {
		return err
	}
	return nil
}

// 返回受影响(删除的比数)
func DeleteByPrimaryKey(id int, table Model) (int64, error) {
	db := driver.GormDb
	result := db.Debug().Table(table.GetTableName()).Delete(table, "id = ?", id)
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}

// 返回受影响(删除的比数)
func DeleteByExample(customizeSQL func(db *gorm.DB) *gorm.DB, table Model) (int64, error) {
	db := driver.GormDb
	result := db.Debug().Table(table.GetTableName()).Scopes(customizeSQL).Delete(table)

	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, nil
}

func CountByExample(customizeSQL func(db *gorm.DB) *gorm.DB, table Model) (int64, error) {
	db := driver.GormDb
	var count int64
	result := db.Debug().Table(table.GetTableName()).Scopes(customizeSQL).Count(&count)
	if result.Error != nil {
		return 0, result.Error
	}
	return count, nil
}

// Insert 插入 DB auto_increment 无须带id ,包含空值(没带的属性 "会" 添加至条件中在 DB NULL)
func Insert(insetCondition interface{}, table Model) (int64, error) {

	var lastInsertID int64
	db := driver.GormDb
	tx := db.Begin()
	if tx.Error != nil {
		return 0, tx.Error
	}

	// 使用 newValue 来创建记录
	if err := tx.Debug().Table(table.GetTableName()).Create(insetCondition).Error; err != nil {
		tx.Rollback()
		return 0, err
	}

	if err := tx.Raw("SELECT LAST_INSERT_ID()").Row().Scan(&lastInsertID); err != nil {
		tx.Rollback()
		return 0, err
	}
	if err := tx.Commit().Error; err != nil {
		return 0, err
	}

	return lastInsertID, nil
}

// InsertSelective auto_increment 无须带id  插入 , 忽略空字段 (没带的属性 "不会" 添加到条件中)
func InsertSelective(insetCondition interface{}, table Model) (int64, error) {
	db := driver.GormDb
	tx := db.Begin()
	if tx.Error != nil {
		return 0, tx.Error
	}

	if err := tx.Debug().Table(table.GetTableName()).Create(utils.BuildNotNullMap(insetCondition)).Error; err != nil {
		tx.Rollback()
		return 0, err
	}

	var lastInsertID int64
	if err := tx.Raw("SELECT LAST_INSERT_ID()").Row().Scan(&lastInsertID); err != nil {
		tx.Rollback()
		return 0, err
	}
	if err := tx.Commit().Error; err != nil {
		return 0, err
	}

	return lastInsertID, nil
}

// InsertSelectiveList auto_increment 无须带id  插入 , 忽略空字段 (没带的属性 "不会" 添加到条件中) 返回所有主建
func InsertSelectiveList[T any](insetCondition []T, table Model) ([]int64, error) {
	var idList []int64
	db := driver.GormDb
	tx := db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	for _, v := range insetCondition {
		if err := tx.Debug().Table(table.GetTableName()).Create(utils.BuildNotNullMap(v)).Error; err != nil {
			tx.Rollback()
			return nil, err
		}

		var lastInsertID int64
		if err := tx.Raw("SELECT LAST_INSERT_ID()").Row().Scan(&lastInsertID); err != nil {
			tx.Rollback()
			return nil, err
		}
		idList = append(idList, lastInsertID)
	}
	if err := tx.Commit().Error; err != nil {
		return nil, err
	}
	return idList, nil
}

// UpdateByExampleSelective 更新 (没带的属性 "不会" 添加到条件中)
func UpdateByExampleSelective(updatesReq interface{}, whereReq interface{}, customizeSQL func(db *gorm.DB) *gorm.DB, table Model) (int64, error) {
	db := driver.GormDb
	result := db.Debug().Table(table.GetTableName()).Where(utils.BuildNotNullMap(whereReq)).Scopes(customizeSQL).Updates(updatesReq)
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}

// UpdateByExample 更新 (没带的属性 "会" 添加至条件中在 DB NULL)
func UpdateByExample(updatesReq interface{}, whereReq interface{}, customizeSQL func(db *gorm.DB) *gorm.DB, table Model) (int64, error) {
	db := driver.GormDb
	upReq := utils.BuildNullMap(updatesReq)
	//更新条件去掉id
	delete(upReq, "id")
	result := db.Debug().Table(table.GetTableName()).Where(utils.BuildNotNullMap(whereReq)).Scopes(customizeSQL).Updates(upReq)
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}

// UpdateByPrimaryKeySelective 更新 (没带的属性 "不会" 添加到条件中)
func UpdateByPrimaryKeySelective(id int, updatesReq interface{}, table Model) (int64, error) {
	db := driver.GormDb
	result := db.Debug().Table(table.GetTableName()).Where("id = ?", id).Updates(updatesReq)
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}

// UpdateByPrimaryKey 更新  (没带的属性 "会" 添加至条件中在 DB NULL)
func UpdateByPrimaryKey(id int, updatesReq interface{}, table Model) (int64, error) {
	db := driver.GormDb.Debug()
	upReq := utils.BuildNullMap(updatesReq)
	//更新条件去掉id
	delete(upReq, "id")
	result := db.Table(table.GetTableName()).Where("id = ?", id).Updates(upReq)
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}
