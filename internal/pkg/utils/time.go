package utils

import "time"

func ParseDateRange(startDate, endDate string) (time.Time, time.Time, error) {
	st, err := time.Parse("2006-01-02", startDate)
	if err != nil {
		return time.Now(), time.Now(), err
	}
	ed, err := time.Parse("2006-01-02", endDate)
	ed = time.Date(ed.Year(), ed.Month(), ed.Day(), 23, 59, 59, 1e9-1, ed.Location())
	if err != nil {
		return time.Now(), time.Now(), err
	}
	return st, ed, nil
}

func IsDateRangeValid(startDate, endDate string) bool {
	if startDate == "" || endDate == "" {
		return false
	}
	st, ed, err := ParseDateRange(startDate, endDate)
	if err != nil {
		return false
	}
	return ed.After(st)
}
func GetLastSevenDays() (time.Time, time.Time) {
	now := time.Now().UTC()
	startDate := now.AddDate(0, 0, -7).Truncate(time.Hour * 24)
	endDate := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 1e9-1, now.Location())

	return startDate, endDate
}
