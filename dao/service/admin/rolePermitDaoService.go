package admin

import (
	"AdminPro/dao/model/adminDao"
)

//role 對應 permit 中間表

func GetRolePermitByRoleId(roleIds string) (rolePermits []adminDao.RolePermitDAO, err error) {
	dao := adminDao.RolePermitDAO{}
	permits, err := dao.GetRolePermitByRoleId(roleIds)
	return permits, err
}

func GetRolePermitByRoleIds(roleIds []string) (rolePermits []adminDao.RolePermitDAO, err error) {
	dao := adminDao.RolePermitDAO{}
	permits, err := dao.GetRolePermitByRoleIds(roleIds)
	return permits, err
}

func GetRolePermitByRoleIdAndPermitIds(roleId string, permitIds []string) (rolePermits []adminDao.RolePermitDAO, err error) {
	dao := adminDao.RolePermitDAO{}
	permits, err := dao.GetRolePermitByRoleIdAndPermitIds(roleId, permitIds)
	return permits, err
}

func InsertRolePermit(dao adminDao.RolePermitDAO) error {
	err := dao.InsertRolePermit()
	if err != nil {
		return err
	}
	return nil

}

func InsertRolePermits(rolePermits []adminDao.RolePermitDAO) error {
	if len(rolePermits) == 0 {
		return nil
	}

	err := rolePermits[0].InsertRolePermits(rolePermits)
	if err != nil {
		return err
	}
	return nil
}

func DeleteByIds(ids []string) (err error) {
	if len(ids) == 0 {
		return nil
	}
	dao := adminDao.RolePermitDAO{}
	err = dao.DeleteByIds(ids)
	if err != nil {
		return err
	}
	return nil
}

// ContainsRolePermitByPermitId 依造permitId 檢查中間表是否重複
func ContainsRolePermitByPermitId(s []adminDao.RolePermitDAO, permitId string) bool {
	for _, r := range s {
		if r.PermitID == permitId {
			return true
		}
	}
	return false
}
