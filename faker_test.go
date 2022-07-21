package faker_test

import (
	"fmt"
	"regexp"
	"testing"

	"go.neotoolkit.com/faker"
)

func TestInteger(t *testing.T) {
	t.Parallel()
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
		{
			name: "min -2, max -1",
			min:  -2,
			max:  -1,
		},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			value := faker.Integer(tc.min, tc.max)
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

func TestNumerify(t *testing.T) {
	t.Parallel()
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
			expr: "[0-9][A-Z][0-9]",
			in:   "#A#",
		},
		{
			name: "",
			expr: "[A-Z][A-Z][A-Z]",
			in:   "AAA",
		},
		{
			name: "",
			expr: "[0-9][0-9][0-9][0-9][0-9]",
			in:   "#####",
		},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			r, err := regexp.Compile(tc.expr)
			if err != nil {
				t.Error(err)
			}
			n := faker.Numerify(tc.in)
			if !r.MatchString(n) {
				t.Errorf("%s not match %s", n, tc.expr)
			}
		})
	}
}

func TestAsciify(t *testing.T) {
	t.Parallel()
	for _, tc := range []struct {
		name string
		expr string
		in   string
		opt  faker.Option
	}{
		{
			name: "",
			expr: "",
			in:   "",
			opt:  func(opts *faker.Config) {},
		},
		{
			name: "",
			expr: "[0-9][0-9][0-9]",
			in:   "111",
			opt:  func(opts *faker.Config) {},
		},
		{
			name: "",
			expr: "[a-zA-Z][a-zA-Z][a-zA-Z][a-zA-Z][a-zA-Z]",
			in:   "*****",
			opt:  func(opts *faker.Config) {},
		},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			r, err := regexp.Compile(tc.expr)
			if err != nil {
				t.Error(err)
			}
			a := faker.Asciify(tc.in, tc.opt)
			if !r.MatchString(a) {
				t.Errorf("%s not match %s", a, tc.expr)
			}
		})
	}
}
