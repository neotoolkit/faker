package faker_test

import (
	"regexp"
	"testing"

	"go.neotoolkit.com/faker"
)

func TestFaker_UUID(t *testing.T) {
	t.Parallel()
	f := faker.New()
	uuid := f.UUID()
	matchUUID(t, uuid)
}

func TestUUID(t *testing.T) {
	t.Parallel()
	uuid := faker.UUID()
	matchUUID(t, uuid)
}

func matchUUID(t *testing.T, uuid string) {
	t.Helper()
	match, err := regexp.MatchString("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}$", uuid)
	if err != nil {
		t.Error(err)
	}
	if !match {
		t.Error("want true, got false")
	}
}

func TestFaker_UUID_UniqueInSequence(t *testing.T) {
	t.Parallel()
	f := faker.New()
	last := f.UUID()
	current := f.UUID()
	if last == current {
		t.Error("want unique uuid")
	}
}

func TestUUID_UniqueInSequence(t *testing.T) {
	t.Parallel()
	last := faker.UUID()
	current := faker.UUID()
	if last == current {
		t.Error("want unique uuid")
	}
}
