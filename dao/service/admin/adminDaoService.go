package admin

import (
	"AdminPro/common/model"
	"AdminPro/dao/model/adminDao"
	"gorm.io/gorm"
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

func InsertAdmin(dao adminDao.AdminDAO, tx *gorm.DB) (err error) {
	err = dao.InsertAdmin(tx)
	if err != nil {
		return err
	}
	return nil
}

func CheckAdminExist(adminId string) bool {
	_, err := GetAdminById(adminId)
	return err == nil
}
