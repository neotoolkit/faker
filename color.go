package faker

import "strings"

// Color is struct for Color
type Color struct {
	Faker      *Faker
	color      []string
	hexLetters []string
}

// Color returns random color
func (c Color) Color() string {
	return c.Faker.RandomStringElement(c.color)
}

// Hex returns random hex
func (c Color) Hex() string {
	var color strings.Builder
	color.Grow(7)
	color.WriteString("#")
	for i := 0; i < 6; i++ {
		color.WriteString(c.Faker.RandomStringElement(c.hexLetters))
	}
	return color.String()
}
