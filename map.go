package helper

import "reflect"

// FieldValuesFloatMap 函数接受一个结构体类型值的 map 和一个字段名，返回字段值的切片
func FieldValuesFloatMap[T any](data map[string]T, fieldName string) []float64 {
	values := make([]float64, 0, len(data))
	for _, v := range data {
		val := reflect.ValueOf(v)
		// 检查是否是 struct 类型
		if val.Kind() == reflect.Struct {
			// 获取字段值
			field := val.FieldByName(fieldName)
			if field.IsValid() && field.CanInterface() {
				values = append(values, field.Float())
			}
		}
	}
	return values
}

// FieldValuesStringMap 函数接受一个结构体类型值的 map 和一个字段名，返回字段值的切片
func FieldValuesStringMap[T any](data map[string]T, fieldName string) []string {
	values := make([]string, 0, len(data))
	for _, v := range data {
		val := reflect.ValueOf(v)
		// 检查是否是 struct 类型
		if val.Kind() == reflect.Struct {
			// 获取字段值
			field := val.FieldByName(fieldName)
			if field.IsValid() && field.CanInterface() {
				values = append(values, field.String())
			}
		}
	}
	return values
}

// MapStringKeys 获取 map 的 key (类型为 string)
func MapStringKeys[T any](data map[string]T) []string {
	keys := make([]string, len(data))
	i := 0
	for k := range data {
		keys[i] = k
		i++
	}

	return keys
}

// MapIntKeys 获取 map 的 key (类型为 int)
func MapIntKeys[T any](data map[int]T) []int {
	keys := make([]int, len(data))
	i := 0
	for k := range data {
		keys[i] = k
		i++
	}

	return keys
}
