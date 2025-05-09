package go_helper

import "fmt"

// GetString 获取泛型内的 string
func GetString(data []interface{}, index int) string {
	val, ok := data[index].(string)
	if !ok {
		panic(fmt.Sprintf("interface get value is failed! index: %d, info: %s", index, data))
	}
	return val
}

// GetFloat 获取泛型内的 float
func GetFloat(data []interface{}, index int) float64 {
	val, ok := data[index].(float64)
	if !ok {
		panic(fmt.Sprintf("interface get value is failed! index: %d, info: %s", index, data))
	}
	return val
}

// GetInt 获取泛型内的 int
func GetInt(data []interface{}, index int) int {
	val, ok := data[index].(int)
	if !ok {
		panic(fmt.Sprintf("interface get value is failed! index: %d, info: %s", index, data))
	}
	return val
}
