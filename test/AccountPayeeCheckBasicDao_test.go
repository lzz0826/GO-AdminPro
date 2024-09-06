package test

import (
	"AdminPro/common/enum"
	"AdminPro/common/model"
	"AdminPro/common/mysql"
	"AdminPro/dao/model/adminDao"
	"fmt"
	"gorm.io/gorm"
	"strconv"
	"testing"
	"time"
)

func TestListAccountPayeeChecks(t *testing.T) {

	userRandomId := "1"
	status := enum.WAIT
	results, err := adminDao.ListAccountPayeeChecks(&userRandomId, &status)
	if err != nil {
		t.Fatalf("ListAccountPayeeChecks 失敗：%v", err)
	}
	for _, result := range results {
		fmt.Println("----------------------------")
		fmt.Printf("%+v\n", result.ID)
		fmt.Printf("%+v\n", result.UID)
		fmt.Printf("%+v\n", result.Type)
		fmt.Printf("%+v\n", result.Description)
		fmt.Printf("%+v\n", result.Status)
		fmt.Printf("%+v\n", result.CheckID)
		fmt.Printf("%+v\n", result.CheckTime)
		fmt.Printf("%+v\n", result.UpdateTime)
		fmt.Printf("%+v\n", result.CreatedTime)
		fmt.Println("----------------------------")
	}
}

func TestSelectCheckPageTest(t *testing.T) {

	userRandomId := "1"
	status := enum.REJECT

	basicDao := adminDao.AccountPayeeCheckBasicDao{}
	//pagination := model.Pagination{Page: 2, Size: 2}
	//basicDao.Page(pagination)
	results, err := basicDao.SelectCheckPageTest(&userRandomId, &status)

	fmt.Println("--------------PageBean--------------")
	fmt.Printf("%+v\n", basicDao.PageBean.Total)
	fmt.Printf("%+v\n", basicDao.PageBean.Pages)
	fmt.Printf("%+v\n", basicDao.PageBean.IsLastPage)
	fmt.Printf("%+v\n", basicDao.PageBean.BeanList)
	fmt.Println("--------------PageBean--------------")
	if err != nil {
		t.Fatalf("TestSelectCheckPageTest 失敗：%v", err)
	}
	for _, result := range results {
		fmt.Println("----------------------------")
		fmt.Printf("%+v\n", result.ID)
		fmt.Printf("%+v\n", result.UID)
		fmt.Printf("%+v\n", result.Type)
		fmt.Printf("%+v\n", result.Description)
		fmt.Printf("%+v\n", *result.Status)
		fmt.Printf("%+v\n", result.CheckID)
		fmt.Printf("%+v\n", result.CheckTime)
		fmt.Printf("%+v\n", result.UpdateTime)
		fmt.Printf("%+v\n", result.CreatedTime)
		fmt.Println("----------------------------")
	}
}

func TestSumTotalStatusSUM(t *testing.T) {

	customizeSQL := func(db *gorm.DB) *gorm.DB {
		return db
	}

	price, err := adminDao.SumTotalStatusSUM(customizeSQL)

	if err != nil {
		t.Fatalf("SumTotalStatusSUM 失敗：%v", err)
	}

	fmt.Printf(" setus SUM : %+v\n", price)

}

func TestDeleteByPrimaryKey(t *testing.T) {
	primaryKey := 40
	i, err := adminDao.DeleteByPrimaryKey(primaryKey, &adminDao.AccountPayeeCheck{})
	if err != nil {
		t.Fatalf("TestDeleteByPrimaryKey 失敗：%v", err)
	}
	fmt.Printf("%+v\n", i)
}

func TestDeleteByList(t *testing.T) {
	columnName := "check_id"
	list := []int{51, 52}
	i, err := adminDao.DeleteByList(columnName, list, &adminDao.AccountPayeeCheck{})
	if err != nil {
		t.Fatalf("TestDeleteByList 失敗：%v", err)
	}
	fmt.Printf("%+v\n", i)
}

