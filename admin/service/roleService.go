package service

import (
	"AdminPro/common/tool"
	"AdminPro/dao/model/adminDao"
	"AdminPro/dao/service/admin"
	"errors"
	"time"
)

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
