package adminDao

import (
	"AdminPro/common/driver"
	"AdminPro/common/model"
	"log"
	"time"
)

type PermitDAO struct {
	ID         string    `gorm:"column:id;type:varchar(40);primary_key;comment:'id'" json:"id"`
	PermitKey  string    `gorm:"column:permit_key;type:varchar(40);not null;comment:'給spring定位用'" json:"permit_key"`
	PermitName string    `gorm:"column:permit_name;type:varchar(40);not null;comment:'權限名稱'" json:"permit_name"`
	Memo       string    `gorm:"column:memo;type:text;comment:'備註'" json:"memo"`
	PermitDesc string    `gorm:"column:permit_desc;type:text;comment:'說明'" json:"permit_desc"`
	Sort       int       `gorm:"column:sort;type:int;comment:'排序'" json:"sort"`
	CreatorID  string    `gorm:"column:creator_id;type:varchar(40);not null;comment:'創建者id(adminVo id)'" json:"creator_id"`
	UpdaterID  string    `gorm:"column:updater_id;type:varchar(40);not null;comment:'更新者id(adminVo id)'" json:"updater_id"`
	UpdateTime time.Time `gorm:"column:update_time;type:timestamp;not null;comment:'更新時間'" json:"update_time"`
	CreateTime time.Time `gorm:"column:create_time;type:timestamp;not null;comment:'創建時間'" json:"create_time"`
}

func (p *PermitDAO) TableName() string {
	return "admin_permit"
}

// InsertPermit 插入 AdminPermitDAO 資料
func (p *PermitDAO) InsertPermit() (err error) {
	err = driver.GormDb.Table(p.TableName()).Create(p).Error
	if err != nil {
		log.Println("InsertPermit error:", err.Error())
		return
	}
	return
}

// GetPermitByID 根據 ID 查詢 AdminPermitDAO
func (p *PermitDAO) GetPermitByID(id string) (permit PermitDAO, err error) {
	err = driver.GormDb.Table(p.TableName()).Where("id = ?", id).First(&permit).Error
	if err != nil {
		log.Println("GetPermitByID error:", err.Error())
		return
	}
	return
}

// GetPermitByPermitKey 根據 PermitKey 查詢 PermitDAO
func (p *PermitDAO) GetPermitByPermitKey(permitKey string) (permit PermitDAO, err error) {
	err = driver.GormDb.Table(p.TableName()).Where("permit_key = ?", permitKey).First(&permit).Error
	if err != nil {
		log.Println("GetPermitByPermitKey error:", err.Error())
		return
	}
	return
}

func (p *PermitDAO) GetPermitByByIds(ids []string) (permits []PermitDAO, err error) {
	err = driver.GormDb.Table(p.TableName()).Where("id IN (?)", ids).Find(&permits).Error
	if err != nil {
		log.Println(err.Error())
		return
	}
	return
}

func (p *PermitDAO) GetAllPermitList(pagination *model.Pagination) (permits []PermitDAO, err error) {
	offset := (pagination.Page - 1) * pagination.Limit
	err = driver.GormDb.Table(p.TableName()).Limit(pagination.Limit).Offset(offset).Order(pagination.Sort).Find(&permits).Error
	if err != nil {
		log.Println(err.Error())
		return
	}
	return
}
