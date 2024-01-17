package admin

import (
	"AdminPro/dao/model/adminDao"
)

func GetById(id string) (user adminDao.User, err error) {
	userModel := adminDao.User{}

	user, err = userModel.GetById(id)

	return user, err
}
