package faker

import (
	"net"
	"net/http"
)

// GenericTopLevelDomain returns random generic top-level domain
func (f *Faker) GenericTopLevelDomain() string {
	return GenericTopLevelDomain(
		WithRand(f.cfg.rand),
		WithGenericTopLevelDomains(f.cfg.genericTopLevelDomains...),
	)
}

// GenericTopLevelDomain returns random generic top-level domain
//
//    faker.GenericTopLevelDomain(
//        faker.WithRand(rand.New(rand.NewSource(time.Now().Unix()))), // Rand instance
//        faker.WithGenericTopLevelDomains("com", "edu", "net", "org"), // Slice of generic top-level domains for RandomElement
//    )
//
func GenericTopLevelDomain(opts ...Option) string {
	cfg := newConfig(opts...)
	if len(cfg.genericTopLevelDomains) == 0 {
		WithGenericTopLevelDomains("com", "edu", "net", "org")(cfg)
	}
	return RandomElement(cfg.genericTopLevelDomains, opts...)
}

// IPv4 returns random IP v4 address
func (f *Faker) IPv4() string {
	return IPv4(WithRand(f.cfg.rand))
}

// IPv4 returns random IP v4 address
//
//    faker.IPv4(
//        faker.WithRand(rand.New(rand.NewSource(time.Now().Unix()))), // Rand instance
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
	return IPv6(WithRand(f.cfg.rand))
}

// IPv6 returns random IP v6 address
//
//    faker.IPv6(
//        faker.WithRand(rand.New(rand.NewSource(time.Now().Unix()))), // Rand instance
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

// HTTPMethod returns random HTTP method
func (f *Faker) HTTPMethod() string {
	return HTTPMethod(
		WithRand(f.cfg.rand),
		WithHTTPMethods(f.cfg.httpMethods...),
	)
}

// HTTPMethod returns random HTTP method
//
//    faker.HTTPMethod(
//        faker.WithRand(rand.New(rand.NewSource(time.Now().Unix()))), // Rand instance
//        faker.SetSetHTTPMethods("GET", "HEAD"), // Slice of HTTP method for RandomElement
//    )
//
func HTTPMethod(opts ...Option) string {
	cfg := newConfig(opts...)
	if len(cfg.httpMethods) == 0 {
		WithHTTPMethods(
			http.MethodGet,
			http.MethodHead,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
			http.MethodConnect,
			http.MethodOptions,
			http.MethodTrace,
		)(cfg)
	}
	return RandomElement(cfg.httpMethods, opts...)
}
