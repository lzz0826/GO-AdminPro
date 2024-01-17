package wallet

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"time"
)

// CreditLog 信用分紀錄 /**
type CreditLog struct {
	ID                  string          `gorm:"column:id;type:varchar(40);primary_key;comment:'信用紀錄id'"`
	UserID              string          `gorm:"column:user_id;type:varchar(40);not null;comment:'代理id'"`
	CreditID            string          `gorm:"column:credit_id;type:varchar(40);not null;comment:'代理信用id'"`
	TranslogID          string          `gorm:"column:translog_id;type:varchar(40);comment:'交易紀錄id'"`
	WalletID            string          `gorm:"column:wallet_id;type:varchar(40);comment:'關聯的玩家錢包紀錄'"`
	PlayerID            string          `gorm:"column:player_id;type:varchar(40);comment:'當時關聯玩家id'"`
	CreditlogType       int             `gorm:"column:creditlog_type;type:int;not null;default:1;comment:'信用紀錄類型'"`
	CreditBefore        decimal.Decimal `gorm:"column:credit_before;type:decimal(20,5);not null;comment:'代理信用變更前'"`
	CreditAfter         decimal.Decimal `gorm:"column:credit_after;type:decimal(20,5);not null;comment:'代理信用變更後'"`
	MFreezeCreditAfter  decimal.Decimal `gorm:"column:m_freeze_credit_after;type:decimal(20,5);not null;default:0.00000;comment:'待減凍結信用餘額(後)'"`
	MFreezeCreditBefore decimal.Decimal `gorm:"column:m_freeze_credit_before;type:decimal(20,5);not null;default:0.00000;comment:'待減凍結信用餘額(前)'"`
	PFreezeCreditAfter  decimal.Decimal `gorm:"column:p_freeze_credit_after;type:decimal(20,5);not null;default:0.00000;comment:'待加凍結信用餘額(後)'"`
	PFreezeCreditBefore decimal.Decimal `gorm:"column:p_freeze_credit_before;type:decimal(20,5);not null;default:0.00000;comment:'待加凍結信用餘額(前)'"`
	Amount              decimal.Decimal `gorm:"column:amount;type:decimal(20,5);not null;default:0.00000;comment:'變更數量'"`
	CreditlogStatus     int             `gorm:"column:creditlog_status;type:int;not null;default:1;comment:'信用紀錄狀態'"`
	CompleteTime        time.Time       `gorm:"column:complete_time;type:timestamp;default:null;comment:'完成時間'"`
	CancelTime          time.Time       `gorm:"column:cancel_time;type:timestamp;default:null;comment:'取消時間'"`
	Memo                string          `gorm:"column:memo;type:text;comment:'備註'"`
	CreateTime          time.Time       `gorm:"column:create_time;type:timestamp;not null;comment:'創建時間'"`
	UpdateTime          time.Time       `gorm:"column:update_time;type:timestamp;not null;comment:'更新時間'"`
}

func (CreditLog) TableName() string {
	return "wallet_creditlog"
}

func (g *CreditLog) BeforeCreate(tx *gorm.DB) error {
	id := uuid.New().String()
	g.ID = id
	return nil
}
