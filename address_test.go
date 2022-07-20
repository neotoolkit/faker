package faker_test

import (
	"testing"

	"neotoolkit.com/faker"
)

func TestFaker_PostCode(t *testing.T) {
	t.Parallel()
	f := faker.New(faker.WithPostCodeFormats("test"))
	postCode := f.PostCode()
	if len(postCode) == 0 {
		t.Error("post code is empty")
	}
	if postCode != "test" {
		t.Errorf("got %s, want test", postCode)
	}
}

func TestPostCode(t *testing.T) {
	t.Parallel()
	postCode := faker.PostCode()
	if len(postCode) == 0 {
		t.Error("post code is empty")
	}
}
