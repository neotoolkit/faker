package faker_test

import (
	"fmt"
	"neotoolkit.com/faker"
	"testing"
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
