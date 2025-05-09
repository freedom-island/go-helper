package go_helper

import (
	"math/rand"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().Unix()))

// RandomStringWithCharset 生成伪随机的字符
func RandomStringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

// RandomString 生成伪随机的字符串
func RandomString(length int) string {
	return RandomStringWithCharset(length, charset)
}

// RandomInt 生成一个随机整数
func RandomInt(max int) int {
	if max <= 0 {
		return 0
	}
	return seededRand.Intn(max)
}

// RandomIntSlice 生成一个随机整数数组（不重复）
func RandomIntSlice(length int, has int) []uint {
	result := make([]uint, 0)
	if length <= 0 || has <= 0 {
		return result
	}
	if length < has {
		has = length
	} else if length > has {
		length = has
	}

	// 随机一个 max i 数组并打乱，摘取固定位数字
	randSlice := seededRand.Perm(length)
	for i := 0; i < has; i++ {
		result = append(result, uint(randSlice[i]))
	}
	return result
}
