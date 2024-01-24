package admin

import (
	"AdminPro/common/model"
	"AdminPro/common/tool"
	"AdminPro/dao/model/adminDao"
	"errors"
)

func InsertRole(dao adminDao.RoleDAO) error {

	err := dao.InsertRole()
	if err != nil {
		return err
	}
	return nil
}

func GetAllRoleList(pagination *model.Pagination) (role []adminDao.RoleDAO, err error) {

	roleDAO := adminDao.RoleDAO{}
	roles, err := roleDAO.GetAllRoleList(pagination)
	return roles, err
}

func GetRoleByID(id string) (role adminDao.RoleDAO, err error) {
	roleDAO := adminDao.RoleDAO{}
	role, err = roleDAO.GetRoleByID(id)
	return role, err

}

func GetRoleByIDs(ids []string) (role []adminDao.RoleDAO, err error) {
	roleDAO := adminDao.RoleDAO{}
	roles, err := roleDAO.GetRoleByIDs(ids)
	return roles, err
}

func CheckRoleIdsExist(roleIds []string) (roles []adminDao.RoleDAO, err error) {
	roles, err = GetRoleByIDs(roleIds)
	if err != nil {
		return nil, errors.New(tool.NotFindRole.Msg)
	}
	if roles == nil || len(roles) != len(roleIds) {
		return nil, errors.New(tool.NotFindRole.Msg)
	}
	return roles, nil

}
