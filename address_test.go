package faker_test

import (
	"testing"

	"github.com/neotoolkit/faker"
)

func TestPostCode(t *testing.T) {
	pc := faker.PostCode()

	if len(pc) == 0 {
		t.Error("post code is empty")
	}
}

func TestCountry(t *testing.T) {
	c := faker.Country()

	if len(c) == 0 {
		t.Error("country is empty")
	}
}

func TestBuildingNumber(t *testing.T) {
	bn := faker.BuildingNumber()

	if len(bn) == 0 {
		t.Error("building number is empty")
	}
}
