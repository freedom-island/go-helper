package helper

import (
	"golang.org/x/exp/constraints"
	"reflect"
)

// RemoveEmpty 移除为空的元素
func RemoveEmpty(slice []string) []string {
	var result []string
	for _, value := range slice {
		if value != "" {
			result = append(result, value)
		}
	}
	return result
}

// ReverseTwiceArr 数组反转
func ReverseTwiceArr(arr [][]interface{}) [][]interface{} {
	n := len(arr)
	newArr := make([][]interface{}, n)
	for i, v := range arr {
		newArr[n-1-i] = v
	}
	return newArr
}

// ReverseArr 数组反转
func ReverseArr[T any](arr []T) []T {
	n := len(arr)
	newArr := make([]T, n)
	for i, v := range arr {
		newArr[n-1-i] = v
	}
	return newArr
}

// SumArr 数组求和
func SumArr[T constraints.Ordered](arr []T) T {
	var sum T
	for _, v := range arr {
		sum += v
	}
	return sum
}

// MaxArr 求最大值
func MaxArr[T constraints.Integer | constraints.Float](arr []T) T {
	t := arr[0]
	for _, v := range arr {
		if v > t {
			t = v
		}
	}
	return t
}

// MaxAndIndexArr 求最大值和索引
func MaxAndIndexArr[T constraints.Integer | constraints.Float](arr []T) (T, int) {
	t := arr[0]
	var index int
	for i, v := range arr {
		if v > t {
			t = v
			index = i
		}
	}
	return t, index
}

// MinArr 数组最小值
func MinArr[T constraints.Integer | constraints.Float](arr []T) T {
	t := arr[0]
	for _, v := range arr {
		if v < t {
			t = v
		}
	}
	return t
}

// UniqueArr 数组去重
func UniqueArr[T comparable](inputSlice []T) []T {
	uniqueSlice := make([]T, 0, len(inputSlice))
	seen := make(map[T]struct{}, len(inputSlice))
	for _, element := range inputSlice {
		if _, ok := seen[element]; !ok {
			uniqueSlice = append(uniqueSlice, element)
			seen[element] = struct{}{}
		}
	}
	return uniqueSlice
}

// FieldValuesFloatArr 函数接受一个结构体类型的 slice 和一个字段名，返回该字段值的 slice
func FieldValuesFloatArr[T any](data []T, fieldName string) []float64 {
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
