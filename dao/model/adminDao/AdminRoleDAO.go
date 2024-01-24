package adminDao

import (
	"AdminPro/common/driver"
	"log"
	"time"
)

// admin 對應 role 中間表
type AdminRoleDAO struct {
	ID         string    `gorm:"column:id;type:varchar(40);primary_key;comment:'id'" json:"id"`
	AdminID    string    `gorm:"column:admin_id;type:varchar(40);not null;comment:'對應的admin id'" json:"admin_id"`
	RoleID     string    `gorm:"column:role_id;type:varchar(40);not null;comment:'對應的腳色id'" json:"role_id"`
	CreatorID  string    `gorm:"column:creator_id;type:varchar(40);comment:'創建者id'" json:"creator_id"`
	UpdaterID  string    `gorm:"column:updater_id;type:varchar(40);comment:'更新者id'" json:"updater_id"`
	CreateTime time.Time `gorm:"column:create_time;not null;comment:'創建時間'" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time;not null;comment:'更新時間'" json:"update_time"`
}

func (ar *AdminRoleDAO) TableName() string {
	return "admin_admin_role"
}

// InsertAdminRole 插入 AdminRoleDAO 資料
func (ar *AdminRoleDAO) InsertAdminRole() (err error) {
	err = driver.GormDb.Table(ar.TableName()).Create(ar).Error
	if err != nil {
		log.Println("InsertAdminRole error:", err.Error())
		return
	}
	return
}

func (ar *AdminRoleDAO) InsertAdminRoles(adminRoles []AdminRoleDAO) (err error) {
	err = driver.GormDb.Table(adminRoles[0].TableName()).Omit("id").Create(&adminRoles).Error
	if err != nil {
		log.Println("InsertAdminRoles error:", err.Error())
		return
	}
	return
}

func (ar *AdminRoleDAO) GetAllAdminRole() (adminRole []AdminRoleDAO, err error) {
	err = driver.GormDb.Table(ar.TableName()).Find(&adminRole).Error
	if err != nil {
		log.Println("GetAdminRoleByID error:", err.Error())
		return
	}
	return
}

// GetAdminRoleByID 根據 ID 查詢 AdminRoleDAO
func (ar *AdminRoleDAO) GetAdminRoleByID(id string) (adminRole AdminRoleDAO, err error) {
	err = driver.GormDb.Table(ar.TableName()).Where("id = ?", id).First(&adminRole).Error
	if err != nil {
		log.Println("GetAdminRoleByID error:", err.Error())
		return
	}
	return
}

func (ar *AdminRoleDAO) GetAdminRoleByAdminId(adminId string) (adminRoles []AdminRoleDAO, err error) {
	err = driver.GormDb.Table(ar.TableName()).Where("admin_id IN (?)", adminId).Find(&adminRoles).Error
	if err != nil {
		log.Println("GetAdminRoleByID error:", err.Error())
		return
	}
	return
}

func (ar *AdminRoleDAO) GetAdminRoleByAdminIdAndRoleIds(admins string, roleIds []string) (adminRoles []AdminRoleDAO, err error) {
	err = driver.GormDb.Table(ar.TableName()).Where("admin_id = ? AND role_id IN (?)", admins, roleIds).Find(&adminRoles).Error
	if err != nil {
		log.Println("GetAdminRoleByAdminIdAndRoleIds error:", err.Error())
		return
	}
	return
}

func (ar *AdminRoleDAO) DeleteAdminRoleByIds(ids []string) (err error) {
	err = driver.GormDb.Table(ar.TableName()).Where("id IN (?)", ids).Delete(&AdminRoleDAO{}).Error
	if err != nil {
		log.Println("DeleteAdminRoleByIds error:", err.Error())
		return err
	}
	return nil
}
