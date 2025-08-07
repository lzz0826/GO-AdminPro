package utils

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
)

const (
	TimeBarFormat          = "2006-01-02 15:04:05"        // 标准的完整时间格式，包含年月日和时分秒，适用于大多数场景。
	TimeMinuteBarFormat    = "2006-01-02 15:04"           // 只包含年月日、小时和分钟的时间格式，适合场景如简化时间显示。
	TimeBarFormatPM        = "2006-01-02 15:04:05 PM"     // 12小时制时间格式，包含 AM/PM 标识，适用于特定场合的时间显示。
	TimeFormatHMS          = "20060102150405"             // 紧凑型时间格式，不包含任何分隔符，通常用于文件命名或生成时间戳。
	TimeUnderlineYearMonth = "2006_01"                    // 年月之间用下划线分隔的格式，适合文件名或按年月分类的日志文件。
	TimeFormatHM           = "2006-01"                    // 只包含年和月的格式，适合按月份处理的报表或统计数据。
	TimeBarYYMMDD          = "2006-01-02"                 // 只包含年月日的格式，用于简单的日期显示。
	TimeHHMMSS             = "15:04:05"                   // 只包含时分秒的时间格式，用于单独表示时间（不包含日期）。
	TimeYYMMDD             = "20060102"                   // 紧凑的年月日格式，用于不带分隔符的日期表示，常用于存储或标识符。
	TimeYYMM               = "200601"                     // 紧凑的年月格式，用于不带分隔符的年月表示。
	TimeYMDHM              = "200601021504"               // 年月日小时分钟格式，用于需要精确到分钟的时间表示。
	TimeBEIJINGFormat      = "2006-01-02 15:04:05 +08:00" // 北京时间的标准格式，包含时区信息，适用于全球化场景。
	TimeGDFormat           = "01/02/2006 15:04:05"        // MM/DD/YYYY 格式的时间表示，适用于不同的日期显示需求。
	TimeTFormat            = "2006-01-02T15:04:05"        // ISO 8601 格式的时间表示，带有 "T" 和时区信息，适合与外部系统交互时使用。
	TimeTBjFormat          = "2006-01-02T15:04:05+08:00"  // 带有北京时区的 ISO 8601 格式，适合中国时区的数据处理。

	Minute   = 60
	HourVal  = Minute * 60
	DayVal   = HourVal * 24
	MonthVal = DayVal * 30
	YearVal  = MonthVal * 365

	BeiJinAreaTime = "Asia/Shanghai"
)

func GetNowEsTime() string {
	return time.Now().Format(TimeTBjFormat)
}

// 返回指定格式化的日期
func GetNowTimeV2(format string) string {
	if format == "" {
		format = TimeBarFormat
	}
	return time.Now().Format(format)
}

// 返回13位时间戳 字符串
func GetMicroSecond() string {
	return fmt.Sprintf("%v", time.Now().UnixNano()/1e6)
}

// 接收一個 time.Time，返回 13 位時間戳（毫秒）字串
func GetMilliTimestamp(t time.Time) string {
	return fmt.Sprintf("%d", t.UnixNano()/1e6)
}

// 接收一個 13 位毫秒時間戳的字串，返回 time.Time 和錯誤
func ParseMilliTimestamp(millisStr string) (time.Time, error) {
	millis, err := strconv.ParseInt(millisStr, 10, 64)
	if err != nil {
		return time.Time{}, err
	}

	seconds := millis / 1000
	nanoseconds := (millis % 1000) * 1e6
	return time.Unix(seconds, nanoseconds), nil
}

// 当前时间增加 n 小时
func NowAddHours(n int) string {
	now := time.Now()
	h, _ := time.ParseDuration(fmt.Sprint(n) + "h")
	hh1 := now.Add(h)
	return hh1.Format(TimeBarFormat)
}

// GetESTimeFomat return 2019-01-14T19:00:33+08:00
func GetESTimeFomat(timestr string) string {
	return fmt.Sprintf("%s+08:00", strings.Replace(strings.TrimSpace(timestr), " ", "T", -1))
}

func GetBeginOfMonth(t time.Time) string {
	return time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location()).Format(TimeBarFormat)
}
func GetEndOfMonth(t time.Time) string {
	return time.Date(t.Year(), t.Month()+1, 0, 0, 0, -1, 0, t.Location()).Format(TimeBarFormat)
}
func GetCurrDate() string {
	return time.Now().Format(TimeBarYYMMDD)
}
func GetCurrMonth() string {
	return time.Now().Format(TimeFormatHM)
}
func GetCurrDateTime() string {
	return time.Now().Format(TimeBarFormat)
}

// DiffMonth 计算 t1 和 t2 之间的月份差
func DiffMonth(t1, t2 time.Time) (month int) {
	return int(t2.Month()-t1.Month()) + 12*(t2.Year()-t1.Year())
}

