package admin

import (
	"AdminPro/common/model"
	"AdminPro/dao/model/adminDao"
)

func GetPermitByByIds(ids []string) (permits []adminDao.PermitDAO, err error) {
	permitModel := adminDao.PermitDAO{}
	permits, err = permitModel.GetPermitByByIds(ids)
	return permits, err
}

func GetAllPermitList(pagination *model.Pagination) (permits []adminDao.PermitDAO, err error) {
	permitModel := adminDao.PermitDAO{}
	permits, err = permitModel.GetAllPermitList(pagination)
	return permits, err

}
