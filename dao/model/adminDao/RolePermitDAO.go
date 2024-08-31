package adminDao

import (
	"AdminPro/common/mysql"
	"errors"
	"log"
	"time"
)

//role 對應 permit 中間表

type RolePermitDAO struct {
	ID         string    `gorm:"column:id;type:varchar(40);primary_key;comment:'id'" json:"id"`
	RoleID     string    `gorm:"column:role_id;type:varchar(40);not null;comment:'腳色id'" json:"role_id"`
	PermitID   string    `gorm:"column:permit_id;type:varchar(40);not null;comment:'權限id'" json:"permit_id"`
	CreatorID  string    `gorm:"column:creator_id;type:varchar(40);not null;comment:'創建者id'" json:"creator_id"`
	UpdaterID  string    `gorm:"column:updater_id;type:varchar(40);not null;comment:'更新者id'" json:"updater_id"`
	CreateTime time.Time `gorm:"column:create_time;type:timestamp;not null;comment:'創建時間'" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time;type:timestamp;not null;comment:'更新時間'" json:"update_time"`
}

func (rp *RolePermitDAO) TableName() string {
	return "admin_role_permit"
}

// InsertRolePermit 插入 RolePermitDAO 資料
func (rp *RolePermitDAO) InsertRolePermit() (err error) {
	err = mysql.GormDb.Table(rp.TableName()).Omit("id").Create(rp).Error
	if err != nil {
		log.Println("InsertRolePermit error:", err.Error())
		return
	}
	return nil
}

// InsertRolePermits 插入多筆 RolePermitDAO 資料
func (rp *RolePermitDAO) InsertRolePermits(rolePermits []RolePermitDAO) error {
	err := mysql.GormDb.Table(rolePermits[0].TableName()).Omit("id").Create(&rolePermits).Error
	if err != nil {
		log.Println("InsertRolePermits error:", err.Error())
		return err
	}
	return nil
}

// GetRolePermitByID 根據 ID 查詢 RolePermitDAO
func (rp *RolePermitDAO) GetRolePermitByID(id string) (rolePermit RolePermitDAO, err error) {
	err = mysql.GormDb.Table(rp.TableName()).Where("id = ?", id).First(&rolePermit).Error
	if err != nil {
		log.Println("GetRolePermitByID error:", err.Error())
		return
	}
	return rolePermit, nil
}

// GetRolePermitByRoleIDAndPermitID 根據 RoleID 和 PermitID 查詢 RolePermitDAO
func (rp *RolePermitDAO) GetRolePermitByRoleIDAndPermitID(roleID, permitID string) (rolePermit RolePermitDAO, err error) {
	err = mysql.GormDb.Table(rp.TableName()).
		Where("role_id = ? AND permit_id = ?", roleID, permitID).
		First(&rolePermit).Error
	if err != nil {
		log.Println("GetRolePermitByRoleIDAndPermitID error:", err.Error())
		return
	}
	return rolePermit, nil
}

func (rp *RolePermitDAO) GetRolePermitByRoleId(roleId string) (rolePermits []RolePermitDAO, err error) {
	err = mysql.GormDb.Table(rp.TableName()).Where("role_id = ?", roleId).Find(&rolePermits).Error
	if err != nil {
		log.Println("GetRolePermitByID error:", err.Error())
		return
	}
	return rolePermits, nil
}

func (rp *RolePermitDAO) GetRolePermitByRoleIds(roleIds []string) (rolePermits []RolePermitDAO, err error) {
	err = mysql.GormDb.Table(rp.TableName()).Where("role_id IN (?)", roleIds).Find(&rolePermits).Error
	if err != nil {
		log.Println("GetRolePermitByID error:", err.Error())
		return
	}
	return rolePermits, nil
}

func (rp *RolePermitDAO) GetRolePermitByRoleIdAndPermitIds(roleIds string, permitIds []string) (rolePermits []RolePermitDAO, err error) {
	err = mysql.GormDb.Table(rp.TableName()).Where("role_id = ? AND permit_id IN (?)", roleIds, permitIds).Find(&rolePermits).Error
	if err != nil {
		log.Println("GetRolePermitByRoleIdAndPermitIds error:", err.Error())
		return
	}
	return rolePermits, nil
}

func (rp *RolePermitDAO) DeleteRolePermitByIds(ids []string) (err error) {
	if len(ids) == 0 {
		return errors.New("ids slice is empty")
	}
	err = mysql.GormDb.Table(rp.TableName()).Where("id IN (?)", ids).Delete(&RolePermitDAO{}).Error
	if err != nil {
		log.Println("deleteIds error:", err.Error())
		return err
	}
	return nil
}
