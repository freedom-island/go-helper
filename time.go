package main

import (
	"time"
)

// ITimeNowTs 获取当前时间戳
func ITimeNowTs() int64 {
	return time.Now().Unix()
}

// ITimeZeroTs 获取今日零点时间戳
func ITimeZeroTs() int64 {
	return time.Now().Unix() / 86400 * 86400
}

// ITimeZeroCnTs 获取今日零点时间戳（东八时间）
func ITimeZeroCnTs() int64 {
	return time.Now().Unix()/86400*86400 - 8*3600
}

// ITimeArithmeticTs 获取加减运算的时间戳
func ITimeArithmeticTs(change time.Duration) int64 {
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

// IGetNextWorkday 获取下一个工作日
func IGetNextWorkday(dt time.Time) time.Time {
	for {
		dt = dt.AddDate(0, 0, 1)
		if dt.Weekday() != time.Saturday && dt.Weekday() != time.Sunday {
			break
		}
	}

	return dt
}

// IGetPrevWorkday 获取上一个工作日
func IGetPrevWorkday(dt time.Time) time.Time {
	for {
		dt = dt.AddDate(0, 0, -1)
		if dt.Weekday() != time.Saturday && dt.Weekday() != time.Sunday {
			break
		}
	}

	return dt
}
