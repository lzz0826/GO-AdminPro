package service

import (
	"AdminPro/common/jwt"
	"AdminPro/common/tool"
	"AdminPro/common/utils"
	"AdminPro/dao/service/admin"
	"AdminPro/vo/model/adminVo"
	"golang.org/x/crypto/bcrypt"
)

func CheckUserAndPassword(username string, password string) (vo adminVo.AdminLoginVO, globalErr *utils.GlobalError) {

	usr, err := admin.GetAdminByUsername(username)

	if err != nil {
		return adminVo.AdminLoginVO{}, utils.NewGlobalError(err, &tool.NotFindAdmin)
	}

	pass, err := admin.GetAdminTokenByAdminID(usr.ID)
	if err != nil {
		return adminVo.AdminLoginVO{}, utils.NewGlobalError(err, &tool.PasswordError)
	}

	err = bcrypt.CompareHashAndPassword([]byte(pass.Token), []byte(password))
	if err != nil {
		return adminVo.AdminLoginVO{}, utils.NewGlobalError(err, &tool.PasswordError)
	}

	tokenStr, err := jwt.LoginHandler(usr)
	if err != nil {
		return adminVo.AdminLoginVO{}, utils.NewGlobalError(err, &tool.PasswordError)
	}

	//記錄管理員權限 登出要刪
	SetPermissionByAdminId(usr.ID)

	adminVo := adminVo.AdminLoginVO{
		Username:   usr.Username,
		AdminName:  usr.AdminName,
		Nickname:   usr.Nickname,
		Token:      tokenStr,
		UpdateTime: usr.CreateTime,
		CreateTime: usr.CreateTime,
	}

	return adminVo, nil
}
