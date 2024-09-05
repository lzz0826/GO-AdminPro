package test

import (
	"AdminPro/common/enum"
	"AdminPro/common/model"
	"AdminPro/dao/model/adminDao"
	"fmt"
	"gorm.io/gorm"
	"strconv"
	"testing"
	"time"
)

func TestSelectByObjWhereReq(t *testing.T) {
	var results []adminDao.AccountPayeeCheck
	customizeSQL := func(db *gorm.DB) *gorm.DB {
		db = db.Order("id desc")
		return db
	}
	uid := 4
	whereReq := adminDao.AccountPayeeCheck{
		UID: &uid,
	}
	err := adminDao.SelectByObjWhereReq(customizeSQL, &whereReq, &results, adminDao.AccountPayeeCheck{})

	if err != nil {
		t.Fatalf("TestSelectByObjWhereReq 失敗：%v", err)
	}

	for _, result := range results {
		fmt.Println("----------------------------")
		fmt.Printf("%+v\n", *result.ID)
		fmt.Printf("%+v\n", *result.UID)
		fmt.Printf("%+v\n", *result.Type)
		fmt.Printf("%+v\n", *result.Description)
		fmt.Printf("%+v\n", *result.Status)
		fmt.Printf("%+v\n", *result.CheckID)
		fmt.Printf("%+v\n", *result.CheckTime)
		fmt.Printf("%+v\n", *result.UpdateTime)
		fmt.Printf("%+v\n", *result.CreatedTime)
		fmt.Println("----------------------------")
	}
}

func TestSelectByObjWhereReqPage(t *testing.T) {
	var results []adminDao.AccountPayeeCheck
	customizeSQL := func(db *gorm.DB) *gorm.DB {
		db = db.Select("description")
		db = db.Order("id asc")
		return db
	}
	i := 4
	whereReq := adminDao.AccountPayeeCheck{
		//UID:    &i,
		Status: &i,
	}
	pagination := model.Pagination{Page: 1, Size: 2}

	total, err := adminDao.SelectByObjWhereReqPage(customizeSQL, &whereReq, &results, &pagination)

	if err != nil {
		t.Fatalf("TestSelectByObjWhereReqPage 失敗：%v", err)
	}

	fmt.Printf("%+v\n total: ", total)

	for _, result := range results {
		fmt.Println("----------------------------")
		//fmt.Printf("%+v\n", *result.ID)
		//fmt.Printf("%+v\n", *result.UID)
		//fmt.Printf("%+v\n", *result.Type)
		fmt.Printf("%+v\n", *result.Description)
		//fmt.Printf("%+v\n", *result.Status)
		//fmt.Printf("%+v\n", *result.CheckID)
		//fmt.Printf("%+v\n", *result.CheckTime)
		//fmt.Printf("%+v\n", *result.UpdateTime)
		//fmt.Printf("%+v\n", *result.CreatedTime)
		fmt.Println("----------------------------")
	}
}

func TestListAccountPayeeChecks(t *testing.T) {

	userRandomId := "1"
	status := enum.WAIT

	results, err := adminDao.ListAccountPayeeChecks(&userRandomId, &status)

	if err != nil {
		t.Fatalf("ListAccountPayeeChecks 失敗：%v", err)
	}

	for _, result := range results {
		fmt.Println("----------------------------")
		fmt.Printf("%+v\n", *result.ID)
		fmt.Printf("%+v\n", *result.UID)
		fmt.Printf("%+v\n", *result.Type)
		fmt.Printf("%+v\n", *result.Description)
		fmt.Printf("%+v\n", *result.Status)
		fmt.Printf("%+v\n", *result.CheckID)
		fmt.Printf("%+v\n", *result.CheckTime)
		fmt.Printf("%+v\n", *result.UpdateTime)
		fmt.Printf("%+v\n", *result.CreatedTime)
		fmt.Println("----------------------------")
	}
}

