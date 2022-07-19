package faker

import "time"

// Weekday -.
func (f *Faker) Weekday() string {
	return Weekday(
		SetRand(f.options.rand),
		SetWeekdays(f.options.weekdays...),
	)
}

// Weekday -.
func Weekday(opts ...Option) string {
	options := setOptions(opts...)
	if len(options.weekdays) == 0 {
		options.SetWeekdays(
			time.Sunday.String(),
			time.Monday.String(),
			time.Tuesday.String(),
			time.Wednesday.String(),
			time.Thursday.String(),
			time.Friday.String(),
			time.Saturday.String(),
		)
	}
	return RandomElement(options.weekdays, opts...)
}

// Month -.
func (f *Faker) Month() string {
	return Month(
		SetRand(f.options.rand),
		SetMonths(f.options.months...),
	)
}

// Month -.
func Month(opts ...Option) string {
	options := setOptions(opts...)
	if len(options.months) == 0 {
		options.SetMonths(
			time.January.String(),
			time.February.String(),
			time.March.String(),
			time.April.String(),
			time.May.String(),
			time.June.String(),
			time.July.String(),
			time.August.String(),
			time.September.String(),
			time.October.String(),
			time.November.String(),
			time.December.String(),
		)
	}
	return RandomElement(options.months, opts...)
}
