package faker_test

import (
	"strings"
	"testing"

	"github.com/neotoolkit/faker"
)

func TestPerson_FirstName(t *testing.T) {
	firstName := faker.NewFaker().Person().FirstName()

	if len(firstName) == 0 {
		t.Error("firstName is empty")
	}
}

func TestPerson_LastName(t *testing.T) {
	lastName := faker.NewFaker().Person().LastName()

	if len(lastName) == 0 {
		t.Error("lastName is empty")
	}
}

func TestPerson_FirstNameMale(t *testing.T) {
	firstNameMale := faker.NewFaker().Person().FirstNameMale()

	if len(firstNameMale) == 0 {
		t.Error("firstNameMale is empty")
	}
}

func TestPerson_FirstNameFemale(t *testing.T) {
	firstNameFemale := faker.NewFaker().Person().FirstNameFemale()

	if len(firstNameFemale) == 0 {
		t.Error("firstNameFemale is empty")
	}
}

func TestPerson_Name(t *testing.T) {
	name := faker.NewFaker().Person().Name()

	if len(name) == 0 {
		t.Error("name is empty")
	}

	if strings.Contains(name, "{{FirstNameMale}}") {
		t.Error("name is format")
	}

	if strings.Contains(name, "{{FirstNameFemale}}") {
		t.Error("name is format")
	}

	if strings.Contains(name, "{{LastName}}") {
		t.Error("name is format")
	}
}

func TestPerson_NameMale(t *testing.T) {
	nameMale := faker.NewFaker().Person().NameMale()

	if len(nameMale) == 0 {
		t.Error("nameMale is empty")
	}
}

func TestPerson_NameFemale(t *testing.T) {
	nameFemale := faker.NewFaker().Person().NameFemale()

	if len(nameFemale) == 0 {
		t.Error("nameFemale is empty")
	}
}

func TestPerson_Gender(t *testing.T) {
	gender := faker.NewFaker().Person().Gender()

	if !(gender == "Male" || gender == "Female") {
		t.Error("gender must be male or female")
	}
}

func TestPerson_GenderMale(t *testing.T) {
	genderMale := faker.NewFaker().Person().GenderMale()

	if genderMale != "Male" {
		t.Error("genderMale must be male")
	}
}

func TestPerson_GenderFemale(t *testing.T) {
	genderFemale := faker.NewFaker().Person().GenderFemale()

	if genderFemale != "Female" {
		t.Error("genderFemale must be female")
	}
}
