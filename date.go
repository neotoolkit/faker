package faker

import "time"

// Weekday returns random weekday
func (f *Faker) Weekday() string {
	return Weekday(
		WithRand(f.options.rand),
		WithWeekdays(f.options.weekdays...),
	)
}

// Weekday returns random weekday
//
//    faker.Weekday(
//        faker.WithRand(rand.New(rand.NewSource(time.Now().Unix()))), // Rand instance
//        faker.WithWeekdays(
//            "Sunday",
//            "Monday",
//            "Tuesday",
//            "Wednesday",
//            "Thursday",
//            "Friday",
//            "Saturday",
//        ), // Slice of weekday for RandomElement
//    )
//
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

// Month returns random month
func (f *Faker) Month() string {
	return Month(
		WithRand(f.options.rand),
		WithMonths(f.options.months...),
	)
}

// Month returns random month
//
//	faker.Month(
//		faker.WithRand(rand.New(rand.NewSource(time.Now().Unix()))), // Rand instance
//		faker.WithMonths(
//			"January",
//			"February",
//			"March",
//			"April",
//			"May",
//			"June",
//			"July",
//			"August",
//			"September",
//			"October",
//			"November",
//			"December",
//		), // Slice of month for RandomElement
//	)
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

// Year returns random year between 1970 and current local
func (f *Faker) Year() int {
	return Integer(1970, time.Now().Year(), WithRand(f.options.rand))
}

// Year returns random year between 1970 and current local
//
//    faker.Year(
//        faker.WithRand(rand.New(rand.NewSource(time.Now().Unix()))), // Rand instance
//    )
//
func Year(opts ...Option) int {
	return Integer(1970, time.Now().Year(), opts...)
}
