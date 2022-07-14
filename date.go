package faker

import (
	"math"
	"time"
)

func Unix(opts ...UnixOption) int64 {
	options := UnixOptions{
		min: 0,
		max: math.MaxInt64,
	}
	for _, opt := range opts {
		opt(&options)
	}
	return IntBetweenInt64(options.min, options.max)
}

type UnixOption func(opts *UnixOptions)

type UnixOptions struct {
	min int64
	max int64
}

func SetUnixMax(max int64) UnixOption {
	return func(opts *UnixOptions) {
		opts.max = max
	}
}

func SetUnixMin(min int64) UnixOption {
	return func(opts *UnixOptions) {
		opts.min = min
	}
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
