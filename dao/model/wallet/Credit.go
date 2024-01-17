package wallet

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"time"
)

// Credit 信用分 /**
type Credit struct {
	ID            string          `gorm:"column:id;type:varchar(40);primary_key;comment:'代理信用額度id'"`
	UserID        string          `gorm:"column:user_id;type:varchar(40);not null;comment:'關聯代理id(與user表一對一)'"`
	Credit        decimal.Decimal `gorm:"column:credit;type:decimal(20,5);not null;default:0.00000;comment:'代理信用額度'"`
	MFreezeCredit decimal.Decimal `gorm:"column:m_freeze_credit;type:decimal(20,5);not null;default:0.00000;comment:'總凍結代理信用額度(待減)'"`
	PFreezeCredit decimal.Decimal `gorm:"column:p_freeze_credit;type:decimal(20,5);not null;default:0.00000;comment:'總凍結代理信用額度(待加)'"`
	CreateTime    time.Time       `gorm:"column:create_time;type:timestamp;not null;comment:'創建時間'"`
	UpdateTime    time.Time       `gorm:"column:update_time;type:timestamp;not null;comment:'更新時間'"`
	Memo          string          `gorm:"column:memo;type:text;comment:'memo'"`
	IsEnabled     int             `gorm:"column:is_enabled;type:tinyint;not null;default:1;comment:'是否啟用'"`
}

func (Credit) TableName() string {
	return "wallet_credit"
}

func (g *Credit) BeforeCreate(tx *gorm.DB) error {
	id := uuid.New().String()
	g.ID = id
	return nil
}
