package adminDao

import (
	"AdminPro/common/driver"
	"log"
	"time"
)

type AdminPermitDAO struct {
	ID         string    `gorm:"column:id;type:varchar(40);primary_key;comment:'id'" json:"id"`
	AdminID    string    `gorm:"column:admin_id;type:varchar(40);not null;comment:'adminVo id'" json:"admin_id"`
	PermitID   string    `gorm:"column:permit_id;type:varchar(40);not null;comment:'permit id'" json:"permit_id"`
	CreatorID  string    `gorm:"column:creator_id;type:varchar(40);not null;comment:'創建者id'" json:"creator_id"`
	UpdaterID  string    `gorm:"column:updater_id;type:varchar(40);not null;comment:'更新者id'" json:"updater_id"`
	CreateTime time.Time `gorm:"column:create_time;not null;comment:'創建時間'" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time;not null;comment:'更新時間'" json:"update_time"`
}

func (ap *AdminPermitDAO) TableName() string {
	return "admin_admin_permit"
}

// InsertAdminPermit 插入 AdminPermitDAO 資料
func (ap *AdminPermitDAO) InsertAdminPermit() (err error) {
	err = driver.GormDb.Table(ap.TableName()).Create(ap).Error
	if err != nil {
		log.Println(err.Error())
		return
	}
	return nil
}

// GetAdminPermitByID 根據 ID 查詢 AdminPermitDAO
func (ap *AdminPermitDAO) GetAdminPermitByID(id string) (adminPermit AdminPermitDAO, err error) {
	err = driver.GormDb.Table(ap.TableName()).Where("id = ?", id).First(&adminPermit).Error
	if err != nil {
		log.Println(err.Error())
		return
	}
	return
}

// GetAdminPermitByAdminID 根據 AdminID 查詢 AdminPermits
func (ap *AdminPermitDAO) GetAdminPermitByAdminID(adminID string) (adminPermit AdminPermitDAO, err error) {
	err = driver.GormDb.Table(ap.TableName()).Where("admin_id = ?", adminID).First(&adminPermit).Error
	if err != nil {
		log.Println(err.Error())
		return
	}
	return
}

func (ap *AdminPermitDAO) GetAdminPermitListByAdminID(adminID string) (adminPermits []AdminPermitDAO, err error) {
	err = driver.GormDb.Table(ap.TableName()).Where("admin_id = ?", adminID).Find(&adminPermits).Error
	if err != nil {
		log.Println(err.Error())
		return
	}
	return
}
