package test

import (
	"AdminPro/admin/service"
	_ "AdminPro/admin/service"
	"fmt"
	_ "fmt"
	"testing"
	_ "testing"
	_ "time"
)

func TestCheckUserAndPassword(t *testing.T) {

	vo, err := service.CheckUserAndPassword("admin", "12345678")
	fmt.Println("vo")
	fmt.Println(vo.Token)

	if err != nil {
		fmt.Println(err)

	}
}
