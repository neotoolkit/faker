package faker

import (
	"math"
	"time"
)

func Unix() int64 {
	return IntBetweenInt64(0, math.MaxInt64)
}

func Weekday() string {
	weekdays := []time.Weekday{
		time.Sunday,
		time.Monday,
		time.Tuesday,
		time.Wednesday,
		time.Thursday,
		time.Friday,
		time.Saturday,
	}
	return weekdays[IntBetween(0, len(weekdays)-1)].String()
}

func Month() string {
	months := []time.Month{
		time.January,
		time.February,
		time.March,
		time.April,
		time.May,
		time.June,
		time.July,
		time.August,
		time.September,
		time.October,
		time.November,
		time.December,
	}
	return months[IntBetween(0, len(months)-1)].String()
}
