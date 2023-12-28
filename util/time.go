package util

import (
	"time"
)

// getNowOjb 获取本地当前时间的time对象
func getNowOjb() time.Time {
	location, _ := time.LoadLocation("Asia/Shanghai")
	return time.Now().In(location)
}

// GetTimeStamp 获取本地当前时间戳(秒)
func GetTimeStamp() int64 {
	nowTime := getNowOjb()
	return nowTime.Unix()
}

// GetMilliTimeStamp 获取本地当前时间戳（毫秒）
func GetMilliTimeStamp() int64 {
	nowTime := getNowOjb()
	return nowTime.UnixMilli()
}

// GetDateTime 获取本地当前时间 DateTime
func GetDateTime() string {
	nowTime := getNowOjb()
	return nowTime.Format(time.DateTime)
}

// GetDate 获取本地当前日期 Date
func GetDate() string {
	nowTime := getNowOjb()
	return nowTime.Format(time.DateOnly)
}

// GetTime 获取本地当前时间 Time
func GetTime() string {
	nowTime := getNowOjb()
	return nowTime.Format(time.TimeOnly)
}

// DateToTime 日期转时间戳（秒）date
func DateToTime(dateTime string) int64 {
	timeStamp, err := time.Parse(time.DateOnly, dateTime)
	CheckErr(err)

	return timeStamp.Unix()
}

// DateToMilliTime 日期转时间戳 (毫秒) date
func DateToMilliTime(dateTime string) int64 {
	timeStamp, err := time.Parse(time.DateOnly, dateTime)
	CheckErr(err)

	return timeStamp.UnixMilli()
}

// DateTimeToTime 日期转时间戳（秒）dateTime
func DateTimeToTime(dateTime string) int64 {
	timeStamp, err := time.Parse(time.DateTime, dateTime)
	CheckErr(err)

	return timeStamp.Unix()
}

// DateTimeToMilliTime 日期转时间戳 (毫秒) dateTime
func DateTimeToMilliTime(dateTime string) int64 {
	timeStamp, err := time.Parse(time.DateTime, dateTime)
	CheckErr(err)

	return timeStamp.UnixMilli()
}

// GtCalculateTime 当前时间 + d 时间
//
// d: 有效单位：ns, us，ms, s, m, h
// 例如："300ms", "-1.5h", "2h45m"
func GtCalculateTime(currencyTimer time.Time, d string) (time.Time, error) {
	duration, err := time.ParseDuration(d)
	if err != nil {
		return time.Time{}, err
	}

	return currencyTimer.Add(duration), nil
}

// DiffDate 两个时间相差天数 date
func DiffDate(date1, date2 string) int {
	layout := time.DateOnly
	start, _ := time.Parse(layout, date1)
	end, _ := time.Parse(layout, date2)

	duration := end.Sub(start)
	days := int(duration.Hours() / 24)

	return days
}

// DiffDateTime 两个时间相差天数 date
func DiffDateTime(dateTime1, dateTime2 string) int {
	layout := time.DateTime
	start, _ := time.Parse(layout, dateTime1)
	end, _ := time.Parse(layout, dateTime2)

	duration := end.Sub(start)
	days := int(duration.Hours() / 24)

	return days
}

// GetTimeDuration 获取d秒时间
func GetTimeDuration(d int64) time.Duration {
	return time.Duration(d) * time.Second
}
