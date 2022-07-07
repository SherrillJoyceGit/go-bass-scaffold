package util

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"math"
	"time"
)

//日期格式
const (
	DATE                      = "2006-01-02"
	DATE_MONTH                = "2006-01"
	DATE_TIME                 = "2006-01-02 15:04:05"
	TIME                      = "15:04:05"
	TIME_WITHOUT_SECOND       = "15:04"
	DATE_TIME_WITHOUT_SECONDS = "2006-01-02 15:04"
	//修改时间的类型
	types = "s" //秒
	typem = "m" //分
	typeh = "h" //时
	typed = "d" //日
	typeM = "M" //月
	typey = "y" //年

	YYYYMMDDHHMISS = "2006-01-02 15:04:05" //常规类型
)

type JSONTime struct {
	time.Time
}

//格式化日期为字符串
func FormatDate(date time.Time, pattern string) (error, string) {
	if date.IsZero() {
		//return nil, ""
		return errors.New("传入的日期为空"), ""
	}
	return nil, date.Format(pattern)
}

//解析字符串日期为日期
func ParseDate(date string, pattern string) (error, time.Time) {
	loc := time.Now().Location()
	theTime, err := time.ParseInLocation(pattern, date, loc)
	if err != nil {
		return err, time.Now()
	}
	return nil, theTime
}

//为Date增减时间(时h、分m、秒s） (增加传正数，减少传负数）
func AddTime(date time.Time, plus int, t string) time.Time {
	add := fmt.Sprintf("%d"+t, plus)
	d, _ := time.ParseDuration(add)
	newDate := date.Add(d)
	return newDate
}

//日期加减类型为（日期：d;月份：M；年份：y)
func AddDate(date time.Time, plus int, t string) (error, time.Time) {
	newDate := time.Time{}
	switch t {
	case "d":
		newDate = date.AddDate(0, 0, plus)
	case "M":
		newDate = date.AddDate(0, plus, 0)
	case "y":
		newDate = date.AddDate(plus, 0, 0)
	default:
		return errors.New("选择的类型错误，应该为：\n日期：d;\n月份：M；\n年份：y"), time.Time{}
	}
	return nil, newDate
}

//获取两个时间差几天,向下取整
func GetSubDay(front time.Time, latest time.Time) int64 {
	//<24小时--不到一天，返回0
	//>=24小时--返回整数部分
	hour := latest.Sub(front).Hours()
	if hour < 24 {
		return 0
	} else {
		d := hour / 24
		d = math.Floor(d)
		return int64(d)
	}
}

//返回某天起始时间
func GetStartTime(date time.Time) time.Time {
	return time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
}

//返回某天结束时间
func GetEndTime(date time.Time) time.Time {
	return time.Date(date.Year(), date.Month(), date.Day(), 23, 59, 59, 999, date.Location())
}

//返回某个月的第一天
func GetFirstDay(date time.Time) time.Time {
	date = date.AddDate(0, 0, -date.Day()+1)
	return GetStartTime(date)
}

//返回某个月的最后一天
func GetLastDay(date time.Time) time.Time {
	return GetFirstDay(date).AddDate(0, 1, -1)
}

// MarshalJSON on JSONTime format Time field with %Y-%m-%d %H:%M:%S
func (t JSONTime) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%s\"", t.Format(YYYYMMDDHHMISS))
	return []byte(formatted), nil
}

// Value insert timestamp into mysql need this function.
func (t JSONTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	format := t.Time.Format("2006-01-02 15:04:05")
	return format, nil
	//return t.Time, nil
}

// Scan valueof time.Time
func (t *JSONTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = JSONTime{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

// UnmarshalJSON implements the json.Unmarshaler interface.
// The time is expected to be a quoted string in RFC 3339 format.
func (t *JSONTime) UnmarshalJSON(data []byte) error {
	// Ignore null, like in the main JSON package.
	if string(data) == "null" {
		return nil
	}
	// Fractional seconds are handled implicitly by Parse.
	var err error
	(*t).Time, err = time.ParseInLocation(`"`+YYYYMMDDHHMISS+`"`, string(data), time.Local)
	return err
}
