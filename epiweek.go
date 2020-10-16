package epiweek

import (
	"fmt"
	"time"
)

const (
	secondsInDay = 24 * 60 * 60
)

type Epiweek struct {
	Time time.Time
}

func (e Epiweek) daysFromDay(w time.Weekday) (days int) {
	days = int(w - e.Time.Weekday())
	return
}

// Epiweek return the year and week number in which Time occurs. Weeks range
// from 1 to 53. Jan 1 to Jan 3  of the year might belong to the prior year.
// Likewise, Dec 29 to Dec 31 might belong to the next year.
func (e Epiweek) Epiweek() (year, week int) {
	// According to the rule that first calendar week in a year is the week
	// that has at least 4 days in that week. (Wednesday)
	var days time.Duration
	days = time.Duration(e.daysFromDay(time.Wednesday))

	// find the Wednesday of the week
	wed := e.Time.Add(days * time.Second * secondsInDay)
	year = wed.Year()
	week = wed.YearDay()/7 + 1
	return
}

func (e Epiweek) String() string {
	year, week := e.Epiweek()
	return fmt.Sprintf("Year [%d], Week [%d]", year, week)
}
