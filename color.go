package faker

import "strings"

// Color returns random color
func (f *Faker) Color() string {
	return Color(
		WithRand(f.options.rand),
		WithColors(f.options.colors...),
	)
}

// Color returns random color
//
//    faker.Color(
//        faker.WithRand(rand.New(rand.NewSource(time.Now().Unix()))), // Rand instance
//        faker.WithColors("Red", "Orange", "Yellow", "Green", "Blue", "Indigo", "Violet"), // Slice of color for RandomElement
//    )
//
func Color(opts ...Option) string {
	options := setOptions(opts...)
	if len(options.colors) == 0 {
		options.SetColors("Red", "Orange", "Yellow", "Green", "Blue", "Indigo", "Violet")
	}
	return RandomElement(options.colors, opts...)
}

// Hex returns random hex
func (f *Faker) Hex() string {
	return Hex(
		WithRand(f.options.rand),
		WithHexSymbols(f.options.hexSymbols),
	)
}

// Hex returns random hex
//
//    faker.Hex(
//        faker.WithRand(rand.New(rand.NewSource(time.Now().Unix()))), // Rand instance
//        faker.WithHexSymbols("0123456789ABCDEF"), // Hex symbols as string
//    )
//
func Hex(opts ...Option) string {
	var hex strings.Builder
	hex.Grow(7)
	hex.WriteString("#")
	options := setOptions(opts...)
	if len(options.hexSymbols) == 0 {
		options.SetHexSymbols("0123456789ABCDEF")
	}
	for i := 0; i < 6; i++ {
		hexSymbol := Integer(0, len(options.hexSymbols)-1, opts...)
		hex.WriteByte(options.hexSymbols[hexSymbol])
	}
	return hex.String()
}

// RGB returns random RGB
func (f *Faker) RGB() []int {
	return RGB(WithRand(f.options.rand))
}

// RGB returns random RGB
//
//    faker.RGB(
//        faker.WithRand(rand.New(rand.NewSource(time.Now().Unix()))), // Rand instance
//    )
//
func RGB(opts ...Option) []int {
	return []int{Integer(0, 255, opts...), Integer(0, 255, opts...), Integer(0, 255, opts...)}
}
