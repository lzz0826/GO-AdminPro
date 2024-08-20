package adminDao

import (
	"time"
)

type AccountPayeeCheck struct {
	ID          *int       `gorm:"column:id;primaryKey;autoIncrement;comment:'id'" json:"id"`
	UID         *int       `gorm:"column:uid;comment:'用户id'" json:"uid"`
	Type        *int       `gorm:"column:type;comment:'审核类型1-银行卡,2-支付保'" json:"type"`
	Description *string    `gorm:"column:description;comment:'审核内容" json:"description"`
	Status      *int       `gorm:"column:status;comment:'??(0-审核中,1-审核成功,2-省合失败)'" json:"status"`
	CheckID     *int       `gorm:"column:check_id;comment:'审核人id'" json:"checkId"`
	CheckTime   *time.Time `gorm:"column:check_time;comment:'审核时间'" json:"checkTime"`
	UpdateTime  *time.Time `gorm:"column:update_time;comment:'更新时间'" json:"updateTime"`
	CreatedTime *time.Time `gorm:"column:created_time;comment:'创建时间'" json:"createdTime"`
}


func (apd *AccountPayeeCheck) GetTableName() string {
	return "account_payee_check"
}