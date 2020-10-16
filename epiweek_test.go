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
