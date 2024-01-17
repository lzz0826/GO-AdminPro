package test

import (
	"AdminPro/dao/model/adminDao"
	"fmt"
	"testing"
	"time"
)

func TestPermitInsert(t *testing.T) {
	newPermit := adminDao.PermitDAO{
		ID:         "556",
		PermitKey:  "example_key2",           // 你需要提供 PermitKey 的值
		PermitName: "Example PermitDAO",      // 你需要提供 PermitName 的值
		Memo:       "Some memo",              // 你需要提供 Memo 的值
		PermitDesc: "Description for permit", // 你需要提供 PermitDesc 的值
		Sort:       1,                        // 你需要提供 Sort 的值
		CreatorID:  "999",                    // 你需要提供 CreatorID 的值
		UpdaterID:  "999",                    // 你需要提供 UpdaterID 的值
		UpdateTime: time.Now(),
		CreateTime: time.Now(),
	}
	err := newPermit.InsertPermit()
	if err != nil {
		t.Fatalf("插入 PermitDAO 記錄失敗：%v", err)
	}
}

func TestGetPermitByID(t *testing.T) {
	newPermit := adminDao.PermitDAO{}
	retrievedPermit, err := newPermit.GetPermitByID("555")
	if err != nil {
		t.Fatalf("根據 ID 查詢 PermitDAO 記錄失敗：%v", err)
	}

	fmt.Println(retrievedPermit.ID)
	fmt.Println(retrievedPermit.PermitKey)
	fmt.Println(retrievedPermit.PermitName)

	t.Logf("%+v\n", retrievedPermit)
}

func TestGetPermitByPermitKey(t *testing.T) {
	newPermit := adminDao.PermitDAO{}
	retrievedPermit, err := newPermit.GetPermitByPermitKey("example_key")
	if err != nil {
		t.Fatalf("根據 PermitKey 查詢 PermitDAO 記錄失敗：%v", err)
	}

	fmt.Println(retrievedPermit.ID)
	fmt.Println(retrievedPermit.PermitName)
	fmt.Println(retrievedPermit.PermitDesc)

	t.Logf("%+v\n", retrievedPermit)
}

func Test(t *testing.T) {
	newPermit := adminDao.PermitDAO{}

	ids := []string{"1", "2"}

	permits, err := newPermit.GetPermitByByIds(ids)

	for _, permit := range permits {

		fmt.Printf("%+v\n", permit)
	}

	if err != nil {
		t.Fatalf("根據 ids 查詢 PermitDAO 記錄失敗：%v", err)
	}

	if len(permits) != len(ids) {
		t.Fatalf("返回的 permits 数量不匹配，期望 %d，实际 %d", len(ids), len(permits))
	}

}
