package test

import (
	"AdminPro/dao/model/adminDao"
	"AdminPro/dao/service/admin"
	"encoding/json"
	"fmt"
	"testing"
)

func TestCh(t *testing.T) {

	permits := admin.Convert()

	fmt.Printf("%+v\n", permits)

	//-----------  切片轉 map      -------
	// 创建一个 map，用 ID 作为键
	permitMap := make(map[string]adminDao.PermitDAO, len(permits))
	// 遍历 permits，将每个 permit 的 ID 作为键存储到 map 中
	for _, permit := range permits {
		permitMap[permit.ID] = permit
	}
	fmt.Printf("%+v\n", permitMap)

	//----------------- map 取id 轉List ---------------
	permitIds := make([]string, 0, len(permitMap))
	for _, club := range permitMap {
		permitIds = append(permitIds, club.ID)
	}
	fmt.Printf("%+v\n", permitIds)

	//----------------- 将 Go 结构体转换为 JSON 格式的字节切片 ---------------
	//物件转JSON
	jsonData, err := json.Marshal(permits)
	if err != nil {
		fmt.Println("JSON marshaling failed:", err)
		return
	}
	fmt.Println("JSON data:", string(jsonData))

}
