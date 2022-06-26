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
	gTLD []string
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

// GTLD returns random generic top-level domain
func (i Internet) GTLD() string {
	return i.Faker.RandomStringElement(i.gTLD)
}

// Domain returns random domain
func (i Internet) Domain() string {
	return i.Faker.Asciify("***") + "." + i.GTLD()
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
