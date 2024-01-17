package test

import (
	"AdminPro/dao/model/adminDao"
	"fmt"
	"testing"
	"time"
)

func TestRolePermitInsert(t *testing.T) {
	newRolePermit := adminDao.RolePermitDAO{
		ID:         "334",
		RoleID:     "example_role_id",    // 你需要提供 RoleID 的值
		PermitID:   "example_permit_id2", // 你需要提供 PermitID 的值
		CreatorID:  "999",                // 你需要提供 CreatorID 的值
		UpdaterID:  "999",                // 你需要提供 UpdaterID 的值
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}
	err := newRolePermit.InsertRolePermit()
	if err != nil {
		t.Fatalf("插入 RolePermitDAO 記錄失敗：%v", err)
	}
}

func TestGetRolePermitByID(t *testing.T) {
	newRolePermit := adminDao.RolePermitDAO{}
	retrievedRolePermit, err := newRolePermit.GetRolePermitByID("333")
	if err != nil {
		t.Fatalf("根據 ID 查詢 RolePermitDAO 記錄失敗：%v", err)
	}

	fmt.Println(retrievedRolePermit.ID)
	fmt.Println(retrievedRolePermit.RoleID)
	fmt.Println(retrievedRolePermit.PermitID)

	t.Logf("%+v\n", retrievedRolePermit)
}

func TestGetRolePermitByRoleIDAndPermitID(t *testing.T) {
	newRolePermit := adminDao.RolePermitDAO{}

	retrievedRolePermit, err := newRolePermit.GetRolePermitByRoleIDAndPermitID("example_role_id", "example_permit_id")
	if err != nil {
		t.Fatalf("根據 RoleID 和 PermitID 查詢 RolePermitDAO 記錄失敗：%v", err)
	}

	fmt.Println(retrievedRolePermit.ID)
	fmt.Println(retrievedRolePermit.CreatorID)
	fmt.Println(retrievedRolePermit.UpdaterID)

	t.Logf("%+v\n", retrievedRolePermit)
}
