package faker_test

import (
	"testing"

	"neotoolkit.com/faker"
)

func TestFaker_Color(t *testing.T) {
	t.Parallel()
	f := faker.New(faker.WithColors("test"))
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
	f := faker.New(faker.WithHexSymbols("T"))
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

func TestFaker_RGB(t *testing.T) {
	t.Parallel()
	f := faker.New()
	rgb := f.RGB()
	if len(rgb) != 3 {
		t.Error("length of RGB slice must be 3")
	}
	for _, val := range rgb {
		if val < 0 || val > 255 {
			t.Error("RGB value must be equal or greeter than 0 and equal or less 255")
		}
	}
}

func TestRGB(t *testing.T) {
	t.Parallel()
	rgb := faker.RGB()
	if len(rgb) != 3 {
		t.Error("length of RGB slice must be 3")
	}
}
