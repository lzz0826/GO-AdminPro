package service

import (
	"AdminPro/common/mysql"
	"AdminPro/common/tool"
	"AdminPro/common/utils"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"strconv"
	"time"

	"AdminPro/dao/model/adminDao"
	"AdminPro/dao/service/admin"
	"errors"
	"fmt"
	_ "github.com/sony/sonyflake"
)

// AddAdmin 添加管理員
func AddAdmin(username string, password string, adminName string, nickName string, currentAdminId string) (err error) {
	return mysql.WithTransaction(func(tx *gorm.DB) error {
		_, err = admin.GetAdminByUsername(username)
		if err == nil {
			return errors.New(tool.AdminIsExits.Msg)
		}

		// 加密
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		encodePW := string(hashedPassword)
		fmt.Println(encodePW)

		// 在需要生成 ID 的地方调用 generateID 函数
		id, err := utils.GenerateID()
		if err != nil {
			return err
		}

		var dao = adminDao.AdminDAO{
			ID:         strconv.FormatInt(id, 10),
			Username:   username,
			AdminName:  adminName,
			Nickname:   nickName,
			CreatorID:  currentAdminId,
			UpdaterID:  currentAdminId,
			CreateTime: time.Now(),
			UpdateTime: time.Now(),
			LoginTime:  time.Now(),
		}

		err = admin.InsertAdmin(dao, tx)

		if err != nil {
			return errors.New(tool.RegisterAdminFail.Msg)
		}

		var adminTokenDao = adminDao.AdminTokenDAO{
			AdminID:    strconv.FormatInt(id, 10),
			Token:      encodePW,
			CreatorID:  currentAdminId,
			UpdaterID:  currentAdminId,
			CreateTime: time.Now(),
			UpdateTime: time.Now(),
			ExpireTime: time.Now(),
		}

		err = admin.InsertAdminToken(adminTokenDao, tx)
		if err != nil {
			return errors.New(tool.RegisterAdminFail.Msg)
		}

		return nil
	})
}
