package main

import (
	"fmt"
	"github.com/jmeekhof/epiweek"
	"time"
)

func main() {
	//e := epiweek.Epiweek{Time: time.Now()}
	days := []epiweek.Epiweek{
		epiweek.NewEpiweek(date(2019, 12, 31)),
		epiweek.NewEpiweek(date(2020, 10, 15)),
		epiweek.NewEpiweek(date(2020, 10, 16)),
		epiweek.NewEpiweek(date(2020, 10, 17)),
		epiweek.NewEpiweek(date(2020, 10, 18)),
		epiweek.NewEpiweek(date(2020, 10, 19)),
		epiweek.NewEpiweek(date(2020, 10, 20)),
		epiweek.NewEpiweek(date(2020, 10, 21)),
		epiweek.NewEpiweek(date(2020, 10, 22)),
	}
	for _, e := range days {
		fmt.Printf("Epiweek: %s\n", e)
	}
}

func date(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}
