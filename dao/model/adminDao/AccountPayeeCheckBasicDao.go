package adminDao

import (
	"AdminPro/common/enum"
	"AdminPro/common/model"
	"AdminPro/common/utils"
	"gorm.io/gorm"
)

type AccountPayeeCheckBasicDao struct {
}

// ListAccountPayeeChecks
func (dao *AccountPayeeCheckBasicDao) ListAccountPayeeChecks(userRandomId *string, status *enum.EAccountPayeeCheckStatusEnum, page *model.Pagination) (int64, []AccountPayeeCheck, error) {
	var results []AccountPayeeCheck
	customizeSQL := func(db *gorm.DB) *gorm.DB {
		if userRandomId != nil {
			db = db.Where("uid = ?", userRandomId)
		}
		if status != nil {
			db = db.Where("status = ?", status)
		}
		db = db.Scopes(utils.WithPagination(page.Page, page.Limit))
		db = db.Order("case when status = 0 then 0 else 1 end asc, created_time desc, id desc")
		return db
	}
	err := SelectByExample(customizeSQL, &results, &AccountPayeeCheck{})

	if err != nil {
		return 0, results, err
	}

	count, err := CountByExample(customizeSQL, &AccountPayeeCheck{})
	if err != nil {
		return count, nil, err
	}

	return count, results, nil
}
