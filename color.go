package faker

import "strings"

// Color returns random color
func (f *Faker) Color() string {
	return Color(
		SetRand(f.options.rand),
		SetColors(f.options.colors...),
	)
}

// Color returns random color
//
//    faker.Color(
//        faker.SetRand(rand.New(rand.NewSource(time.Now().Unix()))), // Rand instance
//        faker.SetColors(
//            "Red",
//            "Orange",
//            "Yellow",
//            "Green",
//            "Blue",
//            "Indigo",
//            "Violet",
//        ), // Slice of color for RandomElement
//    )
//
func Color(opts ...Option) string {
	options := setOptions(opts...)
	if len(options.colors) == 0 {
		options.SetColors(
			"Red",
			"Orange",
			"Yellow",
			"Green",
			"Blue",
			"Indigo",
			"Violet",
		)
	}
	return RandomElement(options.colors, opts...)
}

// Hex returns random hex
func (f *Faker) Hex() string {
	return Hex(
		SetRand(f.options.rand),
		SetHexLetters(f.options.hexLetters...),
	)
}

// Hex returns random hex
//
//    faker.Hex(
//        faker.SetRand(rand.New(rand.NewSource(time.Now().Unix()))), // Rand instance
//        faker.SetHexLetters(
//            "0",
//            "1",
//            "2",
//            "3",
//            "4",
//            "5",
//            "6",
//            "7",
//            "8",
//            "9",
//            "A",
//            "B",
//            "C",
//            "D",
//            "E",
//            "F",
//        ), // Slice of hex letter for RandomElement
//    )
//
func Hex(opts ...Option) string {
	var hex strings.Builder
	hex.Grow(7)
	hex.WriteString("#")
	options := setOptions(opts...)
	if len(options.hexLetters) == 0 {
		options.SetHexLetters("0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "A", "B", "C", "D", "E", "F")
	}
	for i := 0; i < 6; i++ {
		hex.WriteString(RandomElement(options.hexLetters, opts...))
	}
	return hex.String()
}
