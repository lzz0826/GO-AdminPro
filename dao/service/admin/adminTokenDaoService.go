package admin

import (
	"AdminPro/dao/model/adminDao"
)

func GetAdminTokenById(id string) (adminDao.AdminTokenDAO, error) {
	adminTokenModel := adminDao.AdminTokenDAO{}
	admin, err := adminTokenModel.GetAdminTokenByID(id)
	return admin, err
}

func GetAdminTokenByAdminID(adminId string) (adminDao.AdminTokenDAO, error) {
	adminTokenModel := adminDao.AdminTokenDAO{}
	admin, err := adminTokenModel.GetAdminTokenByAdminID(adminId)
	return admin, err
}
