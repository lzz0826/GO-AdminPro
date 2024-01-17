package wallet

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"time"
)

// Wallet 錢包 /**
type Wallet struct {
	ID                  string          `gorm:"column:id;type:varchar(40);primary_key;comment:'錢包id'"`
	PlayerID            string          `gorm:"column:player_id;type:varchar(40);not null;comment:'關聯玩家id(與player表一對一)'"`
	Balance             decimal.Decimal `gorm:"column:balance;type:decimal(20,5);not null;default:0.00000;comment:'未鎖定餘額'"`
	TotalMFreezeBalance decimal.Decimal `gorm:"column:total_m_freeze_balance;type:decimal(20,5);not null;default:0.00000;comment:'總凍結餘額(待減)'"`
	TotalPFreezeBalance decimal.Decimal `gorm:"column:total_p_freeze_balance;type:decimal(20,5);default:0.00000;comment:'總凍結餘額(待加)'"`
	CreateTime          time.Time       `gorm:"column:create_time;type:timestamp;not null;comment:'創建時間'"`
	UpdateTime          time.Time       `gorm:"column:update_time;type:timestamp;not null;comment:'更新時間'"`
	Memo                string          `gorm:"column:memo;type:text;comment:'備註'"`
	IsEnabled           bool            `gorm:"column:is_enabled;type:tinyint(1);not null;default:1;comment:'是否啟用'"`
}

func (Wallet) TableName() string {
	return "wallet_wallet"
}

func (g *Wallet) BeforeCreate(tx *gorm.DB) error {
	id := uuid.New().String()
	g.ID = id
	return nil
}