func TestCount(t *testing.T) {
	db := mysql.GormDb
	db = db.Where("uid = 4")
	i, err := adminDao.Count(db, &adminDao.AccountPayeeCheck{})
	if err != nil {
		t.Fatalf("TestCount 失敗：%v", err)
	}
	fmt.Printf("%+v\n", i)
}

func TestDelete(t *testing.T) {
	uid := 4
	db := mysql.GormDb
	db = db.Where("uid = ?", uid)
	i, err := adminDao.Delete(db, &adminDao.AccountPayeeCheck{})
	if err != nil {
		t.Fatalf("TestDelete 失敗：%v", err)
	}
	fmt.Printf("%+v\n", i)
}

func TestInsertReturnLastId(t *testing.T) {
	i := 4
	status := 0
	//s := "test"
	time := time.Now()
	ap := adminDao.AccountPayeeCheck{
		//ID:          &i,
		UID:  i,
		Type: i,
		//Description: nil,
		Status:      &status,
		CheckID:     i,
		CheckTime:   time,
		UpdateTime:  time,
		CreatedTime: time,
	}
	db := mysql.GormDb
	//db = db.Table("acwc")
	insert, err := adminDao.InsertReturnLastId(db, &ap)
	if err != nil {
		t.Fatalf("TestInsertReturnLastId 失敗：%v", err)
	}

	fmt.Println("----------------------------")
	fmt.Printf("%+v\n", insert)
	fmt.Println("----------------------------")
}

func TestInsertIgnoringNullList(t *testing.T) {

	var q []adminDao.AccountPayeeCheck
	i := 4
	s := "test"
	time := time.Now()
	ap := adminDao.AccountPayeeCheck{
		UID:         i,
		Type:        i,
		Description: s,
		Status:      &i,
		CheckID:     i,
		CheckTime:   time,
		UpdateTime:  time,
		CreatedTime: time,
	}
	i2 := 5
	s2 := "test2"
	ap2 := adminDao.AccountPayeeCheck{
		UID:         i2,
		Type:        i2,
		Description: s2,
		CheckTime:   time,
		UpdateTime:  time,
		CreatedTime: time,
	}
	q = append(q, ap, ap2)
	db := mysql.GormDb
	insert, err := adminDao.InsertsReturnLastIds(db, q)
	if err != nil {
		t.Fatalf("TestInsertIgnoringNullList 失敗：%v", err)
	}
	fmt.Println("----------------------------")
	fmt.Printf("%+v\n", insert)
	fmt.Println("----------------------------")
}

func TestUpdates(t *testing.T) {
	uid := 66
	typet := 6
	description := "test77"
	status := 6
	checkID := 6
	timeAdmin := time.Now()

	updatesReq := adminDao.AccountPayeeCheck{
		UID:         uid,
		Type:        typet,
		Description: description,
		Status:      &status,
		CheckID:     checkID,
		CheckTime:   timeAdmin,
		UpdateTime:  timeAdmin,
		CreatedTime: timeAdmin,
	}
	id := 4
	whereReq := map[string]interface{}{"id": id}

	specificTime := time.Date(2023, time.September, 15, 12, 0, 0, 0, time.UTC)
	db := mysql.GormDb
	db = db.Where("created_time >= ?", specificTime)
	db = db.Where("created_time <= ?", time.Now())
	rep, err := adminDao.Updates(db, &updatesReq, whereReq)
	if err != nil {
		t.Fatalf("TestUpdates 失敗：%v", err)
	}
	fmt.Printf(strconv.FormatInt(rep, 10))
}

func TestUpdateByPrimaryKey(t *testing.T) {
	uid := 55
	typet := 55
	status := 0
	//description := "test55"
	time := time.Now()
	updatesReq := adminDao.AccountPayeeCheck{
		UID:  uid,
		Type: typet,
		//Description: description,
		Status:      &status,
		CheckTime:   time,
		UpdateTime:  time,
		CreatedTime: time,
	}
	id := 489
	db := mysql.GormDb
	//db = db.Where("uid = 0")
	//db = db.Table("ewfwef")
	selective, err := adminDao.UpdateByPrimaryKey(db, id, updatesReq)
	if err != nil {
		t.Fatalf("TestUpdateByPrimaryKey 失敗：%v", err)
	}
	fmt.Printf(strconv.FormatInt(selective, 10))
}

