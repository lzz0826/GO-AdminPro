package utils

import (
	"AdminPro/server/task"
	"fmt"
)

const (
	/**
	 * 1个字节流类型
	 * */
	TYPE_INT_1 = 0

	/**
	 * 4个字节流类型
	 * **/
	TYPE_INT_4 = 1

	/**
	 * UTF-16  字节流转码
	 * */
	TYPE_STRING_UTF16 = 2

	/**
	 * ascii unicode 字节流转码
	 * */
	TYPE_STRING_UNICODE = 3

	/**
	 * 数组
	 */
	TYPE_INT_4_ARRAY = 5

	/**
	 * 1 数组
	 * */

	TYPE_INT_1_ARRAY = 6

	/**
	 * 一个bute itemid < 50
	 */
	TYPE_BYTE = 4

	UTF_16          = "UTF-16"
	ISO_10646_UCS_2 = "ISO-10646-UCS-2"
	UNICODE         = "UNICODE"

	MAX_CP_SIZE = 1024 * 16

	/**
	 * 返回包内容开始位置
	 * */
	FEE_DATA_BT_LENGTH = 20

	/**
	 * 解析最多的山数个数
	 */
	MAX_FOR = 100
	/**
	 * 协议中协议内容开始的位置
	 * */
	START_LOCATION = 30 + 54
)

// PickAll 函数实现
func PickAll(bt2 []byte, int2 [][]int) map[int]interface{} {
	resMap := make(map[int]interface{})
	beforeSite := START_LOCATION
	fmt.Printf("bt2:%+v \n", bt2)

	beforeSite -= 54 //内部通讯协议读取位置沿用原先下标
	isGive := true
	for i := 0; i < MAX_FOR; i++ {
		//itemId 只有一位
		if (beforeSite + 2) >= len(bt2) {
			break
		}
		beforeSite += 1
		itemId := int(bt2[beforeSite] & 0xff)
		//fmt.Printf("itemId:%d \n", itemId)
		isGive = true
		for _, j := range int2 {
			if j[0] == itemId {
				fmt.Printf("itemId:%d \n", j[0])
				isGive = false
				switch j[1] {
				case TYPE_BYTE:
					resMap[j[0]] = itemId
				case TYPE_INT_1:
					beforeSite++
					if (beforeSite + 1) >= len(bt2) {
						break
					}
					resMap[j[0]] = int(bt2[beforeSite] & 0xff)
				case TYPE_INT_4:
					if (beforeSite + 4) > len(bt2) {
						fmt.Println("包长度出错")
						break
					}
					beforeSite = beforeSite + 2
					typeInt4 := ByteArrayToInt(bt2, beforeSite+1)
					beforeSite += 4
					resMap[j[0]] = typeInt4
				case TYPE_STRING_UTF16:
					if (beforeSite + 2) > len(bt2) {
						break
					}
					inputSize := ByteArrayToShortInt(bt2, beforeSite+1)
					beforeSite += 2
					if (beforeSite + inputSize + 1) > len(bt2) {
						break
					}
					if inputSize != 0 {
						content, _, _ := Utf162Utf8(bt2[beforeSite+1 : beforeSite+inputSize+1])
						resMap[j[0]] = string(content)
						//resMap[j[0]] = string(bt2[beforeSite+1 : beforeSite+inputSize+1])
					}
					beforeSite += inputSize
				case TYPE_STRING_UNICODE:
					if (beforeSite + 2) > len(bt2) {
						break
					}
					inputSize := ByteArrayToShortInt(bt2, beforeSite+1)
					beforeSite += 2
					if (beforeSite + inputSize) > len(bt2) {
						break
					}
					var typeStringUnicode string
					if inputSize != 0 {
						typeStringUnicode = string(bt2[beforeSite+1 : beforeSite+inputSize+1])
					}
					resMap[j[0]] = typeStringUnicode
					beforeSite += inputSize
				case TYPE_INT_4_ARRAY:
					if (beforeSite + 2) > len(bt2) {
						break
					}
					inputSize := ByteArrayToShortInt(bt2, beforeSite+1)
					beforeSite += 2
					if (beforeSite + inputSize) > len(bt2) {
						break
					}

					intArray := make([]int, inputSize/4)
					for k := range intArray {
						intArray[k] = ByteArrayToInt(bt2, beforeSite+k*4+1)
					}

					resMap[j[0]] = intArray
					beforeSite += inputSize

				case TYPE_INT_1_ARRAY:
					if (beforeSite + 2) > len(bt2) {
						break
					}
					inputSize := ByteArrayToShortInt(bt2, beforeSite+1)
					beforeSite += 2
					if (beforeSite + inputSize) > len(bt2) {
						break
					}

					intArray := make([]int, inputSize)
					for k := range intArray {
						intArray[k] = int(bt2[beforeSite+k+1] & 0xff)
					}

					resMap[j[0]] = intArray
					beforeSite += inputSize
				}
				break
			}
		}

		if isGive {
			if itemId < 50 {
				continue
			} else if 50 <= itemId && itemId < 128 {
				beforeSite++
				continue
			} else if itemId >= 128 {
				if (beforeSite + 2) > len(bt2) {
					break
				}
				inputSize := ByteArrayToShortInt(bt2, beforeSite+1)
				beforeSite += 2
				if (beforeSite + inputSize) > len(bt2) {
					break
				}
				beforeSite += inputSize
				continue
			}
		}
	}

	return resMap
}

