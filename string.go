package go_helper

import (
	"unicode"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// RemoveSingleChar 移除单个字符
func RemoveSingleChar(input []string, include []string) []string {
	var result []string

	for _, str := range input {
		chineseCount := 0
		for _, r := range str {
			for _, s := range include {
				if s == string(r) {
					chineseCount++
					break
				}
			}

			if chineseCount > 1 {
				continue
			}
		}
		if chineseCount != 1 {
			result = append(result, str)
		}
	}

	return result
}

// RemoveSingleChinese 移除单个中文字符
func RemoveSingleChinese(input []string, exclude []string) []string {
	var result []string

	for _, str := range input {
		chineseCount := 0
		for _, r := range str {
			// 检查字符是否为中文
			if unicode.Is(unicode.Han, r) {
				chineseCount++

				// exclude 中包含的字符，则跳过该字符
				for _, s := range exclude {
					if s == string(r) {
						chineseCount++
						break
					}
				}

				if chineseCount > 1 {
					continue
				}
			}
		}
		// 仅包含一个中文字符的元素不添加到结果数组中
		if chineseCount != 1 {
			result = append(result, str)
		}
	}

	return result
}

// HasUnicodeText 检查字符是否包含中文
func HasUnicodeText(str string) bool {
	for _, r := range str {
		if unicode.Is(unicode.Han, r) {
			return true
		}
	}
	return false
}

// IsComposedOf 检查字符串 s 是否完全由字符 c 构成
func IsComposedOf(s string, c rune) bool {
	if len(s) == 0 {
		return true // 空字符串视为由任何字符构成（根据需求可调整）
	}

	for _, ch := range s {
		if ch != c {
			return false
		}
	}
	return true
}
