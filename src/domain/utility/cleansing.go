package utility

import "time"

// get first & last date in period
func FirstLastInPeriod(period time.Time) (time.Time, time.Time) {
	y, m, _ := period.Date()
	f := time.Date(y, m, 1, 0, 0, 0, 0, time.UTC)
	l := f.AddDate(0, 1, -1)

	return f, l
}