func packHead() []byte {
	bt := make([]byte, MAX_CP_SIZE)
	bt[0] = task.HEADER_INDICATER_0
	bt[1] = task.HEADER_INDICATER_1
	bt[2] = task.HEADER_INDICATER_2
	bt[3] = task.HEADER_INDICATER_3
	bt[task.SRP_PACKE_LEVEL] = byte(0)
	return bt
}

func PackAll(int2 [][]interface{}, requestCode int) []byte {
	bt := packHead()
	bt[task.SRP_REQUEST_HIGH] = (byte)(requestCode >> 8)
	bt[task.SRP_REQUEST_LOW] = (byte)(requestCode)
	feeDataBtLength := FEE_DATA_BT_LENGTH
	bt[feeDataBtLength] = (byte)(len(int2))
	var itemId int
	var feePageDataBt []byte
	for _, objl := range int2 {
		itemId = objl[0].(int)
		feeDataBtLength++
		bt[feeDataBtLength] = byte(itemId)
		types := objl[2]
		switch types {
		case TYPE_BYTE:
			//不限字节
			continue
		case TYPE_INT_1:
			//跟一位
			i_ := objl[1].(int)
			feeDataBtLength++
			bt[feeDataBtLength] = byte(i_)
		case TYPE_INT_4:
			if itemId >= 128 {
				feeDataBtLength++
				bt[feeDataBtLength] = byte(4 >> 8)
				feeDataBtLength++
				bt[feeDataBtLength] = byte(4)
			}
			feePageDataBt = IntToByteArray(objl[1].(int))
			copy(bt[feeDataBtLength+1:], feePageDataBt)
			feeDataBtLength = feeDataBtLength + 4
		case TYPE_STRING_UNICODE:
			feePageDataBt = StringToBytesUNICODE(objl[1].(string))
			//放长度
			feeDataBtLength++
			bt[feeDataBtLength] = byte(len(feePageDataBt) >> 8)
			feeDataBtLength++
			bt[feeDataBtLength] = byte(len(feePageDataBt))
			copy(bt[feeDataBtLength+1:], feePageDataBt)
			feeDataBtLength = feeDataBtLength + len(feePageDataBt)
		case TYPE_STRING_UTF16:
			feePageDataBt = StringToBytesUTF16(objl[1].(string))
			feeDataBtLength++
			bt[feeDataBtLength] = byte(len(feePageDataBt) >> 8)
			feeDataBtLength++
			bt[feeDataBtLength] = byte(len(feePageDataBt))
			copy(bt[feeDataBtLength+1:], feePageDataBt)
			feeDataBtLength = feeDataBtLength + len(feePageDataBt)
		case TYPE_INT_4_ARRAY:
			intArray := objl[1].([]int)
			feePageDataBt = make([]byte, len(intArray)*4)
			for i3 := 0; i3 < len(intArray); i3++ {
				if intArray[i3] == 0 {
					break
				}
				i4 := i3 * 4
				bs := IntToByteArray(intArray[i3])
				feePageDataBt[i4] = bs[0]
				feePageDataBt[i4+1] = bs[1]
				feePageDataBt[i4+2] = bs[2]
				feePageDataBt[i4+3] = bs[3]
			}
			//放长度
			feeDataBtLength++
			bt[feeDataBtLength] = byte(len(feePageDataBt) >> 8)
			feeDataBtLength++
			bt[feeDataBtLength] = byte(len(feePageDataBt))
			copy(bt[feeDataBtLength+1:], feePageDataBt)
			feeDataBtLength = feeDataBtLength + len(feePageDataBt)
		case TYPE_INT_1_ARRAY:
			//获取int数组
			intArray_ := objl[1].([]int)
			feePageDataBt = make([]byte, len(intArray_))
			for i3 := 0; i3 < len(intArray_); i3++ {
				if intArray_[i3] == 0 {
					break
				}
				a := intArray_[i3]
				feePageDataBt[i3] = byte(a)
			}
			//放长度
			feeDataBtLength++
			bt[feeDataBtLength] = byte(len(feePageDataBt) >> 8)
			feeDataBtLength++
			bt[feeDataBtLength] = byte(len(feePageDataBt))
			copy(bt[feeDataBtLength+1:], feePageDataBt)
			feeDataBtLength = feeDataBtLength + len(feePageDataBt)
		}
	}
	feeDataBtLength++
	//设置返回长度
	bt[task.SRP_SIZE_HIGH] = byte(feeDataBtLength >> 8)
	bt[task.SRP_SIZE_LOW] = byte(feeDataBtLength)
	bt2 := make([]byte, feeDataBtLength)
	// 复制数组
	copy(bt2, bt[:feeDataBtLength])
	return bt2
}
