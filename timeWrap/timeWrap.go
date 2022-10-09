package timeWrap

import (
	"time"
)

func Timestamp() int64 {
	return time.Now().Unix()
}

func Date(timestamp int64, args ...interface{}) string {
	layout := "2006-01-02 15:04:05"
	if args != nil {
		layout = args[0].(string)
	}
	return time.Unix(timestamp, 0).Format(layout)
}

func Today(args ...interface{}) string {
	layout := "2006-01-02 15:04:05"
	if args != nil {
		layout = args[0].(string)
	}
	return time.Now().Format(layout)
}

func ToTimestamp(format, date string) int64 {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	t, _ := time.ParseInLocation(format, date, loc)
	return t.Unix()
}

func MonthDays(year, month int) int {
	days := 0
	switch month {
	case 4:
		fallthrough
	case 6:
		fallthrough
	case 9:
		fallthrough
	case 11:
		days = 30
	case 2:
		if ((year%4) == 0 && (year%100) != 0) || (year%400) == 0 {
			days = 29
		} else {
			days = 28
		}
	default:
		days = 31
	}

	return days
}

func DateAddDuration(day string, duration time.Duration, args ...string) string {
	format := "2006-01-02"
	if len(args) > 0 {
		format = args[0]
	}

	t, _ := time.Parse(format, day)
	return t.Add(duration).Format(format)
}

func AddMonths(day string, months int, args ...string) string {
	format := "2006-01-02"
	if len(args) > 0 {
		format = args[0]
	}

	t, _ := time.Parse(format, day)
	return t.AddDate(0, months, 0).Format(format)
}

func AddDays(day string, days int, args ...string) string {
	format := "2006-01-02"
	if len(args) > 0 {
		format = args[0]
	}

	t, _ := time.Parse(format, day)
	return t.AddDate(0, 0, days).Format(format)
}

func DateDays(startDate, endDate, format string) int {
	startTime, _ := time.Parse(format, startDate)
	endTime, _ := time.Parse(format, endDate)
	return int(endTime.Sub(startTime).Hours()/24) + 1
}

func DayStartTimestamp(timestamp int64) int64 {
	s := time.Unix(timestamp, 0).Format("2006/01/02")
	return ToTimestamp("2006/01/02", s)
}
