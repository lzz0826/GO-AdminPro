package test

import (
	"AdminPro/dao/model/adminDao"
	"fmt"
	"testing"
)

//func TestAdminInsert(t *testing.T) {
//	adminMember := adminDao.AdminDAO{
//		ID:            "993",
//		ChannelID:     "channel001",
//		Username:      "adminVo",
//		AdminName:     "最高", // 你需要提供 AdminName 的值
//		Nickname:      "大老", // 你需要提供 Nickname 的值
//		AccountStatus: 0,
//		LoginIP:       "192.168.1.1", // 你需要提供 LoginIP 的值
//		LoginTime:     time.Now(),
//		Memo:          "Some memo", // 你需要提供 Memo 的值
//		CreatorID:     "999",       // 你需要提供 CreatorID 的值
//		UpdaterID:     "999",       // 你需要提供 UpdaterID 的值
//		UpdateTime:    time.Now(),
//		CreateTime:    time.Now(),
//		TwoFactorCode: "123456", // 你需要提供 TwoFactorCode 的值
//	}
//	err := adminMember.InsertAdmin()
//	if err != nil {
//		t.Fatalf("插入 AdminDAO 記錄失敗：%v", err)
//	}
//
//}

func TestGetAdminByID(t *testing.T) {

	adminMember := adminDao.AdminDAO{}

	retrievedAdmin, err := adminMember.GetAdminByID("1")
	if err != nil {
		t.Fatalf("根據 ID 查詢 AdminDAO 記錄失敗：%v", err)
	}
	fmt.Println("retrievedAdmin.AdminName")
	fmt.Println(retrievedAdmin.AdminName)
	fmt.Println("retrievedAdmin.AdminName")

	t.Logf("%+v\n", retrievedAdmin)

}

func TestGetByUsername(t *testing.T) {

	adminMember := adminDao.AdminDAO{}

	admin, err := adminMember.GetAdminByUsername("adminVo")

	if err != nil {
		t.Fatalf("根據 useranme 查詢 AdminDAO 記錄失敗：%v", err)
	}
	fmt.Println(admin)

}