func TestSetMAXType(t *testing.T) {
	adminDao.SetMAXType(64)
}

func TestSelectTypeLast(t *testing.T) {
	result, err := adminDao.SelectTypeLast(4)
	if err != nil {
		t.Fatalf("TestSelectTypeLast 失敗：%v", err)
	}
	fmt.Println("----------------------------")
	fmt.Printf("%+v\n", result.ID)
	fmt.Printf("%+v\n", result.UID)
	fmt.Printf("%+v\n", result.Type)
	fmt.Printf("%+v\n", result.Description)
	fmt.Printf("%+v\n", *result.Status)
	fmt.Printf("%+v\n", result.CheckID)
	fmt.Printf("%+v\n", result.CheckTime)
	fmt.Printf("%+v\n", result.UpdateTime)
	fmt.Printf("%+v\n", result.CreatedTime)
	fmt.Println("----------------------------")
}

func TestTestJoin(t *testing.T) {
	search := "es"
	res, err := adminDao.TestJoin(4, 1, 2, &search)
	if err != nil {
		t.Fatalf("TestTestJoin 失敗：%v", err)
	}
	for _, rueslt := range res {
		fmt.Println("----------------------------")
		fmt.Printf("%+v\n", rueslt.Description)
		fmt.Printf("%+v\n", rueslt.ID)
		fmt.Printf("%+v\n", rueslt.Username)
		fmt.Println("----------------------------")
	}
}

func TestJoinSelectCustomizeSqlCheckPage(t *testing.T) {

	search := "es"
	checkId := 4

	basicDao := adminDao.AccountPayeeCheckBasicDao{}
	pagination := model.Pagination{Page: 2, Size: 2}
	basicDao.Page(pagination)
	results, err := basicDao.JoinSelectCustomizeSqlCheckPage(checkId, &search)

	fmt.Println("--------------PageBean--------------")
	fmt.Printf("%+v\n", basicDao.PageBean.Total)
	fmt.Printf("%+v\n", basicDao.PageBean.Pages)
	fmt.Printf("%+v\n", basicDao.PageBean.IsLastPage)
	fmt.Printf("%+v\n", basicDao.PageBean.BeanList)
	fmt.Println("--------------PageBean--------------")
	if err != nil {
		t.Fatalf("TestJoinSelectCustomizeSqlCheckPage 失敗：%v", err)
	}
	for _, result := range results {
		fmt.Println("----------------------------")
		fmt.Printf("%+v\n", result.ID)
		fmt.Printf("%+v\n", result.Username)
		fmt.Printf("%+v\n", result.Description)

		fmt.Println("----------------------------")
	}
}

func TestTestSubquery(t *testing.T) {
	search := "xxx"
	_, err := adminDao.TestSubquery(search)
	if err != nil {
		t.Fatalf("TestTestSubquery 失敗：%v", err)
	}
}

func TestRawSubquery(t *testing.T) {
	search := 1
	_, err := adminDao.TestRawSubquery(search)
	if err != nil {
		t.Fatalf("TestRawSubquery 失敗：%v", err)
	}
}

func TestUpdateAccountStatusFoAdminAccountStatus(t *testing.T) {
	id := 1
	total, err := adminDao.UpdateAccountStatusFoAdminAccountStatus(id)
	fmt.Printf("%+v\n", total)
	if err != nil {
		t.Fatalf("TestUpdateAccountStatusFoAdminAccountStatus 失敗：%v", err)
	}
}

func TestUpdateByExampleSelectivePoint(t *testing.T) {
	uid := 1
	description := "test"
	total, err := adminDao.TestUpdateByExampleSelectivePoint(uid, description)
	if err != nil {
		t.Fatalf("TestUpdateByExampleSelectivePoint 失敗：%v", err)
	}
	fmt.Printf("%+v\n", total)
}
