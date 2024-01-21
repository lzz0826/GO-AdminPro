package service

import (
	"AdminPro/common/tool"
	"AdminPro/dao/model/adminDao"
	"AdminPro/dao/service/admin"
	"errors"
	"time"
)

// AddRolePermits 為角色添加權限
func AddRolePermits(roleId string, permitIds []string, currentAdminId string) error {

	role, err := admin.GetRoleByID(roleId)
	if err != nil {
		return errors.New(tool.NotFindRole.Msg)
	}

	//TODO 去掉原來已有的權限

	permits, err := admin.GetPermitByByIds(permitIds)
	if err != nil {
		return errors.New(tool.NotFinPermit.Msg)
	}
	if permits == nil || len(permits) != len(permitIds) {
		return errors.New(tool.NotFinPermit.Msg)
	}
	var rolePermits []adminDao.RolePermitDAO
	for _, permit := range permits {
		dao := adminDao.RolePermitDAO{
			RoleID:     role.ID,
			PermitID:   permit.ID,
			CreatorID:  currentAdminId,
			UpdaterID:  currentAdminId,
			CreateTime: time.Now(),
			UpdateTime: time.Now(),
		}
		rolePermits = append(rolePermits, dao)
	}
	err = admin.InsertRolePermits(rolePermits)
	if err != nil {
		return errors.New(tool.AddRolePermitsFail.Msg)
	}

	return nil
}

// AddAdminRoles 為管理員添加角色
func AddAdminRoles(adminId string, roleIds []string, currentAdminId string) error {

	_, err := admin.GetAdminById(adminId)
	if err != nil {
		return errors.New(tool.NotFindAdmin.Msg)
	}

	roles, err := admin.GetRoleByIDs(roleIds)
	if err != nil || len(roleIds) != len(roles) {
		return errors.New(tool.NotFindRole.Msg)
	}

	//去掉原來已有的 角色
	originalRoles, _ := GetRoleByAdminId(adminId)
	if originalRoles != nil {

		remainingRoles := make([]adminDao.RoleDAO, 0)
		for _, r := range roles {
			if !containsRole(originalRoles, r) {
				remainingRoles = append(remainingRoles, r)
			}
		}
		roles = remainingRoles
	}

	var adminRoles []adminDao.AdminRoleDAO

	for _, role := range roles {
		dao := adminDao.AdminRoleDAO{
			AdminID:    adminId,
			RoleID:     role.ID,
			CreatorID:  currentAdminId,
			UpdaterID:  currentAdminId,
			CreateTime: time.Now(),
			UpdateTime: time.Now(),
		}
		adminRoles = append(adminRoles, dao)
	}
	err = admin.InsertAdminRoles(adminRoles)

	if err != nil {
		return errors.New(tool.AddAdminRolesFail.Msg)
	}
	return nil
}

// GetRoleByAdminId 查詢指定 adminId 包含的角色
func GetRoleByAdminId(adminId string) (role []adminDao.RoleDAO, err error) {
	var roleIdList []string
	adminRoles, err := admin.GetAdminRoleByAdminId(adminId)
	if err != nil {
		return
	}
	for _, role := range adminRoles {
		roleID := role.RoleID
		roleIdList = append(roleIdList, roleID)
	}
	roles, err := admin.GetRoleByIDs(roleIdList)
	if err != nil {
		return
	}
	return roles, nil
}

func containsRole(s []adminDao.RoleDAO, role adminDao.RoleDAO) bool {
	for _, r := range s {
		if r.ID == role.ID {
			return true
		}
	}
	return false
}
