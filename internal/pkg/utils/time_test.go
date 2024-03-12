package utils

import (
	"errors"
	"testing"
	"time"
)

func TestIsDateRangeValid(t *testing.T) {
	type testCase struct {
		// Inputs
		startDate, endDate string
		// expected value
		want bool
		// Got value
		got bool
	}
	scenarios := []testCase{
		{
			"2024-01-01", "2024-01-01",
			true,
			true,
		},
		{
			"2024-01-01", "2023-01-01",
			false,
			false,
		},
		{
			"", "",
			false,
			false,
		},
	}
	for id, scenario := range scenarios {
		scenario.got = IsDateRangeValid(scenario.startDate, scenario.endDate)
		if scenario.got != scenario.want {
			t.Errorf("ID: %v .You've got: %v and you expected: %v", id, scenario.got, scenario.want)
		}
	}
}

func TestGetLastSevenDays(t *testing.T) {

	gotSDate, gotEDate := GetLastSevenDays()

	if gotSDate.IsZero() {
		t.Errorf("You've got: %v", gotSDate)
	}
	if gotEDate.IsZero() {
		t.Errorf("You've got: %v", gotEDate)
	}
}

func TestParseDateRange(t *testing.T) {
	type testCase struct {
		// Inputs
		startDate, endDate string
		// Expected values
		wantSdate, wantEdate time.Time
		wantErr              error
	}
	scenarios := []testCase{
		{
			"", "",
			time.Now(), time.Now(),
			errors.New("Parsing time"),
		},
	}
	for _, scenario := range scenarios {
		gotSdate, gotEdate, gotErr := ParseDateRange(scenario.startDate, scenario.endDate)
		if gotErr == nil {
			if gotSdate != scenario.wantSdate {
				t.Errorf("Parsing time %v", scenario.wantSdate)
			}
			if gotEdate != scenario.wantEdate {
				t.Errorf("Parsing time %v", scenario.wantEdate)
			}
		}
	}
}
