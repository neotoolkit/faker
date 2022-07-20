package faker

import "strings"

// Username returns random username from first names and two number
func (f *Faker) Username() string {
	return Username(
		SetRand(f.options.rand),
		SetFirstNames(f.options.firstNames...),
	)
}

// Username returns random username from first names and two number
//
//    faker.Username(
//        faker.SetRand(rand.New(rand.NewSource(time.Now().Unix()))),
//        faker.SetSetFirstNames("Tom")
//    )
//
func Username(opts ...Option) string {
	return FirstName(opts...) + Numerify("**", opts...)
}

// Password returns random password
func (f *Faker) Password() string {
	return Password(
		SetRand(f.options.rand),
		SetPasswordMin(f.options.passwordMin),
		SetPasswordMax(f.options.passwordMax),
		SetPasswordChars(f.options.passwordChars),
	)
}

// Password returns random password
//
//    faker.Password(
//        faker.SetRand(rand.New(rand.NewSource(time.Now().Unix()))), // Rand instance
//        faker.SetPasswordMin(8),
//        faker.SetPasswordMax(16),
//        faker.SetPasswordChars("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz1234567890"),
//    )
//
func Password(opts ...Option) string {
	options := setOptions(opts...)
	if options.passwordMin == 0 {
		options.SetPasswordMin(8)
	}
	if options.passwordMax == 0 {
		options.SetPasswordMax(16)
	}
	if len(options.passwordChars) == 0 {
		options.SetPasswordChars("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz1234567890")
	}
	length := Integer(options.passwordMin, options.passwordMax, opts...)
	var password strings.Builder
	password.Grow(length)
	for i := 0; i < length; i++ {
		char := Integer(0, len(options.passwordChars)-1, opts...)
		password.WriteByte(options.passwordChars[char])
	}
	return password.String()
}
