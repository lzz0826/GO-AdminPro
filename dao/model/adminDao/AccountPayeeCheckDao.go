package adminDao

import (
	"AdminPro/common/driver"
	"AdminPro/common/model"
	"AdminPro/common/utils"
	"database/sql"
	"fmt"
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

func (apd *AccountPayeeCheckDao) SelectByExampleCustomizeSQL(customizeSQL func(db *gorm.DB) *gorm.DB) ([]AccountPayeeCheck, error) {
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
	result := db.Debug().Model(AccountPayeeCheck{}).Table(apd.TableName()).Scopes(customizeSQL).Count(&count)
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

//(没带的属性 "会" 添加至条件中在 DB NULL)
func (apd *AccountPayeeCheckDao) UpdateDBNullTest(uid, status, checkID *int, description *string) (int64, error) {
	updates := map[string]interface{}{}

	//如果没有添加NullString会被忽略 不会添加到条件中
	if description != nil {
		updates["description"] = sql.NullString{String: *description, Valid: true}
	} else {
		updates["description"] = sql.NullString{Valid: false}
	}

	//如果没有添加NullInt32会被忽略 不会添加到条件中
	if status != nil {
		updates["status"] = sql.NullInt32{Int32: int32(*status), Valid: true}
	} else {
		updates["status"] = sql.NullInt32{Valid: false}
	}

	if len(updates) == 0 {
		return 0, fmt.Errorf("no fields to update")
	}

	db := driver.GormDb
	//uid check_id 查询条件会自动代收寻条件  WHERE uid = NULL AND check_id = 3
	result := db.Debug().Table(apd.TableName()).Where("uid = ? AND check_id = ?", uid, checkID).Updates(updates)
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}

func (apd *AccountPayeeCheckDao) FindRecordByStatusAndUey(status int, uid string, page int, pageSize int) (model.PageBean, error) {

	bean := model.PageBean{}

	buildQuery := func(db *gorm.DB) *gorm.DB {
		if status >= 0 {
			db = db.Where("status = ?", status)
		}
		if uid != "" {
			db = db.Where("uid = ?", uid)
		}
		return db
	}

	totalRecords, err := apd.CountByCustomizeSQL(buildQuery)

	if err != nil {
		return bean, err
	}

	customizeSQL := func(db *gorm.DB) *gorm.DB {
		db = db.Scopes(buildQuery)
		db = db.Scopes(utils.WithPagination(page, pageSize))
		return db
	}
	example, err := apd.SelectByExampleCustomizeSQL(customizeSQL)

	if err != nil {
		return bean, err
	}

	pageBean := model.Of(totalRecords, page, pageSize, example)

	return *pageBean, nil
}

//accountStatus 指針會有 WHERE aa.account_status = 1 或 WHERE aa.account_status = NULL的情況
func (apd *AccountPayeeCheckDao) SelectByAccountStatus(accountStatus *int) ([]ClubOnUserStatistics, error) {
	db := driver.GormDb
	query := `
		SELECT
			aa.id AS clubId,
		    apc.status AS normalNum,
		    apc.type AS opNum
		FROM
			account_payee_check apc
		INNER JOIN
			admin_admin aa ON apc.uid = aa.id
		WHERE
			aa.account_status = ?
	`

	var cs []ClubOnUserStatistics
	if err := db.Debug().Raw(query, accountStatus).Scan(&cs).Error; err != nil {
		return nil, err
	}
	return cs, nil
}

//返回最後自增ID DB會有 NULL 情況
func (apd *AccountPayeeCheckDao) AddAccountPayeeCheck(a AccountPayeeCheck) (int64, error) {
	db := driver.GormDb

	// 开始一个事务
	tx := db.Begin()
	if tx.Error != nil {
		return 0, tx.Error
	}

	// 执行插入操作
	if err := tx.Debug().Exec(`
        INSERT INTO 
            account_payee_check (check_id, check_time, description, status, type, uid, created_time, update_time)
        VALUES 
            (?, ?, ?, ?, ?, ?, NOW(), NOW())
    `, a.CheckID, a.CheckTime, a.Description, a.Status, a.Type, a.UID).Error; err != nil {
		tx.Rollback() // 发生错误回滚
		return 0, err
	}

	// 获取自增 ID
	var lastInsertID int64
	if err := tx.Debug().Raw("SELECT LAST_INSERT_ID()").Scan(&lastInsertID).Error; err != nil {
		tx.Rollback() // 发生错误回滚
		return 0, err
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		return 0, err
	}

	return lastInsertID, nil
}

//打印出的SQL
//SELECT aa.id AS clubId,
//COUNT(CASE WHEN apc.status = '1' AND apc.type = '2' THEN 1 ELSE NULL END) AS normalNum,
//COUNT(CASE WHEN apc.status = '1' AND apc.type = '2' THEN 1 ELSE NULL END) AS opNum
//FROM account_payee_check apc
//LEFT JOIN admin_admin aa ON apc.uid = aa.id
//WHERE aa.id IN (1,2,3)
//GROUP BY aa.id
func CustomizeSQL(db *gorm.DB, clubIdLst []int64) ([]ClubOnUserStatistics, error) {
	// Create a comma-separated list of club IDs for the SQL IN clause
	clubIdPlaceholders := ""
	for i := range clubIdLst {
		if i > 0 {
			clubIdPlaceholders += ","
		}
		clubIdPlaceholders += "?"
	}

	// Construct the raw SQL query
	query := `
		SELECT aa.id AS clubId,
		       COUNT(CASE WHEN apc.status = '1' AND apc.type = '2' THEN 1 ELSE NULL END) AS normalNum,
		       COUNT(CASE WHEN apc.status = '1' AND apc.type = '2' THEN 1 ELSE NULL END) AS opNum
		FROM account_payee_check apc
		LEFT JOIN admin_admin aa ON apc.uid = aa.id
		WHERE aa.id IN (` + clubIdPlaceholders + `)
		GROUP BY aa.id
	`

	// Execute the raw SQL query and scan the results into a slice of ClubStats
	var stats []ClubOnUserStatistics
	if err := db.Debug().Raw(query, utils.ConvertToInterfaceSlice(clubIdLst)...).Scan(&stats).Error; err != nil {
		return nil, err
	}

	return stats, nil
}

// 包含了QueryComplaintList 和 CountComplaintList
//func (dao *CheatComplaintDao) QueryComplaintListAndCount(typeVal *int, startTime, endTime *time.Time, start, size int, status *int) ([]CheatComplaintDao, int64, error) {
//	var count int64
//	var results []CheatComplaintDao
//	db := sqldb.GetSqlDB().GameTidb
//
//	query := db.Table(dao.TableName()).
//		Select("t1.id, t2.random_num as complainant, t1.type, t1.content, t1.status, DATE_FORMAT(t1.create_time,'%Y/%c/%d %H:%i') as time, t2.nike_name as complainantName, t4.name as clubName").
//		Joins("join user_details_info t2 on t1.complainant = t2.USER_ID").
//		Joins("join club_members t3 on t2.USER_ID = t3.user_id").
//		Joins("join club_record t4 on t3.club_id = t4.id")
//
//	if typeVal != nil {
//		query = query.Where("t1.type = ?", *typeVal)
//	}
//	if startTime != nil {
//		query = query.Where("t1.create_time >= ?", *startTime)
//	}
//	if endTime != nil {
//		query = query.Where("t1.create_time <= ?", *endTime)
//	}
//	//CountComplaintList
//	if err := query.Count(&count).Error; err != nil {
//		return nil, 0, err
//	}
//	if status != nil && *status != -1 {
//		query = query.Where("t1.status = ?", *status)
//	}
//	//QueryComplaintList
//	if err := query.Order("t1.create_time desc").Offset(start).Limit(size).Scan(&results).Error; err != nil {
//		return nil, 0, err
//	}
//	return results, count, nil
//}
