package faker_test

import (
	"regexp"
	"testing"

	"github.com/neotoolkit/faker"
)

func TestUUID_v4(t *testing.T) {
	f := faker.NewFaker()
	value := f.UUID().V4()

	match, err := regexp.MatchString("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}$", value)
	if err != nil {
		t.Error(err)
	}

	if !match {
		t.Error("want true, got false")
	}
}

func TestUUID_V4UniqueInSequence(t *testing.T) {
	f := faker.NewFaker()
	last := f.UUID().V4()
	current := f.UUID().V4()

	if last == current {
		t.Error("want unique uuid")
	}
}
