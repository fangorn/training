package int_to_ip

import "testing"

func TestInt32ToIp(t *testing.T) {
	tests := []struct {
		octet    uint32
		expected string
	}{
		{octet: 0, expected: "0.0.0.0"},
		{octet: 32, expected: "0.0.0.32"},
		{octet: 2149583361, expected: "128.32.10.1"},
	}

	for _, test := range tests {
		result := Int32ToIp(test.octet)
		if result != test.expected {
			t.Errorf("Int32ToIp(%d) = %s; want %s", test.octet, result, test.expected)
		}
	}
}
