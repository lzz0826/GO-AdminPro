package service

import (
	"AdminPro/common/tool"
	"AdminPro/dao/model/adminDao"
	"AdminPro/dao/service/admin"
	"errors"
	"time"
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

// AddAdminPermits 為管理員添加權限
func AddAdminPermits(adminId string, permitsIds []string, currentAdminId string) error {

	_, err := admin.GetAdminById(adminId)
	if err != nil {
		return errors.New(tool.NotFindAdmin.Msg)
	}

	permits, err := admin.GetPermitByByIds(permitsIds)

	if err != nil || len(permitsIds) != len(permits) {
		return errors.New(tool.NotFinPermit.Msg)
	}

	//去掉原來已有的 權限
	originalPermits, err := GetPermitsByAdminId(adminId)
	if originalPermits != nil {
		remainingPermits := make([]adminDao.PermitDAO, 0)
		for _, p := range permits {
			if !ContainsPermits(originalPermits, p) {
				remainingPermits = append(remainingPermits, p)
			}
		}
		permits = remainingPermits
	}

	var adminPermits []adminDao.AdminPermitDAO

	for _, p := range permits {
		dao := adminDao.AdminPermitDAO{
			AdminID:    adminId,
			PermitID:   p.ID,
			CreatorID:  currentAdminId,
			UpdaterID:  currentAdminId,
			CreateTime: time.Now(),
			UpdateTime: time.Now(),
		}
		adminPermits = append(adminPermits, dao)
	}
	err = admin.InsertAdminPermits(adminPermits)
	if err != nil {
		return errors.New(tool.AddAdminPermitsFail.Msg)
	}
	return nil
}

// RemoveAdminPermits 移除管理員額外的權限
func RemoveAdminPermits(adminId string, permitsIds []string) error {

	_, err := admin.GetAdminById(adminId)
	if err != nil {
		return errors.New(tool.NotFindAdmin.Msg)
	}

	_, err = admin.CheckPermitIdsExist(permitsIds)

	if err != nil {
		return errors.New(tool.GetStatusMsgFromError(err))
	}

	permits, err := admin.GetAdminPermitByAdminIdAndPermitIds(adminId, permitsIds)

	var adminPermitDAOIds []string
	if permits != nil {
		for _, p := range permits {
			adminPermitDAOIds = append(adminPermitDAOIds, p.ID)
		}
	}
	err = admin.DeleteAdminPermit(adminPermitDAOIds)
	if err != nil {
		return errors.New(tool.RemoveAdminPermitsFail.Msg)
	}
	return nil

}

func ContainsPermits(daos []adminDao.PermitDAO, permit adminDao.PermitDAO) bool {
	for _, p := range daos {
		if p.ID == permit.ID {
			return true
		}
	}
	return false
}

func RemoveDuplicatesPermits(permits []adminDao.PermitDAO) []adminDao.PermitDAO {
	var adminDaoMap = make(map[string]adminDao.PermitDAO, 0)

	for _, p := range permits {
		adminDaoMap[p.ID] = p
	}
	var newPermits []adminDao.PermitDAO
	for _, m := range adminDaoMap {
		newPermits = append(newPermits, m)
	}
	return newPermits
}
