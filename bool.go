package faker

// Bool returns random bool
func (f *Faker) Bool() bool {
	return Bool(WithRand(f.options.rand))
}

// Bool returns random bool
//
//    faker.Bool(
//        faker.WithRand(rand.New(rand.NewSource(time.Now().Unix()))), // Rand instance
//    )
//
func Bool(opts ...Option) bool {
	return Integer(0, 100, opts...) > 50
}
