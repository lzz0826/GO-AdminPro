package game

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"time"
)

// BetStake 注碼 /**
type BetStake struct {
	ID         string          `gorm:"column:id;type:varchar(40);primary_key;comment:'注碼id'"`
	GameID     string          `gorm:"column:game_id;type:varchar(40);not null;comment:'遊戲id'"`
	Title      string          `gorm:"column:title;type:varchar(40);not null;comment:'標題'"`
	Stake      decimal.Decimal `gorm:"column:stake;type:decimal(20,5);not null;comment:'注碼'"`
	IsEnabled  bool            `gorm:"column:is_enabled;type:tinyint;not null;default:1;comment:'是否啟用'"`
	UpdateTime time.Time       `gorm:"column:update_time;type:timestamp;not null;comment:'更新時間'"`
	CreateTime time.Time       `gorm:"column:create_time;type:timestamp;not null;comment:'創建時間'"`
}

func (BetStake) TableName() string {
	return "game_betstake"
}

func (g *BetStake) BeforeCreate(tx *gorm.DB) error {
	id := uuid.New().String()
	g.ID = id
	return nil
}
