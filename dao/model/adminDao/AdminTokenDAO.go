package adminDao

import (
	"AdminPro/common/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

type AdminTokenDAO struct {
	ID         string    `gorm:"column:id;type:varchar(40);primary_key;comment:'id'" json:"id"`
	AdminID    string    `gorm:"column:admin_id;type:varchar(40);not null;comment:'代理id'" json:"admin_id"`
	TokenType  int       `gorm:"column:token_type;type:int;comment:'token類型'" json:"token_type"`
	Token      string    `gorm:"column:token;type:text;not null;comment:'token'" json:"token"`
	ExpireTime time.Time `gorm:"column:expire_time;type:timestamp;comment:'過期時間'" json:"expire_time"`
	UpdateTime time.Time `gorm:"column:update_time;type:timestamp;comment:'更新時間'" json:"update_time"`
	CreateTime time.Time `gorm:"column:create_time;type:timestamp;not null;comment:'創建時間'" json:"create_time"`
	CreatorID  string    `gorm:"column:creator_id;type:varchar(40);comment:'創建者id'" json:"creator_id"`
	UpdaterID  string    `gorm:"column:updater_id;type:varchar(40);comment:'更新者id'" json:"updater_id"`
}

func (at *AdminTokenDAO) TableName() string {
	return "admin_admintoken"
}

func (at *AdminTokenDAO) InsertAdminToken(tx *gorm.DB) error {
	return tx.Table(at.TableName()).Omit("id").Create(at).Error
}

// GetAdminTokenByID 根據 ID 查詢 AdminTokenDAO
func (at *AdminTokenDAO) GetAdminTokenByID(id string) (adminToken AdminTokenDAO, err error) {
	err = mysql.GormDb.Table(at.TableName()).Where("id = ?", id).First(&adminToken).Error
	if err != nil {
		log.Println("GetAdminTokenByID errors:", err.Error())
		return
	}
	return
}

func (at *AdminTokenDAO) GetAdminTokenByAdminID(adminID string) (adminToken AdminTokenDAO, err error) {
	err = mysql.GormDb.Table(at.TableName()).Where("admin_id = ?", adminID).First(&adminToken).Error
	if err != nil {
		log.Println("GetAdminTokenByID errors:", err.Error())
		return
	}
	return
}

// GetAdminTokensByAdminID 根據 AdminID 查詢 AdminTokens
func (at *AdminTokenDAO) GetAdminTokensByAdminID(adminID string) (adminTokens []AdminTokenDAO, err error) {
	err = mysql.GormDb.Table(at.TableName()).Where("admin_id = ?", adminID).Find(&adminTokens).Error
	if err != nil {
		log.Println("GetAdminTokensByAdminID errors:", err.Error())
		return
	}
	return
}
