package adminDao

import (
	"AdminPro/common/model"
	"AdminPro/common/mysql"
	"AdminPro/common/utils"
	"fmt"
	"gorm.io/gorm"
)

// 需要分功能的方法 可以用接收者来实现分页
type BasicDao struct {
	Pagination *model.Pagination
	PageBean   model.PageBean
}

// 给GORM映射的表 需实现 GetDbTableName
type Model interface {
	GetDbTableName() string
}

// Page 设置分页信息
func (dao *BasicDao) Page(pagination model.Pagination) *BasicDao {
	dao.Pagination = &pagination
	return dao
}

// 使用构造判断是否使用分页 执行额外操作后调用
func (dao *BasicDao) SelectCustomizeSqlCheckPage(customizeSQL func(db *gorm.DB) *gorm.DB, out interface{}) error {
	if dao.Pagination != nil {
		total, err := SelectCustomizeSqlPage(customizeSQL, out, dao.Pagination)
		dao.PageBean.Set(total, dao.Pagination.Page, dao.Pagination.Size, out)
		if err != nil {
			return err
		}
	} else {
		// 调用原始的 SelectCustomizeSql
		err := SelectCustomizeSql(customizeSQL, out)
		if err != nil {
			return err
		}
	}
	return nil
}

func SelectCustomizeSql(customizeSQL func(db *gorm.DB) *gorm.DB, out interface{}) error {
	db := mysql.GormDb
	query := db.Debug().Scopes(customizeSQL)
	err := query.Find(out).Error
	if err != nil {
		return err
	}
	return nil
}

// customizeSQL
func SelectCustomizeSqlPage(customizeSQL func(db *gorm.DB) *gorm.DB, out interface{}, page *model.Pagination) (int64, error) {
	var total int64
	db := mysql.GormDb

	countQuery := db.Debug().Scopes(customizeSQL)
	// 执行 Count 操作
	err := countQuery.Model(out).Scopes(utils.WithSelect("COUNT(*)")).Row().Scan(&total)
	if err != nil {
		return 0, err
	}

	query := db.Debug().Model(out).Scopes(customizeSQL)
	query = query.Scopes(utils.WithPagination(page))
	err = query.Find(out).Error
	if err != nil {
		return 0, err
	}
	return total, nil
}

func SelectByObjWhereReq(customizeSQL func(db *gorm.DB) *gorm.DB, whereReq, out interface{}) error {
	db := mysql.GormDb
	query := db.Debug().Where(utils.BuildNotNullMap(whereReq)).Scopes(customizeSQL)
	err := query.Find(out).Error
	if err != nil {
		return err
	}
	return nil
}

// customizeSQL
func SelectByObjWhereReqPage(customizeSQL func(db *gorm.DB) *gorm.DB, whereReq, out interface{}, page *model.Pagination) (int64, error) {
	var total int64
	db := mysql.GormDb

	countQuery := db.Debug().Model(out).Where(utils.BuildNotNullMap(whereReq)).Scopes(customizeSQL)
	// 执行 Count 操作
	err := countQuery.Scopes(utils.WithSelect("COUNT(*)")).Row().Scan(&total)
	if err != nil {
		return 0, nil
	}

	query := db.Debug().Model(out).Where(utils.BuildNotNullMap(whereReq)).Scopes(customizeSQL)
	query = query.Where(utils.BuildNotNullMap(whereReq)).Scopes(utils.WithPagination(page))
	err = query.Find(out).Error
	if err != nil {
		return 0, nil
	}
	return total, nil
}

func SelectByPrimaryKey(id int, out interface{}) error {
	db := mysql.GormDb
	err := db.Debug().Where("id = ?", id).First(out).Error
	if err != nil {
		return err
	}
	return nil
}

// 返回受影响(删除的比数)
func DeleteByPrimaryKey(id int, table interface{}) (int64, error) {
	db := mysql.GormDb
	result := db.Debug().Model(table).Delete(table, "id = ?", id)
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}

// 返回受影响(删除的比数)
func DeleteByList(columnName string, list []int, table interface{}) (int64, error) {
	db := mysql.GormDb
	// 使用 fmt.Sprintf 确保正确插入列名
	query := fmt.Sprintf("%s IN ?", columnName)
	result := db.Debug().Model(table).Where(query, list).Delete(table)
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}

// 返回受影响(删除的比数)
func DeleteCustomizeSql(customizeSQL func(db *gorm.DB) *gorm.DB, table interface{}) (int64, error) {
	db := mysql.GormDb
	result := db.Debug().Model(table).Scopes(customizeSQL).Delete(table)

	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, nil
}

