package faker_test

import (
	"testing"

	"neotoolkit.com/faker"
)

func TestFaker_Color(t *testing.T) {
	t.Parallel()
	f := faker.New(faker.SetColors("test"))
	color := f.Color()
	if len(color) == 0 {
		t.Error("color is empty")
	}
	if color != "test" {
		t.Errorf("got %s, want test", color)
	}
}

func TestColor(t *testing.T) {
	t.Parallel()
	color := faker.Color()
	if len(color) == 0 {
		t.Error("color is empty")
	}
}

func TestFaker_Hex(t *testing.T) {
	t.Parallel()
	f := faker.New(faker.SetHexSymbols("T"))
	hex := f.Hex()
	if len(hex) != 7 {
		t.Error("length of hex must be 7")
	}
	if hex != "#TTTTTT" {
		t.Errorf("got %s, want test", hex)
	}
}

func TestHex(t *testing.T) {
	t.Parallel()
	hex := faker.Hex()
	if len(hex) != 7 {
		t.Error("length of hex must be 7")
	}
}
