package admin

import (
	"AdminPro/dao/model/adminDao"
	_ "AdminPro/dao/model/adminDao"
)

func GetAdminByUsername(username string) (adminDao.AdminDAO, error) {
	adminModel := adminDao.AdminDAO{}
	admin, err := adminModel.GetAdminByUsername(username)
	return admin, err
}

func GetAdminById(adminId string) (adminDao.AdminDAO, error) {
	adminModel := adminDao.AdminDAO{}
	admin, err := adminModel.GetAdminByID(adminId)
	return admin, err
}
