package faker_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/neotoolkit/faker"
)

func TestIntBetween(t *testing.T) {
	t.Parallel()

	f := faker.NewFaker()

	for _, tc := range []struct {
		name string
		min  int
		max  int
	}{
		{
			name: "min 1, max 100",
			min:  1,
			max:  100,
		},
		{
			name: "min 1, max 1",
			min:  1,
			max:  1,
		},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			value := f.IntBetween(tc.min, tc.max)

			valueType := fmt.Sprintf("%T", value)

			if valueType != "int" {
				t.Fatalf("value type want int, got %T", value)
			}

			if value < tc.min {
				t.Fatalf("value must be less %d", tc.min)
			}

			if value > tc.max {
				t.Fatalf("value must be greater %d", tc.max)
			}
		})
	}
}

func TestByName(t *testing.T) {
	t.Parallel()

	for _, tc := range []struct {
		name  string
		faker string
	}{
		{
			name:  "",
			faker: "Boolean",
		},
		{
			name:  "",
			faker: "Username",
		},
		{
			name:  "",
			faker: "GTLD",
		},
		{
			name:  "",
			faker: "Domain",
		},
		{
			name:  "",
			faker: "Email",
		},
		{
			name:  "",
			faker: "firstname",
		},
		{
			name:  "",
			faker: "Person.FirstName",
		},
		{
			name:  "",
			faker: "lastname",
		},
		{
			name:  "",
			faker: "Person.LastName",
		},
		{
			name:  "",
			faker: "firstname male",
		},
		{
			name:  "",
			faker: "Person.FirstNameMale",
		},
		{
			name:  "",
			faker: "firstname female",
		},
		{
			name:  "",
			faker: "Person.FirstNameFemale",
		},
		{
			name:  "",
			faker: "name",
		},
		{
			name:  "",
			faker: "Person.Name",
		},
		{
			name:  "",
			faker: "name male",
		},
		{
			name:  "",
			faker: "Person.NameMale",
		},
		{
			name:  "",
			faker: "name female",
		},
		{
			name:  "",
			faker: "Person.NameFemale",
		},
		{
			name:  "",
			faker: "gender",
		},
		{
			name:  "",
			faker: "Person.Gender",
		},
		{
			name:  "",
			faker: "gender male",
		},
		{
			name:  "",
			faker: "Person.GenderMale",
		},
		{
			name:  "",
			faker: "gender female",
		},
		{
			name:  "",
			faker: "Person.GenderFemale",
		},
		{
			name:  "",
			faker: "UUID",
		},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got := faker.NewFaker().Faker(tc.faker)

			if len(got) == 0 {
				t.Error("faker by name is empty")
			}
		})
	}
}

func TestFaker_Numerify(t *testing.T) {
	for _, tc := range []struct {
		name string
		expr string
		in   string
	}{
		{
			name: "",
			expr: "",
			in:   "",
		},
		{
			name: "",
			expr: "[1-9][A-Z][1-9]",
			in:   "*A*",
		},
		{
			name: "",
			expr: "[A-Z][A-Z][A-Z]",
			in:   "AAA",
		},
		{
			name: "",
			expr: "[1-9][1-9][1-9][1-9][1-9]",
			in:   "*****",
		},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			r, err := regexp.Compile(tc.expr)
			if err != nil {
				t.Error(err)
			}

			n := faker.NewFaker().Numerify(tc.in)

			if !r.MatchString(n) {
				t.Errorf("%s not match %s", n, tc.expr)
			}
		})
	}
}
