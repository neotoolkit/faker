package faker

// Address is a faker struct for Address
type Address struct {
	Faker                *Faker
	postCodeFormat       []string
	country              []string
	buildingNumberFormat []string
}

// PostCode returns a random post code for Address
func (a Address) PostCode() string {
	return a.Faker.Numerify(a.Faker.RandomStringElement(a.postCodeFormat))
}

// Country returns a random country for Address
func (a Address) Country() string {
	return a.Faker.RandomStringElement(a.country)
}

// BuildingNumber returns a random building number for Address
func (a Address) BuildingNumber() string {
	return a.Faker.Numerify(a.Faker.RandomStringElement(a.buildingNumberFormat))
}
