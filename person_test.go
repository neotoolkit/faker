package faker_test

import (
	"strings"
	"testing"

	"github.com/neotoolkit/faker"
)

func TestPerson_FirstName(t *testing.T) {
	f := faker.NewFaker()
	firstName := f.Person().FirstName()

	if len(firstName) == 0 {
		t.Fatal("firstName is empty")
	}
}

func TestPerson_LastName(t *testing.T) {
	f := faker.NewFaker()
	p := f.Person()
	lastName := p.LastName()

	if len(lastName) == 0 {
		t.Fatal("lastName is empty")
	}
}

func TestPerson_FirstNameMale(t *testing.T) {
	f := faker.NewFaker()
	firstNameMale := f.Person().FirstNameMale()

	if len(firstNameMale) == 0 {
		t.Fatal("firstNameMale is empty")
	}
}

func TestPerson_FirstNameFemale(t *testing.T) {
	f := faker.NewFaker()
	p := f.Person()
	firstNameFemale := p.FirstNameFemale()

	if len(firstNameFemale) == 0 {
		t.Fatal("firstNameFemale is empty")
	}
}

func TestPerson_Name(t *testing.T) {
	f := faker.NewFaker()
	p := f.Person()
	name := p.Name()

	if len(name) == 0 {
		t.Fatal("name is empty")
	}

	if strings.Contains(name, "{{FirstNameMale}}") {
		t.Fatal("name is format")
	}

	if strings.Contains(name, "{{FirstNameFemale}}") {
		t.Fatal("name is format")
	}

	if strings.Contains(name, "{{LastName}}") {
		t.Fatal("name is format")
	}
}

func TestPerson_NameMale(t *testing.T) {
	f := faker.NewFaker()
	p := f.Person()
	nameMale := p.NameMale()

	if len(nameMale) == 0 {
		t.Fatal("nameMale is empty")
	}
}

func TestPerson_NameFemale(t *testing.T) {
	f := faker.NewFaker()
	p := f.Person()
	nameFemale := p.NameFemale()

	if len(nameFemale) == 0 {
		t.Fatal("nameFemale is empty")
	}
}

func TestPerson_Gender(t *testing.T) {
	f := faker.NewFaker()
	p := f.Person()
	gender := p.Gender()

	if !(gender == "Male" || gender == "Female") {
		t.Fatal("gender must be male or female")
	}
}

func TestPerson_GenderMale(t *testing.T) {
	f := faker.NewFaker()
	p := f.Person()
	genderMale := p.GenderMale()

	if genderMale != "Male" {
		t.Fatal("genderMale must be male")
	}
}

func TestPerson_GenderFemale(t *testing.T) {
	f := faker.NewFaker()
	p := f.Person()
	genderFemale := p.GenderFemale()

	if genderFemale != "Female" {
		t.Fatal("genderFemale must be female")
	}
}
