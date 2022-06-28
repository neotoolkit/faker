package faker

// Boolean is struct for Boolean
type Boolean struct {
	Faker *Faker
}

// Boolean returns random boolean result
func (b Boolean) Boolean() bool {
	return b.Faker.IntBetween(0, 100) > 50
}

// String returns random boolean result as string
func (b Boolean) String() string {
	if b.Faker.IntBetween(0, 100) > 50 {
		return "true"
	}
	return "false"
}
