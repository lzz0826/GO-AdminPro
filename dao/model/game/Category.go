package game

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

// Category 遊戲類別 /**
type Category struct {
	ID         string    `gorm:"column:id;type:varchar(40);primary_key;comment:'類型id'"`
	Title      string    `gorm:"column:title;type:varchar(40);not null;comment:'類型名稱'"`
	Narrative  string    `gorm:"column:narrative;type:text;comment:'說明'"`
	CreateTime time.Time `gorm:"column:create_time;type:timestamp;not null;comment:'創建時間'"`
	UpdateTime time.Time `gorm:"column:update_time;type:timestamp;default:null;comment:'更新時間'"`
}

func (Category) TableName() string {
	return "game_category"
}

func (g *Category) BeforeCreate(tx *gorm.DB) error {
	id := uuid.New().String()
	g.ID = id
	return nil
}
