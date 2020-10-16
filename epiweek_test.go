package epiweek

import (
	"testing"
	"time"
)

type expected struct {
	year int
	week int
}

func TestEpiweek(t *testing.T) {
	tests := []struct {
		name    string
		epiweek Epiweek
		want    expected
	}{
		{
			name:    "Year starts on Wednesday",
			epiweek: Epiweek{Time: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)},
			want: expected{
				year: 2020,
				week: 1,
			},
		},
		{
			name:    "Week start on Sunday",
			epiweek: Epiweek{Time: time.Date(2019, 12, 1, 0, 0, 0, 0, time.UTC)},
			want: expected{
				year: 2019,
				week: 49,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			y, w := tt.epiweek.Epiweek()
			if y != tt.want.year || w != tt.want.week {
				t.Errorf("Wanted year: %d, Wanted week: %d\nGot year: %d, Got week: %d", tt.want.year, tt.want.week, y, w)
			}
		})
	}
}

func TestAdd(t *testing.T) {
	tests := []struct {
		name    string
		epiweek Epiweek
		add     int
		want    expected
	}{
		{
			name:    "Add positive weeks",
			epiweek: Epiweek{Time: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)},
			add:     2,
			want:    expected{year: 2020, week: 3},
		},
		{
			name:    "Add negative weeks",
			epiweek: Epiweek{Time: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)},
			add:     -2,
			want:    expected{year: 2019, week: 51},
		},
		{
			name:    "Add weeks, cross year, 53 week year",
			epiweek: Epiweek{Time: time.Date(2020, 12, 4, 0, 0, 0, 0, time.UTC)},
			add:     4,
			want:    expected{year: 2020, week: 53},
		},
		{
			name:    "Add weeks, cross year, 52 week year",
			epiweek: Epiweek{Time: time.Date(2019, 12, 24, 0, 0, 0, 0, time.UTC)},
			add:     1,
			want:    expected{year: 2020, week: 1},
		},
		{
			name:    "Cross backwards into 53 year week",
			epiweek: Epiweek{Time: time.Date(2021, 1, 3, 0, 0, 0, 0, time.UTC)},
			add:     -1,
			want:    expected{year: 2020, week: 53},
		},
	}

	for _, tt := range tests {
		e := tt.epiweek.Add(tt.add)
		y, w := e.Epiweek()
		if y != tt.want.year || w != tt.want.week {
			t.Errorf("Wanted year: %d, Wanted week: %d\nGot year: %d, Got week: %d", tt.want.year, tt.want.week, y, w)
		}
	}
}
