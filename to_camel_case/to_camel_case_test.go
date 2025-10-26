package to_camel_case

import "testing"

func TestToCamelCase(t *testing.T) {
	tests := []struct {
		str, expected string
	}{
		{"the-stealth-warrior", "theStealthWarrior"},
		{"The_Stealth_Warrior", "TheStealthWarrior"},
		{"The_Stealth-Warrior", "TheStealthWarrior"},
	}

	for _, test := range tests {
		result := ToCamelCase(test.str)
		if result != test.expected {
			t.Errorf("ToCamelCase(%q) = %q, want %q", test.str, result, test.expected)
		}
	}
}
