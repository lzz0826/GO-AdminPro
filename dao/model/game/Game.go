package game

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

// Game 遊戲 /**
type Game struct {
	ID                string    `gorm:"column:id;type:varchar(40);primary_key;comment:'遊戲id'"`
	CategoryID        string    `gorm:"column:category_id;type:varchar(40);not null;comment:'類型id'"`
	Title             string    `gorm:"column:title;type:varchar(40);not null;comment:'遊戲名稱'"`
	GameStartFrom     time.Time `gorm:"column:game_start_from;type:time;not null;default:'00:00:00';comment:'從一天中的幾點開始(排程生成用)'"`
	GameEndTo         time.Time `gorm:"column:game_end_to;type:time;not null;default:'23:59:59';comment:'從一天中的幾點結束(排程生成用)'"`
	GamePeriod        time.Time `gorm:"column:game_period;type:time;not null;comment:'盤口期間(排程生成用)'"`
	GameClass         string    `gorm:"column:game_class;type:varchar(255);not null;comment:'遊戲類'"`
	CloseBeforeStart  time.Time `gorm:"column:close_before_start;type:time;not null;comment:'封盤時間(排程生成用)'"`
	DrawBeforeStart   time.Time `gorm:"column:draw_before_start;type:time;not null;comment:'開獎時間(排程生成用)'"`
	SettleBeforeStart time.Time `gorm:"column:settle_before_start;type:time;default:null;comment:'結算時間(排程生成用)'"`
	IsEnabled         bool      `gorm:"column:is_enabled;type:tinyint;not null;default:1;comment:'是否啟用'"`
	UpdateTime        time.Time `gorm:"column:update_time;type:timestamp;not null;comment:'更新時間'"`
	CreateTime        time.Time `gorm:"column:create_time;type:timestamp;not null;comment:'創建時間'"`
	Memo              string    `gorm:"column:memo;type:text;comment:'備註'"`
	Narrative         string    `gorm:"column:narrative;type:text;comment:'敘述(遊戲說明)'"`
}

func (Game) TableName() string {
	return "game_game"
}

func (g *Game) BeforeCreate(tx *gorm.DB) error {
	id := uuid.New().String()
	g.ID = id
	return nil
}
