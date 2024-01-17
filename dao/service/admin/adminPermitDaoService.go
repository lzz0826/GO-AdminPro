package admin

import (
	"AdminPro/dao/model/adminDao"
)

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
