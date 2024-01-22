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

func GetAdminPermitByAdminIdAndPermitIds(adminId string, permitsId []string) (adminPermits []adminDao.AdminPermitDAO, err error) {
	dao := adminDao.AdminPermitDAO{}
	permits, err := dao.GetAdminPermitByAdminIdAndPermitIds(adminId, permitsId)
	return permits, err
}

func InsertAdminPermits(adminPermits []adminDao.AdminPermitDAO) error {
	if len(adminPermits) == 0 {
		return nil
	}
	err := adminPermits[0].InsertAdminPermits(adminPermits)
	if err != nil {
		return err
	}
	return nil
}

func DeleteAdminPermit(ids []string) error {
	if len(ids) == 0 {
		return nil
	}
	dao := adminDao.AdminPermitDAO{}

	err := dao.DeleteAdminPermit(ids)
	if err != nil {
		return err
	}
	return nil

}
