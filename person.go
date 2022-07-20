package faker

// FirstName returns random first name
func (f *Faker) FirstName() string {
	return FirstName(
		WithRand(f.cfg.rand),
		WithFirstNames(f.cfg.firstNames...),
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
	cfg := newConfig(opts...)
	if len(cfg.firstNames) == 0 {
		WithFirstNames("Tom", "Dick", "Harry")(cfg)
	}
	return RandomElement(cfg.firstNames, opts...)
}

// LastName returns random last name
func (f *Faker) LastName() string {
	return LastName(
		WithRand(f.cfg.rand),
		WithLastNames(f.cfg.lastNames...),
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
	cfg := newConfig(opts...)
	if len(cfg.lastNames) == 0 {
		WithLastNames("Bloggs", "Doe", "Schmoe", "Smith")(cfg)
	}
	return RandomElement(cfg.lastNames, opts...)
}

// Name returns random name
func (f *Faker) Name() string {
	return Name(
		WithRand(f.cfg.rand),
		WithFirstNames(f.cfg.firstNames...),
		WithLastNames(f.cfg.lastNames...),
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
