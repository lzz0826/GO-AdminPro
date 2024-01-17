package adminDao

import (
	"AdminPro/common/driver"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"log"
	"time"
)

// models package's db obj
// all db operation should be done in models pkg
// so db is a pkg inner var
var db *gorm.DB = driver.GormDb

// User 代理 /**
type User struct {
	ID         string    `gorm:"column:id;type:varchar(40);primary_key;comment:'代理id'"`
	Username   string    `gorm:"column:username;type:varchar(40);not null;comment:'帳號'"`
	Avatar     string    `gorm:"column:avatar;type:varchar(100);comment:'頭像'"`
	UpdateTime time.Time `gorm:"column:update_time;type:timestamp;not null;comment:'更新時間'"`
	CreateTime time.Time `gorm:"column:create_time;type:timestamp;not null;comment:'創建時間'"`
}

func (User) TableName() string {
	return "member_user"
}

func (model *User) BeforeCreate(tx *gorm.DB) error {
	id := uuid.New().String()
	model.ID = id
	return nil
}

func (model *User) GetById(id string) (user User, err error) {
	// find one record
	err = db.Table("member_user").Where("id = ?", id).First(&user).Error

	if err != nil {
		log.Println(err.Error())
		return
	}

	return
}
