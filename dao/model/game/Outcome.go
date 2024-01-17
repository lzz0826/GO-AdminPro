package game

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"time"
)

// Outcome 盤口結果 /**
type Outcome struct {
	ID                 string          `gorm:"column:id;type:varchar(40);primary_key;comment:'玩法id'"`
	GameID             string          `gorm:"column:game_id;type:varchar(40);not null;comment:'遊戲id'"`
	MarketID           string          `gorm:"column:market_id;type:varchar(40);not null;comment:'盤口id'"`
	Title              string          `gorm:"column:title;type:varchar(40);not null;comment:'玩法名稱'"`
	Narrative          string          `gorm:"column:narrative;type:text;comment:'描述'"`
	Odds               decimal.Decimal `gorm:"column:odds;type:decimal(20,5);not null;comment:'賠率'"`
	UpdateTime         time.Time       `gorm:"column:update_time;type:timestamp;not null;comment:'更新時間'"`
	CreateTime         time.Time       `gorm:"column:create_time;type:timestamp;not null;comment:'創建時間'"`
	IsEnabled          bool            `gorm:"column:is_enabled;type:tinyint(1);not null;default:true;comment:'是否啟用'"`
	SettlementFunction string          `gorm:"column:settlement_function;type:varchar(60);not null;comment:'開獎邏輯對應的function'"`
	Priority           int             `gorm:"column:priority;type:int;not null;default:0;comment:'優先順序'"`
}

func (Outcome) TableName() string {
	return "game_outcome"
}

func (g *Outcome) BeforeCreate(tx *gorm.DB) error {
	id := uuid.New().String()
	g.ID = id
	return nil
}
