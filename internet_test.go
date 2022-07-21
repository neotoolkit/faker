package faker_test

import (
	"testing"

	"github.com/neotoolkit/faker"
)

func TestFaker_HTTPMethod(t *testing.T) {
	t.Parallel()
	f := faker.New(faker.WithHTTPMethods("test"))
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

func TestFaker_HTTPStatusCode(t *testing.T) {
	t.Parallel()
	f := faker.New(faker.WithHTTPStatusCodes(200))
	httpStatusCode := f.HTTPStatusCode()
	if httpStatusCode != 200 {
		t.Errorf("got %v, want 200", httpStatusCode)
	}
}

func TestHTTPStatusCode(t *testing.T) {
	t.Parallel()
	httpStatusCode := faker.HTTPStatusCode()
	if httpStatusCode == 0 {
		t.Error("HTTP status code is empty")
	}
}