// DiffMonthAndDay 计算 t1 和 t2 之间的月份差，考虑日期
func DiffMonthAndDay(t1, t2 time.Time) int {
	// 计算年和月的差异
	monthDiff := int(t2.Month()-t1.Month()) + 12*(t2.Year()-t1.Year())
	// 如果 t1 的日期大于 t2 的日期，则月份差减一
	if t2.Day() < t1.Day() {
		monthDiff--
	}
	return monthDiff
}

// 将任意时间转化为另一种时间字符串
func GetTimeFomat(value, layout string) string {
	return StrToTime(value).Format(layout)
}

func StrToTime(value string) time.Time {
	if value == "" {
		return time.Time{}
	}
	layouts := []string{
		"2006-01-02 15:04:05 -0700 MST",
		"2006-01-02 15:04:05 -0700",
		"2006-01-02 15:04:05",
		"2006/01/02 15:04:05 -0700 MST",
		"2006/01/02 15:04:05 -0700",
		"2006/01/02 15:04:05",
		"2006-01-02 -0700 MST",
		"2006-01-02 -0700",
		"2006-01-02",
		"2006/01/02 -0700 MST",
		"2006/01/02 -0700",
		"2006/01/02",
		"2006-01-02 15:04:05 -0700 -0700",
		"2006/01/02 15:04:05 -0700 -0700",
		"2006-01-02 -0700 -0700",
		"2006/01/02 -0700 -0700",
		"2006-01",
		time.ANSIC,
		time.UnixDate,
		time.RubyDate,
		time.RFC822,
		time.RFC822Z,
		time.RFC850,
		time.RFC1123,
		time.RFC1123Z,
		time.RFC3339,
		time.RFC3339Nano,
		time.Kitchen,
		time.Stamp,
		time.StampMilli,
		time.StampMicro,
		time.StampNano,
	}

	var t time.Time
	var err error
	for _, layout := range layouts {
		t, err = time.Parse(layout, value)
		if err == nil {
			return t
		}
	}

	return t
}

func CountDay(year int, month int) (days int) {
	if month != 2 {
		if month == 4 || month == 6 || month == 9 || month == 11 {
			days = 30

		} else {
			days = 31
		}
	} else {
		if ((year%4) == 0 && (year%100) != 0) || ((year % 400) == 0) {
			days = 29
		} else {
			days = 28
		}
	}

	return
}

// 判断当前时间是否在指定时间之内
func CurrentTimeBetween(timeStart, timeEnd string) (bool, error) {
	start, err := BjTBarFmtTimeFormat(timeStart, TimeBarFormat)
	if err != nil {
		return false, err
	}
	end, err := BjTBarFmtTimeFormat(timeEnd, TimeBarFormat)
	if err != nil {
		return false, err
	}
	now := BJNowTime()
	if now.Before(start) || now.After(end) {
		return false, nil
	}
	return true, nil
}

// 获取两个时间相差的天数，0表同一天，正数表t1>t2，负数表t1<t2
func GetDiffDays(t1, t2 time.Time) int {
	t1 = time.Date(t1.Year(), t1.Month(), t1.Day(), 0, 0, 0, 0, time.Local)
	t2 = time.Date(t2.Year(), t2.Month(), t2.Day(), 0, 0, 0, 0, time.Local)

	return int(t1.Sub(t2).Hours() / 24)
}

// 比较指定的日期跟当前时间比对 当前时间大返回true
func CompareDate(endAt string) bool {
	if endAt == "" {
		return true
	}
	startTime := BJNowTime()
	endTime, _ := BjTBarFmtTime(endAt)
	res := DiffMonthAndDay(startTime, endTime)
	if res > 0 {
		return true
	}
	return false
}

// GetTimeInterval 根据输入时间字符串和时区信息，计算指定时间与当前时间的时间间隔 描述信息。
func GetTimeInterval(timeStr string, location *time.Location) string {
	if timeStr == "" {
		return ""
	}
	bjTime, err := time.ParseInLocation(TimeTBjFormat, timeStr, location)
	if err != nil {
		return "30分钟前"
	}
	interval := GetBjNowTime().Unix() - bjTime.Unix()
	if interval < 60 {
		return "刚刚"
	}
	if interval/Minute > 0 && interval/Minute < Minute {
		return fmt.Sprintf("%v分钟前", interval/(Minute))
	} else if interval/HourVal > 0 && interval/HourVal < 24 {
		return fmt.Sprintf("%v小时前", interval/HourVal)
	} else if interval/DayVal > 0 && interval/DayVal < 30 {
		return fmt.Sprintf("%v天前", interval/DayVal)
	} else if interval/MonthVal > 0 && interval/MonthVal < 12 {
		return fmt.Sprintf("%v月前", interval/MonthVal)
	} else if interval/YearVal > 0 {
		return fmt.Sprintf("%v年前", interval/YearVal)
	}
	return "刚刚"
}

