package adminDao

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

// UserPlayerToken 代理密碼 /**
type UserPlayerToken struct {
	ID         string    `gorm:"column:id;type:varchar(40);primary_key;comment:'id'"`
	UserID     string    `gorm:"column:user_id;type:varchar(40);not null;comment:'代理id'"`
	TokenType  int       `gorm:"column:token_type;type:int;comment:'token類型'"`
	Token      string    `gorm:"column:token;type:text;not null;comment:'token'"`
	ExpireTime time.Time `gorm:"column:expire_time;type:timestamp;default:null;comment:'過期時間'"`
	UpdateTime time.Time `gorm:"column:update_time;type:timestamp;default:null;comment:'更新時間'"`
	CreateTime time.Time `gorm:"column:create_time;type:timestamp;not null;comment:'創建時間'"`
}

func (UserPlayerToken) TableName() string {
	return "member_userplayertoken"
}

func (g *UserPlayerToken) BeforeCreate(tx *gorm.DB) error {
	id := uuid.New().String()
	g.ID = id
	return nil
}
