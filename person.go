package faker

import (
	"strings"
)

// Person is struct for Person
type Person struct {
	Faker            *Faker
	firstNameMale    []string
	firstNameFemale  []string
	lastName         []string
	maleNameFormat   []string
	femaleNameFormat []string
}

// FirstName returns random first name
func (p Person) FirstName() string {
	firstName := make([]string, len(p.firstNameMale)+len(p.firstNameFemale))
	i := 0

	for j := 0; j < len(p.firstNameMale); j++ {
		firstName[i] = p.firstNameMale[j]
		i++
	}

	for j := 0; j < len(p.firstNameFemale); j++ {
		firstName[i] = p.firstNameFemale[j]
		i++
	}

	return p.Faker.RandomStringElement(firstName)
}

// LastName returns random last name
func (p Person) LastName() string {
	i := p.Faker.IntBetween(0, len(p.lastName)-1)
	return p.lastName[i]
}

// FirstNameMale returns random male first name
func (p Person) FirstNameMale() string {
	i := p.Faker.IntBetween(0, len(p.firstNameMale)-1)
	return p.firstNameMale[i]
}

// FirstNameFemale returns random female first name
func (p Person) FirstNameFemale() string {
	i := p.Faker.IntBetween(0, len(p.firstNameFemale)-1)
	return p.firstNameFemale[i]
}

// Name returns random name
func (p Person) Name() string {
	format := make([]string, len(p.maleNameFormat)+len(p.femaleNameFormat))
	i := 0

	for j := 0; j < len(p.maleNameFormat); j++ {
		format[i] = p.maleNameFormat[j]
		i++
	}

	for j := 0; j < len(p.femaleNameFormat); j++ {
		format[i] = p.femaleNameFormat[j]
		i++
	}

	name := format[p.Faker.IntBetween(0, len(format)-1)]

	// {{firstNameMale}}
	if strings.Contains(name, "{{firstNameMale}}") {
		name = strings.ReplaceAll(name, "{{firstNameMale}}", p.FirstNameMale())
	}

	// {{firstNameFemale}}
	if strings.Contains(name, "{{firstNameFemale}}") {
		name = strings.ReplaceAll(name, "{{firstNameFemale}}", p.FirstNameFemale())
	}

	// {{lastName}}
	if strings.Contains(name, "{{lastName}}") {
		name = strings.ReplaceAll(name, "{{lastName}}", p.LastName())
	}

	return name
}

// NameMale returns random male name
func (p Person) NameMale() string {
	return p.FirstNameMale() + " " + p.LastName()
}

// NameFemale returns random female name
func (p Person) NameFemale() string {
	return p.FirstNameFemale() + " " + p.LastName()
}

// Gender returns random gender
func (p Person) Gender() string {
	return p.Faker.RandomStringElement([]string{p.GenderMale(), p.GenderFemale()})
}

// GenderMale returns male gender
func (p Person) GenderMale() string {
	return "Male"
}

// GenderFemale returns female gender
func (p Person) GenderFemale() string {
	return "Female"
}
