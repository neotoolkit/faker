package faker

// Bool returns random bool
func (f *Faker) Bool() bool {
	return Bool(SetRand(f.options.rand))
}

// Bool returns random bool
//
//    faker.Bool(
//        faker.SetRand(rand.New(rand.NewSource(time.Now().Unix()))), // Rand instance
//    )
//
func Bool(opts ...Option) bool {
	return Integer(0, 100, opts...) > 50
}
