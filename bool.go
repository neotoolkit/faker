package faker

func (f *Faker) Bool() bool {
	return Bool(SetRand(f.options.rand))
}

func Bool(opts ...Option) bool {
	return Integer(0, 100, opts...) > 50
}
