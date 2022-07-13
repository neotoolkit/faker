package faker_test

import (
	"math"
	"testing"

	"github.com/neotoolkit/faker"
)

func TestUnix(t *testing.T) {
	u := faker.Unix()

	if u < 0 || u > math.MaxInt64-1 {
		t.Errorf("unix must be equal 0 or less or equal %v", math.MaxInt64-1)
	}
}

func TestWeekday(t *testing.T) {
	weekday := faker.Weekday()

	m := map[string]struct{}{
		"Sunday":    {},
		"Monday":    {},
		"Tuesday":   {},
		"Wednesday": {},
		"Thursday":  {},
		"Friday":    {},
		"Saturday":  {},
	}

	if _, ok := m[weekday]; !ok {
		t.Error("bad weekday")
	}
}

func TestMonth(t *testing.T) {
	month := faker.Month()

	m := map[string]struct{}{
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

	if _, ok := m[month]; !ok {
		t.Error("bad month")
	}
}
