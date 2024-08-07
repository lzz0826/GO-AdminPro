package test

import (
	"AdminPro/common/model"
	"AdminPro/common/utils"
	"AdminPro/dao/model/adminDao"
	"fmt"
	"gorm.io/gorm"
	"strconv"
	"testing"
	"time"
)

func TestSelectByExample(t *testing.T) {
	i := 3
	adminMember := adminDao.AccountPayeeCheckDao{}

	pagination := model.Pagination{Page: 1, Limit: 2}

	//排序 未审核的优先排前面 其余按创建时间倒续
	sql := utils.WithOrderBySQL("case when status = 0 then 0 else 1 end asc, created_time desc, id desc")

	results, err := adminMember.SelectByExample(&i, &i, sql, &pagination)

	if err != nil {
		t.Fatalf("TestSelectByExample 失敗：%v", err)
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

func TestSelectByExample2(t *testing.T) {

	adminMember := adminDao.AccountPayeeCheckDao{}
	uid := 1
	status := 0

	customizeSQL := func(db *gorm.DB) *gorm.DB {
		db = db.Where("uid = ?", uid)
		db = db.Where("status = ?", status)
		db = db.Scopes(utils.WithPagination(1, 2))
		db = db.Order("case when status = 0 then 0 else 1 end asc, created_time desc, id desc")
		return db
	}

	example, err := adminMember.SelectByExample2(customizeSQL)

	if err != nil {
		fmt.Print(err.Error())
	}

	if err != nil {
		t.Fatalf("TestSelectByExample 失敗：%v", err)
	}

	for _, result := range example {
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

func TestSelectByPrimaryKey(t *testing.T) {
	i := 2
	adminMember := adminDao.AccountPayeeCheckDao{}

	result, err := adminMember.SelectByPrimaryKey(i)

	if err != nil {
		t.Fatalf("SelectByPrimaryKey 失敗：%v", err)
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

func TestDeleteByExample(t *testing.T) {
	i := 2
	adminMember := adminDao.AccountPayeeCheckDao{}

	err := adminMember.DeleteByExample(i)

	if err != nil {
		t.Fatalf("DeleteByExample 失敗：%v", err)
	}

}

func TestInsert(t *testing.T) {
	adminMember := adminDao.AccountPayeeCheckDao{}
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
	insert, err := adminMember.Insert(ap)
	if err != nil {
		t.Fatalf("Insert 失敗：%v", err)
	}

	fmt.Println("----------------------------")
	fmt.Printf("%+v\n", insert)
	fmt.Println("----------------------------")

}

func TestInsertSelective(t *testing.T) {
	adminMember := adminDao.AccountPayeeCheckDao{}
	uid := 99
	typet := 1
	description := "test"
	status := 0
	checkID := 22
	time := time.Now()

	ap := adminDao.AccountPayeeCheck{
		UID:         &uid,
		Type:        &typet,
		Description: &description,
		Status:      &status,
		CheckID:     &checkID,
		CheckTime:   &time,
		UpdateTime:  &time,
		CreatedTime: &time,
	}
	insert, err := adminMember.InsertSelective(ap)
	if err != nil {
		t.Fatalf("InsertSelective 失敗：%v", err)
	}

	fmt.Println("----------------------------")
	fmt.Printf("%+v\n", insert)
	fmt.Println("----------------------------")
}

func TestCountByExample(t *testing.T) {
	adminMember := adminDao.AccountPayeeCheckDao{}
	id := 14
	typet := 1

	ap := adminDao.AccountPayeeCheck{
		ID:   &id,
		Type: &typet,
	}
	example, err := adminMember.CountByExample(ap)
	if err != nil {
		t.Fatalf("CountByExample 失敗：%v", err)
	}

	fmt.Printf(strconv.FormatInt(example, 10))

}

func TestUpdateByExampleSelective(t *testing.T) {
	adminMember := adminDao.AccountPayeeCheckDao{}
	uid := 66
	typet := 6
	description := "test6"
	status := 6
	checkID := 6
	time := time.Now()
	updatesReq := adminDao.AccountPayeeCheck{
		UID:         &uid,
		Type:        &typet,
		Description: &description,
		Status:      &status,
		CheckID:     &checkID,
		CheckTime:   &time,
		UpdateTime:  &time,
		CreatedTime: &time,
	}

	id := 15
	whereReq := adminDao.AccountPayeeCheck{
		ID: &id,
	}
	err := adminMember.UpdateByExampleSelective(updatesReq, whereReq)
	if err != nil {
		t.Fatalf("UpdateByExampleSelective 失敗：%v", err)
	}
}

func TestUpdateByExample(t *testing.T) {

	adminMember := adminDao.AccountPayeeCheckDao{}
	uid := 777
	typet := 7
	description := "test7"
	status := 1
	checkID := 2
	time := time.Now()
	updatesReq := adminDao.AccountPayeeCheck{
		UID:         &uid,
		Type:        &typet,
		Description: &description,
		Status:      &status,
		CheckID:     &checkID,
		CheckTime:   &time,
		UpdateTime:  &time,
		CreatedTime: &time,
	}

	id := 14
	whereReq := adminDao.AccountPayeeCheck{
		ID: &id,
	}

	err := adminMember.UpdateByExample(updatesReq, whereReq)
	if err != nil {
		t.Fatalf("UpdateByExample 失敗：%v", err)
	}
}

func TestUpdateByPrimaryKeySelective(t *testing.T) {

	adminMember := adminDao.AccountPayeeCheckDao{}
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

	id := 14

	err := adminMember.UpdateByPrimaryKeySelective(id, updatesReq)
	if err != nil {
		t.Fatalf("UpdateByPrimaryKeySelective 失敗：%v", err)
	}
}

func TestUpdateByPrimaryKey(t *testing.T) {

	adminMember := adminDao.AccountPayeeCheckDao{}
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

	id := 14

	err := adminMember.UpdateByPrimaryKey(id, updatesReq)
	if err != nil {
		t.Fatalf("UpdateByPrimaryKeySelective 失敗：%v", err)
	}
}
