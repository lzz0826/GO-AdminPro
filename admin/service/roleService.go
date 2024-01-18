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
		return err
	}

	return nil
}

//TODO
//func AddAdminRoles(adminId string, roleIds []string, currentAdminId string) error {
//
//	err := admin.InsertAdminRoles()
//}

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
