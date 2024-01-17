package order

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"time"
)

// Betslip 注單 /**
type Betslip struct {
	ID                string          `gorm:"column:id;type:varchar(40);primary_key;comment:'訂單id'"`
	IssueID           string          `gorm:"column:issue_id;type:varchar(40);not null;comment:'期數id'"`
	GameID            string          `gorm:"column:game_id;type:varchar(40);not null;comment:'對應的遊戲id'"`
	MarketID          string          `gorm:"column:market_id;type:varchar(40);not null;comment:'對應的盤口id'"`
	OutcomeID         string          `gorm:"column:outcome_id;type:varchar(40);not null;comment:'對應的玩法id'"`
	SettlementID      string          `gorm:"column:settlement_id;type:varchar(40);comment:'結果id'"`
	BetTranslogID     string          `gorm:"column:bet_translog_id;type:varchar(40);comment:'下注時的交易id'"`
	SettledTranslogID string          `gorm:"column:settled_translog_id;type:varchar(40);comment:'結算的時候的交易id(如果有的話)'"`
	PlayerID          string          `gorm:"column:player_id;type:varchar(40);not null;comment:'玩家id'"`
	UserID            string          `gorm:"column:user_id;type:varchar(40);not null;comment:'用戶所屬代理'"`
	Odds              decimal.Decimal `gorm:"column:odds;type:decimal(20,5);not null;comment:'賠率'"`
	Result            string          `gorm:"column:result;type:varchar(100);comment:'結果'"`
	SettledStatus     int             `gorm:"column:settled_status;type:int;default:1;comment:'結算狀態'"`
	BetStatus         int             `gorm:"column:bet_status;type:int;not null;comment:'注單狀態'"`
	Price             decimal.Decimal `gorm:"column:price;type:decimal(20,5);not null;comment:'金額'"`
	Memo              string          `gorm:"column:memo;type:varchar(255);comment:'備註'"`
	Detail            string          `gorm:"column:detail;type:text;comment:'其他雜項'"`
	PreDrawCode       string          `gorm:"column:pre_draw_code;type:varchar(100);comment:'預開獎'"`
	SettledTime       time.Time       `gorm:"column:settled_time;type:timestamp;default:null;comment:'結算時間'"`
	CreateTime        time.Time       `gorm:"column:create_time;type:timestamp;not null;comment:'創建時間'"`
	UpdateTime        time.Time       `gorm:"column:update_time;type:timestamp;not null;comment:'更新時間'"`
}

func (Betslip) TableName() string {
	return "order_betslip"
}

func (g *Betslip) BeforeCreate(tx *gorm.DB) error {
	id := uuid.New().String()
	g.ID = id
	return nil
}
