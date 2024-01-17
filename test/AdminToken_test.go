package test

import (
	"AdminPro/dao/model/adminDao"
	"fmt"
	"testing"
	"time"
)

func TestAdminTokenInsert(t *testing.T) {
	newAdminToken := adminDao.AdminTokenDAO{
		ID:         "566",
		AdminID:    "999",                         // 你需要提供 AdminID 的值
		TokenType:  1,                             // 你需要提供 TokenType 的值
		Token:      "example_token",               // 你需要提供 Token 的值
		ExpireTime: time.Now().Add(1 * time.Hour), // 你需要提供 ExpireTime 的值
		UpdateTime: time.Now(),
		CreateTime: time.Now(),
		CreatorID:  "999", // 你需要提供 CreatorID 的值
		UpdaterID:  "999", // 你需要提供 UpdaterID 的值
	}
	err := newAdminToken.InsertAdminToken()
	if err != nil {
		t.Fatalf("插入 AdminTokenDAO 記錄失敗：%v", err)
	}
}

func TestGetAdminTokenByID(t *testing.T) {
	newAdminToken := adminDao.AdminTokenDAO{}
	retrievedAdminToken, err := newAdminToken.GetAdminTokenByID("566")
	if err != nil {
		t.Fatalf("根據 ID 查詢 AdminTokenDAO 記錄失敗：%v", err)
	}

	fmt.Println(retrievedAdminToken.ID)
	fmt.Println(retrievedAdminToken.AdminID)
	fmt.Println(retrievedAdminToken.TokenType)
	fmt.Println(retrievedAdminToken.Token)
	fmt.Println(retrievedAdminToken.ExpireTime)

	t.Logf("%+v\n", retrievedAdminToken)
}

func TestGetAdminTokensByAdminID(t *testing.T) {
	newAdminToken := adminDao.AdminTokenDAO{}

	retrievedAdminTokens, err := newAdminToken.GetAdminTokensByAdminID("999")
	if err != nil {
		t.Fatalf("根據 AdminID 查詢 AdminTokens 記錄失敗：%v", err)
	}

	for _, adminToken := range retrievedAdminTokens {
		fmt.Printf("%+v\n", adminToken)
	}

	t.Logf("總共找到 %d 條 AdminTokens 記錄\n", len(retrievedAdminTokens))
}
