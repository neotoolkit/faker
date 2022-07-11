package faker_test

import (
	"regexp"
	"testing"

	"github.com/neotoolkit/faker"
)

func TestUUID(t *testing.T) {
	uuid := faker.UUID()

	match, err := regexp.MatchString("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}$", uuid)
	if err != nil {
		t.Error(err)
	}

	if !match {
		t.Error("want true, got false")
	}
}

func TestUUID_UniqueInSequence(t *testing.T) {
	last := faker.UUID()
	current := faker.UUID()

	if last == current {
		t.Error("want unique uuid")
	}
}
