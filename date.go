package helper

import (
	"log"
	"strings"
	"time"
)

// DateNowInt 获取今日时间（身份证型时间 19960315）
func DateNowInt() int {
	var now = time.Now()
	return now.Year()*10000 + int(now.Month())*100 + now.Day()
}

// DateArithmeticInt 获取加减运算的今日时间（身份证型时间 19960315）
func DateArithmeticInt(change time.Duration) int {
	var now = time.Now().Add(change)
	return now.Year()*10000 + int(now.Month())*100 + now.Day()
}

// DateEqual 日期相等检查
func DateEqual(date1, date2 time.Time) bool {
	y1, m1, d1 := date1.Date()
	y2, m2, d2 := date2.Date()
	return y1 == y2 && m1 == m2 && d1 == d2
}

// StringToTime 字符串转换为时间
func StringToTime(text string) time.Time {
	text = strings.ReplaceAll(text, "：", ":")
	text = strings.ReplaceAll(text, "D", "0")
	text = strings.ReplaceAll(text, "U", "0")

	parseTime, err := time.Parse(time.TimeOnly, text)
	if err != nil {
		log.Printf("StringToTime error from: '%s'", text)
		log.Panicln(err)
	}

	return parseTime
}
