package test

import (
	"AdminPro/server/admin"
	"fmt"
	_ "fmt"
	"testing"
	_ "testing"
	_ "time"
)

func TestCheckUserAndPassword(t *testing.T) {

	vo, err := admin.CheckUserAndPassword("admin", "12345678")
	fmt.Println("vo")
	fmt.Println(vo.Token)

	if err != nil {
		fmt.Println(err)

	}
}
