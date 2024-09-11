package utils

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
