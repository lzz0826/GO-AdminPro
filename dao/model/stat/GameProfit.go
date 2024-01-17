package stat

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"time"
)

// GameProfit 遊戲利潤統計 /**
type GameProfit struct {
	ID          string          `gorm:"column:id;type:varchar(40);primary_key;comment:'id'"`
	GameID      string          `gorm:"column:game_id;type:varchar(40);not null;comment:'遊戲ID'"`
	IssueID     string          `gorm:"column:issue_id;type:varchar(40);not null;comment:'遊戲ID'"`
	TotalProfit decimal.Decimal `gorm:"column:total_profit;type:decimal(20,5);not null;default:0.00000;comment:'總收益'"`
	TotalLoss   decimal.Decimal `gorm:"column:total_loss;type:decimal(20,5);not null;default:0.00000;comment:'總損失'"`
	TotalIncome decimal.Decimal `gorm:"column:total_income;type:decimal(20,5);not null;default:0.00000;comment:'總毛收益'"`
	CreateTime  time.Time       `gorm:"column:create_time;type:timestamp;not null;comment:'創建時間'"`
	UpdateTime  time.Time       `gorm:"column:update_time;type:timestamp;not null;comment:'更新時間'"`
}

func (GameProfit) TableName() string {
	return "stat_gameprofit"
}

func (g *GameProfit) BeforeCreate(tx *gorm.DB) error {
	id := uuid.New().String()
	g.ID = id
	return nil
}
