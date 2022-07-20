package faker

// FirstName returns random first name
func (f *Faker) FirstName() string {
	return FirstName(
		WithRand(f.options.rand),
		WithFirstNames(f.options.firstNames...),
	)
}

// FirstName returns random first name
//
//    faker.FirstName(
//        faker.WithRand(rand.New(rand.NewSource(time.Now().Unix()))), // Rand instance
//        faker.WithFirstNames("Tom", "Dick", "Harry"), // Slice of first names for RandomElement
//    )
//
func FirstName(opts ...Option) string {
	options := setOptions(opts...)
	if len(options.firstNames) == 0 {
		options.SetFirstNames("Tom", "Dick", "Harry")
	}
	return RandomElement(options.firstNames, opts...)
}

// LastName returns random last name
func (f *Faker) LastName() string {
	return LastName(
		WithRand(f.options.rand),
		WithLastNames(f.options.lastNames...),
	)
}

// LastName returns random last name
//
//    faker.LastName(
//        faker.WithRand(rand.New(rand.NewSource(time.Now().Unix()))), // Rand instance
//        faker.WithLastNames("Bloggs", "Doe", "Schmoe", "Smith"),
//    )
//
func LastName(opts ...Option) string {
	options := setOptions(opts...)
	if len(options.lastNames) == 0 {
		options.SetLastNames("Bloggs", "Doe", "Schmoe", "Smith")
	}
	return RandomElement(options.lastNames, opts...)
}

// Name returns random name
func (f *Faker) Name() string {
	return Name(
		WithRand(f.options.rand),
		WithFirstNames(f.options.firstNames...),
		WithLastNames(f.options.lastNames...),
	)
}

// Name returns random name
//
//    faker.Name(
//        faker.WithRand(rand.New(rand.NewSource(time.Now().Unix()))), // Rand instance
//        faker.WithFirstNames("Tom", "Dick", "Harry"),
//        faker.WithLastNames("Bloggs", "Doe", "Schmoe", "Smith"),
//    )
//
func Name(opts ...Option) string {
	firstName := FirstName(opts...)
	lastName := LastName(opts...)
	return firstName + " " + lastName
}
