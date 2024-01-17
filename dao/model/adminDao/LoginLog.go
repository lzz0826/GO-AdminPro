package adminDao

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

// LoginLog 登入紀錄 /**
type LoginLog struct {
	ID          string    `gorm:"column:id;type:varchar(40);primary_key;comment:'id'"`
	PlayerID    string    `gorm:"column:player_id;type:varchar(40);not null;comment:'玩家id'"`
	UserID      string    `gorm:"column:user_id;type:varchar(40);comment:'當下的代理id'"`
	Token       string    `gorm:"column:token;type:text;not null;comment:'當下的token'"`
	LoginSource int       `gorm:"column:login_source;type:int;not null;comment:'登入來源'"`
	LoginType   int       `gorm:"column:login_type;type:int;not null;comment:'登入類型'"`
	LoginTime   time.Time `gorm:"column:login_time;type:timestamp;not null;comment:'登入時間'"`
	CreateTime  time.Time `gorm:"column:create_time;type:timestamp;not null;comment:'創建時間'"`
	UpdateTime  time.Time `gorm:"column:update_time;type:timestamp;default:null;comment:'更新時間'"`
}

func (LoginLog) TableName() string {
	return "member_loginlog"
}

func (g *LoginLog) BeforeCreate(tx *gorm.DB) error {
	id := uuid.New().String()
	g.ID = id
	return nil
}
