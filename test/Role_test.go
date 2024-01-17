package test

import (
	"AdminPro/dao/model/adminDao"
	"fmt"
	"testing"
	"time"
)

func TestRoleInsert(t *testing.T) {

	newRole := adminDao.RoleDAO{
		ID:         "775",
		RoleKey:    "example_role_key2",    // 你需要提供 RoleKey 的值
		RoleName:   "Example RoleDAO",      // 你需要提供 RoleName 的值
		Sort:       1,                      // 你需要提供 Sort 的值
		RoleStatus: 0,                      // 你需要提供 RoleStatus 的值
		Memo:       "Some memo",            // 你需要提供 Memo 的值
		CreatorID:  "999",                  // 你需要提供 CreatorID 的值
		UpdaterID:  "999",                  // 你需要提供 UpdaterID 的值
		RoleDesc:   "Description for role", // 你需要提供 RoleDesc 的值
		UpdateTime: time.Now(),
		CreateTime: time.Now(),
	}
	err := newRole.InsertRole()
	if err != nil {
		t.Fatalf("插入 RoleDAO 記錄失敗：%v", err)
	}
}

func TestGetRoleByID(t *testing.T) {
	newRole := adminDao.RoleDAO{}
	retrievedRole, err := newRole.GetRoleByID("777")
	if err != nil {
		t.Fatalf("根據 ID 查詢 RoleDAO 記錄失敗：%v", err)
	}

	fmt.Println(retrievedRole.ID)
	fmt.Println(retrievedRole.RoleKey)
	fmt.Println(retrievedRole.RoleName)

	t.Logf("%+v\n", retrievedRole)
}

func TestGetRoleByRoleKey(t *testing.T) {
	newRole := adminDao.RoleDAO{}

	retrievedRole, err := newRole.GetRoleByRoleKey("example_role_key")
	if err != nil {
		t.Fatalf("根據 RoleKey 查詢 RoleDAO 記錄失敗：%v", err)
	}

	fmt.Println(retrievedRole.ID)
	fmt.Println(retrievedRole.RoleName)
	fmt.Println(retrievedRole.RoleDesc)

	t.Logf("%+v\n", retrievedRole)
}
