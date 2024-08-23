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

func TestListAccountPayeeChecks(t *testing.T) {

	userRandomId := "1"
	status := enum.WAIT

	adminMember := adminDao.AccountPayeeCheckBasicDao{}

	pagination := model.Pagination{Page: 1, Limit: 2}

	total, results, err := adminMember.ListAccountPayeeChecks(&userRandomId, &status, &pagination)

	if err != nil {
		t.Fatalf("TestSelectByExample 失敗：%v", err)
	}

	fmt.Printf(" total : %+v\n", total)

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

func TestSumTotalStatusSUM(t *testing.T) {

	adminMember := adminDao.AccountPayeeCheckBasicDao{}

	customizeSQL := func(db *gorm.DB) *gorm.DB {
		return db
	}

	price, err := adminMember.SumTotalStatusSUM(customizeSQL)

	if err != nil {
		t.Fatalf("SumTotalStatusSUM 失敗：%v", err)
	}

	fmt.Printf(" price : %+v\n", price)

}

func TestSelectByPrimaryKey2(t *testing.T) {
	var result adminDao.AccountPayeeCheck

	primaryKey := 1

	err := adminDao.SelectByPrimaryKey(primaryKey, &result, &adminDao.AccountPayeeCheck{})

	if err != nil {
		t.Fatalf("TestSelectByExample 失敗：%v", err)
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
	primaryKey := 1
	err := adminDao.DeleteByPrimaryKey(primaryKey, &adminDao.AccountPayeeCheck{})
	if err != nil {
		t.Fatalf("TestSelectByExample 失敗：%v", err)
	}
}

func TestDeleteByExample2(t *testing.T) {
	uir := 1
	customizeSQL := func(db *gorm.DB) *gorm.DB {
		db = db.Where("uid = ?", uir)
		return db
	}
	err := adminDao.DeleteByExample(customizeSQL, &adminDao.AccountPayeeCheck{})
	if err != nil {
		t.Fatalf("TestSelectByExample 失敗：%v", err)
	}
}

func TestInsert2(t *testing.T) {
	i := 4
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

	insert, err := adminDao.Insert(&ap, &adminDao.AccountPayeeCheck{})
	if err != nil {
		t.Fatalf("TestSelectByExample 失敗：%v", err)
	}

	fmt.Println("----------------------------")
	fmt.Printf("%+v\n", insert)
	fmt.Println("----------------------------")
}

func TestInsertSelective2(t *testing.T) {
	i := 4
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

	insert, err := adminDao.InsertSelective(ap, &adminDao.AccountPayeeCheck{})
	if err != nil {
		t.Fatalf("TestSelectByExample 失敗：%v", err)
	}

	fmt.Println("----------------------------")
	fmt.Printf("%+v\n", insert)
	fmt.Println("----------------------------")
}

func TestUpdateByExampleSelective2(t *testing.T) {
	uid := 66
	typet := 6
	description := "test77"
	status := 6
	checkID := 6
	timeAdmin := time.Now()

	updatesReq := adminDao.AccountPayeeCheck{
		UID:         &uid,
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
	rep, err := adminDao.UpdateByExampleSelective(&updatesReq, &whereReq, customizeSQL, &adminDao.AccountPayeeCheck{})
	if err != nil {
		t.Fatalf("UpdateByCustomizeSQL 失敗：%v", err)
	}
	fmt.Printf(strconv.FormatInt(rep, 10))
}

func TestUpdateByExample2(t *testing.T) {

	uid := 888
	typet := 8
	description := "test888"
	status := 888
	checkID := 8
	timeA := time.Now()
	updatesReq := adminDao.AccountPayeeCheck{
		UID:         &uid,
		Type:        &typet,
		Description: &description,
		Status:      &status,
		CheckID:     &checkID,
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
	rep, err := adminDao.UpdateByExample(updatesReq, whereReq, customizeSQL, &adminDao.AccountPayeeCheck{})
	if err != nil {
		t.Fatalf("UpdateByExampleCustomizeSQL 失敗：%v", err)
	}

	fmt.Printf(strconv.FormatInt(rep, 10))

}

func TestUpdateByPrimaryKeySelective2(t *testing.T) {
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
	selective, err := adminDao.UpdateByPrimaryKeySelective(id, updatesReq, &adminDao.AccountPayeeCheck{})
	if err != nil {
		t.Fatalf("UpdateByPrimaryKeySelective 失敗：%v", err)
	}
	fmt.Printf(strconv.FormatInt(selective, 10))

}

func TestUpdateByPrimaryKey2(t *testing.T) {
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

	rep, err := adminDao.UpdateByPrimaryKey(id, updatesReq, &adminDao.AccountPayeeCheck{})
	if err != nil {
		t.Fatalf("UpdateByPrimaryKeySelective 失敗：%v", err)
	}
	fmt.Printf(strconv.FormatInt(rep, 10))
}
