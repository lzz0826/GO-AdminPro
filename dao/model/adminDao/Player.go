package adminDao

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

// Player 會員 /**
type Player struct {
	ID         string    `gorm:"column:id;type:varchar(40);primary_key;comment:'玩家id'"`
	UserID     string    `gorm:"column:user_id;type:varchar(40);comment:'所屬代理id'"`
	Username   string    `gorm:"column:username;type:varchar(40);comment:'玩家帳號'"`
	PlayerName string    `gorm:"column:player_name;type:varchar(30);comment:'玩家名稱'"`
	Avatar     string    `gorm:"column:avatar;type:varchar(100);comment:'頭像'"`
	CreateTime time.Time `gorm:"column:create_time;type:timestamp;not null;comment:'創建時間'"`
	UpdateTime time.Time `gorm:"column:update_time;type:timestamp;not null;comment:'更新時間'"`
}

func (Player) TableName() string {
	return "member_player"
}

func (g *Player) BeforeCreate(tx *gorm.DB) error {
	id := uuid.New().String()
	g.ID = id
	return nil
}
