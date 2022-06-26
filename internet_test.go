package faker_test

import (
	"testing"

	"github.com/neotoolkit/faker"
)

func TestUsername(t *testing.T) {
	username := faker.NewFaker().Internet().Username()

	if len(username) == 0 {
		t.Error("username is empty")
	}
}

func TestGTLD(t *testing.T) {
	gTLD := faker.NewFaker().Internet().GTLD()

	if len(gTLD) == 0 {
		t.Error("GTLD is empty")
	}
}

func TestDomain(t *testing.T) {
	d := faker.NewFaker().Internet().Domain()

	if len(d) == 0 {
		t.Error("domain is empty")
	}
}

func TestEmail(t *testing.T) {
	e := faker.NewFaker().Internet().Email()

	if len(e) == 0 {
		t.Error("email is empty")
	}
}
