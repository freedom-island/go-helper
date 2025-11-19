package helper

import (
	"golang.org/x/exp/constraints"
	"log"
	"regexp"
	"strconv"
	"strings"
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

// SplitInt 切割整数
func SplitInt[T constraints.Integer | constraints.Float](val T, max T) []T {
	var result []T
	for val > max {
		result = append(result, max)
		val -= max
	}
	if val > 0 {
		result = append(result, max)
	}

	return result
}

// RemoveNonDigits 移除字符串中的非数字字符
func RemoveNonDigits(input string) string {
	// 正则表达式，匹配非数字字符
	re := regexp.MustCompile(`\D`)
	// 用空字符串替换所有匹配的非数字字符
	return re.ReplaceAllString(input, "")
}

// StringToFloat64 字符串转浮点数
func StringToFloat64(text string) float64 {
	text = strings.ReplaceAll(text, "D", "0")
	text = strings.ReplaceAll(text, "U", "0")
	text = strings.ReplaceAll(text, "C", "0")
	text = strings.ReplaceAll(text, "J", "0")
	text = strings.ReplaceAll(text, " ", "")
	parseFloat, err := strconv.ParseFloat(text, 64)
	if err != nil {
		log.Panic("StringToFloat64 error from:", text, err)
	}

	return parseFloat
}

// StringToInt 字符串转整数
func StringToInt(text string) int {
	parseInt := StringToIntNoErr(text)
	if parseInt == -99999999 {
		log.Panic("StringToInt error from:", text)
	}

	return parseInt
}

// StringToIntNoErr 字符串转整数(忽略错误)
func StringToIntNoErr(text string) int {
	text = strings.ReplaceAll(text, "D", "0")
	text = strings.ReplaceAll(text, "U", "0")
	text = strings.ReplaceAll(text, "C", "0")
	text = strings.ReplaceAll(text, " ", "")
	// 检查是否包含小数点
	if strings.Contains(text, ".") {
		text = strings.Split(text, ".")[0]
	}

	parseInt, err := strconv.Atoi(text)
	if err != nil {
		log.Println("StringToInt error from:", text, err)
		return -99999999
	}

	return parseInt
}
