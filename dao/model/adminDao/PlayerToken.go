package adminDao

import "time"
import "gorm.io/gorm"
import "github.com/google/uuid"

// PlayerToken 用戶密碼 /**
type PlayerToken struct {
	ID         string    `gorm:"column:id;type:varchar(40);primary_key;comment:'token id'"`
	PlayerID   string    `gorm:"column:player_id;type:varchar(40);not null;comment:'玩家id'"`
	TokenType  int       `gorm:"column:token_type;type:int;not null;comment:'token類型'"`
	Token      string    `gorm:"column:token;type:text;not null;comment:'token'"`
	TKIndex    string    `gorm:"column:tk_index;type:varchar(500);comment:'索引用'"`
	ExpireTime time.Time `gorm:"column:expire_time;type:timestamp;default:null;comment:'失效時間'"`
	UpdateTime time.Time `gorm:"column:update_time;type:timestamp;not null;comment:'更新時間'"`
	CreateTime time.Time `gorm:"column:create_time;type:timestamp;not null;comment:'創建時間'"`
}

func (PlayerToken) TableName() string {
	return "member_playertoken"
}

func (g *PlayerToken) BeforeCreate(tx *gorm.DB) error {
	id := uuid.New().String()
	g.ID = id
	return nil
}
