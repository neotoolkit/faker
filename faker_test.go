package faker_test

import (
	"fmt"
	"testing"

	"github.com/neotoolkit/faker"
)

func TestIntBetween(t *testing.T) {
	f := faker.NewFaker()

	tests := []struct {
		name string
		min  int
		max  int
	}{
		{
			name: "",
			min:  1,
			max:  100,
		},
		{
			name: "",
			min:  1,
			max:  1,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
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
	tests := []struct {
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
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			f := faker.NewFaker()
			got := f.ByName(tc.faker)

			if nil == got {
				t.Fatal("faker by name is nil")
			}
		})
	}
}
