package game

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

// Market 盤口 /**
type Market struct {
	ID         string    `gorm:"column:id;type:varchar(40);primary_key;comment:'盤口id'"`
	GameID     string    `gorm:"column:game_id;type:varchar(40);not null;comment:'遊戲id'"`
	Title      string    `gorm:"column:title;type:varchar(40);not null;comment:'盤口名稱'"`
	Narrative  string    `gorm:"column:narrative;type:text;comment:'說明'"`
	UpdateTime time.Time `gorm:"column:update_time;type:timestamp;not null;comment:'更新時間'"`
	CreateTime time.Time `gorm:"column:create_time;type:timestamp;not null;comment:'創建時間'"`
	Priority   int       `gorm:"column:priority;type:int;not null;comment:'順序'"`
	IsEnabled  bool      `gorm:"column:is_enabled;type:tinyint(1);default:1;comment:'是否啟用'"`
}

func (Market) TableName() string {
	return "game_market"
}

func (g *Market) BeforeCreate(tx *gorm.DB) error {
	id := uuid.New().String()
	g.ID = id
	return nil
}
