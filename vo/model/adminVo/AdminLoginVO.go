package adminVo

import (
	_ "AdminPro/dao/model/adminDao"
	"time"
)

type AdminLoginVO struct {
	Username   string
	AdminName  string
	Nickname   string
	Token      string
	UpdateTime time.Time
	CreateTime time.Time
}
