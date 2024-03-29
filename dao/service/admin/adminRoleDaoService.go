package admin

import (
	"AdminPro/dao/model/adminDao"
)

//admin 對應 role 中間表

func GetAdminRoleByAdminId(admins string) (adminRoles []adminDao.AdminRoleDAO, err error) {
	dao := adminDao.AdminRoleDAO{}
	roles, err := dao.GetAdminRoleByAdminId(admins)
	return roles, err
}

func GetAdminRoleByAdminIdAndRoleIds(adminId string, roleIds []string) (adminRoles []adminDao.AdminRoleDAO, err error) {
	dao := adminDao.AdminRoleDAO{}
	roles, err := dao.GetAdminRoleByAdminIdAndRoleIds(adminId, roleIds)
	return roles, err
}

func GetAllAdminRole() (adminRoles []adminDao.AdminRoleDAO, err error) {
	dao := adminDao.AdminRoleDAO{}
	roles, err := dao.GetAllAdminRole()
	return roles, err
}

func InsertAdminRoles(adminRoles []adminDao.AdminRoleDAO) error {
	if len(adminRoles) == 0 {
		return nil
	}
	err := adminRoles[0].InsertAdminRoles(adminRoles)
	if err != nil {
		return err
	}
	return nil
}

func DeleteAdminRoleByIds(ids []string) error {
	if len(ids) == 0 {
		return nil
	}
	dao := adminDao.AdminRoleDAO{}
	err := dao.DeleteAdminRoleByIds(ids)
	if err != nil {
		return err
	}
	return nil
}
