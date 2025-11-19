package helper

import (
	"strings"
	"time"
)

// GetStockExchange 获取股票交易所；isSimple 决定是否返回简写
func GetStockExchange(symbol string, isSimple bool) string {
	if len(symbol) == 6 {
		switch symbol[0] {
		case '0', '2', '3':
			if isSimple {
				return "SZ"
			}
			return "XSHE" // 深圳证券交易所
		case '6':
			if isSimple {
				return "SH"
			}
			return "XSHG" // 上海证券交易所
		case '4', '8', '9':
			if isSimple {
				return "BJ"
			}
			// 北交所股票，通常在上海交易所显示
			return "BJSE"
		}
	}

	return "" // 未知交易所
}

// GetLimitOrderAmount 获取[最大、最小]下单数量
func GetLimitOrderAmount(stockCode string) []int {
	// 参数基础校验
	if len(stockCode) != 6 {
		return []int{0, 0}
	}

	// 根据代码前缀判断市场和板块
	switch stockCode[:2] {
	case "68": // 科创板
		return []int{100000, 200} // 限价申报上限10万股
	case "30": // 创业板
		return []int{300000, 100} // 限价申报上限30万股
	case "92": // 北交所
		return []int{300000, 100} // 限价申报上限30万股
	case "00", "60": // 沪市主板
		return []int{1000000, 100} // 主板上限100万股
	default: // 对于未明确匹配的代码，可根据需要返回默认值或视为错误
		return []int{0, 0}
	}
}

// GetTradingMinutes 获取当日已开盘分钟数
func GetTradingMinutes(now time.Time) int {
	// 中国股市交易时间
	amStart := time.Date(now.Year(), now.Month(), now.Day(), 9, 30, 0, 0, time.Local)
	amEnd := time.Date(now.Year(), now.Month(), now.Day(), 11, 30, 0, 0, time.Local)
	pmStart := time.Date(now.Year(), now.Month(), now.Day(), 13, 0, 0, 0, time.Local)
	pmEnd := time.Date(now.Year(), now.Month(), now.Day(), 15, 0, 0, 0, time.Local)

	// 如果当前时间在上午交易时段之前
	if now.Before(amStart) {
		return 0
	}

	// 如果当前时间在上午交易时段
	if now.After(amStart) && now.Before(amEnd) {
		return int(now.Sub(amStart).Minutes()) + 1
	}

	// 如果当前时间在午休时段
	if now.After(amEnd) && now.Before(pmStart) {
		return 120 + 1 // 上午共 120 分钟
	}

	// 如果当前时间在下午交易时段
	if now.After(pmStart) && now.Before(pmEnd) {
		amMinutes := 120
		pmMinutes := int(now.Sub(pmStart).Minutes())
		return amMinutes + pmMinutes + 1
	}

	// 如果当前时间在下午交易时段之后
	if now.After(pmEnd) {
		return 240 + 1 // 全天共 240 分钟
	}

	return 0
}

// IsStockMain 判断股票是否是A股主板（十厘米）
func IsStockMain(symbol string) bool {
	return strings.HasPrefix(symbol, "00") || strings.HasPrefix(symbol, "60")
}

// IsStockA 判断股票是否是A股（ETF、B股等排除）
func IsStockA(symbol string) bool {
	return !strings.HasPrefix(symbol, "1") && !strings.HasPrefix(symbol, "2") && !strings.HasPrefix(symbol, "5") && !strings.HasPrefix(symbol, "7") && !strings.HasPrefix(symbol, "9") && !strings.HasPrefix(symbol, "A")
}
