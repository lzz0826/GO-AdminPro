package test

import (
	"AdminPro/server/tonke"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http/httptest"
	"testing"
)

func TestSetTokenToRides(t *testing.T) {
	ctx := gin.CreateTestContextOnly(httptest.NewRecorder(), gin.Default())
	err := tonke.SetTokenToRides(ctx, "token1", "tony001")
	if err != nil {
		fmt.Println(err)
	}
}

func TestSetGetTokenToRides(t *testing.T) {
	ctx := gin.CreateTestContextOnly(httptest.NewRecorder(), gin.Default())
	rides, err := tonke.GetTokenToRides(ctx, "token1")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(rides)

}

func TestSetRemoveTokenToRides(t *testing.T) {
	ctx := gin.CreateTestContextOnly(httptest.NewRecorder(), gin.Default())
	err := tonke.RemoveTokenToRides(ctx, "token1")
	if err != nil {
		fmt.Println(err)
	}
}
