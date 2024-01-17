package wallet

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"time"
)

// Translog 交易紀錄 /**
type Translog struct {
	ID             string          `gorm:"column:id;type:varchar(40);primary_key;comment:'交易紀錄id'"`
	WalletID       string          `gorm:"column:wallet_id;type:varchar(40);not null;comment:'錢包id'"`
	PlayerID       string          `gorm:"column:player_id;type:varchar(40);not null;comment:'玩家id'"`
	UserID         string          `gorm:"column:user_id;type:varchar(40);not null;comment:'當時關聯代理id'"`
	CreditlogID    string          `gorm:"column:creditlog_id;type:varchar(40);comment:'信用紀錄id'"`
	TranslogType   int             `gorm:"column:translog_type;type:int;not null;default:1;comment:'交易類型'"`
	BalanceBefore  decimal.Decimal `gorm:"column:balance_before;type:decimal(20,5);not null;comment:'餘額變更前'"`
	BalanceAfter   decimal.Decimal `gorm:"column:balance_after;type:decimal(20,5);not null;comment:'餘額變更後'"`
	MFreezeBalance decimal.Decimal `gorm:"column:m_freeze_balance;type:decimal(20,5);not null;default:0.00000;comment:'凍結餘額(待減)'"`
	PFreezeBalance decimal.Decimal `gorm:"column:p_freeze_balance;type:decimal(20,5);not null;default:0.00000;comment:'凍結餘額(待加)'"`
	Amount         decimal.Decimal `gorm:"column:amount;type:decimal(20,5);not null;comment:'變更數量'"`
	TranslogStatus int             `gorm:"column:translog_status;type:int;not null;default:1;comment:'交易紀錄狀態'"`
	CompleteTime   time.Time       `gorm:"column:complete_time;type:timestamp;default:null;comment:'完成時間'"`
	CancelTime     time.Time       `gorm:"column:cancel_time;type:timestamp;default:null;comment:'取消時間'"`
	Memo           string          `gorm:"column:memo;type:text;comment:'備註'"`
	CreateTime     time.Time       `gorm:"column:create_time;type:timestamp;not null;comment:'創建時間'"`
	UpdateTime     time.Time       `gorm:"column:update_time;type:timestamp;not null;comment:'更新時間'"`
}

func (Translog) TableName() string {
	return "wallet_translog"
}

func (g *Translog) BeforeCreate(tx *gorm.DB) error {
	id := uuid.New().String()
	g.ID = id
	return nil
}
