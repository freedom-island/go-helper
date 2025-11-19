package helper

import (
	"time"
)

// TimeNowTs 获取当前时间戳
func TimeNowTs() int64 {
	return time.Now().Unix()
}

// TimeZeroTs 获取今日零点时间戳
func TimeZeroTs() int64 {
	return time.Now().Unix() / 86400 * 86400
}

// TimeZeroCnTs 获取今日零点时间戳（东八时间）
func TimeZeroCnTs() int64 {
	return time.Now().Unix()/86400*86400 - 8*3600
}

// TimeArithmeticTs 获取加减运算的时间戳
func TimeArithmeticTs(change time.Duration) int64 {
	return time.Now().Add(change).Unix()
}

// IsSameDay 同一天
func IsSameDay(now time.Time, target time.Time) bool {
	return now.Year() == target.Year() && now.YearDay() == target.YearDay()
}

// TodayResidueSec 获取今日剩余秒数
func TodayResidueSec() time.Duration {
	return time.Duration(86400-time.Now().Unix()%86400) * time.Second
}

// DatetimeResidueSec 获取某时间的当日剩余秒数
func DatetimeResidueSec(dt time.Time) time.Duration {
	return time.Duration(86400-dt.Unix()%86400) * time.Second
}

// GetNextWorkday 获取下一个工作日
func GetNextWorkday(dt time.Time) time.Time {
	for {
		dt = dt.AddDate(0, 0, 1)
		if dt.Weekday() != time.Saturday && dt.Weekday() != time.Sunday {
			break
		}
	}

	return dt
}

// GetPrevWorkday 获取上一个工作日
func GetPrevWorkday(dt time.Time) time.Time {
	for {
		dt = dt.AddDate(0, 0, -1)
		if dt.Weekday() != time.Saturday && dt.Weekday() != time.Sunday {
			break
		}
	}

	return dt
}
