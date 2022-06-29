package faker

import (
	"net"
	"strings"
)

// Internet is struct for Internet
type Internet struct {
	Faker          *Faker
	usernameFormat []string
	// Generic top-level domain
	gTLD      []string
	scheme    []string
	urlFormat []string
}

// Username returns random username
func (i Internet) Username() string {
	username := i.Faker.RandomStringElement(i.usernameFormat)

	p := i.Faker.Person()

	// {{firstName}}
	if strings.Contains(username, "{{firstName}}") {
		username = strings.ReplaceAll(username, "{{firstName}}", strings.ToLower(p.FirstName()))
	}

	// {{lastName}}
	if strings.Contains(username, "{{lastName}}") {
		username = strings.ReplaceAll(username, "{{lastName}}", strings.ToLower(p.LastName()))
	}

	return username
}

// Password returns random password
func (i Internet) Password() string {
	pattern := strings.Repeat("*", i.Faker.IntBetween(8, 16))
	return i.Faker.Asciify(pattern)
}

// GTLD returns random generic top-level domain
func (i Internet) GTLD() string {
	return i.Faker.RandomStringElement(i.gTLD)
}

// Domain returns random domain
func (i Internet) Domain() string {
	return i.Faker.LowerAsciify("***") + "." + i.GTLD()
}

// Email returns random email
func (i Internet) Email() string {
	return i.Username() + "@" + i.Domain()
}

// IPv4 returns random IP v4 address
func (i Internet) IPv4() string {
	size := 4
	ip := make([]byte, size)
	for idx := 0; idx < size; idx++ {
		ip[idx] = byte(i.Faker.IntBetween(0, 256))
	}
	return net.IP(ip).To4().String()
}

// IPv6 returns random IP v6 address
func (i Internet) IPv6() string {
	size := 16
	ip := make([]byte, size)
	for idx := 0; idx < size; idx++ {
		ip[idx] = byte(i.Faker.IntBetween(0, 256))
	}
	return net.IP(ip).To16().String()
}

// Scheme returns random scheme
func (i Internet) Scheme() string {
	return i.Faker.RandomStringElement(i.scheme)
}

// URL returns random URL
func (i Internet) URL() string {
	url := i.Faker.RandomStringElement(i.urlFormat)

	// {{scheme}}
	if strings.Contains(url, "{{scheme}}") {
		url = strings.ReplaceAll(url, "{{scheme}}", strings.ToLower(i.Scheme()))
	}

	// {{domain}}
	if strings.Contains(url, "{{domain}}") {
		url = strings.ReplaceAll(url, "{{domain}}", strings.ToLower(i.Domain()))
	}

	// {{path}}
	if strings.Contains(url, "{{path}}") {
		in := strings.Repeat("*", i.Faker.IntBetween(1, 10))
		url = strings.ReplaceAll(url, "{{path}}", i.Faker.LowerAsciify(in))
	}

	// {{query}}
	if strings.Contains(url, "{{query}}") {
		keyIn := strings.Repeat("*", i.Faker.IntBetween(1, 10))
		key := i.Faker.LowerAsciify(keyIn)
		valueIn := strings.Repeat("*", i.Faker.IntBetween(1, 10))
		value := i.Faker.LowerAsciify(valueIn)
		url = strings.ReplaceAll(url, "{{query}}", key+"="+value)
	}

	// {{fragment}}
	if strings.Contains(url, "{{fragment}}") {
		in := strings.Repeat("*", i.Faker.IntBetween(1, 10))
		url = strings.ReplaceAll(url, "{{fragment}}", i.Faker.LowerAsciify(in))
	}

	return url
}
