package adminDao

import (
	"AdminPro/common/enum"
	"AdminPro/common/model"
	"AdminPro/common/mysql"
	"fmt"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"strings"
	"time"
)

type AccountPayeeCheckBasicDao struct {
	BasicDao
}

// 使用分页判断
func (dao *AccountPayeeCheckBasicDao) SelectCheckPageTest(userRandomId *string, status *enum.EAccountPayeeCheckStatusEnum) ([]AccountPayeeCheck, error) {
	var results []AccountPayeeCheck
	db := mysql.GormDb
	if userRandomId != nil {
		db = db.Where("uid = ?", userRandomId)
	}
	if status != nil {
		db = db.Where("status = ?", status)
	}
	db = db.Order("case when status = 0 then 0 else 1 end asc, created_time desc, id desc")
	err := dao.SelectCheckPage(db, &results)
	if err != nil {
		return results, err
	}
	return results, nil
}

func (dao *AccountPayeeCheckBasicDao) JoinSelectCustomizeSqlCheckPage(checkId int, search *string) ([]Join, error) {
	var results []Join
	db := mysql.GormDb

	db = db.Table("account_payee_check t1")
	db = db.Select("t2.username as username, t1.id, t1.description as description")
	db = db.Joins("JOIN admin_admin t2 ON t1.uid = t2.id")
	db = db.Where("t1.check_id = ?", checkId)
	if search != nil {
		db = db.Where("(t1.description LIKE ? OR t2.nickname LIKE ?)", "%"+*search+"%", "%"+*search+"%")
	}
	db = db.Order("t1.created_time DESC")
	err := dao.SelectCheckPage(db, &results)
	if err != nil {
		return results, err
	}
	return results, nil
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
	db := mysql.GormDb
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
	db := mysql.GormDb
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
func ListAccountPayeeChecks(userRandomId *string, status *enum.EAccountPayeeCheckStatusEnum) ([]AccountPayeeCheck, error) {
	var results []AccountPayeeCheck
	db := mysql.GormDb

	if userRandomId != nil {
		db = db.Where("uid = ?", userRandomId)
	}
	if status != nil {
		db = db.Where("status = ?", status)
	}
	db = db.Order("case when status = 0 then 0 else 1 end asc, created_time desc, id desc")
	err := Select(db, &results)
	if err != nil {
		return results, err
	}

	return results, nil
}

func ListAccountPayeeChecksPage(userRandomId *string, status *enum.EAccountPayeeCheckStatusEnum, page *model.Pagination) (int64, []AccountPayeeCheck, error) {
	var results []AccountPayeeCheck
	db := mysql.GormDb
	if userRandomId != nil {
		//db = db.Select("description")
		db = db.Where("uid = ?", userRandomId)
	}
	if status != nil {
		db = db.Where("status = ?", status)
	}
	db = db.Order("case when status = 0 then 0 else 1 end asc, created_time desc, id desc")
	total, err := SelectPage(db, &results, page)

	if err != nil {
		return 0, results, err
	}
	return total, results, nil
}

func SumTotalStatusSUM() (*decimal.Decimal, error) {
	var totalAmount decimal.Decimal
	db := mysql.GormDb.Debug()
	db = db.Table("account_payee_check").Select("IFNULL(SUM(status), 0) AS total_amount")
	err := Select(db, &totalAmount)
	if err != nil {
		return nil, fmt.Errorf("failed to convert total amount to decimal: %w", err)
	}
	return &totalAmount, nil
}

// JOIN 测试
func TestJoin(checkId, start, size int, search *string) ([]Join, error) {
	var results []Join
	db := mysql.GormDb
	query := db.Table("account_payee_check t1").
		Select("t2.username as username, t1.id, t1.description as description").
		Joins("JOIN admin_admin t2 ON t1.uid = t2.id").
		Where("t1.check_id = ?", checkId)
	if search != nil {
		query = query.Where("(t1.description LIKE ? OR t2.nickname LIKE ?)", "%"+*search+"%", "%"+*search+"%")
	}
	query = query.Order("t1.created_time DESC").Offset(start).Limit(size)
	if err := query.Debug().Find(&results).Error; err != nil {
		return nil, err
	}
	return results, nil
}

// 子查询 测试
func TestSubquery(search string) (AccountPayeeCheck, error) {
	var result AccountPayeeCheck
	db := mysql.GormDb
	// 子查询获取 club_record 的 id
	subquery := db.Table("club_record").
		Select("id").
		Where("random_id = ?", search).
		Limit(1)

	// 主查询在 pay_channel 和 club_pay_channel 中进行连接查询
	query := db.Table("pay_channel t1").
		Select("t1.id, t1.name"). // 如果需要其他字段，这里添加
		Joins("JOIN club_pay_channel t2 ON t1.id = t2.channel_id").
		Where("t2.club_id = (?)", subquery)

	// 获取查询结果
	if err := query.Debug().First(&result).Error; err != nil {
		return result, err
	}
	return result, nil
}

// 测试 Raw subquery 子查询
func TestRawSubquery(transType int) (int64, error) {
	var result int64
	db := mysql.GormDb
	subquery := "SELECT club_id FROM club_pay_channel"

	db = db.Table("club_record").
		Where("club_status = ?", 0).
		Where("trans_type = ?", transType).
		Where("id IN (?)", db.Raw(subquery))

	if err := db.Debug().Count(&result).Error; err != nil {
		return 0, err
	}
	return result, nil
}

// 测试子查询更新
func UpdateAccountStatusFoAdminAccountStatus(id int) (int64, error) {
	db := mysql.GormDb
	// 构建子查询
	subquery := db.Table("admin_admin t1").
		Select("t1.account_status").
		Where("t1.id = ?", id)

	// 更新操作
	query := db.Debug().Table("account_payee_check t2").
		Where("t2.type = (?)", subquery).
		Update("t2.status", 7)

	affected := query.RowsAffected
	// 执行并检查错误
	if err := query.Error; err != nil {
		return 0, err
	}
	return affected, nil
}

// 测试指针带入
func TestUpdateByExampleSelectivePoint(uid int, description string) (int64, error) {
	updatesReq := AccountPayeeCheck{
		UID: uid,
	}
	whereReq := map[string]interface{}{"description": description}
	db = mysql.GormDb
	result, err := Updates(db, updatesReq, whereReq)

	if err != nil {
		return 0, err
	}
	return result, nil
}

// TODO
func BatchInsertChannelLog(db *gorm.DB, payChannelId int, clubIds []int, logType int, createBy int, createTime time.Time) error {
	var values []string
	for _, clubID := range clubIds {
		values = append(values, fmt.Sprintf("(%d, %d, %d, %d, '%s')", payChannelId, clubID, logType, createBy, createTime.Format("2006-01-02 15:04:05")))
	}
	valuesStr := strings.Join(values, ", ")
	sql := fmt.Sprintf(
		"INSERT INTO club_pay_channel_log (pay_channel_id, club_id, type, op_id, created_time) VALUES %s",
		valuesStr,
	)
	if err := db.Exec(sql).Error; err != nil {
		return err
	}

	return nil
}
