package utils

// 向切片前添加
func SliceAddFrontAny[T comparable](slice *[]T, item T) {
	newSlice := make([]T, len(*slice)+1)
	newSlice[0] = item
	copy(newSlice[1:], *slice)
	*slice = newSlice
}

// 向切片任意位置添加
func SliceAddAnyPoint[T comparable](slice *[]T, item T, index int) {
	if index < 0 || index > len(*slice) {
		print("index out of range")
		return
	}
	newSlice := make([]T, len(*slice)+1)
	copy(newSlice, (*slice)[:index])
	newSlice[index] = item
	copy(newSlice[index+1:], (*slice)[index:])
	*slice = newSlice
}

// 从切片获取某个元素的下标index
func GetSliceAntIndex[T comparable](slice []T, item T) int {
	for i, v := range slice {
		if v == item {
			return i
		}
	}
	return -1
}

// 判断切片中是否包含某个元素(T any)
func SliceContainsAny[T comparable](slice []T, item T) bool {
	for _, a := range slice {
		if a == item {
			return true
		}
	}
	return false
}

// 判断int切片中是否包含某个int元素
func SliceContainsInt(ids []int, id int) bool {
	for _, v := range ids {
		if v == id {
			return true
		}
	}
	return false
}

// 判断字符串切片中是否包含某个字符串
func SliceContainsString(arr []string, e string) bool {
	for _, v := range arr {
		if v == e {
			return true
		}
	}
	return false
}

// 切片中删除某个int元素并返回修改后的切片
func SliceDelInt(arr []int, e int) []int {
	for i := 0; i < len(arr); i++ {
		if arr[i] == e {
			arr = append(arr[:i], arr[i+1:]...)
			i--
		}
	}
	return arr
}

// 切片中删除某个string元素并返回修改后的切片
func SliceDelString(arr []string, e string) []string {
	for i := 0; i < len(arr); i++ {
		if arr[i] == e {
			arr = append(arr[:i], arr[i+1:]...)
			i--
		}
	}
	return arr
}

// RemoveDuplicates 去重切片的通用方法，并直接修改原切片
func RemoveDuplicatesPort[T comparable](slice *[]T) {
	seen := make(map[T]struct{})
	unique := []T{}

	for _, item := range *slice {
		if _, exists := seen[item]; !exists {
			seen[item] = struct{}{}
			unique = append(unique, item)
		}
	}

	// 修改原切片的值
	*slice = unique
}

// RemoveDuplicates 去重切片的通用方法
func RemoveDuplicates[T comparable](slice []T) []T {
	// 使用 map 来跟踪出现过的元素
	seen := make(map[T]struct{})
	unique := []T{}

	for _, item := range slice {
		if _, exists := seen[item]; !exists {
			seen[item] = struct{}{}
			unique = append(unique, item)
		}
	}

	return unique
}

// Filter 切片筛选的通用方法
func Filter[T any](slice []T, predicate func(T) bool) []T {
	var result []T
	for _, item := range slice {
		if predicate(item) {
			result = append(result, item)
		}
	}
	return result
}

// Map 切片映射的通用方法
func Map[T any, U any](slice []T, mapper func(T) U) []U {
	var result []U
	for _, item := range slice {
		result = append(result, mapper(item))
	}
	return result
}

// Merge 合并两个切片
func Merge[T any](a, b []T) []T {
	return append(a, b...)
}

// Contains 检查切片中是否包含指定元素
func Contains[T comparable](slice []T, element T) bool {
	for _, item := range slice {
		if item == element {
			return true
		}
	}
	return false
}

func ContainsAny[T comparable](a []T, x T) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

func IsEmptyArray[T any](arr []T) bool {
	if arr == nil || len(arr) == 0 {
		return true
	}
	return false
}
