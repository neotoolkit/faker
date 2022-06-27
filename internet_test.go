package faker_test

import (
	"net"
	"testing"

	"github.com/neotoolkit/faker"
)

func TestInternet_Username(t *testing.T) {
	username := faker.NewFaker().Internet().Username()

	if len(username) == 0 {
		t.Error("username is empty")
	}
}

func TestInternet_GTLD(t *testing.T) {
	gTLD := faker.NewFaker().Internet().GTLD()

	if len(gTLD) == 0 {
		t.Error("GTLD is empty")
	}
}

func TestInternet_Domain(t *testing.T) {
	d := faker.NewFaker().Internet().Domain()

	if len(d) == 0 {
		t.Error("domain is empty")
	}
}

func TestInternet_Email(t *testing.T) {
	e := faker.NewFaker().Internet().Email()

	if len(e) == 0 {
		t.Error("email is empty")
	}
}

func TestInternet_IPv4(t *testing.T) {
	ip := faker.NewFaker().Internet().IPv4()

	if len(ip) == 0 {
		t.Error("IP is empty")
	}

	if nil == net.ParseIP(ip) {
		t.Errorf("cannot parse %s as IP", ip)
	}
}

func TestInternet_IPv6(t *testing.T) {
	ip := faker.NewFaker().Internet().IPv6()

	if len(ip) == 0 {
		t.Error("IP is empty")
	}

	if nil == net.ParseIP(ip) {
		t.Errorf("cannot parse %s as IP", ip)
	}
}

func TestInternet_Scheme(t *testing.T) {
	s := faker.NewFaker().Internet().Scheme()

	if len(s) == 0 {
		t.Error("scheme is empty")
	}
}

func TestInternet_URL(t *testing.T) {
	url := faker.NewFaker().Internet().URL()

	if len(url) == 0 {
		t.Error("URL is empty")
	}
}
