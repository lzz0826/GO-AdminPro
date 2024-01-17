package test

import (
	"AdminPro/admin/service"
	_ "AdminPro/admin/service"
	"fmt"
	"testing"
	_ "testing"
	_ "time"
)

func TestGetPermitByAdminId(t *testing.T) {

	permits, err := service.GetPermitsByAdminId("0")
	if err != nil {
		t.Fatalf("GetPermitByAdminId失敗：%v", err)
	}
	for _, r := range permits {
		fmt.Printf("%+v\n", r)
	}

}

func TestGetPermitsByRoleId(t *testing.T) {
	roleIds := []string{"0", "1"}
	permits, err := service.GetPermitsByRoleId(roleIds)
	if err != nil {
		t.Fatalf("GetPermitsByRoleId：%v", err)
	}
	for _, r := range permits {
		fmt.Printf("%+v\n", r)
	}

}

func TestGetRoleByAdminId(t *testing.T) {

	roles, err := service.GetRoleByAdminId("0")
	if err != nil {
		t.Fatalf("GetRoleByAdminId失敗：%v", err)
	}
	for _, r := range roles {
		fmt.Printf("%+v\n", r)
	}

}

func TestGetAllPermitByAdminId(t *testing.T) {
	permits, err := service.GetAllPermitByAdminId("0")
	if err != nil {
		t.Fatalf("GetAllPermitByAdminId失敗：%v", err)
	}
	for _, r := range permits {
		fmt.Printf("%+v\n", r)
	}

}

func TestSetPermissionByAdminIdAndGetPermissionListByAdminId(t *testing.T) {
	service.SetPermissionByAdminId("0")
	ids := service.GetPermitKeyListByAdminId("0")
	for _, id := range ids {
		fmt.Printf("%+v\n", id)
	}
}

func TestCheckPermission(t *testing.T) {
	TestSetPermissionByAdminIdAndGetPermissionListByAdminId(t)
	permission := service.CheckPermission("0", "example_key2")
	fmt.Printf("Permission exists: %t\n", permission)

}
