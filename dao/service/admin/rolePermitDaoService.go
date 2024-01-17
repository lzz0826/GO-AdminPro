package admin

import (
	"AdminPro/dao/model/adminDao"
)

func GetRolePermitByRoleIds(roleIds []string) (rolePermits []adminDao.RolePermitDAO, err error) {
	dao := adminDao.RolePermitDAO{}
	permits, err := dao.GetRolePermitByRoleIds(roleIds)
	return permits, err
}

func InsertRolePermit(dao adminDao.RolePermitDAO) error {
	err := dao.InsertRolePermit()
	if err != nil {
		return err
	}
	return nil

}

func InsertRolePermits(roles []adminDao.RolePermitDAO) error {
	err := roles[0].InsertRolePermits(roles)
	if err != nil {
		return err
	}
	return nil
}
