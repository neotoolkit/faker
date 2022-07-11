package faker_test

import (
	"testing"

	"github.com/neotoolkit/faker"
)

func TestColor(t *testing.T) {
	color := faker.Color()

	if len(color) == 0 {
		t.Error("color is empty")
	}
}

func TestHex(t *testing.T) {
	hex := faker.Hex()

	if len(hex) != 7 {
		t.Error("hex is empty")
	}
}
