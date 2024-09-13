package service

import (
	"AdminPro/dao/model/adminDao"
	_ "fmt"
)

// key = adminId , key = permitKey
var permissionMap = make(map[string]map[string]adminDao.PermitDAO)

// SetPermissionByAdminId
func SetPermissionByAdminId(adminId string) {
	var permissionSet = make(map[string]adminDao.PermitDAO)
	permits, _ := GetAllPermitByAdminId(adminId)
	if permits != nil {
		for _, permit := range permits {
			//key = permitKey
			permissionSet[permit.PermitKey] = permit
		}
	}
	//key = adminId
	permissionMap[adminId] = permissionSet
}

// RemovePermissionByAdminId key = adminId
func RemovePermissionByAdminId(adminId string) {
	delete(permissionMap, adminId)
}

// GetPermitKeyListByAdminId key = adminId
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
