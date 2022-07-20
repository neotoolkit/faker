package faker

// PostCode returns random post code
func (f *Faker) PostCode() string {
	return PostCode(
		WithRand(f.cfg.rand),
		WithPostCodeFormats(f.cfg.postCodeFormats...),
	)
}

// PostCode returns random post code
//
//    faker.PostCode(
//        faker.WithRand(rand.New(rand.NewSource(time.Now().Unix()))), // Rand instance
//        faker.WithPostCodeFormats("****", "*****", "******") // Slice of post code format for RandomElement
//    )
//
func PostCode(opts ...Option) string {
	cfg := newConfig(opts...)
	if len(cfg.postCodeFormats) == 0 {
		WithPostCodeFormats("####", "#####", "######")(cfg)
	}
	return Numerify(RandomElement(cfg.postCodeFormats, opts...), opts...)
}
