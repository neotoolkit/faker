package faker_test

import (
	"reflect"
	"testing"

	"github.com/neotoolkit/faker"
)

func TestBoolean_Boolean(t *testing.T) {
	f := faker.NewFaker().Boolean()

	if reflect.TypeOf(f.Boolean()).String() != "bool" {
		t.Fatal("Boolean type must be bool")
	}
}
