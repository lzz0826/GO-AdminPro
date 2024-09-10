package test

import (
	"AdminPro/dao/model/adminDao"
	"AdminPro/dao/service/admin"
	"testing"
	"time"
)

func TestInsertRolePermits(t *testing.T) {

	// 創建兩個 RolePermitDAO 實例作為測試數據
	rolePermit1 := adminDao.RolePermitDAO{
		RoleID:     "role1",
		PermitID:   "permit1",
		CreatorID:  "creator1",
		UpdaterID:  "updater1",
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}

	rolePermit2 := adminDao.RolePermitDAO{
		RoleID:     "role2",
		PermitID:   "permit2",
		CreatorID:  "creator2",
		UpdaterID:  "updater2",
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}

	// 呼叫被測試函數
	err := admin.InsertRolePermits([]adminDao.RolePermitDAO{rolePermit1, rolePermit2})

	// 驗證測試結果
	if err != nil {
		t.Errorf("Expected no errors, got %v", err)
	}

}
