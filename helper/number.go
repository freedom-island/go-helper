package helper

import (
	"golang.org/x/exp/constraints"
	"regexp"
)

// Avg 求平均值
func Avg[T constraints.Integer | constraints.Float](inputSlice []T) T {
	var sum T
	for _, element := range inputSlice {
		sum += element
	}
	return sum / T(len(inputSlice))
}

// Max 求最大值
func Max[T constraints.Integer | constraints.Float](arr ...T) T {
	t := arr[0]
	for _, v := range arr {
		if v > t {
			t = v
		}
	}
	return t
}

// Min 求最小值
func Min[T constraints.Integer | constraints.Float](arr ...T) T {
	t := arr[0]
	for _, v := range arr {
		if v < t {
			t = v
		}
	}
	return t
}

// RemoveNonDigits 移除字符串中的非数字字符
func RemoveNonDigits(input string) string {
	// 正则表达式，匹配非数字字符
	re := regexp.MustCompile(`\D`)
	// 用空字符串替换所有匹配的非数字字符
	return re.ReplaceAllString(input, "")
}
