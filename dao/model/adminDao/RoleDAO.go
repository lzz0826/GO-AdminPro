package adminDao

import (
	"AdminPro/common/driver"
	"AdminPro/common/model"
	"log"
	"time"
)

type RoleDAO struct {
	ID         string    `gorm:"column:id;type:varchar(40);primary_key;comment:'id'" json:"id"`
	RoleKey    string    `gorm:"column:role_key;type:varchar(40);not null;comment:'給spring定位用'" json:"role_key" binding:"required"`
	RoleName   string    `gorm:"column:role_name;type:varchar(40);not null;comment:'腳色名稱'" json:"role_name" binding:"required"`
	Sort       int       `gorm:"column:sort;type:int;comment:'排序'" json:"sort"`
	RoleStatus int       `gorm:"column:role_status;type:int;not null;default:0;comment:'腳色狀態'" json:"role_status"binding:"required" `
	Memo       string    `gorm:"column:memo;type:text;comment:'備註'" json:"memo"`
	CreatorID  string    `gorm:"column:creator_id;type:varchar(40);not null;comment:'創建者id(adminVo id)'" json:"creator_id"`
	UpdaterID  string    `gorm:"column:updater_id;type:varchar(40);not null;comment:'更新者id(adminVo id)'" json:"updater_id"`
	RoleDesc   string    `gorm:"column:role_desc;type:text;comment:'說明'" json:"role_desc"`
	UpdateTime time.Time `gorm:"column:update_time;type:timestamp;not null;comment:'更新時間'" json:"update_time"`
	CreateTime time.Time `gorm:"column:create_time;type:timestamp;not null;comment:'創建時間'" json:"create_time"`
}

func (r *RoleDAO) TableName() string {
	return "admin_role"
}

// InsertRole 插入 RoleDAO 資料
func (r *RoleDAO) InsertRole() (err error) {
	err = driver.GormDb.Table(r.TableName()).Omit("id").Create(r).Error
	if err != nil {
		log.Println("InsertRole error:", err.Error())
		return
	}
	return
}

func (r *RoleDAO) GetAllRoleList(pagination *model.Pagination) (roles []RoleDAO, err error) {
	// 分页查询
	offset := (pagination.Page - 1) * pagination.Size
	err = driver.GormDb.Table(r.TableName()).Limit(pagination.Size).Offset(offset).Order(pagination.Sort).Find(&roles).Error
	if err != nil {
		log.Println("GetRoleByIDs error:", err.Error())
		return
	}
	return
}

// GetRoleByID 根據 ID 查詢 RoleDAO
func (r *RoleDAO) GetRoleByID(id string) (role RoleDAO, err error) {
	err = driver.GormDb.Table(r.TableName()).Where("id = ?", id).First(&role).Error
	if err != nil {
		log.Println("GetRoleByID error:", err.Error())
		return
	}
	return
}

// GetRoleByRoleKey 根據 RoleKey 查詢 RoleDAO
func (r *RoleDAO) GetRoleByRoleKey(roleKey string) (role RoleDAO, err error) {
	err = driver.GormDb.Table(r.TableName()).Where("role_key = ?", roleKey).First(&role).Error
	if err != nil {
		log.Println("GetRoleByRoleKey error:", err.Error())
		return
	}
	return
}

func (r *RoleDAO) GetRoleByIDs(ids []string) (roles []RoleDAO, err error) {
	err = driver.GormDb.Table(r.TableName()).Where("id IN (?)", ids).Find(&roles).Error

	if err != nil {
		log.Println("GetRoleByIDs error:", err.Error())
		return
	}
	return
}
