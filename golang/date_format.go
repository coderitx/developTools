package tools

import (
	"fmt"
	"time"
)

func Date2Str(_date time.Time, layout string) string {
	return _date.Format(layout)
}

func Date2Timestamp(_date time.Time) int64 {
	return _date.Unix()
}

func Str2Date(dateStr string, layout string) time.Time {
	_time, err := time.Parse(layout, dateStr)
	if err != nil {
		fmt.Println("date str parse failed ", err.Error())
	}
	return _time
}

func Str2Timestamp(dateStr string, layout string) int64 {
	_time, err := time.Parse(layout, dateStr)
	if err != nil {
		fmt.Println("date str parse failed ", err.Error())
	}
	return _time.Unix()
}

func Timestamp2Date(timestamp int64) time.Time {
	return time.Unix(timestamp, 0)
}

func Timestamp2DateStr(timestamp int64, layout string) string {
	return time.Unix(timestamp, 0).Format(layout)
}
