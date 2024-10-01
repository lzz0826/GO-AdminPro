package admin

import (
	"AdminPro/common/jwt"
	"AdminPro/common/tool"
	"AdminPro/common/utils"
	"AdminPro/dao/service/admin"
	"AdminPro/server/tonke"
	"AdminPro/vo/model/adminVo"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func CheckUserAndPassword(ctx *gin.Context, username string, password string) (vo adminVo.AdminLoginVO, globalErr *utils.GlobalError) {

	adm, err := admin.GetAdminByUsername(username)

	if err != nil {
		return adminVo.AdminLoginVO{}, utils.NewGlobalError(err, &tool.NotFindAdmin)
	}

	pass, err := admin.GetAdminTokenByAdminID(adm.ID)
	if err != nil {
		return adminVo.AdminLoginVO{}, utils.NewGlobalError(err, &tool.PasswordError)
	}

	err = bcrypt.CompareHashAndPassword([]byte(pass.Token), []byte(password))
	if err != nil {
		return adminVo.AdminLoginVO{}, utils.NewGlobalError(err, &tool.PasswordError)
	}

	tokenStr, err := jwt.LoginHandler(adm)
	if err != nil {
		return adminVo.AdminLoginVO{}, utils.NewGlobalError(err, &tool.PasswordError)
	}

	//記錄管理員權限 登出要刪
	SetPermissionByAdminId(adm.ID)
	//Redis Token 登出要刪
	tonke.SetTokenToRides(ctx, adm.ID, tokenStr)

	adminVo := adminVo.AdminLoginVO{
		Username:   adm.Username,
		AdminName:  adm.AdminName,
		Nickname:   adm.Nickname,
		Token:      tokenStr,
		UpdateTime: adm.CreateTime,
		CreateTime: adm.CreateTime,
	}

	return adminVo, nil
}
