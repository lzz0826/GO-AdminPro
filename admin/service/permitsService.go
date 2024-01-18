package service

import (
	"AdminPro/dao/model/adminDao"
	"AdminPro/dao/service/admin"
)

// GetAllPermitByAdminId 查詢指定 adminId 所有的權限 (包含持有角色下)
func GetAllPermitByAdminId(adminId string) (permits []adminDao.PermitDAO, err error) {

	//依照角色 的權限
	var roleIds []string
	roles, err := GetRoleByAdminId(adminId)
	if err != nil {
		return
	}
	for _, role := range roles {
		roleIds = append(roleIds, role.ID)
	}

	// 获取角色的权限
	roledPermits, err := GetPermitsByRoleIds(roleIds)
	permits = append(permits, roledPermits...)

	// 获取额外的权限
	adminPermits, err := GetPermitsByAdminId(adminId)
	permits = append(permits, adminPermits...)
	return
}

// GetPermitsByAdminId 查詢指定 adminId 包含的 permit
func GetPermitsByAdminId(adminId string) (permits []adminDao.PermitDAO, err error) {
	var permitIdList []string

	adminPermitDAO, err := admin.GetAdminPermitListByAdminID(adminId)
	if err != nil {
		return
	}
	for _, permit := range adminPermitDAO {
		permitId := permit.PermitID
		permitIdList = append(permitIdList, permitId)
	}
	permitDAOList, err := admin.GetPermitByByIds(permitIdList)
	if err != nil {
		return
	}
	return permitDAOList, nil
}

// GetPermitsByRoleIds 查詢指定 角色們 包含的權限
func GetPermitsByRoleIds(roleIds []string) (permits []adminDao.PermitDAO, err error) {
	var permitIdList []string
	rolePermits, err := admin.GetRolePermitByRoleIds(roleIds)
	if err != nil {
		return
	}
	for _, rolePermit := range rolePermits {
		permitIdList = append(permitIdList, rolePermit.PermitID)
	}
	permitDAOList, err := admin.GetPermitByByIds(permitIdList)
	if err != nil {
		return
	}
	return permitDAOList, nil
}
