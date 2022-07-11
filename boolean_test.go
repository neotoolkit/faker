package faker_test

import (
	"reflect"
	"testing"

	"github.com/neotoolkit/faker"
)

func TestBoolean(t *testing.T) {
	b := faker.Boolean()

	if reflect.TypeOf(b).String() != "bool" {
		t.Error("Boolean type must be bool")
	}
}
