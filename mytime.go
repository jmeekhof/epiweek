package epiweek

import "time"

type myTime time.Time

// daysFromDay returns an int that indicates how many days the date is from the
// requested day within the same week.
func (t myTime) daysFromDay(w time.Weekday) (days int) {
	days = int(w - time.Time(t).Weekday())
	return
}
