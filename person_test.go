package faker_test

import (
	"strings"
	"testing"

	"github.com/neotoolkit/faker"
)

func TestFirstName(t *testing.T) {
	firstName := faker.FirstName()

	if len(firstName) == 0 {
		t.Error("firstName is empty")
	}
}

func TestLastName(t *testing.T) {
	lastName := faker.LastName()

	if len(lastName) == 0 {
		t.Error("lastName is empty")
	}
}

func TestFirstNameMale(t *testing.T) {
	firstNameMale := faker.FirstNameMale()

	if len(firstNameMale) == 0 {
		t.Error("firstNameMale is empty")
	}
}

func TestFirstNameFemale(t *testing.T) {
	firstNameFemale := faker.FirstNameFemale()

	if len(firstNameFemale) == 0 {
		t.Error("firstNameFemale is empty")
	}
}

func TestName(t *testing.T) {
	name := faker.Name()

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

func TestNameMale(t *testing.T) {
	nameMale := faker.NameMale()

	if len(nameMale) == 0 {
		t.Error("nameMale is empty")
	}
}

func TestNameFemale(t *testing.T) {
	nameFemale := faker.NameFemale()

	if len(nameFemale) == 0 {
		t.Error("nameFemale is empty")
	}
}

func TestGender(t *testing.T) {
	gender := faker.Gender()

	if !(gender == "Male" || gender == "Female") {
		t.Error("gender must be male or female")
	}
}

func TestGenderMale(t *testing.T) {
	genderMale := faker.GenderMale()

	if genderMale != "Male" {
		t.Error("genderMale must be male")
	}
}

func TestGenderFemale(t *testing.T) {
	genderFemale := faker.GenderFemale()

	if genderFemale != "Female" {
		t.Error("genderFemale must be female")
	}
}
