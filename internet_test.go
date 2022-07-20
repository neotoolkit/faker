package faker_test

import (
	"testing"

	"neotoolkit.com/faker"
)

func TestFaker_HTTPMethod(t *testing.T) {
	t.Parallel()
	f := faker.New(faker.SetHTTPMethods("test"))
	httpMethod := f.HTTPMethod()
	if len(httpMethod) == 0 {
		t.Error("HTTP method is empty")
	}
	if httpMethod != "test" {
		t.Errorf("got %s, want test", httpMethod)
	}
}

func TestHTTPMethod(t *testing.T) {
	t.Parallel()
	httpMethod := faker.HTTPMethod()
	if len(httpMethod) == 0 {
		t.Error("HTTP method is empty")
	}
}
