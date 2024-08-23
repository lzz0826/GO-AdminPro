package utils

// [K comparable, V any] :定义范行
// K comparable：表示键的类型必须是可比较的（map 的键必须是可比较类型）。
// V any：表示值的类型可以是任意类型。
// m map[K]V：表示一个以键 K 为类型、值 V 为类型的 map。
// key K：表示要查找的键。
// defaultValue V：表示当键不存在时要返回的默认值。
// 取得MAP理的某值否则返回 自订预设值
func GetOrDefault[K comparable, V any](m map[K]V, key K, defaultValue V) V {
	if value, ok := m[key]; ok {
		return value
	}
	return defaultValue
}
