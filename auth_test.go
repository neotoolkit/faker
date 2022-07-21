package faker_test

import (
	"regexp"
	"testing"

	"github.com/neotoolkit/faker"
)

func TestFaker_Username(t *testing.T) {
	t.Parallel()
	f := faker.New(faker.WithFirstNames("test"))
	username := f.Username()
	if len(username) == 0 {
		t.Error("username is empty")
	}
	match, err := regexp.MatchString("[a-zA-Z][0-9]{2}$", username)
	if err != nil {
		t.Error(err)
	}
	if !match {
		t.Error("want true, got false")
	}
}

func TestUsername(t *testing.T) {
	t.Parallel()
	username := faker.Username()
	if len(username) == 0 {
		t.Error("username is empty")
	}
}
