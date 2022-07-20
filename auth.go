package faker

import "strings"

// Username returns random username from first names and two number
func (f *Faker) Username() string {
	return Username(
		WithRand(f.cfg.rand),
		WithFirstNames(f.cfg.firstNames...),
	)
}

// Username returns random username from first names and two number
//
//    faker.Username(
//        faker.WithRand(rand.New(rand.NewSource(time.Now().Unix()))),
//        faker.WithFirstNames("Tom")
//    )
//
func Username(opts ...Option) string {
	return FirstName(opts...) + Numerify("##", opts...)
}

// Password returns random password
func (f *Faker) Password() string {
	return Password(
		WithRand(f.cfg.rand),
		WithPasswordMin(f.cfg.passwordMin),
		WithPasswordMax(f.cfg.passwordMax),
		WithPasswordChars(f.cfg.passwordChars),
	)
}

// Password returns random password
//
//    faker.Password(
//        faker.WithRand(rand.New(rand.NewSource(time.Now().Unix()))), // Rand instance
//        faker.WithPasswordMin(8),
//        faker.WithPasswordMax(16),
//        faker.WithPasswordChars("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz1234567890"),
//    )
//
func Password(opts ...Option) string {
	cfg := newConfig(opts...)
	if cfg.passwordMin == 0 {
		WithPasswordMin(8)(cfg)
	}
	if cfg.passwordMax == 0 {
		WithPasswordMax(16)(cfg)
	}
	if len(cfg.passwordChars) == 0 {
		WithPasswordChars("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz1234567890")(cfg)
	}
	length := Integer(cfg.passwordMin, cfg.passwordMax, opts...)
	var password strings.Builder
	password.Grow(length)
	for i := 0; i < length; i++ {
		char := Integer(0, len(cfg.passwordChars)-1, opts...)
		password.WriteByte(cfg.passwordChars[char])
	}
	return password.String()
}
