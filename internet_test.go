package faker_test

import (
	"net"
	"testing"

	"github.com/neotoolkit/faker"
)

func TestUsername(t *testing.T) {
	username := faker.Username()

	if len(username) == 0 {
		t.Error("username is empty")
	}
}

func TestPassword(t *testing.T) {
	password := faker.Password()

	if len(password) == 0 {
		t.Error("password is empty")
	}
	if len(password) < 8 || len(password) > 16 {
		t.Error("password must be greater or equal and less or equal 16")
	}
}

func TestPasswordWithOpts(t *testing.T) {
	password := faker.Password(
		faker.SetPasswordMin(10),
		faker.SetPasswordMax(12),
		faker.SetPasswordChars("abc"),
	)

	if len(password) == 0 {
		t.Error("password is empty")
	}
	if len(password) < 10 || len(password) > 12 {
		t.Error("password must be greater 10 or equal and less or equal 12")
	}
}

func TestGTLD(t *testing.T) {
	gTLD := faker.GTLD()

	if len(gTLD) == 0 {
		t.Error("GTLD is empty")
	}
}

func TestDomain(t *testing.T) {
	d := faker.Domain()

	if len(d) == 0 {
		t.Error("domain is empty")
	}
}

func TestEmail(t *testing.T) {
	e := faker.Email()

	if len(e) == 0 {
		t.Error("email is empty")
	}
}

func TestIPv4(t *testing.T) {
	ip := faker.IPv4()

	if len(ip) == 0 {
		t.Error("IP is empty")
	}

	if nil == net.ParseIP(ip) {
		t.Errorf("cannot parse %s as IP", ip)
	}
}

func TestIPv6(t *testing.T) {
	ip := faker.IPv6()

	if len(ip) == 0 {
		t.Error("IP is empty")
	}

	if nil == net.ParseIP(ip) {
		t.Errorf("cannot parse %s as IP", ip)
	}
}

func TestScheme(t *testing.T) {
	s := faker.Scheme()

	if len(s) == 0 {
		t.Error("scheme is empty")
	}
}

func TestURL(t *testing.T) {
	url := faker.URL()

	if len(url) == 0 {
		t.Error("URL is empty")
	}
}
