package deadfish

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	tests := []struct {
		data     string
		expected []int
	}{
		{data: "iiisdoso", expected: []int{8, 64}},
		{data: "iiisdosodddddiso", expected: []int{8, 64, 3600}},
		{data: "ooo", expected: []int{0, 0, 0}},
		{data: "ioioio", expected: []int{1, 2, 3}},
		{data: "codewars", expected: []int{0}},
	}

	for _, test := range tests {
		result := Parse(test.data)
		if reflect.DeepEqual(result, test.expected) == false {
			t.Errorf("Parse(\"%s\"): got %v; want %v", test.data, result, test.expected)
		}
	}
}
