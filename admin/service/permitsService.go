package service

import (
	"AdminPro/dao/model/adminDao"
	"AdminPro/dao/service/admin"
)

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
