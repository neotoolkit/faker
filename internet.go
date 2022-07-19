package faker

import (
	"net"
	"strings"
)

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
//        faker.SetRand(rand.New(rand.NewSource(time.Now().Unix()))),
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

// GenericTopLevelDomain returns random generic top-level domain
func (f *Faker) GenericTopLevelDomain() string {
	return GenericTopLevelDomain(
		SetRand(f.options.rand),
		SetGenericTopLevelDomains(f.options.genericTopLevelDomains...),
	)
}

// GenericTopLevelDomain returns random generic top-level domain
//
//    faker.GenericTopLevelDomain(
//        faker.SetRand(rand.New(rand.NewSource(time.Now().Unix()))),
//        faker.SetGenericTopLevelDomains("com", "edu", "net", "org"),
//    )
//
func GenericTopLevelDomain(opts ...Option) string {
	options := setOptions(opts...)
	if len(options.genericTopLevelDomains) == 0 {
		options.SetGenericTopLevelDomains("com", "edu", "net", "org")
	}
	return RandomElement(options.genericTopLevelDomains, opts...)
}

// IPv4 returns random IP v4 address
func (f *Faker) IPv4() string {
	return IPv4(SetRand(f.options.rand))
}

// IPv4 returns random IP v4 address
//
//    faker.IPv4(
//        faker.SetRand(rand.New(rand.NewSource(time.Now().Unix()))),
//    )
//
func IPv4(opts ...Option) string {
	size := 4
	ip := make([]byte, size)
	for i := 0; i < size; i++ {
		ip[i] = byte(Integer(0, 256, opts...))
	}
	return net.IP(ip).To4().String()
}

// IPv6 returns random IP v6 address
func (f *Faker) IPv6() string {
	return IPv6(SetRand(f.options.rand))
}

// IPv6 returns random IP v6 address
//
//    faker.IPv6(
//        faker.SetRand(rand.New(rand.NewSource(time.Now().Unix()))),
//    )
//
func IPv6(opts ...Option) string {
	size := 16
	ip := make([]byte, size)
	for i := 0; i < size; i++ {
		ip[i] = byte(Integer(0, 256, opts...))
	}
	return net.IP(ip).To16().String()
}
