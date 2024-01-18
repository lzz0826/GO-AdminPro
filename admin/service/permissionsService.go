package service

import (
	"AdminPro/dao/model/adminDao"
	_ "fmt"
)

// key = adminId
var permissionMap = make(map[string]map[string]adminDao.PermitDAO)

// SetPermissionByAdminId key = permitKey
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
