package service

import (
	"AdminPro/dao/model/adminDao"
	"AdminPro/dao/service/admin"
	_ "fmt"
)

// key = adminId
var permissionMap = make(map[string]map[string]adminDao.PermitDAO)

// key = permitKey
func SetPermissionByAdminId(adminId string) {
	var permissionSet = make(map[string]adminDao.PermitDAO)
	permits, _ := GetAllPermitByAdminId(adminId)
	if permits != nil {
		for _, permit := range permits {
			permissionSet[permit.PermitKey] = permit
		}
	}
	//key = adminId
	permissionMap[adminId] = permissionSet
}

// key = adminId
func GetPermitKeyListByAdminId(adminId string) []string {
	var permitKeyList []string
	userPermissions, exists := permissionMap[adminId]

	if exists {
		for _, value := range userPermissions {
			permitKeyList = append(permitKeyList, value.PermitKey)
		}
	} else {
		return nil
	}
	return permitKeyList
}

func CheckPermission(adminId string, permitKey string) bool {
	permitKeyList := GetPermitKeyListByAdminId(adminId)

	for _, premission := range permitKeyList {
		if premission == permitKey {
			return true
		}
	}

	return false
}

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
	roledPermits, err := GetPermitsByRoleId(roleIds)
	permits = append(permits, roledPermits...)

	// 获取额外的权限
	adminPermits, err := GetPermitsByAdminId(adminId)
	permits = append(permits, adminPermits...)
	return
}

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

func GetPermitsByRoleId(roleIds []string) (permits []adminDao.PermitDAO, err error) {
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
