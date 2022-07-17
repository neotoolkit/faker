package faker

import (
	"net"
	"strings"
)

var (
	usernameFormat = []string{
		"{{lastName}}.{{firstName}}",
		"{{firstName}}.{{lastName}}",
		"{{firstName}}",
		"{{lastName}}",
	}
	scheme = []string{
		"http",
		"https",
	}
	urlFormat = []string{
		"{{scheme}}://{{domain}}",
		"{{scheme}}://{{domain}}/{{path}}",
		"{{scheme}}://{{domain}}/{{path}}?{{query}}",
		"{{scheme}}://{{domain}}/{{path}}?{{query}}#{{fragment}}",
		"{{scheme}}://www.{{domain}}/{{path}}",
		"{{scheme}}://www.{{domain}}/{{path}}?{{query}}",
		"{{scheme}}://www.{{domain}}/{{path}}?{{query}}#{{fragment}}",
	}
)

// Username returns random username
func Username() string {
	username := RandomStringElement(usernameFormat)

	// {{firstName}}
	if strings.Contains(username, "{{firstName}}") {
		username = strings.ReplaceAll(username, "{{firstName}}", strings.ToLower(FirstName()))
	}

	// {{lastName}}
	if strings.Contains(username, "{{lastName}}") {
		username = strings.ReplaceAll(username, "{{lastName}}", strings.ToLower(LastName()))
	}

	return username
}

// Password returns random password
func Password(opts ...PasswordOption) string {
	options := PasswordOptions{
		min:   8,
		max:   16,
		chars: "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz1234567890",
	}
	for _, opt := range opts {
		opt(&options)
	}
	length := IntBetween(options.min, options.max)
	var password strings.Builder
	password.Grow(length)
	for i := 0; i < length; i++ {
		char := IntBetween(0, len(options.chars)-1)
		password.WriteByte(options.chars[char])
	}
	return password.String()
}

type PasswordOption func(opts *PasswordOptions)

type PasswordOptions struct {
	min   int
	max   int
	chars string
}

func SetPasswordMin(min int) PasswordOption {
	return func(opts *PasswordOptions) {
		opts.min = min
	}
}

func SetPasswordMax(max int) PasswordOption {
	return func(opts *PasswordOptions) {
		opts.max = max
	}
}

func SetPasswordChars(chars string) PasswordOption {
	return func(opts *PasswordOptions) {
		opts.chars = chars
	}
}

// GTLD returns random generic top-level domain
func GTLD(opts ...GTLDOption) string {
	options := GTLDOptions{
		domains: []string{"com", "net", "org", "biz", "info"},
	}
	for _, opt := range opts {
		opt(&options)
	}
	return options.domains[IntBetween(0, len(options.domains)-1)]
}

type GTLDOption func(opts *GTLDOptions)

type GTLDOptions struct {
	domains []string
}

func SetGTLDDomains(domains []string) GTLDOption {
	return func(opts *GTLDOptions) {
		opts.domains = domains
	}
}

// Domain returns random domain
func Domain() string {
	return LowerAsciify("***") + "." + GTLD()
}

// Email returns random email
func Email() string {
	return Username() + "@" + Domain()
}

// IPv4 returns random IP v4 address
func IPv4() string {
	size := 4
	ip := make([]byte, size)
	for i := 0; i < size; i++ {
		ip[i] = byte(IntBetween(0, 256))
	}
	return net.IP(ip).To4().String()
}

// IPv6 returns random IP v6 address
func IPv6() string {
	size := 16
	ip := make([]byte, size)
	for i := 0; i < size; i++ {
		ip[i] = byte(IntBetween(0, 256))
	}
	return net.IP(ip).To16().String()
}

// Scheme returns random scheme
func Scheme() string {
	return RandomStringElement(scheme)
}

// URL returns random URL
func URL() string {
	url := RandomStringElement(urlFormat)

	// {{scheme}}
	if strings.Contains(url, "{{scheme}}") {
		url = strings.ReplaceAll(url, "{{scheme}}", strings.ToLower(Scheme()))
	}

	// {{domain}}
	if strings.Contains(url, "{{domain}}") {
		url = strings.ReplaceAll(url, "{{domain}}", strings.ToLower(Domain()))
	}

	// {{path}}
	if strings.Contains(url, "{{path}}") {
		in := strings.Repeat("*", IntBetween(1, 10))
		url = strings.ReplaceAll(url, "{{path}}", LowerAsciify(in))
	}

	// {{query}}
	if strings.Contains(url, "{{query}}") {
		n := IntBetween(1, 10)
		query := make([]string, n)
		for i := 0; i < n; i++ {
			keyIn := strings.Repeat("*", IntBetween(1, 10))
			key := LowerAsciify(keyIn)
			valueIn := strings.Repeat("*", IntBetween(1, 10))
			value := LowerAsciify(valueIn)
			query[i] = key + "=" + value
		}
		url = strings.ReplaceAll(url, "{{query}}", strings.Join(query, "&"))
	}

	// {{fragment}}
	if strings.Contains(url, "{{fragment}}") {
		in := strings.Repeat("*", IntBetween(1, 10))
		url = strings.ReplaceAll(url, "{{fragment}}", LowerAsciify(in))
	}

	return url
}
