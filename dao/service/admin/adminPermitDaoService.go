package admin

import (
	"AdminPro/dao/model/adminDao"
)

// admin 對應 permit 中間表
func GetAdminPermitsByAdminID(adminId string) (adminPermit adminDao.AdminPermitDAO, err error) {
	dao := adminDao.AdminPermitDAO{}
	permit, err := dao.GetAdminPermitByAdminID(adminId)
	return permit, err
}

func GetAdminPermitListByAdminID(adminId string) (adminPermits []adminDao.AdminPermitDAO, err error) {
	dao := adminDao.AdminPermitDAO{}
	permits, err := dao.GetAdminPermitListByAdminID(adminId)
	return permits, err
}
