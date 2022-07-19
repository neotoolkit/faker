package faker

// FirstName returns random first name
func (f *Faker) FirstName() string {
	return FirstName(
		SetRand(f.options.rand),
		SetFirstNames(f.options.firstNames...),
	)
}

// FirstName returns random first name
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
		SetRand(f.options.rand),
		SetLastNames(f.options.lastNames...),
	)
}

// LastName returns random last name
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
		SetRand(f.options.rand),
		SetFirstNames(f.options.firstNames...),
		SetLastNames(f.options.lastNames...),
	)
}

// Name returns random name
func Name(opts ...Option) string {
	firstName := FirstName(opts...)
	lastName := LastName(opts...)
	return firstName + " " + lastName
}
