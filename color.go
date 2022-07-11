package faker

import "strings"

var (
	color = []string{
		"black", "blue",
		"gray",
		"green",
		"purple",
		"silver",
		"white",
		"yellow",
	}
	hexLetters = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "A", "B", "C", "D", "E", "F"}
)

// Color returns random color
func Color() string {
	return RandomStringElement(color)
}

// Hex returns random hex
func Hex() string {
	var hex strings.Builder
	hex.Grow(7)
	hex.WriteString("#")
	for i := 0; i < 6; i++ {
		hex.WriteString(RandomStringElement(hexLetters))
	}
	return hex.String()
}