// 检查结束时间是否大于开始时间
func CheckEndDateGTStartDate(startDate, endDate string) error {
	if startDate == "" || endDate == "" {
		return fmt.Errorf("开始时间或结束时间为空: startDate=%s, endDate=%s", startDate, endDate)
	}
	beforeTimeStamp, err := time.Parse("2006-01-02 15:04:05", startDate)
	if err != nil {
		return fmt.Errorf("查询起始时间格式非法: %s", startDate)
	}
	afterTimeStamp, err := time.Parse("2006-01-02 15:04:05", endDate)
	if err != nil {
		return fmt.Errorf("查询结束时间格式非法: %s", endDate)
	}
	if afterTimeStamp.Before(beforeTimeStamp) {
		return fmt.Errorf("查询起始时间不能大于查询结束时间: startDate=%s, endDate=%s", startDate, endDate)
	}

	return nil
}
func GetNowTimeStr() string {
	return GetBjNowTime().Format(TimeBarFormat)
}

// 比较两个时间，判断开始日期是否小于结束日期，如果小于返回true，注意：start 和 end 类型需一致
func CompareBothTime(start, end interface{}) bool {
	var (
		t1 time.Time
		t2 time.Time
		e1 error
		e2 error
	)
	// 类型不一样时返回false
	if reflect.TypeOf(start).Kind() != reflect.TypeOf(end).Kind() {
		return false
	}

	switch start.(type) {
	// 时间类型的对比
	case time.Time:
		t1, t2 = start.(time.Time), end.(time.Time)
	case string:
		startValue, endValue := start.(string), end.(string)
		if startValue == "" || endValue == "" {
			return false
		}
		t1, e1 = time.Parse(TimeBarFormat, startValue)
		t2, e2 = time.Parse(TimeBarFormat, endValue)
		if e1 != nil || e2 != nil {
			return false
		}
	default:
		return false
	}
	return t1.Before(t2)
}

/*
北京 时区
*/

// 取得时区 北京
func GetBjTimeLoc() *time.Location {
	// 获取北京时间, 在 windows系统上 time.LoadLocation 会加载失败, 最好的办法是用 time.FixedZone
	var bjLoc *time.Location
	var err error
	bjLoc, err = time.LoadLocation(BeiJinAreaTime)
	if err != nil {
		bjLoc = time.FixedZone("CST", 8*3600)
	}
	return bjLoc
}

// BJNowTime 北京当前时间
func BJNowTime() time.Time {
	// 获取北京时间, 在 windows系统上 time.LoadLocation 会加载失败, 最好的办法是用 time.FixedZone, es 中的时间为: "2019-03-01T21:33:18+08:00"
	var beiJinLocation *time.Location
	var err error

	beiJinLocation, err = time.LoadLocation("Asia/Shanghai")
	if err != nil {
		beiJinLocation = time.FixedZone("CST", 8*3600)
	}

	nowTime := time.Now().In(beiJinLocation)

	return nowTime
}

// GetBjNowTime 获取北京时间,
func GetBjNowTime() time.Time {
	// 获取北京时间, 在 windows系统上 time.LoadLocation 会加载失败, 最好的办法是用 time.FixedZone
	var bjLoc *time.Location
	var err error
	bjLoc, err = time.LoadLocation(BeiJinAreaTime)
	if err != nil {
		bjLoc = time.FixedZone("CST", 8*3600)
	}

	return time.Now().In(bjLoc)
}

// 将北京时间 2006-01-02 15:04:05 类型的时间转换为时间
func BjTBarFmtTime(timeStr string) (time.Time, error) {
	if timeStr == "" {
		return time.Time{}, errors.New("time is empty")
	}

	bjTimeLoc := GetBjTimeLoc()
	format := TimeBarFormat
	if len(timeStr) == 10 {
		format = TimeBarYYMMDD
	}
	return time.ParseInLocation(format, timeStr, bjTimeLoc)
}

// 将时间戳转换为北京时间
func FmtUnixToBjTime(timestamp int64) time.Time {
	bjTimeLoc := GetBjTimeLoc()

	utcTime := time.Unix(timestamp, 0)
	return utcTime.In(bjTimeLoc)
}

func GetTodayEndTime() time.Time {
	t := GetBjNowTime()
	return time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 0, t.Location())
}

// locTIMESEND
func LocUnix() int64 {
	location, _ := time.LoadLocation(BeiJinAreaTime)
	return time.Now().In(location).Unix()
}

// 将北京时间 2006-01-02 15:04:05 类型的时间转换为时间
func BjTBarFmtTimeFormat(timeStr string, timeFormat string) (time.Time, error) {
	if timeStr == "" {
		return time.Time{}, errors.New("time is empty")
	}
	bjTimeLoc := GetBjTimeLoc()
	return time.ParseInLocation(timeFormat, timeStr, bjTimeLoc)
}

// 获取指定时间所在月的第一天 北京
func FirstDayOfMonth(t time.Time) time.Time {
	t = t.In(time.FixedZone("shanghai/beijing", 8*3600))
	y, m, _ := t.Date()
	firstOfMonth := time.Date(y, m, 1, 0, 0, 0, 0, time.FixedZone("shanghai/beijing", 8*3600))
	return firstOfMonth
}
