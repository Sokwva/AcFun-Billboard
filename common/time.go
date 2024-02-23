package common

import (
	"time"
)

func getMonthStartDate() time.Time {
	var d time.Time = time.Now()
	d.AddDate(0, 0, -d.Day())
	return d
}

func getMonthEndDate() time.Time {
	return getMonthStartDate().AddDate(0, 1, -1)
}

func getTimeStrTodayToMonthBegin() (string, string) {
	var now time.Time = time.Now()
	today := now.Format("2006-01-02")
	firstDay := getMonthStartDate().Format("2006-01-02")
	return firstDay, today
}
