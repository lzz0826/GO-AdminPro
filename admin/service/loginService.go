package service

import (
	"AdminPro/common/jwt"
	"AdminPro/dao/service/admin"
	"AdminPro/vo/model/adminVo"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func CheckUserAndPassword(username string, password string) (vo adminVo.AdminLoginVO, err error) {

	usr, err := admin.GetAdminByUsername(username)

	if err != nil {
		return adminVo.AdminLoginVO{}, err
	}

	token, err := admin.GetAdminTokenByAdminID(usr.ID)
	if err != nil {
		return adminVo.AdminLoginVO{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(token.Token), []byte(password))
	if err != nil {
		fmt.Println("密码不匹配")
		return adminVo.AdminLoginVO{}, err
	}

	// jwt 暫時只放username
	tokenStr, err := jwt.LoginHandler(usr)
	if err != nil {
		return adminVo.AdminLoginVO{}, err
	}

	//TODO 記錄管理員權限 登出要刪
	SetPermissionByAdminId(usr.ID)

	adminVo := adminVo.AdminLoginVO{
		Username:   usr.Username,
		AdminName:  usr.AdminName,
		Nickname:   usr.Nickname,
		Token:      tokenStr,
		UpdateTime: usr.CreateTime,
		CreateTime: usr.CreateTime,
	}

	return adminVo, err
}
