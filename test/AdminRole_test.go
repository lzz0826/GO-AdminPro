package test

import (
	"AdminPro/dao/model/adminDao"
	"fmt"
	"testing"
	"time"
)

func TestAdminRoleInsert(t *testing.T) {
	newAdminRole := adminDao.AdminRoleDAO{
		ID:         "778",
		AdminID:    "996", // 你需要提供 AdminID 的值
		RoleID:     "001", // 你需要提供 RoleID 的值
		CreatorID:  "999", // 你需要提供 CreatorID 的值
		UpdaterID:  "999", // 你需要提供 UpdaterID 的值
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}
	err := newAdminRole.InsertAdminRole()
	if err != nil {
		t.Fatalf("插入 AdminRoleDAO 記錄失敗：%v", err)
	}
}

func TestGetAdminRoleByID(t *testing.T) {
	newAdminRole := adminDao.AdminRoleDAO{}
	retrievedAdminRole, err := newAdminRole.GetAdminRoleByID("777")
	if err != nil {
		t.Fatalf("根據 ID 查詢 AdminRoleDAO 記錄失敗：%v", err)
	}

	fmt.Println(retrievedAdminRole.ID)
	fmt.Println(retrievedAdminRole.AdminID)
	fmt.Println(retrievedAdminRole.RoleID)

	t.Logf("%+v\n", retrievedAdminRole)
}
