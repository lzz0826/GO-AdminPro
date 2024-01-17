package test

import (
	"AdminPro/dao/model/adminDao"
	"fmt"
	"testing"
	"time"
)

func TestAdminPermitInsert(t *testing.T) {
	adminPermitMode := adminDao.AdminPermitDAO{
		ID:         "887",
		AdminID:    "999", // 你需要提供 AdminID 的值
		PermitID:   "002", // 你需要提供 PermitID 的值
		CreatorID:  "999", // 你需要提供 CreatorID 的值
		UpdaterID:  "999", // 你需要提供 UpdaterID 的值
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}

	err := adminPermitMode.InsertAdminPermit()
	if err != nil {
		t.Fatalf("插入 AdminPermitDAO 記錄失敗：%v", err)
	}
}

func TestGetAdminPermitByID(t *testing.T) {
	adminPermitMode := adminDao.AdminPermitDAO{}
	retrievedAdminPermit, err := adminPermitMode.GetAdminPermitByID("888")
	if err != nil {
		t.Fatalf("根據 ID 查詢 AdminPermitDAO 記錄失敗：%v", err)
	}

	fmt.Println(retrievedAdminPermit.ID)
	fmt.Println(retrievedAdminPermit.AdminID)
	fmt.Println(retrievedAdminPermit.PermitID)

	t.Logf("%+v\n", retrievedAdminPermit)
}

func TestGetAdminPermitsByAdminID(t *testing.T) {
	adminPermitMode := adminDao.AdminPermitDAO{}

	retrievedAdminPermits, err := adminPermitMode.GetAdminPermitByAdminID("0")
	if err != nil {
		t.Fatalf("根據 ID 查詢 AdminPermitDAO 記錄失敗：%v", err)
	}

	fmt.Println(retrievedAdminPermits)

}

func TestGetAdminPermitListByAdminID(t *testing.T) {
	adminPermitMode := adminDao.AdminPermitDAO{}
	retrievedAdminPermits, err := adminPermitMode.GetAdminPermitListByAdminID("0")

	if err != nil {
		t.Fatalf("根據 ID 查詢 AdminPermitDAOList 記錄失敗：%v", err)
	}

	for _, r := range retrievedAdminPermits {

		fmt.Printf("%+v\n", r)
	}

}
