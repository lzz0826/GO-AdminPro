package adminDao

import (
	"AdminPro/common/driver"
	"AdminPro/common/model"
	"AdminPro/common/utils"
)

type AccountPayeeCheckDao struct {
}

func (apd *AccountPayeeCheckDao) TableName() string {
	return "account_payee_check"
}

func (apd *AccountPayeeCheckDao) SelectByExample(uid *int, status *int, page *model.Pagination) ([]AccountPayeeCheck, error) {
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

	//排序 未审核的优先排前面 其余按创建时间倒续
	query.Scopes(utils.WithOrderBySQL("case when status = 0 then 0 else 1 end asc, created_time desc, id desc"))

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

// Insert 插入,包含空值(没带的属性 "会" 添加至条件中在 DB NULL)
func (apd *AccountPayeeCheckDao) Insert(a AccountPayeeCheck) error {
	db := driver.GormDb

	err := db.Debug().Table(apd.TableName()).Create(a).Error
	if err != nil {
		return err
	}

	return nil
}

// InsertSelective 插入 , 忽略空字段 (没带的属性 "不会" 添加到条件中)
func (apd *AccountPayeeCheckDao) InsertSelective(a AccountPayeeCheck) error {
	db := driver.GormDb

	err := db.Debug().Table(apd.TableName()).Create(utils.BuildNotNullMap(a)).Error
	if err != nil {
		return err
	}

	return nil
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

// UpdateByExampleSelective 更新 (没带的属性 "不会" 添加到条件中)
func (apd *AccountPayeeCheckDao) UpdateByExampleSelective(updatesReq AccountPayeeCheck, whereReq AccountPayeeCheck) error {
	db := driver.GormDb
	err := db.Debug().Table(apd.TableName()).Where(utils.BuildNotNullMap(whereReq)).Updates(updatesReq).Error
	if err != nil {
		return err
	}
	return nil
}

// UpdateByExample 更新 (没带的属性 "会" 添加至条件中在 DB NULL)
func (apd *AccountPayeeCheckDao) UpdateByExample(updatesReq AccountPayeeCheck, whereReq AccountPayeeCheck) error {
	db := driver.GormDb
	upReq := utils.BuildNullMap(updatesReq)
	//更新条件去掉id
	delete(upReq, "id")
	err := db.Debug().Table(apd.TableName()).Where(utils.BuildNotNullMap(whereReq)).Updates(upReq).Error
	if err != nil {
		return err
	}
	return nil
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
func (apd *AccountPayeeCheckDao) UpdateByPrimaryKey(id int, updatesReq AccountPayeeCheck) error {
	db := driver.GormDb
	upReq := utils.BuildNullMap(updatesReq)
	//更新条件去掉id
	delete(upReq, "id")
	err := db.Table(apd.TableName()).Where("id = ?", id).Updates(upReq).Error
	if err != nil {
		return err
	}
	return nil
}
