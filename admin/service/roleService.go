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

	permits, err := admin.CheckPermitIdsExist(permitIds)
	if err != nil {
		return errors.New(tool.GetStatusMsgFromError(err))
	}

	//去掉原來已有的權限
	originalPermit, err := admin.GetRolePermitByRoleId(roleId)
	if originalPermit != nil {
		remainingPermits := make([]adminDao.PermitDAO, 0)

		for _, p := range permits {
			if !admin.ContainsRolePermitByPermitId(originalPermit, p.ID) {
				remainingPermits = append(remainingPermits, p)
			}
		}
		permits = remainingPermits
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

// RemoveRolePermits 移除角色的權限
func RemoveRolePermits(roleId string, permitIds []string) error {
	_, err := admin.GetRoleByID(roleId)
	if err != nil {
		return errors.New(tool.NotFindRole.Msg)
	}

	_, err = admin.CheckPermitIdsExist(permitIds)
	if err != nil {
		return errors.New(tool.NotFinPermit.Msg)
	}

	var rolePermitDAOIds []string
	permits, err := admin.GetRolePermitByRoleIdAndPermitIds(roleId, permitIds)
	if permits != nil {
		for _, p := range permits {
			rolePermitDAOIds = append(rolePermitDAOIds, p.ID)
		}
	}

	err = admin.DeleteRolePermitByIds(rolePermitDAOIds)
	if err != nil {
		return errors.New(tool.RemoveRolePermitsFail.Msg)
	}
	return nil

}

func RemoveAdminRoles(adminId string, roleIds []string) (err error) {
	if !admin.CheckAdminExist(adminId) {
		return errors.New(tool.NotFindAdmin.Msg)
	}

	_, err = admin.CheckRoleIdsExist(roleIds)
	if err != nil {
		return errors.New(tool.GetStatusMsgFromError(err))
	}

	adminRoles, err := admin.GetAdminRoleByAdminIdAndRoleIds(adminId, roleIds)

	var adminRolesDAOIds []string

	if adminRoles != nil {
		for _, a := range adminRoles {
			adminRolesDAOIds = append(adminRolesDAOIds, a.ID)
		}
	}

	err = admin.DeleteAdminRoleByIds(adminRolesDAOIds)
	if err != nil {
		return errors.New(tool.RemoveAdminRolesFail.Msg)
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