func TestListAccountPayeeChecksPage(t *testing.T) {

	userRandomId := "1"
	status := enum.WAIT

	pagination := model.Pagination{Page: 2, Size: 4}

	total, results, err := adminDao.ListAccountPayeeChecksPage(&userRandomId, &status, &pagination)

	if err != nil {
		t.Fatalf("TestListAccountPayeeChecksPage 失敗：%v", err)
	}

	fmt.Printf(" total : %+v\n", total)

	for _, result := range results {
		fmt.Println("----------------------------")
		//fmt.Printf("%+v\n", *result.ID)
		//fmt.Printf("%+v\n", *result.UID)
		//fmt.Printf("%+v\n", *result.Type)
		fmt.Printf("%+v\n", *result.Description)
		//fmt.Printf("%+v\n", *result.Status)
		//fmt.Printf("%+v\n", *result.CheckID)
		//fmt.Printf("%+v\n", *result.CheckTime)
		//fmt.Printf("%+v\n", *result.UpdateTime)
		//fmt.Printf("%+v\n", *result.CreatedTime)
		fmt.Println("----------------------------")
	}
}

func TestSelectCustomizeSqlCheckPageTest(t *testing.T) {

	userRandomId := "1"
	status := enum.WAIT

	basicDao := adminDao.AccountPayeeCheckBasicDao{}
	pagination := model.Pagination{Page: 4, Size: 2}
	basicDao.Page(pagination)
	results, err := basicDao.SelectCustomizeSqlCheckPageTest(&userRandomId, &status)

	fmt.Println("--------------PageBean--------------")
	fmt.Printf("%+v\n", basicDao.PageBean.Total)
	fmt.Printf("%+v\n", basicDao.PageBean.Pages)
	fmt.Printf("%+v\n", basicDao.PageBean.IsLastPage)
	fmt.Printf("%+v\n", basicDao.PageBean.BeanList)
	fmt.Println("--------------PageBean--------------")
	if err != nil {
		t.Fatalf("TestSelectCustomizeSqlCheckPageTest 失敗：%v", err)
	}
	for _, result := range results {
		fmt.Println("----------------------------")
		fmt.Printf("%+v\n", *result.ID)
		fmt.Printf("%+v\n", *result.UID)
		fmt.Printf("%+v\n", *result.Type)
		fmt.Printf("%+v\n", *result.Description)
		fmt.Printf("%+v\n", *result.Status)
		fmt.Printf("%+v\n", *result.CheckID)
		fmt.Printf("%+v\n", *result.CheckTime)
		fmt.Printf("%+v\n", *result.UpdateTime)
		fmt.Printf("%+v\n", *result.CreatedTime)
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
		fmt.Printf("%+v\n", *result.ID)
		fmt.Printf("%+v\n", *result.Username)
		fmt.Printf("%+v\n", *result.Description)

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

	fmt.Printf(" price : %+v\n", price)

}

func TestSelectByPrimaryKey2(t *testing.T) {
	var result adminDao.AccountPayeeCheck

	primaryKey := 4

	err := adminDao.SelectByPrimaryKey(primaryKey, &result)

	if err != nil {
		t.Fatalf("TestSelectByPrimaryKey2 失敗：%v", err)
	}
	fmt.Println("----------------------------")
	fmt.Printf("%+v\n", *result.ID)
	fmt.Printf("%+v\n", *result.UID)
	fmt.Printf("%+v\n", *result.Type)
	fmt.Printf("%+v\n", *result.Description)
	fmt.Printf("%+v\n", *result.Status)
	fmt.Printf("%+v\n", *result.CheckID)
	fmt.Printf("%+v\n", *result.CheckTime)
	fmt.Printf("%+v\n", *result.UpdateTime)
	fmt.Printf("%+v\n", *result.CreatedTime)
	fmt.Println("----------------------------")
}

func TestDeleteByPrimaryKey2(t *testing.T) {
	primaryKey := 40
	i, err := adminDao.DeleteByPrimaryKey(primaryKey, &adminDao.AccountPayeeCheck{})
	if err != nil {
		t.Fatalf("TestDeleteByPrimaryKey2 失敗：%v", err)
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

func TestDeleteByExample2(t *testing.T) {
	uir := 4
	customizeSQL := func(db *gorm.DB) *gorm.DB {
		db = db.Where("uid = ?", uir)
		return db
	}
	i, err := adminDao.DeleteCustomizeSql(customizeSQL, &adminDao.AccountPayeeCheck{})
	if err != nil {
		t.Fatalf("TestDeleteByExample2 失敗：%v", err)
	}
	fmt.Printf("%+v\n", i)
}

func TestInsertAllowingNull(t *testing.T) {
	i := 466
	s := "test"
	time := time.Now()
	ap := adminDao.AccountPayeeCheck{
		ID:          &i,
		UID:         &i,
		Type:        &i,
		Description: &s,
		Status:      &i,
		CheckID:     &i,
		CheckTime:   &time,
		UpdateTime:  &time,
		CreatedTime: &time,
	}

	insert, err := adminDao.InsertAllowingNull(&ap)
	if err != nil {
		t.Fatalf("TestInsertAllowingNull 失敗：%v", err)
	}

	fmt.Println("----------------------------")
	fmt.Printf("%+v\n", insert)
	fmt.Println("----------------------------")
}

func TestInsertAllowingNullCustomizeSQL(t *testing.T) {
	i := 4
	//s := "test"
	time := time.Now()
	ap := adminDao.AccountPayeeCheck{
		//ID:          &i,
		UID:  &i,
		Type: &i,
		//Description: nil,
		Status:      &i,
		CheckID:     &i,
		CheckTime:   &time,
		UpdateTime:  &time,
		CreatedTime: &time,
	}
	customizeSQL := func(db *gorm.DB) *gorm.DB {
		db = db.Table("account_payee_check")
		return db
	}
	insert, err := adminDao.InsertAllowingNullCustomizeSQL(customizeSQL, &ap)
	if err != nil {
		t.Fatalf("TestInsertAllowingNullCustomizeSQL 失敗：%v", err)
	}

	fmt.Println("----------------------------")
	fmt.Printf("%+v\n", insert)
	fmt.Println("----------------------------")
}

func TestInsertIgnoringNull(t *testing.T) {
	i := 4
	s := "test"
	time := time.Now()
	ap := adminDao.AccountPayeeCheck{
		//ID:          &i,
		UID:         &i,
		Type:        &i,
		Description: &s,
		//Status:      &i,
		CheckID:     &i,
		CheckTime:   &time,
		UpdateTime:  &time,
		CreatedTime: &time,
	}

	insert, err := adminDao.InsertIgnoringNull(ap)
	if err != nil {
		t.Fatalf("TestInsertIgnoringNull 失敗：%v", err)
	}

	fmt.Println("----------------------------")
	fmt.Printf("%+v\n", insert)
	fmt.Println("----------------------------")
}

func TestInsertIgnoringNullCustomizeSQL(t *testing.T) {
	i := 4
	s := "test"
	time := time.Now()
	ap := adminDao.AccountPayeeCheck{
		//ID:          &i,
		UID:         &i,
		Type:        &i,
		Description: &s,
		//Status:      &i,
		CheckID:     &i,
		CheckTime:   &time,
		UpdateTime:  &time,
		CreatedTime: &time,
	}

	customizeSQL := func(db *gorm.DB) *gorm.DB {
		//db = db.Table("xxx")
		return db
	}

	insert, err := adminDao.InsertIgnoringNullCustomizeSQL(customizeSQL, ap)
	if err != nil {
		t.Fatalf("TestInsertIgnoringNullCustomizeSQL 失敗：%v", err)
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
		UID:         &i,
		Type:        &i,
		Description: &s,
		Status:      &i,
		CheckID:     &i,
		CheckTime:   &time,
		UpdateTime:  &time,
		CreatedTime: &time,
	}
	i2 := 5
	s2 := "test2"
	ap2 := adminDao.AccountPayeeCheck{
		UID:         &i2,
		Type:        &i2,
		Description: &s2,
		CheckTime:   &time,
		UpdateTime:  &time,
		CreatedTime: &time,
	}
	q = append(q, ap, ap2)

	insert, err := adminDao.InsertIgnoringNullList(q)
	if err != nil {
		t.Fatalf("TestInsertIgnoringNullList 失敗：%v", err)
	}

	fmt.Println("----------------------------")
	fmt.Printf("%+v\n", insert)
	fmt.Println("----------------------------")
}

func TestUpdateIgnoringNull(t *testing.T) {
	//uid := 66
	typet := 6
	description := "test77"
	status := 6
	checkID := 6
	timeAdmin := time.Now()

	updatesReq := adminDao.AccountPayeeCheck{
		UID:         nil,
		Type:        &typet,
		Description: &description,
		Status:      &status,
		CheckID:     &checkID,
		CheckTime:   &timeAdmin,
		UpdateTime:  &timeAdmin,
		CreatedTime: &timeAdmin,
	}

	id := 4
	whereReq := adminDao.AccountPayeeCheck{
		ID: &id,
	}

	specificTime := time.Date(2023, time.September, 15, 12, 0, 0, 0, time.UTC)

	customizeSQL := func(db *gorm.DB) *gorm.DB {
		db = db.Where("created_time >= ?", specificTime)
		db = db.Where("created_time <= ?", time.Now())
		//db = db.Where("status = ?", 1)
		return db
	}
	rep, err := adminDao.UpdateIgnoringNull(&updatesReq, &whereReq, customizeSQL)
	if err != nil {
		t.Fatalf("TestUpdateIgnoringNull 失敗：%v", err)
	}
	fmt.Printf(strconv.FormatInt(rep, 10))
}

func TestUpdateAllowingNull(t *testing.T) {

	uid := 888
	typet := 8
	description := "test888"
	status := 888
	//checkID := 8
	timeA := time.Now()
	updatesReq := adminDao.AccountPayeeCheck{
		UID:         &uid,
		Type:        &typet,
		Description: &description,
		Status:      &status,
		CheckID:     nil,
		CheckTime:   &timeA,
		UpdateTime:  &timeA,
		CreatedTime: &timeA,
	}

	id := 4
	whereReq := adminDao.AccountPayeeCheck{
		ID: &id,
	}

	specificTime := time.Date(2023, time.September, 15, 12, 0, 0, 0, time.UTC)

	customizeSQL := func(db *gorm.DB) *gorm.DB {
		db = db.Where("created_time >= ?", specificTime)
		db = db.Where("created_time <= ?", time.Now())
		//db = db.Where("status = ?", 1)

		return db
	}
	rep, err := adminDao.UpdateAllowingNull(updatesReq, whereReq, customizeSQL)
	if err != nil {
		t.Fatalf("TestUpdateAllowingNull 失敗：%v", err)
	}

	fmt.Printf(strconv.FormatInt(rep, 10))

}

func TestUpdateIgnoringNullByPrimaryKey(t *testing.T) {
	uid := 55
	typet := 55
	description := "test55"
	time := time.Now()
	updatesReq := adminDao.AccountPayeeCheck{
		UID:         &uid,
		Type:        &typet,
		Description: &description,
		CheckTime:   &time,
		UpdateTime:  &time,
		CreatedTime: &time,
	}
	id := 10
	selective, err := adminDao.UpdateIgnoringNullByPrimaryKey(id, updatesReq)
	if err != nil {
		t.Fatalf("TestUpdateIgnoringNullByPrimaryKey 失敗：%v", err)
	}
	fmt.Printf(strconv.FormatInt(selective, 10))

}

func TestUpdateAllowingNullByPrimaryKey(t *testing.T) {
	uid := 55
	typet := 55
	description := "test55"
	time := time.Now()
	updatesReq := adminDao.AccountPayeeCheck{
		UID:         &uid,
		Type:        &typet,
		Description: &description,
		CheckTime:   &time,
		UpdateTime:  &time,
		CreatedTime: &time,
	}

	id := 10

	rep, err := adminDao.UpdateAllowingNullByPrimaryKey(id, updatesReq)
	if err != nil {
		t.Fatalf("TestUpdateAllowingNullByPrimaryKey 失敗：%v", err)
	}
	fmt.Printf(strconv.FormatInt(rep, 10))
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
	fmt.Printf("%+v\n", *result.ID)
	fmt.Printf("%+v\n", *result.UID)
	fmt.Printf("%+v\n", *result.Type)
	fmt.Printf("%+v\n", *result.Description)
	fmt.Printf("%+v\n", *result.Status)
	fmt.Printf("%+v\n", *result.CheckID)
	fmt.Printf("%+v\n", *result.CheckTime)
	fmt.Printf("%+v\n", *result.UpdateTime)
	fmt.Printf("%+v\n", *result.CreatedTime)
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
		fmt.Printf("%+v\n", *rueslt.Description)
		fmt.Printf("%+v\n", *rueslt.ID)
		fmt.Printf("%+v\n", *rueslt.Username)
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
