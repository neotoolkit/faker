package faker_test

import (
	"testing"

	"github.com/neotoolkit/faker"
)

func TestUsername(t *testing.T) {
	i := faker.NewFaker().Internet()
	username := i.Username()

	if len(username) == 0 {
		t.Fatal("username is empty")
	}
}

func TestGTLD(t *testing.T) {
	i := faker.NewFaker().Internet()
	gTLD := i.GTLD()

	if len(gTLD) == 0 {
		t.Fatal("gTLD is empty")
	}
}

func TestDomain(t *testing.T) {
	i := faker.NewFaker().Internet()
	d := i.Domain()

	if len(d) == 0 {
		t.Fatal("domain is empty")
	}
}

func TestEmail(t *testing.T) {
	i := faker.NewFaker().Internet()
	e := i.Email()

	if len(e) == 0 {
		t.Fatal("domain is empty")
	}
}
