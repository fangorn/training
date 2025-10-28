package duration_format

import "testing"

func TestFormatDuration(t *testing.T) {
	tests := []struct {
		seconds  int64
		expected string
	}{
		{62, "1 minute and 2 seconds"},
		{3662, "1 hour, 1 minute and 2 seconds"},
	}

	for _, test := range tests {
		result := FormatDuration(test.seconds)
		if result != test.expected {
			t.Errorf("FormatDuration(%d): expected %s, got %s", test.seconds, test.expected, result)
		}
	}
}
