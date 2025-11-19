package helper

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"regexp"
	"strings"
	"unicode"
)

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

// SplitNumberAndChinese 拆分数字和汉字结合在一起的情况
func SplitNumberAndChinese(input string) (number string, chinese string, ok bool) {
	// 正则表达式匹配：数字开头（可能包含小数），后面跟着中文字符
	re := regexp.MustCompile(`^([0-9]+\.?[0-9]*)([\p{Han}]+)$`)
	matches := re.FindStringSubmatch(input)

	if len(matches) != 3 {
		return "", "", false
	}

	numberPart := matches[1]
	chinesePart := matches[2]

	// 检查数字部分是否超过三位数（不包括小数点）
	digitCount := 0
	for _, r := range numberPart {
		if unicode.IsDigit(r) {
			digitCount++
		}
	}

	if digitCount <= 3 {
		return "", "", false
	}

	return numberPart, chinesePart, true
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

// FilterNumbers 只保留数字元素（字符串形式的数字）
func FilterNumbers(arr []string) []string {
	var result []string
	numberRegex := regexp.MustCompile(`^-?\d+(\.\d+)?$`)

	for _, item := range arr {
		if numberRegex.MatchString(item) {
			result = append(result, item)
		}
	}
	return result
}

// ToCamelCase 将下划线分隔的名称转换为驼峰式命名
func ToCamelCase(s string) string {
	words := strings.Split(s, "_")
	caser := cases.Title(language.English)

	for i := 1; i < len(words); i++ {
		if words[i] == "" {
			continue
		}
		words[i] = caser.String(words[i])
	}
	return strings.Join(words, "")
}
