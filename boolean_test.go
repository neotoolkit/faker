package faker_test

import (
	"reflect"
	"testing"

	"github.com/neotoolkit/faker"
)

func TestBoolean_Boolean(t *testing.T) {
	b := faker.NewFaker().Boolean()

	if reflect.TypeOf(b.Boolean()).String() != "bool" {
		t.Error("Boolean type must be bool")
	}
}

func TestBoolean_String(t *testing.T) {
	b := faker.NewFaker().Boolean()

	if reflect.TypeOf(b.String()).String() != "string" {
		t.Error("String type must be string")
	}
}
