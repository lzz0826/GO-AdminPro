package admin

import (
	"AdminPro/common/model"
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

func GetAllAdminList(pagination *model.Pagination) (admins []adminDao.AdminDAO, err error) {
	adminModel := adminDao.AdminDAO{}
	admins, err = adminModel.GetAllAdminList(pagination)
	return admins, err

}
