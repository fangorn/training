package decode

import "testing"

func TestDecodeMorse(t *testing.T) {
	tests := []struct {
		code, expected string
	}{
		{".... . -.--   .--- ..- -.. .", "HEY JUDE"},
	}

	for _, test := range tests {
		result := DecodeMorse(test.code)
		if result != test.expected {
			t.Errorf("Code %s expected %s got %s", test.code, test.expected, result)
		}
	}
}