func CountCustomizeSql(customizeSQL func(db *gorm.DB) *gorm.DB, table interface{}) (int64, error) {
	db := mysql.GormDb
	var count int64
	result := db.Debug().Model(table).Scopes(customizeSQL).Count(&count)
	if result.Error != nil {
		return 0, result.Error
	}
	return count, nil
}

// InsertAllowingNull 插入 DB auto_increment 无须带id ,包含空值(没带的属性 "会" 添加至Insert条件中 some_column = NULL)
func InsertAllowingNull(insetCondition interface{}) (int64, error) {

	var lastInsertID int64
	db := mysql.GormDb
	tx := db.Begin()
	if tx.Error != nil {
		return 0, tx.Error
	}
	if err := tx.Debug().Create(insetCondition).Error; err != nil {
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

// InsertAllowingNullCustomizeSQL 插入 DB auto_increment 无须带id ,包含空值(没带的属性 "会" 添加至Insert条件中 some_column = NULL)
func InsertAllowingNullCustomizeSQL(customizeSQL func(db *gorm.DB) *gorm.DB, insetCondition interface{}) (int64, error) {
	var lastInsertID int64
	db := mysql.GormDb
	tx := db.Begin()
	if tx.Error != nil {
		return 0, tx.Error
	}

	if err := tx.Debug().Scopes(customizeSQL).Create(insetCondition).Error; err != nil {
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

// InsertIgnoringNull auto_increment 无须带id  插入 , 忽略空字段 (没带的属性 "不会" 添加到Insert条件中)
func InsertIgnoringNull(insetCondition interface{}) (int64, error) {
	db := mysql.GormDb
	tx := db.Begin()
	if tx.Error != nil {
		return 0, tx.Error
	}

	if err := tx.Debug().Model(insetCondition).Create(utils.BuildNotNullMap(insetCondition)).Error; err != nil {
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

// InsertIgnoringNullCustomizeSQL auto_increment 无须带id  插入 , 忽略空字段 (没带的属性 "不会" 添加到Insert条件中)
func InsertIgnoringNullCustomizeSQL(customizeSQL func(db *gorm.DB) *gorm.DB, insetCondition interface{}) (int64, error) {
	db := mysql.GormDb
	tx := db.Begin()
	if tx.Error != nil {
		return 0, tx.Error
	}

	if err := tx.Debug().Model(insetCondition).Scopes(customizeSQL).Create(utils.BuildNotNullMap(insetCondition)).Error; err != nil {
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

// InsertIgnoringNullList 批量插入 auto_increment 无须带id  插入 , 忽略空字段 (没带的属性 "不会" 添加到Insert条件中) 返回所有主建
func InsertIgnoringNullList[T any](insetCondition []T) ([]int64, error) {
	var idList []int64
	db := mysql.GormDb
	tx := db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}
	for _, v := range insetCondition {
		if err := tx.Debug().Model(insetCondition).Create(utils.BuildNotNullMap(v)).Error; err != nil {
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

// UpdateIgnoringNull 更新 (没带的属性 "不会" 添加到Update条件中)
func UpdateIgnoringNull(updatesReq interface{}, whereReq interface{}, customizeSQL func(db *gorm.DB) *gorm.DB) (int64, error) {
	db := mysql.GormDb
	result := db.Debug().Where(utils.BuildNotNullMap(whereReq)).Scopes(customizeSQL).Updates(updatesReq)
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}

// UpdateAllowingNull 更新 (没带的属性 "会" 添加至Update条件中 some_column = NULL)
func UpdateAllowingNull(updatesReq interface{}, whereReq interface{}, customizeSQL func(db *gorm.DB) *gorm.DB) (int64, error) {
	db := mysql.GormDb
	upReq := utils.BuildNullMap(updatesReq)
	//更新条件去掉id
	delete(upReq, "id")
	result := db.Debug().Model(updatesReq).Where(utils.BuildNotNullMap(whereReq)).Scopes(customizeSQL).Updates(upReq)
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}

// UpdateIgnoringNullByPrimaryKey 更新 (没带的属性 "不会" 添加Update到条件中)
func UpdateIgnoringNullByPrimaryKey(id int, updatesReq interface{}) (int64, error) {
	db := mysql.GormDb
	result := db.Debug().Where("id = ?", id).Updates(updatesReq)
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}

// UpdateAllowingNullByPrimaryKey 更新  (没带的属性 "会" 添加至Update条件中 some_column = NULL)
func UpdateAllowingNullByPrimaryKey(id int, updatesReq interface{}) (int64, error) {
	db := mysql.GormDb.Debug()
	upReq := utils.BuildNullMap(updatesReq)
	//更新条件去掉id
	delete(upReq, "id")
	result := db.Model(updatesReq).Where("id = ?", id).Updates(upReq)
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}
