package faker_test

import (
	"testing"

	"neotoolkit.com/faker"
)

func TestFaker_Weekday(t *testing.T) {
	t.Parallel()
	f := faker.New(faker.SetWeekdays("test"))
	weekday := f.Weekday()
	if weekday != "test" {
		t.Errorf("got %s, want test", weekday)
	}
}

func TestWeekday(t *testing.T) {
	t.Parallel()
	weekdays := map[string]struct{}{
		"Sunday":    {},
		"Monday":    {},
		"Tuesday":   {},
		"Wednesday": {},
		"Thursday":  {},
		"Friday":    {},
		"Saturday":  {},
	}
	weekday := faker.Weekday()
	if _, ok := weekdays[weekday]; !ok {
		t.Error("bad weekday")
	}
}

func TestFaker_Month(t *testing.T) {
	t.Parallel()
	f := faker.New(faker.SetMonths("test"))
	month := f.Month()
	if month != "test" {
		t.Errorf("got %s, want test", month)
	}
}

func TestMonth(t *testing.T) {
	t.Parallel()
	months := map[string]struct{}{
		"January":   {},
		"February":  {},
		"March":     {},
		"April":     {},
		"May":       {},
		"June":      {},
		"July":      {},
		"August":    {},
		"September": {},
		"October":   {},
		"November":  {},
		"December":  {},
	}
	month := faker.Month()
	if _, ok := months[month]; !ok {
		t.Error("bad month")
	}
}
