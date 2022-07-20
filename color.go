package faker

import "strings"

// Color returns random color
func (f *Faker) Color() string {
	return Color(
		WithRand(f.cfg.rand),
		WithColors(f.cfg.colors...),
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
	cfg := newConfig(opts...)
	if len(cfg.colors) == 0 {
		WithColors("Red", "Orange", "Yellow", "Green", "Blue", "Indigo", "Violet")(cfg)
	}
	return RandomElement(cfg.colors, opts...)
}

// Hex returns random hex
func (f *Faker) Hex() string {
	return Hex(
		WithRand(f.cfg.rand),
		WithHexSymbols(f.cfg.hexSymbols),
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
	cfg := newConfig(opts...)
	if len(cfg.hexSymbols) == 0 {
		WithHexSymbols("0123456789ABCDEF")(cfg)
	}
	for i := 0; i < 6; i++ {
		hexSymbol := Integer(0, len(cfg.hexSymbols)-1, opts...)
		hex.WriteByte(cfg.hexSymbols[hexSymbol])
	}
	return hex.String()
}

// RGB returns random RGB
func (f *Faker) RGB() []int {
	return RGB(WithRand(f.cfg.rand))
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
