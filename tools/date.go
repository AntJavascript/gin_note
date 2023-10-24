package tools

import (
	"fmt"
	"time"
)

// 日期格式转换
func formatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%d-%d", year, month, day)
}

// string 时间转化为Time
func StringAsTime(date string) time.Time {
	t, _ := time.Parse("2006-01-02", date)
	return t
}

// Time 时间转化为string
func TimeAsString(t time.Time) string {
	timeString := t.Format("2006-01-02 00:00:00")
	return timeString
}

// 时间戳转Time
func TimesAsTampTime(timestamp int64) time.Time {
	timeNow := time.Unix(timestamp, 0)
	return timeNow
}

// 获取月份最大天数
func GetMonthDaxDay(year, month string) int {
	t := StringAsTime(year + "-" + month + "-" + "01")
	lastTime := t.AddDate(0, 1, -1)
	_, _, lastDay := lastTime.Date()
	return lastDay
}

// string 转 时间戳
func StringAsTampTime(time string) int64 {
	the_time := StringAsTime(time)
	return the_time.Unix()
}
