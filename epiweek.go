package epiweek

import (
	"fmt"
	"time"
)

const (
	secondsInDay = 24 * 60 * 60
	daysInWeek   = 7
)

type epiType int

const (
	epiweek epiType = iota
	isoweek
)

// Epiweek is initialized with Time and will allow easy operations for
// calculating  CDC Epi weeks
type Epiweek struct {
	time time.Time
	et   epiType
}

// NewEpiweek returns a Epiweek struct based upon CDC definition of an epi week
// with a week beginning on Sunday and ending on Saturday
func NewEpiweek(t time.Time) (e Epiweek) {
	var days time.Duration
	days = time.Duration(myTime(t).daysFromDay(time.Wednesday))

	e = Epiweek{
		time: t.Add(days * time.Second * secondsInDay),
		et:   epiweek,
	}

	return
}

// NewIsoWeek reterns an Epiweek struct based upon ISO standards with a week
// beginning on Monday and ending on Sunday
func NewIsoWeek(t time.Time) (e Epiweek) {
	var days time.Duration
	days = time.Duration(myTime(t).daysFromDay(time.Thursday))

	if days == 4 {
		days = -3
	}

	e = Epiweek{
		time: t.Add(days * time.Second * secondsInDay),
		et:   isoweek,
	}

	return
}

// Epiweek return the year and week number in which Time occurs. Weeks range
// from 1 to 53. Jan 1 to Jan 3  of the year might belong to the prior year.
// Likewise, Dec 29 to Dec 31 might belong to the next year.
func (e Epiweek) Epiweek() (year, week int) {
	// According to the rule that first calendar week in a year is the week
	// that has at least 4 days in that week. (Wednesday for CDC, Thursday for
	// ISO) The day of the week logic was handled in the New function

	year = e.time.Year()
	week = e.time.YearDay()/7 + 1
	return
}

// Add will add the week number of weeks to the e Epiweek
func (e Epiweek) Add(week int) (epiweek Epiweek) {
	epiweek.time = e.time.AddDate(0, 0, daysInWeek*week)

	return
}

func (e Epiweek) String() string {
	year, week := e.Epiweek()
	return fmt.Sprintf("Year [%d], Week [%d]", year, week)
}
