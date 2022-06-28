package faker_test

import (
	"testing"

	"github.com/neotoolkit/faker"
)

func TestColor_Color(t *testing.T) {
	color := faker.NewFaker().Color().Color()

	if len(color) == 0 {
		t.Error("color is empty")
	}
}

func TestColor_Hex(t *testing.T) {
	color := faker.NewFaker().Color().Hex()

	if len(color) != 7 {
		t.Error("hex is empty")
	}
}
