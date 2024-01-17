package adminDao

import (
	"AdminPro/common/driver"
	_ "gorm.io/gorm"
	"log"
	"time"
)

type AdminDAO struct {
	ID            string    `gorm:"column:id;type:varchar(40);primary_key;comment:'管理用戶ID'" json:"id"`
	ChannelID     string    `gorm:"column:channel_id;type:varchar(40)" json:"channel_id"`
	Username      string    `gorm:"column:username;type:varchar(40);not null;comment:'帳號'" json:"username"`
	AdminName     string    `gorm:"column:admin_name;type:varchar(40)" json:"admin_name"`
	Nickname      string    `gorm:"column:nickname;type:varchar(40)" json:"nickname"`
	AccountStatus int       `gorm:"column:account_status;type:int;not null;default:0;comment:'狀態'" json:"account_status"`
	LoginIP       string    `gorm:"column:login_ip;type:varchar(40)" json:"login_ip"`
	LoginTime     time.Time `gorm:"column:login_time" json:"login_time"`
	Memo          string    `gorm:"column:memo;type:text" json:"memo"`
	CreatorID     string    `gorm:"column:creator_id;type:varchar(40)" json:"creator_id"`
	UpdaterID     string    `gorm:"column:updater_id;type:varchar(40)" json:"updater_id"`
	UpdateTime    time.Time `gorm:"column:update_time;not null" json:"update_time"`
	CreateTime    time.Time `gorm:"column:create_time;not null" json:"create_time"`
	TwoFactorCode string    `gorm:"column:two_factor_code;type:varchar(40)" json:"two_factor_code"`
}

func (AdminDAO) TableName() string {
	return "admin_admin"
}

// InsertAdmin 插入 AdminDAO 資料
func (model *AdminDAO) InsertAdmin() (err error) {
	err = driver.GormDb.Table(model.TableName()).Create(model).Error
	if err != nil {
		log.Println(err.Error())
		return
	}
	return
}

// GetAdminByID 根據 ID 查詢 AdminDAO
func (model *AdminDAO) GetAdminByID(id string) (admin AdminDAO, err error) {
	err = driver.GormDb.Table(model.TableName()).Where("id = ?", id).First(&admin).Error
	if err != nil {
		log.Println(err.Error())
		return
	}
	return
}

func (model *AdminDAO) GetAdminByUsername(username string) (admin AdminDAO, err error) {
	err = driver.GormDb.Table(model.TableName()).Where("username = ?", username).First(&admin).Error
	if err != nil {
		log.Println(err.Error())
		return
	}
	return
}
