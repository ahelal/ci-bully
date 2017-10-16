package main

import (
	"testing"
	"time"
)

func TestWorkdaysBetweenDates(t *testing.T) {
	t.Parallel()

	tables := []struct {
		day1     time.Time
		day2     time.Time
		workdays int
	}{
		{
			time.Date(2017, 10, 1, 0, 0, 0, 0, time.UTC),
			time.Date(2017, 10, 31, 0, 0, 0, 0, time.UTC),
			21,
		},
		{
			time.Date(2017, 10, 31, 0, 0, 0, 0, time.UTC),
			time.Date(2017, 10, 1, 0, 0, 0, 0, time.UTC),
			21,
		},
		{
			time.Date(2017, 10, 1, 0, 0, 0, 0, time.UTC),
			time.Date(2017, 11, 1, 0, 0, 0, 0, time.UTC),
			22,
		},
		{
			time.Date(2017, 10, 8, 0, 0, 0, 0, time.UTC),
			time.Date(2017, 10, 13, 0, 0, 0, 0, time.UTC),
			4,
		},
		{
			time.Date(2017, 10, 13, 0, 0, 0, 0, time.UTC),
			time.Date(2017, 10, 14, 0, 0, 0, 0, time.UTC),
			1,
		},
		{
			time.Date(2017, 10, 13, 0, 0, 0, 0, time.UTC),
			time.Date(2017, 10, 13, 0, 0, 0, 0, time.UTC),
			0,
		},
		{
			time.Date(2017, 10, 14, 0, 0, 0, 0, time.UTC),
			time.Date(2017, 10, 14, 0, 0, 0, 0, time.UTC),
			0,
		},
		{
			time.Date(2017, 10, 14, 0, 0, 0, 0, time.UTC),
			time.Date(2017, 10, 15, 0, 0, 0, 0, time.UTC),
			0,
		},
		{
			time.Date(2017, 10, 10, 0, 0, 0, 0, time.UTC),
			time.Date(2017, 10, 12, 0, 0, 0, 0, time.UTC),
			2,
		},
		{
			time.Date(2017, 12, 10, 0, 0, 0, 0, time.UTC),
			time.Date(2018, 1, 10, 0, 0, 0, 0, time.UTC),
			22,
		},
		{
			time.Date(2017, 10, 9, 0, 0, 0, 0, time.UTC),
			time.Date(2017, 10, 15, 0, 0, 0, 0, time.UTC),
			5,
		},
	}

	for _, table := range tables {
		workdays := workdaysBetweenDates(table.day1, table.day2)
		if workdays != table.workdays {
			t.Errorf("Workdays count between %s and %s are incorrect, got: %d, want: %d.",
				table.day1, table.day2, workdays, table.workdays)
		}
	}
}
