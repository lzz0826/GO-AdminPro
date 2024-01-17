package game

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

// Issue 期數 /**
type Issue struct {
	ID                   string    `gorm:"column:id;type:varchar(40);primary_key;comment:'期數id'"`
	GameID               string    `gorm:"column:game_id;type:varchar(40);not null;comment:'遊戲id'"`
	SettlementID         string    `gorm:"column:settlement_id;type:varchar(40);default:null;comment:'結算id'"`
	Memo                 string    `gorm:"column:memo;type:text;comment:'備註'"`
	StartTime            time.Time `gorm:"column:start_time;type:timestamp;default:null;comment:'開盤時間'"`
	CloseTime            time.Time `gorm:"column:close_time;type:timestamp;default:null;comment:'封盤時間'"`
	ScheduledDrawTime    time.Time `gorm:"column:scheduled_draw_time;type:timestamp;not null;comment:'預定開獎時間'"`
	ActualDrawTime       time.Time `gorm:"column:actual_draw_time;type:timestamp;default:null;comment:'實際開獎時間'"`
	ScheduledSettledTime time.Time `gorm:"column:scheduled_settled_time;type:timestamp;not null;comment:'預定結算時間'"`
	ActualSettledTime    time.Time `gorm:"column:actual_settled_time;type:timestamp;default:null;comment:'實際結算時間'"`
	IsEnabled            bool      `gorm:"column:is_enabled;type:tinyint(1);not null;default:true;comment:'是否啟用'"`
	UpdateTime           time.Time `gorm:"column:update_time;type:timestamp;not null;comment:'更新時間'"`
	CreateTime           time.Time `gorm:"column:create_time;type:timestamp;not null;comment:'創建時間'"`
}

func (Issue) TableName() string {
	return "game_issue"
}

func (g *Issue) BeforeCreate(tx *gorm.DB) error {
	id := uuid.New().String()
	g.ID = id
	return nil
}
