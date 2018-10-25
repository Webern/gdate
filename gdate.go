// gdate is Copyright (c) 2018 by Matthew James Briggs, MIT License

package gdate

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Date struct {
	year  int
	month int
	day   int
}

func NewDate(year int, month int, day int) Date {
	var gd Date
	gd.SetYear(year)
	gd.SetMonth(month)
	gd.SetDay(day)
	return gd
}

func GetMonthInt(m time.Month) int {
	switch m {
	case time.January:
		return 1
	case time.February:
		return 2
	case time.March:
		return 3
	case time.April:
		return 4
	case time.May:
		return 5
	case time.June:
		return 6
	case time.July:
		return 7
	case time.August:
		return 8
	case time.September:
		return 9
	case time.October:
		return 10
	case time.November:
		return 11
	case time.December:
		return 12
	}
	return 1
}

func NewDateFromTime(t time.Time) Date {
	var gd Date
	gd.SetYear(t.Year())
	gd.SetMonth(GetMonthInt(t.Month()))
	gd.SetDay(t.Day())
	return gd
}

func (gd *Date) SetYear(year int) {
	if year < 1 {
		year = 1
	}

	if year > 9999 {
		year = 9999
	}

	utc, _ := time.LoadLocation("UTC")
	t := time.Date(year, 1, 2, 0, 0, 0, 0, utc)
	gd.year = t.Year()
}

func (gd *Date) SetMonth(month int) {
	if month < 1 {
		month = 1
	}

	if month > 12 {
		month = 12
	}

	utc, _ := time.LoadLocation("UTC")
	t := time.Date(2006, time.Month(month), 2, 0, 0, 0, 0, utc)
	m := GetMonthInt(t.Month())
	gd.month = m
}

func (gd *Date) SetDay(day int) {

	if day < 1 {
		day = 1
	}

	if day > 31 {
		day = 31
	}

	utc, _ := time.LoadLocation("UTC")
	m := time.Month(gd.month)

	t := time.Date(gd.year, m, day, 0, 0, 0, 0, utc)
	gd.day = t.Day()
}

func (gd Date) String() string {
	return fmt.Sprintf("%04d-%02d-%02d", gd.year, gd.month, gd.day)
}

func (gd Date) MarshalJSON() ([]byte, error) {
	return []byte("\"" + gd.String() + "\""), nil
}

func (gd *Date) Parse(str string) error {
	yearStr := string(str[0:4])
	monthStr := string(str[5:6])
	dayStr := string(str[8:9])

	y, err := strconv.ParseInt(yearStr, 10, 32)

	if err != nil {
		return err
	}

	m, err := strconv.ParseInt(monthStr, 10, 32)

	if err != nil {
		return err
	}

	d, err := strconv.ParseInt(dayStr, 10, 32)

	if err != nil {
		return err
	}

	gd.SetYear(int(y))
	gd.SetMonth(int(m))
	gd.SetDay(int(d))
	return nil
}

func (gd *Date) UnmarshalJSON(b []byte) error {
	str := string(b)
	str = strings.Replace(str, "\"", "", 0)
	err := gd.Parse(str)
	return err
}

func (gd Date) Year() int {
	return gd.year
}

func (gd Date) Month() int {
	return gd.month
}

func (gd Date) MonthMonth() time.Month {
	return time.Month(gd.month)
}

func (gd Date) Day() int {
	return gd.day
}
