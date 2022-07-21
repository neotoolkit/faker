package faker_test

import (
	"reflect"
	"testing"

	"github.com/neotoolkit/faker"
)

func TestFaker_Bool(t *testing.T) {
	t.Parallel()
	f := faker.New()
	b := f.Bool()
	if reflect.TypeOf(b).String() != "bool" {
		t.Error("Bool type must be bool")
	}
}

func TestBool(t *testing.T) {
	t.Parallel()
	b := faker.Bool()
	if reflect.TypeOf(b).String() != "bool" {
		t.Error("Bool type must be bool")
	}
}
