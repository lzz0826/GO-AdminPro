package test

import (
	"AdminPro/server/admin"
	"fmt"
	"testing"
	_ "testing"
	_ "time"
)

func TestGetPermitByAdminId(t *testing.T) {

	permits, err := admin.GetPermitsByAdminId("0")
	if err != nil {
		t.Fatalf("GetPermitByAdminId失敗：%v", err)
	}
	for _, r := range permits {
		fmt.Printf("%+v\n", r)
	}

}

func TestGetPermitsByRoleId(t *testing.T) {
	roleIds := []string{"0", "1"}
	permits, err := admin.GetPermitsByRoleIds(roleIds)
	if err != nil {
		t.Fatalf("GetPermitsByRoleIds：%v", err)
	}
	for _, r := range permits {
		fmt.Printf("%+v\n", r)
	}

}

func TestGetRoleByAdminId(t *testing.T) {

	roles, err := admin.GetRoleByAdminId("0")
	if err != nil {
		t.Fatalf("GetRoleByAdminId失敗：%v", err)
	}
	for _, r := range roles {
		fmt.Printf("%+v\n", r)
	}

}

func TestGetAllPermitByAdminId(t *testing.T) {
	permits, err := admin.GetAllPermitByAdminId("0")
	if err != nil {
		t.Fatalf("GetAllPermitByAdminId失敗：%v", err)
	}
	for _, r := range permits {
		fmt.Printf("%+v\n", r)
	}

}

func TestSetPermissionByAdminIdAndGetPermissionListByAdminId(t *testing.T) {
	admin.SetPermissionByAdminId("0")
	ids := admin.GetPermitKeyListByAdminId("0")
	for _, id := range ids {
		fmt.Printf("%+v\n", id)
	}
}

func TestCheckPermission(t *testing.T) {
	TestSetPermissionByAdminIdAndGetPermissionListByAdminId(t)
	permission := admin.CheckPermission("0", "example_key2")
	fmt.Printf("Permission exists: %t\n", permission)

}
