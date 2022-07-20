package faker

// PostCode returns random post code
func (f *Faker) PostCode() string {
	return PostCode(
		SetRand(f.options.rand),
		SetPostCodeFormats(f.options.postCodeFormats...),
	)
}

// PostCode returns random post code
//
//    faker.PostCode(
//        faker.SetRand(rand.New(rand.NewSource(time.Now().Unix()))), // Rand instance
//        faker.SetPostCodeFormats("****", "*****", "******") // Slice of post code format for RandomElement
//    )
//
func PostCode(opts ...Option) string {
	options := setOptions(opts...)
	if len(options.postCodeFormats) == 0 {
		options.SetPostCodeFormats("****", "*****", "******")
	}
	return Numerify(RandomElement(options.postCodeFormats, opts...), opts...)
}
