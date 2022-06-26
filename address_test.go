package faker_test

import (
	"testing"

	"github.com/neotoolkit/faker"
)

func TestAddress_PostCode(t *testing.T) {
	pc := faker.NewFaker().Address().PostCode()

	if len(pc) == 0 {
		t.Error("post code is empty")
	}
}

func TestAddress_Country(t *testing.T) {
	c := faker.NewFaker().Address().Country()

	if len(c) == 0 {
		t.Error("country is empty")
	}
}

func TestAddress_BuildingNumber(t *testing.T) {
	bn := faker.NewFaker().Address().BuildingNumber()

	if len(bn) == 0 {
		t.Error("building number is empty")
	}
}
