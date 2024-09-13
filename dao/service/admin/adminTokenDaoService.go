package admin

import (
	"AdminPro/dao/model/adminDao"
	"gorm.io/gorm"
)

// 這裡的 Token 是加密後的密碼
func InsertAdminToken(dao adminDao.AdminTokenDAO, tx *gorm.DB) (err error) {
	err = dao.InsertAdminToken(tx)
	if err != nil {
		return err
	}
	return nil
}

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
