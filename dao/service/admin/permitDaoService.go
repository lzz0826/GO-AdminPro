package admin

import (
	"AdminPro/common/model"
	"AdminPro/common/tool"
	"AdminPro/dao/model/adminDao"
	"errors"
	"log"
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

func CheckPermitIdsExist(permitIds []string) (permits []adminDao.PermitDAO, err error) {
	permits, err = GetPermitByByIds(permitIds)
	if err != nil {
		return nil, errors.New(tool.NotFinPermit.Msg)
	}
	if permits == nil || len(permits) != len(permitIds) {
		return nil, errors.New(tool.NotFinPermit.Msg)
	}
	return permits, nil

}

func Convert() []adminDao.PermitDAO {

	page := model.Pagination{
		Page: 1,
		Size: 20,
	}
	permits, err := GetAllPermitList(&page)
	if err != nil {
		log.Println("GetAllPermitList errors:", err.Error())
	}

	return permits
}
