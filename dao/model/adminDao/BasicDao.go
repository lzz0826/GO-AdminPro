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
func (dao *BasicDao) SelectCheckPage(db *gorm.DB, out interface{}) error {
	if dao.Pagination != nil {
		total, err := SelectPage(db, out, dao.Pagination)
		dao.PageBean.Set(total, dao.Pagination.Page, dao.Pagination.Size, out)
		if err != nil {
			return err
		}
	} else {
		// 调用原始的 Select
		err := Select(db, out)
		if err != nil {
			return err
		}
	}
	return nil
}

func Select(db *gorm.DB, out interface{}) error {
	CheckGormDb(db)
	query := db.Debug()
	err := query.Find(out).Error
	if err != nil {
		return err
	}
	return nil
}

// customizeSQL
func SelectPage(db *gorm.DB, out interface{}, page *model.Pagination) (int64, error) {
	var total int64
	CheckGormDb(db)
	countQuery := db.Debug()
	// 执行 Count 操作
	err := countQuery.Model(out).Scopes(utils.WithSelect("COUNT(*)")).Row().Scan(&total)
	if err != nil {
		return 0, err
	}

	query := db.Debug().Model(out)
	query = query.Scopes(utils.WithPagination(page))
	err = query.Find(out).Error
	if err != nil {
		return 0, err
	}
	return total, nil
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

// 返回受影响
func Delete(db *gorm.DB, table interface{}) (int64, error) {
	CheckGormDb(db)
	result := db.Debug().Model(table).Delete(table)
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}

func Count(db *gorm.DB, table interface{}) (int64, error) {
	CheckGormDb(db)
	var count int64
	result := db.Debug().Model(table).Count(&count)
	if result.Error != nil {
		return 0, result.Error
	}
	return count, nil
}

func InsertReturnLastId(db *gorm.DB, insetCondition interface{}) (int64, error) {
	CheckGormDb(db)
	var lastInsertID int64
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

// InsertReturnLastIds 批量插入
func InsertsReturnLastIds[T any](db *gorm.DB, insetCondition []T) ([]int64, error) {
	CheckGormDb(db)
	var idList []int64
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

// Updates
func Updates(db *gorm.DB, updatesReq interface{}, whereReq map[string]interface{}) (int64, error) {
	CheckGormDb(db)
	result := db.Debug().Where(whereReq).Updates(updatesReq)
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}

// UpdateByPrimaryKey
func UpdateByPrimaryKey(db *gorm.DB, id int, updatesReq interface{}) (int64, error) {
	CheckGormDb(db)
	result := db.Debug().Where("id = ?", id).Updates(updatesReq)
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}

func CheckGormDb(db *gorm.DB) {
	if db == nil {
		db = mysql.GormDb
	}
}
