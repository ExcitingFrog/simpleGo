package stime

import "time"

// WeekStart
func WeekStart(year, week int) time.Time {
	t := time.Date(year, 7, 1, 0, 0, 0, 0, time.Now().Location())

	if wd := t.Weekday(); wd == time.Sunday {
		t = t.AddDate(0, 0, -6)
	} else {
		t = t.AddDate(0, 0, -int(wd)+1)
	}

	_, w := t.ISOWeek()
	t = t.AddDate(0, 0, (week-w)*7)

	return t
}

// GetDateString ...
func GetDateString(start, end string) (sTime, eTime string) {
	weekDict := map[string]int64{
		"Sunday":    6,
		"Monday":    0,
		"Tuesday":   1,
		"Wednesday": 2,
		"Thursday":  3,
		"Friday":    4,
		"Saturday":  5,
	}

	whichDay := time.Now().Weekday().String()

	if start == "" || end == "" {
		now := time.Now().UTC()
		start := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
		sUnix := start.Unix() - 3600*24*weekDict[whichDay]
		sTime = time.Unix(sUnix, 0).String()[0:10]

		end := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, now.Location())
		eTime = end.String()[0:10]
	} else {
		sTime = start
		eTime = end
	}
	return
}

// ParseDate ...
func ParseDate(layout, start, end string) (sUnix, eUnix int64, err error) {

	sTime, eTime := GetDateString(start, end)

	if sTime != "" && eTime != "" {
		_start, err := time.ParseInLocation(layout, sTime, time.Local)
		if err != nil {
			return 0, 0, err
		}
		_end, err := time.ParseInLocation(layout, eTime, time.Local)
		if err != nil {
			return 0, 0, err
		}

		start := time.Date(_start.Year(), _start.Month(), _start.Day(), 0, 0, 0, 0, _start.Location())
		end := time.Date(_end.Year(), _end.Month(), _end.Day(), 23, 59, 59, 0, _end.Location())

		start = start.UTC()
		end = end.UTC()

		sUnix = start.UTC().Unix()
		eUnix = end.UTC().Unix()
	}

	return
}
